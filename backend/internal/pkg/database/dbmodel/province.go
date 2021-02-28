package dbmodel

import (
	"github.com/mrverdant13/dash_buttons/backend/graph/gqlmodel"
	"gorm.io/gorm"
)

// Province is a SQL model.
type Province struct {
	gorm.Model
	Name         string
	DepartmentID uint64
	Districts    []District
}

// ToGQL converts the SQL model to a GraphQL model.
func (r Province) ToGQL() gqlmodel.Province {
	return gqlmodel.Province{
		ID:           int64(r.ID),
		Name:         r.Name,
		DepartmentID: int64(r.DepartmentID),
	}
}

// Provinces is a slice of "Province" SQL models.
type Provinces []*Province

// ToGQL converts the SQL model to a GraphQL model.
func (r Provinces) ToGQL() []*gqlmodel.Province {
	var _province []*gqlmodel.Province

	for _, department := range r {
		_department := department.ToGQL()
		_province = append(_province, &_department)
	}

	return _province
}
