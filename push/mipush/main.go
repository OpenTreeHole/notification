package mipush

import (
	"github.com/goccy/go-json"
	"github.com/rs/zerolog/log"
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

func (s *Sender) Send() {
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
		"extra.callback":          config.Config.MipushCallbackUrl,
		"extra.callback.type":     "19",
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
		log.Err(err).Msg("mipush request error")
		return
	}
	req.Header.Set("Authorization", authorization)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)

	if err != nil {
		log.Err(err).Msg("error sending mipush")
		return
	}

	s.Response = readBody(resp.Body)
	if resp.StatusCode != 200 || s.getStatusCode() != 0 {
		log.Error().Any("response", s.Response).Msg("failed sending mipush")
	}
}

func (s *Sender) Service() PushService {
	return ServiceMipush
}
