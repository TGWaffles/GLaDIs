package interactable

import (
	"github.com/JackHumphries9/dapper-go/discord"
	"github.com/JackHumphries9/dapper-go/discord/component_type"
	"github.com/JackHumphries9/dapper-go/helpers"
)

type Button struct {
	Component        *discord.Button
	ComponentOptions ComponentOptions
	OnPress          InteractionHandler
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

func (db Button) GetComponentOptions() ComponentOptions {
	return db.ComponentOptions
}

func (db *Button) CreateComponentInstance(opts ComponentInstanceOptions) discord.MessageComponent {
	return &discord.Button{
		Style:      db.Component.Style,
		Label:      db.Component.Label,
		Emoji:      db.Component.Emoji,
		Url:        db.Component.Url,
		Disabled:   &opts.Disabled,
		ButtonType: db.Component.ButtonType,
		CustomId:   helpers.Ptr(*db.Component.CustomId + ":" + opts.ID),
	}
}
