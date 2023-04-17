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

func (guildClient *GuildClient) GetMemberClient(memberId discord.Snowflake) *MemberClient {
	return &MemberClient{
		MemberId:    memberId,
		GuildClient: guildClient,
	}
}
