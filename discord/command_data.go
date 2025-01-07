package discord

import "github.com/JackHumphries9/dapper-go/discord/command_type"

type ApplicationCommandData struct {
	Id        Snowflake                           `json:"id"`
	Name      string                              `json:"name"`
	Type      command_type.ApplicationCommandType `json:"type"`
	Resolved  *ResolvedData                       `json:"resolved,omitempty"`
	Options   []ApplicationCommandOption          `json:"options,omitempty"`
	GuildId   *Snowflake                          `json:"guild_id,omitempty"`
	TargetId  *Snowflake                          `json:"target_id,omitempty"`
	optionMap map[string]ApplicationCommandOption
}

func (commandData *ApplicationCommandData) GetOption(optionName string) *ApplicationCommandOption {
	if commandData.Options == nil || len(commandData.Options) == 0 {
		return nil
	}

	if commandData.optionMap == nil {
		commandData.optionMap = make(map[string]ApplicationCommandOption)
		for _, option := range commandData.Options {
			commandData.optionMap[option.Name] = option
		}
	}

	option, present := commandData.optionMap[optionName]
	if !present {
		return nil
	}
	return &option
}
