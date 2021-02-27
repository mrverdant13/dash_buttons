package dbmodel

import (
	"github.com/mrverdant13/dash_buttons/backend/graph/model"
	"gorm.io/gorm"
)

// Department is a SQL model.
type Department struct {
	gorm.Model
	Name string
}

// ToGQL converts the SQL model to a GraphQL model.
func (r Department) ToGQL() model.Department {
	return model.Department{
		ID:   int64(r.ID),
		Name: r.Name,
	}
}

// Departments is a slice of "Department" SQL models.
type Departments []*Department

// ToGQL converts the SQL model to a GraphQL model.
func (r Departments) ToGQL() []*model.Department {
	var _departments []*model.Department

	for _, department := range r {
		_department := department.ToGQL()
		_departments = append(_departments, &_department)
	}

	return _departments
}
