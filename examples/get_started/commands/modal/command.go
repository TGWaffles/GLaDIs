package modal_command

import (
	"fmt"

	"github.com/JackHumphries9/dapper-go/client"
	"github.com/JackHumphries9/dapper-go/helpers"
	"github.com/JackHumphries9/dapper-go/interactable"
)

var Command = interactable.Command{
	Command: client.CreateApplicationCommand{
		Name:        "modaltesting",
		Description: helpers.Ptr("Modal Example"),
	},
	OnCommand: func(itc *interactable.InteractionContext) {
		err := itc.ShowModal(Modal)

		if err != nil {
			fmt.Printf("Failed to edit response")
		}
	},
}
