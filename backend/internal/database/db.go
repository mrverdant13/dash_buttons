package database

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/golobby/container"
	"github.com/mrverdant13/dash_buttons/backend/config"
	"github.com/mrverdant13/dash_buttons/backend/internal/database/dbmodel"
	"gorm.io/gorm"

	//
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
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

// MigrationPolicy holds the migration policy type.
type MigrationPolicy string

const (
	// NoMigration indicates that migration will not be executed.
	NoMigration = "none"

	// OrmMigration indicates that migration will be executed by the ORM.
	OrmMigration = "orm"

	// SQLScriptsMigration indicates that migration will be executed based on SQL scripts.
	SQLScriptsMigration = "script"
)

// Migrate uses migration SQL files.
func Migrate(migrationPolicy MigrationPolicy) {
	container.Make(
		func(gormDB *gorm.DB) {

			if migrationPolicy == OrmMigration {
				err := gormDB.AutoMigrate(
					&dbmodel.User{},
					&dbmodel.Department{},
					&dbmodel.Province{},
					&dbmodel.District{},
				)
				if err != nil {
					log.Fatalln(err.Error())
				}
			} else if migrationPolicy == SQLScriptsMigration {
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
					"file://internal/database/migrations/mysql",
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
			}
		},
	)
}
