package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header := http.Header{}
	header.Set("Authorization", "ApiKey 123")
	got, err := GetAPIKey(header)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := "123"
	if want != got {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
