package sonya

import "fmt"

// https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-structure
type ApplicationCommand struct {
	ID                       ApplicationCommandID       `json:"id"`
	Type                     *ApplicationCommandType    `json:"type,omitempty"`
	ApplicationID            ApplicationID              `json:"application_id"`
	GuildID                  *GuildID                   `json:"guild_id,omitempty"`
	Name                     string                     `json:"name"`
	NameLocalizations        map[string]string          `json:"name_localizations,omitempty"`
	Description              string                     `json:"description"`
	DescriptionLocalizations map[string]string          `json:"description_localizations,omitempty"`
	Options                  []ApplicationCommandOption `json:"options,omitempty"`
	DefaultPermission        *bool                      `json:"default_permission,omitempty"`
	Version                  Snowflake                  `json:"version"`
}

// https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-types
type ApplicationCommandType int

const (
	ApplicationCommandTypeChatInput ApplicationCommandType = iota + 1
	ApplicationCommandTypeUser
	ApplicationCommandTypeMessage
)

// https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-option-structure
type ApplicationCommandOption struct {
	Type                     *ApplicationCommandOptionType    `json:"type"`
	Name                     string                           `json:"name"`
	NameLocalizations        map[string]string                `json:"name_localizations,omitempty"`
	Description              string                           `json:"description"`
	DescriptionLocalizations map[string]string                `json:"description_localizations,omitempty"`
	Required                 *bool                            `json:"required,omitempty"`
	Choices                  []ApplicationCommandOptionChoice `json:"choices,omitempty"`
	Options                  []ApplicationCommandOption       `json:"options,omitempty"`
	ChannelTypes             []ChannelType                    `json:"channel_types,omitempty"`
	MinValue                 float64                          `json:"min_value,omitempty"`
	MaxValue                 float64                          `json:"max_value,omitempty"`
	AutoComplete             *bool                            `json:"auto_complete,omitempty"`
}

// https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-option-type
type ApplicationCommandOptionType int

const (
	ApplicationCommandOptionTypeSubCommand ApplicationCommandOptionType = iota + 1
	ApplicationCommandOptionTypeSubCommandGroup
	ApplicationCommandOptionTypeString
	ApplicationCommandOptionTypeInteger
	ApplicationCommandOptionTypeBoolean
	ApplicationCommandOptionTypeUser
	ApplicationCommandOptionTypeChannel
	ApplicationCommandOptionTypeRole
	ApplicationCommandOptionTypeMentionable
	ApplicationCommandOptionTypeNumber
	ApplicationCommandOptionTypeAttachment
)

// https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-option-choice-structure
type ApplicationCommandOptionChoice struct {
	Name              string            `json:"name"`
	NameLocalizations map[string]string `json:"name_localizations,omitempty"`
	Value             any               `json:"value"`
}

// https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-interaction-data-option-structure
type ApplicationCommandInteractionDataOption struct {
	Name    string                                    `json:"name"`
	Type    ApplicationCommandOptionType              `json:"type"`
	Value   any                                       `json:"value,omitempty"`
	Options []ApplicationCommandInteractionDataOption `json:"options,omitempty"`
	Focused *bool                                     `json:"focused,omitempty"`
}

// GetGlobalApplicationCommands fetch all of the global commands for your application.
// Returns an array of application command objects.
//
// https://discord.com/developers/docs/interactions/application-commands#get-global-application-commands
func (d *Discord) GetGlobalApplicationCommands(applicationID ApplicationID, withLocalization bool) ([]ApplicationCommand, error) {
	commands := make([]ApplicationCommand, 0, 16)
	url_ := fmt.Sprintf("/applications/%v/commands", applicationID)
	if withLocalization {
		url_ += "?with_localization=true"
	}
	return commands, d.get(url_, &commands)
}
