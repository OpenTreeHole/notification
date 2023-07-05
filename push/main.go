package push

import (
	"github.com/rs/zerolog/log"
	. "notification/models"
	"notification/push/apns"
	"notification/push/base"
	"notification/push/mipush"
)

// CreateSender creates a sender for a certain push service.
func CreateSender(service PushService) Sender {
	switch service {
	case ServiceAPNS:
		return &apns.Sender{}
	case ServiceMipush:
		return &mipush.Sender{}
	default:
		log.Error().Msgf("%s not implemented", service)
		return &base.Sender{}
	}
}

func Send(message *Message) {

	// load push tokens from database
	var pushTokens []PushToken
	err := DB.Where("user_id IN ?", message.Recipients).Find(&pushTokens).Error
	if err != nil {
		log.Err(err).Msg("Get push tokens failed")
		return
	}
	if len(pushTokens) == 0 {
		return
	}

	serviceTokenMapping := make(map[PushService][]string)
	for _, serviceToken := range pushTokens {
		serviceTokenMapping[serviceToken.Service] = append(
			serviceTokenMapping[serviceToken.Service],
			serviceToken.Token,
		)
	}

	for _, service := range PushServices {
		tokens, ok := serviceTokenMapping[service]
		if !ok {
			continue
		}

		sender := CreateSender(service)
		sender.New(message, tokens)
		sender.Send()
		sender.Clear()
	}
}
