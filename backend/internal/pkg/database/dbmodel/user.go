package dbmodel

import "gorm.io/gorm"

// User is a SQL model.
type User struct {
	gorm.Model
	Email          string `gorm:"not null;unique"`
	HashedPassword string `gorm:"not null;column:password"`
}
