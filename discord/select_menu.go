package discord

import (
	"github.com/JackHumphries9/dapper-go/discord/channel_type"
	"github.com/JackHumphries9/dapper-go/discord/component_type"
)

type SelectMenu struct {
	MenuType     component_type.ComponentType `json:"type"`
	CustomId     string                       `json:"custom_id"`
	Options      []SelectOption               `json:"options,omitempty"`
	ChannelTypes []channel_type.ChannelType   `json:"channel_types,omitempty"`
	Placeholder  *string                      `json:"placeholder,omitempty"`
	MinValues    *uint8                       `json:"min_values,omitempty"`
	MaxValues    *uint8                       `json:"max_values,omitempty"`
	Disabled     *bool                        `json:"disabled,omitempty"`
}

type SelectOption struct {
	Label       string  `json:"label"`
	Value       string  `json:"value"`
	Description *string `json:"description,omitempty"`
	Emoji       *Emoji  `json:"emoji,omitempty"`
	Default     *bool   `json:"default,omitempty"`
}

func (selectMenu *SelectMenu) Type() component_type.ComponentType {
	return selectMenu.MenuType
}

func (selectMenu *SelectMenu) Verify() error {
	if selectMenu.CustomId == "" {
		return ErrComponentMustHaveCustomId{selectMenu}
	}
	if len(selectMenu.CustomId) > 100 {
		return ErrInvalidPropertyLength{
			Component:      selectMenu,
			PropertyName:   "custom_id",
			MaxLength:      100,
			MinLength:      1,
			PropertyLength: len(selectMenu.CustomId),
			PropertyValue:  selectMenu.CustomId,
		}
	}

	if selectMenu.Options == nil && selectMenu.MenuType == component_type.StringSelect {
		return ErrStringSelectMenuMustHaveOptions{selectMenu}
	} else if selectMenu.Options != nil && selectMenu.MenuType != component_type.StringSelect {
		return ErrNonStringSelectMenuCannotHaveOptions{selectMenu}
	}

	if len(selectMenu.Options) > 25 {
		return ErrInvalidPropertyLength{
			Component:      selectMenu,
			PropertyName:   "options",
			MaxLength:      25,
			MinLength:      1,
			PropertyLength: len(selectMenu.Options),
			PropertyValue:  selectMenu.Options,
		}
	}

	for _, option := range selectMenu.Options {
		err := option.Verify()
		if err != nil {
			return err
		}
	}

	return nil
}

func (selectOption SelectOption) Verify() error {
	if selectOption.Label == "" {
		return ErrSelectOptionMustHaveLabel{selectOption}
	}
	if len(selectOption.Label) > 100 {
		return ErrInvalidPropertyLength{
			Component:      selectOption,
			PropertyName:   "label",
			MaxLength:      100,
			MinLength:      1,
			PropertyLength: len(selectOption.Label),
			PropertyValue:  selectOption.Label,
		}
	}
	if selectOption.Value == "" {
		return ErrSelectOptionMustHaveValue{selectOption}
	}
	if len(selectOption.Value) > 100 {
		return ErrInvalidPropertyLength{
			Component:      selectOption,
			PropertyName:   "value",
			MaxLength:      100,
			MinLength:      1,
			PropertyLength: len(selectOption.Value),
			PropertyValue:  selectOption.Value,
		}
	}
	if selectOption.Description != nil && len(*selectOption.Description) > 100 {
		return ErrInvalidPropertyLength{
			Component:      selectOption,
			PropertyName:   "description",
			MaxLength:      100,
			MinLength:      1,
			PropertyLength: len(*selectOption.Description),
			PropertyValue:  *selectOption.Description,
		}
	}
	return nil
}
