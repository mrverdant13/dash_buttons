package users

import "github.com/mrverdant13/dash_buttons/backend/graph/gqlmodel"

// Repo manages users.
type Repo interface {
	CreateUser(newUser gqlmodel.NewUser, newUserIsAdmin bool) (*gqlmodel.User, error)
	GetByID(id uint64) (*gqlmodel.User, error)
	Authenticate(loginData gqlmodel.Login) (uint64, error)
}
