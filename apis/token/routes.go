package token

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app fiber.Router) {
	app.Get("/users/push-tokens", ListTokens)
	app.Post("/users/push-tokens", AddToken)
	app.Put("/users/push-tokens", AddToken)
	app.Delete("/users/push-tokens", DeleteToken)
}
