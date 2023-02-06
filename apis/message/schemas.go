package message

import (
	. "notification/models"

	"github.com/creasty/defaults"
)

type CreateModel struct {
	Type        string `json:"type" validate:"required"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Data        JSON   `json:"data"`
	URL         string `json:"url"`
	Recipients  []int  `json:"recipients" validate:"required"`
}

type ListModel struct {
	NotRead bool `default:"false" query:"not_read"`
}

func (body *CreateModel) SetDefaults() {
	if defaults.CanUpdate(body.Data) {
		body.Data = JSON{}
	}
}
