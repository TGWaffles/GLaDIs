package client

import (
	"github.com/tgwaffles/gladis/discord/oauth_scopes"
	"net/url"
)

type TokenGrantType string

const (
	GrantTypeAuthorizationCode TokenGrantType = "authorization_code"
	GrantTypeRefreshToken      TokenGrantType = "refresh_token"
)

type TokenRequest struct {
	GrantType    TokenGrantType `json:"grant_type"`
	Code         string         `json:"code,omitempty"`
	RedirectUri  string         `json:"redirect_uri,omitempty"`
	RefreshToken string         `json:"refresh_token,omitempty"`
}

type RevokeTokenRequest struct {
	Token         string `json:"token"`
	TokenTypeHint string `json:"token_type_hint"`
}

func (r *RevokeTokenRequest) ToValues() url.Values {
	formValues := url.Values{}

	formValues.Add("token", r.Token)

	if r.TokenTypeHint != "" {
		formValues.Add("token_type_hint", r.TokenTypeHint)
	}

	return formValues
}

func (r *RevokeTokenRequest) ToString() string {
	return r.ToValues().Encode()
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
}

func (t *TokenRequest) ToValues() url.Values {
	formValues := url.Values{}

	formValues.Add("grant_type", string(t.GrantType))

	if t.Code != "" {
		formValues.Add("code", t.Code)
	}

	if t.RedirectUri != "" {
		formValues.Add("redirect_uri", t.RedirectUri)
	}

	if t.RefreshToken != "" {
		formValues.Add("refresh_token", t.RefreshToken)
	}

	return formValues
}

func (t *TokenRequest) ToString() string {
	return t.ToValues().Encode()
}

func (response *TokenResponse) ToAuthorizedUser() *AuthorizedUser {
	return &AuthorizedUser{
		RefreshToken: response.RefreshToken,
		AccessToken:  response.AccessToken,
		ExpiresIn:    response.ExpiresIn,
		Scopes:       oauth_scopes.ParseParamStringToScopes(response.Scope),
	}
}
