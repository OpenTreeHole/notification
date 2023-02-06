package main

import (
	"log"
	"notification/app"
	"os"
	"os/signal"
	"syscall"
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
	a := app.Create()

	go func() {
		err := a.Listen("0.0.0.0:8000")
		if err != nil {
			log.Fatal(err)
		}
	}()

	interrupt := make(chan os.Signal, 1)

	// wait for CTRL-C interrupt
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interrupt

	// close app
	err := a.Shutdown()
	if err != nil {
		log.Println(err)
	}
}
