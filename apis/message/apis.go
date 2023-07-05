package message

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opentreehole/go-common"
	. "notification/models"
	"notification/push"
)

func RegisterRoutes(app fiber.Router) {
	app.Post("/messages", SendMessage)
}

// SendMessage
// @Summary Send a message
// @Description Send to multiple recipients and save to db, admin only.
// @Tags Message
// @Produce application/json
// @Param json body models.Message true "json"
// @Router /messages [post]
// @Success 201 {object} Message
func SendMessage(c *fiber.Ctx) error {
	message, err := common.ValidateBody[Message](c)
	if err != nil {
		return err
	}

	go push.Send(message)

	return c.Status(201).JSON(message)
}
