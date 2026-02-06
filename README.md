# OnlyFans API - Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/ofauth-org/onlyfans-sdk-go.svg)](https://pkg.go.dev/github.com/ofauth-org/onlyfans-sdk-go)
[![License](https://img.shields.io/github/license/OFAuth-org/onlyfans-sdk-go)](LICENSE)

The official **OnlyFans API Go SDK** by [OFAuth](https://ofauth.com). A type-safe Go client for integrating with the OnlyFans API using only the standard library. Build OnlyFans tools, dashboards, analytics, and automations in Go.

> **What is this?** This is an SDK for the [OnlyFans API](https://ofauth.com) via OFAuth — the only authorized way to programmatically access OnlyFans data. Use it to build OnlyFans integrations, manage creator accounts, access earnings data, automate messaging, and more.

## Installation

```bash
go get github.com/ofauth-org/onlyfans-sdk-go
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    ofauth "github.com/ofauth-org/onlyfans-sdk-go"
)

func main() {
    client := ofauth.NewClient("your-api-key")

    ctx := context.Background()
    account, err := client.Whoami(ctx)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Account ID: %s\n", account.Id)
    fmt.Printf("Permissions: %v\n", account.Permissions)
}
```

## Features

- Full type safety with generated response structs (126 typed methods)
- Zero dependencies (standard library only)
- Builder pattern for configuration
- Context-based cancellation
- Typed API methods with response structs
- Webhook verification and routing (Svix-compatible)

## Configuration

```go
// Basic configuration
client := ofauth.NewClient("your-api-key")

// With default connection ID (for access API calls)
client := ofauth.NewClient("your-api-key").
    WithConnectionID("conn_xxx")

// Full configuration
client := ofauth.NewClient("your-api-key").
    WithBaseURL("https://api-next.ofauth.com").
    WithConnectionID("conn_xxx").
    WithHTTPClient(&http.Client{Timeout: 30 * time.Second})
```

## Usage Examples

### Account Operations

```go
ctx := context.Background()
client := ofauth.NewClient("your-api-key")

// Get account info
account, err := client.Whoami(ctx)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Account ID: %s\n", account.Id)

// List all connections
connections, err := client.ListAccountConnections(ctx, nil)
if err != nil {
    log.Fatal(err)
}
for _, conn := range connections.List {
    fmt.Printf("%s: %s (%s)\n", conn.Id, conn.UserData.Username, conn.Status)
}

// List with filters
connections, err := client.ListAccountConnections(ctx, &ofauth.RequestOptions{
    Query: url.Values{
        "status": []string{"active"},
        "limit":  []string{"10"},
    },
})
```

### Access API (OnlyFans Data)

Access endpoints require a connection ID:

```go
ctx := context.Background()

// Set connection ID on client
client := ofauth.NewClient("your-api-key").
    WithConnectionID("conn_xxx")

// Get creator profile
profile, err := client.GetSelf(ctx, nil)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Username: %s\n", profile.Username)

// Or pass connection ID per-request
profile, err := client.GetSelf(ctx, &ofauth.RequestOptions{
    ConnectionID: "conn_xxx",
})

// Get earnings data
earnings, err := client.GetEarningsChart(ctx, &ofauth.RequestOptions{
    Query: url.Values{
        "startDate": []string{"2024-01-01"},
        "endDate":   []string{"2024-01-31"},
    },
})

// List posts
posts, err := client.ListPosts(ctx, &ofauth.RequestOptions{
    Query: url.Values{"limit": []string{"20"}},
})
for _, post := range posts.List {
    fmt.Printf("Post %d: %s\n", post.Id, post.Text)
}
```

### Proxy Requests

Call any OnlyFans API endpoint directly:

```go
ctx := context.Background()
client := ofauth.NewClient("your-api-key")

// GET request
body, err := client.Request(ctx, "GET", "/v2/access/proxy/users/me", &ofauth.RequestOptions{
    ConnectionID: "conn_xxx",
})
if err != nil {
    log.Fatal(err)
}

var user map[string]interface{}
json.Unmarshal(body, &user)
fmt.Printf("User: %v\n", user["name"])

// POST request with body
body, err = client.Request(ctx, "POST", "/v2/access/proxy/messages/queue", &ofauth.RequestOptions{
    ConnectionID: "conn_xxx",
    Body: map[string]interface{}{
        "text":       "Hello!",
        "lockedText": false,
    },
})

// With query parameters
body, err = client.Request(ctx, "GET", "/v2/access/proxy/subscriptions/subscribers", &ofauth.RequestOptions{
    ConnectionID: "conn_xxx",
    Query: url.Values{
        "limit": []string{"10"},
    },
})
```

### Media Upload

```go
ctx := context.Background()
client := ofauth.NewClient("your-api-key")

// Read file
data, err := os.ReadFile("video.mp4")
if err != nil {
    log.Fatal(err)
}

// Upload
result, err := client.UploadMedia(ctx, "conn_xxx", "video.mp4", data, "video/mp4", nil)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Media ID: %s\n", result.MediaId)

// Upload with progress callback
result, err = client.UploadMedia(ctx, "conn_xxx", "video.mp4", data, "video/mp4", 
    func(uploaded, total int64) {
        fmt.Printf("Progress: %d/%d (%.1f%%)\n", uploaded, total, float64(uploaded)/float64(total)*100)
    },
)
```

## Error Handling

```go
import ofauth "github.com/ofauth-org/onlyfans-sdk-go"

ctx := context.Background()
client := ofauth.NewClient("your-api-key")

account, err := client.Whoami(ctx)
if err != nil {
    if apiErr, ok := err.(*ofauth.APIError); ok {
        fmt.Printf("API Error %d: %s\n", apiErr.Status, apiErr.Message)
        fmt.Printf("Error code: %s\n", apiErr.Code)
        fmt.Printf("Details: %v\n", apiErr.Details)
    } else {
        log.Fatal(err)
    }
}
```

## Type Safety with Response Structs

All API methods return typed response structs:

```go
// Response types are generated from OpenAPI spec
account, _ := client.Whoami(ctx)
// account is *ofauth.WhoamiResponse

// Full type safety
account.Id          // string
account.Name        // string
account.Permissions // []string

// Nested types
connections, _ := client.ListAccountConnections(ctx, nil)
// connections is *ofauth.ListAccountConnectionsResponse

for _, conn := range connections.List {
    conn.Id                  // string
    conn.Status              // string ("active", "expired", "awaiting_2fa")
    conn.UserData.Username   // string
    conn.UserData.Avatar     // string
}
```

## Available API Methods

### Account
- `Whoami(ctx) (*WhoamiResponse, error)`
- `ListAccountConnections(ctx, opts) (*ListAccountConnectionsResponse, error)`
- `DeleteAccountConnection(ctx, connectionId, opts) error`
- `InvalidateAccountConnection(ctx, connectionId, opts) error`
- `GetAccountConnectionsSettings(ctx, connectionId, opts) (*GetAccountConnectionsSettingsResponse, error)`
- `UpdateAccountConnectionsSettings(ctx, connectionId, opts) (*UpdateAccountConnectionsSettingsResponse, error)`

### Access (require connection ID)
- `GetSelf(ctx, opts) (*GetSelfResponse, error)`
- `ListPosts(ctx, opts) (*ListPostsResponse, error)`
- `CreatePost(ctx, opts) (*CreatePostResponse, error)`
- `GetEarningsChart(ctx, opts) (*GetEarningsChartResponse, error)`
- `ListChatsMessages(ctx, userId, opts) (*ListChatsMessagesResponse, error)`
- `ListSubscribers(ctx, opts) (*ListSubscribersResponse, error)`
- ... and 100+ more typed methods

### Low-Level
- `Request(ctx, method, path, opts) ([]byte, error)` - Raw API request

## License

MIT
