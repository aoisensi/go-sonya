package sonya

// GetCurrentUser returns the user object of the requester's account.
// For OAuth2, this requires the identify scope,
// which will return the object without an email,
// and optionally the email scope,
// which returns the object with an email.
//
// https://discord.com/developers/docs/resources/user#get-current-user
func (s *Session) GetCurrentUser() (*User, error) {
	return s.GetUser("@me")
}

// GetUser returns a user object for a given user ID.
//
// https://discord.com/developers/docs/resources/user#get-user
func (s *Session) GetUser(id SnowflakeUser) (*User, error) {
	user := new(User)
	return user, s.get("/users/"+id.String(), user)
}

// ModifyCurrentUser modify the requester's user account settings.
// Returns a user object on success.
//
// https://discord.com/developers/docs/resources/user#modify-current-user
func (s *Session) ModifyCurrentUser(opts ...ModifyCurrentUserOption) (*User, error) {
	user := new(User)
	j := make(map[string]interface{})
	for _, opt := range opts {
		opt.optModifyCurrentUser(j)
	}
	return user, s.patch("/users/@me", user, j)
}

type ModifyCurrentUserOption interface {
	optModifyCurrentUser(v map[string]interface{})
}

type User struct {
	ID          SnowflakeUser `json:"id"`
	Username    string        `json:"username"`
	Discrim     string        `json:"discriminator"`
	Avatar      *string       `json:"avatar"`
	Bot         *bool         `json:"bot"`
	System      *bool         `json:"system"`
	MFAEnabled  *bool         `json:"mfa_enabled"`
	Banner      *string       `json:"banner"`
	AccentColor *int          `json:"accent_color"`
	Locale      *string       `json:"locale"`
	Verified    *bool         `json:"verified"`
	Email       *string       `json:"email"`
	Flags       *UserFlag     `json:"flags"`
	PremiumType *int          `json:"premium_type"`
	PublicFlags *int          `json:"public_flags"`
}

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
