package sonya

import "time"

// https://discord.com/developers/docs/resources/guild#guild-object-guild-structure
type Guild struct {
	ID                          SnowflakeGuild    `json:"id,string"`
	Name                        string            `json:"name"`
	Icon                        *string           `json:"icon"`
	IconHash                    *string           `json:"icon_hash"`
	Splash                      *string           `json:"splash"`
	DiscoverySplash             *string           `json:"discovery_splash"`
	Owner                       bool              `json:"owner,omitempty"`
	OwnerID                     SnowflakeUser     `json:"owner_id,string"`
	Permissions                 string            `json:"permissions,omitempty"`
	AFKChannelID                *SnowflakeChannel `json:"afk_channel_id,string"`
	AFKTimeout                  int               `json:"afk_timeout"`
	WidgetEnabled               bool              `json:"widget_enabled"`
	WidgetChannelID             *SnowflakeChannel `json:"widget_channel_id,string"`
	VerificationLevel           int               `json:"verification_level"`
	DefaultMessageNotifications int               `json:"default_message_notifications"`
	ExplicitContentFilter       int               `json:"explicit_content_filter"`
	// TODO
	// Roles
	// Emojis
	// Features
	MFALevel           int               `json:"mfa_level"`
	ApplicationID      *SnowflakeGuild   `json:"application_id,string"`
	SystemChannelID    *SnowflakeChannel `json:"system_channel_id,string"`
	SystemChannelFlags int               `json:"system_channel_flags"`
	RulesChannelID     *SnowflakeChannel `json:"rules_channel_id,string"`
	JoinedAt           time.Time         `json:"joined_at,omitempty"`
	Large              bool              `json:"large,omitempty"`
	Unavailable        bool              `json:"unavailable,omitempty"`
	MemberCount        int               `json:"member_count,omitempty"`
	// TODO
	// VoiceStates
	// Members
	// Channels
	// Threads
	// Presences
	MaxPresences             *int              `json:"max_presences,omitempty"`
	MaxMembers               int               `json:"max_members,omitempty"`
	VanityURLCode            *string           `json:"vanity_url_code"`
	Description              *string           `json:"description"`
	Banner                   *string           `json:"banner"`
	PremiumTier              int               `json:"premium_tier"`
	PremiumSubscriptionCount int               `json:"premium_subscription_count,omitempty"`
	PreferredLocale          string            `json:"preferred_locale"`
	PublicUpdatesChannelID   *SnowflakeChannel `json:"public_updates_channel_id,string"`
	MaxVideoChannelUsers     int               `json:"max_video_channel_users,omitempty"`
	ApproximateMemberCount   int               `json:"approximate_member_count,omitempty"`
	ApproximatePresenceCount int               `json:"approximate_presence_count,omitempty"`
	// WelcomeScreen
	NSFWLevel int `json:"nsfw_level"`
	// StageInstances
	// Stickers
	// GuildScheduledEvents
	PremiumProgressbarEnabled bool `json:"premium_progressbar_enabled"`
}

// https://discord.com/developers/docs/resources/guild#guild-member-object-guild-member-structure
type GuildMember struct {
	User                       *User           `json:"user"`
	Nick                       *string         `json:"nick"`
	Avatar                     *string         `json:"avatar"` //TODO
	Roles                      []SnowflakeRole `json:"roles"`
	JoinedAt                   time.Time       `json:"joined_at"`
	PremiumSince               *time.Time      `json:"premium_since"`
	Deaf                       bool            `json:"deaf"`
	Mute                       bool            `json:"mute"`
	Pending                    *bool           `json:"pending"`
	Permissions                *string
	CommunicationDisabledUntil *time.Time `json:"communication_disabled_until"`
}
