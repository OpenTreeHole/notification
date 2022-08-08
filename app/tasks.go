package app

import "notification/apis/message"

func startTasks() {
	go message.PurgeMessage()
}
