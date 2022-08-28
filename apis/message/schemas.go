package message

import (
	"github.com/creasty/defaults"
	. "notification/models"
)

type CreateModel struct {
	// message type, change "oneof" when MessageType changes
	Type        MessageType `json:"type" validate:"required,oneof=favorite reply mention modify report permission report_dealt"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Data        JSON        `json:"data"`
	URL         string      `json:"url"`
	Recipients  []int       `json:"recipients" validate:"required"`
}

func (body *CreateModel) SetDefaults() {
	if defaults.CanUpdate(body.Data) {
		body.Data = JSON{}
	}
}
