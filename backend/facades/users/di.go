package users

import (
	"log"

	"github.com/golobby/container"
	"github.com/mrverdant13/dash_buttons/backend/config"
	"github.com/mrverdant13/dash_buttons/backend/internal/database/dbmodel"
	"github.com/mrverdant13/dash_buttons/backend/utilities"
	"gorm.io/gorm"
)

// Init creates and injects the an users repo.
func Init(createAdminIfNone bool) {
	container.Singleton(
		func(
			gormDB *gorm.DB,
			adminUser config.AdminUser,
		) Repo {
			if createAdminIfNone {
				user := dbmodel.User{
					Email: adminUser.Email,
				}

				hashedPassword, err := utilities.HashPassword(adminUser.Password)
				if err != nil {
					log.Fatalln(err.Error())
				}

				result := gormDB.
					Where(dbmodel.User{IsAdmin: true}).
					Attrs(dbmodel.User{HashedPassword: hashedPassword, IsAdmin: true}).
					FirstOrCreate(&user, user)

				if result.Error != nil {
					log.Fatalln(result.Error.Error())
				}
			}

			return NewRepo(
				gormDB,
			)
		},
	)
}
