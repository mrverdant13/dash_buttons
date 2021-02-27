package database

import (
	"log"

	"github.com/golobby/container"
	"gorm.io/gorm"
)

// Init creates and injects the "*gorm.DB" for the entire app using the MySQL driver.
func Init(migrationPolicy MigrationPolicy) {
	initMySQLDatabase()
	Migrate(migrationPolicy)
}

// Dispose frees resources used by instances initialized by "Init".
func Dispose() {
	container.Make(
		func(gormDB *gorm.DB) {
			db, err := gormDB.DB()
			if err != nil {
				log.Fatalln(err.Error())
			}

			err = db.Close()
			if err != nil {
				log.Fatalln(err.Error())
			}
		},
	)
}
