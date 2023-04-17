package client

import (
	"bytes"
	"io"
	"strings"
)

const (
	DiscordApiURL = "https://discord.com/api"
)

type DiscordRequest struct {
	ExpectedStatus int
	Method         string
	Endpoint       string
	Body           []byte

	UnmarshalTo interface{}

	DisableAuth        bool
	DisableStatusCheck bool
	AdditionalHeaders  map[string]string
}

func (discordRequest *DiscordRequest) ValidateEndpoint() {
	if len(discordRequest.Endpoint) == 0 {
		return
	}
	// Make sure the endpoint starts with a slash
	if discordRequest.Endpoint[0] != '/' {
		discordRequest.Endpoint = "/" + discordRequest.Endpoint
	}
}

func (discordRequest *DiscordRequest) validateMethod() {
	// Make sure the method is uppercase
	discordRequest.Method = strings.ToUpper(discordRequest.Method)
}

func (discordRequest *DiscordRequest) GetUrl() string {
	discordRequest.ValidateEndpoint()
	return DiscordApiURL + discordRequest.Endpoint
}

func (discordRequest *DiscordRequest) getBodyAsReader() io.Reader {
	if discordRequest.Body == nil || len(discordRequest.Body) == 0 {
		return nil
	}
	return bytes.NewReader(discordRequest.Body)
}
