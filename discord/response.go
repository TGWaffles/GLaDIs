package discord

import (
	"encoding/json"
	"fmt"

	"github.com/JackHumphries9/dapper-go/discord/interaction_callback_type"
)

type InteractionResponse struct {
	Type interaction_callback_type.InteractionCallbackType `json:"type"`
	Data InteractionCallbackData                           `json:"data,omitempty"`
}

type HTTPResponse struct {
	StatusCode uint
	Body       string
	Headers    map[string]string
}

func (ir InteractionResponse) ToHttpResponse() HTTPResponse {
	data, err := json.Marshal(ir)
	if err != nil {
		fmt.Println("Error marshalling interaction response:", err)
		return HTTPResponse{
			StatusCode: 500,
		}
	}
	return HTTPResponse{
		StatusCode: 200,
		Body:       string(data),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}

type InteractionCallbackData interface {
}

type MessageCallbackData struct {
	TTS             *bool              `json:"tts,omitempty"`
	Content         *string            `json:"content,omitempty"`
	Embeds          []Embed            `json:"embeds,omitempty"`
	AllowedMentions *AllowedMentions   `json:"allowed_mentions,omitempty"`
	Flags           *int               `json:"flags,omitempty"`
	Components      []MessageComponent `json:"components,omitempty"`
	Attachments     []Attachment       `json:"attachments,omitempty"`
}

type ResponseEditData struct {
	Content         *string            `json:"content,omitempty"`
	Embeds          []Embed            `json:"embeds,omitempty"`
	AllowedMentions *AllowedMentions   `json:"allowed_mentions"`
	Components      []MessageComponent `json:"components"`
}

func (data ResponseEditData) Verify() error {
	if data.Content != nil && len(*data.Content) > 2000 {
		return fmt.Errorf("content cannot be longer than 2000 characters (you have %d)", len(*data.Content))
	}

	if len(data.Embeds) > 10 {
		return fmt.Errorf("too many embeds (max 10, you have %d)", len(data.Embeds))
	}

	for _, component := range data.Components {
		if err := component.Verify(); err != nil {
			return err
		}
	}

	for _, embed := range data.Embeds {
		if err := embed.Verify(); err != nil {
			return err
		}
	}

	return nil
}

type AutocompleteCallbackData struct {
	Choices []AutoCompleteChoice `json:"choices"`
}

type ModalCallback struct {
	CustomId   string             `json:"custom_id"`
	Title      string             `json:"title"`
	Components []MessageComponent `json:"components"`
}

func CreatePongResponse() InteractionResponse {
	return InteractionResponse{
		Type: interaction_callback_type.Pong,
	}
}

func CreateDeferMessageResponse() InteractionResponse {
	return InteractionResponse{
		Type: interaction_callback_type.DeferredChannelMessageWithSource,
	}
}

func CreateDeferEditResponse() InteractionResponse {
	return InteractionResponse{
		Type: interaction_callback_type.DeferredUpdateMessage,
	}
}
