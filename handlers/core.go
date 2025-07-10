package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Koro-Erp/shared/models"
	"github.com/Koro-Erp/shared/util"

	"github.com/gin-gonic/gin"
)

func SendAppLogRequest(c *gin.Context, data models.AppLog, url string) error {
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
	util.CopyAuthHeaders(c, req)

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send log request: %w", err)
	}
	defer resp.Body.Close()

	// Log response for debugging
	// body, _ := io.ReadAll(resp.Body)
	// fmt.Println("Log response:", string(body))

	return nil
}

func SendGatewayLogRequest(data models.GatewayLog, url string) error {
	// Marshal GatewayLog struct to JSON
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

	// Use the AuthHeaders from the model directly
	for key, value := range data.AuthHeaders {
		req.Header.Set(key, value) // Set all auth headers
	}

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send log request: %w", err)
	}
	defer resp.Body.Close()

	return nil
}