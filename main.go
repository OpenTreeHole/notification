package main

import (
	"notification/app"
)

// @title Notification Center
// @version 2.0.0
// @description This is a notification microservice.

// @contact.name Maintainer OpenTreeHole
// @contact.email dev@fduhole.com

// @license.name Apache 2.0
// @license.url https://www.apache.org/licenses/LICENSE-2.0.html

// @host
// @BasePath /api

func main() {
	myApp := app.Create()
	err := myApp.Listen("0.0.0.0:8000")
	if err != nil {
		panic(err)
	}
}
