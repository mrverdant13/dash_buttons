package dbmodel

import (
	"github.com/mrverdant13/dash_buttons/backend/graph/gqlmodel"
	"gorm.io/gorm"
)

// Department is a SQL model.
type Department struct {
	gorm.Model
	Name      string
	Provinces []Province
}

// ToGQL converts the SQL model to a GraphQL model.
func (r Department) ToGQL() gqlmodel.Department {
	return gqlmodel.Department{
		ID:   int64(r.ID),
		Name: r.Name,
	}
}

// Departments is a slice of "Department" SQL models.
type Departments []*Department

// ToGQL converts the SQL model to a GraphQL model.
func (r Departments) ToGQL() []*gqlmodel.Department {
	var _departments []*gqlmodel.Department

	for _, department := range r {
		_department := department.ToGQL()
		_departments = append(_departments, &_department)
	}

	return _departments
}
