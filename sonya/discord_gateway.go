package sonya

import "encoding/json"

// GetGateway returns an object with a single valid WSS URL,
// which the client can use for Connecting.
// Clients should cache this value and only call this endpoint to retrieve a new URL
// if they are unable to properly establish a connection using the cached version of the URL.
//
// https://discord.com/developers/docs/topics/gateway#get-gateway

func (s *Discord) GetGateway() (*GetGatewayResponse, error) {
	url := new(GetGatewayResponse)
	return url, s.get("/gateway", url)
}

// GetGatewayBot returns an object based on the information in Get Gateway,
// plus additional metadata that can help during the operation of large or sharded bots.
// Unlike the Get Gateway,
// this route should not be cached for extended periods of time as the value is not guaranteed to be the same per-call,
// and changes as the bot joins/leaves guilds.
//
// https://discord.com/developers/docs/topics/gateway#get-gateway-bot

func (d *Discord) GetGatewayBot() (*GetGatewayResponse, error) {
	url := new(GetGatewayResponse)
	return url, d.get("/gateway/bot", url)
}

type GetGatewayResponse struct {
	URL               string `json:"url"`
	Shards            int    `json:"shards,omitempty"`
	SessionStartLimit *struct {
		Total          int `json:"total"`
		Remaining      int `json:"remaining"`
		ResetAfter     int `json:"reset_after"`
		MaxConcurrency int `json:"max_concurrency"`
	} `json:"session_start_limit,omitempty"`
}

func (d *Discord) dispatch(typ string, in []byte) error {
	switch typ {
	case "READY":
		data := new(EventReady)
		if err := json.Unmarshal(in, data); err != nil {
			return err
		}
		for _, h := range d.handlers.Ready {
			go h(d, data)
		}
	case "MESSAGE_CREATE":
		data := new(EventMessageCreate)
		if err := json.Unmarshal(in, data); err != nil {
			return err
		}
		for _, h := range d.handlers.MessageCreate {
			go h(d, data)
		}
	}
	return nil
}
