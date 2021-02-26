package users

import "github.com/mrverdant13/dash_buttons/backend/graph/model"

// Repo manages users.
type Repo interface {
	CreateUser(newUser model.NewUser) (*model.User, error)
	UserWithIDExists(userID uint64) (bool, error)
	Authenticate(email, password string) (bool, error)
}
