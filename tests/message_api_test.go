package tests

import (
	"github.com/stretchr/testify/assert"
	. "notification/models"
	"strconv"
	"testing"
)

func init() {
	for i := 0; i < 12; i++ {
		message := Message{
			Message: "message " + strconv.Itoa(i),
			Type:    MessageTypeFavorite,
			Data:    []byte("{}"),
		}
		userIDs := make([]int, 1)
		userIDs[0] = i%3 + 1
		err := message.Create(userIDs)
		if err != nil {
			panic(err)
		}
	}
}

func TestGetMessage(t *testing.T) {
	messages := testAPIModel[[]Message](t, "get", "/api/messages", 200)
	assert.GreaterOrEqual(t, 5, len(messages))
}

func TestDeleteMessage(t *testing.T) {
	userIDMockLock.Lock()
	defer resetUserIDMock()
	userIDMock = 2
	messageID := userIDMock

	testCommon(t, "delete", "/api/messages/"+strconv.Itoa(messageID), 204)
	assert.NotNilf(
		t,
		DB.
			Where("user_id = ?", userIDMock).
			Where("message_id = ?", messageID).
			First(&MessageUser{}).Error,
		"TestDeleteMessage",
	)
}

func TestClearMessage(t *testing.T) {
	userIDMockLock.Lock()
	defer resetUserIDMock()
	userIDMock = 3

	testCommon(t, "post", "/api/messages/clear", 204)
	var messages []MessageUser
	DB.Where("user_id = ?", userIDMock).Find(&messages)
	assert.Equalf(t, 0, len(messages), "TestClearMessage")
}
