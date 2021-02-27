package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/mrverdant13/dash_buttons/backend/graph/generated"
	"github.com/mrverdant13/dash_buttons/backend/graph/model"
	"github.com/mrverdant13/dash_buttons/backend/internal/middlewares"
)

func (r *mutationResolver) CreateDepartment(ctx context.Context, input model.NewDepartment) (*model.Department, error) {
	user := middlewares.CtxUser(ctx)
	if user == nil {
		err := fmt.Errorf("Access denied")
		log.Println(err.Error())
		return nil, err
	}

	return r.departmentsRepo.Create(input)
}

func (r *mutationResolver) CreateProvince(ctx context.Context, input model.NewProvince) (*model.Province, error) {
	user := middlewares.CtxUser(ctx)
	if user == nil {
		err := fmt.Errorf("Access denied")
		log.Println(err.Error())
		return nil, err
	}

	return r.provincesRepo.Create(input)
}

func (r *mutationResolver) CreateDistrict(ctx context.Context, input model.NewDistrict) (*model.District, error) {
	user := middlewares.CtxUser(ctx)
	if user == nil {
		err := fmt.Errorf("Access denied")
		log.Println(err.Error())
		return nil, err
	}

	return r.districtsRepo.Create(input)
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return r.usersRepo.CreateUser(input)
}

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

func (r *queryResolver) Departments(ctx context.Context) ([]*model.Department, error) {
	return r.departmentsRepo.GetAll()
}

func (r *queryResolver) Department(ctx context.Context, id string) (*model.Department, error) {
	return r.departmentsRepo.GetByID(id)
}

func (r *queryResolver) Provinces(ctx context.Context) ([]*model.Province, error) {
	return r.provincesRepo.GetAll()
}

func (r *queryResolver) Province(ctx context.Context, id string) (*model.Province, error) {
	return r.provincesRepo.GetByID(id)
}

func (r *queryResolver) Districts(ctx context.Context) ([]*model.District, error) {
	return r.districtsRepo.GetAll()
}

func (r *queryResolver) District(ctx context.Context, id string) (*model.District, error) {
	return r.districtsRepo.GetByID(id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
