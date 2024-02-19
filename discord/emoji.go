package discord

type Emoji struct {
	Id            *Snowflake  `json:"id,omitempty"`
	Name          *string     `json:"name,omitempty"`
	Roles         []Snowflake `json:"roles,omitempty"`
	User          *User       `json:"user,omitempty"`
	RequireColons *bool       `json:"require_colons,omitempty"`
	Managed       *bool       `json:"managed,omitempty"`
	Animated      *bool       `json:"animated,omitempty"`
	Available     *bool       `json:"available,omitempty"`
}

func (e Emoji) String() string {
	emojiName := ""
	if e.Name != nil {
		emojiName = *e.Name
	}
	if e.Id == nil {
		return emojiName
	}

	if e.Animated != nil && *e.Animated {
		return "<a:" + emojiName + ":" + e.Id.String() + ">"
	}

	return "<:" + emojiName + ":" + e.Id.String() + ">"
}
