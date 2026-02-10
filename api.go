// Package ofauth provides typed API methods
package ofauth

import (
	"context"
	"encoding/json"
	"fmt"
)

// ============================================================================
// Response Types
// ============================================================================

// GetAccessAnalyticsCampaignsChartResponse represents the response for GetAccessAnalyticsCampaignsChart
type GetAccessAnalyticsCampaignsChartResponse map[string]struct {
	Chart []struct {
	Count float64 `json:"count"`
	Date string `json:"date"`
} `json:"chart"`
	Delta float64 `json:"delta"`
	Gross *float64 `json:"gross,omitempty"`
	Total float64 `json:"total"`
}

// GetAccessAnalyticsCampaignsTopResponse represents the response for GetAccessAnalyticsCampaignsTop
type GetAccessAnalyticsCampaignsTopResponse struct {
	HasMore bool `json:"hasMore"`
	List []struct {
	CampaignCode float64 `json:"campaignCode"`
	CampaignName string `json:"campaignName"`
	CountSubscribers float64 `json:"countSubscribers"`
	CountTransitions float64 `json:"countTransitions"`
	CreatedAt string `json:"createdAt"`
	EndDate string `json:"endDate"`
	Id float64 `json:"id"`
} `json:"list"`
}

// ListAccessAnalyticsEarningsChargebacksResponse represents the response for ListAccessAnalyticsEarningsChargebacks
type ListAccessAnalyticsEarningsChargebacksResponse struct {
	HasMore bool `json:"hasMore"`
	List []struct {
	CreatedAt string `json:"createdAt"`
	Id float64 `json:"id"`
	Payment struct {
	Amounts struct {
	Fee float64 `json:"fee"`
	Gross float64 `json:"gross"`
	Net float64 `json:"net"`
	Tax float64 `json:"tax"`
	Vat float64 `json:"vat"`
} `json:"amounts"`
	CreatedAt string `json:"createdAt"`
	Currency string `json:"currency"`
	Description string `json:"description"`
	Id string `json:"id"`
	Status string `json:"status"`
	User struct {
	Avatar string `json:"avatar"`
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
} `json:"user"`
} `json:"payment"`
	PaymentType string `json:"paymentType"`
} `json:"list"`
	NextMarker float64 `json:"nextMarker"`
}

// GetAccessAnalyticsEarningsChartResponse represents the response for GetAccessAnalyticsEarningsChart
type GetAccessAnalyticsEarningsChartResponse struct {
	Chart []struct {
	Count float64 `json:"count"`
	Date string `json:"date"`
} `json:"chart"`
	Delta *float64 `json:"delta,omitempty"`
	Gross *float64 `json:"gross,omitempty"`
	Total float64 `json:"total"`
}

// ListAccessAnalyticsEarningsTransactionsResponse represents the response for ListAccessAnalyticsEarningsTransactions
type ListAccessAnalyticsEarningsTransactionsResponse struct {
	HasMore bool `json:"hasMore"`
	List []struct {
	Amounts struct {
	Fee float64 `json:"fee"`
	Gross float64 `json:"gross"`
	Net float64 `json:"net"`
	Tax float64 `json:"tax"`
	Vat float64 `json:"vat"`
} `json:"amounts"`
	CreatedAt string `json:"createdAt"`
	Currency string `json:"currency"`
	Description string `json:"description"`
	Id string `json:"id"`
	PayoutPendingDays float64 `json:"payoutPendingDays"`
	Status string `json:"status"`
	Type string `json:"type"`
	User struct {
	Avatar string `json:"avatar"`
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
} `json:"user"`
} `json:"list"`
	NextMarker float64 `json:"nextMarker"`
}

// ListAccessAnalyticsMassMessagesBuyersResponse represents the response for ListAccessAnalyticsMassMessagesBuyers
type ListAccessAnalyticsMassMessagesBuyersResponse struct {
	HasMore bool `json:"hasMore"`
	List []struct {
	Avatar string `json:"avatar"`
	AvatarThumbs *struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs,omitempty"`
	Id float64 `json:"id"`
	IsVerified bool `json:"isVerified"`
	Name string `json:"name"`
	Username string `json:"username"`
} `json:"list"`
	NextMarker *float64 `json:"nextMarker,omitempty"`
}

// GetAccessAnalyticsMassMessagesChartResponse represents the response for GetAccessAnalyticsMassMessagesChart
type GetAccessAnalyticsMassMessagesChartResponse struct {
	Messages struct {
	Chart []struct {
	Count float64 `json:"count"`
	Date string `json:"date"`
} `json:"chart"`
	Delta float64 `json:"delta"`
	Gross *float64 `json:"gross,omitempty"`
	Total float64 `json:"total"`
} `json:"messages"`
	Purchases struct {
	Chart []struct {
	Count float64 `json:"count"`
	Date string `json:"date"`
} `json:"chart"`
	Delta float64 `json:"delta"`
	Gross *float64 `json:"gross,omitempty"`
	Total float64 `json:"total"`
} `json:"purchases"`
}

// GetAccessAnalyticsMassMessagesPurchasedResponse represents the response for GetAccessAnalyticsMassMessagesPurchased
type GetAccessAnalyticsMassMessagesPurchasedResponse struct {
	HasMore bool `json:"hasMore"`
	Items []struct {
	CanSendMessageToBuyers *bool `json:"canSendMessageToBuyers,omitempty"`
	CanUnsend *bool `json:"canUnsend,omitempty"`
	Date string `json:"date"`
	GiphyId *string `json:"giphyId,omitempty"`
	Id float64 `json:"id"`
	IsCanceled *bool `json:"isCanceled,omitempty"`
	IsFree *bool `json:"isFree,omitempty"`
	IsMediaReady *bool `json:"isMediaReady,omitempty"`
	IsReportedByMe *bool `json:"isReportedByMe,omitempty"`
	IsTip *bool `json:"isTip,omitempty"`
	Media []struct {
	CanView bool `json:"canView"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CreatedAt string `json:"createdAt"`
	Duration *float64 `json:"duration,omitempty"`
	Files *struct {
	Full struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"sources"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"full"`
	Preview *struct {
	Height *float64 `json:"height,omitempty"`
	Options []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"options"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"squarePreview,omitempty"`
	Thumb *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"thumb,omitempty"`
} `json:"files,omitempty"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	HasError bool `json:"hasError"`
	Id float64 `json:"id"`
	IsReady bool `json:"isReady"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs"`
	Id float64 `json:"id"`
	IsFromGuest bool `json:"isFromGuest"`
	IsVerified bool `json:"isVerified"`
	IvStatus string `json:"ivStatus"`
	Name string `json:"name"`
	Username string `json:"username"`
	View string `json:"view"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	Type string `json:"type"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
} `json:"media"`
	MediaCount *float64 `json:"mediaCount,omitempty"`
	Previews []struct {
	CanView bool `json:"canView"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CreatedAt string `json:"createdAt"`
	Duration *float64 `json:"duration,omitempty"`
	Files *struct {
	Full struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"sources"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"full"`
	Preview *struct {
	Height *float64 `json:"height,omitempty"`
	Options []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"options"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"squarePreview,omitempty"`
	Thumb *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"thumb,omitempty"`
} `json:"files,omitempty"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	HasError bool `json:"hasError"`
	Id float64 `json:"id"`
	IsReady bool `json:"isReady"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs"`
	Id float64 `json:"id"`
	IsFromGuest bool `json:"isFromGuest"`
	IsVerified bool `json:"isVerified"`
	IvStatus string `json:"ivStatus"`
	Name string `json:"name"`
	Username string `json:"username"`
	View string `json:"view"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	Type string `json:"type"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
} `json:"previews"`
	Price *string `json:"price,omitempty"`
	PurchasedCount *float64 `json:"purchasedCount,omitempty"`
	RawText *string `json:"rawText,omitempty"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs"`
	Id float64 `json:"id"`
	IsFromGuest bool `json:"isFromGuest"`
	IsVerified bool `json:"isVerified"`
	IvStatus string `json:"ivStatus"`
	Name string `json:"name"`
	Username string `json:"username"`
	View string `json:"view"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	ResponseType *string `json:"responseType,omitempty"`
	SentCount *float64 `json:"sentCount,omitempty"`
	Template *string `json:"template,omitempty"`
	Text string `json:"text"`
	UnsendSeconds *float64 `json:"unsendSeconds,omitempty"`
	ViewedCount *float64 `json:"viewedCount,omitempty"`
} `json:"items"`
}

// GetAccessAnalyticsPostsResponse represents the response for GetAccessAnalyticsPosts
type GetAccessAnalyticsPostsResponse struct {
	CommentChart []struct {
	Count float64 `json:"count"`
	Date string `json:"date"`
} `json:"commentChart"`
	CommentCount float64 `json:"commentCount"`
	HasStats bool `json:"hasStats"`
	HasVideo *bool `json:"hasVideo,omitempty"`
	IsAvailable bool `json:"isAvailable"`
	LikeChart []struct {
	Count float64 `json:"count"`
	Date string `json:"date"`
} `json:"likeChart"`
	LikeCount float64 `json:"likeCount"`
	LookChart []struct {
	Count float64 `json:"count"`
	Date string `json:"date"`
} `json:"lookChart"`
	LookCount float64 `json:"lookCount"`
	LookDuration float64 `json:"lookDuration"`
	LookDurationAverage float64 `json:"lookDurationAverage"`
	PurchasedCount float64 `json:"purchasedCount"`
	PurchasedSumm float64 `json:"purchasedSumm"`
	PurchasesChart []struct {
	Count float64 `json:"count"`
	Date string `json:"date"`
} `json:"purchasesChart"`
	TipChart []struct {
	Count float64 `json:"count"`
	Date string `json:"date"`
} `json:"tipChart"`
	TipCount float64 `json:"tipCount"`
	TipSum float64 `json:"tipSum"`
	TipSumChart []struct {
	Count float64 `json:"count"`
	Date string `json:"date"`
} `json:"tipSumChart"`
	UniqueLookChart []struct {
	Count float64 `json:"count"`
	Date string `json:"date"`
} `json:"uniqueLookChart"`
	UniqueLookCount float64 `json:"uniqueLookCount"`
}

// GetAccessAnalyticsPostsChartResponse represents the response for GetAccessAnalyticsPostsChart
type GetAccessAnalyticsPostsChartResponse map[string]struct {
	Chart []struct {
	Count float64 `json:"count"`
	Date string `json:"date"`
} `json:"chart"`
	Delta float64 `json:"delta"`
	Gross *float64 `json:"gross,omitempty"`
	Total float64 `json:"total"`
}

// GetAccessAnalyticsPostsTopResponse represents the response for GetAccessAnalyticsPostsTop
type GetAccessAnalyticsPostsTopResponse struct {
	HasMore bool `json:"hasMore"`
	Items []struct {
	Author *struct {
	View string `json:"_view"`
	Id float64 `json:"id"`
} `json:"author,omitempty"`
	CanComment bool `json:"canComment"`
	CanDelete bool `json:"canDelete"`
	CanEdit bool `json:"canEdit"`
	CanToggleFavorite bool `json:"canToggleFavorite"`
	CanViewMedia bool `json:"canViewMedia"`
	FavoritesCount float64 `json:"favoritesCount"`
	HasVoting *bool `json:"hasVoting,omitempty"`
	Id float64 `json:"id"`
	IsFavorite bool `json:"isFavorite"`
	IsMarkdownDisabled bool `json:"isMarkdownDisabled"`
	IsMediaReady bool `json:"isMediaReady"`
	IsOpened bool `json:"isOpened"`
	Media []struct {
	CanView bool `json:"canView"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CreatedAt string `json:"createdAt"`
	Duration *float64 `json:"duration,omitempty"`
	Files *struct {
	Full struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"sources"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"full"`
	Preview *struct {
	Height *float64 `json:"height,omitempty"`
	Options []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"options"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"squarePreview,omitempty"`
	Thumb *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"thumb,omitempty"`
} `json:"files,omitempty"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	HasError bool `json:"hasError"`
	Id float64 `json:"id"`
	IsReady bool `json:"isReady"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs"`
	Id float64 `json:"id"`
	IsFromGuest bool `json:"isFromGuest"`
	IsVerified bool `json:"isVerified"`
	IvStatus string `json:"ivStatus"`
	Name string `json:"name"`
	Username string `json:"username"`
	View string `json:"view"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	Type string `json:"type"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
} `json:"media"`
	MediaCount float64 `json:"mediaCount"`
	PostedAt string `json:"postedAt"`
	PostedAtPrecise string `json:"postedAtPrecise"`
	RawText string `json:"rawText"`
	ResponseType string `json:"responseType"`
	Stats *struct {
	CommentChart []struct {
	Count float64 `json:"count"`
	Date string `json:"date"`
} `json:"commentChart"`
	CommentCount float64 `json:"commentCount"`
	HasStats bool `json:"hasStats"`
	HasVideo *bool `json:"hasVideo,omitempty"`
	IsAvailable bool `json:"isAvailable"`
	LikeChart []struct {
	Count float64 `json:"count"`
	Date string `json:"date"`
} `json:"likeChart"`
	LikeCount float64 `json:"likeCount"`
	LookChart []struct {
	Count float64 `json:"count"`
	Date string `json:"date"`
} `json:"lookChart"`
	LookCount float64 `json:"lookCount"`
	LookDuration float64 `json:"lookDuration"`
	LookDurationAverage float64 `json:"lookDurationAverage"`
	PurchasedCount float64 `json:"purchasedCount"`
	PurchasedSumm float64 `json:"purchasedSumm"`
	TipChart []struct {
	Count float64 `json:"count"`
	Date string `json:"date"`
} `json:"tipChart"`
	TipCount float64 `json:"tipCount"`
	TipSum float64 `json:"tipSum"`
	UniqueLookChart []struct {
	Count float64 `json:"count"`
	Date string `json:"date"`
} `json:"uniqueLookChart"`
	UniqueLookCount float64 `json:"uniqueLookCount"`
} `json:"stats,omitempty"`
	Text string `json:"text"`
	TipsAmount string `json:"tipsAmount"`
	Voting *struct {
	FinishedAt string `json:"finishedAt"`
	Options []struct {
	Count *float64 `json:"count,omitempty"`
	Id float64 `json:"id"`
	IsWinner *bool `json:"isWinner,omitempty"`
	Percent *float64 `json:"percent,omitempty"`
	Text *string `json:"text,omitempty"`
} `json:"options"`
	Total float64 `json:"total"`
} `json:"voting,omitempty"`
	VotingType *float64 `json:"votingType,omitempty"`
} `json:"items"`
}

// GetAccessAnalyticsPromotionsChartResponse represents the response for GetAccessAnalyticsPromotionsChart
type GetAccessAnalyticsPromotionsChartResponse map[string]struct {
	Chart []struct {
	Count float64 `json:"count"`
	Date string `json:"date"`
} `json:"chart"`
	Delta float64 `json:"delta"`
	Gross *float64 `json:"gross,omitempty"`
	Total float64 `json:"total"`
}

// GetAccessAnalyticsPromotionsTopResponse represents the response for GetAccessAnalyticsPromotionsTop
type GetAccessAnalyticsPromotionsTopResponse struct {
	HasMore bool `json:"hasMore"`
	Items []struct {
	CanClaim bool `json:"canClaim"`
	ClaimsCount float64 `json:"claimsCount"`
	CreatedAt string `json:"createdAt"`
	FinishedAt string `json:"finishedAt"`
	HasRelatedPromo bool `json:"hasRelatedPromo"`
	Id float64 `json:"id"`
	IsFinished bool `json:"isFinished"`
	Message string `json:"message"`
	Price float64 `json:"price"`
	RawMessage string `json:"rawMessage"`
	SubscribeCounts float64 `json:"subscribeCounts"`
	SubscribeDays float64 `json:"subscribeDays"`
	Type string `json:"type"`
} `json:"items"`
}

// GetAccessAnalyticsStoriesChartResponse represents the response for GetAccessAnalyticsStoriesChart
type GetAccessAnalyticsStoriesChartResponse map[string]struct {
	Chart []struct {
	Count float64 `json:"count"`
	Date string `json:"date"`
} `json:"chart"`
	Delta float64 `json:"delta"`
	Gross *float64 `json:"gross,omitempty"`
	Total float64 `json:"total"`
}

// GetAccessAnalyticsStoriesTopResponse represents the response for GetAccessAnalyticsStoriesTop
type GetAccessAnalyticsStoriesTopResponse struct {
	HasMore bool `json:"hasMore"`
	Items []struct {
	CanDelete bool `json:"canDelete"`
	CanvasHeight float64 `json:"canvasHeight"`
	CanvasWidth float64 `json:"canvasWidth"`
	CommentsCount float64 `json:"commentsCount"`
	CreatedAt string `json:"createdAt"`
	HasPost bool `json:"hasPost"`
	Id float64 `json:"id"`
	IsHighlightCover bool `json:"isHighlightCover"`
	IsLastInHighlight bool `json:"isLastInHighlight"`
	IsReady bool `json:"isReady"`
	IsWatched bool `json:"isWatched"`
	LikesCount float64 `json:"likesCount"`
	Media []struct {
	CanView bool `json:"canView"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CreatedAt string `json:"createdAt"`
	Duration *float64 `json:"duration,omitempty"`
	Files *struct {
	Full struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"sources"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"full"`
	Preview *struct {
	Height *float64 `json:"height,omitempty"`
	Options []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"options"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"squarePreview,omitempty"`
	Thumb *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"thumb,omitempty"`
} `json:"files,omitempty"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	HasError bool `json:"hasError"`
	Id float64 `json:"id"`
	IsReady bool `json:"isReady"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs"`
	Id float64 `json:"id"`
	IsFromGuest bool `json:"isFromGuest"`
	IsVerified bool `json:"isVerified"`
	IvStatus string `json:"ivStatus"`
	Name string `json:"name"`
	Username string `json:"username"`
	View string `json:"view"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	Type string `json:"type"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
} `json:"media"`
	Post *struct {
	Author *struct {
	View string `json:"_view"`
	Id float64 `json:"id"`
} `json:"author,omitempty"`
	CoordinateParams *struct {
	Angle float64 `json:"angle"`
	Height float64 `json:"height"`
	Left float64 `json:"left"`
	Scale float64 `json:"scale"`
	Top float64 `json:"top"`
	Width float64 `json:"width"`
} `json:"coordinateParams,omitempty"`
	Id float64 `json:"id"`
} `json:"post,omitempty"`
	Question *struct {
	Entity struct {
	CreatedAt string `json:"createdAt"`
	Id float64 `json:"id"`
	Text string `json:"text"`
} `json:"entity"`
	Positions struct {
	Angle float64 `json:"angle"`
	Color string `json:"color"`
	Height float64 `json:"height"`
	Left float64 `json:"left"`
	Top float64 `json:"top"`
	Width float64 `json:"width"`
	X float64 `json:"x"`
	Y float64 `json:"y"`
	ZIndex float64 `json:"zIndex"`
} `json:"positions"`
	Type string `json:"type"`
} `json:"question,omitempty"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs"`
	Id float64 `json:"id"`
	IsFromGuest bool `json:"isFromGuest"`
	IsVerified bool `json:"isVerified"`
	IvStatus string `json:"ivStatus"`
	Name string `json:"name"`
	Username string `json:"username"`
	View string `json:"view"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	Texts []struct {
	Position struct {
	Angle float64 `json:"angle"`
	Height float64 `json:"height"`
	Width float64 `json:"width"`
	X float64 `json:"x"`
	Y float64 `json:"y"`
} `json:"position"`
	Text string `json:"text"`
} `json:"texts"`
	TipsAmount string `json:"tipsAmount"`
	TipsAmountRaw float64 `json:"tipsAmountRaw"`
	TipsCount float64 `json:"tipsCount"`
	UserId float64 `json:"userId"`
	Viewers []struct {
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs"`
	CanCommentStory bool `json:"canCommentStory"`
	CanLookStory bool `json:"canLookStory"`
	HasNotViewedStory bool `json:"hasNotViewedStory"`
	Id float64 `json:"id"`
	IsStoryBlockedUser bool `json:"isStoryBlockedUser"`
	IsStoryLiked bool `json:"isStoryLiked"`
	IsVerified bool `json:"isVerified"`
	Name string `json:"name"`
	Username string `json:"username"`
} `json:"viewers"`
	ViewersCount float64 `json:"viewersCount"`
} `json:"items"`
}

// GetAccessAnalyticsStreamsChartResponse represents the response for GetAccessAnalyticsStreamsChart
type GetAccessAnalyticsStreamsChartResponse map[string]struct {
	Chart []struct {
	Count float64 `json:"count"`
	Date string `json:"date"`
} `json:"chart"`
	Delta float64 `json:"delta"`
	Gross *float64 `json:"gross,omitempty"`
	Total float64 `json:"total"`
}

// GetAccessAnalyticsStreamsTopResponse represents the response for GetAccessAnalyticsStreamsTop
type GetAccessAnalyticsStreamsTopResponse struct {
	HasMore bool `json:"hasMore"`
	Items []struct {
	Author *struct {
	View string `json:"_view"`
	Id float64 `json:"id"`
} `json:"author,omitempty"`
	CanComment bool `json:"canComment"`
	CanDelete bool `json:"canDelete"`
	CanEdit bool `json:"canEdit"`
	CanToggleFavorite bool `json:"canToggleFavorite"`
	CanViewMedia bool `json:"canViewMedia"`
	FavoritesCount float64 `json:"favoritesCount"`
	Id float64 `json:"id"`
	IsFavorite bool `json:"isFavorite"`
	IsMarkdownDisabled bool `json:"isMarkdownDisabled"`
	IsMediaReady bool `json:"isMediaReady"`
	IsOpened bool `json:"isOpened"`
	Media []struct {
	CanView bool `json:"canView"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CreatedAt string `json:"createdAt"`
	Duration *float64 `json:"duration,omitempty"`
	Files *struct {
	Full struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"sources"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"full"`
	Preview *struct {
	Height *float64 `json:"height,omitempty"`
	Options []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"options"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"squarePreview,omitempty"`
	Thumb *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"thumb,omitempty"`
} `json:"files,omitempty"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	HasError bool `json:"hasError"`
	Id float64 `json:"id"`
	IsReady bool `json:"isReady"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs"`
	Id float64 `json:"id"`
	IsFromGuest bool `json:"isFromGuest"`
	IsVerified bool `json:"isVerified"`
	IvStatus string `json:"ivStatus"`
	Name string `json:"name"`
	Username string `json:"username"`
	View string `json:"view"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	Type string `json:"type"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
} `json:"media"`
	MediaCount float64 `json:"mediaCount"`
	PostedAt string `json:"postedAt"`
	PostedAtPrecise string `json:"postedAtPrecise"`
	RawText string `json:"rawText"`
	ResponseType string `json:"responseType"`
	Stats *struct {
	CommentChart []struct {
	Count float64 `json:"count"`
	Date string `json:"date"`
} `json:"commentChart"`
	CommentCount float64 `json:"commentCount"`
	HasStats bool `json:"hasStats"`
	HasVideo *bool `json:"hasVideo,omitempty"`
	IsAvailable bool `json:"isAvailable"`
	LikeChart []struct {
	Count float64 `json:"count"`
	Date string `json:"date"`
} `json:"likeChart"`
	LikeCount float64 `json:"likeCount"`
	LookChart []struct {
	Count float64 `json:"count"`
	Date string `json:"date"`
} `json:"lookChart"`
	LookCount float64 `json:"lookCount"`
	LookDuration float64 `json:"lookDuration"`
	LookDurationAverage float64 `json:"lookDurationAverage"`
	PurchasedCount float64 `json:"purchasedCount"`
	PurchasedSumm float64 `json:"purchasedSumm"`
	TipChart []struct {
	Count float64 `json:"count"`
	Date string `json:"date"`
} `json:"tipChart"`
	TipCount float64 `json:"tipCount"`
	TipSum float64 `json:"tipSum"`
	UniqueLookChart []struct {
	Count float64 `json:"count"`
	Date string `json:"date"`
} `json:"uniqueLookChart"`
	UniqueLookCount float64 `json:"uniqueLookCount"`
} `json:"stats,omitempty"`
	Text string `json:"text"`
	TipsAmount string `json:"tipsAmount"`
} `json:"items"`
}

// GetAccessAnalyticsTrialsChartResponse represents the response for GetAccessAnalyticsTrialsChart
type GetAccessAnalyticsTrialsChartResponse map[string]struct {
	Chart []struct {
	Count float64 `json:"count"`
	Date string `json:"date"`
} `json:"chart"`
	Delta float64 `json:"delta"`
	Gross *float64 `json:"gross,omitempty"`
	Total float64 `json:"total"`
}

// GetAccessAnalyticsTrialsTopResponse represents the response for GetAccessAnalyticsTrialsTop
type GetAccessAnalyticsTrialsTopResponse struct {
	HasMore bool `json:"hasMore"`
	Items []struct {
	ClaimCounts float64 `json:"claimCounts"`
	CreatedAt string `json:"createdAt"`
	ExpiredAt string `json:"expiredAt"`
	Id float64 `json:"id"`
	IsFinished bool `json:"isFinished"`
	SubscribeCounts float64 `json:"subscribeCounts"`
	SubscribeDays float64 `json:"subscribeDays"`
	TrialLinkName string `json:"trialLinkName"`
	Url string `json:"url"`
} `json:"items"`
}

// GetAccessAnalyticsVisitorCountriesChartResponse represents the response for GetAccessAnalyticsVisitorCountriesChart
type GetAccessAnalyticsVisitorCountriesChartResponse struct {
	Chart struct {
	Duration []struct {
	Count float64 `json:"count"`
	Date string `json:"date"`
} `json:"duration"`
	Visitors []struct {
	Count float64 `json:"count"`
	Date string `json:"date"`
} `json:"visitors"`
} `json:"chart"`
	HasStats bool `json:"hasStats"`
	IsAvailable bool `json:"isAvailable"`
	Total struct {
	Current string `json:"current"`
	Delta float64 `json:"delta"`
} `json:"total"`
}

// GetAccessAnalyticsVisitorCountriesTopResponse represents the response for GetAccessAnalyticsVisitorCountriesTop
type GetAccessAnalyticsVisitorCountriesTopResponse struct {
	HasStats bool `json:"hasStats"`
	IsAvailable bool `json:"isAvailable"`
	TopCountries struct {
	HasMore bool `json:"hasMore"`
	Rows []struct {
	CountryCode string `json:"countryCode"`
	CountryName string `json:"countryName"`
	Rank float64 `json:"rank"`
	ViewsCount struct {
	Guests float64 `json:"guests"`
	Subscribers float64 `json:"subscribers"`
	Total float64 `json:"total"`
	Users float64 `json:"users"`
} `json:"viewsCount"`
} `json:"rows"`
	Totals struct {
	Guests string `json:"guests"`
	Subscribers float64 `json:"subscribers"`
	Total float64 `json:"total"`
	Users string `json:"users"`
} `json:"totals"`
} `json:"topCountries"`
}

// ListAccessChatsResponse represents the response for ListAccessChats
type ListAccessChatsResponse struct {
	HasMore bool `json:"hasMore"`
	List []struct {
	CanGoToProfile *bool `json:"canGoToProfile,omitempty"`
	CanNotSendReason *interface{} `json:"canNotSendReason,omitempty"`
	CanSendMessage *bool `json:"canSendMessage,omitempty"`
	CountPinnedMessages *float64 `json:"countPinnedMessages,omitempty"`
	HasPurchasedFeed *bool `json:"hasPurchasedFeed,omitempty"`
	HasUnreadTips *bool `json:"hasUnreadTips,omitempty"`
	IsMutedNotifications *bool `json:"isMutedNotifications,omitempty"`
	LastMessage *struct {
	CanBePinned *bool `json:"canBePinned,omitempty"`
	CanPurchase *bool `json:"canPurchase,omitempty"`
	CanPurchaseReason *string `json:"canPurchaseReason,omitempty"`
	CanReport *bool `json:"canReport,omitempty"`
	CanUnsendQueue *bool `json:"canUnsendQueue,omitempty"`
	CancelSeconds *float64 `json:"cancelSeconds,omitempty"`
	ChangedAt *string `json:"changedAt,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	FromUser interface{} `json:"fromUser"`
	GiphyId *string `json:"giphyId,omitempty"`
	Id float64 `json:"id"`
	IsCouplePeopleMedia *bool `json:"isCouplePeopleMedia,omitempty"`
	IsFree *bool `json:"isFree,omitempty"`
	IsFromQueue *bool `json:"isFromQueue,omitempty"`
	IsLiked *bool `json:"isLiked,omitempty"`
	IsMarkdownDisabled *bool `json:"isMarkdownDisabled,omitempty"`
	IsMediaReady *bool `json:"isMediaReady,omitempty"`
	IsNew *bool `json:"isNew,omitempty"`
	IsOpened *bool `json:"isOpened,omitempty"`
	IsPinned *bool `json:"isPinned,omitempty"`
	IsReportedByMe *bool `json:"isReportedByMe,omitempty"`
	IsTip *bool `json:"isTip,omitempty"`
	LockedText *bool `json:"lockedText,omitempty"`
	Media []struct {
	CanView bool `json:"canView"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CreatedAt string `json:"createdAt"`
	Duration *float64 `json:"duration,omitempty"`
	Files *struct {
	Full struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"sources"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"full"`
	Preview *struct {
	Height *float64 `json:"height,omitempty"`
	Options []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"options"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"squarePreview,omitempty"`
	Thumb *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"thumb,omitempty"`
} `json:"files,omitempty"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	HasError bool `json:"hasError"`
	Id float64 `json:"id"`
	IsReady bool `json:"isReady"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs"`
	Id float64 `json:"id"`
	IsFromGuest bool `json:"isFromGuest"`
	IsVerified bool `json:"isVerified"`
	IvStatus string `json:"ivStatus"`
	Name string `json:"name"`
	Username string `json:"username"`
	View string `json:"view"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	Type string `json:"type"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
} `json:"media"`
	MediaCount *float64 `json:"mediaCount,omitempty"`
	Previews []float64 `json:"previews"`
	QueueId *float64 `json:"queueId,omitempty"`
	ReleaseForms *interface{} `json:"releaseForms,omitempty"`
	ResponseType *string `json:"responseType,omitempty"`
	Text string `json:"text"`
	UnsendSecondsQueue *float64 `json:"unsendSecondsQueue,omitempty"`
} `json:"lastMessage,omitempty"`
	LastReadMessageId *float64 `json:"lastReadMessageId,omitempty"`
	UnreadMessagesCount *float64 `json:"unreadMessagesCount,omitempty"`
	WithUser struct {
	View string `json:"_view"`
	Id float64 `json:"id"`
} `json:"withUser"`
} `json:"list"`
	NextOffset *float64 `json:"nextOffset,omitempty"`
}

// ListAccessChatsMediaResponse represents the response for ListAccessChatsMedia
type ListAccessChatsMediaResponse struct {
	HasMore bool `json:"hasMore"`
	List []struct {
	CanBePinned *bool `json:"canBePinned,omitempty"`
	CanPurchase *bool `json:"canPurchase,omitempty"`
	CanPurchaseReason *string `json:"canPurchaseReason,omitempty"`
	CanReport *bool `json:"canReport,omitempty"`
	CanUnsendQueue *bool `json:"canUnsendQueue,omitempty"`
	CancelSeconds *float64 `json:"cancelSeconds,omitempty"`
	ChangedAt *string `json:"changedAt,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	FromUser interface{} `json:"fromUser"`
	GiphyId *string `json:"giphyId,omitempty"`
	Id float64 `json:"id"`
	IsCouplePeopleMedia *bool `json:"isCouplePeopleMedia,omitempty"`
	IsFree *bool `json:"isFree,omitempty"`
	IsFromQueue *bool `json:"isFromQueue,omitempty"`
	IsLiked *bool `json:"isLiked,omitempty"`
	IsMarkdownDisabled *bool `json:"isMarkdownDisabled,omitempty"`
	IsMediaReady *bool `json:"isMediaReady,omitempty"`
	IsNew *bool `json:"isNew,omitempty"`
	IsOpened *bool `json:"isOpened,omitempty"`
	IsPinned *bool `json:"isPinned,omitempty"`
	IsReportedByMe *bool `json:"isReportedByMe,omitempty"`
	IsTip *bool `json:"isTip,omitempty"`
	LockedText *bool `json:"lockedText,omitempty"`
	Media []struct {
	CanView bool `json:"canView"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CreatedAt string `json:"createdAt"`
	Duration *float64 `json:"duration,omitempty"`
	Files *struct {
	Full struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"sources"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"full"`
	Preview *struct {
	Height *float64 `json:"height,omitempty"`
	Options []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"options"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"squarePreview,omitempty"`
	Thumb *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"thumb,omitempty"`
} `json:"files,omitempty"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	HasError bool `json:"hasError"`
	Id float64 `json:"id"`
	IsReady bool `json:"isReady"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs"`
	Id float64 `json:"id"`
	IsFromGuest bool `json:"isFromGuest"`
	IsVerified bool `json:"isVerified"`
	IvStatus string `json:"ivStatus"`
	Name string `json:"name"`
	Username string `json:"username"`
	View string `json:"view"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	Type string `json:"type"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
} `json:"media"`
	MediaCount *float64 `json:"mediaCount,omitempty"`
	Previews []float64 `json:"previews"`
	QueueId *float64 `json:"queueId,omitempty"`
	ReleaseForms *interface{} `json:"releaseForms,omitempty"`
	ResponseType *string `json:"responseType,omitempty"`
	Text string `json:"text"`
	UnsendSecondsQueue *float64 `json:"unsendSecondsQueue,omitempty"`
} `json:"list"`
}

// ListAccessChatsMessagesResponse represents the response for ListAccessChatsMessages
type ListAccessChatsMessagesResponse struct {
	HasMore bool `json:"hasMore"`
	List []struct {
	CanBePinned *bool `json:"canBePinned,omitempty"`
	CanPurchase *bool `json:"canPurchase,omitempty"`
	CanPurchaseReason *string `json:"canPurchaseReason,omitempty"`
	CanReport *bool `json:"canReport,omitempty"`
	CanUnsendQueue *bool `json:"canUnsendQueue,omitempty"`
	CancelSeconds *float64 `json:"cancelSeconds,omitempty"`
	ChangedAt *string `json:"changedAt,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	FromUser interface{} `json:"fromUser"`
	GiphyId *string `json:"giphyId,omitempty"`
	Id float64 `json:"id"`
	IsCouplePeopleMedia *bool `json:"isCouplePeopleMedia,omitempty"`
	IsFree *bool `json:"isFree,omitempty"`
	IsFromQueue *bool `json:"isFromQueue,omitempty"`
	IsLiked *bool `json:"isLiked,omitempty"`
	IsMarkdownDisabled *bool `json:"isMarkdownDisabled,omitempty"`
	IsMediaReady *bool `json:"isMediaReady,omitempty"`
	IsNew *bool `json:"isNew,omitempty"`
	IsOpened *bool `json:"isOpened,omitempty"`
	IsPinned *bool `json:"isPinned,omitempty"`
	IsReportedByMe *bool `json:"isReportedByMe,omitempty"`
	IsTip *bool `json:"isTip,omitempty"`
	LockedText *bool `json:"lockedText,omitempty"`
	Media []struct {
	CanView bool `json:"canView"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CreatedAt string `json:"createdAt"`
	Duration *float64 `json:"duration,omitempty"`
	Files *struct {
	Full struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"sources"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"full"`
	Preview *struct {
	Height *float64 `json:"height,omitempty"`
	Options []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"options"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"squarePreview,omitempty"`
	Thumb *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"thumb,omitempty"`
} `json:"files,omitempty"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	HasError bool `json:"hasError"`
	Id float64 `json:"id"`
	IsReady bool `json:"isReady"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs"`
	Id float64 `json:"id"`
	IsFromGuest bool `json:"isFromGuest"`
	IsVerified bool `json:"isVerified"`
	IvStatus string `json:"ivStatus"`
	Name string `json:"name"`
	Username string `json:"username"`
	View string `json:"view"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	Type string `json:"type"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
} `json:"media"`
	MediaCount *float64 `json:"mediaCount,omitempty"`
	Previews []float64 `json:"previews"`
	QueueId *float64 `json:"queueId,omitempty"`
	ReleaseForms *interface{} `json:"releaseForms,omitempty"`
	ResponseType *string `json:"responseType,omitempty"`
	Text string `json:"text"`
	UnsendSecondsQueue *float64 `json:"unsendSecondsQueue,omitempty"`
} `json:"list"`
}

// CreateAccessChatsMessagesResponse represents the response for CreateAccessChatsMessages
type CreateAccessChatsMessagesResponse struct {
	CanBePinned *bool `json:"canBePinned,omitempty"`
	CanPurchase *bool `json:"canPurchase,omitempty"`
	CanPurchaseReason *string `json:"canPurchaseReason,omitempty"`
	CanReport *bool `json:"canReport,omitempty"`
	CanUnsendQueue *bool `json:"canUnsendQueue,omitempty"`
	CancelSeconds *float64 `json:"cancelSeconds,omitempty"`
	ChangedAt *string `json:"changedAt,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	FromUser interface{} `json:"fromUser"`
	GiphyId *string `json:"giphyId,omitempty"`
	Id float64 `json:"id"`
	IsCouplePeopleMedia *bool `json:"isCouplePeopleMedia,omitempty"`
	IsFree *bool `json:"isFree,omitempty"`
	IsFromQueue *bool `json:"isFromQueue,omitempty"`
	IsLiked *bool `json:"isLiked,omitempty"`
	IsMarkdownDisabled *bool `json:"isMarkdownDisabled,omitempty"`
	IsMediaReady *bool `json:"isMediaReady,omitempty"`
	IsNew *bool `json:"isNew,omitempty"`
	IsOpened *bool `json:"isOpened,omitempty"`
	IsPinned *bool `json:"isPinned,omitempty"`
	IsReportedByMe *bool `json:"isReportedByMe,omitempty"`
	IsTip *bool `json:"isTip,omitempty"`
	LockedText *bool `json:"lockedText,omitempty"`
	Media []struct {
	CanView bool `json:"canView"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CreatedAt string `json:"createdAt"`
	Duration *float64 `json:"duration,omitempty"`
	Files *struct {
	Full struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"sources"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"full"`
	Preview *struct {
	Height *float64 `json:"height,omitempty"`
	Options []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"options"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"squarePreview,omitempty"`
	Thumb *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"thumb,omitempty"`
} `json:"files,omitempty"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	HasError bool `json:"hasError"`
	Id float64 `json:"id"`
	IsReady bool `json:"isReady"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs"`
	Id float64 `json:"id"`
	IsFromGuest bool `json:"isFromGuest"`
	IsVerified bool `json:"isVerified"`
	IvStatus string `json:"ivStatus"`
	Name string `json:"name"`
	Username string `json:"username"`
	View string `json:"view"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	Type string `json:"type"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
} `json:"media"`
	MediaCount *float64 `json:"mediaCount,omitempty"`
	Previews []float64 `json:"previews"`
	QueueId *float64 `json:"queueId,omitempty"`
	ReleaseForms *interface{} `json:"releaseForms,omitempty"`
	ResponseType *string `json:"responseType,omitempty"`
	Text string `json:"text"`
	UnsendSecondsQueue *float64 `json:"unsendSecondsQueue,omitempty"`
}

// DeleteAccessChatsMessagesResponse represents the response for DeleteAccessChatsMessages
type DeleteAccessChatsMessagesResponse struct {
	Queue *struct {
	CanBePinned *bool `json:"canBePinned,omitempty"`
	CanPurchase *bool `json:"canPurchase,omitempty"`
	CanPurchaseReason *string `json:"canPurchaseReason,omitempty"`
	CanReport *bool `json:"canReport,omitempty"`
	CanSendMessageToBuyers *bool `json:"canSendMessageToBuyers,omitempty"`
	CanUnsend *bool `json:"canUnsend,omitempty"`
	CanUnsendQueue *bool `json:"canUnsendQueue,omitempty"`
	CancelSeconds *float64 `json:"cancelSeconds,omitempty"`
	ChangedAt *string `json:"changedAt,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	Date *string `json:"date,omitempty"`
	FromUser interface{} `json:"fromUser"`
	GiphyId *string `json:"giphyId,omitempty"`
	HasError *bool `json:"hasError,omitempty"`
	Id float64 `json:"id"`
	IsCanceled *bool `json:"isCanceled,omitempty"`
	IsCouplePeopleMedia *bool `json:"isCouplePeopleMedia,omitempty"`
	IsFree *bool `json:"isFree,omitempty"`
	IsFromQueue *bool `json:"isFromQueue,omitempty"`
	IsLiked *bool `json:"isLiked,omitempty"`
	IsMarkdownDisabled *bool `json:"isMarkdownDisabled,omitempty"`
	IsMediaReady *bool `json:"isMediaReady,omitempty"`
	IsNew *bool `json:"isNew,omitempty"`
	IsOpened *bool `json:"isOpened,omitempty"`
	IsPinned *bool `json:"isPinned,omitempty"`
	IsReportedByMe *bool `json:"isReportedByMe,omitempty"`
	IsTip *bool `json:"isTip,omitempty"`
	LockedText *bool `json:"lockedText,omitempty"`
	Media []struct {
	CanView bool `json:"canView"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CreatedAt string `json:"createdAt"`
	Duration *float64 `json:"duration,omitempty"`
	Files *struct {
	Full struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"sources"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"full"`
	Preview *struct {
	Height *float64 `json:"height,omitempty"`
	Options []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"options"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"squarePreview,omitempty"`
	Thumb *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"thumb,omitempty"`
} `json:"files,omitempty"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	HasError bool `json:"hasError"`
	Id float64 `json:"id"`
	IsReady bool `json:"isReady"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs"`
	Id float64 `json:"id"`
	IsFromGuest bool `json:"isFromGuest"`
	IsVerified bool `json:"isVerified"`
	IvStatus string `json:"ivStatus"`
	Name string `json:"name"`
	Username string `json:"username"`
	View string `json:"view"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	Type string `json:"type"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
} `json:"media"`
	MediaCount *float64 `json:"mediaCount,omitempty"`
	MediaTypes *struct {
	Audio *float64 `json:"audio,omitempty"`
	Gif *float64 `json:"gif,omitempty"`
	Photo *float64 `json:"photo,omitempty"`
	Video *float64 `json:"video,omitempty"`
} `json:"mediaTypes,omitempty"`
	Previews []float64 `json:"previews"`
	Price *string `json:"price,omitempty"`
	PurchasedCount *float64 `json:"purchasedCount,omitempty"`
	QueueId *float64 `json:"queueId,omitempty"`
	ReleaseForms *interface{} `json:"releaseForms,omitempty"`
	ResponseType *string `json:"responseType,omitempty"`
	SentCount *float64 `json:"sentCount,omitempty"`
	Text string `json:"text"`
	TextCropped *string `json:"textCropped,omitempty"`
	UnsendSeconds *float64 `json:"unsendSeconds,omitempty"`
	UnsendSecondsQueue *float64 `json:"unsendSecondsQueue,omitempty"`
	ViewedCount *float64 `json:"viewedCount,omitempty"`
} `json:"queue,omitempty"`
	Success bool `json:"success"`
}

// ListAccessEarningsChargebacksResponse represents the response for ListAccessEarningsChargebacks
type ListAccessEarningsChargebacksResponse struct {
	HasMore bool `json:"hasMore"`
	List []struct {
	CreatedAt string `json:"createdAt"`
	Id float64 `json:"id"`
	Payment struct {
	Amounts struct {
	Fee float64 `json:"fee"`
	Gross float64 `json:"gross"`
	Net float64 `json:"net"`
	Tax float64 `json:"tax"`
	Vat float64 `json:"vat"`
} `json:"amounts"`
	CreatedAt string `json:"createdAt"`
	Currency string `json:"currency"`
	Description string `json:"description"`
	Id string `json:"id"`
	Status string `json:"status"`
	User struct {
	Avatar string `json:"avatar"`
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
} `json:"user"`
} `json:"payment"`
	PaymentType string `json:"paymentType"`
} `json:"list"`
	NextMarker float64 `json:"nextMarker"`
}

// GetAccessEarningsChartResponse represents the response for GetAccessEarningsChart
type GetAccessEarningsChartResponse struct {
	Chart []struct {
	Count float64 `json:"count"`
	Date string `json:"date"`
} `json:"chart"`
	Delta *float64 `json:"delta,omitempty"`
	Gross *float64 `json:"gross,omitempty"`
	Total float64 `json:"total"`
}

// ListAccessEarningsTransactionsResponse represents the response for ListAccessEarningsTransactions
type ListAccessEarningsTransactionsResponse struct {
	HasMore bool `json:"hasMore"`
	List []struct {
	Amounts struct {
	Fee float64 `json:"fee"`
	Gross float64 `json:"gross"`
	Net float64 `json:"net"`
	Tax float64 `json:"tax"`
	Vat float64 `json:"vat"`
} `json:"amounts"`
	CreatedAt string `json:"createdAt"`
	Currency string `json:"currency"`
	Description string `json:"description"`
	Id string `json:"id"`
	PayoutPendingDays float64 `json:"payoutPendingDays"`
	Status string `json:"status"`
	Type string `json:"type"`
	User struct {
	Avatar string `json:"avatar"`
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
} `json:"user"`
} `json:"list"`
	NextMarker float64 `json:"nextMarker"`
}

// ListAccessMassMessagesResponse represents the response for ListAccessMassMessages
type ListAccessMassMessagesResponse struct {
	HasMore bool `json:"hasMore"`
	List []struct {
	CanBePinned *bool `json:"canBePinned,omitempty"`
	CanPurchase *bool `json:"canPurchase,omitempty"`
	CanPurchaseReason *string `json:"canPurchaseReason,omitempty"`
	CanReport *bool `json:"canReport,omitempty"`
	CanSendMessageToBuyers *bool `json:"canSendMessageToBuyers,omitempty"`
	CanUnsend *bool `json:"canUnsend,omitempty"`
	CanUnsendQueue *bool `json:"canUnsendQueue,omitempty"`
	CancelSeconds *float64 `json:"cancelSeconds,omitempty"`
	ChangedAt *string `json:"changedAt,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	Date *string `json:"date,omitempty"`
	FromUser interface{} `json:"fromUser"`
	GiphyId *string `json:"giphyId,omitempty"`
	HasError *bool `json:"hasError,omitempty"`
	Id float64 `json:"id"`
	IsCanceled *bool `json:"isCanceled,omitempty"`
	IsCouplePeopleMedia *bool `json:"isCouplePeopleMedia,omitempty"`
	IsFree *bool `json:"isFree,omitempty"`
	IsFromQueue *bool `json:"isFromQueue,omitempty"`
	IsLiked *bool `json:"isLiked,omitempty"`
	IsMarkdownDisabled *bool `json:"isMarkdownDisabled,omitempty"`
	IsMediaReady *bool `json:"isMediaReady,omitempty"`
	IsNew *bool `json:"isNew,omitempty"`
	IsOpened *bool `json:"isOpened,omitempty"`
	IsPinned *bool `json:"isPinned,omitempty"`
	IsReportedByMe *bool `json:"isReportedByMe,omitempty"`
	IsTip *bool `json:"isTip,omitempty"`
	LockedText *bool `json:"lockedText,omitempty"`
	Media []struct {
	CanView bool `json:"canView"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CreatedAt string `json:"createdAt"`
	Duration *float64 `json:"duration,omitempty"`
	Files *struct {
	Full struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"sources"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"full"`
	Preview *struct {
	Height *float64 `json:"height,omitempty"`
	Options []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"options"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"squarePreview,omitempty"`
	Thumb *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"thumb,omitempty"`
} `json:"files,omitempty"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	HasError bool `json:"hasError"`
	Id float64 `json:"id"`
	IsReady bool `json:"isReady"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs"`
	Id float64 `json:"id"`
	IsFromGuest bool `json:"isFromGuest"`
	IsVerified bool `json:"isVerified"`
	IvStatus string `json:"ivStatus"`
	Name string `json:"name"`
	Username string `json:"username"`
	View string `json:"view"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	Type string `json:"type"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
} `json:"media"`
	MediaCount *float64 `json:"mediaCount,omitempty"`
	MediaTypes *struct {
	Audio *float64 `json:"audio,omitempty"`
	Gif *float64 `json:"gif,omitempty"`
	Photo *float64 `json:"photo,omitempty"`
	Video *float64 `json:"video,omitempty"`
} `json:"mediaTypes,omitempty"`
	Previews []float64 `json:"previews"`
	Price *string `json:"price,omitempty"`
	PurchasedCount *float64 `json:"purchasedCount,omitempty"`
	QueueId *float64 `json:"queueId,omitempty"`
	ReleaseForms *interface{} `json:"releaseForms,omitempty"`
	ResponseType *string `json:"responseType,omitempty"`
	SentCount *float64 `json:"sentCount,omitempty"`
	Text string `json:"text"`
	TextCropped *string `json:"textCropped,omitempty"`
	UnsendSeconds *float64 `json:"unsendSeconds,omitempty"`
	UnsendSecondsQueue *float64 `json:"unsendSecondsQueue,omitempty"`
	ViewedCount *float64 `json:"viewedCount,omitempty"`
} `json:"list"`
}

// CreateAccessMassMessagesResponse represents the response for CreateAccessMassMessages
type CreateAccessMassMessagesResponse struct {
	CanUnsend *bool `json:"canUnsend,omitempty"`
	Date string `json:"date"`
	HasError *bool `json:"hasError,omitempty"`
	Id float64 `json:"id"`
	IsCanceled *bool `json:"isCanceled,omitempty"`
	IsCouplePeopleMedia *bool `json:"isCouplePeopleMedia,omitempty"`
	IsDone *bool `json:"isDone,omitempty"`
	IsReady *bool `json:"isReady,omitempty"`
	Pending *float64 `json:"pending,omitempty"`
	Total *float64 `json:"total,omitempty"`
	UnsendSeconds *float64 `json:"unsendSeconds,omitempty"`
}

// GetAccessMassMessagesResponse represents the response for GetAccessMassMessages
type GetAccessMassMessagesResponse struct {
	CanBePinned *bool `json:"canBePinned,omitempty"`
	CanPurchase *bool `json:"canPurchase,omitempty"`
	CanPurchaseReason *string `json:"canPurchaseReason,omitempty"`
	CanReport *bool `json:"canReport,omitempty"`
	CanUnsendQueue *bool `json:"canUnsendQueue,omitempty"`
	CancelSeconds *float64 `json:"cancelSeconds,omitempty"`
	ChangedAt *string `json:"changedAt,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	FromUser interface{} `json:"fromUser"`
	GiphyId *string `json:"giphyId,omitempty"`
	Id float64 `json:"id"`
	IsCouplePeopleMedia *bool `json:"isCouplePeopleMedia,omitempty"`
	IsFree *bool `json:"isFree,omitempty"`
	IsFromQueue *bool `json:"isFromQueue,omitempty"`
	IsLiked *bool `json:"isLiked,omitempty"`
	IsMarkdownDisabled *bool `json:"isMarkdownDisabled,omitempty"`
	IsMediaReady *bool `json:"isMediaReady,omitempty"`
	IsNew *bool `json:"isNew,omitempty"`
	IsOpened *bool `json:"isOpened,omitempty"`
	IsPinned *bool `json:"isPinned,omitempty"`
	IsReportedByMe *bool `json:"isReportedByMe,omitempty"`
	IsTip *bool `json:"isTip,omitempty"`
	LockedText *bool `json:"lockedText,omitempty"`
	Media []struct {
	CanView bool `json:"canView"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CreatedAt string `json:"createdAt"`
	Duration *float64 `json:"duration,omitempty"`
	Files *struct {
	Full struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"sources"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"full"`
	Preview *struct {
	Height *float64 `json:"height,omitempty"`
	Options []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"options"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"squarePreview,omitempty"`
	Thumb *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"thumb,omitempty"`
} `json:"files,omitempty"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	HasError bool `json:"hasError"`
	Id float64 `json:"id"`
	IsReady bool `json:"isReady"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs"`
	Id float64 `json:"id"`
	IsFromGuest bool `json:"isFromGuest"`
	IsVerified bool `json:"isVerified"`
	IvStatus string `json:"ivStatus"`
	Name string `json:"name"`
	Username string `json:"username"`
	View string `json:"view"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	Type string `json:"type"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
} `json:"media"`
	MediaCount *float64 `json:"mediaCount,omitempty"`
	Previews []float64 `json:"previews"`
	QueueId *float64 `json:"queueId,omitempty"`
	ReleaseForms *interface{} `json:"releaseForms,omitempty"`
	ResponseType *string `json:"responseType,omitempty"`
	Text string `json:"text"`
	UnsendSecondsQueue *float64 `json:"unsendSecondsQueue,omitempty"`
}

// ReplaceAccessMassMessagesResponse represents the response for ReplaceAccessMassMessages
type ReplaceAccessMassMessagesResponse struct {
	CanUnsend *bool `json:"canUnsend,omitempty"`
	Date string `json:"date"`
	HasError *bool `json:"hasError,omitempty"`
	Id float64 `json:"id"`
	IsCanceled *bool `json:"isCanceled,omitempty"`
	IsCouplePeopleMedia *bool `json:"isCouplePeopleMedia,omitempty"`
	IsDone *bool `json:"isDone,omitempty"`
	IsReady *bool `json:"isReady,omitempty"`
	Pending *float64 `json:"pending,omitempty"`
	Total *float64 `json:"total,omitempty"`
	UnsendSeconds *float64 `json:"unsendSeconds,omitempty"`
}

// DeleteAccessMassMessagesResponse represents the response for DeleteAccessMassMessages
type DeleteAccessMassMessagesResponse struct {
	Queue *struct {
	CanBePinned *bool `json:"canBePinned,omitempty"`
	CanPurchase *bool `json:"canPurchase,omitempty"`
	CanPurchaseReason *string `json:"canPurchaseReason,omitempty"`
	CanReport *bool `json:"canReport,omitempty"`
	CanSendMessageToBuyers *bool `json:"canSendMessageToBuyers,omitempty"`
	CanUnsend *bool `json:"canUnsend,omitempty"`
	CanUnsendQueue *bool `json:"canUnsendQueue,omitempty"`
	CancelSeconds *float64 `json:"cancelSeconds,omitempty"`
	ChangedAt *string `json:"changedAt,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	Date *string `json:"date,omitempty"`
	FromUser interface{} `json:"fromUser"`
	GiphyId *string `json:"giphyId,omitempty"`
	HasError *bool `json:"hasError,omitempty"`
	Id float64 `json:"id"`
	IsCanceled *bool `json:"isCanceled,omitempty"`
	IsCouplePeopleMedia *bool `json:"isCouplePeopleMedia,omitempty"`
	IsFree *bool `json:"isFree,omitempty"`
	IsFromQueue *bool `json:"isFromQueue,omitempty"`
	IsLiked *bool `json:"isLiked,omitempty"`
	IsMarkdownDisabled *bool `json:"isMarkdownDisabled,omitempty"`
	IsMediaReady *bool `json:"isMediaReady,omitempty"`
	IsNew *bool `json:"isNew,omitempty"`
	IsOpened *bool `json:"isOpened,omitempty"`
	IsPinned *bool `json:"isPinned,omitempty"`
	IsReportedByMe *bool `json:"isReportedByMe,omitempty"`
	IsTip *bool `json:"isTip,omitempty"`
	LockedText *bool `json:"lockedText,omitempty"`
	Media []struct {
	CanView bool `json:"canView"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CreatedAt string `json:"createdAt"`
	Duration *float64 `json:"duration,omitempty"`
	Files *struct {
	Full struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"sources"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"full"`
	Preview *struct {
	Height *float64 `json:"height,omitempty"`
	Options []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"options"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"squarePreview,omitempty"`
	Thumb *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"thumb,omitempty"`
} `json:"files,omitempty"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	HasError bool `json:"hasError"`
	Id float64 `json:"id"`
	IsReady bool `json:"isReady"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs"`
	Id float64 `json:"id"`
	IsFromGuest bool `json:"isFromGuest"`
	IsVerified bool `json:"isVerified"`
	IvStatus string `json:"ivStatus"`
	Name string `json:"name"`
	Username string `json:"username"`
	View string `json:"view"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	Type string `json:"type"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
} `json:"media"`
	MediaCount *float64 `json:"mediaCount,omitempty"`
	MediaTypes *struct {
	Audio *float64 `json:"audio,omitempty"`
	Gif *float64 `json:"gif,omitempty"`
	Photo *float64 `json:"photo,omitempty"`
	Video *float64 `json:"video,omitempty"`
} `json:"mediaTypes,omitempty"`
	Previews []float64 `json:"previews"`
	Price *string `json:"price,omitempty"`
	PurchasedCount *float64 `json:"purchasedCount,omitempty"`
	QueueId *float64 `json:"queueId,omitempty"`
	ReleaseForms *interface{} `json:"releaseForms,omitempty"`
	ResponseType *string `json:"responseType,omitempty"`
	SentCount *float64 `json:"sentCount,omitempty"`
	Text string `json:"text"`
	TextCropped *string `json:"textCropped,omitempty"`
	UnsendSeconds *float64 `json:"unsendSeconds,omitempty"`
	UnsendSecondsQueue *float64 `json:"unsendSecondsQueue,omitempty"`
	ViewedCount *float64 `json:"viewedCount,omitempty"`
} `json:"queue,omitempty"`
	Success bool `json:"success"`
}

// ListAccessPostsResponse represents the response for ListAccessPosts
type ListAccessPostsResponse struct {
	Counters struct {
	ArchivedPostsCount float64 `json:"archivedPostsCount"`
	AudiosCount float64 `json:"audiosCount"`
	MediasCount float64 `json:"mediasCount"`
	PhotosCount float64 `json:"photosCount"`
	PostsCount float64 `json:"postsCount"`
	PrivateArchivedPostsCount float64 `json:"privateArchivedPostsCount"`
	StreamsCount float64 `json:"streamsCount"`
	VideosCount float64 `json:"videosCount"`
} `json:"counters"`
	HasMore bool `json:"hasMore"`
	HeadMarker string `json:"headMarker"`
	List []struct {
	Author *struct {
	View string `json:"_view"`
	Id float64 `json:"id"`
} `json:"author,omitempty"`
	CanComment bool `json:"canComment"`
	CanDelete bool `json:"canDelete"`
	CanEdit bool `json:"canEdit"`
	CanToggleFavorite bool `json:"canToggleFavorite"`
	CanViewMedia bool `json:"canViewMedia"`
	FavoritesCount float64 `json:"favoritesCount"`
	Id float64 `json:"id"`
	IsFavorite bool `json:"isFavorite"`
	IsMarkdownDisabled bool `json:"isMarkdownDisabled"`
	IsMediaReady bool `json:"isMediaReady"`
	IsOpened bool `json:"isOpened"`
	Media []struct {
	CanView bool `json:"canView"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CreatedAt string `json:"createdAt"`
	Duration *float64 `json:"duration,omitempty"`
	Files *struct {
	Full struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"sources"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"full"`
	Preview *struct {
	Height *float64 `json:"height,omitempty"`
	Options []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"options"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"squarePreview,omitempty"`
	Thumb *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"thumb,omitempty"`
} `json:"files,omitempty"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	HasError bool `json:"hasError"`
	Id float64 `json:"id"`
	IsReady bool `json:"isReady"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs"`
	Id float64 `json:"id"`
	IsFromGuest bool `json:"isFromGuest"`
	IsVerified bool `json:"isVerified"`
	IvStatus string `json:"ivStatus"`
	Name string `json:"name"`
	Username string `json:"username"`
	View string `json:"view"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	Type string `json:"type"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
} `json:"media"`
	MediaCount float64 `json:"mediaCount"`
	PostedAt string `json:"postedAt"`
	PostedAtPrecise string `json:"postedAtPrecise"`
	RawText string `json:"rawText"`
	ResponseType string `json:"responseType"`
	Text string `json:"text"`
	TipsAmount string `json:"tipsAmount"`
} `json:"list"`
	TailMarker string `json:"tailMarker"`
}

// CreateAccessPostsResponse represents the response for CreateAccessPosts
type CreateAccessPostsResponse struct {
	Author *struct {
	View string `json:"_view"`
	Id float64 `json:"id"`
} `json:"author,omitempty"`
	CanComment bool `json:"canComment"`
	CanDelete bool `json:"canDelete"`
	CanEdit bool `json:"canEdit"`
	CanToggleFavorite bool `json:"canToggleFavorite"`
	CanViewMedia bool `json:"canViewMedia"`
	FavoritesCount float64 `json:"favoritesCount"`
	Id float64 `json:"id"`
	IsFavorite bool `json:"isFavorite"`
	IsMarkdownDisabled bool `json:"isMarkdownDisabled"`
	IsMediaReady bool `json:"isMediaReady"`
	IsOpened bool `json:"isOpened"`
	Media []struct {
	CanView bool `json:"canView"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CreatedAt string `json:"createdAt"`
	Duration *float64 `json:"duration,omitempty"`
	Files *struct {
	Full struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"sources"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"full"`
	Preview *struct {
	Height *float64 `json:"height,omitempty"`
	Options []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"options"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"squarePreview,omitempty"`
	Thumb *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"thumb,omitempty"`
} `json:"files,omitempty"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	HasError bool `json:"hasError"`
	Id float64 `json:"id"`
	IsReady bool `json:"isReady"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs"`
	Id float64 `json:"id"`
	IsFromGuest bool `json:"isFromGuest"`
	IsVerified bool `json:"isVerified"`
	IvStatus string `json:"ivStatus"`
	Name string `json:"name"`
	Username string `json:"username"`
	View string `json:"view"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	Type string `json:"type"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
} `json:"media"`
	MediaCount float64 `json:"mediaCount"`
	PostedAt string `json:"postedAt"`
	PostedAtPrecise string `json:"postedAtPrecise"`
	RawText string `json:"rawText"`
	ResponseType string `json:"responseType"`
	Text string `json:"text"`
	TipsAmount string `json:"tipsAmount"`
}

// ReplaceAccessPostsResponse represents the response for ReplaceAccessPosts
type ReplaceAccessPostsResponse struct {
	Author *struct {
	View string `json:"_view"`
	Id float64 `json:"id"`
} `json:"author,omitempty"`
	CanComment bool `json:"canComment"`
	CanDelete bool `json:"canDelete"`
	CanEdit bool `json:"canEdit"`
	CanToggleFavorite bool `json:"canToggleFavorite"`
	CanViewMedia bool `json:"canViewMedia"`
	FavoritesCount float64 `json:"favoritesCount"`
	Id float64 `json:"id"`
	IsFavorite bool `json:"isFavorite"`
	IsMarkdownDisabled bool `json:"isMarkdownDisabled"`
	IsMediaReady bool `json:"isMediaReady"`
	IsOpened bool `json:"isOpened"`
	Media []struct {
	CanView bool `json:"canView"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CreatedAt string `json:"createdAt"`
	Duration *float64 `json:"duration,omitempty"`
	Files *struct {
	Full struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"sources"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"full"`
	Preview *struct {
	Height *float64 `json:"height,omitempty"`
	Options []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"options"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"squarePreview,omitempty"`
	Thumb *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"thumb,omitempty"`
} `json:"files,omitempty"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	HasError bool `json:"hasError"`
	Id float64 `json:"id"`
	IsReady bool `json:"isReady"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs"`
	Id float64 `json:"id"`
	IsFromGuest bool `json:"isFromGuest"`
	IsVerified bool `json:"isVerified"`
	IvStatus string `json:"ivStatus"`
	Name string `json:"name"`
	Username string `json:"username"`
	View string `json:"view"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	Type string `json:"type"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
} `json:"media"`
	MediaCount float64 `json:"mediaCount"`
	PostedAt string `json:"postedAt"`
	PostedAtPrecise string `json:"postedAtPrecise"`
	RawText string `json:"rawText"`
	ResponseType string `json:"responseType"`
	Text string `json:"text"`
	TipsAmount string `json:"tipsAmount"`
}

// DeleteAccessPostsResponse represents the response for DeleteAccessPosts
type DeleteAccessPostsResponse map[string]interface{}

// ReplaceAccessPromotionsResponse represents the response for ReplaceAccessPromotions
type ReplaceAccessPromotionsResponse struct {
	CanClaim bool `json:"canClaim"`
	ClaimsCount float64 `json:"claimsCount"`
	CreatedAt string `json:"createdAt"`
	FinishedAt string `json:"finishedAt"`
	HasRelatedPromo bool `json:"hasRelatedPromo"`
	Id float64 `json:"id"`
	IsFinished bool `json:"isFinished"`
	Message string `json:"message"`
	Price float64 `json:"price"`
	RawMessage string `json:"rawMessage"`
	SubscribeCounts float64 `json:"subscribeCounts"`
	SubscribeDays float64 `json:"subscribeDays"`
	Type string `json:"type"`
}

// DeleteAccessPromotionsResponse represents the response for DeleteAccessPromotions
type DeleteAccessPromotionsResponse struct {
	Success bool `json:"success"`
}

// CreateAccessPromotionsStopResponse represents the response for CreateAccessPromotionsStop
type CreateAccessPromotionsStopResponse struct {
	Success bool `json:"success"`
}

// CreateAccessPromotionsBundlesResponse represents the response for CreateAccessPromotionsBundles
type CreateAccessPromotionsBundlesResponse struct {
	CanBuy bool `json:"canBuy"`
	Discount float64 `json:"discount"`
	Duration float64 `json:"duration"`
	Id float64 `json:"id"`
	Price float64 `json:"price"`
}

// GetAccessPromotionsBundlesResponse represents the response for GetAccessPromotionsBundles
type GetAccessPromotionsBundlesResponse struct {
	CanBuy bool `json:"canBuy"`
	Discount float64 `json:"discount"`
	Duration float64 `json:"duration"`
	Id float64 `json:"id"`
	Price float64 `json:"price"`
}

// ReplaceAccessPromotionsBundlesResponse represents the response for ReplaceAccessPromotionsBundles
type ReplaceAccessPromotionsBundlesResponse struct {
	CanBuy bool `json:"canBuy"`
	Discount float64 `json:"discount"`
	Duration float64 `json:"duration"`
	Id float64 `json:"id"`
	Price float64 `json:"price"`
}

// DeleteAccessPromotionsBundlesResponse represents the response for DeleteAccessPromotionsBundles
type DeleteAccessPromotionsBundlesResponse struct {
	CanBuy bool `json:"canBuy"`
	Discount float64 `json:"discount"`
	Duration float64 `json:"duration"`
	Id float64 `json:"id"`
	Price float64 `json:"price"`
}

// ListAccessPromotionsTrackingLinksResponse represents the response for ListAccessPromotionsTrackingLinks
type ListAccessPromotionsTrackingLinksResponse struct {
	HasMore bool `json:"hasMore"`
	List []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
} `json:"list"`
}

// GetAccessPromotionsTrackingLinksResponse represents the response for GetAccessPromotionsTrackingLinks
type GetAccessPromotionsTrackingLinksResponse struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
}

// ReplaceAccessPromotionsTrackingLinksResponse represents the response for ReplaceAccessPromotionsTrackingLinks
type ReplaceAccessPromotionsTrackingLinksResponse struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
}

// DeleteAccessPromotionsTrackingLinksResponse represents the response for DeleteAccessPromotionsTrackingLinks
type DeleteAccessPromotionsTrackingLinksResponse struct {
	Success bool `json:"success"`
}

// ListAccessPromotionsTrackingLinksClaimersResponse represents the response for ListAccessPromotionsTrackingLinksClaimers
type ListAccessPromotionsTrackingLinksClaimersResponse struct {
	HasMore *bool `json:"hasMore,omitempty"`
	List []interface{} `json:"list"`
}

// CreateAccessPromotionsTrackingLinksShareAccessResponse represents the response for CreateAccessPromotionsTrackingLinksShareAccess
type CreateAccessPromotionsTrackingLinksShareAccessResponse struct {
	Success bool `json:"success"`
}

// DeleteAccessPromotionsTrackingLinksShareAccessResponse represents the response for DeleteAccessPromotionsTrackingLinksShareAccess
type DeleteAccessPromotionsTrackingLinksShareAccessResponse struct {
	Success bool `json:"success"`
}

// CreateAccessPromotionsTrialLinksResponse represents the response for CreateAccessPromotionsTrialLinks
type CreateAccessPromotionsTrialLinksResponse struct {
	ClaimCounts float64 `json:"claimCounts"`
	ClicksCounts float64 `json:"clicksCounts"`
	CreatedAt string `json:"createdAt"`
	ExpiredAt string `json:"expiredAt"`
	Id float64 `json:"id"`
	IsFinished bool `json:"isFinished"`
	SharedWith *string `json:"sharedWith,omitempty"`
	SubscribeCounts float64 `json:"subscribeCounts"`
	SubscribeDays float64 `json:"subscribeDays"`
	TrialLinkName string `json:"trialLinkName"`
	Url string `json:"url"`
	User *interface{} `json:"user,omitempty"`
}

// GetAccessPromotionsTrialLinksResponse represents the response for GetAccessPromotionsTrialLinks
type GetAccessPromotionsTrialLinksResponse struct {
	ClaimCounts float64 `json:"claimCounts"`
	ClicksCounts float64 `json:"clicksCounts"`
	CreatedAt string `json:"createdAt"`
	ExpiredAt string `json:"expiredAt"`
	Id float64 `json:"id"`
	IsFinished bool `json:"isFinished"`
	SharedWith *string `json:"sharedWith,omitempty"`
	SubscribeCounts float64 `json:"subscribeCounts"`
	SubscribeDays float64 `json:"subscribeDays"`
	TrialLinkName string `json:"trialLinkName"`
	Url string `json:"url"`
	User *interface{} `json:"user,omitempty"`
}

// ReplaceAccessPromotionsTrialLinksResponse represents the response for ReplaceAccessPromotionsTrialLinks
type ReplaceAccessPromotionsTrialLinksResponse struct {
	Success bool `json:"success"`
}

// DeleteAccessPromotionsTrialLinksResponse represents the response for DeleteAccessPromotionsTrialLinks
type DeleteAccessPromotionsTrialLinksResponse struct {
	Success bool `json:"success"`
}

// CreateAccessPromotionsTrialLinksShareAccessResponse represents the response for CreateAccessPromotionsTrialLinksShareAccess
type CreateAccessPromotionsTrialLinksShareAccessResponse struct {
	Success bool `json:"success"`
}

// DeleteAccessPromotionsTrialLinksShareAccessResponse represents the response for DeleteAccessPromotionsTrialLinksShareAccess
type DeleteAccessPromotionsTrialLinksShareAccessResponse struct {
	Success bool `json:"success"`
}

// UpdateAccessSelfResponse represents the response for UpdateAccessSelf
type UpdateAccessSelfResponse struct {
	About *string `json:"about,omitempty"`
	ArchivedPostsCount *float64 `json:"archivedPostsCount,omitempty"`
	AudiosCount *float64 `json:"audiosCount,omitempty"`
	Avatar string `json:"avatar"`
	AvatarThumbs *struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs,omitempty"`
	CanAddSubscriber *bool `json:"canAddSubscriber,omitempty"`
	CanChat *bool `json:"canChat,omitempty"`
	CanCommentStory *bool `json:"canCommentStory,omitempty"`
	CanCreatePromotion *bool `json:"canCreatePromotion,omitempty"`
	CanCreateTrial *bool `json:"canCreateTrial,omitempty"`
	CanLookStory *bool `json:"canLookStory,omitempty"`
	CanPayInternal *bool `json:"canPayInternal,omitempty"`
	CanReceiveChatMessage *bool `json:"canReceiveChatMessage,omitempty"`
	CanReport *bool `json:"canReport,omitempty"`
	CanRestrict *bool `json:"canRestrict,omitempty"`
	CanTrialSend *bool `json:"canTrialSend,omitempty"`
	CanUnsubscribe *bool `json:"canUnsubscribe,omitempty"`
	CurrentSubscribePrice *float64 `json:"currentSubscribePrice,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	FavoritedCount *float64 `json:"favoritedCount,omitempty"`
	FavoritesCount *float64 `json:"favoritesCount,omitempty"`
	HasLabels *bool `json:"hasLabels,omitempty"`
	HasNotViewedStory *bool `json:"hasNotViewedStory,omitempty"`
	HasPinnedPosts *bool `json:"hasPinnedPosts,omitempty"`
	HasScheduledStream *bool `json:"hasScheduledStream,omitempty"`
	HasStories *bool `json:"hasStories,omitempty"`
	HasStream *bool `json:"hasStream,omitempty"`
	Header *string `json:"header,omitempty"`
	HeaderSize *struct {
	Height float64 `json:"height"`
	Width float64 `json:"width"`
} `json:"headerSize,omitempty"`
	HeaderThumbs *struct {
	W480 string `json:"w480"`
	W760 string `json:"w760"`
} `json:"headerThumbs,omitempty"`
	Id float64 `json:"id"`
	IsAdultContent *bool `json:"isAdultContent,omitempty"`
	IsBlocked *bool `json:"isBlocked,omitempty"`
	IsFriend *bool `json:"isFriend,omitempty"`
	IsMarkdownDisabledForAbout *bool `json:"isMarkdownDisabledForAbout,omitempty"`
	IsPendingAutoprolong *bool `json:"isPendingAutoprolong,omitempty"`
	IsPerformer *bool `json:"isPerformer,omitempty"`
	IsPrivateRestriction *bool `json:"isPrivateRestriction,omitempty"`
	IsRealPerformer *bool `json:"isRealPerformer,omitempty"`
	IsReferrerAllowed *bool `json:"isReferrerAllowed,omitempty"`
	IsRestricted *bool `json:"isRestricted,omitempty"`
	IsVerified bool `json:"isVerified"`
	JoinDate *string `json:"joinDate,omitempty"`
	LastSeen *string `json:"lastSeen,omitempty"`
	Lists []struct {
	Id interface{} `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
} `json:"lists"`
	Location *string `json:"location,omitempty"`
	MediasCount *float64 `json:"mediasCount,omitempty"`
	Name string `json:"name"`
	Notice *string `json:"notice,omitempty"`
	PhotosCount *float64 `json:"photosCount,omitempty"`
	PostsCount *float64 `json:"postsCount,omitempty"`
	PrivateArchivedPostsCount *float64 `json:"privateArchivedPostsCount,omitempty"`
	ShowMediaCount *bool `json:"showMediaCount,omitempty"`
	ShowPostsInFeed *bool `json:"showPostsInFeed,omitempty"`
	ShowSubscribersCount *bool `json:"showSubscribersCount,omitempty"`
	SubscribePrice *float64 `json:"subscribePrice,omitempty"`
	SubscribedBy *bool `json:"subscribedBy,omitempty"`
	SubscribedByAutoprolong *bool `json:"subscribedByAutoprolong,omitempty"`
	SubscribedByData *struct {
	DiscountFinishedAt string `json:"discountFinishedAt"`
	DiscountPercent float64 `json:"discountPercent"`
	DiscountPeriod float64 `json:"discountPeriod"`
	DiscountStartedAt string `json:"discountStartedAt"`
	Duration string `json:"duration"`
	ExpiredAt string `json:"expiredAt"`
	HasActivePaidSubscriptions bool `json:"hasActivePaidSubscriptions"`
	IsMuted bool `json:"isMuted"`
	NewPrice float64 `json:"newPrice"`
	Price float64 `json:"price"`
	RegularPrice float64 `json:"regularPrice"`
	RenewedAt string `json:"renewedAt"`
	ShowPostsInFeed bool `json:"showPostsInFeed"`
	Status string `json:"status"`
	SubscribeAt string `json:"subscribeAt"`
	SubscribePrice float64 `json:"subscribePrice"`
	Subscribes []struct {
	Action string `json:"action"`
	CancelDate string `json:"cancelDate"`
	Date string `json:"date"`
	Discount float64 `json:"discount"`
	Duration float64 `json:"duration"`
	EarningId float64 `json:"earningId"`
	ExpireDate string `json:"expireDate"`
	Id float64 `json:"id"`
	IsCurrent bool `json:"isCurrent"`
	OfferEnd string `json:"offerEnd"`
	OfferStart string `json:"offerStart"`
	Price float64 `json:"price"`
	RegularPrice float64 `json:"regularPrice"`
	StartDate string `json:"startDate"`
	SubscriberId float64 `json:"subscriberId"`
	Type string `json:"type"`
	UserId float64 `json:"userId"`
} `json:"subscribes"`
	UnsubscribeReason string `json:"unsubscribeReason"`
} `json:"subscribedByData,omitempty"`
	SubscribedByExpire *bool `json:"subscribedByExpire,omitempty"`
	SubscribedByExpireDate *string `json:"subscribedByExpireDate,omitempty"`
	SubscribedIsExpiredNow *bool `json:"subscribedIsExpiredNow,omitempty"`
	SubscribedOn *bool `json:"subscribedOn,omitempty"`
	SubscribedOnData *struct {
	DiscountFinishedAt string `json:"discountFinishedAt"`
	DiscountPercent float64 `json:"discountPercent"`
	DiscountPeriod float64 `json:"discountPeriod"`
	DiscountStartedAt string `json:"discountStartedAt"`
	Duration string `json:"duration"`
	ExpiredAt string `json:"expiredAt"`
	HasActivePaidSubscriptions bool `json:"hasActivePaidSubscriptions"`
	IsMuted bool `json:"isMuted"`
	MessagesSumm float64 `json:"messagesSumm"`
	NewPrice float64 `json:"newPrice"`
	PostsSumm float64 `json:"postsSumm"`
	Price float64 `json:"price"`
	RegularPrice float64 `json:"regularPrice"`
	RenewedAt string `json:"renewedAt"`
	Status string `json:"status"`
	StreamsSumm float64 `json:"streamsSumm"`
	SubscribeAt string `json:"subscribeAt"`
	SubscribePrice float64 `json:"subscribePrice"`
	Subscribes []struct {
	Action string `json:"action"`
	CancelDate string `json:"cancelDate"`
	Date string `json:"date"`
	Discount float64 `json:"discount"`
	Duration float64 `json:"duration"`
	EarningId float64 `json:"earningId"`
	ExpireDate string `json:"expireDate"`
	Id float64 `json:"id"`
	IsCurrent bool `json:"isCurrent"`
	OfferEnd string `json:"offerEnd"`
	OfferStart string `json:"offerStart"`
	Price float64 `json:"price"`
	RegularPrice float64 `json:"regularPrice"`
	StartDate string `json:"startDate"`
	SubscriberId float64 `json:"subscriberId"`
	Type string `json:"type"`
	UserId float64 `json:"userId"`
} `json:"subscribes"`
	SubscribesSumm float64 `json:"subscribesSumm"`
	TipsSumm float64 `json:"tipsSumm"`
	TotalSumm float64 `json:"totalSumm"`
	UnsubscribeReason string `json:"unsubscribeReason"`
} `json:"subscribedOnData,omitempty"`
	SubscribedOnDuration *string `json:"subscribedOnDuration,omitempty"`
	SubscribedOnExpiredNow *bool `json:"subscribedOnExpiredNow,omitempty"`
	SubscribersCount *float64 `json:"subscribersCount,omitempty"`
	TipsEnabled *bool `json:"tipsEnabled,omitempty"`
	TipsMax *float64 `json:"tipsMax,omitempty"`
	TipsMin *float64 `json:"tipsMin,omitempty"`
	TipsMinInternal *float64 `json:"tipsMinInternal,omitempty"`
	TipsTextEnabled *bool `json:"tipsTextEnabled,omitempty"`
	Username string `json:"username"`
	VideosCount *float64 `json:"videosCount,omitempty"`
	Website *string `json:"website,omitempty"`
	Wishlist *string `json:"wishlist,omitempty"`
}

// ListAccessSelfNotificationsResponse represents the response for ListAccessSelfNotifications
type ListAccessSelfNotificationsResponse struct {
	HasMore bool `json:"hasMore"`
	List []struct {
	CanGoToProfile bool `json:"canGoToProfile"`
	CreatedAt string `json:"createdAt"`
	Id float64 `json:"id"`
	IsRead bool `json:"isRead"`
	ReplacePairs map[string]string `json:"replacePairs"`
	SubType string `json:"subType"`
	Text string `json:"text"`
	Type string `json:"type"`
	UserId float64 `json:"userId"`
} `json:"list"`
}

// ListAccessSelfReleaseFormsResponse represents the response for ListAccessSelfReleaseForms
type ListAccessSelfReleaseFormsResponse struct {
	HasMore bool `json:"hasMore"`
	List []struct {
	ApprovedAt string `json:"approvedAt"`
	Code string `json:"code"`
	CreatedAt string `json:"createdAt"`
	HasUser bool `json:"hasUser"`
	Id float64 `json:"id"`
	IsHidden bool `json:"isHidden"`
	LastChangedAt string `json:"lastChangedAt"`
	Name string `json:"name"`
	PartnerId float64 `json:"partnerId"`
	Status string `json:"status"`
	Type string `json:"type"`
	UserName string `json:"userName"`
} `json:"list"`
}

// ListAccessSelfTaggedFriendUsersResponse represents the response for ListAccessSelfTaggedFriendUsers
type ListAccessSelfTaggedFriendUsersResponse struct {
	HasMore bool `json:"hasMore"`
	List []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	User struct {
	Avatar string `json:"avatar"`
	Id float64 `json:"id"`
	IsHidden bool `json:"isHidden"`
	IsVerified bool `json:"isVerified"`
	Name string `json:"name"`
	Username string `json:"username"`
} `json:"user"`
} `json:"list"`
}

// ListAccessSubscribersResponse represents the response for ListAccessSubscribers
type ListAccessSubscribersResponse struct {
	HasMore bool `json:"hasMore"`
	List []struct {
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs"`
	Capabilities struct {
	CanBlock bool `json:"canBlock"`
	CanReceiveMessages bool `json:"canReceiveMessages"`
	CanReport bool `json:"canReport"`
	CanRestrict bool `json:"canRestrict"`
	CanSendTrial bool `json:"canSendTrial"`
	CanUnsubscribe bool `json:"canUnsubscribe"`
} `json:"capabilities"`
	DisplayName string `json:"displayName"`
	Id float64 `json:"id"`
	IsBlocked bool `json:"isBlocked"`
	IsRestricted bool `json:"isRestricted"`
	IsVerified bool `json:"isVerified"`
	LastSeen string `json:"lastSeen"`
	Lists []struct {
	Id interface{} `json:"id"`
	Name string `json:"name"`
} `json:"lists"`
	Name string `json:"name"`
	Spending struct {
	Messages float64 `json:"messages"`
	Posts float64 `json:"posts"`
	Streams float64 `json:"streams"`
	Subscriptions float64 `json:"subscriptions"`
	Tips float64 `json:"tips"`
	Total float64 `json:"total"`
} `json:"spending"`
	Subscription struct {
	DiscountFinishedAt string `json:"discountFinishedAt"`
	DiscountPercent float64 `json:"discountPercent"`
	DiscountPeriod float64 `json:"discountPeriod"`
	DiscountStartedAt string `json:"discountStartedAt"`
	ExpiresAt string `json:"expiresAt"`
	History []struct {
	Action string `json:"action"`
	Discount float64 `json:"discount"`
	ExpireDate string `json:"expireDate"`
	Id float64 `json:"id"`
	IsCurrent bool `json:"isCurrent"`
	Price float64 `json:"price"`
	RegularPrice float64 `json:"regularPrice"`
	StartDate string `json:"startDate"`
	Type string `json:"type"`
} `json:"history"`
	IsActive bool `json:"isActive"`
	IsExpired bool `json:"isExpired"`
	Price float64 `json:"price"`
	RegularPrice float64 `json:"regularPrice"`
	RenewedAt string `json:"renewedAt"`
	SubscribedAt string `json:"subscribedAt"`
} `json:"subscription"`
	Username string `json:"username"`
} `json:"list"`
}

// SetAccessSubscribersCustomNameResponse represents the response for SetAccessSubscribersCustomName
type SetAccessSubscribersCustomNameResponse map[string]interface{}

// SetAccessSubscribersDiscountResponse represents the response for SetAccessSubscribersDiscount
type SetAccessSubscribersDiscountResponse map[string]interface{}

// SetAccessSubscribersNoteResponse represents the response for SetAccessSubscribersNote
type SetAccessSubscribersNoteResponse map[string]interface{}

// ListAccessSubscriptionsResponse represents the response for ListAccessSubscriptions
type ListAccessSubscriptionsResponse struct {
	HasMore bool `json:"hasMore"`
	List []struct {
	Avatar string `json:"avatar"`
	AvatarThumbs *struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs,omitempty"`
	ExpiredAt string `json:"expiredAt"`
	Id float64 `json:"id"`
	IsActive bool `json:"isActive"`
	Lists []struct {
	Id interface{} `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
} `json:"lists"`
	MessagesSumm *float64 `json:"messagesSumm,omitempty"`
	Name string `json:"name"`
	PostsSumm *float64 `json:"postsSumm,omitempty"`
	RenewedAt string `json:"renewedAt"`
	StreamsSumm *float64 `json:"streamsSumm,omitempty"`
	SubscribedAt string `json:"subscribedAt"`
	SubscribesSumm *float64 `json:"subscribesSumm,omitempty"`
	SubscriptionPrice float64 `json:"subscriptionPrice"`
	TipsSumm *float64 `json:"tipsSumm,omitempty"`
	TotalSumm *float64 `json:"totalSumm,omitempty"`
	Username string `json:"username"`
} `json:"list"`
}

// GetAccessSubscriptionsHistoryResponse represents the response for GetAccessSubscriptionsHistory
type GetAccessSubscriptionsHistoryResponse struct {
	HasMore bool `json:"hasMore"`
	List []struct {
	ExpireDate string `json:"expireDate"`
	Price float64 `json:"price"`
	SubscribeDate string `json:"subscribeDate"`
} `json:"list"`
}

// GetAccessSubscriptionsCountResponse represents the response for GetAccessSubscriptionsCount
type GetAccessSubscriptionsCountResponse struct {
	Bookmarks float64 `json:"bookmarks"`
	Subscribers struct {
	Active float64 `json:"active"`
	ActiveOnline float64 `json:"activeOnline"`
	All float64 `json:"all"`
	Blocked float64 `json:"blocked"`
	Expired float64 `json:"expired"`
	Muted float64 `json:"muted"`
	Restricted float64 `json:"restricted"`
} `json:"subscribers"`
	Subscriptions struct {
	Active float64 `json:"active"`
	All float64 `json:"all"`
	Attention float64 `json:"attention"`
	Blocked float64 `json:"blocked"`
	Expired float64 `json:"expired"`
	Muted float64 `json:"muted"`
	Restricted float64 `json:"restricted"`
} `json:"subscriptions"`
}

// ReplaceAccessUploadsResponse represents the response for ReplaceAccessUploads
type ReplaceAccessUploadsResponse struct {
	Media *struct {
	CanView bool `json:"canView"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CreatedAt string `json:"createdAt"`
	Duration *float64 `json:"duration,omitempty"`
	Files *struct {
	Full struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"sources"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"full"`
	Preview *struct {
	Height *float64 `json:"height,omitempty"`
	Options []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"options"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"squarePreview,omitempty"`
	Thumb *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"thumb,omitempty"`
} `json:"files,omitempty"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	HasError bool `json:"hasError"`
	Id float64 `json:"id"`
	IsReady bool `json:"isReady"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs"`
	Id float64 `json:"id"`
	IsFromGuest bool `json:"isFromGuest"`
	IsVerified bool `json:"isVerified"`
	IvStatus string `json:"ivStatus"`
	Name string `json:"name"`
	Username string `json:"username"`
	View string `json:"view"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	Type string `json:"type"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
} `json:"media,omitempty"`
	MediaUploadId string `json:"mediaUploadId"`
}

// ReplaceAccessUploadsPartsResponse represents the response for ReplaceAccessUploadsParts
type ReplaceAccessUploadsPartsResponse struct {
	Etag string `json:"etag"`
	MediaUploadId string `json:"mediaUploadId"`
	PartNumber int64 `json:"partNumber"`
}

// AccessUploadsCheckResponse represents the response for AccessUploadsCheck
type AccessUploadsCheckResponse struct {
	Exists bool `json:"exists"`
	Media *struct {
	CanView bool `json:"canView"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CreatedAt string `json:"createdAt"`
	Duration *float64 `json:"duration,omitempty"`
	Files *struct {
	Full struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"sources"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"full"`
	Preview *struct {
	Height *float64 `json:"height,omitempty"`
	Options []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"options"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"squarePreview,omitempty"`
	Thumb *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"thumb,omitempty"`
} `json:"files,omitempty"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	HasError bool `json:"hasError"`
	Id float64 `json:"id"`
	IsReady bool `json:"isReady"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs"`
	Id float64 `json:"id"`
	IsFromGuest bool `json:"isFromGuest"`
	IsVerified bool `json:"isVerified"`
	IvStatus string `json:"ivStatus"`
	Name string `json:"name"`
	Username string `json:"username"`
	View string `json:"view"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	Type string `json:"type"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
} `json:"media,omitempty"`
}

// AccessUploadsCompleteResponse represents the response for AccessUploadsComplete
type AccessUploadsCompleteResponse struct {
	Media *struct {
	CanView bool `json:"canView"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CreatedAt string `json:"createdAt"`
	Duration *float64 `json:"duration,omitempty"`
	Files *struct {
	Full struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"sources"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"full"`
	Preview *struct {
	Height *float64 `json:"height,omitempty"`
	Options []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"options"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"squarePreview,omitempty"`
	Thumb *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"thumb,omitempty"`
} `json:"files,omitempty"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	HasError bool `json:"hasError"`
	Id float64 `json:"id"`
	IsReady bool `json:"isReady"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs"`
	Id float64 `json:"id"`
	IsFromGuest bool `json:"isFromGuest"`
	IsVerified bool `json:"isVerified"`
	IvStatus string `json:"ivStatus"`
	Name string `json:"name"`
	Username string `json:"username"`
	View string `json:"view"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	Type string `json:"type"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
} `json:"media,omitempty"`
	MediaUploadId string `json:"mediaUploadId"`
}

// AccessUploadsInitResponse represents the response for AccessUploadsInit
type AccessUploadsInitResponse struct {
	MediaUploadId string `json:"mediaUploadId"`
}

// CreateAccessUsersListsResponse represents the response for CreateAccessUsersLists
type CreateAccessUsersListsResponse struct {
	Errors []struct {
	Error string `json:"error"`
	ListId float64 `json:"listId"`
} `json:"errors,omitempty"`
	Success []struct {
	ListId float64 `json:"listId"`
	Success bool `json:"success"`
} `json:"success"`
}

// ListAccessUsersPostsResponse represents the response for ListAccessUsersPosts
type ListAccessUsersPostsResponse struct {
	Counters struct {
	ArchivedPostsCount float64 `json:"archivedPostsCount"`
	AudiosCount float64 `json:"audiosCount"`
	MediasCount float64 `json:"mediasCount"`
	PhotosCount float64 `json:"photosCount"`
	PostsCount float64 `json:"postsCount"`
	PrivateArchivedPostsCount float64 `json:"privateArchivedPostsCount"`
	StreamsCount float64 `json:"streamsCount"`
	VideosCount float64 `json:"videosCount"`
} `json:"counters"`
	HasMore bool `json:"hasMore"`
	HeadMarker string `json:"headMarker"`
	List []struct {
	Author *struct {
	View string `json:"_view"`
	Id float64 `json:"id"`
} `json:"author,omitempty"`
	CanComment bool `json:"canComment"`
	CanDelete bool `json:"canDelete"`
	CanEdit bool `json:"canEdit"`
	CanToggleFavorite bool `json:"canToggleFavorite"`
	CanViewMedia bool `json:"canViewMedia"`
	FavoritesCount float64 `json:"favoritesCount"`
	Id float64 `json:"id"`
	IsFavorite bool `json:"isFavorite"`
	IsMarkdownDisabled bool `json:"isMarkdownDisabled"`
	IsMediaReady bool `json:"isMediaReady"`
	IsOpened bool `json:"isOpened"`
	Media []struct {
	CanView bool `json:"canView"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CreatedAt string `json:"createdAt"`
	Duration *float64 `json:"duration,omitempty"`
	Files *struct {
	Full struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"sources"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"full"`
	Preview *struct {
	Height *float64 `json:"height,omitempty"`
	Options []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"options"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"squarePreview,omitempty"`
	Thumb *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"thumb,omitempty"`
} `json:"files,omitempty"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	HasError bool `json:"hasError"`
	Id float64 `json:"id"`
	IsReady bool `json:"isReady"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs"`
	Id float64 `json:"id"`
	IsFromGuest bool `json:"isFromGuest"`
	IsVerified bool `json:"isVerified"`
	IvStatus string `json:"ivStatus"`
	Name string `json:"name"`
	Username string `json:"username"`
	View string `json:"view"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	Type string `json:"type"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
} `json:"media"`
	MediaCount float64 `json:"mediaCount"`
	PostedAt string `json:"postedAt"`
	PostedAtPrecise string `json:"postedAtPrecise"`
	RawText string `json:"rawText"`
	ResponseType string `json:"responseType"`
	Text string `json:"text"`
	TipsAmount string `json:"tipsAmount"`
} `json:"list"`
	TailMarker string `json:"tailMarker"`
}

// CreateAccessUsersRestrictResponse represents the response for CreateAccessUsersRestrict
type CreateAccessUsersRestrictResponse map[string]interface{}

// DeleteAccessUsersRestrictResponse represents the response for DeleteAccessUsersRestrict
type DeleteAccessUsersRestrictResponse map[string]interface{}

// GetAccessUsersBlockedResponse represents the response for GetAccessUsersBlocked
type GetAccessUsersBlockedResponse struct {
	HasMore bool `json:"hasMore"`
	List []struct {
	Avatar string `json:"avatar"`
	AvatarThumbs *struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs,omitempty"`
	CanReceiveChatMessage *bool `json:"canReceiveChatMessage,omitempty"`
	CanReport *bool `json:"canReport,omitempty"`
	CanRestrict *bool `json:"canRestrict,omitempty"`
	CanTrialSend *bool `json:"canTrialSend,omitempty"`
	CanUnsubscribe *bool `json:"canUnsubscribe,omitempty"`
	CurrentSubscribePrice *float64 `json:"currentSubscribePrice,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	HideChat *bool `json:"hideChat,omitempty"`
	Id float64 `json:"id"`
	IsBlocked *bool `json:"isBlocked,omitempty"`
	IsPendingAutoprolong *bool `json:"isPendingAutoprolong,omitempty"`
	IsPerformer *bool `json:"isPerformer,omitempty"`
	IsRealPerformer *bool `json:"isRealPerformer,omitempty"`
	IsRestricted *bool `json:"isRestricted,omitempty"`
	IsVerified bool `json:"isVerified"`
	LastSeen *string `json:"lastSeen,omitempty"`
	Lists []struct {
	Id interface{} `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
} `json:"lists"`
	Name string `json:"name"`
	Notice *string `json:"notice,omitempty"`
	SubscribedBy *bool `json:"subscribedBy,omitempty"`
	SubscribedByAutoprolong *bool `json:"subscribedByAutoprolong,omitempty"`
	SubscribedByData *struct {
	DiscountFinishedAt string `json:"discountFinishedAt"`
	DiscountPercent float64 `json:"discountPercent"`
	DiscountPeriod float64 `json:"discountPeriod"`
	DiscountStartedAt string `json:"discountStartedAt"`
	Duration string `json:"duration"`
	ExpiredAt string `json:"expiredAt"`
	HasActivePaidSubscriptions bool `json:"hasActivePaidSubscriptions"`
	IsMuted bool `json:"isMuted"`
	NewPrice float64 `json:"newPrice"`
	Price float64 `json:"price"`
	RegularPrice float64 `json:"regularPrice"`
	RenewedAt string `json:"renewedAt"`
	ShowPostsInFeed bool `json:"showPostsInFeed"`
	Status string `json:"status"`
	SubscribeAt string `json:"subscribeAt"`
	SubscribePrice float64 `json:"subscribePrice"`
	Subscribes []struct {
	Action string `json:"action"`
	CancelDate string `json:"cancelDate"`
	Date string `json:"date"`
	Discount float64 `json:"discount"`
	Duration float64 `json:"duration"`
	EarningId float64 `json:"earningId"`
	ExpireDate string `json:"expireDate"`
	Id float64 `json:"id"`
	IsCurrent bool `json:"isCurrent"`
	OfferEnd string `json:"offerEnd"`
	OfferStart string `json:"offerStart"`
	Price float64 `json:"price"`
	RegularPrice float64 `json:"regularPrice"`
	StartDate string `json:"startDate"`
	SubscriberId float64 `json:"subscriberId"`
	Type string `json:"type"`
	UserId float64 `json:"userId"`
} `json:"subscribes"`
	UnsubscribeReason string `json:"unsubscribeReason"`
} `json:"subscribedByData,omitempty"`
	SubscribedByExpire *bool `json:"subscribedByExpire,omitempty"`
	SubscribedByExpireDate *string `json:"subscribedByExpireDate,omitempty"`
	SubscribedIsExpiredNow *bool `json:"subscribedIsExpiredNow,omitempty"`
	SubscribedOn *bool `json:"subscribedOn,omitempty"`
	SubscribedOnData *struct {
	DiscountFinishedAt string `json:"discountFinishedAt"`
	DiscountPercent float64 `json:"discountPercent"`
	DiscountPeriod float64 `json:"discountPeriod"`
	DiscountStartedAt string `json:"discountStartedAt"`
	Duration string `json:"duration"`
	ExpiredAt string `json:"expiredAt"`
	HasActivePaidSubscriptions bool `json:"hasActivePaidSubscriptions"`
	IsMuted bool `json:"isMuted"`
	MessagesSumm float64 `json:"messagesSumm"`
	NewPrice float64 `json:"newPrice"`
	PostsSumm float64 `json:"postsSumm"`
	Price float64 `json:"price"`
	RegularPrice float64 `json:"regularPrice"`
	RenewedAt string `json:"renewedAt"`
	Status string `json:"status"`
	StreamsSumm float64 `json:"streamsSumm"`
	SubscribeAt string `json:"subscribeAt"`
	SubscribePrice float64 `json:"subscribePrice"`
	Subscribes []struct {
	Action string `json:"action"`
	CancelDate string `json:"cancelDate"`
	Date string `json:"date"`
	Discount float64 `json:"discount"`
	Duration float64 `json:"duration"`
	EarningId float64 `json:"earningId"`
	ExpireDate string `json:"expireDate"`
	Id float64 `json:"id"`
	IsCurrent bool `json:"isCurrent"`
	OfferEnd string `json:"offerEnd"`
	OfferStart string `json:"offerStart"`
	Price float64 `json:"price"`
	RegularPrice float64 `json:"regularPrice"`
	StartDate string `json:"startDate"`
	SubscriberId float64 `json:"subscriberId"`
	Type string `json:"type"`
	UserId float64 `json:"userId"`
} `json:"subscribes"`
	SubscribesSumm float64 `json:"subscribesSumm"`
	TipsSumm float64 `json:"tipsSumm"`
	TotalSumm float64 `json:"totalSumm"`
	UnsubscribeReason string `json:"unsubscribeReason"`
} `json:"subscribedOnData,omitempty"`
	SubscribedOnDuration *string `json:"subscribedOnDuration,omitempty"`
	SubscribedOnExpiredNow *bool `json:"subscribedOnExpiredNow,omitempty"`
	Username string `json:"username"`
} `json:"list"`
}

// GetAccessUsersListResponse represents the response for GetAccessUsersList
type GetAccessUsersListResponse struct {
	Users []struct {
	About *string `json:"about,omitempty"`
	ArchivedPostsCount *float64 `json:"archivedPostsCount,omitempty"`
	AudiosCount *float64 `json:"audiosCount,omitempty"`
	Avatar string `json:"avatar"`
	AvatarThumbs *struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs,omitempty"`
	CanAddSubscriber *bool `json:"canAddSubscriber,omitempty"`
	CanChat *bool `json:"canChat,omitempty"`
	CanCommentStory *bool `json:"canCommentStory,omitempty"`
	CanCreatePromotion *bool `json:"canCreatePromotion,omitempty"`
	CanCreateTrial *bool `json:"canCreateTrial,omitempty"`
	CanLookStory *bool `json:"canLookStory,omitempty"`
	CanPayInternal *bool `json:"canPayInternal,omitempty"`
	CanReceiveChatMessage *bool `json:"canReceiveChatMessage,omitempty"`
	CanReport *bool `json:"canReport,omitempty"`
	CanRestrict *bool `json:"canRestrict,omitempty"`
	CanTrialSend *bool `json:"canTrialSend,omitempty"`
	CanUnsubscribe *bool `json:"canUnsubscribe,omitempty"`
	CurrentSubscribePrice *float64 `json:"currentSubscribePrice,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	FavoritedCount *float64 `json:"favoritedCount,omitempty"`
	FavoritesCount *float64 `json:"favoritesCount,omitempty"`
	HasLabels *bool `json:"hasLabels,omitempty"`
	HasNotViewedStory *bool `json:"hasNotViewedStory,omitempty"`
	HasPinnedPosts *bool `json:"hasPinnedPosts,omitempty"`
	HasScheduledStream *bool `json:"hasScheduledStream,omitempty"`
	HasStories *bool `json:"hasStories,omitempty"`
	HasStream *bool `json:"hasStream,omitempty"`
	Header *string `json:"header,omitempty"`
	HeaderSize *struct {
	Height float64 `json:"height"`
	Width float64 `json:"width"`
} `json:"headerSize,omitempty"`
	HeaderThumbs *struct {
	W480 string `json:"w480"`
	W760 string `json:"w760"`
} `json:"headerThumbs,omitempty"`
	Id float64 `json:"id"`
	IsAdultContent *bool `json:"isAdultContent,omitempty"`
	IsBlocked *bool `json:"isBlocked,omitempty"`
	IsFriend *bool `json:"isFriend,omitempty"`
	IsMarkdownDisabledForAbout *bool `json:"isMarkdownDisabledForAbout,omitempty"`
	IsPendingAutoprolong *bool `json:"isPendingAutoprolong,omitempty"`
	IsPerformer *bool `json:"isPerformer,omitempty"`
	IsPrivateRestriction *bool `json:"isPrivateRestriction,omitempty"`
	IsRealPerformer *bool `json:"isRealPerformer,omitempty"`
	IsReferrerAllowed *bool `json:"isReferrerAllowed,omitempty"`
	IsRestricted *bool `json:"isRestricted,omitempty"`
	IsVerified bool `json:"isVerified"`
	JoinDate *string `json:"joinDate,omitempty"`
	LastSeen *string `json:"lastSeen,omitempty"`
	Lists []struct {
	Id interface{} `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
} `json:"lists"`
	Location *string `json:"location,omitempty"`
	MediasCount *float64 `json:"mediasCount,omitempty"`
	Name string `json:"name"`
	Notice *string `json:"notice,omitempty"`
	PhotosCount *float64 `json:"photosCount,omitempty"`
	PostsCount *float64 `json:"postsCount,omitempty"`
	PrivateArchivedPostsCount *float64 `json:"privateArchivedPostsCount,omitempty"`
	ShowMediaCount *bool `json:"showMediaCount,omitempty"`
	ShowPostsInFeed *bool `json:"showPostsInFeed,omitempty"`
	ShowSubscribersCount *bool `json:"showSubscribersCount,omitempty"`
	SubscribePrice *float64 `json:"subscribePrice,omitempty"`
	SubscribedBy *bool `json:"subscribedBy,omitempty"`
	SubscribedByAutoprolong *bool `json:"subscribedByAutoprolong,omitempty"`
	SubscribedByData *struct {
	DiscountFinishedAt string `json:"discountFinishedAt"`
	DiscountPercent float64 `json:"discountPercent"`
	DiscountPeriod float64 `json:"discountPeriod"`
	DiscountStartedAt string `json:"discountStartedAt"`
	Duration string `json:"duration"`
	ExpiredAt string `json:"expiredAt"`
	HasActivePaidSubscriptions bool `json:"hasActivePaidSubscriptions"`
	IsMuted bool `json:"isMuted"`
	NewPrice float64 `json:"newPrice"`
	Price float64 `json:"price"`
	RegularPrice float64 `json:"regularPrice"`
	RenewedAt string `json:"renewedAt"`
	ShowPostsInFeed bool `json:"showPostsInFeed"`
	Status string `json:"status"`
	SubscribeAt string `json:"subscribeAt"`
	SubscribePrice float64 `json:"subscribePrice"`
	Subscribes []struct {
	Action string `json:"action"`
	CancelDate string `json:"cancelDate"`
	Date string `json:"date"`
	Discount float64 `json:"discount"`
	Duration float64 `json:"duration"`
	EarningId float64 `json:"earningId"`
	ExpireDate string `json:"expireDate"`
	Id float64 `json:"id"`
	IsCurrent bool `json:"isCurrent"`
	OfferEnd string `json:"offerEnd"`
	OfferStart string `json:"offerStart"`
	Price float64 `json:"price"`
	RegularPrice float64 `json:"regularPrice"`
	StartDate string `json:"startDate"`
	SubscriberId float64 `json:"subscriberId"`
	Type string `json:"type"`
	UserId float64 `json:"userId"`
} `json:"subscribes"`
	UnsubscribeReason string `json:"unsubscribeReason"`
} `json:"subscribedByData,omitempty"`
	SubscribedByExpire *bool `json:"subscribedByExpire,omitempty"`
	SubscribedByExpireDate *string `json:"subscribedByExpireDate,omitempty"`
	SubscribedIsExpiredNow *bool `json:"subscribedIsExpiredNow,omitempty"`
	SubscribedOn *bool `json:"subscribedOn,omitempty"`
	SubscribedOnData *struct {
	DiscountFinishedAt string `json:"discountFinishedAt"`
	DiscountPercent float64 `json:"discountPercent"`
	DiscountPeriod float64 `json:"discountPeriod"`
	DiscountStartedAt string `json:"discountStartedAt"`
	Duration string `json:"duration"`
	ExpiredAt string `json:"expiredAt"`
	HasActivePaidSubscriptions bool `json:"hasActivePaidSubscriptions"`
	IsMuted bool `json:"isMuted"`
	MessagesSumm float64 `json:"messagesSumm"`
	NewPrice float64 `json:"newPrice"`
	PostsSumm float64 `json:"postsSumm"`
	Price float64 `json:"price"`
	RegularPrice float64 `json:"regularPrice"`
	RenewedAt string `json:"renewedAt"`
	Status string `json:"status"`
	StreamsSumm float64 `json:"streamsSumm"`
	SubscribeAt string `json:"subscribeAt"`
	SubscribePrice float64 `json:"subscribePrice"`
	Subscribes []struct {
	Action string `json:"action"`
	CancelDate string `json:"cancelDate"`
	Date string `json:"date"`
	Discount float64 `json:"discount"`
	Duration float64 `json:"duration"`
	EarningId float64 `json:"earningId"`
	ExpireDate string `json:"expireDate"`
	Id float64 `json:"id"`
	IsCurrent bool `json:"isCurrent"`
	OfferEnd string `json:"offerEnd"`
	OfferStart string `json:"offerStart"`
	Price float64 `json:"price"`
	RegularPrice float64 `json:"regularPrice"`
	StartDate string `json:"startDate"`
	SubscriberId float64 `json:"subscriberId"`
	Type string `json:"type"`
	UserId float64 `json:"userId"`
} `json:"subscribes"`
	SubscribesSumm float64 `json:"subscribesSumm"`
	TipsSumm float64 `json:"tipsSumm"`
	TotalSumm float64 `json:"totalSumm"`
	UnsubscribeReason string `json:"unsubscribeReason"`
} `json:"subscribedOnData,omitempty"`
	SubscribedOnDuration *string `json:"subscribedOnDuration,omitempty"`
	SubscribedOnExpiredNow *bool `json:"subscribedOnExpiredNow,omitempty"`
	SubscribersCount *float64 `json:"subscribersCount,omitempty"`
	TipsEnabled *bool `json:"tipsEnabled,omitempty"`
	TipsMax *float64 `json:"tipsMax,omitempty"`
	TipsMin *float64 `json:"tipsMin,omitempty"`
	TipsMinInternal *float64 `json:"tipsMinInternal,omitempty"`
	TipsTextEnabled *bool `json:"tipsTextEnabled,omitempty"`
	Username string `json:"username"`
	VideosCount *float64 `json:"videosCount,omitempty"`
	Website *string `json:"website,omitempty"`
	Wishlist *string `json:"wishlist,omitempty"`
} `json:"users"`
}

// ListAccessUsersListsResponse represents the response for ListAccessUsersLists
type ListAccessUsersListsResponse struct {
	HasMore bool `json:"hasMore"`
	List []struct {
	CanAddUsers *bool `json:"canAddUsers,omitempty"`
	CanDelete *bool `json:"canDelete,omitempty"`
	CanManageUsers *bool `json:"canManageUsers,omitempty"`
	CanPinnedToChat *bool `json:"canPinnedToChat,omitempty"`
	CanPinnedToFeed *bool `json:"canPinnedToFeed,omitempty"`
	CanUpdate *bool `json:"canUpdate,omitempty"`
	Direction *string `json:"direction,omitempty"`
	Id interface{} `json:"id"`
	IsPinnedToChat *bool `json:"isPinnedToChat,omitempty"`
	IsPinnedToFeed *bool `json:"isPinnedToFeed,omitempty"`
	Name string `json:"name"`
	Order *string `json:"order,omitempty"`
	PostsCount *float64 `json:"postsCount,omitempty"`
	SortList []struct {
	Direction string `json:"direction"`
	Order string `json:"order"`
} `json:"sortList"`
	Type string `json:"type"`
	Users []struct {
	Avatar string `json:"avatar"`
	AvatarThumbs *struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs,omitempty"`
	Id float64 `json:"id"`
	IsVerified bool `json:"isVerified"`
	Name string `json:"name"`
	Username string `json:"username"`
} `json:"users"`
	UsersCount *float64 `json:"usersCount,omitempty"`
} `json:"list"`
	Order *string `json:"order,omitempty"`
	Sort *string `json:"sort,omitempty"`
}

// CreateAccessUsersLists1Response represents the response for CreateAccessUsersLists1
type CreateAccessUsersLists1Response struct {
	CanAddUsers *bool `json:"canAddUsers,omitempty"`
	CanDelete *bool `json:"canDelete,omitempty"`
	CanManageUsers *bool `json:"canManageUsers,omitempty"`
	CanPinnedToChat *bool `json:"canPinnedToChat,omitempty"`
	CanPinnedToFeed *bool `json:"canPinnedToFeed,omitempty"`
	CanUpdate *bool `json:"canUpdate,omitempty"`
	Direction *string `json:"direction,omitempty"`
	Id interface{} `json:"id"`
	IsPinnedToChat *bool `json:"isPinnedToChat,omitempty"`
	IsPinnedToFeed *bool `json:"isPinnedToFeed,omitempty"`
	Name string `json:"name"`
	Order *string `json:"order,omitempty"`
	PostsCount *float64 `json:"postsCount,omitempty"`
	SortList []struct {
	Direction string `json:"direction"`
	Order string `json:"order"`
} `json:"sortList"`
	Type string `json:"type"`
	Users []struct {
	Avatar string `json:"avatar"`
	AvatarThumbs *struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs,omitempty"`
	Id float64 `json:"id"`
	IsVerified bool `json:"isVerified"`
	Name string `json:"name"`
	Username string `json:"username"`
} `json:"users"`
	UsersCount *float64 `json:"usersCount,omitempty"`
}

// GetAccessUsersListsResponse represents the response for GetAccessUsersLists
type GetAccessUsersListsResponse struct {
	CanAddUsers *bool `json:"canAddUsers,omitempty"`
	CanDelete *bool `json:"canDelete,omitempty"`
	CanManageUsers *bool `json:"canManageUsers,omitempty"`
	CanPinnedToChat *bool `json:"canPinnedToChat,omitempty"`
	CanPinnedToFeed *bool `json:"canPinnedToFeed,omitempty"`
	CanUpdate *bool `json:"canUpdate,omitempty"`
	Direction *string `json:"direction,omitempty"`
	Id interface{} `json:"id"`
	IsPinnedToChat *bool `json:"isPinnedToChat,omitempty"`
	IsPinnedToFeed *bool `json:"isPinnedToFeed,omitempty"`
	Name string `json:"name"`
	Order *string `json:"order,omitempty"`
	PostsCount *float64 `json:"postsCount,omitempty"`
	SortList []struct {
	Direction string `json:"direction"`
	Order string `json:"order"`
} `json:"sortList"`
	Type string `json:"type"`
	Users []struct {
	Avatar string `json:"avatar"`
	AvatarThumbs *struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs,omitempty"`
	Id float64 `json:"id"`
	IsVerified bool `json:"isVerified"`
	Name string `json:"name"`
	Username string `json:"username"`
} `json:"users"`
	UsersCount *float64 `json:"usersCount,omitempty"`
}

// UpdateAccessUsersListsResponse represents the response for UpdateAccessUsersLists
type UpdateAccessUsersListsResponse struct {
	CanAddUsers *bool `json:"canAddUsers,omitempty"`
	CanDelete *bool `json:"canDelete,omitempty"`
	CanManageUsers *bool `json:"canManageUsers,omitempty"`
	CanPinnedToChat *bool `json:"canPinnedToChat,omitempty"`
	CanPinnedToFeed *bool `json:"canPinnedToFeed,omitempty"`
	CanUpdate *bool `json:"canUpdate,omitempty"`
	Direction *string `json:"direction,omitempty"`
	Id interface{} `json:"id"`
	IsPinnedToChat *bool `json:"isPinnedToChat,omitempty"`
	IsPinnedToFeed *bool `json:"isPinnedToFeed,omitempty"`
	Name string `json:"name"`
	Order *string `json:"order,omitempty"`
	PostsCount *float64 `json:"postsCount,omitempty"`
	SortList []struct {
	Direction string `json:"direction"`
	Order string `json:"order"`
} `json:"sortList"`
	Type string `json:"type"`
	Users []struct {
	Avatar string `json:"avatar"`
	AvatarThumbs *struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs,omitempty"`
	Id float64 `json:"id"`
	IsVerified bool `json:"isVerified"`
	Name string `json:"name"`
	Username string `json:"username"`
} `json:"users"`
	UsersCount *float64 `json:"usersCount,omitempty"`
}

// DeleteAccessUsersListsResponse represents the response for DeleteAccessUsersLists
type DeleteAccessUsersListsResponse map[string]interface{}

// ListAccessUsersListsUsersResponse represents the response for ListAccessUsersListsUsers
type ListAccessUsersListsUsersResponse struct {
	HasMore bool `json:"hasMore"`
	List []struct {
	Avatar string `json:"avatar"`
	AvatarThumbs *struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs,omitempty"`
	Id float64 `json:"id"`
	IsVerified bool `json:"isVerified"`
	Name string `json:"name"`
	Username string `json:"username"`
} `json:"list"`
	NextOffset *float64 `json:"nextOffset,omitempty"`
}

// GetAccessUsersRestrictResponse represents the response for GetAccessUsersRestrict
type GetAccessUsersRestrictResponse struct {
	HasMore bool `json:"hasMore"`
	List []struct {
	Avatar string `json:"avatar"`
	AvatarThumbs *struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs,omitempty"`
	CanReceiveChatMessage *bool `json:"canReceiveChatMessage,omitempty"`
	CanReport *bool `json:"canReport,omitempty"`
	CanRestrict *bool `json:"canRestrict,omitempty"`
	CanTrialSend *bool `json:"canTrialSend,omitempty"`
	CanUnsubscribe *bool `json:"canUnsubscribe,omitempty"`
	CurrentSubscribePrice *float64 `json:"currentSubscribePrice,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	HideChat *bool `json:"hideChat,omitempty"`
	Id float64 `json:"id"`
	IsBlocked *bool `json:"isBlocked,omitempty"`
	IsPendingAutoprolong *bool `json:"isPendingAutoprolong,omitempty"`
	IsPerformer *bool `json:"isPerformer,omitempty"`
	IsRealPerformer *bool `json:"isRealPerformer,omitempty"`
	IsRestricted *bool `json:"isRestricted,omitempty"`
	IsVerified bool `json:"isVerified"`
	LastSeen *string `json:"lastSeen,omitempty"`
	Lists []struct {
	Id interface{} `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
} `json:"lists"`
	Name string `json:"name"`
	Notice *string `json:"notice,omitempty"`
	SubscribedBy *bool `json:"subscribedBy,omitempty"`
	SubscribedByAutoprolong *bool `json:"subscribedByAutoprolong,omitempty"`
	SubscribedByData *struct {
	DiscountFinishedAt string `json:"discountFinishedAt"`
	DiscountPercent float64 `json:"discountPercent"`
	DiscountPeriod float64 `json:"discountPeriod"`
	DiscountStartedAt string `json:"discountStartedAt"`
	Duration string `json:"duration"`
	ExpiredAt string `json:"expiredAt"`
	HasActivePaidSubscriptions bool `json:"hasActivePaidSubscriptions"`
	IsMuted bool `json:"isMuted"`
	NewPrice float64 `json:"newPrice"`
	Price float64 `json:"price"`
	RegularPrice float64 `json:"regularPrice"`
	RenewedAt string `json:"renewedAt"`
	ShowPostsInFeed bool `json:"showPostsInFeed"`
	Status string `json:"status"`
	SubscribeAt string `json:"subscribeAt"`
	SubscribePrice float64 `json:"subscribePrice"`
	Subscribes []struct {
	Action string `json:"action"`
	CancelDate string `json:"cancelDate"`
	Date string `json:"date"`
	Discount float64 `json:"discount"`
	Duration float64 `json:"duration"`
	EarningId float64 `json:"earningId"`
	ExpireDate string `json:"expireDate"`
	Id float64 `json:"id"`
	IsCurrent bool `json:"isCurrent"`
	OfferEnd string `json:"offerEnd"`
	OfferStart string `json:"offerStart"`
	Price float64 `json:"price"`
	RegularPrice float64 `json:"regularPrice"`
	StartDate string `json:"startDate"`
	SubscriberId float64 `json:"subscriberId"`
	Type string `json:"type"`
	UserId float64 `json:"userId"`
} `json:"subscribes"`
	UnsubscribeReason string `json:"unsubscribeReason"`
} `json:"subscribedByData,omitempty"`
	SubscribedByExpire *bool `json:"subscribedByExpire,omitempty"`
	SubscribedByExpireDate *string `json:"subscribedByExpireDate,omitempty"`
	SubscribedIsExpiredNow *bool `json:"subscribedIsExpiredNow,omitempty"`
	SubscribedOn *bool `json:"subscribedOn,omitempty"`
	SubscribedOnData *struct {
	DiscountFinishedAt string `json:"discountFinishedAt"`
	DiscountPercent float64 `json:"discountPercent"`
	DiscountPeriod float64 `json:"discountPeriod"`
	DiscountStartedAt string `json:"discountStartedAt"`
	Duration string `json:"duration"`
	ExpiredAt string `json:"expiredAt"`
	HasActivePaidSubscriptions bool `json:"hasActivePaidSubscriptions"`
	IsMuted bool `json:"isMuted"`
	MessagesSumm float64 `json:"messagesSumm"`
	NewPrice float64 `json:"newPrice"`
	PostsSumm float64 `json:"postsSumm"`
	Price float64 `json:"price"`
	RegularPrice float64 `json:"regularPrice"`
	RenewedAt string `json:"renewedAt"`
	Status string `json:"status"`
	StreamsSumm float64 `json:"streamsSumm"`
	SubscribeAt string `json:"subscribeAt"`
	SubscribePrice float64 `json:"subscribePrice"`
	Subscribes []struct {
	Action string `json:"action"`
	CancelDate string `json:"cancelDate"`
	Date string `json:"date"`
	Discount float64 `json:"discount"`
	Duration float64 `json:"duration"`
	EarningId float64 `json:"earningId"`
	ExpireDate string `json:"expireDate"`
	Id float64 `json:"id"`
	IsCurrent bool `json:"isCurrent"`
	OfferEnd string `json:"offerEnd"`
	OfferStart string `json:"offerStart"`
	Price float64 `json:"price"`
	RegularPrice float64 `json:"regularPrice"`
	StartDate string `json:"startDate"`
	SubscriberId float64 `json:"subscriberId"`
	Type string `json:"type"`
	UserId float64 `json:"userId"`
} `json:"subscribes"`
	SubscribesSumm float64 `json:"subscribesSumm"`
	TipsSumm float64 `json:"tipsSumm"`
	TotalSumm float64 `json:"totalSumm"`
	UnsubscribeReason string `json:"unsubscribeReason"`
} `json:"subscribedOnData,omitempty"`
	SubscribedOnDuration *string `json:"subscribedOnDuration,omitempty"`
	SubscribedOnExpiredNow *bool `json:"subscribedOnExpiredNow,omitempty"`
	Username string `json:"username"`
} `json:"list"`
}

// ListAccessVaultListsResponse represents the response for ListAccessVaultLists
type ListAccessVaultListsResponse struct {
	All struct {
	AudiosCount float64 `json:"audiosCount"`
	GifsCount float64 `json:"gifsCount"`
	Medias []struct {
	Type interface{} `json:"type"`
} `json:"medias"`
	PhotosCount float64 `json:"photosCount"`
	VideosCount float64 `json:"videosCount"`
} `json:"all"`
	CanCreateVaultLists bool `json:"canCreateVaultLists"`
	HasMore bool `json:"hasMore"`
	List []struct {
	CanDelete bool `json:"canDelete"`
	CanUpdate bool `json:"canUpdate"`
	HasMedia bool `json:"hasMedia"`
	Id float64 `json:"id"`
	Medias []struct {
	Type string `json:"type"`
	Url *string `json:"url,omitempty"`
} `json:"medias"`
	Name string `json:"name"`
	Type string `json:"type"`
} `json:"list"`
	Order string `json:"order"`
	Sort string `json:"sort"`
}

// CreateAccessVaultListsResponse represents the response for CreateAccessVaultLists
type CreateAccessVaultListsResponse struct {
	CanDelete bool `json:"canDelete"`
	CanUpdate bool `json:"canUpdate"`
	HasMedia bool `json:"hasMedia"`
	Id float64 `json:"id"`
	Medias []struct {
	Type string `json:"type"`
	Url *string `json:"url,omitempty"`
} `json:"medias"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// UpdateAccessVaultListsResponse represents the response for UpdateAccessVaultLists
type UpdateAccessVaultListsResponse struct {
	CanDelete bool `json:"canDelete"`
	CanUpdate bool `json:"canUpdate"`
	HasMedia bool `json:"hasMedia"`
	Id float64 `json:"id"`
	Medias []struct {
	Type string `json:"type"`
	Url *string `json:"url,omitempty"`
} `json:"medias"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// DeleteAccessVaultListsResponse represents the response for DeleteAccessVaultLists
type DeleteAccessVaultListsResponse map[string]interface{}

// ListAccessVaultListsMediaResponse represents the response for ListAccessVaultListsMedia
type ListAccessVaultListsMediaResponse struct {
	HasMore bool `json:"hasMore"`
	List []struct {
	CanView bool `json:"canView"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	Counters struct {
	LikesCount float64 `json:"likesCount"`
	TipsSumm float64 `json:"tipsSumm"`
} `json:"counters"`
	CreatedAt string `json:"createdAt"`
	Duration *float64 `json:"duration,omitempty"`
	Files *struct {
	Full struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"sources"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"full"`
	Preview *struct {
	Height *float64 `json:"height,omitempty"`
	Options []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"options"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"squarePreview,omitempty"`
	Thumb *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"thumb,omitempty"`
} `json:"files,omitempty"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	HasError bool `json:"hasError"`
	HasPosts *bool `json:"hasPosts,omitempty"`
	Id float64 `json:"id"`
	IsReady bool `json:"isReady"`
	Lists []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
} `json:"lists"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs"`
	Id float64 `json:"id"`
	IsFromGuest bool `json:"isFromGuest"`
	IsVerified bool `json:"isVerified"`
	IvStatus string `json:"ivStatus"`
	Name string `json:"name"`
	Username string `json:"username"`
	View string `json:"view"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	Type string `json:"type"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
} `json:"list"`
}

// CreateAccessVaultListsMediaResponse represents the response for CreateAccessVaultListsMedia
type CreateAccessVaultListsMediaResponse struct {
	CanDelete bool `json:"canDelete"`
	CanUpdate bool `json:"canUpdate"`
	HasMedia bool `json:"hasMedia"`
	Id float64 `json:"id"`
	Medias []struct {
	Type string `json:"type"`
	Url *string `json:"url,omitempty"`
} `json:"medias"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// ListAccessVaultMediaResponse represents the response for ListAccessVaultMedia
type ListAccessVaultMediaResponse struct {
	HasMore bool `json:"hasMore"`
	List []struct {
	CanView bool `json:"canView"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	Counters struct {
	LikesCount float64 `json:"likesCount"`
	TipsSumm float64 `json:"tipsSumm"`
} `json:"counters"`
	CreatedAt string `json:"createdAt"`
	Duration *float64 `json:"duration,omitempty"`
	Files *struct {
	Full struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"sources"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"full"`
	Preview *struct {
	Height *float64 `json:"height,omitempty"`
	Options []struct {
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"options"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"squarePreview,omitempty"`
	Thumb *struct {
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
} `json:"thumb,omitempty"`
} `json:"files,omitempty"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	HasError bool `json:"hasError"`
	HasPosts *bool `json:"hasPosts,omitempty"`
	Id float64 `json:"id"`
	IsReady bool `json:"isReady"`
	Lists []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
} `json:"lists"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C144 string `json:"c144"`
	C50 string `json:"c50"`
} `json:"avatarThumbs"`
	Id float64 `json:"id"`
	IsFromGuest bool `json:"isFromGuest"`
	IsVerified bool `json:"isVerified"`
	IvStatus string `json:"ivStatus"`
	Name string `json:"name"`
	Username string `json:"username"`
	View string `json:"view"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	Type string `json:"type"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
} `json:"list"`
}

// ListAccountConnectionsResponse represents the response for ListAccountConnections
type ListAccountConnectionsResponse struct {
	HasMore bool `json:"hasMore"`
	List []struct {
	ClientReferenceId string `json:"clientReferenceId"`
	CreatedAt string `json:"createdAt"`
	ExpiredAt string `json:"expiredAt"`
	Id string `json:"id"`
	Imported bool `json:"imported"`
	LastCheckedAt string `json:"lastCheckedAt"`
	Permissions []string `json:"permissions"`
	Status string `json:"status"`
	UpdatedAt string `json:"updatedAt"`
	UserData struct {
	Avatar string `json:"avatar"`
	Id string `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
} `json:"userData"`
} `json:"list"`
}

// CreateAccountConnectionsImportResponse represents the response for CreateAccountConnectionsImport
type CreateAccountConnectionsImportResponse struct {
	ClientReferenceId string `json:"clientReferenceId"`
	CreatedAt string `json:"createdAt"`
	ExpiredAt string `json:"expiredAt"`
	Id string `json:"id"`
	Imported bool `json:"imported"`
	LastCheckedAt string `json:"lastCheckedAt"`
	Permissions []string `json:"permissions"`
	Status string `json:"status"`
	UpdatedAt string `json:"updatedAt"`
	UserData struct {
	Avatar string `json:"avatar"`
	Id string `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
} `json:"userData"`
}

// UpdateAccountConnectionsImportResponse represents the response for UpdateAccountConnectionsImport
type UpdateAccountConnectionsImportResponse struct {
	ClientReferenceId string `json:"clientReferenceId"`
	CreatedAt string `json:"createdAt"`
	ExpiredAt string `json:"expiredAt"`
	Id string `json:"id"`
	Imported bool `json:"imported"`
	LastCheckedAt string `json:"lastCheckedAt"`
	Permissions []string `json:"permissions"`
	Status string `json:"status"`
	UpdatedAt string `json:"updatedAt"`
	UserData struct {
	Avatar string `json:"avatar"`
	Id string `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
} `json:"userData"`
}

// GetAccountConnectionsSettingsResponse represents the response for GetAccountConnectionsSettings
type GetAccountConnectionsSettingsResponse struct {
	ConnectionId string `json:"connectionId"`
	VaultPlus struct {
	Enabled bool `json:"enabled"`
	SettingsOverrides struct {
	AccessExpiryDays *float64 `json:"accessExpiryDays,omitempty"`
	AutoCacheMessages *bool `json:"autoCacheMessages,omitempty"`
	AutoCachePosts *bool `json:"autoCachePosts,omitempty"`
	AutoCacheStories *bool `json:"autoCacheStories,omitempty"`
	AutoCacheVault *bool `json:"autoCacheVault,omitempty"`
	CacheAudio *bool `json:"cacheAudio,omitempty"`
	CacheImages *bool `json:"cacheImages,omitempty"`
	CacheVideos *bool `json:"cacheVideos,omitempty"`
	ImageQualities []string `json:"imageQualities,omitempty"`
	MinAccessCountMessages *float64 `json:"minAccessCountMessages,omitempty"`
	MinAccessCountPosts *float64 `json:"minAccessCountPosts,omitempty"`
	MinAccessCountStories *float64 `json:"minAccessCountStories,omitempty"`
	MinAccessCountVault *float64 `json:"minAccessCountVault,omitempty"`
	PresignedUrlTtlSeconds *float64 `json:"presignedUrlTtlSeconds,omitempty"`
	RetentionDays *float64 `json:"retentionDays,omitempty"`
	StorageLimitBytes *float64 `json:"storageLimitBytes,omitempty"`
	StorageLimitPurgeStrategy *string `json:"storageLimitPurgeStrategy,omitempty"`
	VideoQualities []string `json:"videoQualities,omitempty"`
} `json:"settingsOverrides"`
	Stats struct {
	MediaCount float64 `json:"mediaCount"`
	StorageLimitBytes float64 `json:"storageLimitBytes"`
	StorageUsagePercent float64 `json:"storageUsagePercent"`
	StoredCount float64 `json:"storedCount"`
	TotalStorageBytes float64 `json:"totalStorageBytes"`
	VaultPlusEnabled bool `json:"vaultPlusEnabled"`
} `json:"stats"`
} `json:"vaultPlus"`
}

// UpdateAccountConnectionsSettingsResponse represents the response for UpdateAccountConnectionsSettings
type UpdateAccountConnectionsSettingsResponse struct {
	PurgeResult *struct {
	FreedBytes float64 `json:"freedBytes"`
	PurgedCount float64 `json:"purgedCount"`
} `json:"purgeResult,omitempty"`
	Settings struct {
	ConnectionId string `json:"connectionId"`
	VaultPlus struct {
	Enabled bool `json:"enabled"`
	SettingsOverrides struct {
	AccessExpiryDays *float64 `json:"accessExpiryDays,omitempty"`
	AutoCacheMessages *bool `json:"autoCacheMessages,omitempty"`
	AutoCachePosts *bool `json:"autoCachePosts,omitempty"`
	AutoCacheStories *bool `json:"autoCacheStories,omitempty"`
	AutoCacheVault *bool `json:"autoCacheVault,omitempty"`
	CacheAudio *bool `json:"cacheAudio,omitempty"`
	CacheImages *bool `json:"cacheImages,omitempty"`
	CacheVideos *bool `json:"cacheVideos,omitempty"`
	ImageQualities []string `json:"imageQualities,omitempty"`
	MinAccessCountMessages *float64 `json:"minAccessCountMessages,omitempty"`
	MinAccessCountPosts *float64 `json:"minAccessCountPosts,omitempty"`
	MinAccessCountStories *float64 `json:"minAccessCountStories,omitempty"`
	MinAccessCountVault *float64 `json:"minAccessCountVault,omitempty"`
	PresignedUrlTtlSeconds *float64 `json:"presignedUrlTtlSeconds,omitempty"`
	RetentionDays *float64 `json:"retentionDays,omitempty"`
	StorageLimitBytes *float64 `json:"storageLimitBytes,omitempty"`
	StorageLimitPurgeStrategy *string `json:"storageLimitPurgeStrategy,omitempty"`
	VideoQualities []string `json:"videoQualities,omitempty"`
} `json:"settingsOverrides"`
	Stats struct {
	MediaCount float64 `json:"mediaCount"`
	StorageLimitBytes float64 `json:"storageLimitBytes"`
	StorageUsagePercent float64 `json:"storageUsagePercent"`
	StoredCount float64 `json:"storedCount"`
	TotalStorageBytes float64 `json:"totalStorageBytes"`
	VaultPlusEnabled bool `json:"vaultPlusEnabled"`
} `json:"stats"`
} `json:"vaultPlus"`
} `json:"settings"`
}

// GetAccountSettingsResponse represents the response for GetAccountSettings
type GetAccountSettingsResponse struct {
	VaultPlus struct {
	AutoEnableForNewConnections bool `json:"autoEnableForNewConnections"`
	DefaultSettings struct {
	AccessExpiryDays float64 `json:"accessExpiryDays"`
	AutoCacheMessages bool `json:"autoCacheMessages"`
	AutoCachePosts bool `json:"autoCachePosts"`
	AutoCacheStories bool `json:"autoCacheStories"`
	AutoCacheVault bool `json:"autoCacheVault"`
	CacheAudio bool `json:"cacheAudio"`
	CacheImages bool `json:"cacheImages"`
	CacheVideos bool `json:"cacheVideos"`
	ImageQualities []string `json:"imageQualities"`
	MinAccessCountMessages float64 `json:"minAccessCountMessages"`
	MinAccessCountPosts float64 `json:"minAccessCountPosts"`
	MinAccessCountStories float64 `json:"minAccessCountStories"`
	MinAccessCountVault float64 `json:"minAccessCountVault"`
	PresignedUrlTtlSeconds float64 `json:"presignedUrlTtlSeconds"`
	RetentionDays float64 `json:"retentionDays"`
	StorageLimitBytes float64 `json:"storageLimitBytes"`
	StorageLimitPurgeStrategy string `json:"storageLimitPurgeStrategy"`
	VideoQualities []string `json:"videoQualities"`
} `json:"defaultSettings"`
} `json:"vaultPlus"`
}

// UpdateAccountSettingsResponse represents the response for UpdateAccountSettings
type UpdateAccountSettingsResponse struct {
	BroadcastResult *struct {
	AffectedConnections float64 `json:"affectedConnections"`
	PurgeResults []struct {
	ConnectionId string `json:"connectionId"`
	FreedBytes float64 `json:"freedBytes"`
	PurgedCount float64 `json:"purgedCount"`
} `json:"purgeResults"`
} `json:"broadcastResult,omitempty"`
	Settings struct {
	VaultPlus struct {
	AutoEnableForNewConnections bool `json:"autoEnableForNewConnections"`
	DefaultSettings struct {
	AccessExpiryDays float64 `json:"accessExpiryDays"`
	AutoCacheMessages bool `json:"autoCacheMessages"`
	AutoCachePosts bool `json:"autoCachePosts"`
	AutoCacheStories bool `json:"autoCacheStories"`
	AutoCacheVault bool `json:"autoCacheVault"`
	CacheAudio bool `json:"cacheAudio"`
	CacheImages bool `json:"cacheImages"`
	CacheVideos bool `json:"cacheVideos"`
	ImageQualities []string `json:"imageQualities"`
	MinAccessCountMessages float64 `json:"minAccessCountMessages"`
	MinAccessCountPosts float64 `json:"minAccessCountPosts"`
	MinAccessCountStories float64 `json:"minAccessCountStories"`
	MinAccessCountVault float64 `json:"minAccessCountVault"`
	PresignedUrlTtlSeconds float64 `json:"presignedUrlTtlSeconds"`
	RetentionDays float64 `json:"retentionDays"`
	StorageLimitBytes float64 `json:"storageLimitBytes"`
	StorageLimitPurgeStrategy string `json:"storageLimitPurgeStrategy"`
	VideoQualities []string `json:"videoQualities"`
} `json:"defaultSettings"`
} `json:"vaultPlus"`
} `json:"settings"`
}

// WhoamiResponse represents the response for Whoami
type WhoamiResponse struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Permissions []string `json:"permissions"`
}

// ListDynamicRulesResponse represents the response for ListDynamicRules
type ListDynamicRulesResponse struct {
	IsCurrent bool `json:"is_current"`
	IsEarlyAccess bool `json:"is_early_access"`
	IsPublic bool `json:"is_public"`
	Rules struct {
	AppToken string `json:"app_token"`
	ChecksumConstant float64 `json:"checksum_constant"`
	ChecksumIndexes []float64 `json:"checksum_indexes"`
	End string `json:"end"`
	Format string `json:"format"`
	Prefix string `json:"prefix"`
	Revision string `json:"revision"`
	Start string `json:"start"`
	StaticParam string `json:"static_param"`
	Suffix string `json:"suffix"`
} `json:"rules"`
}

// DynamicRulesSignResponse represents the response for DynamicRulesSign
type DynamicRulesSignResponse struct {
	IsEarlyAccess bool `json:"is_early_access"`
	IsPublic bool `json:"is_public"`
	Signed struct {
	AppToken string `json:"app-token"`
	Sign string `json:"sign"`
	Time string `json:"time"`
	UserId *string `json:"user-id,omitempty"`
} `json:"signed"`
}

// GetDynamicRulesStatusResponse represents the response for GetDynamicRulesStatus
type GetDynamicRulesStatusResponse struct {
	AccessGranted bool `json:"access_granted"`
	EarlyAccessRevision string `json:"early_access_revision"`
	IsCurrent bool `json:"is_current"`
	IsEarlyAccess bool `json:"is_early_access"`
	IsPublic bool `json:"is_public"`
	PublicRevision string `json:"public_revision"`
	Revision string `json:"revision"`
}

// LinkInitResponse represents the response for LinkInit
type LinkInitResponse struct {
	ExpiresAt string `json:"expiresAt"`
	Url string `json:"url"`
}

// GetLinkResponse represents the response for GetLink
type GetLinkResponse struct {
	Data *interface{} `json:"data,omitempty"`
	Status string `json:"status"`
}

// GetVaultPlusResponse represents the response for GetVaultPlus
type GetVaultPlusResponse struct {
	Duration float64 `json:"duration"`
	Id string `json:"id"`
	Media map[string]struct {
	AccessCount float64 `json:"accessCount"`
	ContentType string `json:"contentType"`
	CreatedAt float64 `json:"createdAt"`
	ExpiresAt float64 `json:"expiresAt"`
	Id string `json:"id"`
	LastAccessedAt float64 `json:"lastAccessedAt"`
	Quality string `json:"quality"`
	SizeBytes float64 `json:"sizeBytes"`
	Source string `json:"source"`
	Status string `json:"status"`
	StoredAt float64 `json:"storedAt"`
	Url string `json:"url"`
} `json:"media"`
	Type string `json:"type"`
}

// DeleteVaultPlusResponse represents the response for DeleteVaultPlus
type DeleteVaultPlusResponse struct {
	FreedBytes float64 `json:"freedBytes"`
	MediaId string `json:"mediaId"`
	Success bool `json:"success"`
}

// CreateVaultPlusBatchResponse represents the response for CreateVaultPlusBatch
type CreateVaultPlusBatchResponse struct {
	Items []struct {
	Duration float64 `json:"duration"`
	Id string `json:"id"`
	Media map[string]struct {
	AccessCount float64 `json:"accessCount"`
	ContentType string `json:"contentType"`
	CreatedAt float64 `json:"createdAt"`
	ExpiresAt float64 `json:"expiresAt"`
	Id string `json:"id"`
	LastAccessedAt float64 `json:"lastAccessedAt"`
	Quality string `json:"quality"`
	SizeBytes float64 `json:"sizeBytes"`
	Source string `json:"source"`
	Status string `json:"status"`
	StoredAt float64 `json:"storedAt"`
	Url string `json:"url"`
} `json:"media"`
	Type string `json:"type"`
} `json:"items"`
}

// GetVaultPlusListResponse represents the response for GetVaultPlusList
type GetVaultPlusListResponse struct {
	Items []struct {
	Duration float64 `json:"duration"`
	Id string `json:"id"`
	Media map[string]struct {
	AccessCount float64 `json:"accessCount"`
	ContentType string `json:"contentType"`
	CreatedAt float64 `json:"createdAt"`
	ExpiresAt float64 `json:"expiresAt"`
	Id string `json:"id"`
	LastAccessedAt float64 `json:"lastAccessedAt"`
	Quality string `json:"quality"`
	SizeBytes float64 `json:"sizeBytes"`
	Source string `json:"source"`
	Status string `json:"status"`
	StoredAt float64 `json:"storedAt"`
	Url string `json:"url"`
} `json:"media"`
	Type string `json:"type"`
} `json:"items"`
	NextCursor *string `json:"nextCursor,omitempty"`
}

// VaultPlusPurgeResponse represents the response for VaultPlusPurge
type VaultPlusPurgeResponse struct {
	FreedBytes float64 `json:"freedBytes"`
	PurgedCount float64 `json:"purgedCount"`
	Success bool `json:"success"`
}

// CreateVaultPlusStoreListResponse represents the response for CreateVaultPlusStoreList
type CreateVaultPlusStoreListResponse struct {
	EstimatedItems *float64 `json:"estimatedItems,omitempty"`
	ListId string `json:"listId"`
	Queued bool `json:"queued"`
}

// GetVaultPlusStoreStatsResponse represents the response for GetVaultPlusStoreStats
type GetVaultPlusStoreStatsResponse struct {
	ConnectionsWithVaultPlus float64 `json:"connectionsWithVaultPlus"`
	TotalConnections float64 `json:"totalConnections"`
	TotalMediaCount float64 `json:"totalMediaCount"`
	TotalStorageBytes float64 `json:"totalStorageBytes"`
}

// GetVaultPlusStoreStatusResponse represents the response for GetVaultPlusStoreStatus
type GetVaultPlusStoreStatusResponse struct {
	ConnectionId string `json:"connectionId"`
	PendingCount float64 `json:"pendingCount"`
	StoredCount float64 `json:"storedCount"`
	TotalMedia float64 `json:"totalMedia"`
	TotalSizeBytes float64 `json:"totalSizeBytes"`
}


// ============================================================================
// API Methods
// ============================================================================

// GetAccessAnalyticsCampaignsChart Campaigns chart
func (c *Client) GetAccessAnalyticsCampaignsChart(ctx context.Context, opts *RequestOptions) (*GetAccessAnalyticsCampaignsChartResponse, error) {
	path := "/v2/access/analytics/campaigns/chart"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result GetAccessAnalyticsCampaignsChartResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetAccessAnalyticsCampaignsTop Top campaigns
func (c *Client) GetAccessAnalyticsCampaignsTop(ctx context.Context, opts *RequestOptions) (*GetAccessAnalyticsCampaignsTopResponse, error) {
	path := "/v2/access/analytics/campaigns/top"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result GetAccessAnalyticsCampaignsTopResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ListAccessAnalyticsEarningsChargebacks Chargebacks
func (c *Client) ListAccessAnalyticsEarningsChargebacks(ctx context.Context, opts *RequestOptions) (*ListAccessAnalyticsEarningsChargebacksResponse, error) {
	path := "/v2/access/analytics/earnings/chargebacks"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result ListAccessAnalyticsEarningsChargebacksResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetAccessAnalyticsEarningsChart Earnings chart
func (c *Client) GetAccessAnalyticsEarningsChart(ctx context.Context, opts *RequestOptions) (*GetAccessAnalyticsEarningsChartResponse, error) {
	path := "/v2/access/analytics/earnings/chart"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result GetAccessAnalyticsEarningsChartResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ListAccessAnalyticsEarningsTransactions Transactions
func (c *Client) ListAccessAnalyticsEarningsTransactions(ctx context.Context, opts *RequestOptions) (*ListAccessAnalyticsEarningsTransactionsResponse, error) {
	path := "/v2/access/analytics/earnings/transactions"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result ListAccessAnalyticsEarningsTransactionsResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ListAccessAnalyticsMassMessagesBuyers Mass message buyers
func (c *Client) ListAccessAnalyticsMassMessagesBuyers(ctx context.Context, massMessageId string, opts *RequestOptions) (*ListAccessAnalyticsMassMessagesBuyersResponse, error) {
	path := fmt.Sprintf("/v2/access/analytics/mass-messages/%v/buyers", massMessageId)
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result ListAccessAnalyticsMassMessagesBuyersResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetAccessAnalyticsMassMessagesChart Mass messages chart
func (c *Client) GetAccessAnalyticsMassMessagesChart(ctx context.Context, opts *RequestOptions) (*GetAccessAnalyticsMassMessagesChartResponse, error) {
	path := "/v2/access/analytics/mass-messages/chart"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result GetAccessAnalyticsMassMessagesChartResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetAccessAnalyticsMassMessagesPurchased Mass messages purchased
func (c *Client) GetAccessAnalyticsMassMessagesPurchased(ctx context.Context, opts *RequestOptions) (*GetAccessAnalyticsMassMessagesPurchasedResponse, error) {
	path := "/v2/access/analytics/mass-messages/purchased"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result GetAccessAnalyticsMassMessagesPurchasedResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetAccessAnalyticsPosts Post stats
func (c *Client) GetAccessAnalyticsPosts(ctx context.Context, postId string) (*GetAccessAnalyticsPostsResponse, error) {
	path := fmt.Sprintf("/v2/access/analytics/posts/%v", postId)

	respBody, err := c.Request(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}

	var result GetAccessAnalyticsPostsResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetAccessAnalyticsPostsChart Posts chart
func (c *Client) GetAccessAnalyticsPostsChart(ctx context.Context, opts *RequestOptions) (*GetAccessAnalyticsPostsChartResponse, error) {
	path := "/v2/access/analytics/posts/chart"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result GetAccessAnalyticsPostsChartResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetAccessAnalyticsPostsTop Top posts
func (c *Client) GetAccessAnalyticsPostsTop(ctx context.Context, opts *RequestOptions) (*GetAccessAnalyticsPostsTopResponse, error) {
	path := "/v2/access/analytics/posts/top"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result GetAccessAnalyticsPostsTopResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetAccessAnalyticsPromotionsChart Promotions chart
func (c *Client) GetAccessAnalyticsPromotionsChart(ctx context.Context, opts *RequestOptions) (*GetAccessAnalyticsPromotionsChartResponse, error) {
	path := "/v2/access/analytics/promotions/chart"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result GetAccessAnalyticsPromotionsChartResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetAccessAnalyticsPromotionsTop Top promotions
func (c *Client) GetAccessAnalyticsPromotionsTop(ctx context.Context, opts *RequestOptions) (*GetAccessAnalyticsPromotionsTopResponse, error) {
	path := "/v2/access/analytics/promotions/top"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result GetAccessAnalyticsPromotionsTopResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetAccessAnalyticsStoriesChart Stories chart
func (c *Client) GetAccessAnalyticsStoriesChart(ctx context.Context, opts *RequestOptions) (*GetAccessAnalyticsStoriesChartResponse, error) {
	path := "/v2/access/analytics/stories/chart"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result GetAccessAnalyticsStoriesChartResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetAccessAnalyticsStoriesTop Top stories
func (c *Client) GetAccessAnalyticsStoriesTop(ctx context.Context, opts *RequestOptions) (*GetAccessAnalyticsStoriesTopResponse, error) {
	path := "/v2/access/analytics/stories/top"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result GetAccessAnalyticsStoriesTopResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetAccessAnalyticsStreamsChart Streams chart
func (c *Client) GetAccessAnalyticsStreamsChart(ctx context.Context, opts *RequestOptions) (*GetAccessAnalyticsStreamsChartResponse, error) {
	path := "/v2/access/analytics/streams/chart"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result GetAccessAnalyticsStreamsChartResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetAccessAnalyticsStreamsTop Top streams
func (c *Client) GetAccessAnalyticsStreamsTop(ctx context.Context, opts *RequestOptions) (*GetAccessAnalyticsStreamsTopResponse, error) {
	path := "/v2/access/analytics/streams/top"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result GetAccessAnalyticsStreamsTopResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetAccessAnalyticsTrialsChart Trials chart
func (c *Client) GetAccessAnalyticsTrialsChart(ctx context.Context, opts *RequestOptions) (*GetAccessAnalyticsTrialsChartResponse, error) {
	path := "/v2/access/analytics/trials/chart"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result GetAccessAnalyticsTrialsChartResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetAccessAnalyticsTrialsTop Top trials
func (c *Client) GetAccessAnalyticsTrialsTop(ctx context.Context, opts *RequestOptions) (*GetAccessAnalyticsTrialsTopResponse, error) {
	path := "/v2/access/analytics/trials/top"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result GetAccessAnalyticsTrialsTopResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetAccessAnalyticsVisitorCountriesChart Visitor countries chart
func (c *Client) GetAccessAnalyticsVisitorCountriesChart(ctx context.Context, opts *RequestOptions) (*GetAccessAnalyticsVisitorCountriesChartResponse, error) {
	path := "/v2/access/analytics/visitor-countries/chart"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result GetAccessAnalyticsVisitorCountriesChartResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetAccessAnalyticsVisitorCountriesTop Top visitor countries
func (c *Client) GetAccessAnalyticsVisitorCountriesTop(ctx context.Context, opts *RequestOptions) (*GetAccessAnalyticsVisitorCountriesTopResponse, error) {
	path := "/v2/access/analytics/visitor-countries/top"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result GetAccessAnalyticsVisitorCountriesTopResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ListAccessChats Chats list
func (c *Client) ListAccessChats(ctx context.Context, opts *RequestOptions) (*ListAccessChatsResponse, error) {
	path := "/v2/access/chats"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result ListAccessChatsResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ListAccessChatsMedia Get chat media
func (c *Client) ListAccessChatsMedia(ctx context.Context, userId string, opts *RequestOptions) (*ListAccessChatsMediaResponse, error) {
	path := fmt.Sprintf("/v2/access/chats/%v/media", userId)
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result ListAccessChatsMediaResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ListAccessChatsMessages Chat messages
func (c *Client) ListAccessChatsMessages(ctx context.Context, userId string, opts *RequestOptions) (*ListAccessChatsMessagesResponse, error) {
	path := fmt.Sprintf("/v2/access/chats/%v/messages", userId)
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result ListAccessChatsMessagesResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// CreateAccessChatsMessages Send chat message
func (c *Client) CreateAccessChatsMessages(ctx context.Context, userId string, body struct {
	IsForwardedMessage *bool `json:"isForwardedMessage,omitempty"`
	IsLockedText *bool `json:"isLockedText,omitempty"`
	IsMarkdown *bool `json:"isMarkdown,omitempty"`
	MediaItems []interface{} `json:"mediaItems,omitempty"`
	PreviewMediaCount *int64 `json:"previewMediaCount,omitempty"`
	Price *float64 `json:"price,omitempty"`
	ReleaseForms *struct {
	Guests []int64 `json:"guests,omitempty"`
	Partners []int64 `json:"partners,omitempty"`
	Users []int64 `json:"users,omitempty"`
} `json:"releaseForms,omitempty"`
	Text *string `json:"text,omitempty"`
	UserTags []int64 `json:"userTags,omitempty"`
}) (*CreateAccessChatsMessagesResponse, error) {
	path := fmt.Sprintf("/v2/access/chats/%v/messages", userId)
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}

	var result CreateAccessChatsMessagesResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// DeleteAccessChatsMessages Unsend chat message
func (c *Client) DeleteAccessChatsMessages(ctx context.Context, userId string, messageId string, body struct {
	WithUserId interface{} `json:"withUserId"`
}) (*DeleteAccessChatsMessagesResponse, error) {
	path := fmt.Sprintf("/v2/access/chats/%v/messages/%v", userId, messageId)
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}

	var result DeleteAccessChatsMessagesResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ListAccessEarningsChargebacks List chargebacks
func (c *Client) ListAccessEarningsChargebacks(ctx context.Context, opts *RequestOptions) (*ListAccessEarningsChargebacksResponse, error) {
	path := "/v2/access/earnings/chargebacks"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result ListAccessEarningsChargebacksResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetAccessEarningsChart Get earnings chart
func (c *Client) GetAccessEarningsChart(ctx context.Context, opts *RequestOptions) (*GetAccessEarningsChartResponse, error) {
	path := "/v2/access/earnings/chart"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result GetAccessEarningsChartResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ListAccessEarningsTransactions List transactions
func (c *Client) ListAccessEarningsTransactions(ctx context.Context, opts *RequestOptions) (*ListAccessEarningsTransactionsResponse, error) {
	path := "/v2/access/earnings/transactions"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result ListAccessEarningsTransactionsResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ListAccessMassMessages List mass messages
func (c *Client) ListAccessMassMessages(ctx context.Context, opts *RequestOptions) (*ListAccessMassMessagesResponse, error) {
	path := "/v2/access/mass-messages"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result ListAccessMassMessagesResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// CreateAccessMassMessages Create mass message
func (c *Client) CreateAccessMassMessages(ctx context.Context, body struct {
	ExcludeUserLists []interface{} `json:"excludeUserLists,omitempty"`
	IsForwardedMessage *bool `json:"isForwardedMessage,omitempty"`
	IsLockedText *bool `json:"isLockedText,omitempty"`
	IsMarkdown *bool `json:"isMarkdown,omitempty"`
	MediaItems []interface{} `json:"mediaItems,omitempty"`
	PreviewMediaCount *int64 `json:"previewMediaCount,omitempty"`
	Price *float64 `json:"price,omitempty"`
	ReleaseForms *struct {
	Guests []int64 `json:"guests,omitempty"`
	Partners []int64 `json:"partners,omitempty"`
	Users []int64 `json:"users,omitempty"`
} `json:"releaseForms,omitempty"`
	ScheduledDate *interface{} `json:"scheduledDate,omitempty"`
	SubscribedAfterDate *interface{} `json:"subscribedAfterDate,omitempty"`
	Text *string `json:"text,omitempty"`
	UserIds []interface{} `json:"userIds,omitempty"`
	UserLists []interface{} `json:"userLists,omitempty"`
	UserTags []int64 `json:"userTags,omitempty"`
}) (*CreateAccessMassMessagesResponse, error) {
	path := "/v2/access/mass-messages"
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}

	var result CreateAccessMassMessagesResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetAccessMassMessages Get mass message
func (c *Client) GetAccessMassMessages(ctx context.Context, massMessageId string) (*GetAccessMassMessagesResponse, error) {
	path := fmt.Sprintf("/v2/access/mass-messages/%v", massMessageId)

	respBody, err := c.Request(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}

	var result GetAccessMassMessagesResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ReplaceAccessMassMessages Update mass message
func (c *Client) ReplaceAccessMassMessages(ctx context.Context, massMessageId string, body struct {
	ExcludeUserLists []interface{} `json:"excludeUserLists,omitempty"`
	IsForwardedMessage *bool `json:"isForwardedMessage,omitempty"`
	IsLockedText *bool `json:"isLockedText,omitempty"`
	IsMarkdown *bool `json:"isMarkdown,omitempty"`
	MediaItems []interface{} `json:"mediaItems,omitempty"`
	PreviewMediaCount *int64 `json:"previewMediaCount,omitempty"`
	Price *float64 `json:"price,omitempty"`
	ReleaseForms *struct {
	Guests []int64 `json:"guests,omitempty"`
	Partners []int64 `json:"partners,omitempty"`
	Users []int64 `json:"users,omitempty"`
} `json:"releaseForms,omitempty"`
	ScheduledDate *interface{} `json:"scheduledDate,omitempty"`
	SubscribedAfterDate *interface{} `json:"subscribedAfterDate,omitempty"`
	Text *string `json:"text,omitempty"`
	UserIds []interface{} `json:"userIds,omitempty"`
	UserLists []interface{} `json:"userLists,omitempty"`
	UserTags []int64 `json:"userTags,omitempty"`
}) (*ReplaceAccessMassMessagesResponse, error) {
	path := fmt.Sprintf("/v2/access/mass-messages/%v", massMessageId)
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "PUT", path, opts)
	if err != nil {
		return nil, err
	}

	var result ReplaceAccessMassMessagesResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// DeleteAccessMassMessages Delete mass message
func (c *Client) DeleteAccessMassMessages(ctx context.Context, massMessageId string) (*DeleteAccessMassMessagesResponse, error) {
	path := fmt.Sprintf("/v2/access/mass-messages/%v", massMessageId)

	respBody, err := c.Request(ctx, "DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	var result DeleteAccessMassMessagesResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ListAccessPosts List own posts
func (c *Client) ListAccessPosts(ctx context.Context, opts *RequestOptions) (*ListAccessPostsResponse, error) {
	path := "/v2/access/posts"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result ListAccessPostsResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// CreateAccessPosts Create post
func (c *Client) CreateAccessPosts(ctx context.Context, body struct {
	ExpireAfter *float64 `json:"expireAfter,omitempty"`
	FundRaisingTargetAmount *float64 `json:"fundRaisingTargetAmount,omitempty"`
	FundRaisingTipsPresets []float64 `json:"fundRaisingTipsPresets,omitempty"`
	IsLockedText *bool `json:"isLockedText,omitempty"`
	IsMarkdown *bool `json:"isMarkdown,omitempty"`
	MediaItems []interface{} `json:"mediaItems,omitempty"`
	PreviewMediaCount *int64 `json:"previewMediaCount,omitempty"`
	Price *float64 `json:"price,omitempty"`
	ReleaseForms *struct {
	Guests []int64 `json:"guests,omitempty"`
	Partners []int64 `json:"partners,omitempty"`
	Users []int64 `json:"users,omitempty"`
} `json:"releaseForms,omitempty"`
	ScheduledDate *string `json:"scheduledDate,omitempty"`
	Text *string `json:"text,omitempty"`
	UserTags []int64 `json:"userTags,omitempty"`
}) (*CreateAccessPostsResponse, error) {
	path := "/v2/access/posts"
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}

	var result CreateAccessPostsResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetAccessPosts Get post
func (c *Client) GetAccessPosts(ctx context.Context, postId string) (*map[string]interface{}, error) {
	path := fmt.Sprintf("/v2/access/posts/%v", postId)

	respBody, err := c.Request(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ReplaceAccessPosts Edit post
func (c *Client) ReplaceAccessPosts(ctx context.Context, postId string, body struct {
	ExpireAfter *float64 `json:"expireAfter,omitempty"`
	FundRaisingTargetAmount *float64 `json:"fundRaisingTargetAmount,omitempty"`
	FundRaisingTipsPresets []float64 `json:"fundRaisingTipsPresets,omitempty"`
	IsLockedText *bool `json:"isLockedText,omitempty"`
	IsMarkdown *bool `json:"isMarkdown,omitempty"`
	MediaItems []interface{} `json:"mediaItems,omitempty"`
	PreviewMediaCount *int64 `json:"previewMediaCount,omitempty"`
	Price *float64 `json:"price,omitempty"`
	ReleaseForms *struct {
	Guests []int64 `json:"guests,omitempty"`
	Partners []int64 `json:"partners,omitempty"`
	Users []int64 `json:"users,omitempty"`
} `json:"releaseForms,omitempty"`
	ScheduledDate *string `json:"scheduledDate,omitempty"`
	Text *string `json:"text,omitempty"`
	UserTags []int64 `json:"userTags,omitempty"`
}) (*ReplaceAccessPostsResponse, error) {
	path := fmt.Sprintf("/v2/access/posts/%v", postId)
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "PUT", path, opts)
	if err != nil {
		return nil, err
	}

	var result ReplaceAccessPostsResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// DeleteAccessPosts Delete post
func (c *Client) DeleteAccessPosts(ctx context.Context, postId string) (*DeleteAccessPostsResponse, error) {
	path := fmt.Sprintf("/v2/access/posts/%v", postId)

	respBody, err := c.Request(ctx, "DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	var result DeleteAccessPostsResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ListAccessPromotions List promotions
func (c *Client) ListAccessPromotions(ctx context.Context, opts *RequestOptions) (*map[string]interface{}, error) {
	path := "/v2/access/promotions"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// CreateAccessPromotions Create promotion
func (c *Client) CreateAccessPromotions(ctx context.Context, body struct {
	Discount int64 `json:"discount"`
	FinishDays int64 `json:"finishDays"`
	Message string `json:"message"`
	SubscribeCounts int64 `json:"subscribeCounts"`
	SubscribeDays int64 `json:"subscribeDays"`
	Type []string `json:"type"`
}) (*map[string]interface{}, error) {
	path := "/v2/access/promotions"
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ReplaceAccessPromotions Update promotion
func (c *Client) ReplaceAccessPromotions(ctx context.Context, promotionId string, body struct {
	Discount *int64 `json:"discount,omitempty"`
	FinishDays *int64 `json:"finishDays,omitempty"`
	Message *string `json:"message,omitempty"`
	SubscribeCounts *int64 `json:"subscribeCounts,omitempty"`
	SubscribeDays *int64 `json:"subscribeDays,omitempty"`
	Type []string `json:"type,omitempty"`
}) (*ReplaceAccessPromotionsResponse, error) {
	path := fmt.Sprintf("/v2/access/promotions/%v", promotionId)
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "PUT", path, opts)
	if err != nil {
		return nil, err
	}

	var result ReplaceAccessPromotionsResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// DeleteAccessPromotions Delete promotion
func (c *Client) DeleteAccessPromotions(ctx context.Context, promotionId string) (*DeleteAccessPromotionsResponse, error) {
	path := fmt.Sprintf("/v2/access/promotions/%v", promotionId)

	respBody, err := c.Request(ctx, "DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	var result DeleteAccessPromotionsResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// CreateAccessPromotionsStop Stop promotion
func (c *Client) CreateAccessPromotionsStop(ctx context.Context, promotionId string) (*CreateAccessPromotionsStopResponse, error) {
	path := fmt.Sprintf("/v2/access/promotions/%v/stop", promotionId)

	respBody, err := c.Request(ctx, "POST", path, nil)
	if err != nil {
		return nil, err
	}

	var result CreateAccessPromotionsStopResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ListAccessPromotionsBundles List bundles
func (c *Client) ListAccessPromotionsBundles(ctx context.Context, opts *RequestOptions) (*map[string]interface{}, error) {
	path := "/v2/access/promotions/bundles"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// CreateAccessPromotionsBundles Create bundle
func (c *Client) CreateAccessPromotionsBundles(ctx context.Context, body struct {
	Discount int64 `json:"discount"`
	Duration int64 `json:"duration"`
}) (*CreateAccessPromotionsBundlesResponse, error) {
	path := "/v2/access/promotions/bundles"
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}

	var result CreateAccessPromotionsBundlesResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetAccessPromotionsBundles Get bundle
func (c *Client) GetAccessPromotionsBundles(ctx context.Context, bundleId string) (*GetAccessPromotionsBundlesResponse, error) {
	path := fmt.Sprintf("/v2/access/promotions/bundles/%v", bundleId)

	respBody, err := c.Request(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}

	var result GetAccessPromotionsBundlesResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ReplaceAccessPromotionsBundles Update bundle
func (c *Client) ReplaceAccessPromotionsBundles(ctx context.Context, bundleId string, body struct {
	Discount *int64 `json:"discount,omitempty"`
	Duration *int64 `json:"duration,omitempty"`
}) (*ReplaceAccessPromotionsBundlesResponse, error) {
	path := fmt.Sprintf("/v2/access/promotions/bundles/%v", bundleId)
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "PUT", path, opts)
	if err != nil {
		return nil, err
	}

	var result ReplaceAccessPromotionsBundlesResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// DeleteAccessPromotionsBundles Delete bundle
func (c *Client) DeleteAccessPromotionsBundles(ctx context.Context, bundleId string) (*DeleteAccessPromotionsBundlesResponse, error) {
	path := fmt.Sprintf("/v2/access/promotions/bundles/%v", bundleId)

	respBody, err := c.Request(ctx, "DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	var result DeleteAccessPromotionsBundlesResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ListAccessPromotionsTrackingLinks List tracking links
func (c *Client) ListAccessPromotionsTrackingLinks(ctx context.Context, opts *RequestOptions) (*ListAccessPromotionsTrackingLinksResponse, error) {
	path := "/v2/access/promotions/tracking-links"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result ListAccessPromotionsTrackingLinksResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// CreateAccessPromotionsTrackingLinks Create tracking link
func (c *Client) CreateAccessPromotionsTrackingLinks(ctx context.Context, body struct {
	Name string `json:"name"`
}) (*map[string]interface{}, error) {
	path := "/v2/access/promotions/tracking-links"
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetAccessPromotionsTrackingLinks Get tracking link
func (c *Client) GetAccessPromotionsTrackingLinks(ctx context.Context, trackingLinkId string) (*GetAccessPromotionsTrackingLinksResponse, error) {
	path := fmt.Sprintf("/v2/access/promotions/tracking-links/%v", trackingLinkId)

	respBody, err := c.Request(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}

	var result GetAccessPromotionsTrackingLinksResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ReplaceAccessPromotionsTrackingLinks Update tracking link
func (c *Client) ReplaceAccessPromotionsTrackingLinks(ctx context.Context, trackingLinkId string, body struct {
	Name *string `json:"name,omitempty"`
}) (*ReplaceAccessPromotionsTrackingLinksResponse, error) {
	path := fmt.Sprintf("/v2/access/promotions/tracking-links/%v", trackingLinkId)
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "PUT", path, opts)
	if err != nil {
		return nil, err
	}

	var result ReplaceAccessPromotionsTrackingLinksResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// DeleteAccessPromotionsTrackingLinks Delete tracking link
func (c *Client) DeleteAccessPromotionsTrackingLinks(ctx context.Context, trackingLinkId string) (*DeleteAccessPromotionsTrackingLinksResponse, error) {
	path := fmt.Sprintf("/v2/access/promotions/tracking-links/%v", trackingLinkId)

	respBody, err := c.Request(ctx, "DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	var result DeleteAccessPromotionsTrackingLinksResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ListAccessPromotionsTrackingLinksClaimers Get tracking link claimers
func (c *Client) ListAccessPromotionsTrackingLinksClaimers(ctx context.Context, trackingLinkId string) (*ListAccessPromotionsTrackingLinksClaimersResponse, error) {
	path := fmt.Sprintf("/v2/access/promotions/tracking-links/%v/claimers", trackingLinkId)

	respBody, err := c.Request(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}

	var result ListAccessPromotionsTrackingLinksClaimersResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// CreateAccessPromotionsTrackingLinksShareAccess Share tracking link access
func (c *Client) CreateAccessPromotionsTrackingLinksShareAccess(ctx context.Context, body struct {
	CampaignId int64 `json:"campaignId"`
	UserId int64 `json:"userId"`
}) (*CreateAccessPromotionsTrackingLinksShareAccessResponse, error) {
	path := "/v2/access/promotions/tracking-links/share-access"
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}

	var result CreateAccessPromotionsTrackingLinksShareAccessResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// DeleteAccessPromotionsTrackingLinksShareAccess Revoke tracking link access
func (c *Client) DeleteAccessPromotionsTrackingLinksShareAccess(ctx context.Context, body struct {
	CampaignId int64 `json:"campaignId"`
	UserId int64 `json:"userId"`
}) (*DeleteAccessPromotionsTrackingLinksShareAccessResponse, error) {
	path := "/v2/access/promotions/tracking-links/share-access"
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}

	var result DeleteAccessPromotionsTrackingLinksShareAccessResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ListAccessPromotionsTrialLinks List trial links
func (c *Client) ListAccessPromotionsTrialLinks(ctx context.Context, opts *RequestOptions) (*map[string]interface{}, error) {
	path := "/v2/access/promotions/trial-links"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// CreateAccessPromotionsTrialLinks Create trial link
func (c *Client) CreateAccessPromotionsTrialLinks(ctx context.Context, body struct {
	ExpiredAt *string `json:"expiredAt,omitempty"`
	SubscribeCounts *int64 `json:"subscribeCounts,omitempty"`
	SubscribeDays int64 `json:"subscribeDays"`
	TrialLinkName string `json:"trialLinkName"`
}) (*CreateAccessPromotionsTrialLinksResponse, error) {
	path := "/v2/access/promotions/trial-links"
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}

	var result CreateAccessPromotionsTrialLinksResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetAccessPromotionsTrialLinks Get trial link
func (c *Client) GetAccessPromotionsTrialLinks(ctx context.Context, trialLinkId string) (*GetAccessPromotionsTrialLinksResponse, error) {
	path := fmt.Sprintf("/v2/access/promotions/trial-links/%v", trialLinkId)

	respBody, err := c.Request(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}

	var result GetAccessPromotionsTrialLinksResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ReplaceAccessPromotionsTrialLinks Update trial link
func (c *Client) ReplaceAccessPromotionsTrialLinks(ctx context.Context, trialLinkId string, body struct {
	ExpiredAt *string `json:"expiredAt,omitempty"`
	SubscribeCounts *int64 `json:"subscribeCounts,omitempty"`
	SubscribeDays *int64 `json:"subscribeDays,omitempty"`
	TrialLinkName *string `json:"trialLinkName,omitempty"`
}) (*ReplaceAccessPromotionsTrialLinksResponse, error) {
	path := fmt.Sprintf("/v2/access/promotions/trial-links/%v", trialLinkId)
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "PUT", path, opts)
	if err != nil {
		return nil, err
	}

	var result ReplaceAccessPromotionsTrialLinksResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// DeleteAccessPromotionsTrialLinks Delete trial link
func (c *Client) DeleteAccessPromotionsTrialLinks(ctx context.Context, trialLinkId string) (*DeleteAccessPromotionsTrialLinksResponse, error) {
	path := fmt.Sprintf("/v2/access/promotions/trial-links/%v", trialLinkId)

	respBody, err := c.Request(ctx, "DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	var result DeleteAccessPromotionsTrialLinksResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// CreateAccessPromotionsTrialLinksShareAccess Share trial link access
func (c *Client) CreateAccessPromotionsTrialLinksShareAccess(ctx context.Context, body struct {
	TrialId int64 `json:"trialId"`
	UserId int64 `json:"userId"`
}) (*CreateAccessPromotionsTrialLinksShareAccessResponse, error) {
	path := "/v2/access/promotions/trial-links/share-access"
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}

	var result CreateAccessPromotionsTrialLinksShareAccessResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// DeleteAccessPromotionsTrialLinksShareAccess Revoke trial link access
func (c *Client) DeleteAccessPromotionsTrialLinksShareAccess(ctx context.Context, body struct {
	TrialId int64 `json:"trialId"`
	UserId int64 `json:"userId"`
}) (*DeleteAccessPromotionsTrialLinksShareAccessResponse, error) {
	path := "/v2/access/promotions/trial-links/share-access"
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}

	var result DeleteAccessPromotionsTrialLinksShareAccessResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetAccessSelf Get current user
func (c *Client) GetAccessSelf(ctx context.Context) (*map[string]interface{}, error) {
	path := "/v2/access/self"

	respBody, err := c.Request(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// UpdateAccessSelf Update current user profile
func (c *Client) UpdateAccessSelf(ctx context.Context, body struct {
	About *string `json:"about,omitempty"`
	Name *string `json:"name,omitempty"`
}) (*UpdateAccessSelfResponse, error) {
	path := "/v2/access/self"
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "PATCH", path, opts)
	if err != nil {
		return nil, err
	}

	var result UpdateAccessSelfResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ListAccessSelfNotifications List notifications
func (c *Client) ListAccessSelfNotifications(ctx context.Context, opts *RequestOptions) (*ListAccessSelfNotificationsResponse, error) {
	path := "/v2/access/self/notifications"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result ListAccessSelfNotificationsResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ListAccessSelfReleaseForms List release forms
func (c *Client) ListAccessSelfReleaseForms(ctx context.Context, opts *RequestOptions) (*ListAccessSelfReleaseFormsResponse, error) {
	path := "/v2/access/self/release-forms"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result ListAccessSelfReleaseFormsResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ListAccessSelfTaggedFriendUsers List tagged friend users
func (c *Client) ListAccessSelfTaggedFriendUsers(ctx context.Context, opts *RequestOptions) (*ListAccessSelfTaggedFriendUsersResponse, error) {
	path := "/v2/access/self/tagged-friend-users"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result ListAccessSelfTaggedFriendUsersResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ListAccessSubscribers List subscribers
func (c *Client) ListAccessSubscribers(ctx context.Context, opts *RequestOptions) (*ListAccessSubscribersResponse, error) {
	path := "/v2/access/subscribers"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result ListAccessSubscribersResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// SetAccessSubscribersCustomName Set custom name for subscriber
func (c *Client) SetAccessSubscribersCustomName(ctx context.Context, userId string, body struct {
	DisplayName string `json:"displayName"`
}) (*SetAccessSubscribersCustomNameResponse, error) {
	path := fmt.Sprintf("/v2/access/subscribers/%v/custom-name", userId)
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "PUT", path, opts)
	if err != nil {
		return nil, err
	}

	var result SetAccessSubscribersCustomNameResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// SetAccessSubscribersDiscount Apply discount to subscriber
func (c *Client) SetAccessSubscribersDiscount(ctx context.Context, userId string, body struct {
	Discount int64 `json:"discount"`
	Period int64 `json:"period"`
}) (*SetAccessSubscribersDiscountResponse, error) {
	path := fmt.Sprintf("/v2/access/subscribers/%v/discount", userId)
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "PUT", path, opts)
	if err != nil {
		return nil, err
	}

	var result SetAccessSubscribersDiscountResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// SetAccessSubscribersNote Update subscriber note
func (c *Client) SetAccessSubscribersNote(ctx context.Context, userId string, body struct {
	Notice string `json:"notice"`
}) (*SetAccessSubscribersNoteResponse, error) {
	path := fmt.Sprintf("/v2/access/subscribers/%v/note", userId)
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "PUT", path, opts)
	if err != nil {
		return nil, err
	}

	var result SetAccessSubscribersNoteResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ListAccessSubscriptions List subscriptions
func (c *Client) ListAccessSubscriptions(ctx context.Context, opts *RequestOptions) (*ListAccessSubscriptionsResponse, error) {
	path := "/v2/access/subscriptions"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result ListAccessSubscriptionsResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetAccessSubscriptionsHistory Get subscription history
func (c *Client) GetAccessSubscriptionsHistory(ctx context.Context, subscriptionId string, opts *RequestOptions) (*GetAccessSubscriptionsHistoryResponse, error) {
	path := fmt.Sprintf("/v2/access/subscriptions/%v/history", subscriptionId)
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result GetAccessSubscriptionsHistoryResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetAccessSubscriptionsCount Get subscription counts
func (c *Client) GetAccessSubscriptionsCount(ctx context.Context) (*GetAccessSubscriptionsCountResponse, error) {
	path := "/v2/access/subscriptions/count"

	respBody, err := c.Request(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}

	var result GetAccessSubscriptionsCountResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ReplaceAccessUploads Upload single-part media and finalize (No need to call /complete after upload if using this endpoint)
func (c *Client) ReplaceAccessUploads(ctx context.Context, mediaUploadId string) (*ReplaceAccessUploadsResponse, error) {
	path := fmt.Sprintf("/v2/access/uploads/%v", mediaUploadId)

	respBody, err := c.Request(ctx, "PUT", path, nil)
	if err != nil {
		return nil, err
	}

	var result ReplaceAccessUploadsResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ReplaceAccessUploadsParts Upload chunk to managed media upload
func (c *Client) ReplaceAccessUploadsParts(ctx context.Context, mediaUploadId string, partNumber int64) (*ReplaceAccessUploadsPartsResponse, error) {
	path := fmt.Sprintf("/v2/access/uploads/%v/parts/%v", mediaUploadId, partNumber)

	respBody, err := c.Request(ctx, "PUT", path, nil)
	if err != nil {
		return nil, err
	}

	var result ReplaceAccessUploadsPartsResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// AccessUploadsCheck Check if media already exists in vault
func (c *Client) AccessUploadsCheck(ctx context.Context, body struct {
	Etag string `json:"etag"`
	Size int64 `json:"size"`
}) (*AccessUploadsCheckResponse, error) {
	path := "/v2/access/uploads/check"
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}

	var result AccessUploadsCheckResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// AccessUploadsComplete Finalize media upload
func (c *Client) AccessUploadsComplete(ctx context.Context, body struct {
	MediaUploadId string `json:"mediaUploadId"`
}) (*AccessUploadsCompleteResponse, error) {
	path := "/v2/access/uploads/complete"
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}

	var result AccessUploadsCompleteResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// AccessUploadsInit Initialize media upload
func (c *Client) AccessUploadsInit(ctx context.Context, body struct {
	ContentType string `json:"contentType"`
	Filename string `json:"filename"`
	Size int64 `json:"size"`
	VaultUpload *struct {
	Mode *string `json:"mode,omitempty"`
	UserId *string `json:"userId,omitempty"`
} `json:"vaultUpload,omitempty"`
}) (*AccessUploadsInitResponse, error) {
	path := "/v2/access/uploads/init"
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}

	var result AccessUploadsInitResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetAccessUsers Get user
func (c *Client) GetAccessUsers(ctx context.Context, userId string) (*map[string]interface{}, error) {
	path := fmt.Sprintf("/v2/access/users/%v", userId)

	respBody, err := c.Request(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// CreateAccessUsersLists Add user to multiple lists
func (c *Client) CreateAccessUsersLists(ctx context.Context, userId string, body struct {
	ListIds []int64 `json:"listIds"`
}) (*CreateAccessUsersListsResponse, error) {
	path := fmt.Sprintf("/v2/access/users/%v/lists", userId)
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}

	var result CreateAccessUsersListsResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ListAccessUsersPosts List user posts
func (c *Client) ListAccessUsersPosts(ctx context.Context, userId string, opts *RequestOptions) (*ListAccessUsersPostsResponse, error) {
	path := fmt.Sprintf("/v2/access/users/%v/posts", userId)
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result ListAccessUsersPostsResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// CreateAccessUsersRestrict Restrict user
func (c *Client) CreateAccessUsersRestrict(ctx context.Context, userId string) (*CreateAccessUsersRestrictResponse, error) {
	path := fmt.Sprintf("/v2/access/users/%v/restrict", userId)

	respBody, err := c.Request(ctx, "POST", path, nil)
	if err != nil {
		return nil, err
	}

	var result CreateAccessUsersRestrictResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// DeleteAccessUsersRestrict Unrestrict user
func (c *Client) DeleteAccessUsersRestrict(ctx context.Context, userId string) (*DeleteAccessUsersRestrictResponse, error) {
	path := fmt.Sprintf("/v2/access/users/%v/restrict", userId)

	respBody, err := c.Request(ctx, "DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	var result DeleteAccessUsersRestrictResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetAccessUsersBlocked List blocked users
func (c *Client) GetAccessUsersBlocked(ctx context.Context, opts *RequestOptions) (*GetAccessUsersBlockedResponse, error) {
	path := "/v2/access/users/blocked"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result GetAccessUsersBlockedResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetAccessUsersList List users by IDs
func (c *Client) GetAccessUsersList(ctx context.Context, opts *RequestOptions) (*GetAccessUsersListResponse, error) {
	path := "/v2/access/users/list"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result GetAccessUsersListResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ListAccessUsersLists List user lists
func (c *Client) ListAccessUsersLists(ctx context.Context, opts *RequestOptions) (*ListAccessUsersListsResponse, error) {
	path := "/v2/access/users/lists"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result ListAccessUsersListsResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// CreateAccessUsersLists1 Create user list
func (c *Client) CreateAccessUsersLists1(ctx context.Context, body struct {
	Name string `json:"name"`
}) (*CreateAccessUsersLists1Response, error) {
	path := "/v2/access/users/lists"
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}

	var result CreateAccessUsersLists1Response
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetAccessUsersLists Get user list
func (c *Client) GetAccessUsersLists(ctx context.Context, listId string) (*GetAccessUsersListsResponse, error) {
	path := fmt.Sprintf("/v2/access/users/lists/%v", listId)

	respBody, err := c.Request(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}

	var result GetAccessUsersListsResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// UpdateAccessUsersLists Update user list
func (c *Client) UpdateAccessUsersLists(ctx context.Context, listId string, body struct {
	Name string `json:"name"`
}) (*UpdateAccessUsersListsResponse, error) {
	path := fmt.Sprintf("/v2/access/users/lists/%v", listId)
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "PATCH", path, opts)
	if err != nil {
		return nil, err
	}

	var result UpdateAccessUsersListsResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// DeleteAccessUsersLists Delete user list
func (c *Client) DeleteAccessUsersLists(ctx context.Context, listId string) (*DeleteAccessUsersListsResponse, error) {
	path := fmt.Sprintf("/v2/access/users/lists/%v", listId)

	respBody, err := c.Request(ctx, "DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	var result DeleteAccessUsersListsResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ListAccessUsersListsUsers List users in user list
func (c *Client) ListAccessUsersListsUsers(ctx context.Context, listId string, opts *RequestOptions) (*ListAccessUsersListsUsersResponse, error) {
	path := fmt.Sprintf("/v2/access/users/lists/%v/users", listId)
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result ListAccessUsersListsUsersResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// CreateAccessUsersListsUsers Add user to list
func (c *Client) CreateAccessUsersListsUsers(ctx context.Context, listId string, userId string) (*map[string]interface{}, error) {
	path := fmt.Sprintf("/v2/access/users/lists/%v/users/%v", listId, userId)

	respBody, err := c.Request(ctx, "POST", path, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// DeleteAccessUsersListsUsers Remove user from user list
func (c *Client) DeleteAccessUsersListsUsers(ctx context.Context, listId string, userId string) (*map[string]interface{}, error) {
	path := fmt.Sprintf("/v2/access/users/lists/%v/users/%v", listId, userId)

	respBody, err := c.Request(ctx, "DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetAccessUsersRestrict List restricted users
func (c *Client) GetAccessUsersRestrict(ctx context.Context, opts *RequestOptions) (*GetAccessUsersRestrictResponse, error) {
	path := "/v2/access/users/restrict"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result GetAccessUsersRestrictResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// AccessUsersSearch Search performers
func (c *Client) AccessUsersSearch(ctx context.Context, opts *RequestOptions) (*map[string]interface{}, error) {
	path := "/v2/access/users/search"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ListAccessVaultLists List vault folders
func (c *Client) ListAccessVaultLists(ctx context.Context, opts *RequestOptions) (*ListAccessVaultListsResponse, error) {
	path := "/v2/access/vault/lists"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result ListAccessVaultListsResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// CreateAccessVaultLists Create vault list
func (c *Client) CreateAccessVaultLists(ctx context.Context, body struct {
	Name string `json:"name"`
}) (*CreateAccessVaultListsResponse, error) {
	path := "/v2/access/vault/lists"
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}

	var result CreateAccessVaultListsResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// UpdateAccessVaultLists Update vault list
func (c *Client) UpdateAccessVaultLists(ctx context.Context, listId string, body struct {
	ClearMedia *bool `json:"clearMedia,omitempty"`
	Name string `json:"name"`
}) (*UpdateAccessVaultListsResponse, error) {
	path := fmt.Sprintf("/v2/access/vault/lists/%v", listId)
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "PATCH", path, opts)
	if err != nil {
		return nil, err
	}

	var result UpdateAccessVaultListsResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// DeleteAccessVaultLists Delete vault list
func (c *Client) DeleteAccessVaultLists(ctx context.Context, listId string) (*DeleteAccessVaultListsResponse, error) {
	path := fmt.Sprintf("/v2/access/vault/lists/%v", listId)

	respBody, err := c.Request(ctx, "DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	var result DeleteAccessVaultListsResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ListAccessVaultListsMedia List media in vault list
func (c *Client) ListAccessVaultListsMedia(ctx context.Context, listId string, opts *RequestOptions) (*ListAccessVaultListsMediaResponse, error) {
	path := fmt.Sprintf("/v2/access/vault/lists/%v/media", listId)
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result ListAccessVaultListsMediaResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// CreateAccessVaultListsMedia Add media to vault list
func (c *Client) CreateAccessVaultListsMedia(ctx context.Context, listId string, body struct {
	MediaIds []int64 `json:"mediaIds"`
}) (*CreateAccessVaultListsMediaResponse, error) {
	path := fmt.Sprintf("/v2/access/vault/lists/%v/media", listId)
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}

	var result CreateAccessVaultListsMediaResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ListAccessVaultMedia List vault media
func (c *Client) ListAccessVaultMedia(ctx context.Context, opts *RequestOptions) (*ListAccessVaultMediaResponse, error) {
	path := "/v2/access/vault/media"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result ListAccessVaultMediaResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ListAccountConnections List connections
func (c *Client) ListAccountConnections(ctx context.Context, opts *RequestOptions) (*ListAccountConnectionsResponse, error) {
	path := "/v2/account/connections"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result ListAccountConnectionsResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// CreateAccountConnectionsImport Import connection
func (c *Client) CreateAccountConnectionsImport(ctx context.Context, body struct {
	ClientReferenceId *string `json:"clientReferenceId,omitempty"`
	Cookie string `json:"cookie"`
	Permissions []string `json:"permissions,omitempty"`
	UserAgent string `json:"userAgent"`
}) (*CreateAccountConnectionsImportResponse, error) {
	path := "/v2/account/connections/import"
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}

	var result CreateAccountConnectionsImportResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// UpdateAccountConnectionsImport Update imported connection session
func (c *Client) UpdateAccountConnectionsImport(ctx context.Context, connectionId string, body struct {
	Cookie string `json:"cookie"`
	UserAgent string `json:"userAgent"`
}) (*UpdateAccountConnectionsImportResponse, error) {
	path := fmt.Sprintf("/v2/account/connections/import/%v", connectionId)
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "PATCH", path, opts)
	if err != nil {
		return nil, err
	}

	var result UpdateAccountConnectionsImportResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// DeleteAccountConnections Disconnect connection
func (c *Client) DeleteAccountConnections(ctx context.Context, connectionId string) (*map[string]interface{}, error) {
	path := fmt.Sprintf("/v2/account/connections/%v", connectionId)

	respBody, err := c.Request(ctx, "DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// AccountConnectionsInvalidate Invalidate connection
func (c *Client) AccountConnectionsInvalidate(ctx context.Context, connectionId string) (*map[string]interface{}, error) {
	path := fmt.Sprintf("/v2/account/connections/%v/invalidate", connectionId)

	respBody, err := c.Request(ctx, "POST", path, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetAccountConnectionsSettings Get connection settings
func (c *Client) GetAccountConnectionsSettings(ctx context.Context, connectionId string) (*GetAccountConnectionsSettingsResponse, error) {
	path := fmt.Sprintf("/v2/account/connections/%v/settings", connectionId)

	respBody, err := c.Request(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}

	var result GetAccountConnectionsSettingsResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// UpdateAccountConnectionsSettings Update connection settings
func (c *Client) UpdateAccountConnectionsSettings(ctx context.Context, connectionId string, body struct {
	VaultPlus *struct {
	Enabled *bool `json:"enabled,omitempty"`
	Settings *struct {
	AccessExpiryDays *float64 `json:"accessExpiryDays,omitempty"`
	AutoCacheMessages *bool `json:"autoCacheMessages,omitempty"`
	AutoCachePosts *bool `json:"autoCachePosts,omitempty"`
	AutoCacheStories *bool `json:"autoCacheStories,omitempty"`
	AutoCacheVault *bool `json:"autoCacheVault,omitempty"`
	CacheAudio *bool `json:"cacheAudio,omitempty"`
	CacheImages *bool `json:"cacheImages,omitempty"`
	CacheVideos *bool `json:"cacheVideos,omitempty"`
	ImageQualities []string `json:"imageQualities,omitempty"`
	MinAccessCountMessages *float64 `json:"minAccessCountMessages,omitempty"`
	MinAccessCountPosts *float64 `json:"minAccessCountPosts,omitempty"`
	MinAccessCountStories *float64 `json:"minAccessCountStories,omitempty"`
	MinAccessCountVault *float64 `json:"minAccessCountVault,omitempty"`
	PresignedUrlTtlSeconds *float64 `json:"presignedUrlTtlSeconds,omitempty"`
	RetentionDays *float64 `json:"retentionDays,omitempty"`
	StorageLimitBytes *float64 `json:"storageLimitBytes,omitempty"`
	StorageLimitPurgeStrategy *string `json:"storageLimitPurgeStrategy,omitempty"`
	VideoQualities []string `json:"videoQualities,omitempty"`
} `json:"settings,omitempty"`
} `json:"vaultPlus,omitempty"`
}) (*UpdateAccountConnectionsSettingsResponse, error) {
	path := fmt.Sprintf("/v2/account/connections/%v/settings", connectionId)
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "PATCH", path, opts)
	if err != nil {
		return nil, err
	}

	var result UpdateAccountConnectionsSettingsResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetAccountSettings Get organization settings
func (c *Client) GetAccountSettings(ctx context.Context) (*GetAccountSettingsResponse, error) {
	path := "/v2/account/settings"

	respBody, err := c.Request(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}

	var result GetAccountSettingsResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// UpdateAccountSettings Update organization settings
func (c *Client) UpdateAccountSettings(ctx context.Context, body struct {
	VaultPlus *struct {
	ApplyToExistingConnections *bool `json:"applyToExistingConnections,omitempty"`
	AutoEnableForNewConnections *bool `json:"autoEnableForNewConnections,omitempty"`
	DefaultSettings *struct {
	AccessExpiryDays *float64 `json:"accessExpiryDays,omitempty"`
	AutoCacheMessages *bool `json:"autoCacheMessages,omitempty"`
	AutoCachePosts *bool `json:"autoCachePosts,omitempty"`
	AutoCacheStories *bool `json:"autoCacheStories,omitempty"`
	AutoCacheVault *bool `json:"autoCacheVault,omitempty"`
	CacheAudio *bool `json:"cacheAudio,omitempty"`
	CacheImages *bool `json:"cacheImages,omitempty"`
	CacheVideos *bool `json:"cacheVideos,omitempty"`
	ImageQualities []string `json:"imageQualities,omitempty"`
	MinAccessCountMessages *float64 `json:"minAccessCountMessages,omitempty"`
	MinAccessCountPosts *float64 `json:"minAccessCountPosts,omitempty"`
	MinAccessCountStories *float64 `json:"minAccessCountStories,omitempty"`
	MinAccessCountVault *float64 `json:"minAccessCountVault,omitempty"`
	PresignedUrlTtlSeconds *float64 `json:"presignedUrlTtlSeconds,omitempty"`
	RetentionDays *float64 `json:"retentionDays,omitempty"`
	StorageLimitBytes *float64 `json:"storageLimitBytes,omitempty"`
	StorageLimitPurgeStrategy *string `json:"storageLimitPurgeStrategy,omitempty"`
	VideoQualities []string `json:"videoQualities,omitempty"`
} `json:"defaultSettings,omitempty"`
} `json:"vaultPlus,omitempty"`
}) (*UpdateAccountSettingsResponse, error) {
	path := "/v2/account/settings"
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "PATCH", path, opts)
	if err != nil {
		return nil, err
	}

	var result UpdateAccountSettingsResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// Whoami Whoami
func (c *Client) Whoami(ctx context.Context) (*WhoamiResponse, error) {
	path := "/v2/account/whoami"

	respBody, err := c.Request(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}

	var result WhoamiResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// ListDynamicRules Get dynamic rules
func (c *Client) ListDynamicRules(ctx context.Context) (*ListDynamicRulesResponse, error) {
	path := "/v2/dynamic-rules"

	respBody, err := c.Request(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}

	var result ListDynamicRulesResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// DynamicRulesSign Generate sign headers for a request
func (c *Client) DynamicRulesSign(ctx context.Context, body struct {
	Endpoint string `json:"endpoint"`
	Time *interface{} `json:"time,omitempty"`
	UserId *interface{} `json:"user-id,omitempty"`
}) (*DynamicRulesSignResponse, error) {
	path := "/v2/dynamic-rules/sign"
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}

	var result DynamicRulesSignResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetDynamicRulesStatus Get dynamic rules status
func (c *Client) GetDynamicRulesStatus(ctx context.Context) (*GetDynamicRulesStatusResponse, error) {
	path := "/v2/dynamic-rules/status"

	respBody, err := c.Request(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}

	var result GetDynamicRulesStatusResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// LinkInit Initialize a Link session
func (c *Client) LinkInit(ctx context.Context, body struct {
	ClientReferenceId *string `json:"clientReferenceId,omitempty"`
	ConnectionId *string `json:"connectionId,omitempty"`
	Geolocation *struct {
	City *string `json:"city,omitempty"`
	Country string `json:"country"`
	State *string `json:"state,omitempty"`
} `json:"geolocation,omitempty"`
	RedirectUrl *string `json:"redirectUrl,omitempty"`
}) (*LinkInitResponse, error) {
	path := "/v2/link/init"
	opts := &RequestOptions{Body: body}

	respBody, err := c.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}

	var result LinkInitResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetLink Get login status
func (c *Client) GetLink(ctx context.Context, clientSecret string) (*GetLinkResponse, error) {
	path := fmt.Sprintf("/v2/link/%v", clientSecret)

	respBody, err := c.Request(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}

	var result GetLinkResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// DeleteLink Delete login session
func (c *Client) DeleteLink(ctx context.Context, clientSecret string) (*map[string]interface{}, error) {
	path := fmt.Sprintf("/v2/link/%v", clientSecret)

	respBody, err := c.Request(ctx, "DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetVaultPlus Get media item with all quality variants
func (c *Client) GetVaultPlus(ctx context.Context, connectionID string, mediaId string) (*GetVaultPlusResponse, error) {
	path := fmt.Sprintf("/v2/vault-plus/%v", mediaId)
	opts := &RequestOptions{ConnectionID: connectionID}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result GetVaultPlusResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// DeleteVaultPlus Delete a stored media item
func (c *Client) DeleteVaultPlus(ctx context.Context, connectionID string, mediaId string) (*DeleteVaultPlusResponse, error) {
	path := fmt.Sprintf("/v2/vault-plus/%v", mediaId)
	opts := &RequestOptions{ConnectionID: connectionID}

	respBody, err := c.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}

	var result DeleteVaultPlusResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// CreateVaultPlusBatch Get multiple media items with all quality variants
func (c *Client) CreateVaultPlusBatch(ctx context.Context, connectionID string, body struct {
	MediaIds []string `json:"mediaIds"`
}) (*CreateVaultPlusBatchResponse, error) {
	path := "/v2/vault-plus/batch"
	opts := &RequestOptions{ConnectionID: connectionID}
	opts.Body = body

	respBody, err := c.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}

	var result CreateVaultPlusBatchResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetVaultPlusList List stored media for a connection
func (c *Client) GetVaultPlusList(ctx context.Context, connectionID string, opts *RequestOptions) (*GetVaultPlusListResponse, error) {
	path := "/v2/vault-plus/list"
	if opts == nil {
		opts = &RequestOptions{}
	}
	opts.ConnectionID = connectionID

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result GetVaultPlusListResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// VaultPlusPurge Remove all stored media for a connection
func (c *Client) VaultPlusPurge(ctx context.Context, connectionID string) (*VaultPlusPurgeResponse, error) {
	path := "/v2/vault-plus/purge"
	opts := &RequestOptions{ConnectionID: connectionID}

	respBody, err := c.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}

	var result VaultPlusPurgeResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// CreateVaultPlusStoreList Store all media from a vault list
func (c *Client) CreateVaultPlusStoreList(ctx context.Context, connectionID string, listId string) (*CreateVaultPlusStoreListResponse, error) {
	path := fmt.Sprintf("/v2/vault-plus/store/list/%v", listId)
	opts := &RequestOptions{ConnectionID: connectionID}

	respBody, err := c.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}

	var result CreateVaultPlusStoreListResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetVaultPlusStoreStats Get organization vault stats
func (c *Client) GetVaultPlusStoreStats(ctx context.Context) (*GetVaultPlusStoreStatsResponse, error) {
	path := "/v2/vault-plus/store/stats"

	respBody, err := c.Request(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}

	var result GetVaultPlusStoreStatsResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

// GetVaultPlusStoreStatus Get storage status for a connection
func (c *Client) GetVaultPlusStoreStatus(ctx context.Context, connectionID string) (*GetVaultPlusStoreStatusResponse, error) {
	path := "/v2/vault-plus/store/status"
	opts := &RequestOptions{ConnectionID: connectionID}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result GetVaultPlusStoreStatusResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}

