package message_flags

type MessageFlags uint

const (
	Crossposted                      MessageFlags = 1 << 0
	IsCrossposted                    MessageFlags = 1 << 1
	SupressEmbeds                    MessageFlags = 1 << 2
	SourceMessageDeleted             MessageFlags = 1 << 3
	Urgent                           MessageFlags = 1 << 4
	HasThread                        MessageFlags = 1 << 5
	Ephemeral                        MessageFlags = 1 << 6
	Loading                          MessageFlags = 1 << 7
	FailedToMentionSomeRolesInThread MessageFlags = 1 << 8
	SupressNotification              MessageFlags = 1 << 12
	IsVoiceMessage                   MessageFlags = 1 << 13
)

func (flags MessageFlags) HasFlag(flag MessageFlags) bool {
	return flags&flag != 0
}

func (flags *MessageFlags) AddFlag(flag MessageFlags) {
	*flags |= flag
}

func (flags *MessageFlags) RemoveFlag(flag MessageFlags) {
	*flags &^= flag
}
