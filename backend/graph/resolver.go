package graph

import (
	"github.com/mrverdant13/dash_buttons/backend/facades/auth"
	"github.com/mrverdant13/dash_buttons/backend/facades/departments"
	"github.com/mrverdant13/dash_buttons/backend/facades/provinces"
	"github.com/mrverdant13/dash_buttons/backend/facades/users"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver implementats schema.graphql
type Resolver struct {
	authService     auth.Service
	departmentsRepo departments.Repo
	provincesRepo   provinces.Repo
	usersRepo       users.Repo
}

// NewResolver creates a GraphQL resolver.
func NewResolver(
	authService auth.Service,
	departmentsRepo departments.Repo,
	provincesRepo provinces.Repo,
	usersRepo users.Repo,
) Resolver {
	return Resolver{
		authService:     authService,
		departmentsRepo: departmentsRepo,
		provincesRepo:   provincesRepo,
		usersRepo:       usersRepo,
	}
}
