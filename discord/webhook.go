package discord

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const (
	webhookUrl = apiUrl + "/webhooks/%d/%s"
)

type WebhookRequest struct {
	Content         *string            `json:"content,omitempty"`
	Username        string             `json:"username,omitempty"`
	AvatarUrl       string             `json:"avatar_url,omitempty"`
	TTS             *bool              `json:"tts,omitempty"`
	Embeds          []Embed            `json:"embeds,omitempty"`
	AllowedMentions *AllowedMentions   `json:"allowed_mentions,omitempty"`
	Flags           *int               `json:"flags,omitempty"`
	Components      []MessageComponent `json:"components,omitempty"`
	Attachments     []Attachment       `json:"attachments,omitempty"`
	ThreadName      string             `json:"thread_name,omitempty"`
}

type WebhookMessageResponse struct {
}

type Webhook struct {
	Id            Snowflake  `json:"id"`
	Type          uint8      `json:"type"`
	GuildId       *Snowflake `json:"guild_id,omitempty"`
	ChannelId     *Snowflake `json:"channel_id,omitempty"`
	User          *User      `json:"user,omitempty"`
	Name          *string    `json:"name,omitempty"`
	Avatar        *string    `json:"avatar,omitempty"`
	Token         *string    `json:"token,omitempty"`
	ApplicationId *Snowflake `json:"application_id,omitempty"`
	SourceGuild   *Guild     `json:"source_guild,omitempty"`
	SourceChannel *Channel   `json:"source_channel,omitempty"`
	Url           *string    `json:"url,omitempty"`
}

func WebhookFromUrl(url string) (webhook *Webhook, err error) {
	// Webhook is in the format https://discord.com/api/webhooks/<id>/<token>
	splitUrl := strings.Split(url, "webhooks/")
	if len(splitUrl) != 2 {
		return nil, fmt.Errorf("invalid webhook url - missing 'webhooks/' (given: %s)", url)
	}

	splitUrl = strings.Split(splitUrl[1], "/")
	if len(splitUrl) != 2 {
		return nil, fmt.Errorf("invalid webhook url - missing id or token (given: %s)", url)
	}

	webhook = &Webhook{}
	webhookId, err := strconv.ParseUint(splitUrl[0], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid webhook url - invalid id (given: %s), error: %w", splitUrl[0], err)
	}
	webhook.Id = Snowflake(webhookId)
	token := strings.Trim(splitUrl[1], "/ \n\r\t") // Trim trailing slashes and whitespace
	webhook.Token = &token
	webhook.Url = &url

	return webhook, nil
}

func (hook *Webhook) GetUrl() string {
	if hook.Url == nil {
		url := fmt.Sprintf(webhookUrl, hook.Id, *hook.Token)
		hook.Url = &url
	}

	return *hook.Url
}

func (hook *Webhook) Send(req WebhookRequest) (returnedMessage *Message, err error) {
	return hook.SendWithContext(nil, req)
}

func (hook *Webhook) SendWithContext(ctx context.Context, req WebhookRequest) (returnedMessage *Message, err error) {
	data, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling data to JSON: %w", err)
	}
	var request *http.Request
	if ctx != nil {
		request, err = http.NewRequestWithContext(ctx, "POST", hook.GetUrl(), bytes.NewReader(data))
	} else {
		request, err = http.NewRequest("POST", hook.GetUrl(), bytes.NewReader(data))
	}
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %w", err)
	}

	request.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error sending HTTP request: %w", err)
	}

	if resp.StatusCode > 300 {
		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("error reading HTTP response body: %v\n", err)
			return nil, fmt.Errorf("expected status code 2xx, got %d", resp.StatusCode)
		}
		return nil, fmt.Errorf(
			"error sending webhook, status code %d (expected 2xx)\nresponse body: %s\nrequest body: %s",
			resp.StatusCode, string(responseBody), string(data))
	}

	if resp.StatusCode == 200 {
		returnedMessage = &Message{}
		err = json.NewDecoder(resp.Body).Decode(returnedMessage)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling response: %w", err)
		}
	}

	return returnedMessage, nil
}

type WebhookGetMessageRequest struct {
	// String so it can be "@original"
	MessageId string     `json:"-"` // Not sent in request body
	ThreadId  *Snowflake `json:"thread_id,omitempty"`
}

func (hook *Webhook) GetMessage(req WebhookGetMessageRequest) (message *Message, err error) {
	return hook.GetMessageWithContext(nil, req)
}

func (hook *Webhook) GetMessageWithContext(ctx context.Context, req WebhookGetMessageRequest) (message *Message, err error) {
	var body io.Reader
	if req.ThreadId != nil {
		data, err := json.Marshal(req)
		if err != nil {
			return nil, fmt.Errorf("error marshaling data to JSON: %w", err)
		}
		body = bytes.NewReader(data)
	}
	var request *http.Request
	if ctx != nil {
		request, err = http.NewRequestWithContext(ctx, "GET", hook.GetUrl()+fmt.Sprintf("/messages/%s", req.MessageId), body)
	} else {
		request, err = http.NewRequest("GET", hook.GetUrl()+fmt.Sprintf("/messages/%s", req.MessageId), body)
	}
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %w", err)
	}

	request.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error sending HTTP request: %w", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("expected status code 200, got %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&message)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON: %w", err)
	}

	return message, nil
}

func (hook *Webhook) EditMessage(messageId string, data ResponseEditData) error {
	return hook.EditMessageWithContext(nil, messageId, data)
}

func (hook *Webhook) EditMessageWithContext(ctx context.Context, messageId string, data ResponseEditData) error {
	err := data.Verify()
	if err != nil {
		return fmt.Errorf("error verifying edit data: %w", err)
	}

	body, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error marshaling data to JSON: %w", err)
	}
	var request *http.Request
	if ctx != nil {
		request, err = http.NewRequestWithContext(ctx, "PATCH", hook.GetUrl()+fmt.Sprintf("/messages/%s", messageId), bytes.NewReader(body))
	} else {
		request, err = http.NewRequest("PATCH", hook.GetUrl()+fmt.Sprintf("/messages/%s", messageId), bytes.NewReader(body))
	}
	if err != nil {
		return fmt.Errorf("error creating HTTP request: %w", err)
	}

	request.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("error sending HTTP request: %w", err)
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading HTTP response body: %w", err)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("expected status code 200, got %d. Response body: %s\nRequest body: %s",
			resp.StatusCode, string(responseBody), string(body))
	}

	return nil
}

func (hook *Webhook) DeleteMessage(messageId string) error {
	return hook.DeleteMessageWithContext(nil, messageId)
}

func (hook *Webhook) DeleteMessageWithContext(ctx context.Context, messageId string) (err error) {
	var request *http.Request
	if ctx != nil {
		request, err = http.NewRequestWithContext(ctx, "DELETE", hook.GetUrl()+fmt.Sprintf("/messages/%s", messageId), nil)
	} else {
		request, err = http.NewRequest("DELETE", hook.GetUrl()+fmt.Sprintf("/messages/%s", messageId), nil)
	}
	if err != nil {
		return fmt.Errorf("error creating HTTP request: %w", err)
	}
	request.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("error sending HTTP request: %w", err)
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading HTTP response body: %w", err)
	}

	if resp.StatusCode != 204 {
		return fmt.Errorf("expected status code 204, got %d. Response body: %s", resp.StatusCode, string(responseBody))
	}

	return nil
}
