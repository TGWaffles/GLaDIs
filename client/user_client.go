package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tgwaffles/gladis/discord"
)

type UserClient struct {
	UserId discord.Snowflake
	Bot    *BotClient
}

func (userClient *UserClient) MakeRequest(discordRequest DiscordRequest) (response *http.Response, err error) {
	discordRequest.ValidateEndpoint()
	discordRequest.Endpoint = "/users/" + userClient.GetIdAsString() + discordRequest.Endpoint

	return userClient.Bot.MakeRequest(discordRequest)
}

func (userClient *UserClient) GetIdAsString() string {
	if userClient.UserId == 0 {
		return "@me"
	}
	return userClient.UserId.String()
}

func (userClient *UserClient) CreateDMChannel() (*discord.Channel, error) {
	channel := &discord.Channel{}
	selfUserClient := userClient.Bot.GetSelfUserClient()
	jsonBody := map[string]interface{}{
		"recipient_id": userClient.UserId,
	}

	body, err := json.Marshal(jsonBody)

	if err != nil {
		return nil, fmt.Errorf("error marshalling user id to JSON body: %w", err)
	}

	_, err = selfUserClient.MakeRequest(DiscordRequest{
		Method:         "POST",
		Endpoint:       "/channels",
		Body:           body,
		ExpectedStatus: 200,
		UnmarshalTo:    channel,
	})
	if err != nil {
		return nil, err
	}

	return channel, nil
}

func (userClient *UserClient) SendMessage(messageData SendMessageData) (*discord.Message, error) {
	dmChannel, err := userClient.CreateDMChannel()
	if err != nil {
		return nil, err
	}

	channelClient := userClient.Bot.GetChannelClient(dmChannel.Id)
	return channelClient.SendMessage(messageData)
}
