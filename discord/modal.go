package discord

import (
	"encoding/json"
	"github.com/tgwaffles/gladis/discord/component_type"
	"github.com/tgwaffles/gladis/discord/text_input_style"
)

type ModalSubmitData struct {
	CustomId   string             `json:"custom_id"`
	Components []MessageComponent `json:"components"`
}

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
	TextInputType component_type.ComponentType    `json:"type"`
	CustomId      string                          `json:"custom_id"`
	Style         text_input_style.TextInputStyle `json:"style"`
	Label         string                          `json:"label"`
	MinLength     int                             `json:"min_length"`
	MaxLength     int                             `json:"max_length"`
	Required      bool                            `json:"required"`
	Value         string                          `json:"value"`
	Placeholder   string                          `json:"placeholder"`
}

func (t *TextInput) MarshalJSON() ([]byte, error) {
	type Alias TextInput

	var inner Alias
	inner = Alias(*t)
	inner.TextInputType = component_type.TextInput

	return json.Marshal(inner)
}

func (t *TextInput) Type() component_type.ComponentType {
	return component_type.TextInput
}

func (t *TextInput) Verify() error {
	if t.CustomId == "" {
		return ErrComponentMustHaveCustomId{t}
	}
	if len(t.CustomId) > 100 {
		return ErrInvalidPropertyLength{
			Component:      t,
			PropertyName:   "custom_id",
			MaxLength:      100,
			MinLength:      1,
			PropertyLength: len(t.CustomId),
			PropertyValue:  t.CustomId,
		}
	}
	if t.Style == 0 {
		return ErrComponentMustHaveStyle{t}
	}

	if t.Label == "" {
		return ErrComponentMissingProperty{t, "label"}
	}
	if len(t.Label) > 45 {
		return ErrInvalidPropertyLength{
			Component:      t,
			PropertyName:   "label",
			MaxLength:      45,
			MinLength:      1,
			PropertyLength: len(t.Label),
			PropertyValue:  t.Label,
		}
	}

	if len(t.Value) > 4000 {
		return ErrInvalidPropertyLength{
			Component:      t,
			PropertyName:   "value",
			MaxLength:      4000,
			MinLength:      1,
			PropertyLength: len(t.Value),
			PropertyValue:  t.Value,
		}
	}

	if len(t.Placeholder) > 100 {
		return ErrInvalidPropertyLength{
			Component:      t,
			PropertyName:   "placeholder",
			MaxLength:      100,
			MinLength:      1,
			PropertyLength: len(t.Placeholder),
			PropertyValue:  t.Placeholder,
		}
	}

	return nil
}
