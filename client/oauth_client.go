package client

import (
	"encoding/json"
	"fmt"
	"github.com/tgwaffles/gladis/client/errors"
	"github.com/tgwaffles/gladis/discord/oauth_scopes"
	"net/http"
	"net/url"
)

const DiscordAuthorizationUrl = "https://discordapp.com/api/oauth2/authorize"

type OAuthClient struct {
	ClientId     string
	ClientSecret string
	Client       *http.Client
	redirectUri  string
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

	response, err = oauthClient.Client.Do(request)
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

func NewOAuthClient(clientId string, clientSecret string, redirectUri string) *OAuthClient {
	return &OAuthClient{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		Client:       http.DefaultClient,
		redirectUri:  redirectUri,
	}
}

func (oauthClient *OAuthClient) BuildAuthorizationURL(scopes []oauth_scopes.OAuthScope, state string) string {
	params := url.Values{}

	params.Add("client_id", oauthClient.ClientId)
	params.Add("response_type", "code")
	params.Add("redirect_uri", oauthClient.redirectUri)
	params.Add("state", state)

	return DiscordAuthorizationUrl + "?" + params.Encode() + "&scope=" + oauth_scopes.FormatScopesToParamString(scopes)
}

func (oauthClient *OAuthClient) AuthorizeUserFromCode(code string) (*AuthorizedUser, error) {
	requestBody := &TokenRequest{
		GrantType:    GrantTypeAuthorizationCode,
		Code:         code,
		RedirectUri:  oauthClient.redirectUri,
		ClientId:     oauthClient.ClientId,
		ClientSecret: oauthClient.ClientSecret,
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

	return tokenResponse.ToAuthorizedUser(oauthClient), nil
}

func (oauthClient *OAuthClient) RefreshTokensForUser(refreshToken string) (tokenResponse *TokenResponse, err error) {
	requestBody := &TokenRequest{
		GrantType:    GrantTypeRefreshToken,
		RefreshToken: refreshToken,
		ClientId:     oauthClient.ClientId,
		ClientSecret: oauthClient.ClientSecret,
	}

	_, err = oauthClient.MakeRequest(DiscordRequest{
		ExpectedStatus: 200,
		Method:         "POST",
		Endpoint:       "/oauth2/token",
		Body:           []byte(requestBody.ToString()),
		UnmarshalTo:    tokenResponse,
	})

	if err != nil {
		return nil, err
	}

	return tokenResponse, nil
}

func (oauthClient *OAuthClient) RevokeTokensForUser(accessToken string) (err error) {
	requestBody := &RevokeTokenRequest{
		Token:         accessToken,
		TokenTypeHint: "access_token",
		ClientId:      oauthClient.ClientId,
		ClientSecret:  oauthClient.ClientSecret,
	}

	_, err = oauthClient.MakeRequest(DiscordRequest{
		ExpectedStatus: 204,
		Method:         "POST",
		Endpoint:       "/oauth2/token/revoke",
		Body:           []byte(requestBody.ToString()),
	})

	if err != nil {
		return err
	}

	return nil
}
