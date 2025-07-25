package sku_type

type SkuType int

const (
	Durable           SkuType = 2 // Durable one-time purchase
	Consumable        SkuType = 3 // Consumable one-time purchase
	Subscription      SkuType = 5 // Represents a recurring subscription
	SubscriptionGroup SkuType = 6 // System-generated group for each SUBSCRIPTION SKU created
)
