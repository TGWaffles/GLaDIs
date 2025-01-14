package dapper

import (
	"fmt"

	"github.com/JackHumphries9/dapper-go/client"
	"github.com/JackHumphries9/dapper-go/discord"
	"github.com/JackHumphries9/dapper-go/discord/interaction_callback_type"
)

type DapperCommandManager struct {
	commands []DapperCommand
}

func NewDapperCommandManager() DapperCommandManager {
	return DapperCommandManager{
		commands: make([]DapperCommand, 0),
	}
}

func (dcm *DapperCommandManager) Register(command DapperCommand) {
	dcm.commands = append(dcm.commands, command)
}

func (dcm *DapperCommandManager) RouteInteraction(itx *discord.Interaction) (discord.InteractionResponse, error) {
	commandData := itx.Data.(*discord.ApplicationCommandData)

	for _, cmd := range dcm.commands {
		if cmd.Command.Name == commandData.Name {
			flags := 0

			if cmd.CommandOptions.Ephemeral {
				flags = 64
			}

			go cmd.OnCommand(itx)
			return discord.InteractionResponse{
				Type: interaction_callback_type.DeferredChannelMessageWithSource,
				Data: &discord.MessageCallbackData{
					Flags: &flags,
				},
			}, nil
		}
	}

	return discord.InteractionResponse{}, fmt.Errorf("No command found")
}

func (dcm *DapperCommandManager) RegisterCommandsWithDiscord(appId discord.Snowflake, botClient *client.BotClient) error {
	discordCommands := make([]client.CreateApplicationCommand, 0, len(dcm.commands))

	for _, cmd := range dcm.commands {
		discordCommands = append(discordCommands, cmd.Command)
	}

	return botClient.GetApplicationClient(appId).RegisterCommands(discordCommands)
}
