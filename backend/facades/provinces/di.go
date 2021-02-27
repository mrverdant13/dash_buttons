package provinces

import (
	"github.com/golobby/container"
	"github.com/mrverdant13/dash_buttons/backend/facades/departments"
	"gorm.io/gorm"
)

// Init creates and injects the a provinces repo.
func Init() {
	container.Singleton(
		func(
			gormDB *gorm.DB,
			departmentsRepo departments.Repo,
		) Repo {
			return NewRepo(
				gormDB,
				departmentsRepo,
			)
		},
	)
}
