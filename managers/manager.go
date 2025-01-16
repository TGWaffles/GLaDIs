package managers

import (
	"github.com/JackHumphries9/dapper-go/discord"
	"github.com/JackHumphries9/dapper-go/discord/interaction_type"
)

type InteractionManager interface {
	Type() interaction_type.InteractionType
	RouteInteraction(itx *discord.Interaction) discord.InteractionResponse
}
