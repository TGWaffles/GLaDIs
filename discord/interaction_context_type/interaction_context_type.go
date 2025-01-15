package interaction_context_type

type InteractionContextType int

const (
	Guild           InteractionContextType = 0
	BOT_DM                                 = 1
	PRIVATE_CHANNEL                        = 2
)
