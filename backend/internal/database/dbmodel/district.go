package dbmodel

import (
	"github.com/mrverdant13/dash_buttons/backend/graph/gqlmodel"
	"gorm.io/gorm"
)

// District is a SQL model.
type District struct {
	gorm.Model
	Name       string `gorm:"not null;uniqueIndex:idx_unique_district;size:50"`
	ProvinceID uint64 `gorm:"not null;uniqueIndex:idx_unique_district"`
}

// ToGQL converts the SQL model to a GraphQL model.
func (r District) ToGQL() gqlmodel.District {
	return gqlmodel.District{
		ID:         int64(r.ID),
		ProvinceID: int64(r.ProvinceID),
		Name:       r.Name,
	}
}

// Districts is a slice of "District" SQL models.
type Districts []*District

// ToGQL converts the SQL model to a GraphQL model.
func (r Districts) ToGQL() []*gqlmodel.District {
	var _districts []*gqlmodel.District

	for _, district := range r {
		_district := district.ToGQL()
		_districts = append(_districts, &_district)
	}

	return _districts
}
