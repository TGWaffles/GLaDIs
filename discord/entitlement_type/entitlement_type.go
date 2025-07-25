package entitlement_type

type EntitlementType int

const (
	Purchase                EntitlementType = iota + 1 // 1 - Entitlement was purchased by user
	PremiumSubscription                                // 2 - Entitlement for Discord Nitro subscription
	DeveloperGift                                      // 3 - Entitlement was gifted by developer
	TestModeSubscription                               // 4 - Entitlement was purchased by a dev in application test mode
	FreePurchase                                       // 5 - Entitlement was granted when the SKU was free
	UserGift                                           // 6 - Entitlement was gifted by another user
	PremiumPurchase                                    // 7 - Entitlement was claimed by user for free as a Nitro Subscriber
	ApplicationSubscription                            // 8 - Entitlement was purchased as an app subscription
)
