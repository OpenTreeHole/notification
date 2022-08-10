// Package models contains database models
package models

import (
	"database/sql/driver"
	"github.com/goccy/go-json"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
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
	PushToken | Message |
		[]PushToken | []Message
}

type JSON map[string]any

func (t JSON) Value() (driver.Value, error) {
	return json.Marshal(t)
}

func (t *JSON) Scan(input any) error {
	return json.Unmarshal(input.([]byte), t)
}

// GormDataType gorm common data type
func (JSON) GormDataType() string {
	return "json"
}

// GormDBDataType gorm db data type
//goland:noinspection GoUnusedParameter
func (JSON) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite":
		return "JSON"
	case "mysql":
		return "JSON"
	case "postgres":
		return "JSONB"
	}
	return ""
}
