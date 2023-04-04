package components

import (
	"encoding/json"
)

const (
	ShortTextInputStyle TextInputStyle = iota + 1
	LongTextInputStyle
)

type ModalSubmitData struct {
	CustomId   string             `json:"custom_id"`
	Components []MessageComponent `json:"components"`
}

type TextInputStyle uint8

func (m ModalSubmitData) UnmarshalJSON(data []byte) error {
	dataMap := make(map[string]interface{})
	err := json.Unmarshal(data, &dataMap)
	if err != nil {
		return err
	}

	m.CustomId = dataMap["custom_id"].(string)
	components := dataMap["components"].([]interface{})

	for _, component := range components {
		var messageComponent MessageComponentWrapper
		err = json.Unmarshal([]byte(component.(string)), &messageComponent)
		if err != nil {
			return err
		}
		m.Components = append(m.Components, messageComponent.component)
	}

	return nil
}

type TextInput struct {
	TextInputType ComponentType  `json:"type"`
	CustomId      string         `json:"custom_id"`
	Style         TextInputStyle `json:"style"`
	Label         string         `json:"label"`
	MinLength     int            `json:"min_length"`
	MaxLength     int            `json:"max_length"`
	Required      bool           `json:"required"`
	Value         string         `json:"value"`
	Placeholder   string         `json:"placeholder"`
}

func (t TextInput) MarshalJSON() ([]byte, error) {
	type Alias TextInput

	var inner Alias
	inner = Alias(t)
	inner.TextInputType = TextInputType

	return json.Marshal(inner)
}

func (t TextInput) Type() ComponentType {
	return TextInputType
}
