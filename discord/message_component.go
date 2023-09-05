package discord

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/tgwaffles/gladis/discord/component_type"
	"strconv"
)

type MessageComponentWrapper struct {
	component MessageComponent
}

func (m *MessageComponentWrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.component)
}

func (m *MessageComponentWrapper) UnmarshalDynamoDBAttributeValue(value *dynamodb.AttributeValue) (err error) {
	componentMap := value.M
	if componentMap == nil {
		return fmt.Errorf("error unmarshalling message component into map, map empty: %v", value)
	}

	givenComponentType := componentMap["type"]
	if givenComponentType == nil || givenComponentType.N == nil {
		return fmt.Errorf("error unmarshalling message component into map, type empty: %v", value)
	}

	componentTypeInt, err := strconv.ParseInt(*givenComponentType.N, 10, 8)
	if err != nil {
		return fmt.Errorf("error unmarshalling message component into map, type not a number: %v", *givenComponentType.N)
	}

	switch component_type.ComponentType(componentTypeInt) {
	case component_type.ActionRow:
		var actionRow ActionRow
		err = dynamodbattribute.UnmarshalMap(componentMap, &actionRow)
		m.component = &actionRow
		break
	case component_type.Button:
		var button Button
		err = dynamodbattribute.UnmarshalMap(componentMap, &button)
		m.component = &button
		break
	case component_type.StringSelect, component_type.UserSelect, component_type.RoleSelect, component_type.MentionableSelect, component_type.ChannelSelect:
		var selectMenu SelectMenu
		err = dynamodbattribute.UnmarshalMap(componentMap, &selectMenu)
		m.component = &selectMenu
		break
	case component_type.TextInput:
		var textInput TextInput
		err = dynamodbattribute.UnmarshalMap(componentMap, &textInput)
		m.component = &textInput
		break
	}

	if err != nil {
		return fmt.Errorf("error unmarshalling component to struct, %w", err)
	}
	return nil
}

func (m *MessageComponentWrapper) UnmarshalJSON(data []byte) error {
	var componentMap map[string]interface{}
	err := json.Unmarshal(data, &componentMap)
	if err != nil {
		return err
	}
	switch component_type.ComponentType(componentMap["type"].(float64)) {
	case component_type.ActionRow:
		var actionRow ActionRow
		err = json.Unmarshal(data, &actionRow)
		m.component = &actionRow
		break
	case component_type.Button:
		var button Button
		err = json.Unmarshal(data, &button)
		m.component = &button
		break
	case component_type.StringSelect, component_type.UserSelect, component_type.RoleSelect, component_type.MentionableSelect, component_type.ChannelSelect:
		var selectMenu SelectMenu
		err = json.Unmarshal(data, &selectMenu)
		m.component = &selectMenu
		break
	case component_type.TextInput:
		var textInput TextInput
		err = json.Unmarshal(data, &textInput)
		m.component = &textInput
		break
	}
	return err
}

// MessageComponent can be made of any variables, so it's a map until we parse it into a specific component.
type MessageComponent interface {
	Type() component_type.ComponentType
	Verify() error
}

type ActionRow struct {
	RowType    component_type.ComponentType `json:"type"`
	Components []MessageComponent           `json:"components"`
}

func (a *ActionRow) MarshalJSON() ([]byte, error) {
	type Alias ActionRow

	var inner Alias
	inner = Alias(*a)
	inner.RowType = component_type.ActionRow

	return json.Marshal(inner)
}

func (a *ActionRow) UnmarshalDynamoDBAttributeValue(value *dynamodb.AttributeValue) (err error) {
	if a.Components == nil {
		a.Components = make([]MessageComponent, 0)
	}
	dataMap := value.M

	componentsListAttr := dataMap["components"]
	if componentsListAttr == nil || len(componentsListAttr.L) == 0 {
		return nil
	}

	components := make([]MessageComponent, len(componentsListAttr.L))

	for i, component := range componentsListAttr.L {
		var messageComponent MessageComponentWrapper
		err = dynamodbattribute.UnmarshalMap(component.M, &messageComponent)
		if err != nil {
			return fmt.Errorf("error unmarshalling message component: %w", err)
		}
		components[i] = messageComponent.component
	}

	a.Components = components

	return nil
}

func (a *ActionRow) UnmarshalJSON(data []byte) error {
	if a.Components == nil {
		a.Components = make([]MessageComponent, 0)
	}
	dataMap := make(map[string]json.RawMessage)
	err := json.Unmarshal(data, &dataMap)
	if err != nil {
		return err
	}

	a.RowType = component_type.ActionRow
	var components []json.RawMessage
	err = json.Unmarshal(dataMap["components"], &components)
	if err != nil {
		return err
	}
	for _, component := range components {
		var messageComponent MessageComponentWrapper
		err = json.Unmarshal(component, &messageComponent)
		if err != nil {
			return err
		}
		a.Components = append(a.Components, messageComponent.component)
	}

	return nil
}

func (a *ActionRow) Type() component_type.ComponentType {
	return component_type.ActionRow
}

func (a *ActionRow) Verify() error {
	if len(a.Components) > 5 {
		return ErrTooManyComponents{a.Components}
	}
	for _, component := range a.Components {
		// Check the component is NOT another action row
		if component.Type() == component_type.ActionRow {
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
	CustomId string                       `json:"custom_id"`
	Type     component_type.ComponentType `json:"component_type"`
	Values   []string                     `json:"values"`
	Resolved *ResolvedData                `json:"resolved,omitempty"`
}
