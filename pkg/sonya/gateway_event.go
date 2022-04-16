package sonya

// EventReady is dispatched when a client has completed the initial handshake with the gateway (for new sessions).
// The ready event can be the largest and most complex event the gateway will send,
// as it contains all the state required for a client to begin interacting with the rest of the platform.
//
// https://discord.com/developers/docs/topics/gateway#ready-ready-event-fields
type EventReady struct {
	V           int                `json:"v"`
	User        User               `json:"user"`
	Guilds      []UnavailableGuild `json:"guilds"`
	SessionID   string             `json:"session_id"`
	Shard       *[2]int            `json:"shard"`
	Application Application        `json:"application"`
}

type EventMessageCreate Message
