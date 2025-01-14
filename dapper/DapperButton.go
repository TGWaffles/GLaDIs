package dapper

import (
	"github.com/JackHumphries9/dapper-go/discord"
	"github.com/JackHumphries9/dapper-go/discord/component_type"
)

type DapperButton struct {
	Component discord.MessageComponent
	OnPress   func(itx *discord.Interaction)
}

func (db DapperButton) Type() component_type.ComponentType {
	return component_type.Button
}

func (db DapperButton) OnInteract(itx *discord.Interaction) {
	db.OnPress(itx)
}

func (db DapperButton) GetComponent() discord.MessageComponent {
	return db.Component
}
