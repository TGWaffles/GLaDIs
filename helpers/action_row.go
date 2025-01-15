package helpers

import "github.com/JackHumphries9/dapper-go/discord"

func CreateActionRow(comps ...discord.MessageComponent) []discord.MessageComponent {
	return []discord.MessageComponent{
		&discord.ActionRow{
			Components: comps,
		},
	}
}
