package interactable

import (
	"fmt"

	"github.com/JackHumphries9/dapper-go/discord"
	"github.com/JackHumphries9/dapper-go/discord/command_option_type"
	"github.com/JackHumphries9/dapper-go/discord/interaction_callback_type"
	"github.com/JackHumphries9/dapper-go/discord/interaction_type"
	"github.com/JackHumphries9/dapper-go/discord/message_flags"
	"github.com/JackHumphries9/dapper-go/helpers"
)

type InteractionContext struct {
	Interaction  *discord.Interaction
	DeferChannel chan *discord.InteractionResponse
	hasDeferred  bool
	messageFlags message_flags.MessageFlags
}

func (ic *InteractionContext) SetEphemeral(ep bool) {
	if ep {
		ic.messageFlags.AddFlag(message_flags.Ephemeral)
	} else {
		ic.messageFlags.RemoveFlag(message_flags.Ephemeral)
	}
}

func (ic *InteractionContext) HasDeferred() bool {
	return ic.hasDeferred
}

func (ic *InteractionContext) Defer() {
	if ic.hasDeferred {
		fmt.Printf("Interaction already deferred")
		return
	}

	ic.hasDeferred = true

	if ic.Interaction.Type == interaction_type.ApplicationCommand {
		ic.DeferChannel <- &discord.InteractionResponse{
			Type: interaction_callback_type.DeferredChannelMessageWithSource,
			Data: &discord.MessageCallbackData{
				Flags: helpers.Ptr(int(ic.messageFlags)),
			},
		}
	}

	if ic.Interaction.Type == interaction_type.MessageComponent {
		ic.DeferChannel <- &discord.InteractionResponse{
			Type: interaction_callback_type.DeferredUpdateMessage,
			Data: &discord.MessageCallbackData{
				Flags: helpers.Ptr(int(ic.messageFlags)),
			},
		}
	}

	if ic.Interaction.Type == interaction_type.ModalSubmit {
		ic.DeferChannel <- &discord.InteractionResponse{
			Type: interaction_callback_type.DeferredChannelMessageWithSource,
			Data: &discord.MessageCallbackData{
				Flags: helpers.Ptr(int(ic.messageFlags)),
			},
		}
	}
}

func (ic *InteractionContext) Respond(msg discord.ResponseEditData) error {
	if ic.hasDeferred {
		return ic.Interaction.EditResponse(msg)
	}

	var responseType interaction_callback_type.InteractionCallbackType

	if ic.Interaction.Type == interaction_type.ApplicationCommand {
		responseType = interaction_callback_type.ChannelMessageWithSource
	} else if ic.Interaction.Type == interaction_type.MessageComponent {
		responseType = interaction_callback_type.UpdateMessage
	} else if ic.Interaction.Type == interaction_type.ModalSubmit {
		responseType = interaction_callback_type.ChannelMessageWithSource
	}

	ic.DeferChannel <- &discord.InteractionResponse{
		Type: responseType,
		Data: &discord.MessageCallbackData{
			Content:         msg.Content,
			Flags:           helpers.Ptr(int(ic.messageFlags)),
			Embeds:          msg.Embeds,
			Components:      msg.Components,
			AllowedMentions: msg.AllowedMentions,
		},
	}

	return nil
}

func (ic *InteractionContext) ShowModal(modal Modal) error {
	if ic.hasDeferred {
		return fmt.Errorf("Cannot show modal after deferring")
	}

	ic.DeferChannel <- &discord.InteractionResponse{
		Type: interaction_callback_type.Modal,
		Data: modal.Modal,
	}
	return nil
}

func (ic *InteractionContext) GetModalTextInputValue(id string) *string {
	if ic.Interaction.Type != interaction_type.ModalSubmit {
		return nil
	}

	submitData := ic.Interaction.Data.(*discord.ModalSubmitData)

	for _, component := range submitData.Components {
		textInput, ok := component.(*discord.ActionRow).Components[0].(*discord.TextInput)

		if !ok {
			return nil
		}

		if textInput.CustomId == id {
			return textInput.Value
		}
	}

	return nil
}

func (ic *InteractionContext) GetSelectValues() (*[]string, error) {
	if ic.Interaction.Type != interaction_type.MessageComponent {
		return nil, fmt.Errorf("cannot get command options from a non command interaction")
	}

	selMenu, ok := ic.Interaction.Data.(*discord.MessageComponentData)

	if !ok {
		return nil, fmt.Errorf("cannot convert to values")
	}

	return &selMenu.Values, nil

}

func GetCommandOption(itx *discord.Interaction, name string) (*discord.ApplicationCommandDataOption, error) {
	if itx.Type != interaction_type.ApplicationCommand {
		return nil, fmt.Errorf("cannot get command options from a non command interaction")
	}

	commandData := itx.Data.(*discord.ApplicationCommandData)

	return commandData.GetOption(name), nil
}

func (ic *InteractionContext) GetStringCommandOption(name string) (*string, error) {
	option, err := GetCommandOption(ic.Interaction, name)

	if err != nil {
		return nil, err
	}

	if option.Type == command_option_type.String {
		return helpers.Ptr(option.Value.(string)), nil
	}

	return nil, fmt.Errorf("Cannot find string option: %s", name)
}

func (ic *InteractionContext) GetBoolCommandOption(name string) (*bool, error) {
	option, err := GetCommandOption(ic.Interaction, name)

	if err != nil {
		return nil, err
	}

	if option.Type == command_option_type.Boolean {
		return helpers.Ptr(option.Value.(bool)), nil
	}

	return nil, fmt.Errorf("Cannot find string option: %s", name)
}

func (ic *InteractionContext) GetNumberCommandOption(name string) (*float64, error) {
	option, err := GetCommandOption(ic.Interaction, name)

	if err != nil {
		return nil, err
	}

	if option.Type == command_option_type.Number {
		return helpers.Ptr(option.Value.(float64)), nil
	}

	return nil, fmt.Errorf("Cannot find string option: %s", name)
}

func (ic *InteractionContext) GetIntCommandOption(name string) (*int64, error) {
	option, err := GetCommandOption(ic.Interaction, name)

	if err != nil {
		return nil, err
	}

	if option.Type == command_option_type.Integer {
		return helpers.Ptr(option.Value.(int64)), nil
	}

	return nil, fmt.Errorf("Cannot find string option: %s", name)
}

func (ic *InteractionContext) GetUserCommandOption(name string) (*discord.Snowflake, error) {
	option, err := GetCommandOption(ic.Interaction, name)

	if err != nil {
		return nil, err
	}

	if option.Type == command_option_type.User {
		val, err := discord.GetSnowflake(option.Value)

		if err != nil {
			return nil, fmt.Errorf("failed to get snowflake")
		}

		return &val, nil
	}

	return nil, fmt.Errorf("Cannot find string option: %s", name)
}

func (ic *InteractionContext) GetRoleCommandOption(name string) (*discord.Snowflake, error) {
	option, err := GetCommandOption(ic.Interaction, name)

	if err != nil {
		return nil, err
	}

	if option.Type == command_option_type.Role {
		val, err := discord.GetSnowflake(option.Value)

		if err != nil {
			return nil, fmt.Errorf("failed to get snowflake")
		}

		return &val, nil
	}

	return nil, fmt.Errorf("Cannot find string option: %s", name)
}

func (ic *InteractionContext) GetMentionableCommandOption(name string) (*discord.Snowflake, error) {
	option, err := GetCommandOption(ic.Interaction, name)

	if err != nil {
		return nil, err
	}

	if option.Type == command_option_type.Mentionable {
		val, err := discord.GetSnowflake(option.Value)

		if err != nil {
			return nil, fmt.Errorf("failed to get snowflake")
		}

		return &val, nil
	}

	return nil, fmt.Errorf("Cannot find string option: %s", name)
}

func (ic *InteractionContext) GetChannelCommandOption(name string) (*discord.Snowflake, error) {
	option, err := GetCommandOption(ic.Interaction, name)

	if err != nil {
		return nil, err
	}

	if option.Type == command_option_type.Channel {
		val, err := discord.GetSnowflake(option.Value)

		if err != nil {
			return nil, fmt.Errorf("failed to get snowflake")
		}

		return &val, nil
	}

	return nil, fmt.Errorf("Cannot find string option: %s", name)
}

func (ic *InteractionContext) GetAttachmentCommandOption(name string) (*discord.Attachment, error) {
	option, err := GetCommandOption(ic.Interaction, name)

	if err != nil {
		return nil, err
	}

	if option.Type == command_option_type.Mentionable {
		return helpers.Ptr(option.Value.(discord.Attachment)), nil
	}

	return nil, fmt.Errorf("Cannot find string option: %s", name)
}

// TODO: Add subcommands
