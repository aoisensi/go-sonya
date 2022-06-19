package sonya

func (d *Discord) AddReadyHandler(f func(*Discord, *EventReady)) {
	d.hReady = append(d.hReady, f)
}

func (d *Discord) AddMessageCreateHandler(f func(*Discord, *EventMessageCreate)) {
	d.hMessageCreate = append(d.hMessageCreate, f)
}
