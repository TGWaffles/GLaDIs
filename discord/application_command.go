package discord

import (
	"github.com/tgwaffles/gladis/discord/command_option_type"
)

type ApplicationCommand struct {
	ID                      Snowflake                              `json:"id"`
	Type                    *command_option_type.CommandOptionType `json:"type,omitempty"`
	ApplicationID           Snowflake                              `json:"application_id"`
	GuildID                 *Snowflake                             `json:"guild_id,omitempty"`
	Name                    string                                 `json:"name,omitempty"`
	NameLocalizations       *map[string]string                     `json:"name_localizations,omitempty"`
	Discription             string                                 `json:"description"`
	DescriptionLocalization *map[string]string                     `json:"description_localizations,omitempty"`
	Options                 *[]ApplicationCommandOption            `json:"options,omitempty"`
}

type ApplicationCommandOption struct {
	Type                     command_option_type.CommandOptionType `json:"type"`
	Name                     string                                `json:"name"`
	NameLocalizations        map[string]string                     `json:"name_localizations,omitempty"`
	Description              string                                `json:"description"`
	DescriptionLocalizations map[string]string                     `json:"description_localizations,omitempty"`
	Required                 *bool                                 `json:"required,omitempty"`
	Choices                  []ApplicationCommandOptionChoice      `json:"choices,omitempty"`
	Options                  []ApplicationCommandOption            `json:"options,omitempty"`
	ChannelTypes             []string                              `json:"channel_types,omitempty"`
	MinValue                 *float64                              `json:"min_value,omitempty"`
	MaxValue                 *float64                              `json:"max_value,omitempty"`
	MinLength                *int                                  `json:"min_length,omitempty"`
	MaxLength                *int                                  `json:"max_length,omitempty"`
	Autocomplete             *bool                                 `json:"autocomplete,omitempty"`
}

type ApplicationCommandOptionChoice struct {
	Name              string            `json:"name"`
	Value             interface{}       `json:"value"`
	NameLocalizations map[string]string `json:"name_localizations,omitempty"`
}
