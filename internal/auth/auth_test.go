package auth

import (
	"errors"
	"net/http"
	"testing"
	//"github.com/unloadingthecoding/learn-cicd-starter/internal/auth"
)

// var ErrNoAuthHeaderIncluded = errors.New("no authorization header included")

// GetAPIKey function here (as provided in your question)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedError error
	}{
		{
			name:          "No Authorization Header",
			headers:       http.Header{},
			expectedKey:   "",
			expectedError: ErrNoAuthHeaderIncluded,
		},
		{
			name: "Valid Authorization Header",
			headers: http.Header{
				"Authorization": []string{"ApiKey my-secret-api-key"},
			},
			expectedKey:   "my-secret-api-key",
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, gotErr := GetAPIKey(tt.headers)

			if gotKey != tt.expectedKey {
				t.Errorf("GetAPIKey() gotKey = %v, want %v", gotKey, tt.expectedKey)
			}

			if !errors.Is(gotErr, tt.expectedError) {
				t.Errorf("GetAPIKey() gotErr = %v, want %v", gotErr, tt.expectedError)
			}
		})
	}
}
