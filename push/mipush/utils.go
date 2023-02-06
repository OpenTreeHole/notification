package mipush

import (
	"github.com/goccy/go-json"
	"io"
	"log"
	. "notification/models"
	"strings"
)

func getExpiredTokens(resp Map) []string {
	data, ok := resp["data"].(Map)
	if !ok {
		return []string{}
	}
	tokens, ok := data["bad_regids"].(string)
	if !ok {
		return []string{}
	}
	return strings.Split(tokens, ",")
}

func readBody(body io.ReadCloser) Map {
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(body)

	data, err := io.ReadAll(body)
	if err != nil {
		log.Printf("Read body failed: %s", err)
		return Map{}
	}
	var response Map
	err = json.Unmarshal(data, &response)
	if err != nil {
		log.Printf("Unmarshal body failed: %s", err)
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
