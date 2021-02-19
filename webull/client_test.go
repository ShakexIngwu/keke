package webull

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	model "tradeserver/webullmodel"
)

func TestWebsocketConnect(t *testing.T) {
	if os.Getenv("WEBULL_USERNAME") == "" {
		t.Skip("No username set")
		return
	}
	asrt := assert.New(t)
	c, err := NewClient(&Credentials{
		Username:    os.Getenv("WEBULL_USERNAME"),
		Password:    os.Getenv("WEBULL_PASSWORD"),
		AccountType: model.AccountType(2),
		DeviceName:  deviceName(),
	})
	asrt.Empty(err)
	asrt.NotNil(c)
	err = c.ConnectWebsockets([]string{"913256135", "913256136"})
	asrt.Empty(err)
}
