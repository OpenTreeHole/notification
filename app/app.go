package app

import (
	"notification/apis"
	"notification/utils"

	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
)

func Create() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: utils.MyErrorHandler,
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})
	RegisterMiddlewares(app)
	apis.RegisterRoutes(app)

	return app
}
