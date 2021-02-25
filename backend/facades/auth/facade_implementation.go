package auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mrverdant13/dash_buttons/backend/graph/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type service struct {
	gormDB    *gorm.DB
	secretKey string
}

// NewService creates a new auth service.
func NewService(
	secretKey string,
	gormDB *gorm.DB,
) Service {
	return &service{
		secretKey: secretKey,
		gormDB:    gormDB,
	}
}

const (
	expirationDelta = time.Hour * 24

	usernameKey   = "username"
	expirationKey = "expiration"
)

func (r *service) CreateUser(newUser model.NewUser) (*model.User, error) {
	hashedPassword, err := hashPassword(newUser.Password)
	if err != nil {
		return nil, err
	}

	user := User{
		Email:    newUser.Email,
		Password: hashedPassword,
	}

	result := r.gormDB.Create(
		&user,
	)
	if result.Error != nil {
		return nil, result.Error
	}

	_user := model.User{
		ID:    strconv.FormatInt(int64(user.ID), 10),
		Email: user.Email,
	}

	return &_user, nil
}
func (r *service) Authenticate(email, password string) (bool, error) {
	user := User{
		Email: email,
	}

	result := r.gormDB.Where(&user).First(&user)
	if result.Error != nil {
		return false, result.Error
	}

	userExists := checkPasswordHash(password, user.Password)

	return userExists, nil
}

func (r *service) generateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims[usernameKey] = username
	claims[expirationKey] = time.Now().Add(expirationDelta).Unix()

	tokenString, err := token.SignedString(r.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (r *service) parseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(
		tokenStr,
		func(token *jwt.Token) (interface{}, error) {
			return r.secretKey, nil
		},
	)
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("Token parsing error")
	}

	username := claims[usernameKey].(string)

	return username, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		14,
	)

	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)
	return err == nil
}

func (r *service) getUserIDByEmail(email string) (uint64, error) {
	user := User{
		Email: email,
	}

	result := r.gormDB.Where(&user).First(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	return uint64(user.ID), nil
}
