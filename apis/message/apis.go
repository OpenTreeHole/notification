package message

import (
	. "notification/models"
	"notification/push"
	. "notification/utils"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app fiber.Router) {
	app.Post("/messages", SendMessage)
}

// SendMessage
// @Summary Send a message
// @Description Send to multiple recipients and save to db, admin only.
// @Tags Message
// @Produce application/json
// @Param json body CreateModel true "json"
// @Router /messages [post]
// @Success 201 {object} Message
func SendMessage(c *fiber.Ctx) error {
	var message Message
	err := ValidateBody(c, &message)
	if err != nil {
		return err
	}

	go push.Send(message)

	return c.Status(201).JSON(message)
}
