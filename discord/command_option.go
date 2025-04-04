package discord

import "github.com/tgwaffles/gladis/discord/command_option_type"

type ApplicationCommandDataOption struct {
	Name      string                                `json:"name"`
	Type      command_option_type.CommandOptionType `json:"type"`
	Value     any                                   `json:"value,omitempty"`
	Options   []ApplicationCommandDataOption        `json:"options,omitempty"`
	Focused   bool                                  `json:"focused,omitempty"`
	optionMap map[string]ApplicationCommandDataOption
}

func (commandDataOption *ApplicationCommandDataOption) GetOption(optionName string) *ApplicationCommandDataOption {
	if commandDataOption.Options == nil || len(commandDataOption.Options) == 0 {
		return nil
	}

	if commandDataOption.optionMap == nil {
		commandDataOption.optionMap = make(map[string]ApplicationCommandDataOption)
		for _, option := range commandDataOption.Options {
			commandDataOption.optionMap[option.Name] = option
		}
	}

	option, present := commandDataOption.optionMap[optionName]
	if !present {
		return nil
	}
	return &option
}
