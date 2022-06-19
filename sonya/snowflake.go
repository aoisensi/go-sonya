package sonya

import "strconv"

type Snowflake int64

func (s Snowflake) String() string {
	return strconv.FormatUint(uint64(s), 10)
}

type GuildID Snowflake

func (id GuildID) String() string {
	return Snowflake(id).String()
}

type UserID Snowflake

func (id UserID) String() string {
	return Snowflake(id).String()
}

type ChannelID Snowflake

func (id ChannelID) String() string {
	return Snowflake(id).String()
}

type RoleID Snowflake

func (id RoleID) String() string {
	return Snowflake(id).String()
}

type ApplicationID Snowflake

func (id ApplicationID) String() string {
	return Snowflake(id).String()
}

type MessageID Snowflake

func (id MessageID) String() string {
	return Snowflake(id).String()
}

type EmojiID Snowflake

func (id EmojiID) String() string {
	return Snowflake(id).String()
}

type ApplicationCommandID Snowflake

func (id ApplicationCommandID) String() string {
	return Snowflake(id).String()
}
