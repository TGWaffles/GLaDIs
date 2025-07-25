package sku_flags

type SKUFlags uint

const (
	Available         SKUFlags = 1 << 2
	GuildSubscription SKUFlags = 1 << 7
	UserSubscription  SKUFlags = 1 << 8
)

func (flags SKUFlags) HasFlag(flag SKUFlags) bool {
	return flags&flag != 0
}

func (flags *SKUFlags) AddFlag(flag SKUFlags) {
	*flags |= flag
}

func (flags *SKUFlags) RemoveFlag(flag SKUFlags) {
	*flags &^= flag
}
