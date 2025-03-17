package client

import (
	"encoding/json"
	"net/http"

	"github.com/tgwaffles/gladis/discord"
	"github.com/tgwaffles/gladis/discord/interaction_context_type"
)

type ApplicationClient struct {
	ApplicationId discord.Snowflake
	Bot           *BotClient
}

type CreateApplicationCommand struct {
	Name                     string                                            `json:"name"`
	NameLocalizations        map[string]string                                 `json:"name_localizations,omitempty"`
	Description              *string                                           `json:"description,omitempty"`
	DescriptionLocalizations map[string]string                                 `json:"description_localizations,omitempty"`
	Options                  []discord.ApplicationCommandOption                `json:"options,omitempty"`
	DefaultMemberPermissions *string                                           `json:"default_member_permissions,omitempty"`
	DMPermission             *bool                                             `json:"dm_permission,omitempty"`
	DefaultPermission        *bool                                             `json:"default_permission,omitempty"`
	IntegrationTypes         []string                                          `json:"integration_types,omitempty"`
	Contexts                 []interaction_context_type.InteractionContextType `json:"contexts,omitempty"`
	Type                     *int                                              `json:"type,omitempty"`
	NSFW                     *bool                                             `json:"nsfw,omitempty"`
}

func (appClient *ApplicationClient) MakeRequest(discordRequest DiscordRequest) (response *http.Response, err error) {
	discordRequest.ValidateEndpoint()
	discordRequest.Endpoint = "/applications/" + appClient.ApplicationId.String() + discordRequest.Endpoint

	return appClient.Bot.MakeRequest(discordRequest)
}

func (appClient *ApplicationClient) RegisterCommands(cmds []CreateApplicationCommand) error {
	body, err := json.Marshal(cmds)

	if err != nil {
		return err
	}

	_, err = appClient.MakeRequest(DiscordRequest{
		Method:         "PUT",
		Endpoint:       "/commands",
		Body:           body,
		ExpectedStatus: 200,
	})

	if err != nil {
		return err
	}

	return nil
}

func (appClient *ApplicationClient) RegisterCommand(cmds CreateApplicationCommand) error {
	body, err := json.Marshal(cmds)

	if err != nil {
		return err
	}

	_, err = appClient.MakeRequest(DiscordRequest{
		Method:         "POST",
		Endpoint:       "/commands",
		Body:           body,
		ExpectedStatus: 201,
	})

	if err != nil {
		return err
	}

	return nil
}
