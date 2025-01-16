package ping_command

import (
	"fmt"

	"github.com/JackHumphries9/dapper-go/client"
	"github.com/JackHumphries9/dapper-go/discord"
	"github.com/JackHumphries9/dapper-go/helpers"
	"github.com/JackHumphries9/dapper-go/interactable"
)

var Command = interactable.Command{
	Command: client.CreateApplicationCommand{
		Name:        "ping",
		Description: helpers.Ptr("Ping Pong!"),
	},
	OnCommand: func(itc *interactable.InteractionContext) {
		err := itc.Respond(discord.ResponseEditData{
			Content: helpers.Ptr("Pong!"),
		})

		if err != nil {
			fmt.Printf("Failed to edit response")
		}
	},
}
