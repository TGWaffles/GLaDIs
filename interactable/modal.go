package interactable

import (
	"github.com/JackHumphries9/dapper-go/discord"
	"github.com/JackHumphries9/dapper-go/discord/interaction_callback_type"
)

type Modal struct {
	Modal    discord.ModalCallback
	OnSubmit InteractionHandler
}

func (m *Modal) GetModalResponse() discord.InteractionResponse {
	return discord.InteractionResponse{
		Type: interaction_callback_type.Modal,
		Data: m.Modal,
	}
}
