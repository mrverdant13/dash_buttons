package users

import "github.com/mrverdant13/dash_buttons/backend/graph/model"

// Repo manages users.
type Repo interface {
	CreateUser(newUser model.NewUser) (*model.User, error)
	GetUserByID(userID uint64) (*model.User, error)
	Authenticate(loginData model.Login) (uint64, error)
}
