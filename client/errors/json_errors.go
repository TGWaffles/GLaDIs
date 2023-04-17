package errors

import (
	"fmt"
	"net/http"
)

const (
	GeneralError                                               ErrorCode = 0
	UnknownAccount                                                       = 10001
	UnknownApplication                                                   = 10002
	UnknownChannel                                                       = 10003
	UnknownGuild                                                         = 10004
	UnknownIntegration                                                   = 10005
	UnknownInvite                                                        = 10006
	UnknownMember                                                        = 10007
	UnknownMessage                                                       = 10008
	UnknownPermissionOverwrite                                           = 10009
	UnknownProvider                                                      = 10010
	UnknownRole                                                          = 10011
	UnknownToken                                                         = 10012
	UnknownUser                                                          = 10013
	UnknownEmoji                                                         = 10014
	UnknownWebhook                                                       = 10015
	UnknownWebhookService                                                = 10016
	UnknownSession                                                       = 10020
	UnknownBan                                                           = 10026
	UnknownSKU                                                           = 10027
	UnknownStoreListing                                                  = 10028
	UnknownEntitlement                                                   = 10029
	UnknownBuild                                                         = 10030
	UnknownLobby                                                         = 10031
	UnknownBranch                                                        = 10032
	UnknownStoreDirectoryLayout                                          = 10033
	UnknownRedistributable                                               = 10036
	UnknownGiftCode                                                      = 10038
	UnknownStream                                                        = 10049
	UnknownPremiumServerSubscribeCooldown                                = 10050
	UnknownGuildTemplate                                                 = 10057
	UnknownDiscoverableServerCategory                                    = 10059
	UnknownSticker                                                       = 10060
	UnknownInteraction                                                   = 10062
	UnknownApplicationCommand                                            = 10063
	UnknownVoiceState                                                    = 10065
	UnknownApplicationCommandPermissions                                 = 10066
	UnknownStageInstance                                                 = 10067
	UnknownGuildMemberVerificationForm                                   = 10068
	UnknownGuildWelcomeScreen                                            = 10069
	UnknownGuildScheduledEvent                                           = 10070
	UnknownGuildScheduledEventUser                                       = 10071
	UnknownTag                                                           = 10087
	BotsCannotUseThisEndpoint                                            = 20001
	OnlyBotsCanUseThisEndpoint                                           = 20002
	ExplicitContentCannotBeSent                                          = 20009
	NotAuthorizedToPerformAction                                         = 20012
	ActionCannotBePerformedSlowmodeRateLimit                             = 20016
	OnlyAccountOwnerCanPerformAction                                     = 20018
	MessageCannotBeEditedAnnouncementRateLimits                          = 20022
	UnderMinimumAge                                                      = 20024
	ChannelWriteRateLimitHit                                             = 20028
	ServerWriteRateLimitHit                                              = 20029
	DisallowedWordsInContent                                             = 20031
	GuildPremiumSubscriptionLevelTooLow                                  = 20035
	MaximumNumberOfGuildsReached                                         = 30001
	MaximumNumberOfFriendsReached                                        = 30002
	MaximumNumberOfPinsReached                                           = 30003
	MaximumNumberOfRecipientsReached                                     = 30004
	MaximumNumberOfGuildRolesReached                                     = 30005
	MaximumNumberOfWebhooksReached                                       = 30007
	MaximumNumberOfEmojisReached                                         = 30008
	MaximumNumberOfReactionsReached                                      = 30010
	MaximumNumberOfGroupDMsReached                                       = 30011
	MaximumNumberOfGuildChannelsReached                                  = 30013
	MaximumNumberOfAttachmentsReached                                    = 30015
	MaximumNumberOfInvitesReached                                        = 30016
	MaximumNumberOfAnimatedEmojisReached                                 = 30018
	MaximumNumberOfServerMembersReached                                  = 30019
	MaximumNumberOfServerCategoriesReached                               = 30030
	GuildAlreadyHasTemplate                                              = 30031
	MaximumNumberOfApplicationCommandsReached                            = 30032
	MaximumNumberOfThreadParticipantsReached                             = 30033
	MaximumNumberOfDailyApplicationCommandCreatesReached                 = 30034
	MaximumNumberOfBansForNonGuildMembersExceeded                        = 30035
	MaximumNumberOfBansFetchesReached                                    = 30037
	MaximumNumberOfUncompletedGuildScheduledEventsReached                = 30038
	MaximumNumberOfStickersReached                                       = 30039
	MaximumNumberOfPruneRequestsReached                                  = 30040
	MaximumNumberOfGuildWidgetSettingsUpdatesReached                     = 30042
	MaximumNumberOfEditsToMessagesOlderThan1HourReached                  = 30046
	MaximumNumberOfPinnedThreadsInForumChannelReached                    = 30047
	MaximumNumberOfTagsInForumChannelReached                             = 30048
	BitrateTooHighForChannelOfType                                       = 30052
	MaximumNumberOfPremiumEmojisReached                                  = 30056
	MaximumNumberOfWebhooksPerGuildReached                               = 30058
	MaximumNumberOfChannelPermissionOverwritesReached                    = 30060
	ChannelsForGuildAreTooLarge                                          = 30061
	UnauthorizedRequest                                                  = 40001
	AccountNeedsVerificationToPerformAction                              = 40002
	OpeningDirectMessagesTooFast                                         = 40003
	SendingMessagesTemporarilyDisabled                                   = 40004
	RequestEntityTooLarge                                                = 40005
	FeatureTemporarilyDisabledServerSide                                 = 40006
	UserIsBannedFromGuild                                                = 40007
	ConnectionRevoked                                                    = 40012
	TargetUserNotConnectedToVoice                                        = 40032
	MessageAlreadyCrossposted                                            = 40033
	ApplicationCommandWithNameAlreadyExists                              = 40041
	ApplicationInteractionFailedToSend                                   = 40043
	CannotSendMessageInForumChannel                                      = 40058
	InteractionAlreadyAcknowledged                                       = 40060
	TagNamesMustBeUnique                                                 = 40061
	ServiceResourceIsBeingRateLimited                                    = 40062
	TagsNotAvailableToBeSetByNonModerators                               = 40066
	TagIsRequiredToCreateForumPostInChannel                              = 40067
	MissingAccess                                                        = 50001
	InvalidAccountType                                                   = 50002
	CannotExecuteActionOnDMChannel                                       = 50003
	GuildWidgetDisabled                                                  = 50004
	CannotEditMessageAuthoredByAnotherUser                               = 50005
	CannotSendEmptyMessage                                               = 50006
	CannotSendMessageToUser                                              = 50007
	CannotSendMessageInNonTextChannel                                    = 50008
	ChannelVerificationLevelTooHigh                                      = 50009
	OAuth2ApplicationDoesNotHaveBot                                      = 50010
	OAuth2ApplicationLimitReached                                        = 50011
	InvalidOAuth2State                                                   = 50012
	PermissionsLackToPerformAction                                       = 50013
	InvalidAuthenticationTokenProvided                                   = 50014
	NoteWasTooLong                                                       = 50015
	InvalidNumberOfMessagesToDeleteProvided                              = 50016
	InvalidMFALevel                                                      = 50017
	MessageCanOnlyBePinnedInChannelItWasSentIn                           = 50019
	InvalidInviteCode                                                    = 50020
	CannotExecuteActionOnSystemMessage                                   = 50021
	CannotExecuteActionOnThisChannelType                                 = 50024
	InvalidOAuth2AccessTokenProvided                                     = 50025
	MissingRequiredOAuth2Scope                                           = 50026
	InvalidWebhookTokenProvided                                          = 50027
	InvalidRole                                                          = 50028
	InvalidRecipients                                                    = 50033
	MessageTooOldToBulkDelete                                            = 50034
	InvalidFormBody                                                      = 50035
	InviteAcceptedToGuildBotNotIn                                        = 50036
	InvalidActivityAction                                                = 50039
	InvalidAPIVersionProvided                                            = 50041
	FileUploadedExceedsMaximumSize                                       = 50045
	InvalidUploadedFile                                                  = 50046
	CannotSelfRedeemThisGift                                             = 50054
	InvalidGuild                                                         = 50055
	InvalidRequestOrigin                                                 = 50067
	InvalidMessageType                                                   = 50068
	PaymentSourceRequiredToRedeemGift                                    = 50070
	CannotModifySystemWebhook                                            = 50073
	CannotDeleteChannelRequiredForCommunityGuilds                        = 50074
	CannotEditStickersWithinMessage                                      = 50080
	InvalidStickerSent                                                   = 50081
	CannotPerformOperationOnArchivedThread                               = 50083
	InvalidThreadNotificationSettings                                    = 50084
	BeforeValueEarlierThanThreadCreationDate                             = 50085
	CommunityServerChannelsMustBeTextChannels                            = 50086
	EventEntityTypeDifferentFromEntityToStartFor                         = 50091
	ServerNotAvailableInYourLocation                                     = 50095
	ServerNeedsMonetizationEnabled                                       = 50097
	ServerNeedsMoreBoosts                                                = 50101
	InvalidJSONRequestBody                                               = 50109
	OwnershipCannotBeTransferredToBotUser                                = 50132
	FailedToResizeAssetBelowMaximumSize                                  = 50138
	CannotMixSubscriptionAndNonSubscriptionRolesForEmoji                 = 50144
	CannotConvertBetweenPremiumAndNormalEmoji                            = 50145
	UploadedFileNotFound                                                 = 50146
	CannotDeleteGuildSubscriptionIntegration                             = 50163
	PermissionToUseStickerNotGranted                                     = 50600
	TwoFactorAuthenticationRequired                                      = 60003
	NoUsersWithDiscordTagExist                                           = 80004
	ReactionBlocked                                                      = 90001
	ApplicationNotYetAvailable                                           = 110001
	APIResourceOverloaded                                                = 130000
	StageAlreadyOpen                                                     = 150006
	CannotReplyWithoutPermissionToReadMessageHistory                     = 160002
	ThreadAlreadyCreatedForMessage                                       = 160004
	ThreadLocked                                                         = 160005
	MaximumNumberOfActiveThreadsReached                                  = 160006
	MaximumNumberOfActiveAnnouncementThreadsReached                      = 160007
	InvalidJSONForUploadedLottieFile                                     = 170001
	UploadedLottiesCannotContainRasterizedImages                         = 170002
	StickerMaximumFramerateExceeded                                      = 170003
	StickerFrameCountExceedsMaximum                                      = 170004
	LottieAnimationMaximumDimensionsExceeded                             = 170005
	StickerFrameRateTooSmallOrTooLarge                                   = 170006
	StickerAnimationDurationExceedsMaximum                               = 170007
	CannotUpdateFinishedEvent                                            = 180000
	FailedToCreateStageForStageEvent                                     = 180002
	MessageBlockedByAutomaticModeration                                  = 200000
	TitleBlockedByAutomaticModeration                                    = 200001
	WebhooksPostedToForumChannelsMustHaveThreadNameOrID                  = 220001
	WebhooksPostedToForumChannelsCannotHaveBothThreadNameAndID           = 220002
	WebhooksCanOnlyCreateThreadsInForumChannels                          = 220003
	WebhookServicesCannotBeUsedInForumChannels                           = 220004
	MessageBlockedByHarmfulLinksFilter                                   = 240000
)

type ErrorCode int

type DiscordError struct {
	Code     ErrorCode
	Response *http.Response
}

func (err ErrorCode) GetEnglishError() string {
	switch err {
	case GeneralError:
		return "General error"
	case UnknownAccount:
		return "Unknown account"
	case UnknownApplication:
		return "Unknown application"
	case UnknownChannel:
		return "Unknown channel"
	case UnknownGuild:
		return "Unknown guild"
	case UnknownIntegration:
		return "Unknown integration"
	case UnknownInvite:
		return "Unknown invite"
	case UnknownMember:
		return "Unknown member"
	case UnknownMessage:
		return "Unknown message"
	case UnknownPermissionOverwrite:
		return "Unknown permission overwrite"
	case UnknownProvider:
		return "Unknown provider"
	case UnknownRole:
		return "Unknown role"
	case UnknownToken:
		return "Unknown token"
	case UnknownUser:
		return "Unknown user"
	case UnknownEmoji:
		return "Unknown emoji"
	case UnknownWebhook:
		return "Unknown webhook"
	case UnknownWebhookService:
		return "Unknown webhook service"
	case UnknownSession:
		return "Unknown session"
	case UnknownBan:
		return "Unknown ban"
	case UnknownSKU:
		return "Unknown SKU"
	case UnknownStoreListing:
		return "Unknown store listing"
	case UnknownEntitlement:
		return "Unknown entitlement"
	case UnknownBuild:
		return "Unknown build"
	case UnknownLobby:
		return "Unknown lobby"
	case UnknownBranch:
		return "Unknown branch"
	case UnknownStoreDirectoryLayout:
		return "Unknown store directory layout"
	case UnknownRedistributable:
		return "Unknown redistributable"
	case UnknownGiftCode:
		return "Unknown gift code"
	case UnknownStream:
		return "Unknown stream"
	case UnknownPremiumServerSubscribeCooldown:
		return "Unknown premium server subscribe cooldown"
	case UnknownGuildTemplate:
		return "Unknown guild template"
	case UnknownDiscoverableServerCategory:
		return "Unknown discoverable server category"
	case UnknownSticker:
		return "Unknown sticker"
	case UnknownInteraction:
		return "Unknown interaction"
	case UnknownApplicationCommand:
		return "Unknown application command"
	case UnknownVoiceState:
		return "Unknown voice state"
	case UnknownApplicationCommandPermissions:
		return "Unknown application command permissions"
	case UnknownStageInstance:
		return "Unknown stage instance"
	case UnknownGuildMemberVerificationForm:
		return "Unknown guild member verification form"
	case UnknownGuildWelcomeScreen:
		return "Unknown guild welcome screen"
	case UnknownGuildScheduledEvent:
		return "Unknown guild scheduled event"
	case UnknownGuildScheduledEventUser:
		return "Unknown guild scheduled event user"
	case UnknownTag:
		return "Unknown tag"
	case BotsCannotUseThisEndpoint:
		return "Bots cannot use this endpoint"
	case OnlyBotsCanUseThisEndpoint:
		return "Only bots can use this endpoint"
	case ExplicitContentCannotBeSent:
		return "Explicit content cannot be sent"
	case NotAuthorizedToPerformAction:
		return "Not authorized to perform action"
	case ActionCannotBePerformedSlowmodeRateLimit:
		return "Action cannot be performed: slowmode rate limit"
	case OnlyAccountOwnerCanPerformAction:
		return "Only account owner can perform action"
	case MessageCannotBeEditedAnnouncementRateLimits:
		return "Message cannot be edited: announcement rate limits"
	case UnderMinimumAge:
		return "Under minimum age"
	case ChannelWriteRateLimitHit:
		return "Channel write rate limit hit"
	case ServerWriteRateLimitHit:
		return "Server write rate limit hit"
	case DisallowedWordsInContent:
		return "Disallowed words in content"
	case GuildPremiumSubscriptionLevelTooLow:
		return "Guild premium subscription level too low"
	case MaximumNumberOfGuildsReached:
		return "Maximum number of guilds reached"
	case MaximumNumberOfFriendsReached:
		return "Maximum number of friends reached"
	case MaximumNumberOfPinsReached:
		return "Maximum number of pins reached"
	case MaximumNumberOfRecipientsReached:
		return "Maximum number of recipients reached"
	case MaximumNumberOfGuildRolesReached:
		return "Maximum number of guild roles reached"
	case MaximumNumberOfWebhooksReached:
		return "Maximum number of webhooks reached"
	case MaximumNumberOfEmojisReached:
		return "Maximum number of emojis reached"
	case MaximumNumberOfReactionsReached:
		return "Maximum number of reactions reached"
	case MaximumNumberOfGroupDMsReached:
		return "Maximum number of group DMs reached"
	case MaximumNumberOfGuildChannelsReached:
		return "Maximum number of guild channels reached"
	case MaximumNumberOfAttachmentsReached:
		return "Maximum number of attachments reached"
	case MaximumNumberOfInvitesReached:
		return "Maximum number of invites reached"
	case MaximumNumberOfAnimatedEmojisReached:
		return "Maximum number of animated emojis reached"
	case MaximumNumberOfServerMembersReached:
		return "Maximum number of server members reached"
	case MaximumNumberOfServerCategoriesReached:
		return "Maximum number of server categories reached"
	case GuildAlreadyHasTemplate:
		return "Guild already has template"
	case MaximumNumberOfApplicationCommandsReached:
		return "Maximum number of application commands reached"
	case MaximumNumberOfThreadParticipantsReached:
		return "Maximum number of thread participants reached"
	case MaximumNumberOfDailyApplicationCommandCreatesReached:
		return "Maximum number of daily application command creates reached"
	case MaximumNumberOfBansForNonGuildMembersExceeded:
		return "Maximum number of bans for non-guild members exceeded"
	case MaximumNumberOfBansFetchesReached:
		return "Maximum number of bans fetches reached"
	case MaximumNumberOfUncompletedGuildScheduledEventsReached:
		return "Maximum number of uncompleted guild scheduled events reached"
	case MaximumNumberOfStickersReached:
		return "Maximum number of stickers reached"
	case MaximumNumberOfPruneRequestsReached:
		return "Maximum number of prune requests reached"
	case MaximumNumberOfGuildWidgetSettingsUpdatesReached:
		return "Maximum number of guild widget settings updates reached"
	case MaximumNumberOfEditsToMessagesOlderThan1HourReached:
		return "Maximum number of edits to messages older than 1 hour reached"
	case MaximumNumberOfPinnedThreadsInForumChannelReached:
		return "Maximum number of pinned threads in forum channel reached"
	case MaximumNumberOfTagsInForumChannelReached:
		return "Maximum number of tags in forum channel reached"
	case BitrateTooHighForChannelOfType:
		return "Bitrate too high for channel of type"
	case MaximumNumberOfPremiumEmojisReached:
		return "Maximum number of premium emojis reached"
	case MaximumNumberOfWebhooksPerGuildReached:
		return "Maximum number of webhooks per guild reached"
	case MaximumNumberOfChannelPermissionOverwritesReached:
		return "Maximum number of channel permission overwrites reached"
	case ChannelsForGuildAreTooLarge:
		return "Channels for guild are too large"
	case UnauthorizedRequest:
		return "Unauthorized request"
	case AccountNeedsVerificationToPerformAction:
		return "Account needs verification to perform action"
	case OpeningDirectMessagesTooFast:
		return "Opening direct messages too fast"
	case SendingMessagesTemporarilyDisabled:
		return "Sending messages temporarily disabled"
	case RequestEntityTooLarge:
		return "Request entity too large"
	case FeatureTemporarilyDisabledServerSide:
		return "Feature temporarily disabled server-side"
	case UserIsBannedFromGuild:
		return "User is banned from guild"
	case ConnectionRevoked:
		return "Connection revoked"
	case TargetUserNotConnectedToVoice:
		return "Target user not connected to voice"
	case MessageAlreadyCrossposted:
		return "Message already crossposted"
	case ApplicationCommandWithNameAlreadyExists:
		return "Application command with name already exists"
	case ApplicationInteractionFailedToSend:
		return "Application interaction failed to send"
	case CannotSendMessageInForumChannel:
		return "Cannot send message in forum channel"
	case InteractionAlreadyAcknowledged:
		return "Interaction already acknowledged"
	case TagNamesMustBeUnique:
		return "Tag names must be unique"
	case ServiceResourceIsBeingRateLimited:
		return "Service resource is being rate limited"
	case TagsNotAvailableToBeSetByNonModerators:
		return "Tags not available to be set by non-moderators"
	case TagIsRequiredToCreateForumPostInChannel:
		return "Tag is required to create forum post in channel"
	case MissingAccess:
		return "Missing access"
	case InvalidAccountType:
		return "Invalid account type"
	case CannotExecuteActionOnDMChannel:
		return "Cannot execute action on DM channel"
	case GuildWidgetDisabled:
		return "Guild widget disabled"
	case CannotEditMessageAuthoredByAnotherUser:
		return "Cannot edit message authored by another user"
	case CannotSendEmptyMessage:
		return "Cannot send empty message"
	case CannotSendMessageToUser:
		return "Cannot send message to user"
	case CannotSendMessageInNonTextChannel:
		return "Cannot send message in non-text channel"
	case ChannelVerificationLevelTooHigh:
		return "Channel verification level too high"
	case OAuth2ApplicationDoesNotHaveBot:
		return "OAuth2 application does not have bot"
	case OAuth2ApplicationLimitReached:
		return "OAuth2 application limit reached"
	case InvalidOAuth2State:
		return "Invalid OAuth2 state"
	case PermissionsLackToPerformAction:
		return "Permissions lack to perform action"
	case InvalidAuthenticationTokenProvided:
		return "Invalid authentication token provided"
	case NoteWasTooLong:
		return "Note was too long"
	case InvalidNumberOfMessagesToDeleteProvided:
		return "Invalid number of messages to delete provided"
	case InvalidMFALevel:
		return "Invalid MFA level"
	case MessageCanOnlyBePinnedInChannelItWasSentIn:
		return "Message can only be pinned in channel it was sent in"
	case InvalidInviteCode:
		return "Invalid invite code"
	case CannotExecuteActionOnSystemMessage:
		return "Cannot execute action on system message"
	case CannotExecuteActionOnThisChannelType:
		return "Cannot execute action on this channel type"
	case InvalidOAuth2AccessTokenProvided:
		return "Invalid OAuth2 access token provided"
	case MissingRequiredOAuth2Scope:
		return "Missing required OAuth2 scope"
	case InvalidWebhookTokenProvided:
		return "Invalid webhook token provided"
	case InvalidRole:
		return "Invalid role"
	case InvalidRecipients:
		return "Invalid recipients"
	case MessageTooOldToBulkDelete:
		return "Message too old to bulk delete"
	case InvalidFormBody:
		return "Invalid form body"
	case InviteAcceptedToGuildBotNotIn:
		return "Invite accepted to guild bot not in"
	case InvalidActivityAction:
		return "Invalid activity action"
	case InvalidAPIVersionProvided:
		return "Invalid API version provided"
	case FileUploadedExceedsMaximumSize:
		return "File uploaded exceeds maximum size"
	case InvalidUploadedFile:
		return "Invalid uploaded file"
	case CannotSelfRedeemThisGift:
		return "Cannot self-redeem this gift"
	case InvalidGuild:
		return "Invalid guild"
	case InvalidRequestOrigin:
		return "Invalid request origin"
	case InvalidMessageType:
		return "Invalid message type"
	case PaymentSourceRequiredToRedeemGift:
		return "Payment source required to redeem gift"
	case CannotModifySystemWebhook:
		return "Cannot modify system webhook"
	case CannotDeleteChannelRequiredForCommunityGuilds:
		return "Cannot delete channel required for community guilds"
	case CannotEditStickersWithinMessage:
		return "Cannot edit stickers within message"
	case InvalidStickerSent:
		return "Invalid sticker sent"
	case CannotPerformOperationOnArchivedThread:
		return "Cannot perform operation on archived thread"
	case InvalidThreadNotificationSettings:
		return "Invalid thread notification settings"
	case BeforeValueEarlierThanThreadCreationDate:
		return "Before value earlier than thread creation date"
	case CommunityServerChannelsMustBeTextChannels:
		return "Community server channels must be text channels"
	case EventEntityTypeDifferentFromEntityToStartFor:
		return "Event entity type different from entity to start for"
	case ServerNotAvailableInYourLocation:
		return "Server not available in your location"
	case ServerNeedsMonetizationEnabled:
		return "Server needs monetization enabled"
	case ServerNeedsMoreBoosts:
		return "Server needs more boosts"
	case InvalidJSONRequestBody:
		return "Invalid JSON request body"
	case OwnershipCannotBeTransferredToBotUser:
		return "Ownership cannot be transferred to bot user"
	case FailedToResizeAssetBelowMaximumSize:
		return "Failed to resize asset below maximum size"
	case CannotMixSubscriptionAndNonSubscriptionRolesForEmoji:
		return "Cannot mix subscription and non-subscription roles for emoji"
	case CannotConvertBetweenPremiumAndNormalEmoji:
		return "Cannot convert between premium and normal emoji"
	case UploadedFileNotFound:
		return "Uploaded file not found"
	case CannotDeleteGuildSubscriptionIntegration:
		return "Cannot delete guild subscription integration"
	case PermissionToUseStickerNotGranted:
		return "Permission to use sticker not granted"
	case TwoFactorAuthenticationRequired:
		return "Two-factor authentication required"
	case NoUsersWithDiscordTagExist:
		return "No users with Discord tag exist"
	case ReactionBlocked:
		return "Reaction blocked"
	case ApplicationNotYetAvailable:
		return "Application not yet available"
	case APIResourceOverloaded:
		return "API resource overloaded"
	case StageAlreadyOpen:
		return "Stage already open"
	case CannotReplyWithoutPermissionToReadMessageHistory:
		return "Cannot reply without permission to read message history"
	case ThreadAlreadyCreatedForMessage:
		return "Thread already created for message"
	case ThreadLocked:
		return "Thread locked"
	case MaximumNumberOfActiveThreadsReached:
		return "Maximum number of active threads reached"
	case MaximumNumberOfActiveAnnouncementThreadsReached:
		return "Maximum number of active announcement threads reached"
	case InvalidJSONForUploadedLottieFile:
		return "Invalid JSON for uploaded Lottie file"
	case UploadedLottiesCannotContainRasterizedImages:
		return "Uploaded Lotties cannot contain rasterized images"
	case StickerMaximumFramerateExceeded:
		return "Sticker maximum framerate exceeded"
	case StickerFrameCountExceedsMaximum:
		return "Sticker frame count exceeds maximum"
	case LottieAnimationMaximumDimensionsExceeded:
		return "Lottie animation maximum dimensions exceeded"
	case StickerFrameRateTooSmallOrTooLarge:
		return "Sticker frame rate too small or too large"
	case StickerAnimationDurationExceedsMaximum:
		return "Sticker animation duration exceeds maximum"
	case CannotUpdateFinishedEvent:
		return "Cannot update finished event"
	case FailedToCreateStageForStageEvent:
		return "Failed to create stage for stage event"
	case MessageBlockedByAutomaticModeration:
		return "Message blocked by automatic moderation"
	case TitleBlockedByAutomaticModeration:
		return "Title blocked by automatic moderation"
	case WebhooksPostedToForumChannelsMustHaveThreadNameOrID:
		return "Webhooks posted to forum channels must have thread name or ID"
	case WebhooksPostedToForumChannelsCannotHaveBothThreadNameAndID:
		return "Webhooks posted to forum channels cannot have both thread name and ID"
	case WebhooksCanOnlyCreateThreadsInForumChannels:
		return "Webhooks can only create threads in forum channels"
	case WebhookServicesCannotBeUsedInForumChannels:
		return "Webhook services cannot be used in forum channels"
	case MessageBlockedByHarmfulLinksFilter:
		return "Message blocked by harmful links filter"
	default:
		return "Unknown error code"
	}
}

func (e *DiscordError) Error() string {
	return fmt.Sprintf("Discord Error (%d): %s", e.Code, e.Code.GetEnglishError())
}
