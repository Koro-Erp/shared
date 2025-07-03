package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"github.com/Koro-Erp/shared/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendLogRequest(c *gin.Context, data models.AppLog, url string) error {
	// Marshal AppLog struct to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal log data: %w", err)
	}

	// Create outbound POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create log request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Use the helper function to set token headers
	CopyAuthHeaders(c, req)

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send log request: %w", err)
	}
	defer resp.Body.Close()

	// Log response for debugging
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Log response:", string(body))

	return nil
}

// CopyAuthHeaders copies Authorization and X-User-Token from the incoming Gin context to the outbound HTTP request
func CopyAuthHeaders(c *gin.Context, req *http.Request) {
	if authToken := c.GetHeader("Authorization"); authToken != "" {
		req.Header.Set("Authorization", authToken)
	}
	if userToken := c.GetHeader("X-User-Token"); userToken != "" {
		req.Header.Set("X-User-Token", userToken)
	}
}