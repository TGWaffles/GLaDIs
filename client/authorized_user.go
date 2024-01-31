package client

import (
	"encoding/json"
	"fmt"
	"github.com/tgwaffles/gladis/client/errors"
	"github.com/tgwaffles/gladis/discord"
	"net/http"
)

type AuthorizedUser struct {
	RefreshToken string
	AccessToken  string
	ExpiresIn    int
	Scopes       []discord.OAuthScope
	Client       *http.Client
	OAuthClient  *OAuthClient
}

func NewAuthorizedUser(refreshToken string, accessToken string, expiresIn int, scopes []discord.OAuthScope) *AuthorizedUser {
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
	request, err := http.NewRequest(discordRequest.Method, discordRequest.GetUrl(), discordRequest.getBodyAsReader())
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %w", err)
	}

	if !discordRequest.DisableAuth {
		request.Header.Set("Authorization", "Bearer "+authedUser.AccessToken)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", getUserAgent())
	request.Header.Set("Accept", "application/json")

	for key, value := range discordRequest.AdditionalHeaders {
		request.Header.Set(key, value)
	}

	response, err = authedUser.Client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error sending HTTP request: %w", err)
	}

	if !discordRequest.DisableStatusCheck && response.StatusCode != discordRequest.ExpectedStatus {
		return nil, errors.StatusError{
			Code:     errors.StatusErrorCode(response.StatusCode),
			Response: response,
		}
	}

	if discordRequest.UnmarshalTo != nil {
		err = json.NewDecoder(response.Body).Decode(discordRequest.UnmarshalTo)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling response: %w", err)
		}
		return nil, nil
	}

	return response, nil
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
