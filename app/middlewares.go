package app

import (
	"notification/config"
	"notification/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func RegisterMiddlewares(app *fiber.App) {
	app.Use(recover.New())
	if config.Config.Mode != "perf" {
		app.Use(logger.New())
	}
	if config.Config.Debug {
		app.Use(pprof.New())
	}
	app.Use(getUserID)
}

func getUserID(c *fiber.Ctx) error {
	var userID int
	var err error

	if config.Config.Debug {
		userID = 1
	} else {
		userID, err = strconv.Atoi(c.Get("X-Consumer-Username"))
		if err != nil {
			return utils.Unauthorized("Unauthorized")
		}
	}

	c.Locals("userID", userID)

	return c.Next()
}
