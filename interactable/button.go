package interactable

import (
	"github.com/JackHumphries9/dapper-go/discord"
	"github.com/JackHumphries9/dapper-go/discord/component_type"
)

type Button struct {
	Component discord.MessageComponent
	OnPress   InteractionHandler
}

func (db Button) Type() component_type.ComponentType {
	return component_type.Button
}

func (db Button) OnInteract(itc *InteractionContext) {
	db.OnPress(itc)
}

func (db Button) GetComponent() discord.MessageComponent {
	return db.Component
}
