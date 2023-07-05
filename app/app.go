package app

import (
	"github.com/goccy/go-json"
	"github.com/opentreehole/go-common"
	"notification/apis"

	"github.com/gofiber/fiber/v2"
)

func Create() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:               "notification",
		ErrorHandler:          common.ErrorHandler,
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		DisableStartupMessage: true,
	})
	RegisterMiddlewares(app)
	apis.RegisterRoutes(app)
	apis.RegisterTasks()

	return app
}
