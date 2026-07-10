package auth

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

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
