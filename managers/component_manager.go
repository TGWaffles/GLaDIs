package managers

import (
	"fmt"

	"github.com/JackHumphries9/dapper-go/discord"
	"github.com/JackHumphries9/dapper-go/discord/component_type"
	"github.com/JackHumphries9/dapper-go/interactable"
)

type ComponentManager struct {
	components map[string]interactable.Component
}

func (dcm *ComponentManager) RouteInteraction(itx *discord.Interaction) (discord.InteractionResponse, error) {

	commandData := itx.Data.(*discord.MessageComponentData)

	if comp, ok := dcm.components[commandData.CustomId]; ok {
		if comp.Type() == component_type.Button {
			btn := comp.(interactable.Button)

			itc := interactable.InteractionContext{
				Interaction:  itx,
				DeferChannel: make(chan *discord.InteractionResponse),
			}

			go btn.OnPress(&itc)

			response := <-itc.DeferChannel

			return *response, nil
		}
	}

	return discord.InteractionResponse{}, fmt.Errorf("Cannot find interaction")
}

func (dcm *ComponentManager) Register(comp interactable.Component) {
	if comp.Type() == component_type.Button {
		customId := comp.GetComponent().(*discord.Button).CustomId

		dcm.components[*customId] = comp
	}
}

func NewDapperComponentManager() ComponentManager {
	return ComponentManager{
		components: make(map[string]interactable.Component, 0),
	}
}
