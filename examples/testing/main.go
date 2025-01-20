package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

var Button = interactable.Button{
	Component: &discord.Button{
		Label:    helpers.Ptr("Ya boii"),
		CustomId: helpers.Ptr("yaboi"),
		Style:    button_style.Primary,
	},
	OnPress: func(itc *interactable.InteractionContext) {
		itc.SetEphemeral(true)

		err := itc.Respond(discord.ResponseEditData{
			Content: helpers.Ptr("Ya boiiiiiiiii " + *itc.GetIdContext()),
		})

		if err != nil {
			fmt.Printf("cannot respond to message")
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
			itc.Defer()

			err = itc.Respond(discord.ResponseEditData{
				Components: helpers.CreateActionRow(Button.CreateComponentInstance(interactable.ComponentInstanceOptions{
					ID: "hello",
				})),
			})

			if err != nil {
				fmt.Printf("cannot respond to message %v", err)
			}
		},
		AssociatedComponents: []interactable.Component{
			&Button,
		},
	})

	botServer.RegisterCommandsWithDiscord(appId, botClient)

	http.HandleFunc("/", botServer.Handle)

	// Start the server on port 8080
	fmt.Println("Starting server on :3000")
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
