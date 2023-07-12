package mipush

import (
	. "notification/models"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestMipush(t *testing.T) {
	// get recipients from env
	recipientsString := os.Getenv("RECIPIENTS")
	recipientsStringList := strings.Split(recipientsString, ",")
	recipients := make([]int, len(recipientsStringList))
	for i, recipientString := range recipientsStringList {
		recipient, err := strconv.Atoi(recipientString)
		if err != nil {
			t.Errorf("Convert recipient failed: %v", err)
		}
		recipients[i] = recipient
	}

	sender := Sender{}
	sender.Message = &Message{
		Title:       "test",
		Description: "test",
		URL:         "https://www.google.com",
		Type:        "123",
		Data: Map{
			"test": "test",
			"url":  "https://www.google.com",
		},
		Recipients: recipients,
	}

	sender.Tokens = []PushToken{{Token: "Hh3mT3w/bMmTCdp7/D3HcG2VmJ3wmQHscOXBT5oCo5mE7XsgNDvlIzgGe5+gqCYh"}}

	sender.Send()
}
