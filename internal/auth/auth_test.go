package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("returns API key when Authorization header is correctly formatted", func(t *testing.T) {
		headers := http.Header{}
		headers.Add("Authorization", "ApiKey testkey")
		got, err := GetAPIKey(headers)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		want := "testkey"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns error when no Authorization header is included", func(t *testing.T) {
		headers := http.Header{}
		_, err := GetAPIKey(headers)
		if err != ErrNoAuthHeaderIncluded {
			t.Errorf("got %v, want %v", err, ErrNoAuthHeaderIncluded)
		}
	})

	t.Run("returns error when Authorization header is malformed", func(t *testing.T) {
		headers := http.Header{}
		headers.Add("Authorization", "malformed header")
		_, err := GetAPIKey(headers)
		if err == nil || err.Error() != "malformed authorization header" {
			t.Errorf("got %v, want error with message 'malformed authorization header'", err)
		}
	})
}
