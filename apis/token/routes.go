package token

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app fiber.Router) {
	app.Get("/users/push-tokens", ListTokens)
	app.Post("/users/push-tokens", CreateToken)
	app.Put("/users/push-tokens", CreateToken)
	app.Patch("/users/push-tokens/_webvpn", CreateToken)
	app.Delete("/users/push-tokens", DeleteToken)
	app.Delete("/users/push-tokens/_all", DeleteAllTokens)
}

func RegisterTasks() {
	// delete expired tokens
	go deleteExpiredTokens(context.Background())
}
