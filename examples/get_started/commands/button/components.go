package button_command

import (
	"fmt"

	"github.com/JackHumphries9/dapper-go/dapper"
	"github.com/JackHumphries9/dapper-go/discord"
	"github.com/JackHumphries9/dapper-go/discord/button_style"
	"github.com/JackHumphries9/dapper-go/helpers"
)

var nextPageButton = discord.Button{
	Label: helpers.Ptr("Next"),
	Emoji: &discord.Emoji{
		Name: helpers.Ptr("➡️"),
	},
	Style:    button_style.Primary,
	CustomId: helpers.Ptr("button-next"),
}

var nextPageButtonComponent = dapper.DapperButton{
	Component: &nextPageButton,
	OnPress: func(itx *discord.Interaction) {
		err := itx.EditResponse(discord.ResponseEditData{
			Embeds:     []discord.Embed{secondEmbed},
			Components: helpers.CreateActionRow(&nextPageButton),
		})

		if err != nil {
			fmt.Printf("Failed to edit response")
		}
	},
}

var backPageButton = discord.Button{
	Label: helpers.Ptr("Back"),
	Emoji: &discord.Emoji{
		Name: helpers.Ptr("➡️"),
	},
	Style:    button_style.Primary,
	CustomId: helpers.Ptr("button-back"),
}

var backPageButtonComponent = dapper.DapperButton{
	Component: &backPageButton,
	OnPress: func(itx *discord.Interaction) {
		err := itx.EditResponse(discord.ResponseEditData{
			Embeds:     []discord.Embed{secondEmbed},
			Components: helpers.CreateActionRow(&backPageButton),
		})

		if err != nil {
			fmt.Printf("Failed to edit response")
		}
	},
}
