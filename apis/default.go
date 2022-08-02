package apis

import (
	"github.com/gofiber/fiber/v2"
	"notification/data"
)

// Index
// @Produce application/json
// @Success 200 {object} models.MessageModel
// @Router / [get]
func Index(c *fiber.Ctx) error {
	return c.Send(data.MetaFile)
}
