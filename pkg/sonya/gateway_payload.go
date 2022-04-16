package sonya

import (
	"encoding/json"
)

type gwPayload struct {
	Opcode         int             `json:"op"`
	DataRaw        json.RawMessage `json:"d"`
	SequenceNumber *int            `json:"s,omitempty"`
	Type           *string         `json:"t,omitempty"`
	Data           interface{}     `json:"-"`
}

type plHello struct {
	HeartbeatInterval int `json:"heartbeat_interval"`
}

type plIdentify struct {
	Token      string `json:"token"`
	Intents    int    `json:"intents"`
	Properties struct {
		OS      string `json:"$os"`
		Browser string `json:"$browser"`
		Device  string `json:"$device"`
	} `json:"properties"`
}

type plHeartbeat struct {
}

type plHeartbeatACK struct {
}
