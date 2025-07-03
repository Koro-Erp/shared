package util

import (
	"crypto/rsa"
	"fmt"
	"log"
	"os"

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