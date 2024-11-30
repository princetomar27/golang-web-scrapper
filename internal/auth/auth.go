package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey - extracts the API key from
// the headers of an API request

func GetAPIKeyFromHeaders(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
        return "", errors.New("user not authorized")
    }

	value := strings.Split(val," ")
	if len(value) != 2 {
		return "", errors.New("invalid Authorization header format") 
	}
	if value[0]!= "Bearer"{
		return "", errors.New("invalid Authorization header prefix")
	}

	return value[1],nil
}