package mosir_sdk_go

import (
	"context"

	generated "github.com/mosir-social/mosir_sdk_go/internal/generated"
)

func (c *Client) GetAccountProfile(ctx context.Context, accountId string, username string) (data_ *generated.GetAccountProfileResponse, err_ error) {
	if c == nil {
		var zero0 *generated.GetAccountProfileResponse
		var zero1 error
		return zero0, zero1
	}
	return generated.GetAccountProfile(ctx, c.graphqlClient(), accountId, username)
}

func (c *Client) GetBlockedAccounts(ctx context.Context, cursor string, limit int) (data_ *generated.GetBlockedAccountsResponse, err_ error) {
	if c == nil {
		var zero0 *generated.GetBlockedAccountsResponse
		var zero1 error
		return zero0, zero1
	}
	return generated.GetBlockedAccounts(ctx, c.graphqlClient(), cursor, limit)
}

func (c *Client) GetCurrentAccount(ctx context.Context) (data_ *generated.GetCurrentAccountResponse, err_ error) {
	if c == nil {
		var zero0 *generated.GetCurrentAccountResponse
		var zero1 error
		return zero0, zero1
	}
	return generated.GetCurrentAccount(ctx, c.graphqlClient())
}

func (c *Client) GetDiscussions(ctx context.Context, cursor string, limit int) (data_ *generated.GetDiscussionsResponse, err_ error) {
	if c == nil {
		var zero0 *generated.GetDiscussionsResponse
		var zero1 error
		return zero0, zero1
	}
	return generated.GetDiscussions(ctx, c.graphqlClient(), cursor, limit)
}

func (c *Client) GetFeedPosts(ctx context.Context, cursor string, limit int) (data_ *generated.GetFeedPostsResponse, err_ error) {
	if c == nil {
		var zero0 *generated.GetFeedPostsResponse
		var zero1 error
		return zero0, zero1
	}
	return generated.GetFeedPosts(ctx, c.graphqlClient(), cursor, limit)
}

func (c *Client) GetFollowedAccounts(ctx context.Context, accountId string, cursor string, limit int) (data_ *generated.GetFollowedAccountsResponse, err_ error) {
	if c == nil {
		var zero0 *generated.GetFollowedAccountsResponse
		var zero1 error
		return zero0, zero1
	}
	return generated.GetFollowedAccounts(ctx, c.graphqlClient(), accountId, cursor, limit)
}

func (c *Client) GetFollowedPostCollections(ctx context.Context, cursor string, limit int) (data_ *generated.GetFollowedPostCollectionsResponse, err_ error) {
	if c == nil {
		var zero0 *generated.GetFollowedPostCollectionsResponse
		var zero1 error
		return zero0, zero1
	}
	return generated.GetFollowedPostCollections(ctx, c.graphqlClient(), cursor, limit)
}

func (c *Client) GetFollowingAccounts(ctx context.Context, accountId string, cursor string, limit int) (data_ *generated.GetFollowingAccountsResponse, err_ error) {
	if c == nil {
		var zero0 *generated.GetFollowingAccountsResponse
		var zero1 error
		return zero0, zero1
	}
	return generated.GetFollowingAccounts(ctx, c.graphqlClient(), accountId, cursor, limit)
}

func (c *Client) GetFollowingPosts(ctx context.Context, cursor string, limit int) (data_ *generated.GetFollowingPostsResponse, err_ error) {
	if c == nil {
		var zero0 *generated.GetFollowingPostsResponse
		var zero1 error
		return zero0, zero1
	}
	return generated.GetFollowingPosts(ctx, c.graphqlClient(), cursor, limit)
}

func (c *Client) GetHistoryPosts(ctx context.Context, cursor string, includeOwnPosts bool, limit int) (data_ *generated.GetHistoryPostsResponse, err_ error) {
	if c == nil {
		var zero0 *generated.GetHistoryPostsResponse
		var zero1 error
		return zero0, zero1
	}
	return generated.GetHistoryPosts(ctx, c.graphqlClient(), cursor, includeOwnPosts, limit)
}

func (c *Client) GetLinkPreview(ctx context.Context, url string) (data_ *generated.GetLinkPreviewResponse, err_ error) {
	if c == nil {
		var zero0 *generated.GetLinkPreviewResponse
		var zero1 error
		return zero0, zero1
	}
	return generated.GetLinkPreview(ctx, c.graphqlClient(), url)
}

func (c *Client) GetMedia(ctx context.Context, mediaId string) (data_ *generated.GetMediaResponse, err_ error) {
	if c == nil {
		var zero0 *generated.GetMediaResponse
		var zero1 error
		return zero0, zero1
	}
	return generated.GetMedia(ctx, c.graphqlClient(), mediaId)
}

func (c *Client) GetMutualFollowers(ctx context.Context, accountId string, cursor string, limit int) (data_ *generated.GetMutualFollowersResponse, err_ error) {
	if c == nil {
		var zero0 *generated.GetMutualFollowersResponse
		var zero1 error
		return zero0, zero1
	}
	return generated.GetMutualFollowers(ctx, c.graphqlClient(), accountId, cursor, limit)
}

func (c *Client) GetMyPostCollections(ctx context.Context, cursor string, limit int) (data_ *generated.GetMyPostCollectionsResponse, err_ error) {
	if c == nil {
		var zero0 *generated.GetMyPostCollectionsResponse
		var zero1 error
		return zero0, zero1
	}
	return generated.GetMyPostCollections(ctx, c.graphqlClient(), cursor, limit)
}

func (c *Client) GetNotifications(ctx context.Context, cursor string, filter generated.NotificationFilterInput, limit int) (data_ *generated.GetNotificationsResponse, err_ error) {
	if c == nil {
		var zero0 *generated.GetNotificationsResponse
		var zero1 error
		return zero0, zero1
	}
	return generated.GetNotifications(ctx, c.graphqlClient(), cursor, filter, limit)
}

func (c *Client) GetPopularPosts(ctx context.Context, cursor string, language string, limit int) (data_ *generated.GetPopularPostsResponse, err_ error) {
	if c == nil {
		var zero0 *generated.GetPopularPostsResponse
		var zero1 error
		return zero0, zero1
	}
	return generated.GetPopularPosts(ctx, c.graphqlClient(), cursor, language, limit)
}

func (c *Client) GetPost(ctx context.Context, postId string) (data_ *generated.GetPostResponse, err_ error) {
	if c == nil {
		var zero0 *generated.GetPostResponse
		var zero1 error
		return zero0, zero1
	}
	return generated.GetPost(ctx, c.graphqlClient(), postId)
}

func (c *Client) GetPostCollection(ctx context.Context, id string) (data_ *generated.GetPostCollectionResponse, err_ error) {
	if c == nil {
		var zero0 *generated.GetPostCollectionResponse
		var zero1 error
		return zero0, zero1
	}
	return generated.GetPostCollection(ctx, c.graphqlClient(), id)
}

func (c *Client) GetPostCollectionsByAuthor(ctx context.Context, authorID string, cursor string, limit int) (data_ *generated.GetPostCollectionsByAuthorResponse, err_ error) {
	if c == nil {
		var zero0 *generated.GetPostCollectionsByAuthorResponse
		var zero1 error
		return zero0, zero1
	}
	return generated.GetPostCollectionsByAuthor(ctx, c.graphqlClient(), authorID, cursor, limit)
}

func (c *Client) GetPostDraft(ctx context.Context, id string) (data_ *generated.GetPostDraftResponse, err_ error) {
	if c == nil {
		var zero0 *generated.GetPostDraftResponse
		var zero1 error
		return zero0, zero1
	}
	return generated.GetPostDraft(ctx, c.graphqlClient(), id)
}

func (c *Client) GetPostDrafts(ctx context.Context, cursor string, filter generated.PostDraftFilterInput, limit int) (data_ *generated.GetPostDraftsResponse, err_ error) {
	if c == nil {
		var zero0 *generated.GetPostDraftsResponse
		var zero1 error
		return zero0, zero1
	}
	return generated.GetPostDrafts(ctx, c.graphqlClient(), cursor, filter, limit)
}

func (c *Client) GetPostDraftsCount(ctx context.Context, filter generated.PostDraftFilterInput) (data_ *generated.GetPostDraftsCountResponse, err_ error) {
	if c == nil {
		var zero0 *generated.GetPostDraftsCountResponse
		var zero1 error
		return zero0, zero1
	}
	return generated.GetPostDraftsCount(ctx, c.graphqlClient(), filter)
}

func (c *Client) GetPostReactionDetails(ctx context.Context, cursor string, limit int, postId string, reactionType generated.ReactionTypeInput) (data_ *generated.GetPostReactionDetailsResponse, err_ error) {
	if c == nil {
		var zero0 *generated.GetPostReactionDetailsResponse
		var zero1 error
		return zero0, zero1
	}
	return generated.GetPostReactionDetails(ctx, c.graphqlClient(), cursor, limit, postId, reactionType)
}

func (c *Client) GetPostReactions(ctx context.Context, first int, postId string) (data_ *generated.GetPostReactionsResponse, err_ error) {
	if c == nil {
		var zero0 *generated.GetPostReactionsResponse
		var zero1 error
		return zero0, zero1
	}
	return generated.GetPostReactions(ctx, c.graphqlClient(), first, postId)
}

func (c *Client) GetProfileTagById(ctx context.Context, id string) (data_ *generated.GetProfileTagByIdResponse, err_ error) {
	if c == nil {
		var zero0 *generated.GetProfileTagByIdResponse
		var zero1 error
		return zero0, zero1
	}
	return generated.GetProfileTagById(ctx, c.graphqlClient(), id)
}

func (c *Client) GetProfileTagProfiles(ctx context.Context, cursor string, limit int, tagId string) (data_ *generated.GetProfileTagProfilesResponse, err_ error) {
	if c == nil {
		var zero0 *generated.GetProfileTagProfilesResponse
		var zero1 error
		return zero0, zero1
	}
	return generated.GetProfileTagProfiles(ctx, c.graphqlClient(), cursor, limit, tagId)
}

func (c *Client) GetReactedPosts(ctx context.Context, cursor string, limit int, reactionType generated.ReactionTypeInput) (data_ *generated.GetReactedPostsResponse, err_ error) {
	if c == nil {
		var zero0 *generated.GetReactedPostsResponse
		var zero1 error
		return zero0, zero1
	}
	return generated.GetReactedPosts(ctx, c.graphqlClient(), cursor, limit, reactionType)
}

func (c *Client) GetTopicFeedPosts(ctx context.Context, cursor string, limit int, topicId string) (data_ *generated.GetTopicFeedPostsResponse, err_ error) {
	if c == nil {
		var zero0 *generated.GetTopicFeedPostsResponse
		var zero1 error
		return zero0, zero1
	}
	return generated.GetTopicFeedPosts(ctx, c.graphqlClient(), cursor, limit, topicId)
}

func (c *Client) GetTopics(ctx context.Context, limit int) (data_ *generated.GetTopicsResponse, err_ error) {
	if c == nil {
		var zero0 *generated.GetTopicsResponse
		var zero1 error
		return zero0, zero1
	}
	return generated.GetTopics(ctx, c.graphqlClient(), limit)
}

func (c *Client) GetUnreadNotificationCount(ctx context.Context) (data_ *generated.GetUnreadNotificationCountResponse, err_ error) {
	if c == nil {
		var zero0 *generated.GetUnreadNotificationCountResponse
		var zero1 error
		return zero0, zero1
	}
	return generated.GetUnreadNotificationCount(ctx, c.graphqlClient())
}

func (c *Client) GetUserPosts(ctx context.Context, accountId string, cursor string, limit int, postType generated.PostType) (data_ *generated.GetUserPostsResponse, err_ error) {
	if c == nil {
		var zero0 *generated.GetUserPostsResponse
		var zero1 error
		return zero0, zero1
	}
	return generated.GetUserPosts(ctx, c.graphqlClient(), accountId, cursor, limit, postType)
}

func (c *Client) GetUserReactions(ctx context.Context, cursor string, limit int) (data_ *generated.GetUserReactionsResponse, err_ error) {
	if c == nil {
		var zero0 *generated.GetUserReactionsResponse
		var zero1 error
		return zero0, zero1
	}
	return generated.GetUserReactions(ctx, c.graphqlClient(), cursor, limit)
}

func (c *Client) ListProfileTags(ctx context.Context, cursor string, limit int) (data_ *generated.ListProfileTagsResponse, err_ error) {
	if c == nil {
		var zero0 *generated.ListProfileTagsResponse
		var zero1 error
		return zero0, zero1
	}
	return generated.ListProfileTags(ctx, c.graphqlClient(), cursor, limit)
}
