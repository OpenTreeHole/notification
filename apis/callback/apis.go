package callback

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/opentreehole/go-common"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	. "notification/models"
	"strings"
)

func RegisterRoutes(app fiber.Router) {
	app.Post("/callback/mipush", MipushCallback)
}

// MipushCallback
// @Summary Mipush Callback
// @Tags Callback
// @Produce application/json
// @Router /callback/mipush [post]
// @Success 200
func MipushCallback(c *fiber.Ctx) (err error) {
	dataString := c.FormValue("data")
	data := map[string]MipushCallbackData{}
	err = json.Unmarshal([]byte(dataString), &data)
	if err != nil {
		return err
	}

	// parse data
	expiredTokens := make([]string, 0, 10)
	replaceTokens := make(map[string]string)
	for _, v := range data {
		if v.Type == 16 {
			expiredTokens = append(expiredTokens, strings.Split(v.Targets, ",")...)
			for oldRegId, newRegId := range v.ReplaceTarget {
				replaceTokens[oldRegId] = newRegId
			}
		}
	}

	// replace token
	if len(replaceTokens) > 0 {
		replaceTokenKeys := common.Keys(replaceTokens)

		stringBuilder := strings.Builder{}
		stringBuilder.WriteString("UPDATE `push_token` SET `token` = CASE `token`")
		for k, v := range replaceTokens {
			stringBuilder.WriteString(" WHEN '")
			stringBuilder.WriteString(k)
			stringBuilder.WriteString("' THEN '")
			stringBuilder.WriteString(v)
			stringBuilder.WriteString("'")
		}
		stringBuilder.WriteString(" END WHERE `token` IN ('")
		stringBuilder.WriteString(strings.Join(replaceTokenKeys, "','"))
		stringBuilder.WriteString("') AND `service` = 2") // mipush

		logDict := zerolog.Dict()
		for k, v := range replaceTokens {
			logDict.Str(k, v)
		}

		err = DB.Exec(stringBuilder.String()).Error
		if err != nil {
			log.Err(err).Dict("replaced_tokens", logDict).Msg("replace token failed")
		} else {
			log.Info().Dict("replaced_tokens", logDict).Msg("replace token success")
		}
	}

	// delete expired token
	if len(expiredTokens) > 0 {
		err = DB.Where("token IN ? AND service = ?", expiredTokens, ServiceMipush).Delete(&PushToken{}).Error
		if err != nil {
			log.Err(err).Strs("expired_tokens", expiredTokens).Msg("delete token failed")
		} else {
			log.Info().Strs("expired_tokens", expiredTokens).Msg("delete token success")
		}
	}

	return c.SendStatus(200)
}
