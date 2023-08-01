package token

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/opentreehole/go-common"
	"gorm.io/gorm"

	. "notification/models"
)

// ListTokens
// @Summary List Tokens of a User
// @Tags Token
// @Produce application/json
// @Router /users/push-tokens [get]
// @Success 200 {array} PushToken
func ListTokens(c *fiber.Ctx) (err error) {
	// get user_id from jwt
	userID, ok := c.Locals("user_id").(int)
	if !ok {
		return common.Unauthorized()
	}

	var tokens []PushToken
	err = DB.Where("user_id = ?", userID).Find(&tokens).Error
	if err != nil {
		return err
	}
	return c.JSON(tokens)
}

// CreateToken
// @Summary Add Token of a User
// @Tags Token
// @Produce application/json
// @Param json body CreateTokenRequest true "json"
// @Router /users/push-tokens [post]
// @Router /users/push-tokens [put]
// @Success 200 {object} PushToken
func CreateToken(c *fiber.Ctx) (err error) {
	var body CreateTokenRequest
	err = common.ValidateBody(c, &body)
	if err != nil {
		return err
	}

	userID, ok := c.Locals("user_id").(int)
	if !ok {
		return common.Unauthorized()
	}

	token := PushToken{
		UserID:      userID,
		Service:     NewPushService(body.Service),
		DeviceID:    body.DeviceID,
		Token:       body.Token,
		PackageName: body.PackageName,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = DB.Transaction(func(tx *gorm.DB) (err error) {
		// remove all device_id duplicates
		err = tx.Where("device_id = ?", token.DeviceID).Delete(&PushToken{}).Error
		if err != nil {
			return err
		}

		// create or update new token
		return tx.Save(&token).Error
	})
	if err != nil {
		return err
	}

	return c.JSON(&token)
}

// DeleteToken
// @Summary Delete the token of a user's certain device
// @Tags Token
// @Produce application/json
// @Param json body DeleteModel true "json"
// @Router /users/push-tokens [delete]
// @Success 204
func DeleteToken(c *fiber.Ctx) (err error) {
	var body DeleteModel
	err = common.ValidateBody(c, &body)
	if err != nil {
		return err
	}

	userID, ok := c.Locals("user_id").(int)
	if !ok {
		return common.Unauthorized()
	}

	querySet := DB.Where("user_id = ?", userID)
	if body.DeviceID != "" {
		querySet = querySet.Where("device_id = ?", body.DeviceID)
	}
	err = querySet.Delete(&PushToken{}).Error
	if err != nil {
		return err
	}

	return c.Status(204).JSON(nil)
}

type DeleteModel struct {
	DeviceID string `json:"device_id" validate:"max=64"`
}

// DeleteAllTokens
// @Summary Delete all tokens of a user
// @Tags Token
// @Produce application/json
// @Router /users/push-tokens/_all [delete]
// @Success 204
func DeleteAllTokens(c *fiber.Ctx) (err error) {
	userID, ok := c.Locals("user_id").(int)
	if !ok {
		return common.Unauthorized()
	}

	err = DB.Where("user_id = ?", userID).Delete(&PushToken{}).Error
	if err != nil {
		return err
	}

	return c.Status(204).JSON(nil)
}
