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

func (r *mutationResolver) CreateProvince(ctx context.Context, input model.NewProvince) (*model.Province, error) {
	user := middlewares.CtxUser(ctx)
	if user == nil {
		err := fmt.Errorf("Access denied")
		log.Println(err.Error())
		return nil, err
	}

	return r.provincesRepo.Create(input)
}

func (r *mutationResolver) DeleteProvince(ctx context.Context, id int64) (*model.Province, error) {
	user := middlewares.CtxUser(ctx)
	if user == nil {
		err := fmt.Errorf("Access denied")
		log.Println(err.Error())
		return nil, err
	}

	return r.provincesRepo.DeleteByID(uint64(id))
}

func (r *queryResolver) Provinces(ctx context.Context) ([]*model.Province, error) {
	return r.provincesRepo.GetAll()
}

func (r *queryResolver) Province(ctx context.Context, id int64) (*model.Province, error) {
	return r.provincesRepo.GetByID(uint64(id))
}
