package departments

import (
	"github.com/mrverdant13/dash_buttons/backend/graph/model"
)

type repo struct {
	list []*model.Department
}

// NewRepo creates a new departments repo.
func NewRepo() Repo {
	return &repo{
		list: []*model.Department{},
	}
}

func (r *repo) Create(name string) (*model.Department, error) {
	department := model.Department{
		ID:   name,
		Name: name,
	}

	r.list = append(r.list, &department)

	return &department, nil
}

func (r *repo) GetAll() ([]*model.Department, error) {
	return r.list, nil
}
