package interactions

import "github.com/tgwaffles/lambda-discord-interactions-go/discord"

type ResolvedData struct {
	Users   *map[string]*discord.User   `json:"users,omitempty"`
	Members *map[string]*discord.Member `json:"members,omitempty"`
	Roles   *map[string]*discord.Role   `json:"roles,omitempty"`
}