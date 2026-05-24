# mosir_sdk_go

Go SDK for the Mosir public GraphQL API.

## Install

```bash
go get github.com/mosir-social/mosir_sdk_go
```

## Import

```go
import mosir "github.com/mosir-social/mosir_sdk_go"
```

## Quick start

Use the beta endpoint for testing and development:

```go
import (
	"context"
	"fmt"

	mosir "github.com/mosir-social/mosir_sdk_go"
)

client := mosir.NewClient("https://beta.mosir.app/api/v1", "", nil)
```

### Read public data

```go
import (
	"context"
	"fmt"

	mosir "github.com/mosir-social/mosir_sdk_go"
)

client := mosir.NewClient("https://beta.mosir.app/api/v1", "", nil)
post, err := client.GetPost(context.Background(), "VLO8u7UXqclQ7byjfMEX0")
if err != nil {
	panic(err)
}

fmt.Println(post.GetPost.Content)
```

### Authenticated requests

```go
import (
	"context"
	"fmt"

	mosir "github.com/mosir-social/mosir_sdk_go"
)

client := mosir.NewClient("https://beta.mosir.app/api/v1", "YOUR_TOKEN", nil)
count, err := client.GetUnreadNotificationCount(context.Background())
if err != nil {
	panic(err)
}

fmt.Println(count.GetUnreadNotificationCount)
```

## Media helpers

Fetch the best media file for a media object:

```go
import (
	"context"
	"fmt"

	mosir "github.com/mosir-social/mosir_sdk_go"
)

client := mosir.NewClient("https://beta.mosir.app/api/v1", "", nil)
bytes, err := client.FetchMedia(context.Background(), media, nil)
if err != nil {
	panic(err)
}
fmt.Println(len(bytes))
```

Fetch preview images for posts, profiles, and collections:

```go
import (
	"context"
	"fmt"

	mosir "github.com/mosir-social/mosir_sdk_go"
)

client := mosir.NewClient("https://beta.mosir.app/api/v1", "", nil)
url := mosir.GetPreviewImageUrl("https://beta.mosir.app/api/v1", "post", "VLO8u7UXqclQ7byjfMEX0")
preview, err := client.FetchPreviewImage(context.Background(), "post", "VLO8u7UXqclQ7byjfMEX0", nil)
if err != nil {
	panic(err)
}
_ = url
_ = preview
```

## SSE subscriptions

```go
import (
	"context"
	"fmt"

	"github.com/tmaxmax/go-sse"
	mosir "github.com/mosir-social/mosir_sdk_go"
)

client := mosir.NewClient("https://beta.mosir.app/api/v1", "", nil)
err := client.SubscribeSSE(
	context.Background(),
	"https://beta.mosir.app/api/v1/sse/postCreatedByAuthor?authorId=...",
	func(event *sse.Event) error {
		fmt.Println(event.Data)
		return nil
	},
)
if err != nil {
	panic(err)
}
```

Notes:
- keep reconnect logic in your app
- public subscriptions do not require a token
- Mosir SSE connections are limited to 1 hour

## Development

```bash
go test ./...
MOSIR_SMOKE=1 go test -run SmokeBetaEndpoint -v ./...
task codegen
```

`task codegen` uses the vendored local genqlient fork in `tools/genqlient/`.

## Layout

- `client.go` — thin client wrapper
- `client_generated.go` — generated operation wrappers, committed to git
- `aliases.go` — small re-exports for generated input types
- `helpers.go` — media and preview helpers
- `sse.go` — thin SSE wrapper
- `internal/generated/` — generated GraphQL types and operations
- `tools/genqlient/` — local codegen fork
- `tools/genwrappers/` — wrapper generator for the root package
