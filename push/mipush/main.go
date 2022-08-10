package mipush

import (
	"fmt"
	"github.com/goccy/go-json"
	"net/http"
	"net/url"
	"notification/config"
	. "notification/models"
	"notification/push/base"
	. "notification/utils"
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
	req, _ := http.NewRequest(
		"POST",
		mipushURL,
		strings.NewReader(form.Encode()),
	)
	req.Header.Set("Authorization", authorization)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)

	s.Response = readBody(resp.Body)

	if err != nil {
		Logger.Error("error sending mipush" + err.Error())
		return false
	} else if resp.StatusCode != 200 || s.getStatusCode() != 0 {
		Logger.Warn("failed sending mipush")
		fmt.Println(s.Response)
		return false
	} else {
		Logger.Debug("mipush sent successfully")
		return true
	}
}

func (s *Sender) Clear() {
	s.ExpiredTokens = getExpiredTokens(s.Response)
	s.Sender.Clear()
}
