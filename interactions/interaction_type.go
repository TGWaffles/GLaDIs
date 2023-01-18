package interactions

const (
	PingInteractionType InteractionType = iota + 1
	ApplicationCommandInteractionType
	MessageComponentInteractionType
	ApplicationCommandAutocompleteInteractionType
	ModalSubmitInteractionType
)

type InteractionType uint8
