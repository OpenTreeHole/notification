package models

type Messages []Message

type Message struct {
	Title       string `json:"message" validate:"required"`
	Description string `json:"description" validate:"required"`
	Data        Map    `json:"data" validate:"required"`
	Type        string `json:"code" validate:"required"`
	URL         string `json:"url" validate:"required"`
	Recipients  []int  `json:"recipients" validate:"required"`
}
