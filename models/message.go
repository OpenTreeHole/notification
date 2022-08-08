package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Message struct {
	BaseModel
	Type    MessageType    `json:"code" gorm:"size:16;not null"`
	Message string         `json:"message" gorm:"size:64;not null"`
	Data    datatypes.JSON `json:"data" gorm:"not null"`
}

type MessageUser struct {
	MessageID int `json:"message_id" gorm:"primaryKey"`
	UserID    int `json:"user_id" gorm:"primaryKey"`
}

type MessageType string

const (
	MessageTypeFavorite   MessageType = "favorite"
	MessageTypeReply      MessageType = "reply"
	MessageTypeMention    MessageType = "mention"
	MessageTypeModify     MessageType = "modify" // including fold and delete
	MessageTypeReport     MessageType = "report"
	MessageTypePermission MessageType = "permission"
)

func (m *Message) Create(userIDs []int) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(m).Error; err != nil {
			return err
		}

		mapping := make([]MessageUser, len(userIDs))
		for i, userID := range userIDs {
			mapping[i] = MessageUser{
				MessageID: m.ID,
				UserID:    userID,
			}
		}

		if err := tx.Create(&mapping).Error; err != nil {
			return err
		}

		return nil
	})
}
