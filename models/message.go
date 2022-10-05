package models

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Messages []Message

type Message struct {
	BaseModel
	Title       string      `json:"message" gorm:"size:32;not null"`
	Description string      `json:"description" gorm:"size:64;not null"`
	Data        JSON        `json:"data"`
	Type        MessageType `json:"code" gorm:"size:16;not null"`
	URL         string      `json:"url" gorm:"size:64;default:'';not null"`
	Recipients  []int       `json:"-" gorm:"-:all" `
	MessageID   int         `json:"message_id" gorm:"-:all"`       // 兼容旧版 id
	HasRead     bool        `json:"has_read" gorm:"default:false"` // 兼容旧版
}

type MessageUser struct {
	MessageID int `json:"message_id" gorm:"primaryKey"`
	UserID    int `json:"user_id" gorm:"primaryKey"`
}

type MessageType string

const (
	MessageTypeFavorite    MessageType = "favorite"
	MessageTypeReply       MessageType = "reply"
	MessageTypeMention     MessageType = "mention"
	MessageTypeModify      MessageType = "modify" // including fold and delete
	MessageTypePermission  MessageType = "permission"
	MessageTypeReport      MessageType = "report"
	MessageTypeReportDealt MessageType = "report_dealt"
)

func (messages Messages) Preprocess(c *fiber.Ctx) error {
	for i := 0; i < len(messages); i++ {
		err := messages[i].Preprocess(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func (message *Message) Preprocess(c *fiber.Ctx) error {
	message.MessageID = message.ID
	return nil
}

func (m *Message) AfterCreate(tx *gorm.DB) (err error) {
	mapping := make([]MessageUser, len(m.Recipients))
	for i, userID := range m.Recipients {
		mapping[i] = MessageUser{
			MessageID: m.ID,
			UserID:    userID,
		}
	}
	return tx.Create(&mapping).Error
}
