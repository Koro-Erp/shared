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
