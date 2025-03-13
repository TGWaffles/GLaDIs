package time_type

// Should probably use rune
type TimeType string

const (
	ShortTime     = "t"
	LongTime      = "T"
	ShortDate     = "d"
	LongDate      = "D"
	ShortDateTime = "f"
	LongDateTime  = "F"
	RelativeTime  = "R"
)
