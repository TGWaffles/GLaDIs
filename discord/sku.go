package discord

import (
	"github.com/tgwaffles/gladis/discord/sku_flags"
	"github.com/tgwaffles/gladis/discord/sku_types"
)

type SKU struct {
	ID            Snowflake          `json:"id"`
	Type          sku_type.SkuType   `json:"type"`
	ApplicationID Snowflake          `json:"application_id"`
	Name          string             `json:"name"`
	Slug          string             `json:"slug"`
	Flags         sku_flags.SKUFlags `json:"flags"`
}
