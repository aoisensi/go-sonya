package sonya

type Snowflake string

func (s Snowflake) String() string {
	return string(s)
}

type SnowflakeUser Snowflake

func (s SnowflakeUser) String() string {
	return string(s)
}
