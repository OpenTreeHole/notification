package tests

import (
	"github.com/stretchr/testify/assert"
	. "notification/models"
	"strconv"
	"testing"
)

func init() {
	for i := 0; i < 12; i++ {
		userIDs := make([]int, 1)
		userIDs[0] = i%3 + 1
		message := Message{
			Title:      "message " + strconv.Itoa(i),
			Type:       MessageTypeFavorite,
			Data:       Map{},
			Recipients: userIDs,
		}
		err := DB.Create(&message).Error
		if err != nil {
			panic(err)
		}
	}
}

func TestGetMessage(t *testing.T) {
	messages := testAPIModel[[]Message](t, "get", "/api/messages", 200)
	assert.GreaterOrEqual(t, 5, len(messages))
}

func TestAddMessage(t *testing.T) {
	recipients := []int{1, 2}
	data := Map{"type": MessageTypeReply, "recipients": recipients}
	message := testAPIModel[Message](t, "post", "/api/messages", 201, data)
	assert.Equal(t, MessageTypeReply, message.Type)
	// don't return recipients
	//assert.Equal(t, nil, message.Recipients)
	assert.Equal(t, "你的帖子被回复了", message.Title)
	assert.Equal(t, "", message.Description)

	// test no recipients
	data = Map{"type": MessageTypeFavorite}
	testCommon(t, "post", "/api/messages", 400, data)

	// test no type
	data = Map{"recipients": recipients}
	testCommon(t, "post", "/api/messages", 400, data)
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
