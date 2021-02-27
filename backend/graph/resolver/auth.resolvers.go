package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/mrverdant13/dash_buttons/backend/graph/generated"
	"github.com/mrverdant13/dash_buttons/backend/graph/model"
)

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	userID, err := r.usersRepo.Authenticate(input)
	if err != nil {
		return "", err
	}

	return r.authService.GenerateToken(userID)
}

func (r *mutationResolver) RefreshToken(ctx context.Context, expiredToken string) (string, error) {
	userID, err := r.authService.GetUserIDByToken(expiredToken)
	if err != nil {
		return "", err
	}

	return r.authService.GenerateToken(userID)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
