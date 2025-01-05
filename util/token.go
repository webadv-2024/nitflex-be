package util

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken generate token
func GenerateToken(userID string, username string, expiredTime time.Time) (string, error) {
	jwtClaims := &JwtClaims{
		Id:       userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    userID,
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	secretKey := os.Getenv("JWT_SECRET_KEY")
	signedString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedString, nil
}

func ParseToken(tokenString string) (*JwtClaims, error) {
	claims := &JwtClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secretKey := os.Getenv("JWT_SECRET_KEY")
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func Verify(context context.Context, tokenString string) (*JwtClaims, error) {
	claims, err := ParseToken(tokenString)
	if err != nil {
		return nil, err
	}
	expiredTime := claims.ExpiresAt.Time

	if isExpired := time.Now().After(expiredTime); isExpired {
		return nil, errors.New("token expired")
	}
	return claims, nil
}
