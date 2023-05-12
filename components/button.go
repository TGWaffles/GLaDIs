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

func (button Button) Verify() error {
	if button.Style == LinkButtonStyle {
		if button.CustomId != nil {
			return ErrLinkButtonCannotHaveCustomId{button}
		}
		if button.Url == nil {
			return ErrLinkButtonMustHaveUrl{button}
		}
	} else {
		if button.Url != nil {
			return ErrNonLinkButtonCannotHaveUrl{button}
		}
		if button.CustomId == nil {
			return ErrComponentMustHaveCustomId{button}
		}
	}
	return nil
}

type ButtonStyle uint8
