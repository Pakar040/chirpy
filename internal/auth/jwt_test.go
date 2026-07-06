package auth

import (
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
