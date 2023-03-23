package database

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	migrateDatabaseClickhouse "github.com/golang-migrate/migrate/v4/database/clickhouse"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	gormClickhouseDriver "gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

var DB *gorm.DB
var SqlDB *sql.DB

func Connect() {
	var err error
	DB, err = gorm.Open(
		gormClickhouseDriver.Open("clickhouse://default:default@clickhouse:9000/default"),
		&gorm.Config{},
	)

	if err != nil {
		panic(err.Error())
	}

	SqlDB, err = DB.DB()
}

func AutoMigrate() {
	driver, err := migrateDatabaseClickhouse.WithInstance(SqlDB, &migrateDatabaseClickhouse.Config{})

	if err != nil {
		panic(err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://src/migrations",
		"default",
		driver,
	)

	if err != nil {
		panic(err.Error())
	}

	// m.Down()
	m.Up()
}
