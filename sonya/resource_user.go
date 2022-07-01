package sonya

import "net/url"

// GetCurrentUser returns the user object of the requester's account.
// For OAuth2, this requires the identify scope,
// which will return the object without an email,
// and optionally the email scope,
// which returns the object with an email.
//
// https://discord.com/developers/docs/resources/user#get-current-user
func (s *Discord) GetCurrentUser() (*User, error) {
	return s.getUser("@me")
}

// GetUser returns a user object for a given user ID.
//
// https://discord.com/developers/docs/resources/user#get-user
func (s *Discord) GetUser(id UserID) (*User, error) {
	return s.getUser(string(id))
}

func (s *Discord) getUser(id string) (*User, error) {
	user := new(User)
	return user, s.get("/users/"+id, user)
}

// ModifyCurrentUser modify the requester's user account settings.
// Returns a user object on success.
//
// https://discord.com/developers/docs/resources/user#modify-current-user
func (s *Discord) ModifyCurrentUser(opts ...ModifyCurrentUserOption) (*User, error) {
	user := new(User)
	j := make(map[string]any)
	for _, opt := range opts {
		opt.optModifyCurrentUser(j)
	}
	return user, s.patch("/users/@me", user, j)
}

type ModifyCurrentUserOption interface {
	optModifyCurrentUser(v map[string]any)
}

// GetCurrentUserGuilds returns a list of partial guild objects the current user is a member of.
// Requires the guilds OAuth2 scope.
//
// https://discord.com/developers/docs/resources/user#get-current-user-guilds
func (s *Discord) GetCurrentUserGuilds(opts ...GetCurrentUserGuildsOption) ([]*Guild, error) {
	guilds := make([]*Guild, 0, 200)
	v := url.Values{}
	for _, opt := range opts {
		opt.optGetCurrentUserGuilds(v)
	}
	return guilds, s.get("/users/@me/guilds?"+v.Encode(), &guilds)
}

type GetCurrentUserGuildsOption interface {
	optGetCurrentUserGuilds(v url.Values)
}

// User
//
// https://discord.com/developers/docs/resources/user#user-object
type User struct {
	ID          UserID    `json:"id"`
	Username    string    `json:"username"`
	Discrim     string    `json:"discriminator"`
	Avatar      *string   `json:"avatar"`
	Bot         *bool     `json:"bot"`
	System      *bool     `json:"system"`
	MFAEnabled  *bool     `json:"mfa_enabled"`
	Banner      *string   `json:"banner"`
	AccentColor *int      `json:"accent_color"`
	Locale      *string   `json:"locale"`
	Verified    *bool     `json:"verified"`
	Email       *string   `json:"email"`
	Flags       *UserFlag `json:"flags"`
	PremiumType *int      `json:"premium_type"`
	PublicFlags *int      `json:"public_flags"`
}

// GetCurrentUserGuilds Returns a guild member object for the current user.
// Requires the guilds.members.read OAuth2 scope.
//
// https://discord.com/developers/docs/resources/user#get-current-user-guild-member
func (d *Discord) GetCurrentUserGuildMember(guildID GuildID) (*GuildMember, error) {
	member := new(GuildMember)
	return member, d.get("/users/@me/guilds/"+string(guildID)+"/members", member)
}

// https://discord.com/developers/docs/resources/user#user-object-user-flags
type UserFlag int

const (
	UserFlagStaff UserFlag = 1 << iota
	UserFlagPartner
	UserFlagHypesquad
	UserFlagBugHunterLevel1
	_
	_
	UserFlagHypersquadOnlineHouse1
	UserFlagHypersquadOnlineHouse2
	UserFlagHypersquadOnlineHouse3
	UserFlagPremiumEarlySupporter
	UserFlagTeamPseudoUser
	_
	_
	_
	UserFlagBugHunterLevel2
	_
	UserFlagVerifiedBot
	UserFlagVerifiedDeveloper
	UserFlagCertifiedModerator
	UserFlagBotHTTPInteractions
)
