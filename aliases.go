package mosir_sdk_go

import generated "github.com/mosir-social/mosir_sdk_go/internal/generated"

// Re-exported input types used by generated client methods.
type (
	NotificationFilterInput = generated.NotificationFilterInput
	PostDraftFilterInput = generated.PostDraftFilterInput
	ReactionTypeInput = generated.ReactionTypeInput
	PostType = generated.PostType
	NotificationReceivedWsResponse = generated.NotificationReceivedWsResponse
	PostCreatedByAuthorWsResponse = generated.PostCreatedByAuthorWsResponse
	PostCreatedInCollectionWsResponse = generated.PostCreatedInCollectionWsResponse
	PostDeletedWsResponse = generated.PostDeletedWsResponse
	PostUpdatedWsResponse = generated.PostUpdatedWsResponse
	ReplyCreatedUnderRootPostWsResponse = generated.ReplyCreatedUnderRootPostWsResponse
)

// Re-exported enum values used by tests and examples.
const (
	PostTypePost = generated.PostTypePost
	PostTypeReply = generated.PostTypeReply
)
