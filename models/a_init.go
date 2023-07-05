package models

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"notification/config"
	"os"
	"time"
)

var DB *gorm.DB

var zeroLogger = zerolog.New(os.Stdout)

var gormConfig = &gorm.Config{
	NamingStrategy: schema.NamingStrategy{
		SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
	},
	PrepareStmt: true, // use PrepareStmt for `Save`, `Update` and `Delete`
	Logger: logger.New(
		&zeroLogger,
		logger.Config{
			SlowThreshold:             time.Second,  // 慢 SQL 阈值
			LogLevel:                  logger.Error, // 日志级别
			IgnoreRecordNotFoundError: true,         // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,        // 禁用彩色打印
		},
	),
}

func mysqlDB() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(config.Config.DbUrl), gormConfig)
}

func sqliteDB() (*gorm.DB, error) {
	err := os.MkdirAll("data", 0750)
	if err != nil {
		log.Fatal().Err(err).Msg("create sqlite data dir failed")
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
		log.Fatal().Str("scope", "gorm init").Str("mode", config.Config.Mode).Msg("unknown mode")
	}
	if err != nil {
		log.Fatal().Err(err).Str("scope", "init gorm").Msg("connect to database failed")
	}

	switch config.Config.Mode {
	case "test":
		fallthrough
	case "dev":
		DB = DB.Debug()
	}

	err = DB.AutoMigrate(&PushToken{})
	if err != nil {
		log.Fatal().Err(err).Str("scope", "init gorm").Msg("auto migrate failed")
	}
}
