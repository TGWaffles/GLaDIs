package interaction_callback_type

const (
	Pong                     InteractionCallbackType = iota + 1
	ChannelMessageWithSource                         = iota + 3
	DeferredChannelMessageWithSource
	DeferredUpdateMessage
	UpdateMessage
	ApplicationCommandAutocompleteResult
	Modal
)

type InteractionCallbackType uint8
