package auth

import "github.com/mrverdant13/dash_buttons/backend/graph/model"

// Service manages authentication.
type Service interface {
	CreateUser(newUser model.NewUser) (*model.User, error)
	// GetUserIDByEmail(email string) (uint64, error)
	Authenticate(email, password string) (bool, error)
}
