package util

import (
	"github.com/dgrijalva/jwt-go"
)

type StandardClaims struct {
	Audience  string `json:"aud,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	Id        string `json:"jti,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	Issuer    string `json:"iss,omitempty"`
	NotBefore int64  `json:"nbf,omitempty"`
	Subject   string `json:"sub,omitempty"`
}

type MyCustomClaims struct {
	ID             int64
	Username       string
	StandardClaims jwt.StandardClaims
}

func GenerateToken() (string, error) {
	return "xx", nil
}

func ParseToken(token string) {

	return
}
