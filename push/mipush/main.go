package mipush

import (
	"github.com/goccy/go-json"
	"github.com/opentreehole/go-common"
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

	// https://dev.mi.com/distribute/doc/details?pId=1559
	data := map[string]string{
		// 根据registration_id，发送消息到指定设备上。
		// 可以提供多个registration_id，发送给一组设备，不同的registration_id之间用“,”分割。
		"registration_id": strings.Join(s.Tokens, ","),

		// App的包名。V3版本支持多包名（中间用逗号分割）
		"restricted_package_name": config.Config.AndroidPackageName,

		// 通知栏展示的通知的标题，不允许全是空白字符，长度小于50， 一个中英文字符均计算为1（通知栏消息必填）
		"title": common.StripContent(s.Message.Title, 49),

		// 通知栏展示的通知的描述，不允许全是空白字符，长度小于128，一个中英文字符均计算为1（通知栏消息必填）。
		"description": common.StripContent(s.Message.Description, 127),

		// 消息的内容。（注意：需要对payload字符串做urlencode处理）
		"payload": url.QueryEscape(string(payload)),

		// 可选项，预定义通知栏消息的点击行为。通过设置extra.notify_effect的值以得到不同的预定义点击行为。
		// “1″：通知栏点击后打开app的Launcher Activity。
		// “2″：通知栏点击后打开app的任一Activity（开发者还需要传入extra.intent_uri）。
		// “3″：通知栏点击后打开网页（开发者还需要传入extra.web_uri）。
		"extra.notify_effect": "1",

		// 可选项，开启消息回执。消息发送后，推送系统能发送回执给开发者，告知开发者这些消息的送达、点击或发送失败状态。
		// 将extra.callback的值设置为第三方接收回执的http接口。（注意：仅支持http协议，不支持https协议）
		"extra.callback": config.Config.MipushCallbackUrl,

		// 可选项，表示回执类型。详细用法请参见《服务端Java SDK文档》中“消息回执”一节中的callback.type字段。
		"extra.callback.type": "19",
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
