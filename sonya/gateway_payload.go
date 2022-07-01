package sonya

import (
	"encoding/json"
	"time"
)

type gwPayload struct {
	Opcode         int             `json:"op"`
	DataRaw        json.RawMessage `json:"d"`
	SequenceNumber *int            `json:"s,omitempty"`
	Type           *string         `json:"t,omitempty"`
	Data           any             `json:"-"`
}

type plHello struct {
	HeartbeatInterval time.Duration `json:"heartbeat_interval"`
}

func (p *plHello) UnmarshalJSON(data []byte) error {
	j := &struct {
		HeartbeatInterval int64 `json:"heartbeat_interval"`
	}{}
	if err := json.Unmarshal(data, &j); err != nil {
		return err
	}
	p.HeartbeatInterval = time.Duration(j.HeartbeatInterval) * time.Millisecond
	return nil
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
