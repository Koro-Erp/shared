package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"koro-erp/shared/internal/models"
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

	// Extract and forward both tokens if present
	if authToken := c.GetHeader("Authorization"); authToken != "" {
		req.Header.Set("Authorization", authToken)
	}
	if userToken := c.GetHeader("X-User-Token"); userToken != "" {
		req.Header.Set("X-User-Token", userToken)
	}

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