package resolver

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

func (r *departmentResolver) Provinces(ctx context.Context, obj *model.Department) ([]*model.Province, error) {
	return r.provincesRepo.GetAllByDepartmentID(uint64(obj.ID))
}

func (r *mutationResolver) CreateDepartment(ctx context.Context, input model.NewDepartment) (*model.Department, error) {
	user := middlewares.CtxUser(ctx)
	if user == nil {
		err := fmt.Errorf("Access denied")
		log.Println(err.Error())
		return nil, err
	}

	return r.departmentsRepo.Create(input)
}

func (r *mutationResolver) DeleteDepartment(ctx context.Context, id int64) (*model.Department, error) {
	user := middlewares.CtxUser(ctx)
	if user == nil {
		err := fmt.Errorf("Access denied")
		log.Println(err.Error())
		return nil, err
	}

	return r.departmentsRepo.DeleteByID(uint64(id))
}

func (r *queryResolver) Departments(ctx context.Context) ([]*model.Department, error) {
	return r.departmentsRepo.GetAll()
}

func (r *queryResolver) Department(ctx context.Context, id int64) (*model.Department, error) {
	return r.departmentsRepo.GetByID(uint64(id))
}

// Department returns generated.DepartmentResolver implementation.
func (r *Resolver) Department() generated.DepartmentResolver { return &departmentResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type departmentResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
