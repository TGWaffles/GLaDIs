package interactable

import (
	"fmt"
	"strings"

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

func (db *Select) SetContextId(id string) {
	db.Component.CustomId = fmt.Sprintf("%s:%s", db.Component.CustomId, id)
}

func (db *Select) GetContextId() string {
	sp := strings.Split(db.Component.CustomId, ":")

	return sp[len(sp)-1]
}
