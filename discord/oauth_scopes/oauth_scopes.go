package oauth_scopes

import "strings"

type OAuthScope string

const (
	ACTIVITIES_READ                          OAuthScope = "activities.read"
	ACTIVITIES_WRITE                                    = "activities.write"
	APPLICATIONS_BUILDS_READ                            = "applications.builds.read"
	APPLICATIONS_BUILDS_UPLOAD                          = "applications.builds.upload"
	APPLICATIONS_COMMANDS                               = "applications.commands"
	APPLICATIONS_COMMANDS_UPDATE                        = "applications.commands.update"
	APPLICATIONS_COMMANDS_PERMISSIONS_UPDATE            = "applications.commands.permissions.update"
	APPLICATIONS_ENTITLEMENTS                           = "applications.entitlements"
	APPLICATIONS_ENTITLEMENTS_UPDATE                    = "applications.entitlements.update"
	BOT                                                 = "bot"
	CONNECTIONS                                         = "connections"
	DM_CHANNEL_READ                                     = "dm_channels.read"
	EMAIL                                               = "email"
	GDM_JOIN                                            = "gdm.join"
	GUILDS                                              = "guilds"
	GUILDS_JOIN                                         = "guilds.join"
	IDENFITY                                            = "identify"
	MESSAGES_READ                                       = "messages.read"
	RELATIONSHIPS_READ                                  = "relationships.read"
	ROLE_CONNECTIONS_WRITE                              = "role_connections.write"
	RPC                                                 = "rpc"
	RPC_ACTIVITIES_WRITE                                = "rpc.activities.write"
	RPC_NOTIFICATIONS_READ                              = "rpc.notifications.read"
	RPC_VOICE_READ                                      = "rpc.voice.read"
	RPC_VOICE_WRITE                                     = "rpc.voice.write"
	VOICE_READ                                          = "voice"
	WEBHOOK_INCOMING                                    = "webhook.incoming"
)

func FormatScopesToParamString(scopes []OAuthScope) string {
	scopeString := ""
	for i, scope := range scopes {
		if i == len(scopes) - 1 {
			scopeString += string(scope)
			continue
		}
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
