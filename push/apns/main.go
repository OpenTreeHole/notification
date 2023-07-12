package apns

import (
	"errors"
	"github.com/rs/zerolog/log"
	"github.com/sideshow/apns2"
	"notification/config"
	. "notification/models"
	"notification/push/base"
)

type Sender struct {
	base.Sender
}

func (s *Sender) Send() {
	for _, token := range s.Tokens {
		packageName := config.Config.IOSPackageName
		if token.PackageName != "" {
			packageName = token.PackageName
		}
		res, err := client.Push(&apns2.Notification{
			DeviceToken: token.Token,
			Topic:       packageName,
			Payload:     constructPayload(s.Message),
		})
		if err != nil || res == nil {
			log.Err(err).
				Str("scope", "APNs").
				Msgf("push error")
			return
		}
		if res.StatusCode != 200 {
			log.Err(errors.New("APNs push failed")).
				Str("scope", "APNs").
				Str("token", token.Token).
				Int("status", res.StatusCode).
				Str("reason", res.Reason).
				Msgf("APNs push failed")
			// see https://developer.apple.com/documentation/usernotifications/setting_up_a_remote_notification_server/handling_notification_responses_from_apns
			if res.StatusCode == 410 { // expired or unregistered
				// device token is expired, remove it from database
				s.ExpiredTokens = append(s.ExpiredTokens, token.Token)
			}
		}
	}
}

func constructPayload(message *Message) any {
	return Map{"aps": Map{"alert": Map{
		"title":    message.Title,
		"subtitle": message.Description,
		"body":     message.Data,
	}}}
}
