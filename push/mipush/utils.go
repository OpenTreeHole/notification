package mipush

import (
	"github.com/goccy/go-json"
	"github.com/rs/zerolog/log"
	"io"
	. "notification/models"
)

func readBody(body io.ReadCloser) Map {
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(body)

	data, err := io.ReadAll(body)
	if err != nil {
		log.Err(err).Msgf("Read body failed")
		return Map{}
	}
	var response Map
	err = json.Unmarshal(data, &response)
	if err != nil {
		log.Err(err).Msgf("Unmarshal body failed")
		return Map{}
	}
	return response
}

func (s *Sender) getStatusCode() int {
	if s.Response == nil {
		return -1
	}
	code, ok := s.Response["code"].(float64)
	if !ok {
		return -1
	}
	return int(code)
}
