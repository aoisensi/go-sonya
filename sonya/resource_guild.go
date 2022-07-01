package sonya

import "time"

// Guild
//
// https://discord.com/developers/docs/resources/guild#guild-object
type Guild struct {
	ID                          GuildID    `json:"id"`
	Name                        string     `json:"name"`
	Icon                        *string    `json:"icon"`
	IconHash                    *string    `json:"icon_hash"`
	Splash                      *string    `json:"splash"`
	DiscoverySplash             *string    `json:"discovery_splash"`
	Owner                       bool       `json:"owner,omitempty"`
	OwnerID                     UserID     `json:"owner_id"`
	Permissions                 string     `json:"permissions,omitempty"`
	AFKChannelID                *ChannelID `json:"afk_channel_id"`
	AFKTimeout                  int        `json:"afk_timeout"`
	WidgetEnabled               bool       `json:"widget_enabled"`
	WidgetChannelID             *ChannelID `json:"widget_channel_id"`
	VerificationLevel           int        `json:"verification_level"`
	DefaultMessageNotifications int        `json:"default_message_notifications"`
	ExplicitContentFilter       int        `json:"explicit_content_filter"`
	// TODO
	// Roles
	// Emojis
	// Features
	MFALevel           int        `json:"mfa_level"`
	ApplicationID      *GuildID   `json:"application_id"`
	SystemChannelID    *ChannelID `json:"system_channel_id"`
	SystemChannelFlags int        `json:"system_channel_flags"`
	RulesChannelID     *ChannelID `json:"rules_channel_id"`
	JoinedAt           time.Time  `json:"joined_at,omitempty"`
	Large              bool       `json:"large,omitempty"`
	Unavailable        bool       `json:"unavailable,omitempty"`
	MemberCount        int        `json:"member_count,omitempty"`
	// TODO
	// VoiceStates
	// Members
	// Channels
	// Threads
	// Presences
	MaxPresences             *int       `json:"max_presences,omitempty"`
	MaxMembers               int        `json:"max_members,omitempty"`
	VanityURLCode            *string    `json:"vanity_url_code"`
	Description              *string    `json:"description"`
	Banner                   *string    `json:"banner"`
	PremiumTier              int        `json:"premium_tier"`
	PremiumSubscriptionCount int        `json:"premium_subscription_count,omitempty"`
	PreferredLocale          string     `json:"preferred_locale"`
	PublicUpdatesChannelID   *ChannelID `json:"public_updates_channel_id"`
	MaxVideoChannelUsers     int        `json:"max_video_channel_users,omitempty"`
	ApproximateMemberCount   int        `json:"approximate_member_count,omitempty"`
	ApproximatePresenceCount int        `json:"approximate_presence_count,omitempty"`
	// WelcomeScreen
	NSFWLevel int `json:"nsfw_level"`
	// StageInstances
	// Stickers
	// GuildScheduledEvents
	PremiumProgressbarEnabled bool `json:"premium_progressbar_enabled"`
}

// GuildMember
//
// https://discord.com/developers/docs/resources/guild#guild-member-object
type GuildMember struct {
	User   *User   `json:"user"`
	Nick   *string `json:"nick"`
	Avatar *string `json:"avatar"` //TODO
	//Roles                      []RoleID   `json:"roles"`
	JoinedAt                   time.Time  `json:"joined_at"`
	PremiumSince               *time.Time `json:"premium_since"`
	Deaf                       bool       `json:"deaf"`
	Mute                       bool       `json:"mute"`
	Pending                    *bool      `json:"pending"`
	Permissions                *string
	CommunicationDisabledUntil *time.Time `json:"communication_disabled_until"`
}

// UnavailableGuild
// A partial guild object.
// Represents an Offline Guild,
// or a Guild whose information has not been provided through Guild Create events during the Gateway connect.
//
// https://discord.com/developers/docs/resources/guild#unavailable-guild-object
type UnavailableGuild struct {
	ID          GuildID `json:"id"`
	Unavailable bool    `json:"unavailable"`
}
