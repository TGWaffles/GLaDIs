package discord

type Attachment struct {
	ID          Snowflake `json:"id"`
	Filename    string    `json:"filename"`
	Description *string   `json:"description,omitempty"`
	ContentType *string   `json:"content_type,omitempty"`
	Size        int       `json:"size"`
	URL         string    `json:"url"`
	ProxyURL    string    `json:"proxy_url"`
	Height      *int      `json:"height,omitempty"`
	Width       *int      `json:"width,omitempty"`
	Ephemeral   *bool     `json:"ephemeral,omitempty"`
}
