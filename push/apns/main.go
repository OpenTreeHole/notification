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
	var err error

	for _, token := range s.Tokens {
		packageName := config.Config.IOSPackageName
		if token.PackageName != "" {
			packageName = token.PackageName
		}

		notification := apns2.Notification{
			DeviceToken: token.Token,
			Topic:       packageName,
			Payload:     constructPayload(s.Message),
		}

		var (
			res           *apns2.Response
			currentClient *apns2.Client
		)

		if packageName == config.Config.IOSPackageNameV2 {
			currentClient = clientV2
		} else {
			currentClient = client
		}

		if currentClient == nil {
			log.Err(err).
				Str("scope", "APNs").
				Str("package_name", packageName).
				Msgf("client is nil")
			return
		}

		res, err = currentClient.Push(&notification)

		if err != nil || res == nil {
			log.Err(err).
				Str("scope", "APNs").
				Str("package_name", packageName).
				Msgf("push error")
			return
		}
		if res.StatusCode != 200 {
			log.Err(errors.New("APNs push failed")).
				Str("scope", "APNs").
				Str("package_name", packageName).
				Str("token", token.Token).
				Int("status", res.StatusCode).
				Str("reason", res.Reason).
				Msgf("APNs push failed")
			// see https://developer.apple.com/documentation/usernotifications/setting_up_a_remote_notification_server/handling_notification_responses_from_apns
			if res.StatusCode == 410 || // expired or unregistered device token
				(res.StatusCode == 400 && res.Reason == "BadDeviceToken") { // bad device token
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
