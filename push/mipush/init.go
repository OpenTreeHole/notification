package mipush

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"notification/config"
	"notification/utils"
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
	keyPath := utils.ToAbsolutePath(config.Config.MipushKeyPath)
	data, err := ioutil.ReadFile(keyPath)
	if err != nil {
		panic(err)
	}
	return string(data)
}
