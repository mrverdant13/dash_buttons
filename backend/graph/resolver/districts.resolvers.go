package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/mrverdant13/dash_buttons/backend/graph/model"
	"github.com/mrverdant13/dash_buttons/backend/internal/middlewares"
)

func (r *mutationResolver) CreateDistrict(ctx context.Context, input model.NewDistrict) (*model.District, error) {
	user := middlewares.CtxUser(ctx)
	if user == nil {
		err := fmt.Errorf("Access denied")
		log.Println(err.Error())
		return nil, err
	}

	return r.districtsRepo.Create(input)
}

func (r *mutationResolver) DeleteDistrict(ctx context.Context, id string) (*model.District, error) {
	user := middlewares.CtxUser(ctx)
	if user == nil {
		err := fmt.Errorf("Access denied")
		log.Println(err.Error())
		return nil, err
	}

	return r.districtsRepo.DeleteByID(id)
}

func (r *queryResolver) Districts(ctx context.Context) ([]*model.District, error) {
	return r.districtsRepo.GetAll()
}

func (r *queryResolver) District(ctx context.Context, id string) (*model.District, error) {
	return r.districtsRepo.GetByID(id)
}
