package models

import (
	"time"
)

type AppLog struct {
	ID             int            `json:"id"`
	Timestamp      time.Time      `json:"timestamp"`
	ServiceName    string         `json:"service_name"`
	UserIdentifier string         `json:"user_identifier"`
	Level          string         `json:"level"`
	Message        string         `json:"message"`
	StackTrace     string         `json:"stack_trace"`
	ExtraData      map[string]interface{} `json:"extra_data"`
}

type GatewayLog struct {
	ID             int       `json:"id"`
	Timestamp      time.Time `json:"timestamp"`
	ServiceName    string    `json:"service_name"`
	UserIdentifier string    `json:"user_identifier"`
	ClientIP       string    `json:"client_ip"`
	RequestMethod  string    `json:"request_method"`
	RequestURL     string    `json:"request_url"`
	RequestHeaders string    `json:"request_headers"` // store as raw JSON string
	RequestBody    string    `json:"request_body"`
	ResponseStatus int       `json:"response_status"`
	ResponseBody   string    `json:"response_body"`
	DurationMs     int64     `json:"duration_ms"`
}

