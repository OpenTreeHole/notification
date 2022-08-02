package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

const (
	TypeMysql = iota
	TypeSqlite
)

var DBType uint

var gormConfig = &gorm.Config{
	NamingStrategy: schema.NamingStrategy{
		SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
	},
	PrepareStmt: true, // use PrepareStmt for `Save`, `Update` and `Delete`
}

func mysqlDB() (*gorm.DB, error) {
	DBType = TypeMysql
	return gorm.Open(mysql.Open(Config.DBURL), gormConfig)
}

func sqliteDB() (*gorm.DB, error) {
	DBType = TypeSqlite
	err := os.MkdirAll("data", 0750)
	if err != nil {
		panic(err)
	}
	return gorm.Open(sqlite.Open("data/sqlite.db"), gormConfig)
}

func memoryDB() (*gorm.DB, error) {
	DBType = TypeSqlite
	return gorm.Open(sqlite.Open("file::memory:?cache=shared"), gormConfig)
}

func init() {
	fmt.Println("init db config...")
	var err error
	switch Config.Mode {
	case "production":
		DB, err = mysqlDB()
	case "test":
		DB, err = memoryDB()
		DB = DB.Debug()
	case "dev":
		if Config.DBURL == "" {
			DB, err = sqliteDB()
		} else {
			DB, err = mysqlDB()
		}
		DB = DB.Debug()
	case "perf":
		DB, err = sqliteDB()
	default: // sqlite as default
		panic("unknown mode")
	}
	if err != nil {
		panic(err)
	}
}
