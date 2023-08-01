package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Config struct {
	Mode  string `env:"MODE" envDefault:"dev"`
	Debug bool   `env:"DEBUG" envDefault:"false"`
	// example: user:pass@tcp(127.0.0.1:3306)/dbname?parseTime=true
	// for more detail, see https://github.com/go-sql-driver/mysql#dsn-data-source-name
	DbUrl string `env:"DB_URL"`
	// mipush callback only support http
	MipushCallbackUrl string `env:"MIPUSH_CALLBACK_URL" envDefault:"http://notification.fduhole.com/api/callback/mipush"`

	// in production mode, use docker secrets
	MipushKeyPath      string `env:"MIPUSH_KEY_PATH" envDefault:"data/mipush.pem"`
	APNSKeyPath        string `env:"APNS_KEY_PATH" envDefault:"data/apns.pem"`
	APNSKeyPathV2      string `env:"APNS_KEY_PATH_V2" envDefault:"data/apns_v2.p8"`
	IOSPackageName     string `env:"IOS_PACKAGE_NAME" envDefault:"io.github.danxi-dev.dan-xi"`
	IOSPackageNameV2   string `env:"IOS_PACKAGE_NAME_V2" envDefault:"com.fduhole.danxi"`
	AndroidPackageName string `env:"ANDROID_PACKAGE_NAME" envDefault:"io.github.danxi_dev.dan_xi"`
}

func init() { // load config from environment variables
	err := env.Parse(&Config)
	if err != nil {
		log.Fatal().Err(err).Msg("load config from environment variables failed")
	}

	if Config.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Debug().Any("config", Config).Msg("config loaded")
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}
