package mipush

import (
	"github.com/goccy/go-json"
	"io"
	"io/ioutil"
	. "notification/models"
	. "notification/utils"
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
		err := body.Close()
		if err != nil {
			Logger.Error("Close error: " + err.Error())
		}
	}(body)

	data, err := ioutil.ReadAll(body)
	if err != nil {
		Logger.Error("Read body failed: " + err.Error())
		return Map{}
	}
	var response Map
	err = json.Unmarshal(data, &response)
	if err != nil {
		Logger.Error("Unmarshal body failed: " + err.Error())
		return Map{}
	}
	return response
}

func (s *Sender) getStatusCode() int {
	if s.Response == nil {
		return -1
	}
	code, ok := s.Response["code"].(int)
	if !ok {
		return -1
	}
	return code
}
