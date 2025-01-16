package managers

import (
	"fmt"

	"github.com/JackHumphries9/dapper-go/client"
	"github.com/JackHumphries9/dapper-go/discord"
	"github.com/JackHumphries9/dapper-go/discord/interaction_type"
	"github.com/JackHumphries9/dapper-go/interactable"
)

type CommandManager struct {
	commands map[string]interactable.Command
}

func NewDapperCommandManager() CommandManager {
	return CommandManager{
		commands: make(map[string]interactable.Command, 0),
	}
}

func (dcm *CommandManager) Type() interaction_type.InteractionType {
	return interaction_type.ApplicationCommand
}

func (dcm *CommandManager) RouteInteraction(itx *discord.Interaction) (discord.InteractionResponse, error) {
	commandData := itx.Data.(*discord.ApplicationCommandData)

	if cmd, ok := dcm.commands[commandData.Name]; ok {

		itc := interactable.InteractionContext{
			Interaction:  itx,
			DeferChannel: make(chan *discord.InteractionResponse),
		}

		go cmd.OnCommand(&itc)

		response := <-itc.DeferChannel

		return *response, nil
	}

	return discord.InteractionResponse{}, fmt.Errorf("No command found")
}

func (dcm *CommandManager) Register(command interactable.Command) {
	dcm.commands[command.Command.Name] = command
}

func (dcm *CommandManager) RegisterCommandsWithDiscord(appId discord.Snowflake, botClient *client.BotClient) error {
	discordCommands := make([]client.CreateApplicationCommand, 0, len(dcm.commands))

	for _, cmd := range dcm.commands {
		discordCommands = append(discordCommands, cmd.Command)
	}

	return botClient.GetApplicationClient(appId).RegisterCommands(discordCommands)
}
