package main

import (
	"crypto/ecdsa"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const strTokenPublicKey = `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEYv7IAdxabngdzvJYhSOQRcCXOqUU
rQw0Zxj1avDi9Zqyz+8rjyW7UEdeCU5onYPnjv/POR6NvtGPm0Ws+TOKSA==
-----END PUBLIC KEY-----`

var TokenPublicKey *ecdsa.PublicKey = nil

const strTokenPrivateKey = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIP5QzT+MH7xct+tBirb4sli+l5QPzcuHR9hRlcA+NsBzoAoGCCqGSM49
AwEHoUQDQgAEYv7IAdxabngdzvJYhSOQRcCXOqUUrQw0Zxj1avDi9Zqyz+8rjyW7
UEdeCU5onYPnjv/POR6NvtGPm0Ws+TOKSA==
-----END EC PRIVATE KEY-----`

var TokenPrivateKey *ecdsa.PrivateKey = nil

func init() {
	TokenPublicKey, _ = jwt.ParseECPublicKeyFromPEM([]byte(strTokenPublicKey))
	TokenPrivateKey, _ = jwt.ParseECPrivateKeyFromPEM([]byte(strTokenPrivateKey))
}

func MakeToken(name string) (string, error) {
	jwtToken := &jwt.Token{
		Header: map[string]interface{}{
			"alg": "ES256",
		},
		Claims: jwt.MapClaims{
			"Expired": time.Now().Local().AddDate(0, 1, 0).UnixMilli(),
			"User":    name,
		},
		Method: jwt.SigningMethodES256,
	}
	bearer, err := jwtToken.SignedString(TokenPrivateKey)
	if err != nil {
		return "", err
	}
	return bearer, nil
}

func ParseToken(Token string) (valid bool, refresh bool, name string) {
	token, err := jwt.Parse(Token, func(token *jwt.Token) (interface{}, error) {
		return TokenPublicKey, nil
	})
	if err != nil {
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		name := claims["User"].(string)
		Time := int64(claims["Expired"].(float64))
		if Time < time.Now().Local().AddDate(0, 0, 1).UnixMilli() {
			return true, true, name
		}
		return true, false, name
	} else {
		return false, false, ""
	}
}
