package mipush

import (
	"fmt"
	"github.com/goccy/go-json"
	"log"
	"net/http"
	"net/url"
	"notification/config"
	. "notification/models"
	"notification/push/base"
	"strings"
)

type Sender struct {
	base.Sender
	Response Map
}

func (s *Sender) Send() bool {
	payload, _ := json.Marshal(&Map{
		"data": s.Message.Data,
		"code": s.Message.Type,
		"url":  s.Message.URL,
	})
	data := map[string]string{
		"registration_id":         strings.Join(s.Tokens, ","),
		"restricted_package_name": config.Config.AndroidPackageName,
		"title":                   s.Message.Title,
		"description":             s.Message.Description,
		"payload":                 url.QueryEscape(string(payload)),
		"extra.notify_effect":     "1",
	}
	form := url.Values{}
	for k, v := range data {
		form.Add(k, v)
	}
	req, err := http.NewRequest(
		"POST",
		mipushURL,
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		log.Println("mipush request error: " + err.Error())
	}
	req.Header.Set("Authorization", authorization)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)

	if err != nil {
		log.Printf("error sending mipush: %s\n", err)
		return false
	} else {
		s.Response = readBody(resp.Body)
		if resp.StatusCode != 200 || s.getStatusCode() != 0 {
			log.Println("failed sending mipush")
			fmt.Println(s.Response)
			return false
		}
	}
	return true
}

func (s *Sender) Clear() {
	s.ExpiredTokens = getExpiredTokens(s.Response)
	s.Sender.Clear()
}
