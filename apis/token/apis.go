package token

import (
	"github.com/gofiber/fiber/v2"
	. "notification/models"
	. "notification/utils"
)

// ListTokens
// @Summary List Tokens of a User
// @Tags Token
// @Produce application/json
// @Router /users/push-tokens [get]
// @Success 200 {array} PushToken
func ListTokens(c *fiber.Ctx) error {
	var tokens []PushToken
	DB.Where("user_id = ?", c.Locals("userID").(int)).Find(&tokens)
	return c.JSON(tokens)
}

// AddToken
// @Summary Add Token of a User
// @Tags Token
// @Produce application/json
// @Param json body CreateModel true "json"
// @Router /users/push-tokens [post]
// @Success 200 {object} PushToken
func AddToken(c *fiber.Ctx) error {
	var body CreateModel
	err := ValidateBody(c, &body)
	if err != nil {
		return err
	}

	token := PushToken{
		UserID:   c.Locals("userID").(int),
		Service:  ParsePushService(body.Service),
		DeviceID: body.DeviceID,
		Token:    body.Token,
	}
	result := DB.Create(&token)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(token)
}

// DeleteToken
// @Summary Delete the token of a user's certain device
// @Tags Token
// @Produce application/json
// @Param json body DeleteModel true "json"
// @Router /users/push-tokens [delete]
// @Success 204
func DeleteToken(c *fiber.Ctx) error {
	var body DeleteModel
	err := ValidateBody(c, &body)
	if err != nil {
		return err
	}

	querySet := DB.Where("user_id = ?", c.Locals("userID").(int))
	if body.DeviceID != "" {
		querySet = querySet.Where("device_id = ?", body.DeviceID)
	}
	result := querySet.Delete(&PushToken{})
	if result.Error != nil {
		return result.Error
	}

	return c.Status(204).JSON(nil)
}
