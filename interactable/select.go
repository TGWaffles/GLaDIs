package interactable

import (
	"github.com/JackHumphries9/dapper-go/discord"
	"github.com/JackHumphries9/dapper-go/discord/component_type"
)

type Select struct {
	Component *discord.SelectMenu
	OnSelect  InteractionHandler
}

func (db Select) Type() component_type.ComponentType {
	return db.Component.MenuType
}

func (db Select) OnInteract(itc *InteractionContext) {
	db.OnSelect(itc)
}

func (db Select) GetComponent() discord.MessageComponent {
	return db.Component
}
