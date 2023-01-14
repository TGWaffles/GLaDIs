package interactions

import (
	"encoding/json"
	"github.com/tgwaffles/lambda-discord-interactions-go/discord"
)

type Interaction struct {
	Id            discord.Snowflake       `json:"id"`
	ApplicationId discord.Snowflake       `json:"application_id"`
	Type          InteractionType         `json:"type"`
	Data          *map[string]interface{} `json:"data,omitempty"`
	GuildId       *discord.Snowflake      `json:"guild_id,omitempty"`
	ChannelId     *discord.Snowflake      `json:"channel_id,omitempty"`
	Member        *discord.Member         `json:"member,omitempty"`
	User          *discord.User           `json:"user,omitempty"`
	Token         string                  `json:"token"`
	Version       int                     `json:"version"`
}

func Parse(data string) (interaction Interaction, err error) {
	err = json.Unmarshal([]byte(data), &interaction)
	return interaction, err
}

func (interaction Interaction) IsPing() bool {
	return interaction.Type == Ping
}
