package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"os"
)

var jwtKey = []byte(os.Getenv("JWL_SECRET"))

type Claims struct {
	UserID  int  `json:"user_id"`
	IsAdmin bool `json:"is_admin"`
	jwt.RegisteredClaims
}

func GenerateToken(userID int, is_admin bool) (string, error) {
	claims := Claims{
		UserID:  userID,
		IsAdmin: is_admin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
