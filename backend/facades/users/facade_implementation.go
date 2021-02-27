package users

import (
	"fmt"
	"log"
	"strconv"

	"github.com/mrverdant13/dash_buttons/backend/graph/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type repo struct {
	gormDB *gorm.DB
}

// NewRepo creates a new users repo.
func NewRepo(gormDB *gorm.DB) Repo {
	return &repo{
		gormDB: gormDB,
	}
}

func (r *repo) CreateUser(newUser model.NewUser) (*model.User, error) {
	hashedPassword, err := hashPassword(newUser.Password)
	if err != nil {
		log.Println(err.Error())
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
		log.Println(result.Error.Error())
		return nil, result.Error
	}

	_user := model.User{
		ID:    strconv.FormatInt(int64(user.ID), 10),
		Email: user.Email,
	}

	return &_user, nil
}

func (r *repo) GetByID(id string) (*model.User, error) {
	var user User

	result := r.gormDB.First(&user, id)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	}

	_user := model.User{
		ID:    strconv.FormatInt(int64(user.ID), 10),
		Email: user.Email,
	}

	return &_user, result.Error
}

func (r *repo) Authenticate(loginData model.Login) (string, error) {
	user := User{
		Email: loginData.Email,
	}

	result := r.gormDB.Where(&user).First(&user)
	if result.Error != nil {
		return "", result.Error
	}

	rightCreds := checkPasswordHash(
		loginData.Password,
		user.HashedPassword,
	)

	if rightCreds {
		return strconv.FormatInt(int64(user.ID), 10), nil
	}

	return "", fmt.Errorf("Wrong credentials")
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
