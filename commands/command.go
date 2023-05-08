package commands

import (
	"github.com/tgwaffles/gladis/discord"
)

type ApplicationCommandData struct {
	Id        discord.Snowflake          `json:"id"`
	Name      string                     `json:"name"`
	Type      ApplicationCommandType     `json:"type"`
	Resolved  *discord.ResolvedData      `json:"resolved,omitempty"`
	Options   []ApplicationCommandOption `json:"options,omitempty"`
	GuildId   *discord.Snowflake         `json:"guild_id,omitempty"`
	TargetId  *discord.Snowflake         `json:"target_id,omitempty"`
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
