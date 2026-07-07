package auth

import (
	"net/http"
	"strings"
)

func GetBearerToken(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", http.ErrNoCookie
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", http.ErrNoCookie
	}

	return strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer ")), nil
}
