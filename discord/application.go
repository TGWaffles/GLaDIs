package discord

type Application struct {
	Id                             Snowflake      `json:"id"`
	Name                           string         `json:"name"`
	Icon                           string         `json:"icon"`
	Description                    string         `json:"description"`
	RpcOrigins                     *[]string      `json:"rpc_origins,omitempty"`
	BotPublic                      bool           `json:"bot_public"`
	BotRequireCodeGrant            bool           `json:"bot_require_code_grant"`
	TermsOfServiceUrl              *string        `json:"terms_of_service_url,omitempty"`
	PrivacyPolicyUrl               *string        `json:"privacy_policy_url,omitempty"`
	Owner                          *User          `json:"owner,omitempty"`
	VerifyKey                      string         `json:"verify_key"`
	Team                           *Team          `json:"team,omitempty"`
	GuildId                        *Snowflake     `json:"guild_id,omitempty"`
	PrimarySkuId                   *Snowflake     `json:"primary_sku_id,omitempty"`
	Slug                           *string        `json:"slug,omitempty"`
	CoverImage                     *string        `json:"cover_image,omitempty"`
	Flags                          *uint          `json:"flags,omitempty"`
	Tags                           *[]string      `json:"tags,omitempty"`
	InstallParams                  *InstallParams `json:"install_params,omitempty"`
	CustomInstallUrl               *string        `json:"custom_install_url,omitempty"`
	RoleConnectionsVerificationUrl *string        `json:"role_connections_verification_url,omitempty"`
}

type Team struct {
	Icon        string       `json:"icon"`
	Id          Snowflake    `json:"id"`
	Members     []TeamMember `json:"members"`
	Name        string       `json:"name"`
	OwnerUserId Snowflake    `json:"owner_user_id"`
}

type TeamMember struct {
	MembershipState int       `json:"membership_state"`
	Permissions     []string  `json:"permissions"`
	TeamId          Snowflake `json:"team_id"`
	User            *User     `json:"user"`
}

type InstallParams struct {
	Scopes      []string `json:"scopes"`
	Permissions string   `json:"permissions"`
}
