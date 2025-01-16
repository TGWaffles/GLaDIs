package managers

import (
	"fmt"

	"github.com/JackHumphries9/dapper-go/discord"
	"github.com/JackHumphries9/dapper-go/discord/component_type"
	"github.com/JackHumphries9/dapper-go/discord/interaction_callback_type"
	"github.com/JackHumphries9/dapper-go/helpers"
	"github.com/JackHumphries9/dapper-go/interactable"
)

type ComponentManager struct {
	components map[string]interactable.Component
}

func (dcm *ComponentManager) RouteInteraction(itx *discord.Interaction) (discord.InteractionResponse, error) {

	commandData := itx.Data.(*discord.MessageComponentData)

	if comp, ok := dcm.components[commandData.CustomId]; ok {
		itc := interactable.InteractionContext{
			Interaction:  itx,
			DeferChannel: make(chan *discord.InteractionResponse),
			HasDeferred:  !comp.GetComponentOptions().CancelDefer,
		}

		if comp.GetComponentOptions().Ephemeral {
			itc.SetEphemeral(true)
		}

		go comp.OnInteract(&itc)

		if comp.GetComponentOptions().CancelDefer {
			response := <-itc.DeferChannel

			return *response, nil
		}

		return discord.InteractionResponse{
			Type: interaction_callback_type.DeferredUpdateMessage,
			Data: &discord.MessageCallbackData{
				Flags: helpers.Ptr(int(itc.GetMessageFlags())),
			},
		}, nil

	}

	return discord.InteractionResponse{}, fmt.Errorf("Cannot find interaction: %s", commandData.CustomId)
}

func (dcm *ComponentManager) Register(comp interactable.Component) {
	if comp.Type() == component_type.Button {
		customId := comp.GetComponent().(*discord.Button).CustomId

		dcm.components[*customId] = comp
	}

	if comp.Type() == component_type.StringSelect || comp.Type() == component_type.RoleSelect || comp.Type() == component_type.UserSelect || comp.Type() == component_type.MentionableSelect || comp.Type() == component_type.ChannelSelect {
		customId := comp.GetComponent().(*discord.SelectMenu).CustomId

		dcm.components[customId] = comp
	}
}

func NewDapperComponentManager() ComponentManager {
	return ComponentManager{
		components: make(map[string]interactable.Component, 0),
	}
}
