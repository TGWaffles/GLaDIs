package components

import "github.com/tgwaffles/gladis/discord"

type SelectMenu struct {
	MenuType     ComponentType         `json:"type"`
	CustomId     string                `json:"custom_id"`
	Options      []SelectOption        `json:"options,omitempty"`
	ChannelTypes []discord.ChannelType `json:"channel_types,omitempty"`
	Placeholder  *string               `json:"placeholder,omitempty"`
	MinValues    *uint8                `json:"min_values,omitempty"`
	MaxValues    *uint8                `json:"max_values,omitempty"`
	Disabled     *bool                 `json:"disabled,omitempty"`
}

type SelectOption struct {
	Label       string         `json:"label"`
	Value       string         `json:"value"`
	Description *string        `json:"description,omitempty"`
	Emoji       *discord.Emoji `json:"emoji,omitempty"`
	Default     *bool          `json:"default,omitempty"`
}

func (selectMenu SelectMenu) Type() ComponentType {
	return selectMenu.MenuType
}

func (selectMenu SelectMenu) Verify() error {
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

	if selectMenu.Options == nil && selectMenu.MenuType == StringSelectType {
		return ErrStringSelectMenuMustHaveOptions{selectMenu}
	} else if selectMenu.Options != nil && selectMenu.MenuType != StringSelectType {
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
