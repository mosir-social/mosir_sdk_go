package mosir_sdk_go

import (
	"net/http"

	"github.com/Khan/genqlient/graphql"
)

// Client is the main entry point for the Mosir SDK.
type Client struct {
	httpClient *http.Client
	endpoint   string
	token      string
}

// NewClient creates a new Mosir client.
func NewClient(endpoint, token string, httpClient *http.Client) *Client {
	client := cloneHTTPClient(httpClient)
	if token != "" {
		wrapTokenTransport(client, token)
	}
	return &Client{
		httpClient: client,
		endpoint:   endpoint,
		token:      token,
	}
}

// HTTPClient exposes the underlying HTTP client.
func (c *Client) HTTPClient() *http.Client {
	if c == nil {
		return nil
	}
	return c.httpClient
}

func (c *Client) graphqlClient() graphql.Client {
	if c == nil {
		return graphql.NewClient("", nil)
	}
	return graphql.NewClient(c.endpoint, c.httpClient)
}

// cloneHTTPClient preserves all client settings.
func cloneHTTPClient(src *http.Client) *http.Client {
	if src == nil {
		src = http.DefaultClient
	}
	clone := *src
	return &clone
}

// wrapTokenTransport adds a bearer token transport to the client.
func wrapTokenTransport(client *http.Client, token string) {
	if client == nil || token == "" {
		return
	}
	next := client.Transport
	if next == nil {
		next = http.DefaultTransport
	}
	client.Transport = &tokenRoundTripper{next: next, token: token}
}

// tokenRoundTripper is a RoundTripper that adds the Authorization header.
type tokenRoundTripper struct {
	next  http.RoundTripper
	token string
}

func (t *tokenRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.token != "" {
		req.Header.Set("Authorization", "Bearer "+t.token)
	}
	return t.next.RoundTrip(req)
}

// WithAuth returns a new client with authentication enabled.
// It clones the current http.Client and only swaps the transport.
func (c *Client) WithAuth() *Client {
	if c == nil {
		return nil
	}

	clone := cloneHTTPClient(c.httpClient)
	wrapTokenTransport(clone, c.token)

	return &Client{
		httpClient: clone,
		endpoint:   c.endpoint,
		token:      c.token,
	}
}
