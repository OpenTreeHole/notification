package models

import (
	"errors"
	"strings"
	"time"
)

type PushToken struct {
	UserID    int         `json:"user_id" gorm:"primaryKey;not null"` // not required
	Service   PushService `json:"service" gorm:"not null"`
	DeviceID  string      `json:"device_id" gorm:"uniqueIndex:,length:10;not null;size:64" validate:"required,max=64"`
	Token     string      `json:"token" gorm:"primaryKey;not null;size:64;index:,length:10" validate:"required,max=64"`
	CreatedAt time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
}

type PushService uint8

const (
	ServiceAPNS PushService = iota
	ServiceFCM
	ServiceMipush
)

var PushServices = []PushService{ServiceAPNS, ServiceFCM, ServiceMipush}

func (s PushService) MarshalText() ([]byte, error) {
	return []byte(s.String()), nil
}

func (s PushService) String() string {
	switch s {
	case ServiceAPNS:
		return "apns"
	case ServiceFCM:
		return "fcm"
	case ServiceMipush:
		return "mipush"
	default:
		return "unknown"
	}
}

func (s *PushService) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {
	case "apns":
		*s = ServiceAPNS
	case "fcm":
		*s = ServiceFCM
	case "mipush":
		*s = ServiceMipush
	default:
		return errors.New("unknown push service")
	}
	return nil
}
