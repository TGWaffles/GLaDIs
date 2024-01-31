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

func NewAuthorizedUser(refreshToken string, accessToken string, expiresIn int, scopes []oauth_scopes.OAuthScope) *AuthorizedUser {
	return &AuthorizedUser{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
		ExpiresIn:    expiresIn,
		Scopes:       scopes,
		Client:       http.DefaultClient,
	}
}

func (authedUser *AuthorizedUser) MakeTokenRequest(discordRequest DiscordRequest) (response *http.Response, err error) {
	discordRequest.ValidateEndpoint()

	return authedUser.OAuthClient.MakeRequest(discordRequest)
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

	requestBody := &TokenRequest{
		GrantType:    GrantTypeRefreshToken,
		RefreshToken: authedUser.RefreshToken,
	}

	var tokenResponse TokenResponse

	_, err := authedUser.MakeTokenRequest(DiscordRequest{
		ExpectedStatus: 200,
		Method:         "POST",
		Endpoint:       "/oauth2/token",
		Body:           []byte(requestBody.ToString()),
		UnmarshalTo:    tokenResponse,
	})

	if err != nil {
		return err
	}

	authedUser.RefreshToken = tokenResponse.RefreshToken
	authedUser.AccessToken = tokenResponse.AccessToken
	authedUser.ExpiresIn = tokenResponse.ExpiresIn

	return nil
}

func (authedUser *AuthorizedUser) RevokeTokens() error {
	requestBody := &RevokeTokenRequest{
		Token:         authedUser.AccessToken,
		TokenTypeHint: "access_token",
	}

	_, err := authedUser.MakeTokenRequest(DiscordRequest{
		ExpectedStatus: 204,
		Method:         "POST",
		Endpoint:       "/oauth2/token/revoke",
		Body:           []byte(requestBody.ToString()),
		UnmarshalTo:    nil,
	})
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
