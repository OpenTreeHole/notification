package models

import (
	"gorm.io/datatypes"
)

type Message struct {
	BaseModel
	UserID  int            `json:"user_id" gorm:"index;not null"`
	Message string         `json:"message" gorm:"size:64;not null"`
	Data    datatypes.JSON `json:"data" gorm:"not null"`
}
