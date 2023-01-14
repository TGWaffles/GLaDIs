package commands

import "github.com/tgwaffles/lambda-discord-interactions-go/discord"

type ApplicationCommandData struct {
	Id   discord.Snowflake      `json:"id"`
	Name string                 `json:"name"`
	Type ApplicationCommandType `json:"type"`
}
