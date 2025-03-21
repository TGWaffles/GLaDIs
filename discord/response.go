package discord

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/tgwaffles/gladis/discord/interaction_callback_type"
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

func (res HTTPResponse) WriteResponse(w http.ResponseWriter) {
	w.WriteHeader(int(res.StatusCode))

	for key, val := range res.Headers {
		w.Header().Add(key, val)
	}

	if res.Body != "" {
		fmt.Fprint(w, res.Body)
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
	Content           *string             `json:"content,omitempty"`
	Embeds            []Embed             `json:"embeds,omitempty"`
	AllowedMentions   *AllowedMentions    `json:"allowed_mentions"`
	Components        []MessageComponent  `json:"components"`
	Attachments       []MessageAttachment `json:"-"`
	DiscordAttachment []Attachment        `json:"attachments"`
}

func (data *ResponseEditData) ParseAttachments() {
	data.DiscordAttachment = make([]Attachment, 0, len(data.Attachments))

	for i, attachment := range data.Attachments {
		data.DiscordAttachment = append(data.DiscordAttachment, attachment.ToDiscordAttachment(Snowflake(i)))
	}
}

func (data *ResponseEditData) BuildHTTPRequest(ctx context.Context, method string, url string) (*http.Request, error) {

	err := data.Verify()
	if err != nil {
		return nil, fmt.Errorf("error verifying edit data: %w", err)
	}

	body, err := json.Marshal(*data)
	if err != nil {
		return nil, fmt.Errorf("error marshaling data to JSON: %w", err)
	}

	var request *http.Request
	if len(data.Attachments) > 0 {
		data.ParseAttachments()

		// Here we do form stuff
		var requestBody bytes.Buffer
		writer := multipart.NewWriter(&requestBody)

		partHeaders := make(map[string][]string)
		partHeaders["Content-Disposition"] = []string{`form-data; name="payload_json"`}
		partHeaders["Content-Type"] = []string{`application/json`}

		part, err := writer.CreatePart(partHeaders)
		if err != nil {
			fmt.Println("Error creating JSON field:", err)
			return nil, fmt.Errorf("Failed to create JSON field")
		}
		_, err = part.Write(body)
		if err != nil {
			fmt.Println("Error writing JSON field:", err)
			return nil, fmt.Errorf("Failed to write json field")
		}

		//Write attachments
		for i, attachment := range data.Attachments {

			partHeaders := make(map[string][]string)
			partHeaders["Content-Disposition"] = []string{fmt.Sprintf(`form-data; name="files[%d]"; filename="%s"`, i, attachment.GetFileName())}
			partHeaders["Content-Type"] = []string{attachment.GetContentType()}

			part, err := writer.CreatePart(partHeaders)

			if err != nil {
				return nil, fmt.Errorf("error creating form file %s: %w", attachment.GetFileName(), err)
			}

			_, err = part.Write(attachment.GetBytes())
			if err != nil {
				return nil, fmt.Errorf("error writing file bytes for %s: %w", attachment.GetFileName(), err)
			}
		}

		// Close the writer
		err = writer.Close()
		if err != nil {
			return nil, fmt.Errorf("Failed to close writer")
		}

		if ctx != nil {
			request, err = http.NewRequestWithContext(ctx, method, url, &requestBody)
		} else {
			request, err = http.NewRequest(method, url, &requestBody)
		}
		if err != nil {
			return nil, fmt.Errorf("error creating HTTP request: %w", err)
		}

		request.Header.Set("Content-Type", writer.FormDataContentType())

		return request, nil
	}

	// JSON parse instead
	if ctx != nil {
		request, err = http.NewRequestWithContext(ctx, method, url, bytes.NewReader(body))
	} else {
		request, err = http.NewRequest(method, url, bytes.NewReader(body))
	}
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %w", err)
	}

	request.Header.Set("Content-Type", "application/json")

	return request, nil
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
