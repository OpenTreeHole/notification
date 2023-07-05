package token

import (
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
func ListTokens(c *fiber.Ctx) error {
	var tokens []PushToken
	userID, ok := c.Locals("user_id").(int)
	if !ok {
		return common.Unauthorized()
	}
	DB.Where("user_id = ?", userID).Find(&tokens)
	return c.JSON(tokens)
}

// AddToken
// @Summary Add Token of a User
// @Tags Token
// @Produce application/json
// @Param json body models.PushToken true "json"
// @Router /users/push-tokens [post]
// @Success 200 {object} PushToken
func AddToken(c *fiber.Ctx) (err error) {
	token, err := common.ValidateBody[PushToken](c)
	if err != nil {
		return err
	}

	userID, ok := c.Locals("user_id").(int)
	if !ok {
		return common.Unauthorized()
	}
	token.UserID = userID
	err = DB.Transaction(func(tx *gorm.DB) (err error) {
		// remove all device_id duplicates
		err = tx.Where("device_id = ? and token <> ?", token.DeviceID, token.Token).Delete(&PushToken{}).Error
		if err != nil {
			return err
		}

		// create or update new token
		return tx.Save(&token).Error
	})
	if err != nil {
		return err
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
	body, err := common.ValidateBody[DeleteModel](c)
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

type DeleteModel struct {
	DeviceID string `json:"device_id" validate:"max=64"`
}
