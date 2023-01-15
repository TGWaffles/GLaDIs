package components

const (
	ActionRowType ComponentType = iota + 1
	ButtonType
	StringSelectType
	TextInputType
	UserSelectType
	RoleSelectType
	MentionableSelectType
	ChannelSelectType
)

// MessageComponent can be made of any variables, so it's a map until we parse it into a specific component.
type MessageComponent map[string]interface{}

type ComponentType uint8

type ActionRow struct {
	Components []MessageComponent `json:"components"`
}
