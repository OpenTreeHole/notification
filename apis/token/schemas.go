package token

type CreateModel struct {
	Service  string `json:"service" validate:"required,oneof=apns fcm mipush"`
	DeviceID string `json:"device_id" validate:"required,max=64"`
	Token    string `json:"token" validate:"required,max=64"`
}

type DeleteModel struct {
	DeviceID string `json:"device_id" validate:"max=64"`
}
