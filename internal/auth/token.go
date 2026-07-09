package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetBearerToken(headers http.Header) (string, error) {
	authString := headers.Get("Authorization")

	auth := strings.Split(authString, " ")
	if len(auth) != 2 && auth[0] != "Bearer" {
		return "", errors.New("Authorization header deos not contain a Bearer Token")
	}

	tokenString := auth[1]
	if tokenString == "" {
		return "", errors.New("No bearer token provided")
	}

	return tokenString, nil
}
