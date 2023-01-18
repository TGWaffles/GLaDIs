package components

import "encoding/json"

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
	}
	return err
}

// MessageComponent can be made of any variables, so it's a map until we parse it into a specific component.
type MessageComponent interface {
	Type() ComponentType
}

type ComponentType uint8

type ActionRow struct {
	Components []MessageComponent `json:"components"`
}

func (a ActionRow) Type() ComponentType {
	return ActionRowType
}

type MessageComponentData struct {
	CustomId string        `json:"custom_id"`
	Type     ComponentType `json:"component_type"`
	Values   SelectOption  `json:"values"`
}
