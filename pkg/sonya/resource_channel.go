package sonya

import "time"

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
	ID              MessageID    `json:"id,string"`
	ChannelID       ChannelID    `json:"channel_id,string"`
	GuildID         GuildID      `json:"guild_id,string"`
	Author          User         `json:"author"`
	Member          *GuildMember `json:"member"`
	Content         string       `json:"content"`
	Timestamp       time.Time    `json:"timestamp"`
	EditedTimestamp *time.Time   `json:"edited_timestamp"`
	TTS             bool         `json:"tts"`
	MentionEveryone bool         `json:"mention_everyone"`
	Mentions        []User       `json:"mentions"`
	//MentionRoles    []RoleID     `json:"mention_roles,string"`
	MentionChannels []ChannelID `json:"mention_channels,string"`
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
