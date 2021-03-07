package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/mrverdant13/dash_buttons/backend/graph/generated"
	"github.com/mrverdant13/dash_buttons/backend/graph/gqlmodel"
)

func (r *mutationResolver) CreateProvince(ctx context.Context, input gqlmodel.NewProvince) (*gqlmodel.Province, error) {
	return r.provincesRepo.Create(input)
}

func (r *mutationResolver) DeleteProvince(ctx context.Context, id int64) (*gqlmodel.Province, error) {
	return r.provincesRepo.DeleteByID(uint64(id))
}

func (r *provinceResolver) Districts(ctx context.Context, obj *gqlmodel.Province) ([]*gqlmodel.District, error) {
	return r.districtsRepo.GetAllByProvinceID(uint64(obj.ID))
}

func (r *queryResolver) Provinces(ctx context.Context) ([]*gqlmodel.Province, error) {
	return r.provincesRepo.GetAll()
}

func (r *queryResolver) Province(ctx context.Context, id int64) (*gqlmodel.Province, error) {
	return r.provincesRepo.GetByID(uint64(id))
}

// Province returns generated.ProvinceResolver implementation.
func (r *Resolver) Province() generated.ProvinceResolver { return &provinceResolver{r} }

type provinceResolver struct{ *Resolver }
