package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type service struct {
	secretKey string
}

// NewService creates a new auth service.
func NewService(secretKey string) Service {
	return &service{
		secretKey: secretKey,
	}
}

const (
	expirationDelta = time.Hour * 24

	userIDKey     = "userId"
	expirationKey = "expiration"
)

func (r *service) GetUserIDByToken(token string) (uint64, error) {
	return r.parseToken(token)

}

func (r *service) GenerateToken(userID uint64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims[userIDKey] = userID
	claims[expirationKey] = time.Now().Add(expirationDelta).Unix()

	tokenString, err := token.SignedString(r.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (r *service) parseToken(tokenStr string) (uint64, error) {
	token, err := jwt.Parse(
		tokenStr,
		func(token *jwt.Token) (interface{}, error) {
			return r.secretKey, nil
		},
	)
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, fmt.Errorf("Token parsing error")
	}

	userID := claims[userIDKey].(uint64)

	return userID, nil
}
