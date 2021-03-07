package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/mrverdant13/dash_buttons/backend/graph/gqlmodel"
	"github.com/mrverdant13/dash_buttons/backend/internal/middlewares"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input gqlmodel.NewUser) (*gqlmodel.User, error) {
	if *input.IsAdmin {
		adminUser := middlewares.CtxAdminUser(ctx)
		if adminUser == nil {
			err := fmt.Errorf("Access denied")
			log.Println(err.Error())
			return nil, err
		}
	}
	return r.usersRepo.CreateUser(input)
}
