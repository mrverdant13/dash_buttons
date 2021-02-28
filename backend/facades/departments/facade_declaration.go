package departments

import "github.com/mrverdant13/dash_buttons/backend/graph/gqlmodel"

// Repo manages departments.
type Repo interface {
	Create(newDepartmentData gqlmodel.NewDepartment) (*gqlmodel.Department, error)
	GetByID(id uint64) (*gqlmodel.Department, error)
	GetAll() ([]*gqlmodel.Department, error)
	DeleteByID(id uint64) (*gqlmodel.Department, error)
}
