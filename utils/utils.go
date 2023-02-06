package utils

import (
	"log"
	"os"
	"path"
)

func getBasePath() string {
	basePath := os.Getenv("BASE_PATH")
	if basePath == "" {
		log.Println("BASE_PATH not set, relative path may be incorrect")
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
