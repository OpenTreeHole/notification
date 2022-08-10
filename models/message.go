package models

import (
	"gorm.io/gorm"
)

type Message struct {
	BaseModel
	Title       string      `json:"message" gorm:"size:32;not null"`
	Description string      `json:"description" gorm:"size:64;not null"`
	Data        JSON        `json:"data"`
	Type        MessageType `json:"code" gorm:"size:16;not null"`
	URL         string      `json:"url" gorm:"size:64;default:'';not null"`
	Recipients  []int       `json:"-" gorm:"-:all" `
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
