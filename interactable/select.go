package interactable

import (
	"github.com/JackHumphries9/dapper-go/discord"
	"github.com/JackHumphries9/dapper-go/discord/component_type"
)

type Select struct {
	Component        *discord.SelectMenu
	ComponentOptions ComponentOptions
	OnSelect         InteractionHandler
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

func (db Select) GetComponentOptions() ComponentOptions {
	return db.ComponentOptions
}

// TODO: Re add this but implement it at a lower level

// func (db *Select) CreateComponentInstance(opts ComponentInstanceOptions) discord.MessageComponent {
// 	return &discord.SelectMenu{
// 		MenuType:     db.Component.MenuType,
// 		Options:      db.Component.Options,
// 		ChannelTypes: db.Component.ChannelTypes,
// 		Placeholder:  db.Component.Placeholder,
// 		MinValues:    db.Component.MinValues,
// 		MaxValues:    db.Component.MaxValues,
// 		Disabled:     &opts.Disabled,
// 		CustomId:     db.Component.CustomId + ":" + opts.ID,
// 	}
// }
