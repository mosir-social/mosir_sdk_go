package mosir_sdk_go

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClient(t *testing.T) {
	endpoint := "https://example.com/api/v1"
	token := "test-token"
	client := NewClient(endpoint, token, nil)

	if client.endpoint != endpoint {
		t.Errorf("expected endpoint %s, got %s", endpoint, client.endpoint)
	}
	if client.token != token {
		t.Errorf("expected token %s, got %s", token, client.token)
	}
	if client.httpClient == nil {
		t.Error("expected httpClient to be initialized")
	}
}

func TestWithAuth(t *testing.T) {
	endpoint := "https://example.com/api/v1"
	token := "test-token"
	client := NewClient(endpoint, token, nil)
	authClient := client.WithAuth()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader != "Bearer "+token {
			t.Errorf("expected Authorization header 'Bearer %s', got '%s'", token, authHeader)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	// We need to use the client's endpoint which is now the test server's URL
	req, _ := http.NewRequest("GET", ts.URL, nil)
	_, err := authClient.httpClient.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
}
