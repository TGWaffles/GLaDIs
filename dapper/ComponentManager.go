package dapper

import (
	"fmt"

	"github.com/JackHumphries9/dapper-go/discord"
	"github.com/JackHumphries9/dapper-go/discord/component_type"
	"github.com/JackHumphries9/dapper-go/discord/interaction_callback_type"
)

type DapperComponentManager struct {
	components []DapperComponent
}

func (dcm *DapperComponentManager) Register(comp DapperComponent) {
	dcm.components = append(dcm.components, comp)
}

func (dcm *DapperComponentManager) RouteInteraction(itx *discord.Interaction) (discord.InteractionResponse, error) {
	for _, comp := range dcm.components {

		if comp.Type() == component_type.Button {
			data := comp.GetComponent().(*discord.Button)

			// We can assume it's a message component
			itx_data := itx.Data.(*discord.MessageComponentData)

			if *data.CustomId == itx_data.CustomId {
				go comp.OnInteract(itx)
				return discord.InteractionResponse{
					Type: interaction_callback_type.DeferredUpdateMessage,
					Data: &discord.MessageCallbackData{},
				}, nil
			}

		}
	}

	return discord.InteractionResponse{}, fmt.Errorf("Cannot find interaction")
}

func NewDapperComponentManager() DapperComponentManager {
	return DapperComponentManager{
		components: make([]DapperComponent, 0),
	}
}
