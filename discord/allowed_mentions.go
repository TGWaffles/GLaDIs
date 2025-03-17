package discord

import "github.com/tgwaffles/gladis/discord/allowed_mention_type"

type AllowedMentions struct {
	Parse       []allowed_mention_type.AllowedMentionType `json:"parse"`
	Roles       []Snowflake                               `json:"roles"`
	Users       []Snowflake                               `json:"users"`
	RepliedUser bool                                      `json:"replied_user"`
}

func (allowedMentions AllowedMentions) Verify() error {
	for _, parseGroup := range allowedMentions.Parse {
		if parseGroup == allowed_mention_type.User && len(allowedMentions.Users) > 0 {
			return ErrInvalidParseGroup{
				parseGroup,
				allowedMentions,
			}
		}
		if parseGroup == allowed_mention_type.Role && len(allowedMentions.Roles) > 0 {
			return ErrInvalidParseGroup{
				parseGroup,
				allowedMentions,
			}
		}
	}

	return nil
}
