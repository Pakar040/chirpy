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
		return "", errors.New("Authorization header does not contain a Bearer Token")
	}

	tokenString := auth[1]
	if tokenString == "" {
		return "", errors.New("No bearer token provided")
	}

	return tokenString, nil
}

func GetAPIKey(headers http.Header) (string, error) {
	authString := headers.Get("Authorization")

	auth := strings.Split(authString, " ")
	if len(auth) != 2 && auth[0] != "ApiKey" {
		return "", errors.New("Authorization header does not contain a ApiKey")
	}

	tokenString := auth[1]
	if tokenString == "" {
		return "", errors.New("No ApiKey provided")
	}

	return tokenString, nil
}
