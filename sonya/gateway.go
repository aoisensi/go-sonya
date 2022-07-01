package sonya

import (
	"encoding/json"
	"net/url"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Connect to the gateway
func (d *Discord) Connect() error {
	var resp *GetGatewayResponse
	var err error
	if d.isBot {
		resp, err = d.GetGatewayBot()
	} else {
		resp, err = d.GetGateway()
	}
	if err != nil {
		return err
	}
	url := resp.URL + "?" + url.Values{
		"v":        {strconv.Itoa(d.APIVersion)},
		"encoding": {"json"},
	}.Encode()
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return err
	}
	g := &gateway{ws: ws, discord: d}
	data, _, err := g.read()
	if err != nil {
		return err
	}
	hello, ok := data.(*plHello)
	if !ok {
		panic("invalid hello")
	}
	// wait random time
	time.Sleep(time.Second)
	g.sendHeartbeat()
	go g.heart(hello.HeartbeatInterval)

	data, _, err = g.read()
	if err != nil {
		return err
	}
	_, ok = data.(*plHeartbeatACK)
	if !ok {
		panic("invalid hello")
	}

	identify := &plIdentify{
		Token:   d.token,
		Intents: 513,
	}
	identify.Properties.OS = runtime.GOOS
	identify.Properties.Browser = "sonya"
	identify.Properties.Device = "sonya"
	g.send(&gwPayload{
		Opcode: 2, // Identify
		Data:   identify,
	})
	d.gateway = g
	return g.recieve()
}

func (g *gateway) close() error {
	err := g.ws.Close()
	g.ws = nil
	return err
}

func (g *gateway) heart(interval time.Duration) {
	for {
		time.Sleep(interval)
		if g.ws == nil {
			return
		}
		g.sendHeartbeat()
	}
}

func (g *gateway) recieve() error {
	for {
		data, typ, err := g.read()
		if err != nil {
			return err
		}
		switch data := data.(type) {
		case *plHeartbeat:
			if err := g.sendHeartbeat(); err != nil {
				return err
			}
			continue
		case json.RawMessage:
			err := g.discord.dispatch(typ, data)
			if err != nil {
				return err
			}
			continue
		}
	}
}

func (g *gateway) read() (any, string, error) {
	pl := new(gwPayload)
	if err := g.ws.ReadJSON(pl); err != nil {
		return nil, "", err
	}
	if pl.SequenceNumber != nil {
		g.seq = pl.SequenceNumber
	}
	switch pl.Opcode {
	case 0: // Dispatch
		return pl.DataRaw, *pl.Type, nil
	case 1: // Heartbeat
		data := new(plHeartbeat)
		return data, "", json.Unmarshal(pl.DataRaw, data)
	case 10: // Hello
		data := new(plHello)
		return data, "", json.Unmarshal(pl.DataRaw, &data)
	case 11: // Heartbeat ACK
		data := new(plHeartbeatACK)
		return data, "", json.Unmarshal(pl.DataRaw, data)
	default:
		panic("invalid opcode")
	}
}

func (g *gateway) sendHeartbeat() error {
	return g.send(&gwPayload{
		Opcode:         1,
		SequenceNumber: g.seq,
	})
}

func (g *gateway) send(pl *gwPayload) error {
	if pl.DataRaw == nil {
		data, err := json.Marshal(pl.Data)
		if err != nil {
			return err
		}
		pl.DataRaw = data
	}
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.ws.WriteJSON(pl)
}

type gateway struct {
	ws      *websocket.Conn
	mu      sync.Mutex
	seq     *int
	discord *Discord
}
