package departments

import (
	"github.com/golobby/container"
)

// Init creates and injects the a departments repo.
func Init() {
	container.Singleton(
		func() Repo {
			return NewRepo()
		},
	)
}
