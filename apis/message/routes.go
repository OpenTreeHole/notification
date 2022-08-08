package message

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app fiber.Router) {
	app.Get("/messages", ListMessages)
	app.Post("/messages/clear", ClearMessages)
	app.Delete("/messages/:id", DeleteMessage)
}
