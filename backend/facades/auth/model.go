package auth

import "gorm.io/gorm"

// User is a SQL model.
type User struct {
	gorm.Model
	Email          string
	HashedPassword string
}
