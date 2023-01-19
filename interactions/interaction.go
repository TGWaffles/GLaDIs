package interactions

import (
	"encoding/json"
	"github.com/tgwaffles/gladis/commands"
	"github.com/tgwaffles/gladis/components"
	"github.com/tgwaffles/gladis/discord"
)

type Interaction struct {
	Id            discord.Snowflake  `json:"id"`
	ApplicationId discord.Snowflake  `json:"application_id"`
	Type          InteractionType    `json:"type"`
	DataInternal  *json.RawMessage   `json:"data,omitempty"`
	Data          InteractionData    `json:"-"`
	GuildId       *discord.Snowflake `json:"guild_id,omitempty"`
	ChannelId     *discord.Snowflake `json:"channel_id,omitempty"`
	Member        *discord.Member    `json:"member,omitempty"`
	User          *discord.User      `json:"user,omitempty"`
	Token         string             `json:"token"`
	Version       int                `json:"version"`
}

func Parse(data string) (interaction *Interaction, err error) {
	err = json.Unmarshal([]byte(data), &interaction)
	if err == nil {
		err = interaction.createData()
	}
	return interaction, err
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
