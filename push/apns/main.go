package apns

import (
	"fmt"
	"github.com/sideshow/apns2"
	"notification/config"
	. "notification/models"
	"notification/push/base"
	"notification/utils"
	"strings"
)

type Sender struct {
	base.Sender
}

func (s *Sender) Send() bool {
	var success = true

	for _, token := range s.Tokens {
		res, err := client.Push(&apns2.Notification{
			DeviceToken: token,
			Topic:       config.Config.IOSPackageName,
			Payload:     constructPayload(s.Message),
		})
		if err != nil {
			utils.Logger.Error("APNS push error: " + err.Error())
			success = false
		}
		if res.StatusCode != 200 {
			utils.Logger.Warn(fmt.Sprintf(
				"APNS push failed: %d %s",
				res.StatusCode, res.Reason,
			))
			if strings.Contains(res.Reason, "DeviceToken") {
				// device token is expired, remove it from database
				s.ExpiredTokens = append(s.ExpiredTokens, token)
			}
			success = false
		}
		utils.Logger.Debug("APNS push success for " + token)
	}

	return success
}

func constructPayload(message *Message) any {
	return Map{"aps": Map{"alert": Map{
		"title":    message.Title,
		"subtitle": message.Description,
		"body":     message.Data,
	}}}

}
