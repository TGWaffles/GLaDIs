package button_command

import (
	"fmt"

	"github.com/JackHumphries9/dapper-go/discord"
	"github.com/JackHumphries9/dapper-go/discord/button_style"
	"github.com/JackHumphries9/dapper-go/helpers"
	"github.com/JackHumphries9/dapper-go/interactable"
)

var nextPageButton = discord.Button{
	Emoji: &discord.Emoji{
		Name: helpers.Ptr("➡️"),
	},
	Style:    button_style.Primary,
	CustomId: helpers.Ptr("button-next"),
}

var nextPageButtonComponent = interactable.Button{
	Component: &nextPageButton,
	OnPress: func(itx *interactable.InteractionContext) {
		err := itx.Respond(discord.ResponseEditData{
			Embeds:     []discord.Embed{secondEmbed},
			Components: helpers.CreateActionRow(&backPageButton),
		})

		if err != nil {
			fmt.Printf("Failed to edit response")
		}
	},
}

var backPageButton = discord.Button{
	Emoji: &discord.Emoji{
		Name: helpers.Ptr("⬅️"),
	},
	Style:    button_style.Secondary,
	CustomId: helpers.Ptr("button-back"),
}

var backPageButtonComponent = interactable.Button{
	Component: &backPageButton,
	OnPress: func(itx *interactable.InteractionContext) {
		err := itx.Respond(discord.ResponseEditData{
			Embeds:     []discord.Embed{firstEmbed},
			Components: helpers.CreateActionRow(&nextPageButton),
		})

		if err != nil {
			fmt.Printf("Failed to edit response")
		}
	},
}
