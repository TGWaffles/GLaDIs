package interactions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tgwaffles/gladis/commands"
	"github.com/tgwaffles/gladis/components"
	"github.com/tgwaffles/gladis/discord"
	"io"
	"io/ioutil"
	"net/http"
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
}

const (
	apiUrl                       = "https://discord.com/api"
	createInteractionResponseUrl = apiUrl + "/interactions/%d/%s/callback"
	hookUrl                      = apiUrl + "/webhooks/%d/%s/messages/@original"
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
		return err
	}
	request, err := http.NewRequest("POST", fmt.Sprintf(createInteractionResponseUrl, interaction.ApplicationId, interaction.Token), bytes.NewReader(data))
	request.Header.Set("Content-Type", "application/json")

	if err != nil {
		return err
	}

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}

	if resp.StatusCode != 204 {
		return fmt.Errorf("expected status code 204, got %d", resp.StatusCode)
	}

	return nil
}

func (interaction *Interaction) GetResponse() (*discord.Message, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf(hookUrl, interaction.ApplicationId, interaction.Token), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("expected status code 200, got %d", resp.StatusCode)
	}

	var message discord.Message
	err = json.NewDecoder(resp.Body).Decode(&message)
	if err != nil {
		return nil, err
	}

	return &message, nil
}

func (interaction *Interaction) EditResponse(data ResponseEditData) error {
	body, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error marshaling data to JSON: %s", err.Error())
	}

	request, err := http.NewRequest("PATCH", fmt.Sprintf(hookUrl, interaction.ApplicationId, interaction.Token), bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("error creating HTTP request: %s", err.Error())
	}

	request.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("error sending HTTP request: %s", err.Error())
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading HTTP response body: %s", err.Error())
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("expected status code 200, got %d. Response body: %s", resp.StatusCode, string(responseBody))
	}

	return nil
}

func (interaction *Interaction) DeleteResponse() error {
	request, err := http.NewRequest("DELETE", fmt.Sprintf(hookUrl, interaction.ApplicationId, interaction.Token), nil)
	if err != nil {
		return fmt.Errorf("error creating HTTP request: %s", err.Error())
	}
	request.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("error sending HTTP request: %s", err.Error())
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading HTTP response body: %s", err.Error())
	}

	if resp.StatusCode != 204 {
		return fmt.Errorf("expected status code 204, got %d. Response body: %s", resp.StatusCode, string(responseBody))
	}

	return nil
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
