package discord

import (
	"encoding/json"
	"fmt"
	"github.com/tgwaffles/gladis/discord/allowed_mention_type"
)

type ErrTooManyComponents struct {
	Components []MessageComponent
}

func (e ErrTooManyComponents) Error() string {
	var componentsText string
	componentsMarshalled, err := json.Marshal(e.Components)
	if err != nil {
		componentsText = fmt.Sprintf("%v", e.Components)
	} else {
		componentsText = string(componentsMarshalled)
	}
	return fmt.Sprintf("too many components in ActionRow (max 5, you have: %d)\nComponents:%s\n", len(e.Components), componentsText)
}

type ErrNestedActionRow struct {
	Components []MessageComponent
}

func (e ErrNestedActionRow) Error() string {
	var componentsText string
	componentsMarshalled, err := json.Marshal(e.Components)
	if err != nil {
		componentsText = fmt.Sprintf("%v", e.Components)
	} else {
		componentsText = string(componentsMarshalled)
	}
	return fmt.Sprintf("nested ActionRow in ActionRow\nComponents:%s\n", componentsText)
}

type ErrLinkButtonCannotHaveCustomId struct {
	Component *Button
}

func (e ErrLinkButtonCannotHaveCustomId) Error() string {
	var componentText string
	componentMarshalled, err := json.Marshal(e.Component)
	if err != nil {
		componentText = fmt.Sprintf("%v", e.Component)
	} else {
		componentText = string(componentMarshalled)
	}
	return fmt.Sprintf("link button cannot have custom id\nComponent:%s\n", componentText)
}

type ErrLinkButtonMustHaveUrl struct {
	Component *Button
}

func (e ErrLinkButtonMustHaveUrl) Error() string {
	var componentText string
	componentMarshalled, err := json.Marshal(e.Component)
	if err != nil {
		componentText = fmt.Sprintf("%v", e.Component)
	} else {
		componentText = string(componentMarshalled)
	}
	return fmt.Sprintf("link button must have url\nComponent:%s\n", componentText)
}

type ErrNonLinkButtonCannotHaveUrl struct {
	Component *Button
}

func (e ErrNonLinkButtonCannotHaveUrl) Error() string {
	var componentText string
	componentMarshalled, err := json.Marshal(e.Component)
	if err != nil {
		componentText = fmt.Sprintf("%v", e.Component)
	} else {
		componentText = string(componentMarshalled)
	}
	return fmt.Sprintf("non-link button cannot have url\nComponent:%s\n", componentText)
}

type ErrComponentMissingProperty struct {
	Component    MessageComponent
	PropertyName string
}

func (e ErrComponentMissingProperty) Error() string {
	var componentText string
	componentMarshalled, err := json.Marshal(e.Component)
	if err != nil {
		componentText = fmt.Sprintf("%v", e.Component)
	} else {
		componentText = string(componentMarshalled)
	}
	return fmt.Sprintf("component type %d must have %s\nComponent:%s\n", e.Component.Type(), e.PropertyName, componentText)
}

type ErrComponentMustHaveStyle struct {
	Component MessageComponent
}

func (e ErrComponentMustHaveStyle) Error() string {
	var componentText string
	componentMarshalled, err := json.Marshal(e.Component)
	if err != nil {
		componentText = fmt.Sprintf("%v", e.Component)
	} else {
		componentText = string(componentMarshalled)
	}

	return fmt.Sprintf("Component type %d must have style\nComponent:%s\n", e.Component.Type(), componentText)
}

type ErrComponentMustHaveCustomId struct {
	Component MessageComponent
}

func (e ErrComponentMustHaveCustomId) Error() string {
	var componentText string
	componentMarshalled, err := json.Marshal(e.Component)
	if err != nil {
		componentText = fmt.Sprintf("%v", e.Component)
	} else {
		componentText = string(componentMarshalled)
	}
	return fmt.Sprintf("Component type %d must have custom id\nComponent:%s\n", e.Component.Type(), componentText)
}

type ErrInvalidPropertyLength struct {
	Component      interface{}
	PropertyName   string
	MaxLength      int
	MinLength      int
	PropertyLength int
	PropertyValue  interface{}
}

func (e ErrInvalidPropertyLength) Error() string {
	var componentText string
	componentMarshalled, err := json.Marshal(e.Component)
	if err != nil {
		componentText = fmt.Sprintf("%v", e.Component)
	} else {
		componentText = string(componentMarshalled)
	}
	return fmt.Sprintf(
		"invalid length (%d) for property %s\nComponent:%s\nPropertyValue:%s\nMinLength:%d\nMaxLength:%d\n",
		e.PropertyLength, e.PropertyName, componentText, e.PropertyValue, e.MinLength, e.MaxLength,
	)
}

type ErrStringSelectMenuMustHaveOptions struct {
	Component *SelectMenu
}

func (e ErrStringSelectMenuMustHaveOptions) Error() string {
	var componentText string
	componentMarshalled, err := json.Marshal(e.Component)
	if err != nil {
		componentText = fmt.Sprintf("%v", e.Component)
	} else {
		componentText = string(componentMarshalled)
	}
	return fmt.Sprintf("string select menu must have options\nComponent:%s\n", componentText)
}

type ErrNonStringSelectMenuCannotHaveOptions struct {
	Component *SelectMenu
}

func (e ErrNonStringSelectMenuCannotHaveOptions) Error() string {
	var componentText string
	componentMarshalled, err := json.Marshal(e.Component)
	if err != nil {
		componentText = fmt.Sprintf("%v", e.Component)
	} else {
		componentText = string(componentMarshalled)
	}
	return fmt.Sprintf("non-string select menu cannot have options\nComponent:%s\n", componentText)
}

type ErrSelectOptionMustHaveLabel struct {
	Option SelectOption
}

func (e ErrSelectOptionMustHaveLabel) Error() string {
	var optionRepresentation string
	optionMarshalled, err := json.Marshal(e.Option)
	if err != nil {
		optionRepresentation = fmt.Sprintf("%v", e.Option)
	} else {
		optionRepresentation = string(optionMarshalled)
	}

	return fmt.Sprintf("select option must have label\nOption:%s\n", optionRepresentation)
}

type ErrSelectOptionMustHaveValue struct {
	Option SelectOption
}

func (e ErrSelectOptionMustHaveValue) Error() string {
	var optionRepresentation string
	optionMarshalled, err := json.Marshal(e.Option)
	if err != nil {
		optionRepresentation = fmt.Sprintf("%v", e.Option)
	} else {
		optionRepresentation = string(optionMarshalled)
	}

	return fmt.Sprintf("select option must have value\nOption:%s\n", optionRepresentation)
}

type ErrInvalidParseGroup struct {
	InvalidGroup  allowed_mention_type.AllowedMentionType
	MentionObject AllowedMentions
}

func (e ErrInvalidParseGroup) Error() string {
	var jsonRepresentation string
	jsonMarshalled, err := json.Marshal(e.MentionObject)
	if err != nil {
		jsonRepresentation = fmt.Sprintf("%v", e.MentionObject)
	} else {
		jsonRepresentation = string(jsonMarshalled)
	}

	return fmt.Sprintf("invalid parse group \"%s\"\nif you have this parse group, you cannot also provide options for the %s field in the allowed mentions object: %s\n", e.InvalidGroup, e.InvalidGroup, jsonRepresentation)
}
