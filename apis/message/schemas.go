package message

import (
	"github.com/creasty/defaults"
	"gorm.io/datatypes"
	. "notification/models"
)

type CreateModel struct {
	// message type, change "oneof" when MessageType changes
	Type        MessageType    `json:"type" validate:"required,oneof=favorite reply mention modify report permission"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Data        datatypes.JSON `json:"data"`
	Recipients  []int          `json:"recipients" validate:"required"`
}

func (body *CreateModel) SetDefaults() {
	if defaults.CanUpdate(body.Data) {
		body.Data = []byte("{}")
	}
}
