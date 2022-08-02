// Package models contains database models
package models

import (
	"notification/config"
	"time"
)

var DB = config.DB

type Map = map[string]any

type BaseModel struct {
	ID        int       `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"time_created"`
	UpdatedAt time.Time `json:"time_updated"`
}

func (model BaseModel) GetID() int {
	return model.ID
}

type Models interface {
}
