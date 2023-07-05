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
	authorization = "key=" + getMipushKey()
}

func getMipushKey() string {
	data, err := os.ReadFile(config.Config.MipushKeyPath)
	if err != nil {
		pwd, _ := os.Getwd()
		log.Fatal().Str("pwd", pwd).Err(err).Str("scope", "init mipush").Msg("failed to read mipush key")
	}
	return string(data)
}
