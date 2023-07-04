package tests

import (
	"github.com/opentreehole/go-common"
	"notification/app"
)

func init() {
	common.RegisterApp(app.Create())
}
