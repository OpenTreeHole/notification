package app

import (
	"fmt"
	. "notification/models"
)

// models must be registered here to migrate into the database
func init() {
	fmt.Println("migrate database...")
	err := DB.AutoMigrate(&PushToken{}, &Message{}, &MessageUser{})
	if err != nil {
		panic(err)
	}
}
