package dbmodel

import "gorm.io/gorm"

// Province is a SQL model.
type Province struct {
	gorm.Model
	Name         string
	DepartmentID uint64
}
