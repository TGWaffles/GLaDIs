package components

import "github.com/tgwaffles/gladis/discord"

type SelectMenu struct {
	MenuType     ComponentType          `json:"type"`
	CustomId     string                 `json:"custom_id"`
	Options      *[]SelectOption        `json:"options,omitempty"`
	ChannelTypes *[]discord.ChannelType `json:"channel_types,omitempty"`
	Placeholder  *string                `json:"placeholder,omitempty"`
	MinValues    *uint8                 `json:"min_values,omitempty"`
	MaxValues    *uint8                 `json:"max_values,omitempty"`
	Disabled     *bool                  `json:"disabled,omitempty"`
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
