package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func CreateJWT(secretKey, username, password string) string {
	claims := jwt.MapClaims{
		"username": username,
		"password": password,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return err.Error()
	}
	return signedToken
}
