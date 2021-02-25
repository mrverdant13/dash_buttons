package departments

import (
	"github.com/golobby/container"
	"gorm.io/gorm"
)

// Init creates and injects the a departments repo.
func Init() {
	container.Singleton(
		func(gormDB *gorm.DB) Repo {
			return NewRepo(
				gormDB,
			)
		},
	)
}
