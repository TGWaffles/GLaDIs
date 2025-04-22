package button_style

const (
	Primary ButtonStyle = iota + 1
	Secondary
	Success
	Danger
	Link
	Premium
)

type ButtonStyle uint8
