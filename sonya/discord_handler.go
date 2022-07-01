package sonya

func (d *Discord) AddHandler(f any) {
	switch f := f.(type) {
	case func(*Discord, *EventReady):
		d.handlers.Ready = append(d.handlers.Ready, f)
	case func(*Discord, *EventMessageCreate):
		d.handlers.MessageCreate = append(d.handlers.MessageCreate, f)
	default:
		panic("sonya: this handler type is unsupported or invalid.")
	}
}

type handlers struct {
	Ready         []func(*Discord, *EventReady)
	MessageCreate []func(*Discord, *EventMessageCreate)
}

func newHandlers() *handlers {
	return &handlers{
		Ready:         make([]func(*Discord, *EventReady), 0),
		MessageCreate: make([]func(*Discord, *EventMessageCreate), 0),
	}

}
