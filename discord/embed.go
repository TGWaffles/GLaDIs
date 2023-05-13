package discord

import (
	"encoding/json"
	"fmt"
)

type Embed struct {
	Title       string         `json:"title,omitempty"`
	Type        string         `json:"type,omitempty"`
	Description string         `json:"description,omitempty"`
	URL         string         `json:"url,omitempty"`
	Timestamp   string         `json:"timestamp,omitempty"`
	Color       int            `json:"color,omitempty"`
	Footer      *EmbedFooter   `json:"footer,omitempty"`
	Image       *EmbedImage    `json:"image,omitempty"`
	Thumbnail   *EmbedImage    `json:"thumbnail,omitempty"`
	Video       *EmbedVideo    `json:"video,omitempty"`
	Provider    *EmbedProvider `json:"provider,omitempty"`
	Author      *EmbedAuthor   `json:"author,omitempty"`
	Fields      []EmbedField   `json:"fields,omitempty"`
}

type EmbedFooter struct {
	Text         string `json:"text"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

type EmbedImage struct {
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type EmbedVideo struct {
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type EmbedProvider struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

type EmbedAuthor struct {
	Name         string `json:"name,omitempty"`
	URL          string `json:"url,omitempty"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

func (embed *Embed) Verify() error {
	combinedLength := 0
	if len(embed.Title) > 256 {
		return EmbedError{embed, "embed title cannot be longer than 256 characters"}
	}
	combinedLength += len(embed.Title)

	if len(embed.Description) > 4096 {
		return EmbedError{embed, "embed description cannot be longer than 2048 characters"}
	}
	combinedLength += len(embed.Description)

	if len(embed.Fields) > 25 {
		return EmbedError{embed, "embed cannot have more than 25 fields"}
	}

	for _, field := range embed.Fields {
		if len(field.Name) > 256 {
			return EmbedError{embed, "embed field name cannot be longer than 256 characters"}
		}
		combinedLength += len(field.Name)

		if len(field.Value) > 1024 {
			return EmbedError{embed, "embed field value cannot be longer than 1024 characters"}
		}
		combinedLength += len(field.Value)
	}

	if embed.Footer != nil {
		if len(embed.Footer.Text) > 2048 {
			return EmbedError{embed, "embed footer text cannot be longer than 2048 characters"}
		}
		combinedLength += len(embed.Footer.Text)
	}

	if embed.Author != nil {
		if len(embed.Author.Name) > 256 {
			return EmbedError{embed, "embed author name cannot be longer than 256 characters"}
		}
		combinedLength += len(embed.Author.Name)
	}

	if combinedLength > 6000 {
		return EmbedError{embed, "embed cannot be longer than 6000 characters"}
	}

	return nil
}

type EmbedError struct {
	Embed   *Embed
	Message string
}

func (e EmbedError) Error() string {
	var embedRepresentation string
	componentMarshalled, err := json.Marshal(e.Embed)
	if err != nil {
		embedRepresentation = fmt.Sprintf("%v", e.Embed)
	} else {
		embedRepresentation = string(componentMarshalled)
	}
	return fmt.Sprintf("embed error: %s\nEmbed was: %s", e.Message, embedRepresentation)
}
