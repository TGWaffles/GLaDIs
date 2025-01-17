package interactable

import (
	"fmt"
	"strings"

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

func (db *Button) SetContextId(id string) {
	db.Component.CustomId = helpers.Ptr(fmt.Sprintf("%s:%s", *db.Component.CustomId, id))
}

func (db *Button) GetContextId() *string {
	sp := strings.Split(*db.Component.CustomId, ":")

	if len(sp) < 2 {
		return nil
	}

	return &sp[len(sp)-1]
}
