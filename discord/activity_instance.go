package discord

import "github.com/tgwaffles/gladis/discord/activity_location_kind"

type ActivityLocation struct {
	Id        string                                      `json:"id"`
	Kind      activity_location_kind.ActivityLocationKind `json:"kind"`
	ChannelId Snowflake                                   `json:"channel_id"`
	GuildId   *Snowflake                                  `json:"guild_id,omitempty"`
}

type ActivityInstance struct {
	ApplicationId Snowflake        `json:"application_id"`
	InstanceId    string           `json:"instance_id"`
	LaunchId      Snowflake        `json:"launch_id"`
	Location      ActivityLocation `json:"location"`
	Users         []Snowflake      `json:"users"`
}
