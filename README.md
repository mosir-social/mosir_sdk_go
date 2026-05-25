# mosir-sdk-go

Go SDK for the Mosir public GraphQL API.

## What this SDK provides

- generated Go types and operation wrappers from `public.graphqls` and `public.operations.graphql`
- optional Bearer token auth
- GraphQL SSE subscription helpers on the main `Client`
- media and preview image helper utilities
- raw HTTP access via `client.HTTPClient()` when direct control is needed

## Transport choice

This SDK uses:

- `genqlient` for queries and mutations
- `go-sse` for subscriptions

This keeps the package small while still supporting the preferred subscription transport.
WebSocket support is intentionally not bundled.

## Install

```bash
go get github.com/mosir-social/mosir_sdk_go
```

## Quick start

### Anonymous/public requests

Only public data needs no token.

```go
package main

import (
	"context"
	"fmt"

	mosir "github.com/mosir-social/mosir_sdk_go"
)

func main() {
	client := mosir.NewClient("https://beta.mosir.app/api/v1", "", nil)

	post, err := client.GetPost(context.Background(), "VLO8u7UXqclQ7byjfMEX0")
	if err != nil {
		panic(err)
	}

	fmt.Println(post.GetPost.Content)
}
```

### Authenticated requests

Use a token for authenticated operations such as notifications.

```go
client := mosir.NewClient("https://beta.mosir.app/api/v1", os.Getenv("MOSIR_API_TOKEN"), nil)

notifications, err := client.GetNotifications(context.Background(), "", mosir.NotificationFilterInput{}, 20)
if err != nil {
	panic(err)
}
fmt.Println(notifications.GetNotifications.Edges)
```

## Custom endpoint

```go
client := mosir.NewClient("https://example.com/api/v1", os.Getenv("MOSIR_API_TOKEN"), nil)
```

## Common usage examples

### Get a post

```go
post, err := client.GetPost(context.Background(), "VLO8u7UXqclQ7byjfMEX0")
if err != nil {
	panic(err)
}

fmt.Println(post.GetPost.Author.Username)
fmt.Println(post.GetPost.Content)
```

### Get replies under a post

Replies are exposed as nested GraphQL fields on `Post`, so this is a good case for direct GraphQL usage:

```go
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	mosir "github.com/mosir-social/mosir_sdk_go"
)

func main() {
	client := mosir.NewClient("https://beta.mosir.app/api/v1", "", nil)

	payload := map[string]any{
		"query": `query GetPostReplies($postId: ID!, $limit: Int) {
  getPost(postId: $postId) {
    id
    commentsRecent(limit: $limit) {
      edges {
        id
        content
        createdAt
        author {
          id
          username
          displayName
        }
      }
      pageInfo {
        endCursor
        hasNextPage
        totalCount
      }
    }
  }
}`,
		"variables": map[string]any{
			"postId": "VLO8u7UXqclQ7byjfMEX0",
			"limit":  3,
		},
	}

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "https://beta.mosir.app/api/v1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.HTTPClient().Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var out map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		panic(err)
	}

	fmt.Println(out["data"])
}
```

### Get notifications

```go
notifications, err := client.GetNotifications(context.Background(), "", mosir.NotificationFilterInput{}, 20)
if err != nil {
	panic(err)
}
fmt.Println(notifications.GetNotifications.Edges)
```

### Fetch media bytes from a `Media` result

```go
var media *mosir.MediaMetadata

bytes, err := client.FetchMedia(context.Background(), media, nil)
if err != nil {
	panic(err)
}
fmt.Println(len(bytes))
```

### Fetch preview image for a post, profile, or collection

```go
previewURL := mosir.GetPreviewImageUrl("https://beta.mosir.app/api/v1", "post", "VLO8u7UXqclQ7byjfMEX0")
fmt.Println(previewURL)

previewBytes, err := client.FetchPreviewImage(context.Background(), "post", "VLO8u7UXqclQ7byjfMEX0", nil)
if err != nil {
	panic(err)
}
fmt.Println(len(previewBytes))
```

All generated operations are available directly on `client` as methods (for example, `client.GetCurrentAccount(...)`).

## SSE subscriptions

Subscriptions let your app receive updates from Mosir in near real time without polling.
This SDK uses **SSE** (Server-Sent Events) for subscriptions by default.

A good example is a Discord bot:
- subscribe to `PostCreatedByAuthor`
- when a creator publishes something new, format it
- send a message into a Discord channel

That way the bot reacts as soon as something changes, instead of repeatedly calling the API every few seconds.
SSE is especially useful for backend workers, bots, notification relays, and other long-running processes that want a simple one-way stream of events from the server.
For public subscriptions like `PostCreatedByAuthor`, a token is not required.

Note: each SSE connection lasts at most 1 hour. In practice, network conditions may cause it to end earlier.
If you build a bot, worker, or relay process, make sure you implement reconnect logic.

```go
ctx := context.Background()
client := mosir.NewClient("https://beta.mosir.app/api/v1", "", nil)

profile, err := client.GetAccountProfile(ctx, "", "leemiyinghao")
if err != nil {
	panic(err)
}
authorID := profile.GetAccountProfile.Id

err = client.PostCreatedByAuthor(ctx, authorID, mosir.PostTypePost, func(event mosir.PostCreatedByAuthorWsResponse) error {
	if event.Data == nil {
		return nil
	}
	fmt.Println(event.Data.PostCreatedByAuthor.Id)
	fmt.Println(event.Data.PostCreatedByAuthor.Content)
	return nil
})
if err != nil {
	panic(err)
}
```

## Raw GraphQL access

Authentication is optional. Pass a token when creating `Client` for authenticated operations, or omit it when accessing only public data.

### Typed operation usage

Use generated client methods directly:

```go
data, err := client.GetNotifications(context.Background(), "", mosir.NotificationFilterInput{}, 20)
```

### Raw GraphQL string usage

Use `client.HTTPClient()` to send your own GraphQL request payload when needed.

## WebSocket usage

WebSocket transport is not bundled.
If you want it, use your own GraphQL WebSocket client against the same endpoint.

## Notes

- endpoint is passed explicitly into `NewClient(...)` (examples use `https://beta.mosir.app/api/v1`)
- token is optional for public data and required only for authenticated operations
- the same applies to subscriptions: public subscription data does not require a token
- media helpers are available through `FetchMedia(...)`
- preview image helpers are available through `GetPreviewImageUrl(...)` and `FetchPreviewImage(...)`
- subscriptions use SSE in this SDK
- direct GraphQL usage is available through generated methods and manual HTTP requests

## Development

### Generate code

```bash
task codegen
```

### Test

```bash
task test
```

### Tidy dependencies

```bash
task tidy
```

## Repo artifacts

- `public.graphqls` — copied public schema artifact
- `public.operations.graphql` — copied curated operation document
- `internal/generated/` — generated GraphQL types and operations
- `client_generated.go` — generated wrapper methods on `Client`
- `sse.go` — GraphQL SSE subscription wrappers

## License

This project is licensed under the GNU Lesser General Public License v3.0 (LGPL-3.0).
See [`LICENSE`](./LICENSE) for details.
