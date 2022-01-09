package sonya

import "net/http"

type Session struct {
	BaseURL       string
	APIVersion    string
	HTTPClient    *http.Client
	authorization string
}

func NewBot(token string) *Session {
	return New("Bot " + token)
}

func NewBearer(token string) *Session {
	return New("Bearer " + token)
}

func New(authorization string) *Session {
	return &Session{
		BaseURL:       "https://discordapp.com/api",
		APIVersion:    "v9",
		HTTPClient:    http.DefaultClient,
		authorization: authorization,
	}
}
