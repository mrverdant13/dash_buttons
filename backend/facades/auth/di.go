package auth

import (
	"github.com/golobby/container"
	"github.com/mrverdant13/dash_buttons/backend/config"
	"gorm.io/gorm"
)

// Init creates and injects the an auth service.
func Init() {
	container.Singleton(
		func(
			jwtConf config.JWTConf,
			gormDB *gorm.DB,
		) Service {
			return NewService(
				jwtConf.SecretKey,
				gormDB,
			)
		},
	)
}
