package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CanPreprocess interface {
	Preprocess(c *fiber.Ctx) error
}

func Serialize(c *fiber.Ctx, obj CanPreprocess) error {
	err := obj.Preprocess(c)
	if err != nil {
		return err
	}
	return c.JSON(obj)
}

func ReText2IntArray(IDs [][]string) ([]int, error) {
	ansIDs := make([]int, 0)
	for _, v := range IDs {
		id, err := strconv.Atoi(v[1])
		if err != nil {
			return nil, err
		}
		ansIDs = append(ansIDs, id)
	}
	return ansIDs, nil
}
