package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/tgwaffles/gladis/discord/interaction_callback_type"
)

type InteractionResponse struct {
	Type interaction_callback_type.InteractionCallbackType `json:"type"`
	Data InteractionCallbackData                           `json:"data,omitempty"`
}

type InteractionCallbackData interface {
}

type MessageResponse interface {
	GetMessageAttachments() []MessageAttachment
	SetDiscordAttachments([]Attachment)
}

func ensureMessageAttachments(response MessageResponse) {
	existingAttachments := response.GetMessageAttachments()
	discordAttachments := make([]Attachment, 0, len(existingAttachments))

	for i, attachment := range existingAttachments {
		discordAttachments = append(discordAttachments, attachment.ToDiscordAttachment(Snowflake(i)))
	}

	response.SetDiscordAttachments(discordAttachments)
}

type MessageCallbackData struct {
	TTS                *bool               `json:"tts,omitempty"`
	Content            *string             `json:"content,omitempty"`
	Embeds             []Embed             `json:"embeds,omitempty"`
	AllowedMentions    *AllowedMentions    `json:"allowed_mentions,omitempty"`
	Flags              *int                `json:"flags,omitempty"`
	Components         []MessageComponent  `json:"components,omitempty"`
	Attachments        []MessageAttachment `json:"-"`
	DiscordAttachments []Attachment        `json:"attachments,omitempty"`
}

func (data *MessageCallbackData) GetMessageAttachments() []MessageAttachment {
	return data.Attachments
}

func (data *MessageCallbackData) SetDiscordAttachments(attachments []Attachment) {
	data.DiscordAttachments = attachments
}

type ResponseEditData struct {
	Content            *string             `json:"content,omitempty"`
	Embeds             []Embed             `json:"embeds,omitempty"`
	AllowedMentions    *AllowedMentions    `json:"allowed_mentions"`
	Components         []MessageComponent  `json:"components"`
	Attachments        []MessageAttachment `json:"-"`
	DiscordAttachments []Attachment        `json:"attachments"`
}

func (data *ResponseEditData) GetMessageAttachments() []MessageAttachment {
	return data.Attachments
}

func (data *ResponseEditData) SetDiscordAttachments(attachments []Attachment) {
	data.DiscordAttachments = attachments
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

type HTTPResponse struct {
	StatusCode uint
	Body       []byte
	Headers    map[string]string
}

func (res HTTPResponse) WriteResponse(w http.ResponseWriter) {
	w.WriteHeader(int(res.StatusCode))

	for key, val := range res.Headers {
		w.Header().Add(key, val)
	}

	if len(res.Body) > 0 {
		w.Write(res.Body)
	}
}

func (ir InteractionResponse) ToHttpResponse() HTTPResponse {
	data, err := json.Marshal(ir)
	if err != nil {
		fmt.Println("Error marshalling interaction response:", err)
		return HTTPResponse{
			StatusCode: 500,
		}
	}

	if messageResponse, ok := ir.Data.(MessageResponse); ok {
		// It's a message response
		if len(messageResponse.GetMessageAttachments()) > 0 {
			// It has attachments
			var buffer bytes.Buffer
			var contentType string
			contentType, err = WriteFormResponse(&buffer, data, messageResponse.GetMessageAttachments())
			if err != nil {
				fmt.Println("Error writing form response:", err)
				return HTTPResponse{
					StatusCode: 500,
				}
			}
			return HTTPResponse{
				StatusCode: 200,
				Body:       buffer.Bytes(),
				Headers: map[string]string{
					"Content-Type": contentType,
				},
			}
		}
	}
	return HTTPResponse{
		StatusCode: 200,
		Body:       data,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}

func WriteFormResponse(bodyWriter io.Writer, responseJson []byte, attachments []MessageAttachment) (contentType string, err error) {

	writer := multipart.NewWriter(bodyWriter)

	// Write the main message content first

	mainMessageHeaders := make(map[string][]string)
	mainMessageHeaders["Content-Disposition"] = []string{`form-data; name="payload_json"`}
	mainMessageHeaders["Content-Type"] = []string{`application/json`}

	mainMessagePart, err := writer.CreatePart(mainMessageHeaders)
	if err != nil {
		fmt.Println("Error creating JSON field:", err)
		return "", fmt.Errorf("error creating JSON field")
	}
	_, err = mainMessagePart.Write(responseJson)
	if err != nil {
		fmt.Println("Error writing JSON field:", err)
		return "", fmt.Errorf("error writing JSON field")
	}

	// Write attachments
	for i, attachment := range attachments {
		attachmentPartHeaders := make(map[string][]string)
		attachmentPartHeaders["Content-Disposition"] = []string{fmt.Sprintf(`form-data; name="files[%d]"; filename="%s"`, i, attachment.GetFileName())}
		attachmentPartHeaders["Content-Type"] = []string{attachment.GetContentType()}

		attachmentPart, err := writer.CreatePart(attachmentPartHeaders)

		if err != nil {
			return "", fmt.Errorf("error creating form file %s: %w", attachment.GetFileName(), err)
		}

		_, err = attachmentPart.Write(attachment.GetBytes())
		if err != nil {
			return "", fmt.Errorf("error writing file bytes for %s: %w", attachment.GetFileName(), err)
		}
	}

	// Close the writer
	err = writer.Close()
	if err != nil {
		return "", fmt.Errorf("error closing multipart writer: %w", err)
	}

	return writer.FormDataContentType(), nil
}

func ConvertDataToBodyBytes(data InteractionCallbackData) (body []byte, contentType string, err error) {
	body, err = json.Marshal(data)
	if err != nil {
		return nil, "", fmt.Errorf("error marshalling data to body bytes: %w", err)
	}

	if messageResponse, ok := data.(MessageResponse); ok {
		if len(messageResponse.GetMessageAttachments()) > 0 {
			// If the data is a message response with attachments, we need to write it as a form
			var buffer bytes.Buffer
			contentType, err = WriteFormResponse(&buffer, body, messageResponse.GetMessageAttachments())
			if err != nil {
				return nil, "", fmt.Errorf("error writing form response: %w", err)
			}
			return buffer.Bytes(), contentType, nil
		}
	}

	return body, "application/json", nil
}
