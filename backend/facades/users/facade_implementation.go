package users

import (
	"fmt"
	"log"

	"github.com/mrverdant13/dash_buttons/backend/graph/gqlmodel"
	"github.com/mrverdant13/dash_buttons/backend/internal/database/dbmodel"
	"github.com/mrverdant13/dash_buttons/backend/utilities"
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
	hashedPassword, err := utilities.HashPassword(newUser.Password)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	user := dbmodel.User{
		Email:          newUser.Email,
		HashedPassword: hashedPassword,
		IsAdmin:        *newUser.IsAdmin,
	}

	result := r.gormDB.Create(
		&user,
	)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	}

	_user := gqlmodel.User{
		ID:      int64(user.ID),
		Email:   user.Email,
		IsAdmin: user.IsAdmin,
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
		ID:      int64(user.ID),
		Email:   user.Email,
		IsAdmin: user.IsAdmin,
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

	rightCreds := utilities.CheckPasswordHash(
		loginData.Password,
		user.HashedPassword,
	)

	if rightCreds {
		return uint64(user.ID), nil
	}

	return 0, fmt.Errorf("Wrong credentials")
}
