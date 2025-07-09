package models

import "crypto/rsa"

type KeyConfig struct{
	EncryptionKey string
	PublicKey *rsa.PublicKey
	JwtKey string
}