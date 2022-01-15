package sonya

import "time"

// https://discord.com/developers/docs/resources/guild#guild-object-guild-structure
type Guild struct {
	ID                          SnowflakeGuild    `json:"id"`
	Name                        string            `json:"name"`
	Icon                        *string           `json:"icon"`
	IconHash                    *string           `json:"icon_hash"`
	Splash                      *string           `json:"splash"`
	DiscoverySplash             *string           `json:"discovery_splash"`
	Owner                       bool              `json:"owner,omitempty"`
	OwnerID                     SnowflakeUser     `json:"owner_id"`
	Permissions                 string            `json:"permissions,omitempty"`
	AFKChannelID                *SnowflakeChannel `json:"afk_channel_id"`
	AFKTimeout                  int               `json:"afk_timeout"`
	WidgetEnabled               bool              `json:"widget_enabled"`
	WidgetChannelID             *SnowflakeChannel `json:"widget_channel_id"`
	VerificationLevel           int               `json:"verification_level"`
	DefaultMessageNotifications int               `json:"default_message_notifications"`
	ExplicitContentFilter       int               `json:"explicit_content_filter"`
	// Roles
	// Emojis
	// Features
	MFALevel           int               `json:"mfa_level"`
	ApplicationID      *SnowflakeGuild   `json:"application_id"`
	SystemChannelID    *SnowflakeChannel `json:"system_channel_id"`
	SystemChannelFlags int               `json:"system_channel_flags"`
	RulesChannelID     *SnowflakeChannel `json:"rules_channel_id"`
	JoinedAt           time.Time         `json:"joined_at,omitempty"`
	Large              bool              `json:"large,omitempty"`
	Unavailable        bool              `json:"unavailable,omitempty"`
	MemberCount        int               `json:"member_count,omitempty"`
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
	PublicUpdatesChannelID   *SnowflakeChannel `json:"public_updates_channel_id"`
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
