package main

import (
	"fmt"

	"github.com/JackHumphries9/dapper-go/discord"
	"github.com/JackHumphries9/dapper-go/discord/command_type"
	"github.com/JackHumphries9/dapper-go/managers"
	"github.com/JackHumphries9/dapper-go/server"
)

const PUBLIC_KEY = "e9573621727df5e8b915f2a52b481d262f0d26e6d429913e1b960062ca6d4ab3"

func main() {
	server := server.NewInteractionServer(PUBLIC_KEY)

	server.RegisterCommand(managers.DapperCommand{
		Command: discord.ApplicationCommandData{
			Name: "hello",
			Type: command_type.Message,
		},
		CommandOptions: managers.DapperCommandOptions{
			Ephemeral: true,
		},
		Executor: func(itx *discord.Interaction) {
			err := itx.EditResponse(discord.ResponseEditData{
				Embeds: []discord.Embed{
					{
						Title: "Hello From Dapper",
					},
				},
			})

			if err != nil {
				fmt.Printf("failed to send edit response")
			}

		},
	})

	server.Listen(3000)
}
