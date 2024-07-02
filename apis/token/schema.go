package token

type CreateTokenRequest struct {
	DeviceID    string `json:"device_id" validate:"max=64"`
	Service     string `json:"service" validate:"required,oneof=apns fcm mipush"`
	Token       string `json:"token" validate:"max=256"`
	PackageName string `json:"package_name"`
}
