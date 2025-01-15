package button_command

import (
	"fmt"

	"github.com/JackHumphries9/dapper-go/client"
	"github.com/JackHumphries9/dapper-go/dapper"
	"github.com/JackHumphries9/dapper-go/discord"
	"github.com/JackHumphries9/dapper-go/helpers"
)

var Command = dapper.DapperCommand{
	Command: client.CreateApplicationCommand{
		Name:        "button",
		Description: helpers.Ptr("Button Command"),
	},
	CommandOptions: dapper.DapperCommandOptions{
		Ephemeral: true,
	},
	OnCommand: CommandHandler,
	AssociatedComponents: []dapper.DapperComponent{
		nextPageButtonComponent, backPageButtonComponent,
	},
}

func CommandHandler(itx *discord.Interaction) {
	err := itx.EditResponse(discord.ResponseEditData{
		Embeds:     []discord.Embed{firstEmbed},
		Components: helpers.CreateActionRow(&nextPageButton),
	})

	if err != nil {
		fmt.Printf("Failed to edit response")
	}
}
