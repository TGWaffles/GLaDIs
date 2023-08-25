package interaction_type

const (
	Ping InteractionType = iota + 1
	ApplicationCommand
	MessageComponent
	ApplicationCommandAutocomplete
	ModalSubmit
)

type InteractionType uint8
