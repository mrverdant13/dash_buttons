package departments

import "github.com/mrverdant13/dash_buttons/backend/graph/model"

// Repo manages departments.
type Repo interface {
	Create(newDepartmentData model.NewDepartment) (*model.Department, error)
	GetByID(id string) (*model.Department, error)
	GetAll() ([]*model.Department, error)
}
