package sonya

// https://discord.com/developers/docs/resources/application#application-object
type Application struct {
	ID                  ApplicationID `json:"id,string"`
	Name                string        `json:"name"`
	Icon                *string       `json:"icon"`
	Description         string        `json:"description"`
	RPCOrigins          []string      `json:"rpc_origins"`
	BotPublic           bool          `json:"bot_public"`
	BotRequireCodeGrant bool          `json:"bot_require_code_grant"`
	TermsOfServiceURL   *string       `json:"terms_of_service_url"`
	PrivacyPolicyURL    *string       `json:"privacy_policy_url"`
	Owner               *User         `json:"owner"`
	VerifyKey           string
	//Team //TODO
	GuildID *GuildID `json:"guild_id,string"`
	//PrimarySKUID //TODO
	Slug       *string  `json:"slug"`
	CoverImage *string  `json:"cover_image"`
	Flags      *int     `json:"flags"`
	Tags       []string `json:"tags"`
	//InstallParams //TODO
	CustomInstallURL *string `json:"custom_install_url"`
}
