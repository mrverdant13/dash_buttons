package auth

import "github.com/mrverdant13/dash_buttons/backend/graph/model"

// Service manages authentication.
type Service interface {
	CreateUser(newUser model.NewUser) (*model.User, error)
	Authenticate(email, password string) (bool, error)
	GetUserByToken(token string) (*model.User, error)
}
