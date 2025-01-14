package dapper

import (
	"github.com/JackHumphries9/dapper-go/discord"
	"github.com/JackHumphries9/dapper-go/discord/component_type"
)

type DapperButtonOnPress func(itx *discord.Interaction)

type DapperComponent interface {
	Type() component_type.ComponentType
	OnInteract(itx *discord.Interaction) // The OnInteract method allows us to handle a generic interaction and then further distribute it to other handlers
	GetComponent() discord.MessageComponent
}
