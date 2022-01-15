package sonya

type Snowflake string

func (s Snowflake) String() string {
	return string(s)
}

type SnowflakeGuild Snowflake

func (s SnowflakeGuild) String() string {
	return string(s)
}

type SnowflakeUser Snowflake

func (s SnowflakeUser) String() string {
	return string(s)
}

type SnowflakeChannel Snowflake

func (s SnowflakeChannel) String() string {
	return string(s)
}
