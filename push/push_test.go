package push

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	. "notification/models"
	"os"
	"strings"
	"testing"
)

const userID = 1

var serviceTokenMapping = make(map[PushService]string)

func deviceTokenEnvName(service PushService) string {
	return fmt.Sprintf("%s_DEVICE_TOKEN", strings.ToUpper(service.String()))
}

func init() {
	fmt.Println("init push test")
	for _, service := range PushServices {
		deviceToken := os.Getenv(deviceTokenEnvName(service))
		if deviceToken == "" {
			fmt.Println(
				deviceTokenEnvName(service),
				"not set, this service could not be tested",
			)
			continue
		}
		serviceTokenMapping[service] = deviceToken
		// create a test token
		DB.Create(&PushToken{
			UserID:   userID,
			Service:  service,
			DeviceID: "device_id",
			Token:    deviceToken,
		})
	}
}

func TestPushNotification(t *testing.T) {
	success := Send(Message{
		Type:        MessageTypeReply,
		Title:       "title",
		Description: "description",
		Data:        Map{},
		Recipients:  []int{userID},
	})
	assert.True(t, success)
}
