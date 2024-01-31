package client

import (
	"fmt"
	"github.com/tgwaffles/gladis/client/errors"
	"github.com/tgwaffles/gladis/discord"
	"net/http"
	"net/url"
)

const DISCORD_AUTHORIZATION_URL = "https://discordapp.com/api/oauth2/authorize"
const DISCORD_TOKEN_URL = "https://discordapp.com/api/oauth2/token"

type OAuthClient struct {
	clientId     string
	clientSecret string
	client       *http.Client
	redirectUri  string
	Scopes       []discord.OAuthScope
}

func (oauthClient *OAuthClient) MakeRequest(discordRequest DiscordRequest) (response *http.Response, err error) {
	discordRequest.ValidateEndpoint()

	request, err := http.NewRequest(discordRequest.Method, discordRequest.GetUrl(), discordRequest.getBodyAsReader())

	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %w", err)
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("User-Agent", getUserAgent())
	request.Header.Set("Accept", "application/json")

	for key, value := range discordRequest.AdditionalHeaders {
		request.Header.Set(key, value)
	}

	response, err = oauthClient.client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error sending HTTP request: %w", err)
	}

	if !discordRequest.DisableStatusCheck && response.StatusCode != discordRequest.ExpectedStatus {
		return nil, errors.StatusError{
			Code:     errors.StatusErrorCode(response.StatusCode),
			Response: response,
		}
	}

	return response, nil
}

func NewOAuthClient(clientId string, clientSecret string, redirectUri string, scopes []discord.OAuthScope) *OAuthClient {
	return &OAuthClient{
		clientId:     clientId,
		clientSecret: clientSecret,
		client:       http.DefaultClient,
		redirectUri:  redirectUri,
		Scopes:       scopes,
	}
}

func (oauthClient *OAuthClient) BuildAuthorizationURL(state string) string {
	params := url.Values{}

	params.Add("client_id", oauthClient.clientId)
	params.Add("response_type", "code")
	params.Add("redirect_uri", oauthClient.redirectUri)
	params.Add("state", state)

	return DISCORD_AUTHORIZATION_URL + "?" + params.Encode() + "&scope=" + discord.FormatScopesToParamString(oauthClient.Scopes)
}

func (oauthClient *OAuthClient) AuthorizeUserFromCode(code string) (*AuthorizedUser, error) {
	requestBody := &TokenRequest{
		GrantType:   GrantTypeAuthorizationCode,
		Code:        code,
		RedirectUri: oauthClient.redirectUri,
	}

	var tokenResponse TokenResponse

	_, err := oauthClient.MakeRequest(DiscordRequest{
		Method:         "POST",
		Endpoint:       "/oauth2/token",
		Body:           []byte(requestBody.ToString()),
		ExpectedStatus: 200,
		UnmarshalTo:    &tokenResponse,
	})
	if err != nil {
		return nil, err
	}

	return tokenResponse.ToAuthorizedUser(), nil
}
