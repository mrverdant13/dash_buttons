package users

import (
	"fmt"
	"log"

	"github.com/mrverdant13/dash_buttons/backend/graph/gqlmodel"
	"github.com/mrverdant13/dash_buttons/backend/internal/pkg/database/dbmodel"
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

func (r *repo) CreateUser(newUser gqlmodel.NewUser) (*gqlmodel.User, error) {
	hashedPassword, err := hashPassword(newUser.Password)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	user := dbmodel.User{
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

	_user := gqlmodel.User{
		ID:    int64(user.ID),
		Email: user.Email,
	}

	return &_user, nil
}

func (r *repo) GetByID(id uint64) (*gqlmodel.User, error) {
	var user dbmodel.User

	result := r.gormDB.First(&user, id)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	}

	_user := gqlmodel.User{
		ID:    int64(user.ID),
		Email: user.Email,
	}

	return &_user, result.Error
}

func (r *repo) Authenticate(loginData gqlmodel.Login) (uint64, error) {
	user := dbmodel.User{
		Email: loginData.Email,
	}

	result := r.gormDB.Where(&user).First(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	rightCreds := checkPasswordHash(
		loginData.Password,
		user.HashedPassword,
	)

	if rightCreds {
		return uint64(user.ID), nil
	}

	return 0, fmt.Errorf("Wrong credentials")
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
