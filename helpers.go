package mosir_sdk_go

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// FetchOptions contains options for fetching media or preview images.
type FetchOptions struct {
	Header    http.Header
	Transport http.RoundTripper
	Endpoint  string
}

// GetPreviewImageUrl returns the full URL for a preview image.
func GetPreviewImageUrl(endpoint, kind, id string) string {
	var route string
	switch kind {
	case "post":
		route = "postopengraph"
	case "profile":
		route = "profileopengraph"
	case "postCollection":
		route = "collectionopengraph"
	default:
		return ""
	}
	return fmt.Sprintf("%s/%s/%s", endpoint, route, id)
}

// selectMediaFile selects the best media file from the given media metadata.
func selectMediaFile(media *MediaMetadata, profile *MediaFileProfile) *MediaFileMetadata {
	if media == nil || len(media.Files) == 0 {
		return nil
	}

	if profile != nil {
		for i := range media.Files {
			if media.Files[i].Profile == *profile {
				return &media.Files[i]
			}
		}
		return nil
	}

	fallbacks := []MediaFileProfile{
		MediaFileProfileQuality,
		MediaFileProfileCompatible,
		MediaFileProfileThumbnail,
		MediaFileProfileAnimatedCompatible,
	}
	for _, fallback := range fallbacks {
		for i := range media.Files {
			if media.Files[i].Profile == fallback {
				return &media.Files[i]
			}
		}
	}

	return &media.Files[0]
}

func fetchURL(ctx context.Context, baseClient *http.Client, url string, opts *FetchOptions) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	if opts != nil && opts.Header != nil {
		for k, v := range opts.Header {
			req.Header[k] = append([]string(nil), v...)
		}
	}

	client := cloneHTTPClient(baseClient)
	if opts != nil && opts.Transport != nil {
		client.Transport = opts.Transport
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

// FetchMedia fetches the media file from the given media metadata.
func (c *Client) FetchMedia(ctx context.Context, media *MediaMetadata, opts *FetchOptions) ([]byte, error) {
	file := selectMediaFile(media, nil)
	if file == nil {
		return nil, fmt.Errorf("No media file is available for the requested media object.")
	}
	return fetchURL(ctx, c.httpClient, file.URL, opts)
}

// FetchPreviewImage fetches the preview image.
func (c *Client) FetchPreviewImage(ctx context.Context, kind, id string, opts *FetchOptions) ([]byte, error) {
	endpoint := c.endpoint
	if opts != nil && opts.Endpoint != "" {
		endpoint = opts.Endpoint
	}
	return fetchURL(ctx, c.httpClient, GetPreviewImageUrl(endpoint, kind, id), opts)
}
