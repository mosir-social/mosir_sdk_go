package mosir_sdk_go

import generated "github.com/mosir-social/mosir_sdk_go/internal/generated"

// Re-exported types used by generated client methods and subscription wrappers.
type (
	NotificationFilterInput             = generated.NotificationFilterInput
	PostDraftFilterInput                = generated.PostDraftFilterInput
	ReactionTypeInput                   = generated.ReactionTypeInput
	PostType                            = generated.PostType
	NotificationReceivedWsResponse      = generated.NotificationReceivedWsResponse
	PostCreatedByAuthorWsResponse       = generated.PostCreatedByAuthorWsResponse
	PostCreatedInCollectionWsResponse   = generated.PostCreatedInCollectionWsResponse
	PostDeletedWsResponse               = generated.PostDeletedWsResponse
	PostUpdatedWsResponse               = generated.PostUpdatedWsResponse
	ReplyCreatedUnderRootPostWsResponse = generated.ReplyCreatedUnderRootPostWsResponse
)

const (
	PostTypePost  = generated.PostTypePost
	PostTypeReply = generated.PostTypeReply
)
