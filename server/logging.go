package server

import (
	"log"

	"github.com/JackHumphries9/dapper-go/discord"
	"github.com/JackHumphries9/dapper-go/discord/interaction_type"
)

type DapperLogger struct {
	OnInteractionRecieved func(itx *discord.Interaction)
	Info                  func(message string)
	Error                 func(message string)
}

var DefaultLogger = DapperLogger{
	OnInteractionRecieved: func(itx *discord.Interaction) {
		if itx.Type == interaction_type.ApplicationCommand {
			log.Printf("Recieved command %s\n", itx.Data.(*discord.ApplicationCommandData).Name)
		} else if itx.Type == interaction_type.MessageComponent {
			log.Printf("Recieved message component %s\n", itx.Data.(*discord.MessageComponentData).CustomId)
		} else if itx.Type == interaction_type.ModalSubmit {
			log.Printf("Recieved modal submit %s\n", itx.Data.(*discord.ModalSubmitData).CustomId)
		}
	},
	Info: func(message string) {
		log.Println(message)
	},
	Error: func(message string) {
		log.Printf("Error: %s\n", message)
	},
}
