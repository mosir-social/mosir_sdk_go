package mosir_sdk_go

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
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

	var parsed map[string]any
	if err := json.Unmarshal(body, &parsed); err != nil {
		t.Fatalf("response is not valid JSON: %v; body=%s", err, string(body))
	}
	if _, ok := parsed["data"]; !ok {
		t.Fatalf("expected data field in response: %s", string(body))
	}
	if !strings.Contains(string(body), "__typename") && parsed["data"] == nil {
		t.Fatalf("unexpected response body: %s", string(body))
	}
}
