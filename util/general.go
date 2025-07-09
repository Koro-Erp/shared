package util

import (
	"encoding/json"
	"fmt"
	"net/mail"
	"strings"
)

func IsValidEmail(email string) bool {
    _, err := mail.ParseAddress(email)
    return err == nil
}

func PrintStruct(data interface{}){
    jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
}

func LowerFirst(s string) string {
	if s == "" {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}