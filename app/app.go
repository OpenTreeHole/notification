package app

import (
	"github.com/goccy/go-json"
	"notification/apis"
	"notification/utils"

	"github.com/gofiber/fiber/v2"
)

func Create() *fiber.App {
	utils.Logger, _ = utils.InitLog()

	app := fiber.New(fiber.Config{
		ErrorHandler: utils.MyErrorHandler,
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})
	RegisterMiddlewares(app)
	apis.RegisterRoutes(app)

	startTasks()

	return app
}
