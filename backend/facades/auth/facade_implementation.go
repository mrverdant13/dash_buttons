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

	userKey       = "user"
	expirationKey = "expiration"
)

func (r *service) CreateUser(newUser model.NewUser) (*model.User, error) {
	hashedPassword, err := hashPassword(newUser.Password)
	if err != nil {
		return nil, err
	}

	user := User{
		Email:          newUser.Email,
		HashedPassword: hashedPassword,
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

	userExists := checkPasswordHash(
		password,
		user.HashedPassword,
	)

	return userExists, nil
}

func (r *service) GetUserByToken(token string) (*model.User, error) {
	user, err := r.parseToken(token)
	if err != nil {
		return nil, err
	}

	result := r.gormDB.Where(&user).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	_user := model.User{
		ID:    strconv.FormatInt(int64(user.ID), 10),
		Email: user.Email,
	}

	return &_user, nil
}

func (r *service) generateToken(user User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims[userKey] = user
	claims[expirationKey] = time.Now().Add(expirationDelta).Unix()

	tokenString, err := token.SignedString(r.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (r *service) parseToken(tokenStr string) (*User, error) {
	token, err := jwt.Parse(
		tokenStr,
		func(token *jwt.Token) (interface{}, error) {
			return r.secretKey, nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("Token parsing error")
	}

	user := claims[userKey].(User)

	return &user, nil
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
