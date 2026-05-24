package mosir_sdk_go

import (
	"context"
	"fmt"
	"net/http"

	"github.com/tmaxmax/go-sse"
)

// SSEClient is a thin wrapper around go-sse for event stream consumption.
type SSEClient struct {
	client *sse.Client
}

// NewSSEClient creates a new SSE client.
func NewSSEClient(token string, httpClient *http.Client) *SSEClient {
	cloned := cloneHTTPClient(httpClient)
	if token != "" {
		next := cloned.Transport
		if next == nil {
			next = http.DefaultTransport
		}
		cloned.Transport = &tokenRoundTripper{next: next, token: token}
	}

	return &SSEClient{client: &sse.Client{HTTPClient: cloned}}
}

// Subscribe connects to an SSE stream and invokes handler for each event.
func (s *SSEClient) Subscribe(ctx context.Context, url string, handler func(event *sse.Event) error) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	conn := s.client.NewConnection(req)
	var handlerErr error
	conn.SubscribeToAll(func(ev sse.Event) {
		if handlerErr != nil {
			return
		}
		if err := handler(&ev); err != nil {
			handlerErr = err
		}
	})

	if err := conn.Connect(); err != nil {
		return err
	}
	return handlerErr
}

// SubscribeSSE is a convenience wrapper for the Mosir client.
func (c *Client) SubscribeSSE(ctx context.Context, url string, handler func(event *sse.Event) error) error {
	if c == nil {
		return fmt.Errorf("client is nil")
	}
	return NewSSEClient(c.token, c.httpClient).Subscribe(ctx, url, handler)
}
