package ofauth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// UploadConfig holds upload configuration
type UploadConfig struct {
	ConnectionID string
	Filename     string
	MimeType     string
	VaultUpload  map[string]interface{}
	OnProgress   func(uploaded, total int64)
}

// UploadResult holds the upload result
type UploadResult struct {
	Media map[string]interface{} `json:"media"`
}

// UploadMedia uploads a media file (handles single/multi-part automatically)
func (c *Client) UploadMedia(ctx context.Context, file io.Reader, filesize int64, cfg *UploadConfig) (*UploadResult, error) {
	// Read file into memory
	fileData, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}
	
	// Initialize upload
	initBody := map[string]interface{}{
		"filename":  cfg.Filename,
		"filesize":  filesize,
		"mimeType":  cfg.MimeType,
	}
	if cfg.VaultUpload != nil {
		initBody["vaultUpload"] = cfg.VaultUpload
	}
	
	initBodyBytes, _ := json.Marshal(initBody)
	initReq, err := http.NewRequestWithContext(ctx, http.MethodPost,
		c.config.BaseURL+"/v2/access/uploads/init",
		bytes.NewReader(initBodyBytes))
	if err != nil {
		return nil, err
	}
	
	initReq.Header.Set("apiKey", c.config.APIKey)
	initReq.Header.Set("x-connection-id", cfg.ConnectionID)
	initReq.Header.Set("Content-Type", "application/json")
	
	initResp, err := c.config.HTTPClient.Do(initReq)
	if err != nil {
		return nil, err
	}
	defer initResp.Body.Close()
	
	if initResp.StatusCode >= 400 {
		body, _ := io.ReadAll(initResp.Body)
		return nil, &APIError{Status: initResp.StatusCode, Message: string(body)}
	}
	
	var initData struct {
		MediaUploadID string `json:"mediaUploadId"`
	}
	if err := json.NewDecoder(initResp.Body).Decode(&initData); err != nil {
		return nil, err
	}
	
	totalParts, _ := strconv.ParseInt(initResp.Header.Get("x-ofauth-upload-total-parts"), 10, 64)
	if totalParts == 0 {
		totalParts = 1
	}
	partSize, _ := strconv.ParseInt(initResp.Header.Get("x-ofauth-upload-part-size"), 10, 64)
	if partSize == 0 {
		partSize = filesize
	}
	
	// Single-part upload
	if totalParts == 1 {
		uploadReq, err := http.NewRequestWithContext(ctx, http.MethodPut,
			fmt.Sprintf("%s/v2/access/uploads/%s", c.config.BaseURL, initData.MediaUploadID),
			bytes.NewReader(fileData))
		if err != nil {
			return nil, err
		}
		
		uploadReq.Header.Set("apiKey", c.config.APIKey)
		uploadReq.Header.Set("x-connection-id", cfg.ConnectionID)
		uploadReq.Header.Set("Content-Type", cfg.MimeType)
		
		uploadResp, err := c.config.HTTPClient.Do(uploadReq)
		if err != nil {
			return nil, err
		}
		defer uploadResp.Body.Close()
		
		if uploadResp.StatusCode >= 400 {
			body, _ := io.ReadAll(uploadResp.Body)
			return nil, &APIError{Status: uploadResp.StatusCode, Message: string(body)}
		}
		
		if cfg.OnProgress != nil {
			cfg.OnProgress(filesize, filesize)
		}
		
		var result UploadResult
		if err := json.NewDecoder(uploadResp.Body).Decode(&result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	
	// Multi-part upload
	var uploaded int64 = 0
	for partNum := int64(1); partNum <= totalParts; partNum++ {
		start := (partNum - 1) * partSize
		end := start + partSize
		if end > filesize {
			end = filesize
		}
		chunk := fileData[start:end]
		
		partReq, err := http.NewRequestWithContext(ctx, http.MethodPut,
			fmt.Sprintf("%s/v2/access/uploads/%s/parts/%d", c.config.BaseURL, initData.MediaUploadID, partNum),
			bytes.NewReader(chunk))
		if err != nil {
			return nil, err
		}
		
		partReq.Header.Set("apiKey", c.config.APIKey)
		partReq.Header.Set("x-connection-id", cfg.ConnectionID)
		partReq.Header.Set("Content-Type", cfg.MimeType)
		
		partResp, err := c.config.HTTPClient.Do(partReq)
		if err != nil {
			return nil, err
		}
		partResp.Body.Close()
		
		if partResp.StatusCode >= 400 {
			return nil, &APIError{Status: partResp.StatusCode, Message: "chunk upload failed"}
		}
		
		uploaded += int64(len(chunk))
		if cfg.OnProgress != nil {
			cfg.OnProgress(uploaded, filesize)
		}
	}
	
	// Complete upload
	completeBody, _ := json.Marshal(map[string]string{"mediaUploadId": initData.MediaUploadID})
	completeReq, err := http.NewRequestWithContext(ctx, http.MethodPost,
		c.config.BaseURL+"/v2/access/uploads/complete",
		bytes.NewReader(completeBody))
	if err != nil {
		return nil, err
	}
	
	completeReq.Header.Set("apiKey", c.config.APIKey)
	completeReq.Header.Set("x-connection-id", cfg.ConnectionID)
	completeReq.Header.Set("Content-Type", "application/json")
	
	completeResp, err := c.config.HTTPClient.Do(completeReq)
	if err != nil {
		return nil, err
	}
	defer completeResp.Body.Close()
	
	if completeResp.StatusCode >= 400 {
		body, _ := io.ReadAll(completeResp.Body)
		return nil, &APIError{Status: completeResp.StatusCode, Message: string(body)}
	}
	
	var result UploadResult
	if err := json.NewDecoder(completeResp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
