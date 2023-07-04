package tests

import (
	. "github.com/opentreehole/go-common"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndex(t *testing.T) {
	DefaultTester.Get(t, RequestConfig{Route: "/", ExpectedStatus: 302})
	DefaultTester.Get(t, RequestConfig{Route: "/api", ExpectedStatus: 200})

	var data struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	}
	DefaultTester.Get(t, RequestConfig{Route: "/404", ExpectedStatus: 404, ResponseModel: &data})

	log.Debug().Any("data", data).Msg("data")
	assert.EqualValues(t, "Cannot GET /404", data.Message)
}

func TestDocs(t *testing.T) {
	DefaultTester.Get(t, RequestConfig{Route: "/docs", ExpectedStatus: 302})
	DefaultTester.Get(t, RequestConfig{Route: "/docs/index.html", ExpectedStatus: 200})
}
