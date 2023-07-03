package base

import (
	"github.com/rs/zerolog/log"
	. "notification/models"
)

// Sender is a base struct for Sender.
type Sender struct {
	Message       *Message
	Tokens        []string
	ExpiredTokens []string
}

// New initializes a Sender.
func (s *Sender) New(message *Message, tokens []string) {
	s.Message = message
	s.Tokens = tokens
	s.Message.Data["url"] = s.Message.URL
}

// Send sends notification.
func (s *Sender) Send() {}

// Clear expired tokens.
func (s *Sender) Clear() {
	if len(s.ExpiredTokens) == 0 {
		return
	}
	err := DB.Exec(
		"DELETE FROM push_token WHERE token IN (?)",
		s.ExpiredTokens,
	).Error
	if err != nil {
		log.Err(err).Msg("Delete expired tokens failed")
	}
}
