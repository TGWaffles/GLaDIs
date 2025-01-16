package modal_command

import (
	"fmt"

	"github.com/JackHumphries9/dapper-go/discord"
	"github.com/JackHumphries9/dapper-go/discord/text_input_style"
	"github.com/JackHumphries9/dapper-go/helpers"
	"github.com/JackHumphries9/dapper-go/interactable"
)

var Modal = interactable.Modal{
	Modal: discord.ModalCallback{
		Title:    "Example Modal",
		CustomId: "modal-example",
		Components: helpers.CreateActionRow(
			&discord.TextInput{
				Label:       "Name",
				Placeholder: "Your name here...",
				Style:       text_input_style.Short,
				CustomId:    "modal-example-name",
				Required:    true,
			},
			&discord.TextInput{
				Label:       "A fact about you",
				Placeholder: "Your fact here...",
				Style:       text_input_style.Short,
				CustomId:    "modal-example-fact",
				Required:    true,
			},
		),
	},
	OnSubmit: func(itc *interactable.InteractionContext) {
		submissionData := itc.Interaction.Data.(*discord.ModalSubmitData)

		valueMap := map[string]string{}

		for _, component := range submissionData.Components {
			textInput := component.(*discord.TextInput)
			valueMap[textInput.CustomId] = textInput.Value
		}

		err := itc.Respond(discord.ResponseEditData{
			Embeds: []discord.Embed{
				{
					Title:       "Your Results!",
					Description: fmt.Sprintf("You are %s and your fact is: %s", valueMap["modal-example-name"], valueMap["modal-example-fact"]),
				},
			},
		})

		if err != nil {
			fmt.Printf("Failed to edit response")
		}
	},
}
