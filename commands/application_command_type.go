package commands

const (
	ChatInput ApplicationCommandType = iota + 1 // (slash-command)
	User
	Message
)

type ApplicationCommandType uint8
