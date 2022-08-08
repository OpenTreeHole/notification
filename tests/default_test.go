package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndex(t *testing.T) {
	testCommon(t, "get", "/", 302)
	testCommon(t, "get", "/api", 200)
	data := testAPI(t, "get", "/404", 404)
	assert.Equal(t, "Cannot GET /404", data["message"])
}

func TestDocs(t *testing.T) {
	testCommon(t, "get", "/docs", 302)
	testCommon(t, "get", "/docs/index.html", 200)
}
