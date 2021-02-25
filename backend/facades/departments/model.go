package departments

import "gorm.io/gorm"

// Department is a SQL model.
type Department struct {
	gorm.Model
	Name string
}
