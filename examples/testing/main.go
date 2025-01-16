package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/JackHumphries9/dapper-go/client"
	"github.com/JackHumphries9/dapper-go/discord"
	"github.com/JackHumphries9/dapper-go/discord/text_input_style"
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

var Modal = interactable.Modal{
	Modal: discord.ModalCallback{
		Title:    "Example Modal",
		CustomId: "modal-example",
		Components: []discord.MessageComponent{
			&discord.ActionRow{
				Components: []discord.MessageComponent{
					&discord.TextInput{
						Label:       "Name",
						Placeholder: "Your name here...",
						Style:       text_input_style.Short,
						CustomId:    "me-name",
						Required:    true,
						MinLength:   helpers.Ptr(3),
						MaxLength:   helpers.Ptr(100),
					},
				},
			},
			&discord.ActionRow{
				Components: []discord.MessageComponent{
					&discord.TextInput{
						Label:       "A fact about you",
						Placeholder: "Your fact here...",
						Style:       text_input_style.Long,
						CustomId:    "me-fact",
						Required:    true,
						MinLength:   helpers.Ptr(3),
						MaxLength:   helpers.Ptr(250),
					},
				},
			},
		},
	},
	OnSubmit: func(itc *interactable.InteractionContext) {
		submissionData := itc.Interaction.Data.(*discord.ModalSubmitData)

		fmt.Printf("Submitted modal")

		valueMap := map[string]string{}

		for _, component := range submissionData.Components {
			textInput := component.(*discord.ActionRow).Components[0].(*discord.TextInput)
			valueMap[textInput.CustomId] = *textInput.Value
		}

		err := itc.Respond(discord.ResponseEditData{
			Embeds: []discord.Embed{
				{
					Title:       "Your Results!",
					Description: fmt.Sprintf("You are %s and your fact is: %s", valueMap["me-name"], valueMap["me-fact"]),
				},
			},
		})

		if err != nil {
			fmt.Printf("Failed to edit response")
		}
	},
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
			Name:        "modal",
			Description: helpers.Ptr("Modal Example"),
		},
		OnCommand: func(itc *interactable.InteractionContext) {
			err := itc.ShowModal(Modal)

			if err != nil {
				fmt.Printf("Failed to edit response")
			}
		},
		AssociatedModals: []interactable.Modal{Modal},
	})

	//botServer.RegisterCommandsWithDiscord(appId, botClient)

	botServer.Listen(3000)
}
