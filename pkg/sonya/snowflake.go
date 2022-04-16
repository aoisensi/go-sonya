package sonya

import "strconv"

type Snowflake int64

func (s Snowflake) String() string {
	return strconv.FormatUint(uint64(s), 10)
}

type SnowflakeGuild Snowflake

func (s SnowflakeGuild) String() string {
	return Snowflake(s).String()
}

type SnowflakeUser Snowflake

func (s SnowflakeUser) String() string {
	return Snowflake(s).String()
}

type SnowflakeChannel Snowflake

func (s SnowflakeChannel) String() string {
	return Snowflake(s).String()
}

type SnowflakeRole Snowflake

func (s SnowflakeRole) String() string {
	return Snowflake(s).String()
}
