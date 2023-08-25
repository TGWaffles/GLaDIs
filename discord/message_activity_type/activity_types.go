package message_activity_type

const (
	Join MessageActivityType = iota + 1
	Spectate
	Listen
	JoinRequest
)

type MessageActivityType int
