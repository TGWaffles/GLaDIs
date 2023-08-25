package discord

import "github.com/tgwaffles/gladis/discord/command_option_type"

type ApplicationCommandOption struct {
	Name    string                                `json:"name"`
	Type    command_option_type.CommandOptionType `json:"type"`
	Value   interface{}                           `json:"value,omitempty"`
	Options []ApplicationCommandOption            `json:"options,omitempty"`
	Focused bool                                  `json:"focused,omitempty"`
}
