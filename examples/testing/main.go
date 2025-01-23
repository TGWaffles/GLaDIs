package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/JackHumphries9/dapper-go/client"
	"github.com/JackHumphries9/dapper-go/discord"
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

			fileBytes, err := os.ReadFile("./examples/testing/test-image.png")
			if err != nil {
				fmt.Printf("Error reading file: %v\n", err)
				return
			}

			err = itc.Interaction.EditResponse(discord.ResponseEditData{
				Embeds: []discord.Embed{
					{
						Title:       "Hello World!",
						Description: "Hello World!",
						Image: &discord.EmbedImage{
							URL: "attachment://test-image.png",
						},
					},
				},
				Attachments: []discord.MessageAttachment{
					discord.NewBytesAttachment(fileBytes, "test-image.png", "image/png"),
				},
			})

			if err != nil {
				fmt.Printf("cannot respond to message %v", err)
			}
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
