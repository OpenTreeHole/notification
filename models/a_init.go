package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"notification/config"
	"os"
)

var DB *gorm.DB

var gormConfig = &gorm.Config{
	NamingStrategy: schema.NamingStrategy{
		SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
	},
	PrepareStmt: true, // use PrepareStmt for `Save`, `Update` and `Delete`
}

func mysqlDB() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(config.Config.DbUrl), gormConfig)
}

func sqliteDB() (*gorm.DB, error) {
	err := os.MkdirAll("data", 0750)
	if err != nil {
		panic(err)
	}
	return gorm.Open(sqlite.Open("data/sqlite.db"), gormConfig)
}

func memoryDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("file::memory:?cache=shared"), gormConfig)
}

func init() {
	var err error
	switch config.Config.Mode {
	case "production":
		DB, err = mysqlDB()
	case "test":
		DB, err = memoryDB()
	case "dev":
		if config.Config.DbUrl == "" {
			DB, err = sqliteDB()
		} else {
			DB, err = mysqlDB()
		}
	case "perf":
		DB, err = sqliteDB()
	default:
		panic("unknown mode")
	}
	if err != nil {
		panic(err)
	}

	switch config.Config.Mode {
	case "test":
		fallthrough
	case "dev":
		DB = DB.Debug()
	}

	err = DB.AutoMigrate(&PushToken{})
	if err != nil {
		panic(err)
	}
}
