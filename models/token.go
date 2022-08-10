package models

type PushToken struct {
	ID       int         `json:"id" gorm:"primarykey"`
	UserID   int         `json:"user_id" gorm:"index;not null"`
	Service  PushService `json:"service" gorm:"not null"`
	DeviceID string      `json:"device_id" gorm:"size:64;not null"`
	Token    string      `json:"token" gorm:"size:64;not null"`
}

type PushService uint8

const (
	ServiceAPNS PushService = iota
	ServiceFCM
	ServiceMipush
)

var PushServices = []PushService{ServiceAPNS, ServiceFCM, ServiceMipush}

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

func ParsePushService(s string) PushService {
	switch s {
	case "apns":
		return ServiceAPNS
	case "fcm":
		return ServiceFCM
	case "mipush":
		return ServiceMipush
	default:
		return ServiceAPNS
	}
}

func (s PushService) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}
