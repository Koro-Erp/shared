package util

import (
	"crypto/rsa"
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