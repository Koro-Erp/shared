package util

import (
	"encoding/json"
	"strings"
)

var sensitiveHeaders = []string{"authorization", "x-api-key", "x-user-token"}
var sensitiveBodyFields = []string{"password", "token", "secret", "api_key","access_token","refresh_token"}

// ✅ Mask sensitive fields in JSON request/response bodies
func MaskSensitiveData(body string) string {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		// Not JSON, return as-is
		return body
	}

	for _, field := range sensitiveBodyFields {
		if _, exists := data[field]; exists {
			data[field] = "***"
		}
	}

	maskedBody, _ := json.Marshal(data)
	return string(maskedBody)
}

// ✅ Mask sensitive headers
func MaskSensitiveHeaders(headers map[string][]string) string {
	masked := make(map[string][]string)

	for key, values := range headers {
		lowerKey := strings.ToLower(key)
		if contains(sensitiveHeaders, lowerKey) {
			masked[key] = []string{"***"}
		} else {
			masked[key] = values
		}
	}

	result, err := json.Marshal(masked)
	if err != nil {
		return "{}"
	}
	return string(result)
}

func contains(slice []string, target string) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}
