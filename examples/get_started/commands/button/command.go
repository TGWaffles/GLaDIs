package button_command

import (
	"fmt"

	"github.com/JackHumphries9/dapper-go/client"
	"github.com/JackHumphries9/dapper-go/discord"
	"github.com/JackHumphries9/dapper-go/helpers"
	"github.com/JackHumphries9/dapper-go/interactable"
)

var Command = interactable.Command{
	Command: client.CreateApplicationCommand{
		Name:        "button",
		Description: helpers.Ptr("Button Command"),
	},
	OnCommand: CommandHandler,
	AssociatedComponents: []interactable.Component{
		nextPageButtonComponent, backPageButtonComponent,
	},
}

func CommandHandler(itx *interactable.InteractionContext) {
	err := itx.Respond(discord.ResponseEditData{
		Embeds:     []discord.Embed{firstEmbed},
		Components: helpers.CreateActionRow(&nextPageButton),
	})

	if err != nil {
		fmt.Printf("Failed to edit response")
	}
}
