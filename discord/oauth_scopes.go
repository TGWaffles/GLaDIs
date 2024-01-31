package discord

import "strings"

type OAuthScope string

const (
	OAUTHSCOPE_ACTIVITIES_READ                          OAuthScope = "activities.read"
	OAUTHSCOPE_ACTIVITIES_WRITE                                    = "activities.write"
	OAUTHSCOPE_APPLICATIONS_BUILDS_READ                            = "applications.builds.read"
	OAUTHSCOPE_APPLICATIONS_BUILDS_UPLOAD                          = "applications.builds.upload"
	OAUTHSCOPE_APPLICATIONS_COMMANDS                               = "applications.commands"
	OAUTHSCOPE_APPLICATIONS_COMMANDS_UPDATE                        = "applications.commands.update"
	OAUTHSCOPE_APPLICATIONS_COMMANDS_PERMISSIONS_UPDATE            = "applications.commands.permissions.update"
	OAUTHSCOPE_APPLICATIONS_ENTITLEMENTS                           = "applications.entitlements"
	OAUTHSCOPE_APPLICATIONS_ENTITLEMENTS_UPDATE                    = "applications.entitlements.update"
	OAUTHSCOPE_BOT                                                 = "bot"
	OAUTHSCOPE_CONNECTIONS                                         = "connections"
	OAUTHSCOPE_DM_CHANNEL_READ                                     = "dm_channels.read"
	OAUTHSCOPE_EMAIL                                               = "email"
	OAUTHSCOPE_GDM_JOIN                                            = "gdm.join"
	OAUTHSCOPE_GUILDS                                              = "guilds"
	OAUTHSCOPE_GUILDS_JOIN                                         = "guilds.join"
	OAUTHSCOPE_IDENFITY                                            = "identify"
	OAUTHSCOPE_MESSAGES_READ                                       = "messages.read"
	OAUTHSCOPE_RELATIONSHIPS_READ                                  = "relationships.read"
	OAUTHSCOPE_ROLE_CONNECTIONS_WRITE                              = "role_connections.write"
	OAUTHSCOPE_RPC                                                 = "rpc"
	OAUTHSCOPE_RPC_ACTIVITIES_WRITE                                = "rpc.activities.write"
	OAUTHSCOPE_RPC_NOTIFICATIONS_READ                              = "rpc.notifications.read"
	OAUTHSCOPE_RPC_VOICE_READ                                      = "rpc.voice.read"
	OAUTHSCOPE_RPC_VOICE_WRITE                                     = "rpc.voice.write"
	OAUTHSCOPE_VOICE_READ                                          = "voice"
	OAUTHSCOPE_WEBHOOK_INCOMING                                    = "webhook.incoming"
)

func FormatScopesToParamString(scopes []OAuthScope) string {
	scopeString := ""
	for _, scope := range scopes {
		scopeString += string(scope) + "%20"
	}
	return scopeString
}

func ParseParamStringToScopes(scopes string) []OAuthScope {
	scopeSlice := make([]OAuthScope, 0)
	for _, scope := range strings.Split(scopes, " ") {
		scopeSlice = append(scopeSlice, OAuthScope(scope))
	}
	return scopeSlice
}
