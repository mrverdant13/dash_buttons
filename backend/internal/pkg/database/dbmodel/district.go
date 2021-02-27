package dbmodel

import (
	"github.com/mrverdant13/dash_buttons/backend/graph/model"
	"gorm.io/gorm"
)

// District is a SQL model.
type District struct {
	gorm.Model
	Name       string
	ProvinceID uint64
}

// ToGQL converts the SQL model to a GraphQL model.
func (r District) ToGQL() model.District {
	return model.District{
		ID:         int64(r.ID),
		ProvinceID: int64(r.ProvinceID),
		Name:       r.Name,
	}
}

// Districts is a slice of "District" SQL models.
type Districts []*District

// ToGQL converts the SQL model to a GraphQL model.
func (r Districts) ToGQL() []*model.District {
	var _districts []*model.District

	for _, district := range r {
		_district := district.ToGQL()
		_districts = append(_districts, &_district)
	}

	return _districts
}
