package database

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	"github.com/golobby/container"
	"github.com/mrverdant13/dash_buttons/backend/config"

	//
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

const (
	mySQLDriver           = "mysql"
	mySQLConnStringFormat = "%s:%s@tcp(%s:%s)/%s"
)

func initMySQLDatabase() {
	container.Singleton(
		func(destinationConf config.DbConf) *sql.DB {
			connectionString := fmt.Sprintf(
				mySQLConnStringFormat,
				destinationConf.Username,
				destinationConf.Password,
				destinationConf.Host,
				strconv.Itoa(int(destinationConf.Port)),
				destinationConf.Database,
			)

			db, err := sql.Open(
				mySQLDriver,
				connectionString,
			)
			if err != nil {
				log.Fatalln(err.Error())
			}

			err = db.Ping()
			if err != nil {
				log.Panic(err)
			}

			return db
		},
	)
}

// Migrate uses migration SQL files.
func Migrate() {
	container.Make(
		func(db *sql.DB) {
			err := db.Ping()
			if err != nil {
				log.Fatalln(err)
			}

			driver, _ := mysql.WithInstance(db, &mysql.Config{})

			m, err := migrate.NewWithDatabaseInstance(
				"file://internal/pkg/database/migrations/mysql",
				mySQLDriver,
				driver,
			)
			if err != nil {
				log.Fatalln(err.Error())
			}

			err = m.Up()
			if err != nil && err != migrate.ErrNoChange {
				log.Fatalln(err)
			}
		},
	)
}
