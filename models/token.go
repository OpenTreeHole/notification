package models

import "strings"

type PushToken struct {
	UserID   int         `json:"user_id" gorm:"primaryKey;not null"` // not required
	Service  PushService `json:"service" gorm:"not null" validate:"required,oneof=apns fcm mipush"`
	DeviceID string      `json:"device_id" gorm:"primaryKey;not null" validate:"required,max=64"`
	Token    string      `json:"token" gorm:"not null" validate:"required,max=64"`
}

type PushService uint8

const (
	ServiceAPNS PushService = iota
	ServiceFCM
	ServiceMipush
)

var PushServices = []PushService{ServiceAPNS, ServiceFCM, ServiceMipush}

func (s PushService) MarshalText() ([]byte, error) {
	var name string
	switch s {
	case ServiceAPNS:
		name = "apns"
	case ServiceFCM:
		name = "fcm"
	case ServiceMipush:
		name = "mipush"
	default:
		name = "unknown"
	}
	return []byte(name), nil
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
		*s = ServiceAPNS
	}
	return nil
}
