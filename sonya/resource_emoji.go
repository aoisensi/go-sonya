package sonya

import "fmt"

// https://discord.com/developers/docs/resources/emoji#emoji-object
type Emoji struct {
	ID            *EmojiID `json:"id"`
	Name          *string  `json:"name"`
	Roles         []RoleID `json:"roles"`
	User          *User    `json:"user"`
	RequireColons *bool    `json:"require_colons"`
	Managed       *bool    `json:"managed"`
	Animated      *bool    `json:"animated"`
	Available     *bool    `json:"available"`
}

// ListGuildEmojis returns a list of emoji objects for the given guild.
//
// https://discord.com/developers/docs/resources/emoji#list-guild-emojis
func (s *Discord) ListGuildEmojis(guildID GuildID) ([]*Emoji, error) {
	emojis := make([]*Emoji, 0, 200)
	return emojis, s.get(fmt.Sprintf("/guilds/%v/emojis", guildID), &emojis)
}

// GetGuildEmoji returns an emoji object for the given guild and emoji ID.
//
// https://discord.com/developers/docs/resources/emoji#get-guild-emoji
func (s *Discord) GetGuildEmoji(guildID GuildID, emojiID EmojiID) (*Emoji, error) {
	emoji := new(Emoji)
	return emoji, s.get(fmt.Sprintf("/guilds/%v/emojis/%v", guildID, emojiID), emoji)
}
