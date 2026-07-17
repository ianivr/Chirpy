package auth

import (
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	apiKey := headers.Get("Authorization")
	if apiKey == "" {
		return "", http.ErrNoCookie
	}

	if !strings.HasPrefix(apiKey, "ApiKey ") {
		return "", http.ErrNoCookie
	}

	return strings.TrimSpace(strings.TrimPrefix(apiKey, "ApiKey ")), nil
}
