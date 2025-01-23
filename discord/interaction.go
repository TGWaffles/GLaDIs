package discord

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/JackHumphries9/dapper-go/discord/interaction_callback_type"
	"github.com/JackHumphries9/dapper-go/discord/interaction_context_type"
	"github.com/JackHumphries9/dapper-go/discord/interaction_type"
)

type Interaction struct {
	Id             Snowflake                                        `json:"id"`
	ApplicationId  Snowflake                                        `json:"application_id"`
	Type           interaction_type.InteractionType                 `json:"type"`
	DataInternal   *json.RawMessage                                 `json:"data,omitempty"`
	Data           InteractionData                                  `json:"-"`
	GuildId        *Snowflake                                       `json:"guild_id,omitempty"`
	Channel        *Channel                                         `json:"channel,omitempty"`
	ChannelId      *Snowflake                                       `json:"channel_id,omitempty"`
	Member         *Member                                          `json:"member,omitempty"`
	User           *User                                            `json:"user,omitempty"`
	Token          string                                           `json:"token"`
	Version        int                                              `json:"version"`
	Message        *Message                                         `json:"message,omitempty"`
	AppPermissions *Permissions                                     `json:"permissions,omitempty"`
	Locale         string                                           `json:"locale"`
	GuildLocale    string                                           `json:"guild_locale"`
	hook           *Webhook                                         // Used for responding to the interaction
	Context        *interaction_context_type.InteractionContextType `json:"context,omitempty"`
}

type InteractionData interface {
}

const (
	apiUrl                       = "https://discord.com/api"
	createInteractionResponseUrl = apiUrl + "/interactions/%d/%s/callback"
)

var (
	httpClient = http.Client{}
)

func SetHttpClient(newClient http.Client) {
	httpClient = newClient
}

func ParseInteraction(data string) (interaction *Interaction, err error) {
	err = json.Unmarshal([]byte(data), &interaction)
	return interaction, err
}

func (interaction *Interaction) UnmarshalJSON(d []byte) error {
	type InnerInteraction Interaction

	var inner InnerInteraction

	if err := json.Unmarshal(d, &inner); err != nil {
		return err
	}

	castInteraction := Interaction(inner)

	err := castInteraction.createData()
	if err != nil {
		return err
	}

	*interaction = castInteraction

	return nil
}

func (interaction *Interaction) createData() (err error) {
	if interaction.Data != nil || interaction.DataInternal == nil {
		return
	}

	switch interaction.Type {
	case interaction_type.Ping:
		// Ping has no data
		return
	case interaction_type.ApplicationCommand, interaction_type.ApplicationCommandAutocomplete:
		appCommandData := ApplicationCommandData{}
		err = json.Unmarshal(*interaction.DataInternal, &appCommandData)
		interaction.Data = &appCommandData
		break
	case interaction_type.MessageComponent:
		messageComponentData := MessageComponentData{}
		err = json.Unmarshal(*interaction.DataInternal, &messageComponentData)
		interaction.Data = &messageComponentData
		break
	case interaction_type.ModalSubmit:
		modalSubmitData := ModalSubmitData{}
		err = json.Unmarshal(*interaction.DataInternal, &modalSubmitData)
		interaction.Data = &modalSubmitData
		break
	}
	return err
}

func (interaction *Interaction) IsPing() bool {
	return interaction.Type == interaction_type.Ping
}

func (interaction *Interaction) CreateResponse(response InteractionResponse) error {
	return interaction.CreateResponseWithContext(nil, response)
}

func (interaction *Interaction) CreateResponseWithContext(ctx context.Context, response InteractionResponse) error {
	data, err := json.Marshal(response)
	if err != nil {
		return fmt.Errorf("error marshaling data to JSON: %w", err)
	}
	var request *http.Request
	if ctx != nil {
		request, err = http.NewRequestWithContext(ctx, "POST", fmt.Sprintf(createInteractionResponseUrl, interaction.Id, interaction.Token), bytes.NewReader(data))
	} else {
		request, err = http.NewRequest("POST", fmt.Sprintf(createInteractionResponseUrl, interaction.Id, interaction.Token), bytes.NewReader(data))
	}
	if err != nil {
		return fmt.Errorf("error creating HTTP request: %w", err)
	}

	request.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("error sending HTTP request: %w", err)
	}

	if resp.StatusCode != 204 {
		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("expected status code 204, got %d", resp.StatusCode)
		}
		return fmt.Errorf(
			"error sending interaction response, status code %d (expected 204)\nresponse body: %s\nrequest body: %s",
			resp.StatusCode, string(responseBody), string(data))
	}

	return nil
}

func (interaction *Interaction) GetWebhook() *Webhook {
	if interaction.hook == nil {
		interaction.hook = &Webhook{
			Id:    interaction.ApplicationId,
			Token: &interaction.Token,
		}
	}

	return interaction.hook
}

func (interaction *Interaction) GetResponse() (*Message, error) {
	if interaction.hook == nil {
		interaction.hook = &Webhook{
			Id:    interaction.ApplicationId,
			Token: &interaction.Token,
		}
	}

	return interaction.hook.GetMessage(WebhookGetMessageRequest{MessageId: "@original"})
}

func (interaction *Interaction) EditResponse(data ResponseEditData) error {
	if interaction.hook == nil {
		interaction.hook = &Webhook{
			Id:    interaction.ApplicationId,
			Token: &interaction.Token,
		}
	}

	return interaction.hook.EditMessage("@original", data)
}

func (interaction *Interaction) DeleteResponse() error {
	if interaction.hook == nil {
		interaction.hook = &Webhook{
			Id:    interaction.ApplicationId,
			Token: &interaction.Token,
		}
	}

	return interaction.hook.DeleteMessage("@original")
}

func (interaction *Interaction) DeferResponse(isEphemeral bool) error {
	var flags int
	if isEphemeral {
		flags = 64
	}

	response := InteractionResponse{
		Type: interaction_callback_type.DeferredChannelMessageWithSource,
		Data: &MessageCallbackData{
			Flags: &flags,
		},
	}

	return interaction.CreateResponse(response)
}
