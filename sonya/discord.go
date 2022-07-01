package sonya

import (
	"net"
	"net/http"
)

type Discord struct {
	BaseURL    string
	APIVersion int
	HTTPClient *http.Client
	token      string
	isBot      bool
	gateway    *gateway
	handlers   *handlers
}

func New(token string, isBot bool) *Discord {
	return &Discord{
		BaseURL:    "https://discordapp.com/api",
		APIVersion: 9,
		HTTPClient: http.DefaultClient,
		token:      token,
		isBot:      isBot,
		handlers:   newHandlers(),
	}
}

func (d *Discord) Close() error {
	if d.gateway == nil {
		return net.ErrClosed
	}
	return d.gateway.close()
}
