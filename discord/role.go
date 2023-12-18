package discord

import "encoding/json"

type Role struct {
	Id           Snowflake   `json:"id"`
	Name         string      `json:"name"`
	Color        int         `json:"color"`
	Hoist        bool        `json:"hoist"`
	Icon         *string     `json:"icon,omitempty"`
	UnicodeEmoji *string     `json:"unicode_emoji,omitempty"`
	Position     int         `json:"position"`
	Permissions  Permissions `json:"permissions"`
	Managed      bool        `json:"managed"`
	Mentionable  bool        `json:"mentionable"`
	Tags         *RoleTags   `json:"tags,omitempty"`
}

type RoleTags struct {
	BotId                 *Snowflake `json:"bot_id,omitempty"`
	IntegrationId         *Snowflake `json:"integration_id,omitempty"`
	SubscriptionListingId *Snowflake `json:"subscription_listing_id,omitempty"`
	// These fields should be set to "null" in the JSON if they are true, according to the Discord docs...
	PremiumSubscriber    bool `json:"premium_subscriber,omitempty"`
	AvailableForPurchase bool `json:"available_for_purchase,omitempty"`
	GuildConnections     bool `json:"guild_connections,omitempty"`
}

func (tags RoleTags) MarshalJSON() ([]byte, error) {
	raw := map[string]interface{}{}

	if tags.BotId != nil {
		raw["bot_id"] = tags.BotId
	}

	if tags.IntegrationId != nil {
		raw["integration_id"] = tags.IntegrationId
	}

	if tags.SubscriptionListingId != nil {
		raw["subscription_listing_id"] = tags.SubscriptionListingId
	}

	if tags.PremiumSubscriber {
		raw["premium_subscriber"] = nil
	}

	if tags.AvailableForPurchase {
		raw["available_for_purchase"] = nil
	}

	if tags.GuildConnections {
		raw["guild_connections"] = nil
	}

	return json.Marshal(raw)
}

func (tags *RoleTags) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	if v, present := raw["bot_id"]; present {
		botId, err := GetSnowflake(v)
		if err != nil {
			return err
		}
		tags.BotId = &botId
	}

	if v, present := raw["integration_id"]; present {
		integrationId, err := GetSnowflake(v)
		if err != nil {
			return err
		}
		tags.IntegrationId = &integrationId
	}

	if v, present := raw["subscription_listing_id"]; present {
		subId, err := GetSnowflake(v)
		if err != nil {
			return err
		}
		tags.SubscriptionListingId = &subId
	}

	_, present := raw["premium_subscriber"]
	tags.PremiumSubscriber = present

	_, present = raw["available_for_purchase"]
	tags.AvailableForPurchase = present

	_, present = raw["guild_connections"]
	tags.GuildConnections = present

	return nil
}
