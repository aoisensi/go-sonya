package sonya

import "time"

// Channel
// Represents a guild or DM channel within Discord.
//
// https://discord.com/developers/docs/resources/channel#channel-object
type Channel struct {
	ID       ChannelID   `json:"id"`
	Type     ChannelType `json:"type"`
	GuildID  *GuildID    `json:"guild_id"`
	Position int         `json:"position"`
	// PermissionOverwrites // TODO
	Name             *string       `json:"name"`
	Topic            *string       `json:"topic"`
	NSFW             bool          `json:"nsfw"`
	LastMessageID    *MessageID    `json:"last_message_id"`
	Bitrate          int           `json:"bitrate"`
	UserLimit        int           `json:"user_limit"`
	RateLimitPerUser int           `json:"rate_limit_per_user"`
	Recipients       []*User       `json:"recipients"`
	Icon             *string       `json:"icon"`
	OwnerID          UserID        `json:"ownerId"`
	ApplicationID    ApplicationID `json:"application_id"`
	//keep going... TODO
}

// https://discord.com/developers/docs/resources/channel#channel-object-channel-types
type ChannelType int

const (
	ChannelTypeGuildText ChannelType = iota
	ChannelTypeDM
	ChannelTypeGuildVoice
	ChannelTypeGroupDM
	ChannelTypeGuildCategory
	ChannelTypeGuildNews
	_
	_
	_
	_
	ChannelTypeGuildNewsThread
	ChannelTypeGubldPublicThread
	ChannelTypeGuildPrivateThread
	ChannelTypeGuildStageVoice
	ChannelTypeGuildDirectory
	ChannelTypeGuildForum
)

// Message
// Represents a message sent in a channel within Discord.
//
// https://discord.com/developers/docs/resources/channel#message-object
type Message struct {
	ID              MessageID    `json:"id"`
	ChannelID       ChannelID    `json:"channel_id"`
	GuildID         GuildID      `json:"guild_id"`
	Author          User         `json:"author"`
	Member          *GuildMember `json:"member"`
	Content         string       `json:"content"`
	Timestamp       time.Time    `json:"timestamp"`
	EditedTimestamp *time.Time   `json:"edited_timestamp"`
	TTS             bool         `json:"tts"`
	MentionEveryone bool         `json:"mention_everyone"`
	Mentions        []User       `json:"mentions"`
	//MentionRoles    []RoleID     `json:"mention_roles"`
	MentionChannels []ChannelID `json:"mention_channels"`
	//Attachments //TODO
	//Embeds //TODO
	//Reactions //TODO
	//Nonce //TODO
	Pinned bool `json:"pinned"`
	//WebhookID //TODO
	//Type //TODO
	//Activity //TODO
	//Application //TODO
	ApplicationID *ApplicationID `json:"application_id"`
	//MessageReference //TODO
	Flags int `json:"flags"`
	//keep going... TODO
}

func (d *Discord) GetChannel(channelID ChannelID) (*Channel, error) {
	channel := new(Channel)
	return channel, d.get("/channels/"+string(channelID), channel)
}

func (d *Discord) CreateMessage(channelID ChannelID, message *MessageParams) (*Message, error) {
	mes := new(Message)
	return mes, d.post("/channels/"+string(channelID)+"/messages", mes, message)
}

type MessageParams struct {
	Content string `json:"content,omitempty"`
	// TODO
}
