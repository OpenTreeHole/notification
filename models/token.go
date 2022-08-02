package models

type PushService uint8

const (
	ServiceAPNS PushService = iota
	ServiceFCM
	ServiceMipush
)

func (s PushService) String() string {
	switch s {
	case ServiceAPNS:
		return "APNS"
	case ServiceFCM:
		return "FCM"
	case ServiceMipush:
		return "MiPush"
	default:
		return "Unknown"
	}
}

type PushToken struct {
	ID       int         `json:"id" gorm:"primarykey"`
	UserID   int         `json:"user_id" gorm:"index;not null"`
	Service  PushService `json:"service" gorm:"not null"`
	DeviceID string      `json:"device_id" gorm:"size:64;not null"`
	Token    string      `json:"token" gorm:"size:64;not null"`
}
