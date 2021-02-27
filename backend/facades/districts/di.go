package districts

import (
	"github.com/golobby/container"
	"github.com/mrverdant13/dash_buttons/backend/facades/provinces"
	"gorm.io/gorm"
)

// Init creates and injects the a districts repo.
func Init() {
	container.Singleton(
		func(
			gormDB *gorm.DB,
			provincesRepo provinces.Repo,
		) Repo {
			return NewRepo(
				gormDB,
				provincesRepo,
			)
		},
	)
}
