package message

import (
	"github.com/gofiber/fiber/v2"
	. "notification/models"
	"notification/push"
	. "notification/utils"
)

// ListMessages
// @Summary List Messages of a User
// @Tags Message
// @Produce application/json
// @Router /messages [get]
// @Success 200 {array} Message
func ListMessages(c *fiber.Ctx) error {
	var messages []Message
	DB.Raw(`
		SELECT * FROM message
		INNER JOIN message_user ON message.id = message_user.message_id 
		WHERE message_user.user_id = ?`,
		c.Locals("userID").(int),
	).Scan(&messages)
	return c.JSON(messages)
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

	err = DB.Create(&message).Error
	if err != nil {
		return err
	}

	go push.Send(message)

	return c.Status(201).JSON(message)
}

// ClearMessages
// @Summary Clear Messages of a User
// @Tags Message
// @Produce application/json
// @Router /messages/clear [post]
// @Success 204
func ClearMessages(c *fiber.Ctx) error {
	result := DB.Exec(
		"DELETE FROM message_user WHERE user_id = ?",
		c.Locals("userID").(int),
	)
	if result.Error != nil {
		return result.Error
	}
	return c.Status(204).JSON(nil)
}

// DeleteMessage
// @Summary Delete a message of a user
// @Tags Message
// @Produce application/json
// @Router /messages/{id} [delete]
// @Param id path int true "message id"
// @Success 204
func DeleteMessage(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	result := DB.Exec(
		"DELETE FROM message_user WHERE user_id = ? AND message_id = ?",
		c.Locals("userID").(int), id,
	)
	if result.Error != nil {
		return result.Error
	}
	return c.Status(204).JSON(nil)
}
