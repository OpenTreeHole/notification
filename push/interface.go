package push

import (
	. "notification/models"
)

// I like OOP.

// Sender sends a message to multiple tokens(devices) through a certain push service
type Sender interface {
	New(message *Message, tokens []string) // Initialize the sender.
	Send() bool                            // Send notification.
	Clear()                                // Clear expired tokens.
}

// SenderFactoryInterface is an interface for SenderFactory.
type SenderFactoryInterface interface {
	CreateSender(service PushService) Sender
}

// SenderFactory is a factory for Sender.
type SenderFactory struct {
}
