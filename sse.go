package mosir_sdk_go

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Khan/genqlient/graphql"
	gosse "github.com/tmaxmax/go-sse"
)

const graphqlSSEAccept = "text/event-stream"

// subscribeGraphQLSSE opens a graphql-sse subscription against the client's endpoint.
func (c *Client) subscribeGraphQLSSE(ctx context.Context, operationName string, variables any, onMessage func(raw json.RawMessage) error) error {
	if c == nil {
		return fmt.Errorf("client is nil")
	}

	payload := struct {
		Query         string `json:"query"`
		OperationName string `json:"operationName"`
		Variables     any    `json:"variables,omitempty"`
	}{
		Query:         publicOperationsDocument,
		OperationName: operationName,
		Variables:     variables,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	client := cloneHTTPClient(c.httpClient)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.endpoint, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Accept", graphqlSSEAccept)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cache-Control", "no-cache")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("graphql-sse subscription failed with status %d", resp.StatusCode)
	}

	var finalErr error
	parser := gosse.Read(resp.Body, nil)
	parser(func(event gosse.Event, parseErr error) bool {
		if parseErr != nil {
			finalErr = parseErr
			return false
		}

		switch event.Type {
		case "complete":
			return false
		case "error":
			if event.Data != "" {
				finalErr = fmt.Errorf("%s", event.Data)
			} else {
				finalErr = fmt.Errorf("graphql-sse subscription error")
			}
			return false
		default:
			if event.Data == "" {
				return true
			}
			if err := onMessage(json.RawMessage(event.Data)); err != nil {
				finalErr = err
				return false
			}
			return true
		}
	})

	if finalErr != nil {
		return finalErr
	}
	if ctx.Err() != nil {
		return ctx.Err()
	}
	return nil
}

func decodeSubscriptionMessage[T any](raw json.RawMessage, out *T) error {
	var gqlResp graphql.Response
	if err := json.Unmarshal(raw, &gqlResp); err != nil {
		return err
	}
	if len(gqlResp.Errors) == 0 {
		if err := json.Unmarshal(raw, out); err != nil {
			return err
		}
		return nil
	}
	return json.Unmarshal(raw, out)
}

// NotificationReceived subscribes to authenticated notification events.
func (c *Client) NotificationReceived(ctx context.Context, handler func(NotificationReceivedWsResponse) error) error {
	return c.subscribeGraphQLSSE(ctx, "NotificationReceived", nil, func(raw json.RawMessage) error {
		var resp NotificationReceivedWsResponse
		if err := decodeSubscriptionMessage(raw, &resp); err != nil {
			return err
		}
		return handler(resp)
	})
}

// PostCreatedByAuthor subscribes to posts created by a specific author.
func (c *Client) PostCreatedByAuthor(ctx context.Context, authorId string, postType PostType, handler func(PostCreatedByAuthorWsResponse) error) error {
	return c.subscribeGraphQLSSE(ctx, "PostCreatedByAuthor", map[string]any{
		"authorId": authorId,
		"postType": postType,
	}, func(raw json.RawMessage) error {
		var resp PostCreatedByAuthorWsResponse
		if err := decodeSubscriptionMessage(raw, &resp); err != nil {
			return err
		}
		return handler(resp)
	})
}

// PostCreatedInCollection subscribes to posts created in a post collection.
func (c *Client) PostCreatedInCollection(ctx context.Context, postCollectionID string, handler func(PostCreatedInCollectionWsResponse) error) error {
	return c.subscribeGraphQLSSE(ctx, "PostCreatedInCollection", map[string]any{
		"postCollectionID": postCollectionID,
	}, func(raw json.RawMessage) error {
		var resp PostCreatedInCollectionWsResponse
		if err := decodeSubscriptionMessage(raw, &resp); err != nil {
			return err
		}
		return handler(resp)
	})
}

// PostDeleted subscribes to deletion events for a post.
func (c *Client) PostDeleted(ctx context.Context, postId string, handler func(PostDeletedWsResponse) error) error {
	return c.subscribeGraphQLSSE(ctx, "PostDeleted", map[string]any{
		"postId": postId,
	}, func(raw json.RawMessage) error {
		var resp PostDeletedWsResponse
		if err := decodeSubscriptionMessage(raw, &resp); err != nil {
			return err
		}
		return handler(resp)
	})
}

// PostUpdated subscribes to updates for a post.
func (c *Client) PostUpdated(ctx context.Context, postId string, handler func(PostUpdatedWsResponse) error) error {
	return c.subscribeGraphQLSSE(ctx, "PostUpdated", map[string]any{
		"postId": postId,
	}, func(raw json.RawMessage) error {
		var resp PostUpdatedWsResponse
		if err := decodeSubscriptionMessage(raw, &resp); err != nil {
			return err
		}
		return handler(resp)
	})
}

// ReplyCreatedUnderRootPost subscribes to replies under a root post.
func (c *Client) ReplyCreatedUnderRootPost(ctx context.Context, rootPostId string, handler func(ReplyCreatedUnderRootPostWsResponse) error) error {
	return c.subscribeGraphQLSSE(ctx, "ReplyCreatedUnderRootPost", map[string]any{
		"rootPostId": rootPostId,
	}, func(raw json.RawMessage) error {
		var resp ReplyCreatedUnderRootPostWsResponse
		if err := decodeSubscriptionMessage(raw, &resp); err != nil {
			return err
		}
		return handler(resp)
	})
}

// publicOperationsDocument is the full GraphQL operations document.
// It is embedded as-is so that subscription wrappers can address any named subscription.
