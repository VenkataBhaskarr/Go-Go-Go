package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyJWT(tokenString, secretKey string) jwt.MapClaims {
	// first step to check whether the token is encoded with same algo we now use for derc
	decryptToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is correct
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})
	if err != nil {
		return nil
	}
	if claims, ok := decryptToken.Claims.(jwt.MapClaims); ok && decryptToken.Valid {
		return claims // Token verification succeeded
	}
	return nil
}
