package sonya

import (
	"encoding/json"
	"fmt"
	"net/url"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/gollira/websocket"
)

// Connect to the gateway
func (d *Discord) Connect() (*Gateway, error) {
	var resp *GetGatewayResponse
	var err error
	if d.isBot {
		resp, err = d.GetGatewayBot()
	} else {
		resp, err = d.GetGateway()
	}
	if err != nil {
		return nil, err
	}
	url := resp.URL + "?" + url.Values{
		"v":        {strconv.Itoa(d.APIVersion)},
		"encoding": {"json"},
	}.Encode()
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}
	g := &Gateway{ws: ws}
	data, _, err := g.read()
	if err != nil {
		return nil, err
	}
	hello, ok := data.(*plHello)
	if !ok {
		panic("invalid hello")
	}

	go g.heart(time.Duration(hello.HeartbeatInterval) * time.Millisecond)

	data, _, err = g.read()
	if err != nil {
		return nil, err
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
	go g.recieve()
	return g, nil
}

func (g *Gateway) Close() error {
	err := g.ws.Close()
	g.ws = nil
	return err
}

func (g *Gateway) heart(interval time.Duration) {
	time.Sleep(interval * 60 / time.Duration(time.Now().Second()+1))
	for {
		if g.ws == nil {
			return
		}
		g.sendHeartbeat()
		time.Sleep(interval)
	}
}

func (g *Gateway) recieve() {
	for {
		if g.ws == nil {
			return
		}
		data, typ, err := g.read()
		if err != nil {
			return
		}
		switch data.(type) {
		case *plHeartbeat:
			g.sendHeartbeat()
			continue
		case *json.RawMessage:
			fmt.Println(typ)
			continue
		}
	}
}

func (g *Gateway) read() (interface{}, string, error) {
	pl := new(gwPayload)
	if err := g.ws.ReadJSON(pl); err != nil {
		return nil, "", err
	}
	fmt.Println("<<<", pl.Opcode)
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

func (g *Gateway) sendHeartbeat() error {
	return g.send(&gwPayload{
		Opcode:         1,
		SequenceNumber: g.seq,
	})
}

func (g *Gateway) send(pl *gwPayload) error {
	if pl.DataRaw == nil {
		data, err := json.Marshal(pl.Data)
		if err != nil {
			return err
		}
		pl.DataRaw = data
	}
	g.mu.Lock()
	fmt.Println(">>>", pl.Opcode)
	defer g.mu.Unlock()
	return g.ws.WriteJSON(pl)
}

type Gateway struct {
	ws  *websocket.Conn
	mu  sync.Mutex
	seq *int
}
