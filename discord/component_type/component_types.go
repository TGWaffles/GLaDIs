package component_type

const (
	ActionRow ComponentType = iota + 1
	Button
	StringSelect
	TextInput
	UserSelect
	RoleSelect
	MentionableSelect
	ChannelSelect
)

type ComponentType uint8
