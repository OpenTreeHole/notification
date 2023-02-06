package config

import "github.com/caarlos0/env/v6"

var Config struct {
	Mode  string `env:"MODE" envDefault:"dev"`
	Debug bool   `env:"DEBUG" envDefault:"false"`
	// example: user:pass@tcp(127.0.0.1:3306)/dbname?parseTime=true
	// for more detail, see https://github.com/go-sql-driver/mysql#dsn-data-source-name
	DbUrl string `env:"DB_URL,required"`
	// in production mode, use docker secrets
	MipushKeyPath      string `env:"MIPUSH_KEY_PATH" envDefault:"data/mipush.pem"`
	APNSKeyPath        string `env:"APNS_KEY_PATH" envDefault:"data/apns.pem"`
	IOSPackageName     string `env:"IOS_PACKAGE_NAME" envDefault:"io.github.danxi-dev.dan-xi"`
	AndroidPackageName string `env:"ANDROID_PACKAGE_NAME" envDefault:"io.github.danxi_dev.dan_xi"`
}

func init() { // load config from environment variables
	err := env.Parse(&Config)
	if err != nil {
		panic(err)
	}
}
