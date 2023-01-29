package message

import (
	. "notification/models"
	"notification/push"
	. "notification/utils"

	"github.com/gofiber/fiber/v2"
)

// SendMessage
// @Summary Send a message
// @Description Send to multiple recipients and save to db, admin only.
// @Tags Message
// @Produce application/json
// @Param json body CreateModel true "json"
// @Router /messages [post]
// @Success 201 {object} Message
func SendMessage(c *fiber.Ctx) error {
	var body CreateModel
	err := ValidateBody(c, &body)
	if err != nil {
		return err
	}

	message := Message{
		Type:        body.Type,
		Title:       body.Title,
		Description: body.Description,
		Data:        body.Data,
		URL:         body.URL,
		Recipients:  body.Recipients,
	}
	if message.Title == "" {
		message.Title = generateTitle(body.Type)
	}
	if message.Description == "" {
		message.Description = generateDescription(body.Type, body.Data)
	}

	go push.Send(message)

	return Serialize(c.Status(201), &message)
}
