package app

import (
	"notification/apis"
	"notification/utils"

	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
)

func Create() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:               "notification",
		ErrorHandler:          utils.MyErrorHandler,
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		DisableStartupMessage: true,
	})
	RegisterMiddlewares(app)
	apis.RegisterRoutes(app)
	apis.RegisterTasks()

	return app
}
