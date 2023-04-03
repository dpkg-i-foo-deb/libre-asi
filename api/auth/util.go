package auth

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt"
)

func IsAdmin(tk string) bool {

	role, err := RoleFromToken(tk)

	if err != nil {
		return false
	}

	if role != "admin" {
		return false
	}

	return true
}

func EmailFromToken(tokeString string) (string, error) {
	claims, err := GetTokenClaims(tokeString)

	if err != nil {
		return "", err
	}

	return claims.Email, nil
}

func RoleFromToken(tokeString string) (string, error) {
	claims, err := GetTokenClaims(tokeString)

	if err != nil {
		return "", err
	}

	return claims.Role, nil
}

func NeedsPasswordResetFromToken(tokenString string) (bool, error) {
	claims, err := GetTokenClaims(tokenString)

	if err != nil {
		return false, err
	}

	return claims.NeedsPasswordReset, nil
}

func GetTokenClaims(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("AUTH_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	return token.Claims.(*CustomClaims), nil
}
