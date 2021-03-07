package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/mrverdant13/dash_buttons/backend/graph/gqlmodel"
)

func (r *mutationResolver) CreateDistrict(ctx context.Context, input gqlmodel.NewDistrict) (*gqlmodel.District, error) {
	return r.districtsRepo.Create(input)
}

func (r *mutationResolver) DeleteDistrict(ctx context.Context, id int64) (*gqlmodel.District, error) {
	return r.districtsRepo.DeleteByID(uint64(id))
}

func (r *queryResolver) Districts(ctx context.Context) ([]*gqlmodel.District, error) {
	return r.districtsRepo.GetAll()
}

func (r *queryResolver) District(ctx context.Context, id int64) (*gqlmodel.District, error) {
	return r.districtsRepo.GetByID(uint64(id))
}
