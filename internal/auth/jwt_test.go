package auth

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

func TestMakeJWTValidateJWT(t *testing.T) {
	uuidString := uuid.New().String()
	tests := map[string]struct {
		input    string
		expected string
	}{
		"simple": {input: uuidString, expected: uuidString},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			inputId, err := uuid.Parse(tc.input)
			if err != nil {
				t.Fatalf("Failed to parse test input into uuid: %s", err)
			}

			tokenSecret := "ABC123"
			tokenString, err := MakeJWT(inputId, tokenSecret, time.Minute)
			if err != nil {
				t.Fatalf("Failed to make JWT token: %s", err)
			}

			actualId, err := ValidateJWT(tokenString, tokenSecret)
			if err != nil {
				t.Fatalf("Failed to retreive uuid from token: %s", err)
			}

			diff := cmp.Diff(tc.expected, actualId.String())
			if diff != "" {
				t.Fatalf("%s", diff)
			}
		})
	}
}

func TestGetBearerToken(t *testing.T) {
	tests := map[string]struct {
		input    func() http.Header
		expected string
	}{
		"simple": {
			input: func() http.Header {
				h := http.Header{}
				h.Set("Authorization", fmt.Sprintf("Bearer %s", "kmJq1VgZD/xZgw=="))
				return h
			},
			expected: "kmJq1VgZD/xZgw==",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := GetBearerToken(tc.input())
			if err != nil {
				t.Fatalf("Failed to get bearer token from header: %s", err)
			}
			diff := cmp.Diff(tc.expected, actual)
			if diff != "" {
				t.Fatalf("%s", diff)
			}
		})
	}
}
