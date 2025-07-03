package util

import (
	"crypto/rsa"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)
func LoadPublicKey(path string) *rsa.PublicKey {
	pem, err := os.ReadFile(path)
	if err != nil {
		panic("Failed to load public key")
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(pem)
	if err != nil {
		log.Println(err)
		panic("Invalid public key")
	}
	return key
}

func ExtractSubFromJWT(tokenString string) string {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		fmt.Println("JWT parse error:", err)
		return ""
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if sub, ok := claims["sub"].(string); ok {
			return sub
		}
	}

	return ""
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