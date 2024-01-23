package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetApiKey extracts ApiKey info from
// the header of an HTTP request header
// Example: Auhorization: ApiKey <insert apikey here>

func GetApiKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("no authentication info found")
	}

	authSplit := strings.Split(authHeader, " ")
	if len(authSplit) != 2 {
		return "", errors.New("malformed authorization header")
	}

	if authSplit[0] != "ApiKey" {
		return "", errors.New("malformed first part of authorization header")
	}

	return authSplit[1], nil
}
