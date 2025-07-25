package discord

type ResolvedData struct {
	Users       *map[string]*User       `json:"users,omitempty"`
	Members     *map[string]*Member     `json:"members,omitempty"`
	Roles       *map[string]*Role       `json:"roles,omitempty"`
	Channels    *map[string]*Channel    `json:"channels,omitempty"`
	Attachments *map[string]*Attachment `json:"attachments,omitempty"`
}
