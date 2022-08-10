package utils

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"path"
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

func getBasePath() string {
	basePath := os.Getenv("BASE_PATH")
	if basePath == "" {
		Logger.Warn("BASE_PATH not set, relative path may be incorrect")
	}
	return basePath
}

func ToAbsolutePath(relativePath string) string {
	if path.IsAbs(relativePath) {
		return relativePath
	}
	basePath := getBasePath()
	if basePath == "" {
		return relativePath
	}
	return path.Join(basePath, relativePath)
}
