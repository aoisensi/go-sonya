package sonya

import (
	"net/http"
)

type Discord struct {
	BaseURL    string
	APIVersion int
	HTTPClient *http.Client
	token      string
	isBot      bool
}

func NewBot(token string) *Discord {
	return New("Bot "+token, true)
}

func NewBearer(token string) *Discord {
	return New("Bearer "+token, false)
}

func New(token string, isBot bool) *Discord {
	return &Discord{
		BaseURL:    "https://discordapp.com/api",
		APIVersion: 9,
		HTTPClient: http.DefaultClient,
		token:      token,
		isBot:      isBot,
	}
}
