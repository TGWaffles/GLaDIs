package client

import (
	"github.com/tgwaffles/gladis/discord"
	"net/http"
)

type GuildClient struct {
	GuildId discord.Snowflake
	Bot     *BotClient
}

func (guildClient *GuildClient) MakeRequest(discordRequest DiscordRequest) (response *http.Response, err error) {
	discordRequest.ValidateEndpoint()
	discordRequest.Endpoint = "/guilds/" + guildClient.GuildId.String() + discordRequest.Endpoint

	return guildClient.Bot.MakeRequest(discordRequest)
}

func (guildClient *GuildClient) FetchGuild() (*discord.Guild, error) {
	guild := &discord.Guild{}
	_, err := guildClient.MakeRequest(DiscordRequest{
		Method:         "GET",
		Endpoint:       "",
		Body:           nil,
		ExpectedStatus: 200,
		UnmarshalTo:    guild,
	})

	if err != nil {
		return nil, err
	}

	return guild, nil
}

func (guildClient *GuildClient) GetMemberClient(memberId discord.Snowflake) *MemberClient {
	return &MemberClient{
		MemberId:    memberId,
		GuildClient: guildClient,
	}
}

type ActiveThreadsResponse struct {
	Threads       []discord.Channel      `json:"threads"`
	ThreadMembers []discord.ThreadMember `json:"members"`
}

func (guildClient *GuildClient) GetActiveThreads() (ActiveThreadsResponse, error) {
	response := ActiveThreadsResponse{}
	_, err := guildClient.MakeRequest(DiscordRequest{
		Method:         "GET",
		Endpoint:       "/threads/active",
		Body:           nil,
		ExpectedStatus: 200,
		UnmarshalTo:    &response,
	})

	if err != nil {
		return ActiveThreadsResponse{}, err
	}

	return response, nil
}
