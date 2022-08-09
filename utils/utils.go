package utils

import (
	"github.com/gofiber/fiber/v2"
)

type CanPreprocess interface {
	Preprocess(c *fiber.Ctx) error
}

type numbers interface {
	int | uint | int8 | uint8 |
		int16 | uint16 | int32 | uint32 |
		int64 | uint64 | float32 | float64
}

func Min[T numbers](x T, y T) T {
	if x > y {
		return y
	} else {
		return x
	}
}
