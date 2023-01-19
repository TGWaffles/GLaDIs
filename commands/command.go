package commands

import (
	"github.com/tgwaffles/gladis/discord"
)

type ApplicationCommandData struct {
	Id       discord.Snowflake           `json:"id"`
	Name     string                      `json:"name"`
	Type     ApplicationCommandType      `json:"type"`
	Resolved *discord.ResolvedData       `json:"resolved,omitempty"`
	Options  *[]ApplicationCommandOption `json:"options,omitempty"`
	GuildId  *discord.Snowflake          `json:"guild_id,omitempty"`
	TargetId *discord.Snowflake          `json:"target_id,omitempty"`
}
