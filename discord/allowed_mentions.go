package discord

const (
	RoleMentionType     AllowedMentionType = "roles"
	UserMentionType                        = "users"
	EveryoneMentionType                    = "everyone"
)

type AllowedMentionType string

type AllowedMentions struct {
	Parse       []AllowedMentionType `json:"parse"`
	Roles       []Snowflake          `json:"roles"`
	Users       []Snowflake          `json:"users"`
	RepliedUser bool                 `json:"replied_user"`
}
