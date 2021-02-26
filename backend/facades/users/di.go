package users

import (
	"github.com/golobby/container"
	"gorm.io/gorm"
)

// Init creates and injects the an users repo.
func Init() {
	container.Singleton(
		func(gormDB *gorm.DB) Repo {
			return NewRepo(
				gormDB,
			)
		},
	)
}
