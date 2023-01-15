package commands

const (
	SubCommandType ApplicationCommandOptionType = iota + 1
	SubCommandGroupType
	StringType
	IntegerType
	BooleanType
	UserType
	ChannelType
	RoleType
	MentionableType
	NumberType
	AttachmentType
)

type ApplicationCommandOptionType uint8

type ApplicationCommandOption struct {
	Name    string                       `json:"name"`
	Type    ApplicationCommandOptionType `json:"type"`
	Value   *interface{}                 `json:"value,omitempty"`
	Options []ApplicationCommandOption   `json:"options,omitempty"`
	Focused bool                         `json:"focused,omitempty"`
}
