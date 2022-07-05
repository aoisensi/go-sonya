package sonya

type Role struct {
	ID    RoleID `json:"id"`
	Name  string `json:"name"`
	Color int    `json:"color"`
	Hoist bool   `json:"hoist"`
	// Icon TODO
	// UnicodeEmoji TODO
	Position    int    `json:"position"`
	Permissions string `json:"permissions"`
	Managed     bool   `json:"managed"`
	Mentionable bool   `json:"mentable"`
	// Tags TODO
}
