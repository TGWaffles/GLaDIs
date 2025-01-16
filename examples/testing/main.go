package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/JackHumphries9/dapper-go/client"
	"github.com/JackHumphries9/dapper-go/discord"
	"github.com/JackHumphries9/dapper-go/discord/button_style"
	"github.com/JackHumphries9/dapper-go/helpers"
	"github.com/JackHumphries9/dapper-go/interactable"
	"github.com/JackHumphries9/dapper-go/server"
)

const FILENAME = "./examples/env.json"

type Env struct {
	PublicKey string `json:"PUBLIC_KEY"`
	BotToken  string `json:"BOT_TOKEN"`
	AppId     string `json:"APP_ID"`
}

func LoadJSONEnv() Env {
	plan, err := os.ReadFile(FILENAME)

	if err != nil {
		panic("no env file")
	}

	var data Env
	err = json.Unmarshal(plan, &data)

	if err != nil {
		panic("cannot unmarshal")
	}

	return data
}

var firstEmbed = discord.Embed{
	Title:       "Page 1",
	Description: "I'm the first page!",
}

var secondEmbed = discord.Embed{
	Title:       "Page 2",
	Description: "I'm the second page!",
}

var nextPageButton = discord.Button{
	Emoji: &discord.Emoji{
		Name: helpers.Ptr("➡️"),
	},
	Style:    button_style.Primary,
	CustomId: helpers.Ptr("button-next"),
}

var nextPageButtonComponent = interactable.Button{
	Component: &nextPageButton,
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

func main() {
	var env = LoadJSONEnv()

	botServer := server.NewInteractionServer(env.PublicKey)
	// botClient := client.NewBot(env.BotToken)
	// appId, err := discord.GetSnowflake(env.AppId)

	// if err != nil {
	// 	panic("Heyo you messed up")
	// }

	botServer.RegisterCommand(interactable.Command{
		Command: client.CreateApplicationCommand{
			Name:        "ping",
			Description: helpers.Ptr("Ping Pong!"),
		},
		CommandOptions: interactable.CommandOptions{
			ExperimentalAutoDeferral: true,
		},
		OnCommand: func(itc *interactable.InteractionContext) {
			itc.SetEphemeral(true)

			//time.Sleep(6 * time.Second)

			itc.Respond(discord.ResponseEditData{
				Content: helpers.Ptr("Pong!"),
			})
		},
	})

	//botServer.RegisterCommandsWithDiscord(appId, botClient)

	botServer.Listen(3000)
}
