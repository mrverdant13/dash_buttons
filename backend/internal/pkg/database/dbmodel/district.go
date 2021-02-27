package dbmodel

import "gorm.io/gorm"

// District is a SQL model.
type District struct {
	gorm.Model
	Name       string
	ProvinceID uint64
}
