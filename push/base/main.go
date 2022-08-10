package base

import (
	. "notification/models"
	. "notification/utils"
)

// Sender is a base struct for Sender.
//goland:noinspection GoNameStartsWithPackageName
type Sender struct {
	Message       *Message
	Tokens        []string
	ExpiredTokens []string
}

// New initializes a Sender.
func (s *Sender) New(message *Message, tokens []string) {
	s.Message = message
	s.Tokens = tokens
}

// Send sends notification.
func (s *Sender) Send() bool {
	return true
}

// Clear expired tokens.
func (s *Sender) Clear() {
	if len(s.ExpiredTokens) == 0 {
		Logger.Debug("No expired tokens, skip clear")
		return
	}
	err := DB.Exec(
		"DELETE FROM push_token WHERE token IN (?)",
		s.ExpiredTokens,
	).Error
	if err != nil {
		Logger.Error("Delete expired tokens failed: " + err.Error())
	}
}
