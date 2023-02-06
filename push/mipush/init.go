package mipush

import (
	"fmt"
	"net/http"
	"notification/config"
	"os"
	"time"
)

const (
	mipushURL = "https://api.xmpush.xiaomi.com/v3/message/regid"
	timeout   = time.Second * 10
)

var client = http.Client{Timeout: timeout}
var authorization string

func init() {
	fmt.Println("init mipush")
	authorization = "key=" + getMipushKey()
}

func getMipushKey() string {
	data, err := os.ReadFile(config.Config.MipushKeyPath)
	if err != nil {
		panic(err)
	}
	return string(data)
}
