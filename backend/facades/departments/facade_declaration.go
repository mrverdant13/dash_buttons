package departments

import "github.com/mrverdant13/dash_buttons/backend/graph/model"

// Repo manages departments.
type Repo interface {
	Create(newDepartmentData model.NewDepartment) (*model.Department, error)
	GetByID(id uint64) (*model.Department, error)
	GetAll() ([]*model.Department, error)
	DeleteByID(id uint64) (*model.Department, error)
}
