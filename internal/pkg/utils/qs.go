package utils

import (
	"fmt"
	"net/url"
	"strings"
)

// GenerateQueryString creates a URL-encoded query string from a map of key-value pairs.
func GenerateQueryString(query map[string]string) string {
	var params []string
	for key, value := range query {
		param := fmt.Sprintf("%s=%s", url.QueryEscape(key), url.QueryEscape(value))
		params = append(params, param)
	}
	return strings.Join(params, "&")
}
