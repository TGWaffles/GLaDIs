package ping_command

import (
	"fmt"

	"github.com/JackHumphries9/dapper-go/client"
	"github.com/JackHumphries9/dapper-go/dapper"
	"github.com/JackHumphries9/dapper-go/discord"
	"github.com/JackHumphries9/dapper-go/helpers"
)

var Comand = dapper.DapperCommand{
	Command: client.CreateApplicationCommand{
		Name:        "ping",
		Description: helpers.Ptr("Ping Pong!"),
	},
	OnCommand: func(itx *discord.Interaction) {
		err := itx.EditResponse(discord.ResponseEditData{
			Content: helpers.Ptr("Pong!"),
		})

		if err != nil {
			fmt.Printf("Failed to edit response")
		}
	},
}
