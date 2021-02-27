package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type service struct {
	secretKey []byte
}

// NewService creates a new auth service.
func NewService(secretKey string) Service {
	return &service{
		secretKey: []byte(secretKey),
	}
}

const (
	expirationDelta = time.Hour * 24

	userIDKey     = "userId"
	expirationKey = "expiration"
)

func (r *service) GetUserIDByToken(token string) (string, error) {
	_token, err := jwt.Parse(
		token,
		func(token *jwt.Token) (interface{}, error) {
			return r.secretKey, nil
		},
	)
	if err != nil {
		return "", err
	}

	claims, ok := _token.Claims.(jwt.MapClaims)
	if !ok || !_token.Valid {
		return "", fmt.Errorf("Token parsing error")
	}

	userID := claims[userIDKey].(string)

	return userID, nil
}

func (r *service) GenerateToken(userID string) (string, error) {
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
