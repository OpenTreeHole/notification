package push

import (
	. "notification/models"
)

func manualSend(service PushService, title string, description string, data Map, tokens []string) {
	m := Message{
		Title:       title,
		Description: description,
		Data:        data,
	}
	sender := CreateSender(service)
	sender.New(&m, tokens)
	sender.Send()
	sender.Clear()
}
