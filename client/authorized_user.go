package client

import (
	"github.com/tgwaffles/gladis/discord"
	"github.com/tgwaffles/gladis/discord/oauth_scopes"
	"net/http"
)

type AuthorizedUser struct {
	RefreshToken string
	AccessToken  string
	ExpiresIn    int
	Scopes       []oauth_scopes.OAuthScope
	Client       *http.Client
	OAuthClient  *OAuthClient
}

func NewAuthorizedUser(oauthClient *OAuthClient, refreshToken string, accessToken string, expiresIn int, scopes []oauth_scopes.OAuthScope) *AuthorizedUser {
	return &AuthorizedUser{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
		ExpiresIn:    expiresIn,
		Scopes:       scopes,
		Client:       http.DefaultClient,
		OAuthClient:  oauthClient,
	}
}

func (authedUser *AuthorizedUser) MakeRequest(discordRequest DiscordRequest) (response *http.Response, err error) {
	discordRequest.ValidateEndpoint()

	if !discordRequest.DisableAuth {
		discordRequest.AdditionalHeaders["Authorization"] = "Bearer " + authedUser.AccessToken
	}

	discordRequest.AdditionalHeaders["Content-Type"] = "application/json"

	return authedUser.OAuthClient.MakeRequest(discordRequest)
}

func (authedUser *AuthorizedUser) RefreshTokens() error {

	tokenResponse, err := authedUser.OAuthClient.RefreshTokensForUser(authedUser.RefreshToken)

	if err != nil {
		return err
	}

	authedUser.RefreshToken = tokenResponse.RefreshToken
	authedUser.AccessToken = tokenResponse.AccessToken
	authedUser.ExpiresIn = tokenResponse.ExpiresIn

	return nil
}

func (authedUser *AuthorizedUser) RevokeTokens() error {
	err := authedUser.OAuthClient.RevokeTokensForUser(authedUser.AccessToken)

	if err != nil {
		return err
	}

	return nil
}

func (authedUser *AuthorizedUser) FetchUser() (*discord.User, error) {
	user := &discord.User{}

	_, err := authedUser.MakeRequest(DiscordRequest{
		Method:         "GET",
		Endpoint:       "/users/@me",
		ExpectedStatus: 200,
		UnmarshalTo:    user,
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (authedUser *AuthorizedUser) FetchGuilds() ([]discord.Guild, error) {
	var guilds []discord.Guild

	_, err := authedUser.MakeRequest(DiscordRequest{
		Method:         "GET",
		Endpoint:       "/users/@me/guilds",
		ExpectedStatus: 200,
		UnmarshalTo:    guilds,
	})
	if err != nil {
		return nil, err
	}

	return guilds, nil
}
