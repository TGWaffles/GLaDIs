package client

import (
	"testing"

	"github.com/JackHumphries9/dapper-go/discord/oauth_scopes"
)

const (
	ClientId     = "changeme"
	ClientSecret = "changeme"
	RedirectUri  = "changeme"
)

var client = NewOAuthClient(ClientId, ClientSecret, RedirectUri)

func TestOAuthClient_CreateLink(t *testing.T) {

	t.Log(client.BuildAuthorizationURL([]oauth_scopes.OAuthScope{oauth_scopes.IDENFITY, oauth_scopes.GUILDS}, "test"))
}

func TestOAuthClient_AuthorizeUserFromCode(t *testing.T) {
	authedUser, err := client.AuthorizeUserFromCode("changeme")

	if err != nil {
		t.Error(err)
		panic(err)
	}

	t.Log(authedUser)

	t.Log("refreshing token")

	err = authedUser.RefreshTokens()

	if err != nil {
		t.Error()
		panic(err)

	}

	t.Log(authedUser)

	t.Log("getting user")

	user, err := authedUser.FetchUser()

	if err != nil {
		t.Error(err)
		panic(err)

	}

	t.Log(user)

	t.Log("getting guilds")

	//GetGuilds

	guilds, err := authedUser.FetchGuilds()

	if err != nil {
		t.Error(err)
		panic(err)
	}

	t.Log(guilds)

	t.Log("revoking token")

	err = authedUser.RevokeTokens()

	if err != nil {
		t.Error(err)
		panic(err)
	}

	t.Log(authedUser)
}
