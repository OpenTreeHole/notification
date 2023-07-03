package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/opentreehole/go-common"
	"notification/config"
)

func RegisterMiddlewares(app *fiber.App) {
	app.Use(recover.New(recover.Config{EnableStackTrace: true}))
	app.Use(common.MiddlewareGetUserID)
	app.Use(common.MiddlewareCustomLogger)
	if config.Config.Debug {
		app.Use(pprof.New())
	}
}
