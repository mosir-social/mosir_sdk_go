package mosir_sdk_go

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestSmokeBetaEndpoint(t *testing.T) {
	if os.Getenv("MOSIR_SMOKE") != "1" {
		t.Skip("set MOSIR_SMOKE=1 to run beta.mosir.app smoke test")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	client := NewClient("https://beta.mosir.app/api/v1", "", nil)
	payload := []byte(`{"query":"query Smoke { __typename }"}`)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, client.endpoint, bytes.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.httpClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status %d: %s", resp.StatusCode, string(body))
	}
}

func TestSmokeGraphQLSSEPostCreatedByAuthor(t *testing.T) {
	if os.Getenv("MOSIR_SMOKE") != "1" {
		t.Skip("set MOSIR_SMOKE=1 to run beta.mosir.app smoke test")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := NewClient("https://beta.mosir.app/api/v1", "", nil)
	authorID := "GBWfRinN_Ya65D3SJaNS4"
	var got int
	err := client.PostCreatedByAuthor(ctx, authorID, PostTypePost, func(event PostCreatedByAuthorWsResponse) error {
		got++
		_ = event
		return nil
	})
	if err != nil && !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("subscription smoke failed: %v", err)
	}
	if got == 0 && err == nil {
		t.Fatal("expected a timeout or event")
	}
}
