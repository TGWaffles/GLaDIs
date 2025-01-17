package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/JackHumphries9/dapper-go/client"
	"github.com/JackHumphries9/dapper-go/discord"
	"github.com/JackHumphries9/dapper-go/discord/component_type"
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

var Select = interactable.Select{
	Component: &discord.SelectMenu{
		MenuType: component_type.StringSelect,
		CustomId: "fruits",
		Options: []discord.SelectOption{
			{
				Label: "Apple",
				Value: "apple",
			},
			{
				Label: "Orange",
				Value: "orange",
			},
			{
				Label: "Banana",
				Value: "banana",
			},
		},
	},
	ComponentOptions: interactable.ComponentOptions{
		Ephemeral:   true,
		CancelDefer: true,
	},
	OnSelect: func(itc *interactable.InteractionContext) {
		vals, err := itc.GetSelectValues()

		if err != nil {
			fmt.Printf("failed to get values")
		}

		err = itc.Respond(discord.ResponseEditData{
			Content: helpers.Ptr(fmt.Sprintf("You chose: %s", (*vals)[0])),
		})

		if err != nil {
			fmt.Printf("cannot respond")
		}
	},
}

func main() {
	var env = LoadJSONEnv()

	botServer := server.NewInteractionServer(env.PublicKey)
	botClient := client.NewBot(env.BotToken)
	appId, err := discord.GetSnowflake(env.AppId)

	if err != nil {
		panic("Heyo you messed up")
	}

	botServer.RegisterCommand(interactable.Command{
		Command: client.CreateApplicationCommand{
			Name:        "fruit",
			Description: helpers.Ptr("Ping Pong"),
		},
		OnCommand: func(itc *interactable.InteractionContext) {
			itc.SetEphemeral(true)

			err = itc.Respond(discord.ResponseEditData{
				Components: helpers.CreateActionRow(Select.Component),
			})

			if err != nil {
				fmt.Printf("cannot respond to message")
			}
		},
		AssociatedComponents: []interactable.Component{
			&Select,
		},
	})

	botServer.RegisterCommandsWithDiscord(appId, botClient)

	botServer.Listen(3000)
}
