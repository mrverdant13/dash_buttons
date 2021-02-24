package database

import (
	"database/sql"
	"log"

	"github.com/golobby/container"
)

// Init creates and injects the "*sql.DB" for the entire app using the MySQL driver.
func Init(migrate bool) {
	initMySQLDatabase()
	if migrate {
		Migrate()
	}
}

// Dispose frees resources used by instances initialized by "Init".
func Dispose() {
	container.Make(
		func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				log.Fatalln(err.Error())
			}
		},
	)
}
