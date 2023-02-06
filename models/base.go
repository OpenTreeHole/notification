// Package models contains database models
package models

type Map = map[string]any

type MessageModel struct {
	Message string `json:"message"`
}

type Models interface {
	PushToken | []PushToken
}
