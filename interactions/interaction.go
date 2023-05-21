package interactions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tgwaffles/gladis/commands"
	"github.com/tgwaffles/gladis/components"
	"github.com/tgwaffles/gladis/discord"
)

type Interaction struct {
	Id             discord.Snowflake    `json:"id"`
	ApplicationId  discord.Snowflake    `json:"application_id"`
	Type           InteractionType      `json:"type"`
	DataInternal   *json.RawMessage     `json:"data,omitempty"`
	Data           InteractionData      `json:"-"`
	GuildId        *discord.Snowflake   `json:"guild_id,omitempty"`
	ChannelId      *discord.Snowflake   `json:"channel_id,omitempty"`
	Member         *discord.Member      `json:"member,omitempty"`
	User           *discord.User        `json:"user,omitempty"`
	Token          string               `json:"token"`
	Version        int                  `json:"version"`
	Message        *discord.Message     `json:"message,omitempty"`
	AppPermissions *discord.Permissions `json:"permissions,omitempty"`
	Locale         string               `json:"locale"`
	GuildLocale    string               `json:"guild_locale"`
	hook           *Webhook             `json:"-"` // Used for responding to the interaction
}

const (
	apiUrl                       = "https://discord.com/api"
	createInteractionResponseUrl = apiUrl + "/interactions/%d/%s/callback"
)

var (
	httpClient = http.Client{}
)

func Parse(data string) (interaction *Interaction, err error) {
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
	case PingInteractionType:
		// PingInteractionType has no data
		return
	case ApplicationCommandInteractionType, ApplicationCommandAutocompleteInteractionType:
		appCommandData := commands.ApplicationCommandData{}
		err = json.Unmarshal(*interaction.DataInternal, &appCommandData)
		interaction.Data = &appCommandData
		break
	case MessageComponentInteractionType:
		messageComponentData := components.MessageComponentData{}
		err = json.Unmarshal(*interaction.DataInternal, &messageComponentData)
		interaction.Data = &messageComponentData
		break
	case ModalSubmitInteractionType:
		modalSubmitData := components.ModalSubmitData{}
		err = json.Unmarshal(*interaction.DataInternal, &modalSubmitData)
		interaction.Data = &modalSubmitData
		break
	}
	return err
}

func (interaction *Interaction) IsPing() bool {
	return interaction.Type == PingInteractionType
}

func (interaction *Interaction) CreateResponse(response InteractionResponse) error {
	data, err := json.Marshal(response)
	if err != nil {
		return fmt.Errorf("error marshaling data to JSON: %w", err)
	}
	request, err := http.NewRequest("POST", fmt.Sprintf(createInteractionResponseUrl, interaction.Id, interaction.Token), bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("error creating HTTP request: %w", err)
	}

	request.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("error sending HTTP request: %w", err)
	}

	if resp.StatusCode != 204 {
		return fmt.Errorf("expected status code 204, got %d", resp.StatusCode)
	}

	return nil
}

func (interaction *Interaction) GetResponse() (*discord.Message, error) {
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
	var flags *int
	if isEphemeral {
		*flags = 64
	}

	response := InteractionResponse{
		Type: DeferredChannelMessageWithSourceInteractionCallbackType,
		Data: &MessageCallbackData{
			Flags: flags,
		},
	}

	return interaction.CreateResponse(response)
}
