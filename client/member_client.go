package client

import (
	"github.com/tgwaffles/gladis/discord"
	"net/http"
)

type MemberClient struct {
	MemberId    discord.Snowflake
	GuildClient *GuildClient
}

func (memberClient *MemberClient) MakeRequest(discordRequest DiscordRequest) (response *http.Response, err error) {
	discordRequest.ValidateEndpoint()
	discordRequest.Endpoint = "/members/" + memberClient.MemberId.String() + discordRequest.Endpoint

	return memberClient.GuildClient.MakeRequest(discordRequest)
}

func (memberClient *MemberClient) FetchMember() (*discord.Member, error) {
	member := &discord.Member{}
	_, err := memberClient.MakeRequest(DiscordRequest{
		Method:         "GET",
		Endpoint:       "",
		Body:           nil,
		ExpectedStatus: 200,
		UnmarshalTo:    member,
	})

	if err != nil {
		return nil, err
	}

	return member, nil
}

type ModifyMemberRoleOpts struct {
	RoleID discord.Snowflake

	Reason string
}

func (memberClient *MemberClient) AddRoleToMember(opts ModifyMemberRoleOpts) error {
	additionalHeaders := map[string]string{}
	if len(opts.Reason) > 0 {
		additionalHeaders["X-Audit-Log-Reason"] = opts.Reason
	}
	req := DiscordRequest{
		Method:            "PUT",
		Endpoint:          "/roles/" + opts.RoleID.String(),
		Body:              nil,
		ExpectedStatus:    204,
		UnmarshalTo:       nil,
		AdditionalHeaders: additionalHeaders,
	}
	_, err := memberClient.MakeRequest(req)
	if err != nil {
		return err
	}
	return nil
}
