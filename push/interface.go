package push

import (
	. "notification/models"
)

// hasbai: I like OOP.
// jingyijun: I don't like OOP.

// Sender sends a message to multiple tokens(devices) through a certain push service
type Sender interface {
	New(message *Message, tokens []PushToken) // New Initialize the sender.
	Send()                                    // Send notification.
	Clear()                                   // Clear expired tokens.
}
