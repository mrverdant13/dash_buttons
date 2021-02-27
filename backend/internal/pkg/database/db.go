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
	"gorm.io/gorm"

	//
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	mysqlgorm "gorm.io/driver/mysql"
)

const (
	mySQLDriver           = "mysql"
	mySQLConnStringFormat = "%s:%s@tcp(%s:%s)/%s?parseTime=true"

	// DepartmentsTable is the departments table name.
	DepartmentsTable = "Departments"
)

func initMySQLDatabase() {
	container.Singleton(
		func(destinationConf config.DbConf) *gorm.DB {
			connectionString := fmt.Sprintf(
				mySQLConnStringFormat,
				destinationConf.Username,
				destinationConf.Password,
				destinationConf.Host,
				strconv.Itoa(int(destinationConf.Port)),
				destinationConf.Database,
			)

			sqlDB, err := sql.Open(
				mySQLDriver,
				connectionString,
			)
			if err != nil {
				log.Fatalln(err.Error())
			}

			gormDB, err := gorm.Open(
				mysqlgorm.New(
					mysqlgorm.Config{
						Conn: sqlDB,
					},
				), &gorm.Config{},
			)

			err = sqlDB.Ping()
			if err != nil {
				log.Fatalln(err.Error())
			}

			return gormDB
		},
	)
}

// Migrate uses migration SQL files.
func Migrate() {
	container.Make(
		func(gormDB *gorm.DB) {
			sqlDB, err := gormDB.DB()
			if err != nil {
				log.Fatalln(err.Error())
			}

			err = sqlDB.Ping()
			if err != nil {
				log.Fatalln(err.Error())
			}

			driver, _ := mysql.WithInstance(
				sqlDB, &mysql.Config{},
			)

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
				log.Fatalln(err.Error())
			}
		},
	)
}
