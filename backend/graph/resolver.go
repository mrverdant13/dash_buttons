package graph

import (
	"github.com/mrverdant13/dash_buttons/backend/facades/departments"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver implementats schema.graphql
type Resolver struct {
	departmentsRepo departments.Repo
}

// NewResolver creates a GraphQL resolver.
func NewResolver(departmentsRepo departments.Repo) Resolver {
	return Resolver{
		departmentsRepo: departmentsRepo,
	}
}
