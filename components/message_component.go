package components

import (
	"encoding/json"
	"github.com/tgwaffles/gladis/discord"
)

const (
	ActionRowType ComponentType = iota + 1
	ButtonType
	StringSelectType
	TextInputType
	UserSelectType
	RoleSelectType
	MentionableSelectType
	ChannelSelectType
)

type MessageComponentWrapper struct {
	component MessageComponent
}

func (m *MessageComponentWrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.component)
}

func (m *MessageComponentWrapper) UnmarshalJSON(data []byte) error {
	var component MessageComponent
	err := json.Unmarshal(data, &component)
	if err != nil {
		return err
	}
	switch component.Type() {
	case ActionRowType:
		var actionRow ActionRow
		err = json.Unmarshal(data, &actionRow)
		m.component = actionRow
		break
	case ButtonType:
		var button Button
		err = json.Unmarshal(data, &button)
		m.component = button
		break
	case StringSelectType, UserSelectType, RoleSelectType, MentionableSelectType, ChannelSelectType:
		var selectMenu SelectMenu
		err = json.Unmarshal(data, &selectMenu)
		m.component = selectMenu
		break
	case TextInputType:
		var textInput TextInput
		err = json.Unmarshal(data, &textInput)
		m.component = textInput
		break
	}
	return err
}

// MessageComponent can be made of any variables, so it's a map until we parse it into a specific component.
type MessageComponent interface {
	Type() ComponentType
	Verify() error
}

type ComponentType uint8

type ActionRow struct {
	RowType    ComponentType      `json:"type"`
	Components []MessageComponent `json:"components"`
}

func (a ActionRow) MarshalJSON() ([]byte, error) {
	type Alias ActionRow

	var inner Alias
	inner = Alias(a)
	inner.RowType = ActionRowType

	return json.Marshal(inner)
}

func (a ActionRow) UnmarshalJSON(data []byte) error {
	dataMap := make(map[string]interface{})
	err := json.Unmarshal(data, &dataMap)
	if err != nil {
		return err
	}

	a.RowType = ActionRowType
	components := dataMap["components"].([]interface{})
	for _, component := range components {
		var messageComponent MessageComponentWrapper
		err = json.Unmarshal([]byte(component.(string)), &messageComponent)
		if err != nil {
			return err
		}
		a.Components = append(a.Components, messageComponent.component)
	}

	return nil
}

func (a ActionRow) Type() ComponentType {
	return ActionRowType
}

func (a ActionRow) Verify() error {
	if len(a.Components) > 5 {
		return ErrTooManyComponents{a.Components}
	}
	for _, component := range a.Components {
		// Check the component is NOT another action row
		if component.Type() == ActionRowType {
			return ErrNestedActionRow{a.Components}
		}

		err := component.Verify()
		if err != nil {
			return err
		}
	}
	return nil
}

type MessageComponentData struct {
	CustomId string                `json:"custom_id"`
	Type     ComponentType         `json:"component_type"`
	Values   []string              `json:"values"`
	Resolved *discord.ResolvedData `json:"resolved,omitempty"`
}
