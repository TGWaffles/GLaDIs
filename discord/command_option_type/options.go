package command_option_type

const (
	SubCommand CommandOptionType = iota + 1
	SubCommandGroup
	String
	Integer
	Boolean
	User
	Channel
	Role
	Mentionable
	Number
	Attachment
)

type CommandOptionType uint8
