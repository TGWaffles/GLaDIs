package discord

import "fmt"

type MessageAttachment interface {
	GetBytes() []byte
	GetFileName() string
	GetDiscordRef() string
	GetContentType() string
	ToDiscordAttachment(id Snowflake) Attachment
}

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

type BytesAttachment struct {
	bytes       []byte
	fileName    string
	contentType string
}

func NewBytesAttachment(data []byte, fileName string, contentType string) *BytesAttachment {
	return &BytesAttachment{
		bytes:       data,
		fileName:    fileName,
		contentType: contentType,
	}
}

func (ba *BytesAttachment) GetBytes() []byte {
	return ba.bytes
}

func (ba *BytesAttachment) GetFileName() string {
	return ba.fileName
}

func (ba *BytesAttachment) GetContentType() string {
	return ba.contentType
}

var a_desc = "A Description"

func (ba *BytesAttachment) ToDiscordAttachment(id Snowflake) Attachment {
	return Attachment{
		Filename:    ba.fileName,
		ContentType: &ba.contentType,
		ID:          id,
	}
}

func (ba *BytesAttachment) GetDiscordRef() string {
	return fmt.Sprintf("attachment://%s", ba.fileName)
}
