package interactable

import (
	"fmt"

	"github.com/JackHumphries9/dapper-go/discord"
	"github.com/JackHumphries9/dapper-go/discord/interaction_callback_type"
	"github.com/JackHumphries9/dapper-go/discord/interaction_type"
	"github.com/JackHumphries9/dapper-go/discord/message_flags"
	"github.com/JackHumphries9/dapper-go/helpers"
)

type InteractionHandler func(itc *InteractionContext)

type InteractionContext struct {
	Interaction  *discord.Interaction
	DeferChannel chan *discord.InteractionResponse
	hasDeferred  bool
	messageFlags message_flags.MessageFlags
}

func (ic *InteractionContext) SetEphemeral(ep bool) {
	if ep {
		ic.messageFlags.AddFlag(message_flags.Ephemeral)
	} else {
		ic.messageFlags.RemoveFlag(message_flags.Ephemeral)
	}
}

func (ic *InteractionContext) HasDeferred() bool {
	return ic.hasDeferred
}

func (ic *InteractionContext) Defer() {
	if ic.hasDeferred {
		fmt.Printf("Interaction already deferred")
		return
	}

	ic.hasDeferred = true

	if ic.Interaction.Type == interaction_type.ApplicationCommand {
		ic.DeferChannel <- &discord.InteractionResponse{
			Type: interaction_callback_type.DeferredChannelMessageWithSource,
			Data: &discord.MessageCallbackData{
				Flags: helpers.Ptr(int(ic.messageFlags)),
			},
		}
	}

	if ic.Interaction.Type == interaction_type.MessageComponent {
		ic.DeferChannel <- &discord.InteractionResponse{
			Type: interaction_callback_type.DeferredUpdateMessage,
			Data: &discord.MessageCallbackData{
				Flags: helpers.Ptr(int(ic.messageFlags)),
			},
		}
	}
}

func (ic *InteractionContext) Respond(msg discord.ResponseEditData) error {
	if ic.hasDeferred {
		return ic.Interaction.EditResponse(msg)
	}

	var responseType int

	if ic.Interaction.Type == interaction_type.ApplicationCommand {
		responseType = interaction_callback_type.ChannelMessageWithSource
	} else if ic.Interaction.Type == interaction_type.MessageComponent {
		responseType = interaction_callback_type.UpdateMessage
	}

	ic.DeferChannel <- &discord.InteractionResponse{
		Type: interaction_callback_type.ChannelMessageWithSource,
		Data: &discord.MessageCallbackData{
			Content:         msg.Content,
			Flags:           &responseType,
			Embeds:          msg.Embeds,
			Components:      msg.Components,
			AllowedMentions: msg.AllowedMentions,
		},
	}

	return nil
}

func (ic *InteractionContext) ShowModal(modal Modal) error {
	if ic.hasDeferred {
		return fmt.Errorf("Cannot show modal after deferring")
	}

	mD := modal.GetModalResponse()
	ic.DeferChannel <- &mD

	return nil
}
