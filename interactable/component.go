package interactable

import (
	"github.com/JackHumphries9/dapper-go/discord"
	"github.com/JackHumphries9/dapper-go/discord/component_type"
)

type ComponentOptions struct {
	CancelDefer bool
	Ephemeral   bool
}

type Component interface {
	Type() component_type.ComponentType
	OnInteract(itc *InteractionContext) // The OnInteract method allows us to handle a generic interaction and then further distribute it to other handlers
	GetComponent() discord.MessageComponent
	GetComponentOptions() ComponentOptions
}

type ComponentInstanceOptions struct {
	ID       string
	Disabled bool
}
