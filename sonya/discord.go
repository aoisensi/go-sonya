package sonya

import (
	"net"
	"net/http"
)

type Discord struct {
	BaseURL        string
	APIVersion     int
	HTTPClient     *http.Client
	token          string
	isBot          bool
	gateway        *gateway
	hReady         []func(*Discord, *EventReady)
	hMessageCreate []func(*Discord, *EventMessageCreate)
}

func NewBot(token string) *Discord {
	return New("Bot "+token, true)
}

func NewBearer(token string) *Discord {
	return New("Bearer "+token, false)
}

func New(authorization string, isBot bool) *Discord {
	return &Discord{
		BaseURL:        "https://discordapp.com/api",
		APIVersion:     9,
		HTTPClient:     http.DefaultClient,
		token:          authorization,
		isBot:          isBot,
		hReady:         make([]func(*Discord, *EventReady), 0, 4),
		hMessageCreate: make([]func(*Discord, *EventMessageCreate), 0, 4),
	}
}

func (d *Discord) Close() error {
	if d.gateway == nil {
		return net.ErrClosed
	}
	return d.gateway.close()
}
