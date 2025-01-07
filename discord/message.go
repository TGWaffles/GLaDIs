package discord

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/JackHumphries9/dapper-go/discord/channel_type"
	"github.com/JackHumphries9/dapper-go/discord/message_activity_type"
	"github.com/JackHumphries9/dapper-go/discord/message_type"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Message struct {
	Id                   Snowflake                `json:"id"`
	ChannelId            Snowflake                `json:"channel_id"`
	Author               *User                    `json:"author"`
	Content              string                   `json:"content"`
	Timestamp            time.Time                `json:"timestamp"`
	EditedTimestamp      time.Time                `json:"edited_timestamp"`
	Tts                  bool                     `json:"tts"`
	MentionEveryone      bool                     `json:"mention_everyone"`
	Mentions             []User                   `json:"mentions"`
	MentionRoles         []Snowflake              `json:"mention_roles"`
	MentionChannels      []ChannelMention         `json:"mention_channels"`
	Attachments          []Attachment             `json:"attachments"`
	Embeds               []Embed                  `json:"embeds"`
	Reactions            []Reaction               `json:"reactions,omitempty"`
	Nonce                interface{}              `json:"nonce,omitempty"`
	Pinned               bool                     `json:"pinned"`
	WebhookId            *Snowflake               `json:"webhook_id,omitempty"`
	Type                 message_type.MessageType `json:"type"`
	Activity             *MessageActivity         `json:"activity,omitempty"`
	Application          *Application             `json:"application,omitempty"`
	ApplicationId        *Snowflake               `json:"application_id,omitempty"`
	MessageReference     *MessageReference        `json:"message_reference,omitempty"`
	Flags                *int                     `json:"flags,omitempty"`
	ReferencedMessage    *Message                 `json:"referenced_message,omitempty"`
	Interaction          *MessageInteraction      `json:"interaction,omitempty"`
	Thread               *Channel                 `json:"thread,omitempty"`
	Components           []MessageComponent       `json:"components,omitempty"`
	StickerItems         []StickerItem            `json:"sticker_items,omitempty"`
	Position             *int                     `json:"position,omitempty"`
	RoleSubscriptionData *RoleSubscriptionData    `json:"role_subscription_data,omitempty"`
}

func (m *Message) UnmarshalDynamoDBAttributeValue(value *dynamodb.AttributeValue) (err error) {
	type Alias Message

	var inner Alias

	dataMap := value.M
	if dataMap == nil {
		return fmt.Errorf("error unmarshalling message into map, map empty: %v", value)
	}

	componentsListAttr := dataMap["components"]
	if componentsListAttr == nil || len(componentsListAttr.L) == 0 {
		err = dynamodbattribute.UnmarshalMap(dataMap, &inner)
		if err != nil {
			return fmt.Errorf("error unmarshalling message without components: %w", err)
		}
		*m = Message(inner)
		return nil
	}

	components := make([]MessageComponent, len(componentsListAttr.L))
	for i, component := range componentsListAttr.L {
		var messageComponent MessageComponentWrapper
		err = dynamodbattribute.UnmarshalMap(component.M, &messageComponent)
		if err != nil {
			return fmt.Errorf("error unmarshalling message component: %w", err)
		}
		components[i] = messageComponent.component
	}

	delete(dataMap, "components")

	err = dynamodbattribute.UnmarshalMap(dataMap, &inner)
	if err != nil {
		return fmt.Errorf("error unmarshalling message without components: %w", err)
	}
	*m = Message(inner)

	m.Components = components

	return nil
}

func (m *Message) UnmarshalJSON(data []byte) error {
	type Alias Message

	var inner Alias

	dataMap := make(map[string]json.RawMessage)
	err := json.Unmarshal(data, &dataMap)
	if err != nil {
		return fmt.Errorf("error unmarshalling message into map: %w", err)
	}

	componentsListIface := dataMap["components"]
	if componentsListIface == nil {
		err = json.Unmarshal(data, &inner)
		if err != nil {
			return fmt.Errorf("error unmarshalling message: %w", err)
		}
		*m = Message(inner)
		return nil
	}
	var componentsList []json.RawMessage
	err = json.Unmarshal(componentsListIface, &componentsList)
	if len(componentsList) == 0 {
		err = json.Unmarshal(data, &inner)
		if err != nil {
			return fmt.Errorf("error unmarshalling message: %w", err)
		}
		*m = Message(inner)
		return nil
	}

	components := make([]MessageComponent, len(componentsList))
	for i, component := range componentsList {
		var messageComponent MessageComponentWrapper
		err = json.Unmarshal(component, &messageComponent)
		if err != nil {
			return err
		}
		components[i] = messageComponent.component
	}

	delete(dataMap, "components")
	newData, err := json.Marshal(dataMap)
	if err != nil {
		return fmt.Errorf("error marshalling message without components: %w", err)
	}

	err = json.Unmarshal(newData, &inner)
	if err != nil {
		return fmt.Errorf("error unmarshalling message without components: %w", err)
	}
	*m = Message(inner)

	m.Components = components

	return nil
}

type MessageReference struct {
	MessageId Snowflake `json:"message_id,omitempty"`
	// Optional when sending a reply
	ChannelId       Snowflake `json:"channel_id,omitempty"`
	GuildId         Snowflake `json:"guild_id,omitempty"`
	FailIfNotExists bool      `json:"fail_if_not_exists,omitempty"`
}

type ChannelMention struct {
	Id      Snowflake                `json:"id"`
	GuildId Snowflake                `json:"guild_id"`
	Type    channel_type.ChannelType `json:"type"`
	Name    string                   `json:"name"`
}

type MessageActivity struct {
	Type    message_activity_type.MessageActivityType `json:"type"`
	PartyId string                                    `json:"party_id"`
}

type MessageInteraction struct {
	Id     Snowflake `json:"id"`
	Type   uint8     `json:"type"`
	Name   string    `json:"name"`
	User   User      `json:"user"`
	Member *Member   `json:"member"`
}

type RoleSubscriptionData struct {
	RoleSubscriptionListingId Snowflake `json:"role_subscription_listing_id"`
	TierName                  string    `json:"tier_name"`
	TotalMonthsSubscribed     int       `json:"total_months_subscribed"`
	IsRenewal                 bool      `json:"is_renewal"`
}
