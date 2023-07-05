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
	err := DB.Delete(&PushToken{}, "token IN ? and service = ?", s.ExpiredTokens, s.Service()).Error
	if err != nil {
		log.Err(err).Msg("delete expired tokens failed")
	} else {
		log.Info().
			Str("scope", s.Service().String()).
			Strs("expired_tokens", s.ExpiredTokens).
			Msg("delete expired tokens success")
	}
}

func (s *Sender) Service() PushService {
	return ServiceUnknown
}
