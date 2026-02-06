package ofauth

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// ============================================================================
// Types
// ============================================================================

// WebhookVerificationError is returned when webhook signature verification fails.
type WebhookVerificationError struct {
	Message string
	Code    string
}

func (e *WebhookVerificationError) Error() string {
	return fmt.Sprintf("webhook verification error [%s]: %s", e.Code, e.Message)
}

// IsWebhookVerificationError checks if an error is a WebhookVerificationError.
func IsWebhookVerificationError(err error) bool {
	_, ok := err.(*WebhookVerificationError)
	return ok
}

// WebhookHeaders holds the extracted Svix headers.
type WebhookHeaders struct {
	SvixID        string
	SvixTimestamp string
	SvixSignature string
}

// WebhookUserData represents user data from OnlyFans.
type WebhookUserData struct {
	UserID   string `json:"userId"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

// WebhookConnection represents a connection object in webhook events.
type WebhookConnection struct {
	ID             string           `json:"id"`
	PlatformUserID string           `json:"platformUserId"`
	Status         string           `json:"status"`
	UserData       *WebhookUserData `json:"userData,omitempty"`
	Permissions    []string         `json:"permissions,omitempty"`
}

// WebhookDynamicRules represents dynamic rules configuration.
type WebhookDynamicRules struct {
	StaticParam      string `json:"static_param"`
	Format           string `json:"format"`
	Start            string `json:"start"`
	End              string `json:"end"`
	Prefix           string `json:"prefix"`
	Suffix           string `json:"suffix"`
	ChecksumConstant int    `json:"checksum_constant"`
	ChecksumIndexes  []int  `json:"checksum_indexes"`
	AppToken         string `json:"app_token"`
	Revision         string `json:"revision"`
}

// WebhookEvent represents a raw webhook event with unparsed data.
type WebhookEvent struct {
	EventType string          `json:"eventType"`
	Live      bool            `json:"live"`
	Data      json.RawMessage `json:"data"`
}

// ConnectionCreatedEventData is the data payload for connection.created events.
type ConnectionCreatedEventData struct {
	Connection        WebhookConnection `json:"connection"`
	ClientReferenceID *string           `json:"clientReferenceId"`
}

// ConnectionUpdatedEventData is the data payload for connection.updated events.
type ConnectionUpdatedEventData struct {
	Connection        WebhookConnection `json:"connection"`
	ClientReferenceID *string           `json:"clientReferenceId"`
}

// ConnectionExpiredEventData is the data payload for connection.expired events.
type ConnectionExpiredEventData struct {
	Connection struct {
		ID             string `json:"id"`
		PlatformUserID string `json:"platformUserId"`
		Status         string `json:"status"`
	} `json:"connection"`
	ClientReferenceID *string `json:"clientReferenceId"`
}

// RulesUpdatedEventData is the data payload for rules.updated events.
type RulesUpdatedEventData struct {
	Rules    WebhookDynamicRules `json:"rules"`
	Revision string              `json:"revision"`
}

// ============================================================================
// Verification
// ============================================================================

// ExtractWebhookHeaders extracts and validates Svix webhook headers from an HTTP request.
func ExtractWebhookHeaders(r *http.Request) (*WebhookHeaders, error) {
	svixID := r.Header.Get("svix-id")
	svixTimestamp := r.Header.Get("svix-timestamp")
	svixSignature := r.Header.Get("svix-signature")

	if svixID == "" || svixTimestamp == "" || svixSignature == "" {
		return nil, &WebhookVerificationError{
			Message: "Missing required webhook headers (svix-id, svix-timestamp, svix-signature)",
			Code:    "MISSING_HEADERS",
		}
	}

	return &WebhookHeaders{
		SvixID:        svixID,
		SvixTimestamp: svixTimestamp,
		SvixSignature: svixSignature,
	}, nil
}

// ExtractWebhookHeadersFromMap extracts headers from a string map (useful for testing).
func ExtractWebhookHeadersFromMap(headers map[string]string) (*WebhookHeaders, error) {
	get := func(name string) string {
		if v, ok := headers[name]; ok {
			return v
		}
		return headers[strings.ToLower(name)]
	}

	svixID := get("svix-id")
	svixTimestamp := get("svix-timestamp")
	svixSignature := get("svix-signature")

	if svixID == "" || svixTimestamp == "" || svixSignature == "" {
		return nil, &WebhookVerificationError{
			Message: "Missing required webhook headers (svix-id, svix-timestamp, svix-signature)",
			Code:    "MISSING_HEADERS",
		}
	}

	return &WebhookHeaders{
		SvixID:        svixID,
		SvixTimestamp: svixTimestamp,
		SvixSignature: svixSignature,
	}, nil
}

// VerifyWebhookSignature verifies a webhook signature using the Svix HMAC-SHA256 protocol.
// tolerance is the maximum allowed timestamp age in seconds (use 0 for default of 300).
func VerifyWebhookSignature(payload []byte, headers *WebhookHeaders, secret string, tolerance int) error {
	if tolerance <= 0 {
		tolerance = 300
	}

	// Timestamp check
	ts, err := strconv.ParseInt(headers.SvixTimestamp, 10, 64)
	if err != nil {
		return &WebhookVerificationError{Message: "Invalid timestamp format", Code: "INVALID_TIMESTAMP"}
	}
	now := time.Now().Unix()
	if math.Abs(float64(now-ts)) > float64(tolerance) {
		return &WebhookVerificationError{Message: "Webhook timestamp too old", Code: "TIMESTAMP_TOO_OLD"}
	}

	// Compute expected signature
	cleanSecret := secret
	if strings.HasPrefix(secret, "whsec_") {
		cleanSecret = secret[6:]
	}
	key, err := base64.StdEncoding.DecodeString(cleanSecret)
	if err != nil {
		return &WebhookVerificationError{Message: "Invalid webhook secret", Code: "INVALID_SECRET"}
	}

	signed := fmt.Sprintf("%s.%s.%s", headers.SvixID, headers.SvixTimestamp, string(payload))
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(signed))
	expected := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	// Parse provided signatures
	var validSigs []string
	for _, part := range strings.Split(headers.SvixSignature, " ") {
		pieces := strings.SplitN(part, ",", 2)
		if len(pieces) == 2 && pieces[0] == "v1" {
			validSigs = append(validSigs, pieces[1])
		}
	}

	if len(validSigs) == 0 {
		return &WebhookVerificationError{Message: "No valid v1 signatures found", Code: "NO_VALID_SIGNATURES"}
	}

	// Constant-time comparison
	expectedBytes := []byte(expected)
	for _, sig := range validSigs {
		if subtle.ConstantTimeCompare(expectedBytes, []byte(sig)) == 1 {
			return nil
		}
	}

	return &WebhookVerificationError{Message: "Signature verification failed", Code: "SIGNATURE_MISMATCH"}
}

// VerifyWebhookPayload verifies a webhook and returns the parsed event.
func VerifyWebhookPayload(payload []byte, headers *WebhookHeaders, secret string, tolerance int) (*WebhookEvent, error) {
	if err := VerifyWebhookSignature(payload, headers, secret, tolerance); err != nil {
		return nil, err
	}

	var event WebhookEvent
	if err := json.Unmarshal(payload, &event); err != nil {
		return nil, &WebhookVerificationError{Message: "Failed to parse webhook payload as JSON", Code: "INVALID_JSON"}
	}

	return &event, nil
}

// ============================================================================
// Router
// ============================================================================

// WebhookHandlerFunc handles a verified webhook event.
type WebhookHandlerFunc func(event *WebhookEvent) error

// WebhookErrorHandlerFunc handles errors during webhook processing.
type WebhookErrorHandlerFunc func(err error, event *WebhookEvent)

// RouterOption configures a WebhookRouter.
type RouterOption func(*WebhookRouter)

// WithTolerance sets the timestamp tolerance in seconds.
func WithTolerance(seconds int) RouterOption {
	return func(r *WebhookRouter) {
		r.tolerance = seconds
	}
}

// WebhookRouter routes verified webhook events to registered handlers.
type WebhookRouter struct {
	secret         string
	tolerance      int
	handlers       map[string]WebhookHandlerFunc
	defaultHandler WebhookHandlerFunc
	errorHandler   WebhookErrorHandlerFunc
}

// NewWebhookRouter creates a new WebhookRouter.
func NewWebhookRouter(secret string, opts ...RouterOption) *WebhookRouter {
	r := &WebhookRouter{
		secret:    secret,
		tolerance: 300,
		handlers:  make(map[string]WebhookHandlerFunc),
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

// On registers a handler for a specific event type.
func (r *WebhookRouter) On(eventType string, handler WebhookHandlerFunc) *WebhookRouter {
	r.handlers[eventType] = handler
	return r
}

// OnDefault registers a default handler for unmatched event types.
func (r *WebhookRouter) OnDefault(handler WebhookHandlerFunc) *WebhookRouter {
	r.defaultHandler = handler
	return r
}

// OnError registers an error handler.
func (r *WebhookRouter) OnError(handler WebhookErrorHandlerFunc) *WebhookRouter {
	r.errorHandler = handler
	return r
}

// HandleRequest verifies and routes a webhook from an HTTP request.
func (r *WebhookRouter) HandleRequest(w http.ResponseWriter, req *http.Request) {
	headers, err := ExtractWebhookHeaders(req)
	if err != nil {
		r.respondError(w, err)
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		r.respondError(w, &WebhookVerificationError{Message: "Failed to read request body", Code: "READ_ERROR"})
		return
	}

	event, err := VerifyWebhookPayload(body, headers, r.secret, r.tolerance)
	if err != nil {
		r.respondError(w, err)
		return
	}

	if procErr := r.processEvent(event); procErr != nil {
		if r.errorHandler != nil {
			r.errorHandler(procErr, event)
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// ServeHTTP implements the http.Handler interface.
func (r *WebhookRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.HandleRequest(w, req)
}

// UpdateSecret updates the webhook secret at runtime.
func (r *WebhookRouter) UpdateSecret(secret string) {
	r.secret = secret
}

// UpdateTolerance updates the timestamp tolerance at runtime.
func (r *WebhookRouter) UpdateTolerance(tolerance int) {
	r.tolerance = tolerance
}

func (r *WebhookRouter) processEvent(event *WebhookEvent) error {
	handler, ok := r.handlers[event.EventType]
	if !ok {
		handler = r.defaultHandler
	}
	if handler == nil {
		return nil
	}

	if err := handler(event); err != nil {
		if r.errorHandler != nil {
			r.errorHandler(err, event)
			return nil // Error was handled
		}
		return err
	}
	return nil
}

func (r *WebhookRouter) respondError(w http.ResponseWriter, err error) {
	if r.errorHandler != nil {
		r.errorHandler(err, nil)
	}

	if wvErr, ok := err.(*WebhookVerificationError); ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": wvErr.Message,
			"code":  wvErr.Code,
		})
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Internal server error"))
}

// ParseEventData is a helper to unmarshal event.Data into a typed struct.
//
//	var data ConnectionCreatedEventData
//	if err := ParseEventData(event, &data); err != nil { ... }
func ParseEventData(event *WebhookEvent, target interface{}) error {
	return json.Unmarshal(event.Data, target)
}
