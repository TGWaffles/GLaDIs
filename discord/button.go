package discord

import (
	"encoding/json"

	"github.com/tgwaffles/gladis/discord/button_style"
	"github.com/tgwaffles/gladis/discord/component_type"
)

type Button struct {
	Style      button_style.ButtonStyle     `json:"style"`
	Label      *string                      `json:"label,omitempty"`
	Emoji      *Emoji                       `json:"emoji,omitempty"`
	CustomId   *string                      `json:"custom_id,omitempty"`
	Url        *string                      `json:"url,omitempty"`
	Disabled   *bool                        `json:"disabled,omitempty"`
	ButtonType component_type.ComponentType `json:"type"`
	SKUId      *Snowflake                   `json:"sku_id"`
}

func (button Button) MarshalJSON() ([]byte, error) {
	type Alias Button

	var inner Alias
	inner = Alias(button)
	inner.ButtonType = component_type.Button

	return json.Marshal(inner)
}

func (button *Button) Type() component_type.ComponentType {
	return component_type.Button
}

func (button *Button) Verify() error {
	if button.Style == button_style.Link {
		if button.CustomId != nil {
			return ErrButtonCannotHaveCustomId{button}
		}
		if button.SKUId != nil {
			return ErrButtonCannotHaveSKUId{button}
		}
		if button.Url == nil {
			return ErrLinkButtonMustHaveUrl{button}
		}
	} else if button.Style == button_style.Premium {
		if button.CustomId != nil {
			return ErrButtonCannotHaveCustomId{button}
		}
		if button.Url != nil {
			return ErrNonLinkButtonCannotHaveUrl{button}
		}
		if button.SKUId == nil {
			return ErrPremiumButtonMustHaveSKUId{button}
		}
	} else {
		if button.Url != nil {
			return ErrNonLinkButtonCannotHaveUrl{button}
		}
		if button.CustomId == nil {
			return ErrComponentMustHaveCustomId{button}
		}
		if button.SKUId != nil {
			return ErrButtonCannotHaveSKUId{button}
		}
	}
	return nil
}
