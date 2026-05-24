package mosir_sdk_go

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSelectMediaFile(t *testing.T) {
	media := &MediaMetadata{
		Files: []MediaFileMetadata{
			{ID: "1", Profile: MediaFileProfileThumbnail, URL: "url1"},
			{ID: "2", Profile: MediaFileProfileQuality, URL: "url2"},
			{ID: "3", Profile: MediaFileProfileCompatible, URL: "url3"},
			{ID: "4", Profile: MediaFileProfileAnimatedCompatible, URL: "url4"},
		},
	}

	t.Run("no files", func(t *testing.T) {
		if res := selectMediaFile(nil, nil); res != nil {
			t.Error("expected nil for nil media")
		}
		if res := selectMediaFile(&MediaMetadata{}, nil); res != nil {
			t.Error("expected nil for empty files")
		}
	})

	t.Run("exact profile match", func(t *testing.T) {
		profile := MediaFileProfileQuality
		res := selectMediaFile(media, &profile)
		if res == nil || res.ID != "2" {
			t.Errorf("expected ID 2, got %v", res)
		}
	})

	t.Run("fallback order", func(t *testing.T) {
		res := selectMediaFile(media, nil)
		if res == nil || res.ID != "2" {
			t.Errorf("expected ID 2 (Quality), got %v", res)
		}

		media2 := &MediaMetadata{
			Files: []MediaFileMetadata{
				{ID: "thumb", Profile: MediaFileProfileThumbnail, URL: "url1"},
				{ID: "comp", Profile: MediaFileProfileCompatible, URL: "url2"},
			},
		}
		res2 := selectMediaFile(media2, nil)
		if res2 == nil || res2.ID != "comp" {
			t.Errorf("expected ID comp, got %v", res2)
		}
	})

	t.Run("fallback to first file", func(t *testing.T) {
		media3 := &MediaMetadata{
			Files: []MediaFileMetadata{{ID: "last", Profile: "UNKNOWN", URL: "url1"}},
		}
		res := selectMediaFile(media3, nil)
		if res == nil || res.ID != "last" {
			t.Errorf("expected ID last, got %v", res)
		}
	})
}

func TestGetPreviewImageUrl(t *testing.T) {
	tests := []struct {
		kind     string
		id       string
		expected string
	}{
		{"post", "123", "https://api.com/postopengraph/123"},
		{"profile", "abc", "https://api.com/profileopengraph/abc"},
		{"postCollection", "xyz", "https://api.com/collectionopengraph/xyz"},
	}

	for _, tt := range tests {
		res := GetPreviewImageUrl("https://api.com", tt.kind, tt.id)
		if res != tt.expected {
			t.Errorf("for %s/%s: expected %s, got %s", tt.kind, tt.id, tt.expected, res)
		}
	}
}

func TestFetchMedia(t *testing.T) {
	mediaServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("X-Test"); got != "1" {
			t.Fatalf("missing header, got %q", got)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("media-bytes"))
	}))
	defer mediaServer.Close()

	client := NewClient("https://example.com", "", mediaServer.Client())
	media := &MediaMetadata{Files: []MediaFileMetadata{{ID: "m1", Profile: MediaFileProfileQuality, URL: mediaServer.URL}}}

	data, err := client.FetchMedia(context.Background(), media, &FetchOptions{Header: http.Header{"X-Test": []string{"1"}}})
	if err != nil {
		t.Fatalf("FetchMedia failed: %v", err)
	}
	if string(data) != "media-bytes" {
		t.Fatalf("unexpected body: %q", string(data))
	}
}

func TestFetchPreviewImage(t *testing.T) {
	previewServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("X-Preview"); got != "yes" {
			t.Fatalf("missing header, got %q", got)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("preview-bytes"))
	}))
	defer previewServer.Close()

	client := NewClient("https://unused.example", "", previewServer.Client())
	data, err := client.FetchPreviewImage(context.Background(), "post", "123", &FetchOptions{
		Endpoint: previewServer.URL,
		Header:   http.Header{"X-Preview": []string{"yes"}},
	})
	if err != nil {
		t.Fatalf("FetchPreviewImage failed: %v", err)
	}
	if string(data) != "preview-bytes" {
		t.Fatalf("unexpected body: %q", string(data))
	}
}

func TestFetchMediaNoFilesError(t *testing.T) {
	client := NewClient("https://example.com", "", nil)
	_, err := client.FetchMedia(context.Background(), &MediaMetadata{}, nil)
	if err == nil {
		t.Fatal("expected error")
	}
	if got := err.Error(); got != "No media file is available for the requested media object." {
		t.Fatalf("unexpected error: %s", got)
	}
}

func TestFetchPreviewImageEndpointOverride(t *testing.T) {
	previewServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/profileopengraph/abc" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		_, _ = io.WriteString(w, "ok")
	}))
	defer previewServer.Close()

	client := NewClient("https://unused.example", "", previewServer.Client())
	data, err := client.FetchPreviewImage(context.Background(), "profile", "abc", &FetchOptions{Endpoint: previewServer.URL})
	if err != nil {
		t.Fatalf("FetchPreviewImage failed: %v", err)
	}
	if string(data) != "ok" {
		t.Fatalf("unexpected body: %q", string(data))
	}
}
