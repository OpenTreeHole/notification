package push

import (
	"log"
	. "notification/models"
	"notification/push/apns"
	"notification/push/base"
	"notification/push/mipush"
)

var factory = SenderFactory{}

// CreateSender creates a sender for a certain push service.
func (factory SenderFactory) CreateSender(service PushService) Sender {
	switch service {
	case ServiceAPNS:
		return &apns.Sender{}
	case ServiceMipush:
		return &mipush.Sender{}
	default:
		log.Printf("%s not implemented", service.String())
		return &base.Sender{}
	}
}

func Send(message Message) bool {
	var pushTokens []PushToken
	DB.Where("user_id IN ?", message.Recipients).Find(&pushTokens)
	serviceTokenMapping := make(map[PushService][]string)
	for _, serviceToken := range pushTokens {
		serviceTokenMapping[serviceToken.Service] = append(
			serviceTokenMapping[serviceToken.Service],
			serviceToken.Token,
		)
	}

	var success = true
	for _, service := range PushServices {
		tokens, ok := serviceTokenMapping[service]
		if !ok {
			continue
		}

		sender := factory.CreateSender(service)
		sender.New(&message, tokens)
		success = sender.Send() && success
		sender.Clear()
	}
	return success
}
