package mipush

import (
	"github.com/rs/zerolog/log"
	"net/http"
	"notification/config"
	"os"
	"time"
)

const (
	mipushURL = "https://api.xmpush.xiaomi.com/v3/message/regid"
	timeout   = time.Second * 60
)

var client = http.Client{Timeout: timeout}
var authorization string

func init() {

	log.Debug().Msg("init mipush")
	key, _ := getMipushKey()
	authorization = "key=" + key
}

func getMipushKey() (string, error) {
	data, err := os.ReadFile(config.Config.MipushKeyPath)

	if err != nil {
		pwd, _ := os.Getwd()

		switch config.Config.Mode {
		case "production":
			// 在生产环境中严格要求 key 存在
			log.Fatal().
				Str("pwd", pwd).
				Err(err).
				Str("scope", "init mipush").
				Msg("failed to read mipush key in production")
			// 不会执行到这里
			return "", err

		case "dev", "test":
			// 在开发或测试环境中只警告，不终止
			log.Warn().
				Str("pwd", pwd).
				Err(err).
				Str("scope", "init mipush").
				Msg("failed to read mipush key, using empty key in non-production mode")
			return "", nil

		case "perf":
			// 压测环境用 mock key
			log.Info().
				Str("scope", "init mipush").
				Msg("using mock mipush key for perf mode")
			return "mock-mipush-key", nil

		default:
			// 未知模式时明确报错退出
			log.Fatal().
				Str("scope", "init mipush").
				Str("mode", config.Config.Mode).
				Msg("unknown mode while reading mipush key")
			return "", err
		}
	}
	return string(data), nil
}
