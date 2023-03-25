package components

import (
	"encoding/json"
	"github.com/tgwaffles/gladis/discord"
)

const (
	PrimaryButtonStyle ButtonStyle = iota + 1
	SecondaryButtonStyle
	SuccessButtonStyle
	DangerButtonStyle
	LinkButtonStyle
)

type Button struct {
	Style      ButtonStyle    `json:"style"`
	Label      *string        `json:"label,omitempty"`
	Emoji      *discord.Emoji `json:"emoji,omitempty"`
	CustomId   *string        `json:"custom_id,omitempty"`
	Url        *string        `json:"url,omitempty"`
	Disabled   *bool          `json:"disabled,omitempty"`
	ButtonType ComponentType  `json:"type"`
}

func (button Button) MarshalJSON() ([]byte, error) {
	type Alias Button

	var inner Alias
	inner = Alias(button)
	inner.ButtonType = ButtonType

	return json.Marshal(inner)
}

func (button Button) Type() ComponentType {
	return ButtonType
}

type ButtonStyle uint8
