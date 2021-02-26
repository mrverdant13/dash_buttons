package auth

import (
	"github.com/golobby/container"
	"github.com/mrverdant13/dash_buttons/backend/config"
)

// Init creates and injects the an auth service.
func Init() {
	container.Singleton(
		func(jwtConf config.JWTConf) Service {
			return NewService(
				jwtConf.SecretKey,
			)
		},
	)
}
