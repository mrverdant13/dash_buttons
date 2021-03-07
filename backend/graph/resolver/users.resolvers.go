package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/mrverdant13/dash_buttons/backend/graph/gqlmodel"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input gqlmodel.NewUser) (*gqlmodel.User, error) {
	return r.usersRepo.CreateUser(input, false)
}

func (r *mutationResolver) CreateAdminUser(ctx context.Context, input gqlmodel.NewUser) (*gqlmodel.User, error) {
	return r.usersRepo.CreateUser(input, true)
}
