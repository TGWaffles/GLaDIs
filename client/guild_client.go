package client

import (
	"encoding/json"
	"fmt"
	"github.com/tgwaffles/gladis/discord"
	"net/http"
	"net/url"
	"strconv"
)

type GuildClient struct {
	GuildId discord.Snowflake
	Bot     *BotClient
}

type JoinUserToGuildRequest struct {
	AccessToken string              `json:"access_token"`
	Nick        string              `json:"nick,omitempty"`
	Roles       []discord.Snowflake `json:"roles,omitempty"`
	Mute        bool                `json:"mute,omitempty"`
	Deaf        bool                `json:"deaf,omitempty"`
}

func (guildClient *GuildClient) MakeRequest(discordRequest DiscordRequest) (response *http.Response, err error) {
	discordRequest.ValidateEndpoint()
	discordRequest.Endpoint = "/guilds/" + guildClient.GuildId.String() + discordRequest.Endpoint

	return guildClient.Bot.MakeRequest(discordRequest)
}

func (guildClient *GuildClient) FetchGuild(withCounts bool) (*discord.Guild, error) {
	guild := &discord.Guild{}
	endpoint := ""
	if withCounts {
		endpoint = "?with_counts=true"
	}
	_, err := guildClient.MakeRequest(DiscordRequest{
		Method:         "GET",
		Endpoint:       endpoint,
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

type ListMembersRequest struct {
	// The last member fetched
	After *discord.Snowflake
	// Max members to fetch in one request
	Limit *int
}

func (guildClient *GuildClient) ListMembers(request ListMembersRequest) ([]discord.Member, error) {
	members := make([]discord.Member, 0)

	query := make(url.Values)
	if request.After != nil {
		query.Add("after", request.After.String())
	}
	if request.Limit != nil {
		query.Add("limit", strconv.Itoa(*request.Limit))
	}
	endpoint := "/members"
	encodedQuery := query.Encode()
	if len(encodedQuery) > 0 {
		endpoint += "?" + encodedQuery
	}

	req := DiscordRequest{
		ExpectedStatus: 200,
		Method:         "GET",
		Endpoint:       endpoint,
		Body:           nil,
		UnmarshalTo:    &members,
	}

	_, err := guildClient.MakeRequest(req)
	if err != nil {
		return nil, err
	}

	return members, nil
}

func (guildClient *GuildClient) GetChannels() ([]discord.Channel, error) {
	channels := make([]discord.Channel, 0)
	_, err := guildClient.MakeRequest(DiscordRequest{
		Method:         "GET",
		Endpoint:       "/channels",
		Body:           nil,
		ExpectedStatus: 200,
		UnmarshalTo:    &channels,
	})

	if err != nil {
		return nil, err
	}

	return channels, nil
}

func (guildClient *GuildClient) JoinUserToGuild(authedUser *AuthorizedUser, userId discord.Snowflake, guildId discord.Snowflake) (err error) {
	requestBody := &JoinUserToGuildRequest{
		AccessToken: authedUser.AccessToken,
	}

	data, err := json.Marshal(requestBody)

	if err != nil {
		return err
	}

	var response *http.Response
	response, err = guildClient.Bot.MakeRequest(DiscordRequest{
		Method:             "PUT",
		Endpoint:           "/guilds/" + guildId.String() + "/members/" + userId.String(),
		Body:               data,
		DisableStatusCheck: true,
	})

	if err != nil {
		return err
	}

	if response.StatusCode == 204 {
		return fmt.Errorf("user already exists in guild")
	}

	return nil
}
