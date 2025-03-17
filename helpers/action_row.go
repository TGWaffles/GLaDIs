package helpers

import "github.com/tgwaffles/gladis/discord"

func CreateActionRow(comps ...discord.MessageComponent) []discord.MessageComponent {
	return []discord.MessageComponent{
		&discord.ActionRow{
			Components: comps,
		},
	}
}
