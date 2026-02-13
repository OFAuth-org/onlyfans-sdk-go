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

// WhoamiResponse represents the response for Whoami
type WhoamiResponse struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Permissions []string `json:"permissions"`
}

// ListAccountConnectionsResponse represents the response for ListAccountConnections
type ListAccountConnectionsResponse struct {
	List []struct {
	Status string `json:"status"`
	Id string `json:"id"`
	UserData struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	Avatar string `json:"avatar"`
} `json:"userData"`
	Permissions []string `json:"permissions"`
	ExpiredAt string `json:"expiredAt"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	ClientReferenceId string `json:"clientReferenceId"`
	Imported bool `json:"imported"`
	LastCheckedAt string `json:"lastCheckedAt"`
} `json:"list"`
	HasMore bool `json:"hasMore"`
}

// GetAccountConnectionsSettingsResponse represents the response for GetAccountConnectionsSettings
type GetAccountConnectionsSettingsResponse struct {
	ConnectionId string `json:"connectionId"`
	VaultPlus struct {
	Enabled bool `json:"enabled"`
	SettingsOverrides struct {
	AutoCacheVault *bool `json:"autoCacheVault,omitempty"`
	AutoCacheMessages *bool `json:"autoCacheMessages,omitempty"`
	AutoCachePosts *bool `json:"autoCachePosts,omitempty"`
	AutoCacheStories *bool `json:"autoCacheStories,omitempty"`
	MinAccessCountVault *float64 `json:"minAccessCountVault,omitempty"`
	MinAccessCountMessages *float64 `json:"minAccessCountMessages,omitempty"`
	MinAccessCountPosts *float64 `json:"minAccessCountPosts,omitempty"`
	MinAccessCountStories *float64 `json:"minAccessCountStories,omitempty"`
	CacheImages *bool `json:"cacheImages,omitempty"`
	CacheVideos *bool `json:"cacheVideos,omitempty"`
	CacheAudio *bool `json:"cacheAudio,omitempty"`
	ImageQualities []string `json:"imageQualities,omitempty"`
	VideoQualities []string `json:"videoQualities,omitempty"`
	RetentionDays *float64 `json:"retentionDays,omitempty"`
	AccessExpiryDays *float64 `json:"accessExpiryDays,omitempty"`
	PresignedUrlTtlSeconds *float64 `json:"presignedUrlTtlSeconds,omitempty"`
	StorageLimitBytes *float64 `json:"storageLimitBytes,omitempty"`
	StorageLimitPurgeStrategy *string `json:"storageLimitPurgeStrategy,omitempty"`
} `json:"settingsOverrides"`
	Stats struct {
	VaultPlusEnabled bool `json:"vaultPlusEnabled"`
	TotalStorageBytes float64 `json:"totalStorageBytes"`
	StorageLimitBytes float64 `json:"storageLimitBytes"`
	MediaCount float64 `json:"mediaCount"`
	StoredCount float64 `json:"storedCount"`
	StorageUsagePercent float64 `json:"storageUsagePercent"`
} `json:"stats"`
} `json:"vaultPlus"`
}

// UpdateAccountConnectionsSettingsResponse represents the response for UpdateAccountConnectionsSettings
type UpdateAccountConnectionsSettingsResponse struct {
	Settings struct {
	ConnectionId string `json:"connectionId"`
	VaultPlus struct {
	Enabled bool `json:"enabled"`
	SettingsOverrides struct {
	AutoCacheVault *bool `json:"autoCacheVault,omitempty"`
	AutoCacheMessages *bool `json:"autoCacheMessages,omitempty"`
	AutoCachePosts *bool `json:"autoCachePosts,omitempty"`
	AutoCacheStories *bool `json:"autoCacheStories,omitempty"`
	MinAccessCountVault *float64 `json:"minAccessCountVault,omitempty"`
	MinAccessCountMessages *float64 `json:"minAccessCountMessages,omitempty"`
	MinAccessCountPosts *float64 `json:"minAccessCountPosts,omitempty"`
	MinAccessCountStories *float64 `json:"minAccessCountStories,omitempty"`
	CacheImages *bool `json:"cacheImages,omitempty"`
	CacheVideos *bool `json:"cacheVideos,omitempty"`
	CacheAudio *bool `json:"cacheAudio,omitempty"`
	ImageQualities []string `json:"imageQualities,omitempty"`
	VideoQualities []string `json:"videoQualities,omitempty"`
	RetentionDays *float64 `json:"retentionDays,omitempty"`
	AccessExpiryDays *float64 `json:"accessExpiryDays,omitempty"`
	PresignedUrlTtlSeconds *float64 `json:"presignedUrlTtlSeconds,omitempty"`
	StorageLimitBytes *float64 `json:"storageLimitBytes,omitempty"`
	StorageLimitPurgeStrategy *string `json:"storageLimitPurgeStrategy,omitempty"`
} `json:"settingsOverrides"`
	Stats struct {
	VaultPlusEnabled bool `json:"vaultPlusEnabled"`
	TotalStorageBytes float64 `json:"totalStorageBytes"`
	StorageLimitBytes float64 `json:"storageLimitBytes"`
	MediaCount float64 `json:"mediaCount"`
	StoredCount float64 `json:"storedCount"`
	StorageUsagePercent float64 `json:"storageUsagePercent"`
} `json:"stats"`
} `json:"vaultPlus"`
} `json:"settings"`
	PurgeResult *struct {
	PurgedCount float64 `json:"purgedCount"`
	FreedBytes float64 `json:"freedBytes"`
} `json:"purgeResult,omitempty"`
}

// CreateAccountConnectionsImportResponse represents the response for CreateAccountConnectionsImport
type CreateAccountConnectionsImportResponse struct {
	Status string `json:"status"`
	Id string `json:"id"`
	UserData struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	Avatar string `json:"avatar"`
} `json:"userData"`
	Permissions []string `json:"permissions"`
	ExpiredAt string `json:"expiredAt"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	ClientReferenceId string `json:"clientReferenceId"`
	Imported bool `json:"imported"`
	LastCheckedAt string `json:"lastCheckedAt"`
}

// UpdateAccountConnectionsImportResponse represents the response for UpdateAccountConnectionsImport
type UpdateAccountConnectionsImportResponse struct {
	Status string `json:"status"`
	Id string `json:"id"`
	UserData struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	Avatar string `json:"avatar"`
} `json:"userData"`
	Permissions []string `json:"permissions"`
	ExpiredAt string `json:"expiredAt"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	ClientReferenceId string `json:"clientReferenceId"`
	Imported bool `json:"imported"`
	LastCheckedAt string `json:"lastCheckedAt"`
}

// GetAccountSettingsResponse represents the response for GetAccountSettings
type GetAccountSettingsResponse struct {
	VaultPlus struct {
	AutoEnableForNewConnections bool `json:"autoEnableForNewConnections"`
	DefaultSettings struct {
	AutoCacheVault bool `json:"autoCacheVault"`
	AutoCacheMessages bool `json:"autoCacheMessages"`
	AutoCachePosts bool `json:"autoCachePosts"`
	AutoCacheStories bool `json:"autoCacheStories"`
	MinAccessCountVault float64 `json:"minAccessCountVault"`
	MinAccessCountMessages float64 `json:"minAccessCountMessages"`
	MinAccessCountPosts float64 `json:"minAccessCountPosts"`
	MinAccessCountStories float64 `json:"minAccessCountStories"`
	CacheImages bool `json:"cacheImages"`
	CacheVideos bool `json:"cacheVideos"`
	CacheAudio bool `json:"cacheAudio"`
	ImageQualities []string `json:"imageQualities"`
	VideoQualities []string `json:"videoQualities"`
	RetentionDays float64 `json:"retentionDays"`
	AccessExpiryDays float64 `json:"accessExpiryDays"`
	PresignedUrlTtlSeconds float64 `json:"presignedUrlTtlSeconds"`
	StorageLimitBytes float64 `json:"storageLimitBytes"`
	StorageLimitPurgeStrategy string `json:"storageLimitPurgeStrategy"`
} `json:"defaultSettings"`
} `json:"vaultPlus"`
}

// UpdateAccountSettingsResponse represents the response for UpdateAccountSettings
type UpdateAccountSettingsResponse struct {
	Settings struct {
	VaultPlus struct {
	AutoEnableForNewConnections bool `json:"autoEnableForNewConnections"`
	DefaultSettings struct {
	AutoCacheVault bool `json:"autoCacheVault"`
	AutoCacheMessages bool `json:"autoCacheMessages"`
	AutoCachePosts bool `json:"autoCachePosts"`
	AutoCacheStories bool `json:"autoCacheStories"`
	MinAccessCountVault float64 `json:"minAccessCountVault"`
	MinAccessCountMessages float64 `json:"minAccessCountMessages"`
	MinAccessCountPosts float64 `json:"minAccessCountPosts"`
	MinAccessCountStories float64 `json:"minAccessCountStories"`
	CacheImages bool `json:"cacheImages"`
	CacheVideos bool `json:"cacheVideos"`
	CacheAudio bool `json:"cacheAudio"`
	ImageQualities []string `json:"imageQualities"`
	VideoQualities []string `json:"videoQualities"`
	RetentionDays float64 `json:"retentionDays"`
	AccessExpiryDays float64 `json:"accessExpiryDays"`
	PresignedUrlTtlSeconds float64 `json:"presignedUrlTtlSeconds"`
	StorageLimitBytes float64 `json:"storageLimitBytes"`
	StorageLimitPurgeStrategy string `json:"storageLimitPurgeStrategy"`
} `json:"defaultSettings"`
} `json:"vaultPlus"`
} `json:"settings"`
	BroadcastResult *struct {
	AffectedConnections float64 `json:"affectedConnections"`
	PurgeResults []struct {
	ConnectionId string `json:"connectionId"`
	PurgedCount float64 `json:"purgedCount"`
	FreedBytes float64 `json:"freedBytes"`
} `json:"purgeResults"`
} `json:"broadcastResult,omitempty"`
}

// UpdateAccessSelfResponse represents the response for UpdateAccessSelf
type UpdateAccessSelfResponse struct {
	Id float64 `json:"id"`
	Username string `json:"username"`
	Name string `json:"name"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs *struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs,omitempty"`
	Lists []struct {
	Id interface{} `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
} `json:"lists"`
	DisplayName *string `json:"displayName,omitempty"`
	Notice *string `json:"notice,omitempty"`
	About *string `json:"about,omitempty"`
	IsMarkdownDisabledForAbout *bool `json:"isMarkdownDisabledForAbout,omitempty"`
	Website *string `json:"website,omitempty"`
	Wishlist *string `json:"wishlist,omitempty"`
	Location *string `json:"location,omitempty"`
	Header *string `json:"header,omitempty"`
	HeaderSize *struct {
	Width float64 `json:"width"`
	Height float64 `json:"height"`
} `json:"headerSize,omitempty"`
	HeaderThumbs *struct {
	W480 string `json:"w480"`
	W760 string `json:"w760"`
} `json:"headerThumbs,omitempty"`
	SubscribersCount *float64 `json:"subscribersCount,omitempty"`
	PostsCount *float64 `json:"postsCount,omitempty"`
	ArchivedPostsCount *float64 `json:"archivedPostsCount,omitempty"`
	PrivateArchivedPostsCount *float64 `json:"privateArchivedPostsCount,omitempty"`
	PhotosCount *float64 `json:"photosCount,omitempty"`
	VideosCount *float64 `json:"videosCount,omitempty"`
	AudiosCount *float64 `json:"audiosCount,omitempty"`
	MediasCount *float64 `json:"mediasCount,omitempty"`
	FavoritesCount *float64 `json:"favoritesCount,omitempty"`
	FavoritedCount *float64 `json:"favoritedCount,omitempty"`
	JoinDate *string `json:"joinDate,omitempty"`
	LastSeen *string `json:"lastSeen,omitempty"`
	SubscribedBy *bool `json:"subscribedBy,omitempty"`
	SubscribedByExpire *bool `json:"subscribedByExpire,omitempty"`
	SubscribedByExpireDate *string `json:"subscribedByExpireDate,omitempty"`
	SubscribedByAutoprolong *bool `json:"subscribedByAutoprolong,omitempty"`
	SubscribedIsExpiredNow *bool `json:"subscribedIsExpiredNow,omitempty"`
	SubscribedByData *struct {
	Price float64 `json:"price"`
	NewPrice float64 `json:"newPrice"`
	RegularPrice float64 `json:"regularPrice"`
	SubscribePrice float64 `json:"subscribePrice"`
	DiscountPercent float64 `json:"discountPercent"`
	DiscountPeriod float64 `json:"discountPeriod"`
	SubscribeAt string `json:"subscribeAt"`
	ExpiredAt string `json:"expiredAt"`
	RenewedAt string `json:"renewedAt"`
	DiscountFinishedAt string `json:"discountFinishedAt"`
	DiscountStartedAt string `json:"discountStartedAt"`
	Status string `json:"status"`
	IsMuted bool `json:"isMuted"`
	UnsubscribeReason string `json:"unsubscribeReason"`
	Duration string `json:"duration"`
	HasActivePaidSubscriptions bool `json:"hasActivePaidSubscriptions"`
	ShowPostsInFeed bool `json:"showPostsInFeed"`
	Subscribes []struct {
	Id float64 `json:"id"`
	UserId float64 `json:"userId"`
	SubscriberId float64 `json:"subscriberId"`
	Date string `json:"date"`
	Duration float64 `json:"duration"`
	StartDate string `json:"startDate"`
	ExpireDate string `json:"expireDate"`
	CancelDate string `json:"cancelDate"`
	Price float64 `json:"price"`
	RegularPrice float64 `json:"regularPrice"`
	Discount float64 `json:"discount"`
	EarningId float64 `json:"earningId"`
	Action string `json:"action"`
	Type string `json:"type"`
	OfferStart string `json:"offerStart"`
	OfferEnd string `json:"offerEnd"`
	IsCurrent bool `json:"isCurrent"`
} `json:"subscribes"`
} `json:"subscribedByData,omitempty"`
	SubscribedOn *bool `json:"subscribedOn,omitempty"`
	SubscribedOnExpiredNow *bool `json:"subscribedOnExpiredNow,omitempty"`
	SubscribedOnDuration *string `json:"subscribedOnDuration,omitempty"`
	SubscribedOnData *struct {
	Price float64 `json:"price"`
	NewPrice float64 `json:"newPrice"`
	RegularPrice float64 `json:"regularPrice"`
	SubscribePrice float64 `json:"subscribePrice"`
	DiscountPercent float64 `json:"discountPercent"`
	DiscountPeriod float64 `json:"discountPeriod"`
	SubscribeAt string `json:"subscribeAt"`
	ExpiredAt string `json:"expiredAt"`
	RenewedAt string `json:"renewedAt"`
	DiscountFinishedAt string `json:"discountFinishedAt"`
	DiscountStartedAt string `json:"discountStartedAt"`
	Status string `json:"status"`
	IsMuted bool `json:"isMuted"`
	UnsubscribeReason string `json:"unsubscribeReason"`
	Duration string `json:"duration"`
	HasActivePaidSubscriptions bool `json:"hasActivePaidSubscriptions"`
	Subscribes []struct {
	Id float64 `json:"id"`
	UserId float64 `json:"userId"`
	SubscriberId float64 `json:"subscriberId"`
	Date string `json:"date"`
	Duration float64 `json:"duration"`
	StartDate string `json:"startDate"`
	ExpireDate string `json:"expireDate"`
	CancelDate string `json:"cancelDate"`
	Price float64 `json:"price"`
	RegularPrice float64 `json:"regularPrice"`
	Discount float64 `json:"discount"`
	EarningId float64 `json:"earningId"`
	Action string `json:"action"`
	Type string `json:"type"`
	OfferStart string `json:"offerStart"`
	OfferEnd string `json:"offerEnd"`
	IsCurrent bool `json:"isCurrent"`
} `json:"subscribes"`
	TipsSumm float64 `json:"tipsSumm"`
	SubscribesSumm float64 `json:"subscribesSumm"`
	MessagesSumm float64 `json:"messagesSumm"`
	PostsSumm float64 `json:"postsSumm"`
	StreamsSumm float64 `json:"streamsSumm"`
	TotalSumm float64 `json:"totalSumm"`
} `json:"subscribedOnData,omitempty"`
	SubscribePrice *float64 `json:"subscribePrice,omitempty"`
	CurrentSubscribePrice *float64 `json:"currentSubscribePrice,omitempty"`
	CanAddSubscriber *bool `json:"canAddSubscriber,omitempty"`
	TipsEnabled *bool `json:"tipsEnabled,omitempty"`
	TipsTextEnabled *bool `json:"tipsTextEnabled,omitempty"`
	TipsMin *float64 `json:"tipsMin,omitempty"`
	TipsMinInternal *float64 `json:"tipsMinInternal,omitempty"`
	TipsMax *float64 `json:"tipsMax,omitempty"`
	CanLookStory *bool `json:"canLookStory,omitempty"`
	CanCommentStory *bool `json:"canCommentStory,omitempty"`
	HasNotViewedStory *bool `json:"hasNotViewedStory,omitempty"`
	HasStories *bool `json:"hasStories,omitempty"`
	IsRestricted *bool `json:"isRestricted,omitempty"`
	CanRestrict *bool `json:"canRestrict,omitempty"`
	IsBlocked *bool `json:"isBlocked,omitempty"`
	CanReport *bool `json:"canReport,omitempty"`
	CanUnsubscribe *bool `json:"canUnsubscribe,omitempty"`
	IsPendingAutoprolong *bool `json:"isPendingAutoprolong,omitempty"`
	IsPerformer *bool `json:"isPerformer,omitempty"`
	IsRealPerformer *bool `json:"isRealPerformer,omitempty"`
	CanReceiveChatMessage *bool `json:"canReceiveChatMessage,omitempty"`
	CanChat *bool `json:"canChat,omitempty"`
	ShowPostsInFeed *bool `json:"showPostsInFeed,omitempty"`
	HasPinnedPosts *bool `json:"hasPinnedPosts,omitempty"`
	HasLabels *bool `json:"hasLabels,omitempty"`
	IsPrivateRestriction *bool `json:"isPrivateRestriction,omitempty"`
	ShowSubscribersCount *bool `json:"showSubscribersCount,omitempty"`
	ShowMediaCount *bool `json:"showMediaCount,omitempty"`
	IsReferrerAllowed *bool `json:"isReferrerAllowed,omitempty"`
	CanCreatePromotion *bool `json:"canCreatePromotion,omitempty"`
	CanCreateTrial *bool `json:"canCreateTrial,omitempty"`
	IsAdultContent *bool `json:"isAdultContent,omitempty"`
	CanTrialSend *bool `json:"canTrialSend,omitempty"`
	IsFriend *bool `json:"isFriend,omitempty"`
	HasScheduledStream *bool `json:"hasScheduledStream,omitempty"`
	HasStream *bool `json:"hasStream,omitempty"`
	CanPayInternal *bool `json:"canPayInternal,omitempty"`
}

// ListAccessSelfNotificationsResponse represents the response for ListAccessSelfNotifications
type ListAccessSelfNotificationsResponse struct {
	List []struct {
	Id float64 `json:"id"`
	Type string `json:"type"`
	SubType string `json:"subType"`
	CreatedAt string `json:"createdAt"`
	IsRead bool `json:"isRead"`
	Text string `json:"text"`
	ReplacePairs map[string]string `json:"replacePairs"`
	CanGoToProfile bool `json:"canGoToProfile"`
	UserId float64 `json:"userId"`
} `json:"list"`
	HasMore bool `json:"hasMore"`
}

// ListAccessSelfReleaseFormsResponse represents the response for ListAccessSelfReleaseForms
type ListAccessSelfReleaseFormsResponse struct {
	List []struct {
	Id float64 `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
	PartnerId float64 `json:"partnerId"`
	Code string `json:"code"`
	Status string `json:"status"`
	CreatedAt string `json:"createdAt"`
	ApprovedAt string `json:"approvedAt"`
	LastChangedAt string `json:"lastChangedAt"`
	UserName string `json:"userName"`
	HasUser bool `json:"hasUser"`
	IsHidden bool `json:"isHidden"`
} `json:"list"`
	HasMore bool `json:"hasMore"`
}

// ListAccessSelfTaggedFriendUsersResponse represents the response for ListAccessSelfTaggedFriendUsers
type ListAccessSelfTaggedFriendUsersResponse struct {
	List []struct {
	Id float64 `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
	User struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	IsHidden bool `json:"isHidden"`
} `json:"user"`
} `json:"list"`
	HasMore bool `json:"hasMore"`
}

// GetAccessEarningsChartResponse represents the response for GetAccessEarningsChart
type GetAccessEarningsChartResponse struct {
	Chart []struct {
	Date string `json:"date"`
	Count float64 `json:"count"`
} `json:"chart"`
	Total float64 `json:"total"`
	Gross *float64 `json:"gross,omitempty"`
	Delta *float64 `json:"delta,omitempty"`
}

// ListAccessEarningsTransactionsResponse represents the response for ListAccessEarningsTransactions
type ListAccessEarningsTransactionsResponse struct {
	List []struct {
	Id string `json:"id"`
	Type string `json:"type"`
	CreatedAt string `json:"createdAt"`
	Amounts struct {
	Gross float64 `json:"gross"`
	Net float64 `json:"net"`
	Fee float64 `json:"fee"`
	Vat float64 `json:"vat"`
	Tax float64 `json:"tax"`
} `json:"amounts"`
	Currency string `json:"currency"`
	Description string `json:"description"`
	Status string `json:"status"`
	PayoutPendingDays float64 `json:"payoutPendingDays"`
	User struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	Avatar string `json:"avatar"`
} `json:"user"`
} `json:"list"`
	HasMore bool `json:"hasMore"`
	NextMarker float64 `json:"nextMarker"`
}

// ListAccessEarningsChargebacksResponse represents the response for ListAccessEarningsChargebacks
type ListAccessEarningsChargebacksResponse struct {
	List []struct {
	Id float64 `json:"id"`
	CreatedAt string `json:"createdAt"`
	PaymentType string `json:"paymentType"`
	Payment struct {
	Id string `json:"id"`
	CreatedAt string `json:"createdAt"`
	Amounts struct {
	Gross float64 `json:"gross"`
	Net float64 `json:"net"`
	Fee float64 `json:"fee"`
	Vat float64 `json:"vat"`
	Tax float64 `json:"tax"`
} `json:"amounts"`
	Currency string `json:"currency"`
	Description string `json:"description"`
	Status string `json:"status"`
	User struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	Avatar string `json:"avatar"`
} `json:"user"`
} `json:"payment"`
} `json:"list"`
	HasMore bool `json:"hasMore"`
	NextMarker float64 `json:"nextMarker"`
}

// GetAccessAnalyticsPostsChartResponse represents the response for GetAccessAnalyticsPostsChart
type GetAccessAnalyticsPostsChartResponse map[string]struct {
	Chart []struct {
	Date string `json:"date"`
	Count float64 `json:"count"`
} `json:"chart"`
	Total float64 `json:"total"`
	Delta float64 `json:"delta"`
	Gross *float64 `json:"gross,omitempty"`
}

// GetAccessAnalyticsPostsTopResponse represents the response for GetAccessAnalyticsPostsTop
type GetAccessAnalyticsPostsTopResponse struct {
	HasMore bool `json:"hasMore"`
	Items []struct {
	Id float64 `json:"id"`
	CanDelete bool `json:"canDelete"`
	CanEdit bool `json:"canEdit"`
	MediaCount float64 `json:"mediaCount"`
	Media []struct {
	Id float64 `json:"id"`
	Type string `json:"type"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CanView bool `json:"canView"`
	HasError bool `json:"hasError"`
	CreatedAt string `json:"createdAt"`
	IsReady bool `json:"isReady"`
	Duration *float64 `json:"duration,omitempty"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	View string `json:"view"`
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs"`
	IvStatus string `json:"ivStatus"`
	IsFromGuest bool `json:"isFromGuest"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
	Files *struct {
	Full struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"sources"`
} `json:"full"`
	Thumb *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"thumb,omitempty"`
	Preview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Options []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"options"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"squarePreview,omitempty"`
} `json:"files,omitempty"`
} `json:"media"`
	CanViewMedia bool `json:"canViewMedia"`
	Author *struct {
	Id float64 `json:"id"`
	View string `json:"_view"`
} `json:"author,omitempty"`
	ResponseType string `json:"responseType"`
	PostedAt string `json:"postedAt"`
	PostedAtPrecise string `json:"postedAtPrecise"`
	IsMarkdownDisabled bool `json:"isMarkdownDisabled"`
	IsOpened bool `json:"isOpened"`
	CanToggleFavorite bool `json:"canToggleFavorite"`
	TipsAmount string `json:"tipsAmount"`
	Text string `json:"text"`
	IsFavorite bool `json:"isFavorite"`
	CanComment bool `json:"canComment"`
	FavoritesCount float64 `json:"favoritesCount"`
	IsMediaReady bool `json:"isMediaReady"`
	RawText string `json:"rawText"`
	Stats *struct {
	IsAvailable bool `json:"isAvailable"`
	HasStats bool `json:"hasStats"`
	HasVideo *bool `json:"hasVideo,omitempty"`
	LookCount float64 `json:"lookCount"`
	UniqueLookCount float64 `json:"uniqueLookCount"`
	UniqueLookChart []struct {
	Date string `json:"date"`
	Count float64 `json:"count"`
} `json:"uniqueLookChart"`
	LookChart []struct {
	Date string `json:"date"`
	Count float64 `json:"count"`
} `json:"lookChart"`
	LookDuration float64 `json:"lookDuration"`
	LookDurationAverage float64 `json:"lookDurationAverage"`
	LikeCount float64 `json:"likeCount"`
	LikeChart []struct {
	Date string `json:"date"`
	Count float64 `json:"count"`
} `json:"likeChart"`
	CommentCount float64 `json:"commentCount"`
	CommentChart []struct {
	Date string `json:"date"`
	Count float64 `json:"count"`
} `json:"commentChart"`
	TipCount float64 `json:"tipCount"`
	TipChart []struct {
	Date string `json:"date"`
	Count float64 `json:"count"`
} `json:"tipChart"`
	TipSum float64 `json:"tipSum"`
	PurchasedCount float64 `json:"purchasedCount"`
	PurchasedSumm float64 `json:"purchasedSumm"`
} `json:"stats,omitempty"`
	HasVoting *bool `json:"hasVoting,omitempty"`
	VotingType *float64 `json:"votingType,omitempty"`
	Voting *struct {
	FinishedAt string `json:"finishedAt"`
	Options []struct {
	Id float64 `json:"id"`
	Text *string `json:"text,omitempty"`
	Count *float64 `json:"count,omitempty"`
	Percent *float64 `json:"percent,omitempty"`
	IsWinner *bool `json:"isWinner,omitempty"`
} `json:"options"`
	Total float64 `json:"total"`
} `json:"voting,omitempty"`
} `json:"items"`
}

// GetAccessAnalyticsPostsResponse represents the response for GetAccessAnalyticsPosts
type GetAccessAnalyticsPostsResponse struct {
	IsAvailable bool `json:"isAvailable"`
	HasStats bool `json:"hasStats"`
	HasVideo *bool `json:"hasVideo,omitempty"`
	LookCount float64 `json:"lookCount"`
	UniqueLookCount float64 `json:"uniqueLookCount"`
	UniqueLookChart []struct {
	Date string `json:"date"`
	Count float64 `json:"count"`
} `json:"uniqueLookChart"`
	LookChart []struct {
	Date string `json:"date"`
	Count float64 `json:"count"`
} `json:"lookChart"`
	LookDuration float64 `json:"lookDuration"`
	LookDurationAverage float64 `json:"lookDurationAverage"`
	LikeCount float64 `json:"likeCount"`
	LikeChart []struct {
	Date string `json:"date"`
	Count float64 `json:"count"`
} `json:"likeChart"`
	CommentCount float64 `json:"commentCount"`
	CommentChart []struct {
	Date string `json:"date"`
	Count float64 `json:"count"`
} `json:"commentChart"`
	TipCount float64 `json:"tipCount"`
	TipChart []struct {
	Date string `json:"date"`
	Count float64 `json:"count"`
} `json:"tipChart"`
	TipSum float64 `json:"tipSum"`
	PurchasedCount float64 `json:"purchasedCount"`
	PurchasedSumm float64 `json:"purchasedSumm"`
	TipSumChart []struct {
	Date string `json:"date"`
	Count float64 `json:"count"`
} `json:"tipSumChart"`
	PurchasesChart []struct {
	Date string `json:"date"`
	Count float64 `json:"count"`
} `json:"purchasesChart"`
}

// GetAccessAnalyticsStreamsChartResponse represents the response for GetAccessAnalyticsStreamsChart
type GetAccessAnalyticsStreamsChartResponse map[string]struct {
	Chart []struct {
	Date string `json:"date"`
	Count float64 `json:"count"`
} `json:"chart"`
	Total float64 `json:"total"`
	Delta float64 `json:"delta"`
	Gross *float64 `json:"gross,omitempty"`
}

// GetAccessAnalyticsStreamsTopResponse represents the response for GetAccessAnalyticsStreamsTop
type GetAccessAnalyticsStreamsTopResponse struct {
	HasMore bool `json:"hasMore"`
	Items []struct {
	Id float64 `json:"id"`
	CanDelete bool `json:"canDelete"`
	CanEdit bool `json:"canEdit"`
	MediaCount float64 `json:"mediaCount"`
	Media []struct {
	Id float64 `json:"id"`
	Type string `json:"type"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CanView bool `json:"canView"`
	HasError bool `json:"hasError"`
	CreatedAt string `json:"createdAt"`
	IsReady bool `json:"isReady"`
	Duration *float64 `json:"duration,omitempty"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	View string `json:"view"`
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs"`
	IvStatus string `json:"ivStatus"`
	IsFromGuest bool `json:"isFromGuest"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
	Files *struct {
	Full struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"sources"`
} `json:"full"`
	Thumb *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"thumb,omitempty"`
	Preview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Options []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"options"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"squarePreview,omitempty"`
} `json:"files,omitempty"`
} `json:"media"`
	CanViewMedia bool `json:"canViewMedia"`
	Author *struct {
	Id float64 `json:"id"`
	View string `json:"_view"`
} `json:"author,omitempty"`
	ResponseType string `json:"responseType"`
	PostedAt string `json:"postedAt"`
	PostedAtPrecise string `json:"postedAtPrecise"`
	IsMarkdownDisabled bool `json:"isMarkdownDisabled"`
	IsOpened bool `json:"isOpened"`
	CanToggleFavorite bool `json:"canToggleFavorite"`
	TipsAmount string `json:"tipsAmount"`
	Text string `json:"text"`
	IsFavorite bool `json:"isFavorite"`
	CanComment bool `json:"canComment"`
	FavoritesCount float64 `json:"favoritesCount"`
	IsMediaReady bool `json:"isMediaReady"`
	RawText string `json:"rawText"`
	Stats *struct {
	IsAvailable bool `json:"isAvailable"`
	HasStats bool `json:"hasStats"`
	HasVideo *bool `json:"hasVideo,omitempty"`
	LookCount float64 `json:"lookCount"`
	UniqueLookCount float64 `json:"uniqueLookCount"`
	UniqueLookChart []struct {
	Date string `json:"date"`
	Count float64 `json:"count"`
} `json:"uniqueLookChart"`
	LookChart []struct {
	Date string `json:"date"`
	Count float64 `json:"count"`
} `json:"lookChart"`
	LookDuration float64 `json:"lookDuration"`
	LookDurationAverage float64 `json:"lookDurationAverage"`
	LikeCount float64 `json:"likeCount"`
	LikeChart []struct {
	Date string `json:"date"`
	Count float64 `json:"count"`
} `json:"likeChart"`
	CommentCount float64 `json:"commentCount"`
	CommentChart []struct {
	Date string `json:"date"`
	Count float64 `json:"count"`
} `json:"commentChart"`
	TipCount float64 `json:"tipCount"`
	TipChart []struct {
	Date string `json:"date"`
	Count float64 `json:"count"`
} `json:"tipChart"`
	TipSum float64 `json:"tipSum"`
	PurchasedCount float64 `json:"purchasedCount"`
	PurchasedSumm float64 `json:"purchasedSumm"`
} `json:"stats,omitempty"`
} `json:"items"`
}

// GetAccessAnalyticsStoriesChartResponse represents the response for GetAccessAnalyticsStoriesChart
type GetAccessAnalyticsStoriesChartResponse map[string]struct {
	Chart []struct {
	Date string `json:"date"`
	Count float64 `json:"count"`
} `json:"chart"`
	Total float64 `json:"total"`
	Delta float64 `json:"delta"`
	Gross *float64 `json:"gross,omitempty"`
}

// GetAccessAnalyticsStoriesTopResponse represents the response for GetAccessAnalyticsStoriesTop
type GetAccessAnalyticsStoriesTopResponse struct {
	HasMore bool `json:"hasMore"`
	Items []struct {
	Id float64 `json:"id"`
	UserId float64 `json:"userId"`
	IsReady bool `json:"isReady"`
	HasPost bool `json:"hasPost"`
	IsWatched bool `json:"isWatched"`
	Media []struct {
	Id float64 `json:"id"`
	Type string `json:"type"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CanView bool `json:"canView"`
	HasError bool `json:"hasError"`
	CreatedAt string `json:"createdAt"`
	IsReady bool `json:"isReady"`
	Duration *float64 `json:"duration,omitempty"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	View string `json:"view"`
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs"`
	IvStatus string `json:"ivStatus"`
	IsFromGuest bool `json:"isFromGuest"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
	Files *struct {
	Full struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"sources"`
} `json:"full"`
	Thumb *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"thumb,omitempty"`
	Preview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Options []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"options"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"squarePreview,omitempty"`
} `json:"files,omitempty"`
} `json:"media"`
	CreatedAt string `json:"createdAt"`
	CanvasHeight float64 `json:"canvasHeight"`
	CanvasWidth float64 `json:"canvasWidth"`
	Question *struct {
	Entity struct {
	Id float64 `json:"id"`
	Text string `json:"text"`
	CreatedAt string `json:"createdAt"`
} `json:"entity"`
	Type string `json:"type"`
	Positions struct {
	ZIndex float64 `json:"zIndex"`
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Top float64 `json:"top"`
	Left float64 `json:"left"`
	Width float64 `json:"width"`
	Height float64 `json:"height"`
	Angle float64 `json:"angle"`
	Color string `json:"color"`
} `json:"positions"`
} `json:"question,omitempty"`
	ViewersCount float64 `json:"viewersCount"`
	Viewers []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs"`
	IsVerified bool `json:"isVerified"`
	CanLookStory bool `json:"canLookStory"`
	CanCommentStory bool `json:"canCommentStory"`
	HasNotViewedStory bool `json:"hasNotViewedStory"`
	IsStoryLiked bool `json:"isStoryLiked"`
	IsStoryBlockedUser bool `json:"isStoryBlockedUser"`
} `json:"viewers"`
	CommentsCount float64 `json:"commentsCount"`
	CanDelete bool `json:"canDelete"`
	IsHighlightCover bool `json:"isHighlightCover"`
	IsLastInHighlight bool `json:"isLastInHighlight"`
	TipsAmount string `json:"tipsAmount"`
	TipsAmountRaw float64 `json:"tipsAmountRaw"`
	TipsCount float64 `json:"tipsCount"`
	LikesCount float64 `json:"likesCount"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	View string `json:"view"`
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs"`
	IvStatus string `json:"ivStatus"`
	IsFromGuest bool `json:"isFromGuest"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	Post *struct {
	Id float64 `json:"id"`
	Author *struct {
	Id float64 `json:"id"`
	View string `json:"_view"`
} `json:"author,omitempty"`
	CoordinateParams *struct {
	Top float64 `json:"top"`
	Left float64 `json:"left"`
	Angle float64 `json:"angle"`
	Scale float64 `json:"scale"`
	Width float64 `json:"width"`
	Height float64 `json:"height"`
} `json:"coordinateParams,omitempty"`
} `json:"post,omitempty"`
	Texts []struct {
	Text string `json:"text"`
	Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Width float64 `json:"width"`
	Height float64 `json:"height"`
	Angle float64 `json:"angle"`
} `json:"position"`
} `json:"texts"`
} `json:"items"`
}

// GetAccessAnalyticsMassMessagesChartResponse represents the response for GetAccessAnalyticsMassMessagesChart
type GetAccessAnalyticsMassMessagesChartResponse struct {
	GroupMessages struct {
	Chart []struct {
	Date string `json:"date"`
	Count float64 `json:"count"`
} `json:"chart"`
	Total float64 `json:"total"`
	Delta float64 `json:"delta"`
	Gross *float64 `json:"gross,omitempty"`
} `json:"group_messages"`
	GroupMessagesPurchases struct {
	Chart []struct {
	Date string `json:"date"`
	Count float64 `json:"count"`
} `json:"chart"`
	Total float64 `json:"total"`
	Delta float64 `json:"delta"`
	Gross *float64 `json:"gross,omitempty"`
} `json:"group_messages_purchases"`
	DirectMessages struct {
	Chart []struct {
	Date string `json:"date"`
	Count float64 `json:"count"`
} `json:"chart"`
	Total float64 `json:"total"`
	Delta float64 `json:"delta"`
	Gross *float64 `json:"gross,omitempty"`
} `json:"direct_messages"`
	DirectMessagesPurchases struct {
	Chart []struct {
	Date string `json:"date"`
	Count float64 `json:"count"`
} `json:"chart"`
	Total float64 `json:"total"`
	Delta float64 `json:"delta"`
	Gross *float64 `json:"gross,omitempty"`
} `json:"direct_messages_purchases"`
}

// GetAccessAnalyticsMassMessagesSentResponse represents the response for GetAccessAnalyticsMassMessagesSent
type GetAccessAnalyticsMassMessagesSentResponse struct {
	HasMore bool `json:"hasMore"`
	Items []struct {
	Id float64 `json:"id"`
	Date string `json:"date"`
	ResponseType *string `json:"responseType,omitempty"`
	Text string `json:"text"`
	RawText *string `json:"rawText,omitempty"`
	GiphyId *string `json:"giphyId,omitempty"`
	IsFree *bool `json:"isFree,omitempty"`
	IsMediaReady *bool `json:"isMediaReady,omitempty"`
	MediaCount *float64 `json:"mediaCount,omitempty"`
	Media []struct {
	Id float64 `json:"id"`
	Type string `json:"type"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CanView bool `json:"canView"`
	HasError bool `json:"hasError"`
	CreatedAt string `json:"createdAt"`
	IsReady bool `json:"isReady"`
	Duration *float64 `json:"duration,omitempty"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	View string `json:"view"`
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs"`
	IvStatus string `json:"ivStatus"`
	IsFromGuest bool `json:"isFromGuest"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
	Files *struct {
	Full struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"sources"`
} `json:"full"`
	Thumb *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"thumb,omitempty"`
	Preview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Options []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"options"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"squarePreview,omitempty"`
} `json:"files,omitempty"`
} `json:"media"`
	Previews []struct {
	Id float64 `json:"id"`
	Type string `json:"type"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CanView bool `json:"canView"`
	HasError bool `json:"hasError"`
	CreatedAt string `json:"createdAt"`
	IsReady bool `json:"isReady"`
	Duration *float64 `json:"duration,omitempty"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	View string `json:"view"`
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs"`
	IvStatus string `json:"ivStatus"`
	IsFromGuest bool `json:"isFromGuest"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
	Files *struct {
	Full struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"sources"`
} `json:"full"`
	Thumb *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"thumb,omitempty"`
	Preview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Options []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"options"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"squarePreview,omitempty"`
} `json:"files,omitempty"`
} `json:"previews"`
	IsTip *bool `json:"isTip,omitempty"`
	IsReportedByMe *bool `json:"isReportedByMe,omitempty"`
	ViewedCount *float64 `json:"viewedCount,omitempty"`
	SentCount *float64 `json:"sentCount,omitempty"`
	IsCanceled *bool `json:"isCanceled,omitempty"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	View string `json:"view"`
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs"`
	IvStatus string `json:"ivStatus"`
	IsFromGuest bool `json:"isFromGuest"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	Template *string `json:"template,omitempty"`
	CanUnsend *bool `json:"canUnsend,omitempty"`
	UnsendSeconds *float64 `json:"unsendSeconds,omitempty"`
	Price *string `json:"price,omitempty"`
	PurchasedCount *float64 `json:"purchasedCount,omitempty"`
	CanSendMessageToBuyers *bool `json:"canSendMessageToBuyers,omitempty"`
} `json:"items"`
}

// GetAccessAnalyticsMassMessagesPurchasedResponse represents the response for GetAccessAnalyticsMassMessagesPurchased
type GetAccessAnalyticsMassMessagesPurchasedResponse struct {
	HasMore bool `json:"hasMore"`
	Items []struct {
	Id float64 `json:"id"`
	Date string `json:"date"`
	ResponseType *string `json:"responseType,omitempty"`
	Text string `json:"text"`
	RawText *string `json:"rawText,omitempty"`
	GiphyId *string `json:"giphyId,omitempty"`
	IsFree *bool `json:"isFree,omitempty"`
	IsMediaReady *bool `json:"isMediaReady,omitempty"`
	MediaCount *float64 `json:"mediaCount,omitempty"`
	Media []struct {
	Id float64 `json:"id"`
	Type string `json:"type"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CanView bool `json:"canView"`
	HasError bool `json:"hasError"`
	CreatedAt string `json:"createdAt"`
	IsReady bool `json:"isReady"`
	Duration *float64 `json:"duration,omitempty"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	View string `json:"view"`
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs"`
	IvStatus string `json:"ivStatus"`
	IsFromGuest bool `json:"isFromGuest"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
	Files *struct {
	Full struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"sources"`
} `json:"full"`
	Thumb *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"thumb,omitempty"`
	Preview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Options []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"options"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"squarePreview,omitempty"`
} `json:"files,omitempty"`
} `json:"media"`
	Previews []struct {
	Id float64 `json:"id"`
	Type string `json:"type"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CanView bool `json:"canView"`
	HasError bool `json:"hasError"`
	CreatedAt string `json:"createdAt"`
	IsReady bool `json:"isReady"`
	Duration *float64 `json:"duration,omitempty"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	View string `json:"view"`
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs"`
	IvStatus string `json:"ivStatus"`
	IsFromGuest bool `json:"isFromGuest"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
	Files *struct {
	Full struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"sources"`
} `json:"full"`
	Thumb *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"thumb,omitempty"`
	Preview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Options []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"options"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"squarePreview,omitempty"`
} `json:"files,omitempty"`
} `json:"previews"`
	IsTip *bool `json:"isTip,omitempty"`
	IsReportedByMe *bool `json:"isReportedByMe,omitempty"`
	ViewedCount *float64 `json:"viewedCount,omitempty"`
	SentCount *float64 `json:"sentCount,omitempty"`
	IsCanceled *bool `json:"isCanceled,omitempty"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	View string `json:"view"`
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs"`
	IvStatus string `json:"ivStatus"`
	IsFromGuest bool `json:"isFromGuest"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	Template *string `json:"template,omitempty"`
	CanUnsend *bool `json:"canUnsend,omitempty"`
	UnsendSeconds *float64 `json:"unsendSeconds,omitempty"`
	Price *string `json:"price,omitempty"`
	PurchasedCount *float64 `json:"purchasedCount,omitempty"`
	CanSendMessageToBuyers *bool `json:"canSendMessageToBuyers,omitempty"`
} `json:"items"`
}

// ListAccessAnalyticsMassMessagesBuyersResponse represents the response for ListAccessAnalyticsMassMessagesBuyers
type ListAccessAnalyticsMassMessagesBuyersResponse struct {
	List []struct {
	Id float64 `json:"id"`
	Username string `json:"username"`
	Name string `json:"name"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs *struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs,omitempty"`
} `json:"list"`
	HasMore bool `json:"hasMore"`
	NextMarker *float64 `json:"nextMarker,omitempty"`
}

// GetAccessAnalyticsPromotionsChartResponse represents the response for GetAccessAnalyticsPromotionsChart
type GetAccessAnalyticsPromotionsChartResponse map[string]struct {
	Chart []struct {
	Date string `json:"date"`
	Count float64 `json:"count"`
} `json:"chart"`
	Total float64 `json:"total"`
	Delta float64 `json:"delta"`
	Gross *float64 `json:"gross,omitempty"`
}

// GetAccessAnalyticsPromotionsTopResponse represents the response for GetAccessAnalyticsPromotionsTop
type GetAccessAnalyticsPromotionsTopResponse struct {
	HasMore bool `json:"hasMore"`
	Items []struct {
	Id float64 `json:"id"`
	Message string `json:"message"`
	RawMessage string `json:"rawMessage"`
	HasRelatedPromo bool `json:"hasRelatedPromo"`
	Price float64 `json:"price"`
	Type string `json:"type"`
	CanClaim bool `json:"canClaim"`
	ClaimsCount float64 `json:"claimsCount"`
	SubscribeCounts float64 `json:"subscribeCounts"`
	SubscribeDays float64 `json:"subscribeDays"`
	CreatedAt string `json:"createdAt"`
	FinishedAt string `json:"finishedAt"`
	IsFinished bool `json:"isFinished"`
} `json:"items"`
}

// GetAccessAnalyticsTrialsChartResponse represents the response for GetAccessAnalyticsTrialsChart
type GetAccessAnalyticsTrialsChartResponse map[string]struct {
	Chart []struct {
	Date string `json:"date"`
	Count float64 `json:"count"`
} `json:"chart"`
	Total float64 `json:"total"`
	Delta float64 `json:"delta"`
	Gross *float64 `json:"gross,omitempty"`
}

// GetAccessAnalyticsTrialsTopResponse represents the response for GetAccessAnalyticsTrialsTop
type GetAccessAnalyticsTrialsTopResponse struct {
	HasMore bool `json:"hasMore"`
	Items []struct {
	Id float64 `json:"id"`
	TrialLinkName string `json:"trialLinkName"`
	Url string `json:"url"`
	SubscribeDays float64 `json:"subscribeDays"`
	SubscribeCounts float64 `json:"subscribeCounts"`
	ClaimCounts float64 `json:"claimCounts"`
	ExpiredAt string `json:"expiredAt"`
	CreatedAt string `json:"createdAt"`
	IsFinished bool `json:"isFinished"`
} `json:"items"`
}

// GetAccessAnalyticsCampaignsChartResponse represents the response for GetAccessAnalyticsCampaignsChart
type GetAccessAnalyticsCampaignsChartResponse map[string]struct {
	Chart []struct {
	Date string `json:"date"`
	Count float64 `json:"count"`
} `json:"chart"`
	Total float64 `json:"total"`
	Delta float64 `json:"delta"`
	Gross *float64 `json:"gross,omitempty"`
}

// GetAccessAnalyticsCampaignsTopResponse represents the response for GetAccessAnalyticsCampaignsTop
type GetAccessAnalyticsCampaignsTopResponse struct {
	HasMore bool `json:"hasMore"`
	List []struct {
	Id float64 `json:"id"`
	CountSubscribers float64 `json:"countSubscribers"`
	CountTransitions float64 `json:"countTransitions"`
	CampaignCode float64 `json:"campaignCode"`
	CampaignName string `json:"campaignName"`
	CreatedAt string `json:"createdAt"`
	EndDate string `json:"endDate"`
} `json:"list"`
}

// GetAccessAnalyticsVisitorCountriesChartResponse represents the response for GetAccessAnalyticsVisitorCountriesChart
type GetAccessAnalyticsVisitorCountriesChartResponse struct {
	IsAvailable bool `json:"isAvailable"`
	Chart struct {
	Visitors []struct {
	Date string `json:"date"`
	Count float64 `json:"count"`
} `json:"visitors"`
	Duration []struct {
	Date string `json:"date"`
	Count float64 `json:"count"`
} `json:"duration"`
} `json:"chart"`
	Total struct {
	Current string `json:"current"`
	Delta float64 `json:"delta"`
} `json:"total"`
	HasStats bool `json:"hasStats"`
}

// GetAccessAnalyticsVisitorCountriesTopResponse represents the response for GetAccessAnalyticsVisitorCountriesTop
type GetAccessAnalyticsVisitorCountriesTopResponse struct {
	IsAvailable bool `json:"isAvailable"`
	HasStats bool `json:"hasStats"`
	TopCountries struct {
	HasMore bool `json:"hasMore"`
	Totals struct {
	Total float64 `json:"total"`
	Guests string `json:"guests"`
	Users string `json:"users"`
	Subscribers float64 `json:"subscribers"`
} `json:"totals"`
	Rows []struct {
	Rank float64 `json:"rank"`
	CountryName string `json:"countryName"`
	CountryCode string `json:"countryCode"`
	ViewsCount struct {
	Total float64 `json:"total"`
	Guests float64 `json:"guests"`
	Users float64 `json:"users"`
	Subscribers float64 `json:"subscribers"`
} `json:"viewsCount"`
} `json:"rows"`
} `json:"topCountries"`
}

// ListAccessPostsResponse represents the response for ListAccessPosts
type ListAccessPostsResponse struct {
	List []struct {
	Id float64 `json:"id"`
	CanDelete bool `json:"canDelete"`
	CanEdit bool `json:"canEdit"`
	MediaCount float64 `json:"mediaCount"`
	Media []struct {
	Id float64 `json:"id"`
	Type string `json:"type"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CanView bool `json:"canView"`
	HasError bool `json:"hasError"`
	CreatedAt string `json:"createdAt"`
	IsReady bool `json:"isReady"`
	Duration *float64 `json:"duration,omitempty"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	View string `json:"view"`
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs"`
	IvStatus string `json:"ivStatus"`
	IsFromGuest bool `json:"isFromGuest"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
	Files *struct {
	Full struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"sources"`
} `json:"full"`
	Thumb *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"thumb,omitempty"`
	Preview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Options []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"options"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"squarePreview,omitempty"`
} `json:"files,omitempty"`
} `json:"media"`
	CanViewMedia bool `json:"canViewMedia"`
	Author *struct {
	Id float64 `json:"id"`
	View string `json:"_view"`
} `json:"author,omitempty"`
	ResponseType string `json:"responseType"`
	PostedAt string `json:"postedAt"`
	PostedAtPrecise string `json:"postedAtPrecise"`
	IsMarkdownDisabled bool `json:"isMarkdownDisabled"`
	IsOpened bool `json:"isOpened"`
	CanToggleFavorite bool `json:"canToggleFavorite"`
	TipsAmount string `json:"tipsAmount"`
	Text string `json:"text"`
	IsFavorite bool `json:"isFavorite"`
	CanComment bool `json:"canComment"`
	FavoritesCount float64 `json:"favoritesCount"`
	IsMediaReady bool `json:"isMediaReady"`
	RawText string `json:"rawText"`
} `json:"list"`
	HasMore bool `json:"hasMore"`
	HeadMarker string `json:"headMarker"`
	TailMarker string `json:"tailMarker"`
	Counters struct {
	AudiosCount float64 `json:"audiosCount"`
	PhotosCount float64 `json:"photosCount"`
	VideosCount float64 `json:"videosCount"`
	MediasCount float64 `json:"mediasCount"`
	PostsCount float64 `json:"postsCount"`
	StreamsCount float64 `json:"streamsCount"`
	ArchivedPostsCount float64 `json:"archivedPostsCount"`
	PrivateArchivedPostsCount float64 `json:"privateArchivedPostsCount"`
} `json:"counters"`
}

// CreateAccessPostsResponse represents the response for CreateAccessPosts
type CreateAccessPostsResponse struct {
	Id float64 `json:"id"`
	CanDelete bool `json:"canDelete"`
	CanEdit bool `json:"canEdit"`
	MediaCount float64 `json:"mediaCount"`
	Media []struct {
	Id float64 `json:"id"`
	Type string `json:"type"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CanView bool `json:"canView"`
	HasError bool `json:"hasError"`
	CreatedAt string `json:"createdAt"`
	IsReady bool `json:"isReady"`
	Duration *float64 `json:"duration,omitempty"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	View string `json:"view"`
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs"`
	IvStatus string `json:"ivStatus"`
	IsFromGuest bool `json:"isFromGuest"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
	Files *struct {
	Full struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"sources"`
} `json:"full"`
	Thumb *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"thumb,omitempty"`
	Preview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Options []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"options"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"squarePreview,omitempty"`
} `json:"files,omitempty"`
} `json:"media"`
	CanViewMedia bool `json:"canViewMedia"`
	Author *struct {
	Id float64 `json:"id"`
	View string `json:"_view"`
} `json:"author,omitempty"`
	ResponseType string `json:"responseType"`
	PostedAt string `json:"postedAt"`
	PostedAtPrecise string `json:"postedAtPrecise"`
	IsMarkdownDisabled bool `json:"isMarkdownDisabled"`
	IsOpened bool `json:"isOpened"`
	CanToggleFavorite bool `json:"canToggleFavorite"`
	TipsAmount string `json:"tipsAmount"`
	Text string `json:"text"`
	IsFavorite bool `json:"isFavorite"`
	CanComment bool `json:"canComment"`
	FavoritesCount float64 `json:"favoritesCount"`
	IsMediaReady bool `json:"isMediaReady"`
	RawText string `json:"rawText"`
}

// ReplaceAccessPostsResponse represents the response for ReplaceAccessPosts
type ReplaceAccessPostsResponse struct {
	Id float64 `json:"id"`
	CanDelete bool `json:"canDelete"`
	CanEdit bool `json:"canEdit"`
	MediaCount float64 `json:"mediaCount"`
	Media []struct {
	Id float64 `json:"id"`
	Type string `json:"type"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CanView bool `json:"canView"`
	HasError bool `json:"hasError"`
	CreatedAt string `json:"createdAt"`
	IsReady bool `json:"isReady"`
	Duration *float64 `json:"duration,omitempty"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	View string `json:"view"`
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs"`
	IvStatus string `json:"ivStatus"`
	IsFromGuest bool `json:"isFromGuest"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
	Files *struct {
	Full struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"sources"`
} `json:"full"`
	Thumb *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"thumb,omitempty"`
	Preview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Options []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"options"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"squarePreview,omitempty"`
} `json:"files,omitempty"`
} `json:"media"`
	CanViewMedia bool `json:"canViewMedia"`
	Author *struct {
	Id float64 `json:"id"`
	View string `json:"_view"`
} `json:"author,omitempty"`
	ResponseType string `json:"responseType"`
	PostedAt string `json:"postedAt"`
	PostedAtPrecise string `json:"postedAtPrecise"`
	IsMarkdownDisabled bool `json:"isMarkdownDisabled"`
	IsOpened bool `json:"isOpened"`
	CanToggleFavorite bool `json:"canToggleFavorite"`
	TipsAmount string `json:"tipsAmount"`
	Text string `json:"text"`
	IsFavorite bool `json:"isFavorite"`
	CanComment bool `json:"canComment"`
	FavoritesCount float64 `json:"favoritesCount"`
	IsMediaReady bool `json:"isMediaReady"`
	RawText string `json:"rawText"`
}

// DeleteAccessPostsResponse represents the response for DeleteAccessPosts
type DeleteAccessPostsResponse map[string]interface{}

// ListAccessUsersPostsResponse represents the response for ListAccessUsersPosts
type ListAccessUsersPostsResponse struct {
	List []struct {
	Id float64 `json:"id"`
	CanDelete bool `json:"canDelete"`
	CanEdit bool `json:"canEdit"`
	MediaCount float64 `json:"mediaCount"`
	Media []struct {
	Id float64 `json:"id"`
	Type string `json:"type"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CanView bool `json:"canView"`
	HasError bool `json:"hasError"`
	CreatedAt string `json:"createdAt"`
	IsReady bool `json:"isReady"`
	Duration *float64 `json:"duration,omitempty"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	View string `json:"view"`
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs"`
	IvStatus string `json:"ivStatus"`
	IsFromGuest bool `json:"isFromGuest"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
	Files *struct {
	Full struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"sources"`
} `json:"full"`
	Thumb *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"thumb,omitempty"`
	Preview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Options []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"options"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"squarePreview,omitempty"`
} `json:"files,omitempty"`
} `json:"media"`
	CanViewMedia bool `json:"canViewMedia"`
	Author *struct {
	Id float64 `json:"id"`
	View string `json:"_view"`
} `json:"author,omitempty"`
	ResponseType string `json:"responseType"`
	PostedAt string `json:"postedAt"`
	PostedAtPrecise string `json:"postedAtPrecise"`
	IsMarkdownDisabled bool `json:"isMarkdownDisabled"`
	IsOpened bool `json:"isOpened"`
	CanToggleFavorite bool `json:"canToggleFavorite"`
	TipsAmount string `json:"tipsAmount"`
	Text string `json:"text"`
	IsFavorite bool `json:"isFavorite"`
	CanComment bool `json:"canComment"`
	FavoritesCount float64 `json:"favoritesCount"`
	IsMediaReady bool `json:"isMediaReady"`
	RawText string `json:"rawText"`
} `json:"list"`
	HasMore bool `json:"hasMore"`
	HeadMarker string `json:"headMarker"`
	TailMarker string `json:"tailMarker"`
	Counters struct {
	AudiosCount float64 `json:"audiosCount"`
	PhotosCount float64 `json:"photosCount"`
	VideosCount float64 `json:"videosCount"`
	MediasCount float64 `json:"mediasCount"`
	PostsCount float64 `json:"postsCount"`
	StreamsCount float64 `json:"streamsCount"`
	ArchivedPostsCount float64 `json:"archivedPostsCount"`
	PrivateArchivedPostsCount float64 `json:"privateArchivedPostsCount"`
} `json:"counters"`
}

// ListAccessChatsMessagesResponse represents the response for ListAccessChatsMessages
type ListAccessChatsMessagesResponse struct {
	List []struct {
	Id float64 `json:"id"`
	Text string `json:"text"`
	FromUser interface{} `json:"fromUser"`
	Media []struct {
	Id float64 `json:"id"`
	Type string `json:"type"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CanView bool `json:"canView"`
	HasError bool `json:"hasError"`
	CreatedAt string `json:"createdAt"`
	IsReady bool `json:"isReady"`
	Duration *float64 `json:"duration,omitempty"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	View string `json:"view"`
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs"`
	IvStatus string `json:"ivStatus"`
	IsFromGuest bool `json:"isFromGuest"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
	Files *struct {
	Full struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"sources"`
} `json:"full"`
	Thumb *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"thumb,omitempty"`
	Preview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Options []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"options"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"squarePreview,omitempty"`
} `json:"files,omitempty"`
} `json:"media"`
	IsMediaReady *bool `json:"isMediaReady,omitempty"`
	MediaCount *float64 `json:"mediaCount,omitempty"`
	Previews []float64 `json:"previews"`
	LockedText *bool `json:"lockedText,omitempty"`
	ResponseType *string `json:"responseType,omitempty"`
	GiphyId *string `json:"giphyId,omitempty"`
	IsFree *bool `json:"isFree,omitempty"`
	IsTip *bool `json:"isTip,omitempty"`
	IsReportedByMe *bool `json:"isReportedByMe,omitempty"`
	IsCouplePeopleMedia *bool `json:"isCouplePeopleMedia,omitempty"`
	QueueId *float64 `json:"queueId,omitempty"`
	IsMarkdownDisabled *bool `json:"isMarkdownDisabled,omitempty"`
	ReleaseForms *interface{} `json:"releaseForms,omitempty"`
	IsFromQueue *bool `json:"isFromQueue,omitempty"`
	CanUnsendQueue *bool `json:"canUnsendQueue,omitempty"`
	UnsendSecondsQueue *float64 `json:"unsendSecondsQueue,omitempty"`
	IsOpened *bool `json:"isOpened,omitempty"`
	IsNew *bool `json:"isNew,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	ChangedAt *string `json:"changedAt,omitempty"`
	CancelSeconds *float64 `json:"cancelSeconds,omitempty"`
	IsLiked *bool `json:"isLiked,omitempty"`
	CanPurchase *bool `json:"canPurchase,omitempty"`
	CanPurchaseReason *string `json:"canPurchaseReason,omitempty"`
	CanReport *bool `json:"canReport,omitempty"`
	CanBePinned *bool `json:"canBePinned,omitempty"`
	IsPinned *bool `json:"isPinned,omitempty"`
} `json:"list"`
	HasMore bool `json:"hasMore"`
}

// CreateAccessChatsMessagesResponse represents the response for CreateAccessChatsMessages
type CreateAccessChatsMessagesResponse struct {
	Id float64 `json:"id"`
	Text string `json:"text"`
	FromUser interface{} `json:"fromUser"`
	Media []struct {
	Id float64 `json:"id"`
	Type string `json:"type"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CanView bool `json:"canView"`
	HasError bool `json:"hasError"`
	CreatedAt string `json:"createdAt"`
	IsReady bool `json:"isReady"`
	Duration *float64 `json:"duration,omitempty"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	View string `json:"view"`
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs"`
	IvStatus string `json:"ivStatus"`
	IsFromGuest bool `json:"isFromGuest"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
	Files *struct {
	Full struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"sources"`
} `json:"full"`
	Thumb *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"thumb,omitempty"`
	Preview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Options []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"options"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"squarePreview,omitempty"`
} `json:"files,omitempty"`
} `json:"media"`
	IsMediaReady *bool `json:"isMediaReady,omitempty"`
	MediaCount *float64 `json:"mediaCount,omitempty"`
	Previews []float64 `json:"previews"`
	LockedText *bool `json:"lockedText,omitempty"`
	ResponseType *string `json:"responseType,omitempty"`
	GiphyId *string `json:"giphyId,omitempty"`
	IsFree *bool `json:"isFree,omitempty"`
	IsTip *bool `json:"isTip,omitempty"`
	IsReportedByMe *bool `json:"isReportedByMe,omitempty"`
	IsCouplePeopleMedia *bool `json:"isCouplePeopleMedia,omitempty"`
	QueueId *float64 `json:"queueId,omitempty"`
	IsMarkdownDisabled *bool `json:"isMarkdownDisabled,omitempty"`
	ReleaseForms *interface{} `json:"releaseForms,omitempty"`
	IsFromQueue *bool `json:"isFromQueue,omitempty"`
	CanUnsendQueue *bool `json:"canUnsendQueue,omitempty"`
	UnsendSecondsQueue *float64 `json:"unsendSecondsQueue,omitempty"`
	IsOpened *bool `json:"isOpened,omitempty"`
	IsNew *bool `json:"isNew,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	ChangedAt *string `json:"changedAt,omitempty"`
	CancelSeconds *float64 `json:"cancelSeconds,omitempty"`
	IsLiked *bool `json:"isLiked,omitempty"`
	CanPurchase *bool `json:"canPurchase,omitempty"`
	CanPurchaseReason *string `json:"canPurchaseReason,omitempty"`
	CanReport *bool `json:"canReport,omitempty"`
	CanBePinned *bool `json:"canBePinned,omitempty"`
	IsPinned *bool `json:"isPinned,omitempty"`
}

// DeleteAccessChatsMessagesResponse represents the response for DeleteAccessChatsMessages
type DeleteAccessChatsMessagesResponse struct {
	Success bool `json:"success"`
	Queue *struct {
	Id float64 `json:"id"`
	Text string `json:"text"`
	FromUser interface{} `json:"fromUser"`
	Media []struct {
	Id float64 `json:"id"`
	Type string `json:"type"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CanView bool `json:"canView"`
	HasError bool `json:"hasError"`
	CreatedAt string `json:"createdAt"`
	IsReady bool `json:"isReady"`
	Duration *float64 `json:"duration,omitempty"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	View string `json:"view"`
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs"`
	IvStatus string `json:"ivStatus"`
	IsFromGuest bool `json:"isFromGuest"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
	Files *struct {
	Full struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"sources"`
} `json:"full"`
	Thumb *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"thumb,omitempty"`
	Preview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Options []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"options"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"squarePreview,omitempty"`
} `json:"files,omitempty"`
} `json:"media"`
	IsMediaReady *bool `json:"isMediaReady,omitempty"`
	MediaCount *float64 `json:"mediaCount,omitempty"`
	Previews []float64 `json:"previews"`
	LockedText *bool `json:"lockedText,omitempty"`
	ResponseType *string `json:"responseType,omitempty"`
	GiphyId *string `json:"giphyId,omitempty"`
	IsFree *bool `json:"isFree,omitempty"`
	IsTip *bool `json:"isTip,omitempty"`
	IsReportedByMe *bool `json:"isReportedByMe,omitempty"`
	IsCouplePeopleMedia *bool `json:"isCouplePeopleMedia,omitempty"`
	QueueId *float64 `json:"queueId,omitempty"`
	IsMarkdownDisabled *bool `json:"isMarkdownDisabled,omitempty"`
	ReleaseForms *interface{} `json:"releaseForms,omitempty"`
	IsFromQueue *bool `json:"isFromQueue,omitempty"`
	CanUnsendQueue *bool `json:"canUnsendQueue,omitempty"`
	UnsendSecondsQueue *float64 `json:"unsendSecondsQueue,omitempty"`
	IsOpened *bool `json:"isOpened,omitempty"`
	IsNew *bool `json:"isNew,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	ChangedAt *string `json:"changedAt,omitempty"`
	CancelSeconds *float64 `json:"cancelSeconds,omitempty"`
	IsLiked *bool `json:"isLiked,omitempty"`
	CanPurchase *bool `json:"canPurchase,omitempty"`
	CanPurchaseReason *string `json:"canPurchaseReason,omitempty"`
	CanReport *bool `json:"canReport,omitempty"`
	CanBePinned *bool `json:"canBePinned,omitempty"`
	IsPinned *bool `json:"isPinned,omitempty"`
	Date *string `json:"date,omitempty"`
	TextCropped *string `json:"textCropped,omitempty"`
	SentCount *float64 `json:"sentCount,omitempty"`
	ViewedCount *float64 `json:"viewedCount,omitempty"`
	CanUnsend *bool `json:"canUnsend,omitempty"`
	UnsendSeconds *float64 `json:"unsendSeconds,omitempty"`
	IsCanceled *bool `json:"isCanceled,omitempty"`
	MediaTypes *struct {
	Video *float64 `json:"video,omitempty"`
	Photo *float64 `json:"photo,omitempty"`
	Gif *float64 `json:"gif,omitempty"`
	Audio *float64 `json:"audio,omitempty"`
} `json:"mediaTypes,omitempty"`
	HasError *bool `json:"hasError,omitempty"`
	Price *string `json:"price,omitempty"`
	PurchasedCount *float64 `json:"purchasedCount,omitempty"`
	CanSendMessageToBuyers *bool `json:"canSendMessageToBuyers,omitempty"`
} `json:"queue,omitempty"`
}

// CreateAccessMassMessagesResponse represents the response for CreateAccessMassMessages
type CreateAccessMassMessagesResponse struct {
	Id float64 `json:"id"`
	Date string `json:"date"`
	IsReady *bool `json:"isReady,omitempty"`
	IsDone *bool `json:"isDone,omitempty"`
	Total *float64 `json:"total,omitempty"`
	Pending *float64 `json:"pending,omitempty"`
	CanUnsend *bool `json:"canUnsend,omitempty"`
	UnsendSeconds *float64 `json:"unsendSeconds,omitempty"`
	HasError *bool `json:"hasError,omitempty"`
	IsCanceled *bool `json:"isCanceled,omitempty"`
	IsCouplePeopleMedia *bool `json:"isCouplePeopleMedia,omitempty"`
}

// GetAccessMassMessagesResponse represents the response for GetAccessMassMessages
type GetAccessMassMessagesResponse struct {
	Id float64 `json:"id"`
	Text string `json:"text"`
	FromUser interface{} `json:"fromUser"`
	Media []struct {
	Id float64 `json:"id"`
	Type string `json:"type"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CanView bool `json:"canView"`
	HasError bool `json:"hasError"`
	CreatedAt string `json:"createdAt"`
	IsReady bool `json:"isReady"`
	Duration *float64 `json:"duration,omitempty"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	View string `json:"view"`
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs"`
	IvStatus string `json:"ivStatus"`
	IsFromGuest bool `json:"isFromGuest"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
	Files *struct {
	Full struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"sources"`
} `json:"full"`
	Thumb *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"thumb,omitempty"`
	Preview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Options []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"options"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"squarePreview,omitempty"`
} `json:"files,omitempty"`
} `json:"media"`
	IsMediaReady *bool `json:"isMediaReady,omitempty"`
	MediaCount *float64 `json:"mediaCount,omitempty"`
	Previews []float64 `json:"previews"`
	LockedText *bool `json:"lockedText,omitempty"`
	ResponseType *string `json:"responseType,omitempty"`
	GiphyId *string `json:"giphyId,omitempty"`
	IsFree *bool `json:"isFree,omitempty"`
	IsTip *bool `json:"isTip,omitempty"`
	IsReportedByMe *bool `json:"isReportedByMe,omitempty"`
	IsCouplePeopleMedia *bool `json:"isCouplePeopleMedia,omitempty"`
	QueueId *float64 `json:"queueId,omitempty"`
	IsMarkdownDisabled *bool `json:"isMarkdownDisabled,omitempty"`
	ReleaseForms *interface{} `json:"releaseForms,omitempty"`
	IsFromQueue *bool `json:"isFromQueue,omitempty"`
	CanUnsendQueue *bool `json:"canUnsendQueue,omitempty"`
	UnsendSecondsQueue *float64 `json:"unsendSecondsQueue,omitempty"`
	IsOpened *bool `json:"isOpened,omitempty"`
	IsNew *bool `json:"isNew,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	ChangedAt *string `json:"changedAt,omitempty"`
	CancelSeconds *float64 `json:"cancelSeconds,omitempty"`
	IsLiked *bool `json:"isLiked,omitempty"`
	CanPurchase *bool `json:"canPurchase,omitempty"`
	CanPurchaseReason *string `json:"canPurchaseReason,omitempty"`
	CanReport *bool `json:"canReport,omitempty"`
	CanBePinned *bool `json:"canBePinned,omitempty"`
	IsPinned *bool `json:"isPinned,omitempty"`
}

// ReplaceAccessMassMessagesResponse represents the response for ReplaceAccessMassMessages
type ReplaceAccessMassMessagesResponse struct {
	Id float64 `json:"id"`
	Date string `json:"date"`
	IsReady *bool `json:"isReady,omitempty"`
	IsDone *bool `json:"isDone,omitempty"`
	Total *float64 `json:"total,omitempty"`
	Pending *float64 `json:"pending,omitempty"`
	CanUnsend *bool `json:"canUnsend,omitempty"`
	UnsendSeconds *float64 `json:"unsendSeconds,omitempty"`
	HasError *bool `json:"hasError,omitempty"`
	IsCanceled *bool `json:"isCanceled,omitempty"`
	IsCouplePeopleMedia *bool `json:"isCouplePeopleMedia,omitempty"`
}

// DeleteAccessMassMessagesResponse represents the response for DeleteAccessMassMessages
type DeleteAccessMassMessagesResponse struct {
	Success bool `json:"success"`
	Queue *struct {
	Id float64 `json:"id"`
	Text string `json:"text"`
	FromUser interface{} `json:"fromUser"`
	Media []struct {
	Id float64 `json:"id"`
	Type string `json:"type"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CanView bool `json:"canView"`
	HasError bool `json:"hasError"`
	CreatedAt string `json:"createdAt"`
	IsReady bool `json:"isReady"`
	Duration *float64 `json:"duration,omitempty"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	View string `json:"view"`
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs"`
	IvStatus string `json:"ivStatus"`
	IsFromGuest bool `json:"isFromGuest"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
	Files *struct {
	Full struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"sources"`
} `json:"full"`
	Thumb *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"thumb,omitempty"`
	Preview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Options []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"options"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"squarePreview,omitempty"`
} `json:"files,omitempty"`
} `json:"media"`
	IsMediaReady *bool `json:"isMediaReady,omitempty"`
	MediaCount *float64 `json:"mediaCount,omitempty"`
	Previews []float64 `json:"previews"`
	LockedText *bool `json:"lockedText,omitempty"`
	ResponseType *string `json:"responseType,omitempty"`
	GiphyId *string `json:"giphyId,omitempty"`
	IsFree *bool `json:"isFree,omitempty"`
	IsTip *bool `json:"isTip,omitempty"`
	IsReportedByMe *bool `json:"isReportedByMe,omitempty"`
	IsCouplePeopleMedia *bool `json:"isCouplePeopleMedia,omitempty"`
	QueueId *float64 `json:"queueId,omitempty"`
	IsMarkdownDisabled *bool `json:"isMarkdownDisabled,omitempty"`
	ReleaseForms *interface{} `json:"releaseForms,omitempty"`
	IsFromQueue *bool `json:"isFromQueue,omitempty"`
	CanUnsendQueue *bool `json:"canUnsendQueue,omitempty"`
	UnsendSecondsQueue *float64 `json:"unsendSecondsQueue,omitempty"`
	IsOpened *bool `json:"isOpened,omitempty"`
	IsNew *bool `json:"isNew,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	ChangedAt *string `json:"changedAt,omitempty"`
	CancelSeconds *float64 `json:"cancelSeconds,omitempty"`
	IsLiked *bool `json:"isLiked,omitempty"`
	CanPurchase *bool `json:"canPurchase,omitempty"`
	CanPurchaseReason *string `json:"canPurchaseReason,omitempty"`
	CanReport *bool `json:"canReport,omitempty"`
	CanBePinned *bool `json:"canBePinned,omitempty"`
	IsPinned *bool `json:"isPinned,omitempty"`
	Date *string `json:"date,omitempty"`
	TextCropped *string `json:"textCropped,omitempty"`
	SentCount *float64 `json:"sentCount,omitempty"`
	ViewedCount *float64 `json:"viewedCount,omitempty"`
	CanUnsend *bool `json:"canUnsend,omitempty"`
	UnsendSeconds *float64 `json:"unsendSeconds,omitempty"`
	IsCanceled *bool `json:"isCanceled,omitempty"`
	MediaTypes *struct {
	Video *float64 `json:"video,omitempty"`
	Photo *float64 `json:"photo,omitempty"`
	Gif *float64 `json:"gif,omitempty"`
	Audio *float64 `json:"audio,omitempty"`
} `json:"mediaTypes,omitempty"`
	HasError *bool `json:"hasError,omitempty"`
	Price *string `json:"price,omitempty"`
	PurchasedCount *float64 `json:"purchasedCount,omitempty"`
	CanSendMessageToBuyers *bool `json:"canSendMessageToBuyers,omitempty"`
} `json:"queue,omitempty"`
}

// ListAccessChatsResponse represents the response for ListAccessChats
type ListAccessChatsResponse struct {
	List []struct {
	WithUser struct {
	Id float64 `json:"id"`
	View string `json:"_view"`
} `json:"withUser"`
	CanNotSendReason *interface{} `json:"canNotSendReason,omitempty"`
	CanSendMessage *bool `json:"canSendMessage,omitempty"`
	CanGoToProfile *bool `json:"canGoToProfile,omitempty"`
	UnreadMessagesCount *float64 `json:"unreadMessagesCount,omitempty"`
	HasUnreadTips *bool `json:"hasUnreadTips,omitempty"`
	IsMutedNotifications *bool `json:"isMutedNotifications,omitempty"`
	LastMessage *struct {
	Id float64 `json:"id"`
	Text string `json:"text"`
	FromUser interface{} `json:"fromUser"`
	Media []struct {
	Id float64 `json:"id"`
	Type string `json:"type"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CanView bool `json:"canView"`
	HasError bool `json:"hasError"`
	CreatedAt string `json:"createdAt"`
	IsReady bool `json:"isReady"`
	Duration *float64 `json:"duration,omitempty"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	View string `json:"view"`
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs"`
	IvStatus string `json:"ivStatus"`
	IsFromGuest bool `json:"isFromGuest"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
	Files *struct {
	Full struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"sources"`
} `json:"full"`
	Thumb *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"thumb,omitempty"`
	Preview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Options []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"options"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"squarePreview,omitempty"`
} `json:"files,omitempty"`
} `json:"media"`
	IsMediaReady *bool `json:"isMediaReady,omitempty"`
	MediaCount *float64 `json:"mediaCount,omitempty"`
	Previews []float64 `json:"previews"`
	LockedText *bool `json:"lockedText,omitempty"`
	ResponseType *string `json:"responseType,omitempty"`
	GiphyId *string `json:"giphyId,omitempty"`
	IsFree *bool `json:"isFree,omitempty"`
	IsTip *bool `json:"isTip,omitempty"`
	IsReportedByMe *bool `json:"isReportedByMe,omitempty"`
	IsCouplePeopleMedia *bool `json:"isCouplePeopleMedia,omitempty"`
	QueueId *float64 `json:"queueId,omitempty"`
	IsMarkdownDisabled *bool `json:"isMarkdownDisabled,omitempty"`
	ReleaseForms *interface{} `json:"releaseForms,omitempty"`
	IsFromQueue *bool `json:"isFromQueue,omitempty"`
	CanUnsendQueue *bool `json:"canUnsendQueue,omitempty"`
	UnsendSecondsQueue *float64 `json:"unsendSecondsQueue,omitempty"`
	IsOpened *bool `json:"isOpened,omitempty"`
	IsNew *bool `json:"isNew,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	ChangedAt *string `json:"changedAt,omitempty"`
	CancelSeconds *float64 `json:"cancelSeconds,omitempty"`
	IsLiked *bool `json:"isLiked,omitempty"`
	CanPurchase *bool `json:"canPurchase,omitempty"`
	CanPurchaseReason *string `json:"canPurchaseReason,omitempty"`
	CanReport *bool `json:"canReport,omitempty"`
	CanBePinned *bool `json:"canBePinned,omitempty"`
	IsPinned *bool `json:"isPinned,omitempty"`
} `json:"lastMessage,omitempty"`
	LastReadMessageId *float64 `json:"lastReadMessageId,omitempty"`
	HasPurchasedFeed *bool `json:"hasPurchasedFeed,omitempty"`
	CountPinnedMessages *float64 `json:"countPinnedMessages,omitempty"`
} `json:"list"`
	HasMore bool `json:"hasMore"`
	NextOffset *float64 `json:"nextOffset,omitempty"`
}

// ListAccessChatsMediaResponse represents the response for ListAccessChatsMedia
type ListAccessChatsMediaResponse struct {
	List []struct {
	Id float64 `json:"id"`
	Text string `json:"text"`
	FromUser interface{} `json:"fromUser"`
	Media []struct {
	Id float64 `json:"id"`
	Type string `json:"type"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CanView bool `json:"canView"`
	HasError bool `json:"hasError"`
	CreatedAt string `json:"createdAt"`
	IsReady bool `json:"isReady"`
	Duration *float64 `json:"duration,omitempty"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	View string `json:"view"`
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs"`
	IvStatus string `json:"ivStatus"`
	IsFromGuest bool `json:"isFromGuest"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
	Files *struct {
	Full struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"sources"`
} `json:"full"`
	Thumb *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"thumb,omitempty"`
	Preview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Options []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"options"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"squarePreview,omitempty"`
} `json:"files,omitempty"`
} `json:"media"`
	IsMediaReady *bool `json:"isMediaReady,omitempty"`
	MediaCount *float64 `json:"mediaCount,omitempty"`
	Previews []float64 `json:"previews"`
	LockedText *bool `json:"lockedText,omitempty"`
	ResponseType *string `json:"responseType,omitempty"`
	GiphyId *string `json:"giphyId,omitempty"`
	IsFree *bool `json:"isFree,omitempty"`
	IsTip *bool `json:"isTip,omitempty"`
	IsReportedByMe *bool `json:"isReportedByMe,omitempty"`
	IsCouplePeopleMedia *bool `json:"isCouplePeopleMedia,omitempty"`
	QueueId *float64 `json:"queueId,omitempty"`
	IsMarkdownDisabled *bool `json:"isMarkdownDisabled,omitempty"`
	ReleaseForms *interface{} `json:"releaseForms,omitempty"`
	IsFromQueue *bool `json:"isFromQueue,omitempty"`
	CanUnsendQueue *bool `json:"canUnsendQueue,omitempty"`
	UnsendSecondsQueue *float64 `json:"unsendSecondsQueue,omitempty"`
	IsOpened *bool `json:"isOpened,omitempty"`
	IsNew *bool `json:"isNew,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	ChangedAt *string `json:"changedAt,omitempty"`
	CancelSeconds *float64 `json:"cancelSeconds,omitempty"`
	IsLiked *bool `json:"isLiked,omitempty"`
	CanPurchase *bool `json:"canPurchase,omitempty"`
	CanPurchaseReason *string `json:"canPurchaseReason,omitempty"`
	CanReport *bool `json:"canReport,omitempty"`
	CanBePinned *bool `json:"canBePinned,omitempty"`
	IsPinned *bool `json:"isPinned,omitempty"`
} `json:"list"`
	HasMore bool `json:"hasMore"`
}

// ListAccessSubscribersResponse represents the response for ListAccessSubscribers
type ListAccessSubscribersResponse struct {
	List []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	DisplayName string `json:"displayName"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs"`
	LastSeen string `json:"lastSeen"`
	Subscription struct {
	IsActive bool `json:"isActive"`
	IsExpired bool `json:"isExpired"`
	SubscribedAt string `json:"subscribedAt"`
	ExpiresAt string `json:"expiresAt"`
	RenewedAt string `json:"renewedAt"`
	Price float64 `json:"price"`
	RegularPrice float64 `json:"regularPrice"`
	DiscountPercent float64 `json:"discountPercent"`
	DiscountPeriod float64 `json:"discountPeriod"`
	DiscountStartedAt string `json:"discountStartedAt"`
	DiscountFinishedAt string `json:"discountFinishedAt"`
	History []struct {
	Id float64 `json:"id"`
	StartDate string `json:"startDate"`
	ExpireDate string `json:"expireDate"`
	Price float64 `json:"price"`
	RegularPrice float64 `json:"regularPrice"`
	Discount float64 `json:"discount"`
	Type string `json:"type"`
	Action string `json:"action"`
	IsCurrent bool `json:"isCurrent"`
} `json:"history"`
} `json:"subscription"`
	Spending struct {
	Total float64 `json:"total"`
	Tips float64 `json:"tips"`
	Subscriptions float64 `json:"subscriptions"`
	Messages float64 `json:"messages"`
	Posts float64 `json:"posts"`
	Streams float64 `json:"streams"`
} `json:"spending"`
	IsRestricted bool `json:"isRestricted"`
	IsBlocked bool `json:"isBlocked"`
	Capabilities struct {
	CanRestrict bool `json:"canRestrict"`
	CanBlock bool `json:"canBlock"`
	CanReport bool `json:"canReport"`
	CanUnsubscribe bool `json:"canUnsubscribe"`
	CanReceiveMessages bool `json:"canReceiveMessages"`
	CanSendTrial bool `json:"canSendTrial"`
} `json:"capabilities"`
	Lists []struct {
	Id interface{} `json:"id"`
	Name string `json:"name"`
} `json:"lists"`
} `json:"list"`
	HasMore bool `json:"hasMore"`
}

// SetAccessSubscribersNoteResponse represents the response for SetAccessSubscribersNote
type SetAccessSubscribersNoteResponse map[string]interface{}

// SetAccessSubscribersDiscountResponse represents the response for SetAccessSubscribersDiscount
type SetAccessSubscribersDiscountResponse map[string]interface{}

// SetAccessSubscribersCustomNameResponse represents the response for SetAccessSubscribersCustomName
type SetAccessSubscribersCustomNameResponse map[string]interface{}

// ListAccessSubscriptionsResponse represents the response for ListAccessSubscriptions
type ListAccessSubscriptionsResponse struct {
	List []struct {
	Id float64 `json:"id"`
	Username string `json:"username"`
	Name string `json:"name"`
	Avatar string `json:"avatar"`
	AvatarThumbs *struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs,omitempty"`
	Lists []struct {
	Id interface{} `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
} `json:"lists"`
	SubscribedAt string `json:"subscribedAt"`
	ExpiredAt string `json:"expiredAt"`
	RenewedAt string `json:"renewedAt"`
	IsActive bool `json:"isActive"`
	SubscriptionPrice float64 `json:"subscriptionPrice"`
	TipsSumm *float64 `json:"tipsSumm,omitempty"`
	SubscribesSumm *float64 `json:"subscribesSumm,omitempty"`
	MessagesSumm *float64 `json:"messagesSumm,omitempty"`
	PostsSumm *float64 `json:"postsSumm,omitempty"`
	StreamsSumm *float64 `json:"streamsSumm,omitempty"`
	TotalSumm *float64 `json:"totalSumm,omitempty"`
} `json:"list"`
	HasMore bool `json:"hasMore"`
}

// GetAccessSubscriptionsCountResponse represents the response for GetAccessSubscriptionsCount
type GetAccessSubscriptionsCountResponse struct {
	Subscriptions struct {
	Active float64 `json:"active"`
	Muted float64 `json:"muted"`
	Restricted float64 `json:"restricted"`
	Expired float64 `json:"expired"`
	Blocked float64 `json:"blocked"`
	Attention float64 `json:"attention"`
	All float64 `json:"all"`
} `json:"subscriptions"`
	Subscribers struct {
	Active float64 `json:"active"`
	Muted float64 `json:"muted"`
	Restricted float64 `json:"restricted"`
	Expired float64 `json:"expired"`
	Blocked float64 `json:"blocked"`
	All float64 `json:"all"`
	ActiveOnline float64 `json:"activeOnline"`
} `json:"subscribers"`
	Bookmarks float64 `json:"bookmarks"`
}

// GetAccessSubscriptionsHistoryResponse represents the response for GetAccessSubscriptionsHistory
type GetAccessSubscriptionsHistoryResponse struct {
	List []struct {
	SubscribeDate string `json:"subscribeDate"`
	ExpireDate string `json:"expireDate"`
	Price float64 `json:"price"`
} `json:"list"`
	HasMore bool `json:"hasMore"`
}

// ListAccessPromotionsTrackingLinksResponse represents the response for ListAccessPromotionsTrackingLinks
type ListAccessPromotionsTrackingLinksResponse struct {
	List []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
} `json:"list"`
	HasMore bool `json:"hasMore"`
}

// CreateAccessPromotionsTrackingLinksShareAccessResponse represents the response for CreateAccessPromotionsTrackingLinksShareAccess
type CreateAccessPromotionsTrackingLinksShareAccessResponse struct {
	Success bool `json:"success"`
}

// DeleteAccessPromotionsTrackingLinksShareAccessResponse represents the response for DeleteAccessPromotionsTrackingLinksShareAccess
type DeleteAccessPromotionsTrackingLinksShareAccessResponse struct {
	Success bool `json:"success"`
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
	List []interface{} `json:"list"`
	HasMore *bool `json:"hasMore,omitempty"`
}

// CreateAccessPromotionsTrialLinksResponse represents the response for CreateAccessPromotionsTrialLinks
type CreateAccessPromotionsTrialLinksResponse struct {
	Id float64 `json:"id"`
	TrialLinkName string `json:"trialLinkName"`
	Url string `json:"url"`
	SubscribeDays float64 `json:"subscribeDays"`
	SubscribeCounts float64 `json:"subscribeCounts"`
	ClaimCounts float64 `json:"claimCounts"`
	ClicksCounts float64 `json:"clicksCounts"`
	ExpiredAt string `json:"expiredAt"`
	CreatedAt string `json:"createdAt"`
	IsFinished bool `json:"isFinished"`
	User *interface{} `json:"user,omitempty"`
	SharedWith *string `json:"sharedWith,omitempty"`
}

// CreateAccessPromotionsTrialLinksShareAccessResponse represents the response for CreateAccessPromotionsTrialLinksShareAccess
type CreateAccessPromotionsTrialLinksShareAccessResponse struct {
	Success bool `json:"success"`
}

// DeleteAccessPromotionsTrialLinksShareAccessResponse represents the response for DeleteAccessPromotionsTrialLinksShareAccess
type DeleteAccessPromotionsTrialLinksShareAccessResponse struct {
	Success bool `json:"success"`
}

// GetAccessPromotionsTrialLinksResponse represents the response for GetAccessPromotionsTrialLinks
type GetAccessPromotionsTrialLinksResponse struct {
	Id float64 `json:"id"`
	TrialLinkName string `json:"trialLinkName"`
	Url string `json:"url"`
	SubscribeDays float64 `json:"subscribeDays"`
	SubscribeCounts float64 `json:"subscribeCounts"`
	ClaimCounts float64 `json:"claimCounts"`
	ClicksCounts float64 `json:"clicksCounts"`
	ExpiredAt string `json:"expiredAt"`
	CreatedAt string `json:"createdAt"`
	IsFinished bool `json:"isFinished"`
	User *interface{} `json:"user,omitempty"`
	SharedWith *string `json:"sharedWith,omitempty"`
}

// ReplaceAccessPromotionsTrialLinksResponse represents the response for ReplaceAccessPromotionsTrialLinks
type ReplaceAccessPromotionsTrialLinksResponse struct {
	Success bool `json:"success"`
}

// DeleteAccessPromotionsTrialLinksResponse represents the response for DeleteAccessPromotionsTrialLinks
type DeleteAccessPromotionsTrialLinksResponse struct {
	Success bool `json:"success"`
}

// CreateAccessPromotionsBundlesResponse represents the response for CreateAccessPromotionsBundles
type CreateAccessPromotionsBundlesResponse struct {
	Id float64 `json:"id"`
	Discount float64 `json:"discount"`
	Duration float64 `json:"duration"`
	Price float64 `json:"price"`
	CanBuy bool `json:"canBuy"`
}

// GetAccessPromotionsBundlesResponse represents the response for GetAccessPromotionsBundles
type GetAccessPromotionsBundlesResponse struct {
	Id float64 `json:"id"`
	Discount float64 `json:"discount"`
	Duration float64 `json:"duration"`
	Price float64 `json:"price"`
	CanBuy bool `json:"canBuy"`
}

// ReplaceAccessPromotionsBundlesResponse represents the response for ReplaceAccessPromotionsBundles
type ReplaceAccessPromotionsBundlesResponse struct {
	Id float64 `json:"id"`
	Discount float64 `json:"discount"`
	Duration float64 `json:"duration"`
	Price float64 `json:"price"`
	CanBuy bool `json:"canBuy"`
}

// DeleteAccessPromotionsBundlesResponse represents the response for DeleteAccessPromotionsBundles
type DeleteAccessPromotionsBundlesResponse struct {
	Id float64 `json:"id"`
	Discount float64 `json:"discount"`
	Duration float64 `json:"duration"`
	Price float64 `json:"price"`
	CanBuy bool `json:"canBuy"`
}

// ReplaceAccessPromotionsResponse represents the response for ReplaceAccessPromotions
type ReplaceAccessPromotionsResponse struct {
	Id float64 `json:"id"`
	Message string `json:"message"`
	RawMessage string `json:"rawMessage"`
	HasRelatedPromo bool `json:"hasRelatedPromo"`
	Price float64 `json:"price"`
	Type string `json:"type"`
	CanClaim bool `json:"canClaim"`
	ClaimsCount float64 `json:"claimsCount"`
	SubscribeCounts float64 `json:"subscribeCounts"`
	SubscribeDays float64 `json:"subscribeDays"`
	CreatedAt string `json:"createdAt"`
	FinishedAt string `json:"finishedAt"`
	IsFinished bool `json:"isFinished"`
}

// DeleteAccessPromotionsResponse represents the response for DeleteAccessPromotions
type DeleteAccessPromotionsResponse struct {
	Success bool `json:"success"`
}

// CreateAccessPromotionsStopResponse represents the response for CreateAccessPromotionsStop
type CreateAccessPromotionsStopResponse struct {
	Success bool `json:"success"`
}

// GetAccessUsersRestrictResponse represents the response for GetAccessUsersRestrict
type GetAccessUsersRestrictResponse struct {
	List []struct {
	Id float64 `json:"id"`
	Username string `json:"username"`
	Name string `json:"name"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs *struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs,omitempty"`
	Lists []struct {
	Id interface{} `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
} `json:"lists"`
	DisplayName *string `json:"displayName,omitempty"`
	Notice *string `json:"notice,omitempty"`
	IsRestricted *bool `json:"isRestricted,omitempty"`
	CanRestrict *bool `json:"canRestrict,omitempty"`
	SubscribedBy *bool `json:"subscribedBy,omitempty"`
	SubscribedByExpire *bool `json:"subscribedByExpire,omitempty"`
	SubscribedByExpireDate *string `json:"subscribedByExpireDate,omitempty"`
	SubscribedByAutoprolong *bool `json:"subscribedByAutoprolong,omitempty"`
	SubscribedIsExpiredNow *bool `json:"subscribedIsExpiredNow,omitempty"`
	CurrentSubscribePrice *float64 `json:"currentSubscribePrice,omitempty"`
	SubscribedOn *bool `json:"subscribedOn,omitempty"`
	SubscribedOnExpiredNow *bool `json:"subscribedOnExpiredNow,omitempty"`
	SubscribedOnDuration *string `json:"subscribedOnDuration,omitempty"`
	CanReport *bool `json:"canReport,omitempty"`
	CanReceiveChatMessage *bool `json:"canReceiveChatMessage,omitempty"`
	HideChat *bool `json:"hideChat,omitempty"`
	LastSeen *string `json:"lastSeen,omitempty"`
	IsPerformer *bool `json:"isPerformer,omitempty"`
	IsRealPerformer *bool `json:"isRealPerformer,omitempty"`
	SubscribedByData *struct {
	Price float64 `json:"price"`
	NewPrice float64 `json:"newPrice"`
	RegularPrice float64 `json:"regularPrice"`
	SubscribePrice float64 `json:"subscribePrice"`
	DiscountPercent float64 `json:"discountPercent"`
	DiscountPeriod float64 `json:"discountPeriod"`
	SubscribeAt string `json:"subscribeAt"`
	ExpiredAt string `json:"expiredAt"`
	RenewedAt string `json:"renewedAt"`
	DiscountFinishedAt string `json:"discountFinishedAt"`
	DiscountStartedAt string `json:"discountStartedAt"`
	Status string `json:"status"`
	IsMuted bool `json:"isMuted"`
	UnsubscribeReason string `json:"unsubscribeReason"`
	Duration string `json:"duration"`
	HasActivePaidSubscriptions bool `json:"hasActivePaidSubscriptions"`
	ShowPostsInFeed bool `json:"showPostsInFeed"`
	Subscribes []struct {
	Id float64 `json:"id"`
	UserId float64 `json:"userId"`
	SubscriberId float64 `json:"subscriberId"`
	Date string `json:"date"`
	Duration float64 `json:"duration"`
	StartDate string `json:"startDate"`
	ExpireDate string `json:"expireDate"`
	CancelDate string `json:"cancelDate"`
	Price float64 `json:"price"`
	RegularPrice float64 `json:"regularPrice"`
	Discount float64 `json:"discount"`
	EarningId float64 `json:"earningId"`
	Action string `json:"action"`
	Type string `json:"type"`
	OfferStart string `json:"offerStart"`
	OfferEnd string `json:"offerEnd"`
	IsCurrent bool `json:"isCurrent"`
} `json:"subscribes"`
} `json:"subscribedByData,omitempty"`
	SubscribedOnData *struct {
	Price float64 `json:"price"`
	NewPrice float64 `json:"newPrice"`
	RegularPrice float64 `json:"regularPrice"`
	SubscribePrice float64 `json:"subscribePrice"`
	DiscountPercent float64 `json:"discountPercent"`
	DiscountPeriod float64 `json:"discountPeriod"`
	SubscribeAt string `json:"subscribeAt"`
	ExpiredAt string `json:"expiredAt"`
	RenewedAt string `json:"renewedAt"`
	DiscountFinishedAt string `json:"discountFinishedAt"`
	DiscountStartedAt string `json:"discountStartedAt"`
	Status string `json:"status"`
	IsMuted bool `json:"isMuted"`
	UnsubscribeReason string `json:"unsubscribeReason"`
	Duration string `json:"duration"`
	HasActivePaidSubscriptions bool `json:"hasActivePaidSubscriptions"`
	Subscribes []struct {
	Id float64 `json:"id"`
	UserId float64 `json:"userId"`
	SubscriberId float64 `json:"subscriberId"`
	Date string `json:"date"`
	Duration float64 `json:"duration"`
	StartDate string `json:"startDate"`
	ExpireDate string `json:"expireDate"`
	CancelDate string `json:"cancelDate"`
	Price float64 `json:"price"`
	RegularPrice float64 `json:"regularPrice"`
	Discount float64 `json:"discount"`
	EarningId float64 `json:"earningId"`
	Action string `json:"action"`
	Type string `json:"type"`
	OfferStart string `json:"offerStart"`
	OfferEnd string `json:"offerEnd"`
	IsCurrent bool `json:"isCurrent"`
} `json:"subscribes"`
	TipsSumm float64 `json:"tipsSumm"`
	SubscribesSumm float64 `json:"subscribesSumm"`
	MessagesSumm float64 `json:"messagesSumm"`
	PostsSumm float64 `json:"postsSumm"`
	StreamsSumm float64 `json:"streamsSumm"`
	TotalSumm float64 `json:"totalSumm"`
} `json:"subscribedOnData,omitempty"`
	CanTrialSend *bool `json:"canTrialSend,omitempty"`
	IsBlocked *bool `json:"isBlocked,omitempty"`
	CanUnsubscribe *bool `json:"canUnsubscribe,omitempty"`
	IsPendingAutoprolong *bool `json:"isPendingAutoprolong,omitempty"`
} `json:"list"`
	HasMore bool `json:"hasMore"`
}

// GetAccessUsersBlockedResponse represents the response for GetAccessUsersBlocked
type GetAccessUsersBlockedResponse struct {
	List []struct {
	Id float64 `json:"id"`
	Username string `json:"username"`
	Name string `json:"name"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs *struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs,omitempty"`
	Lists []struct {
	Id interface{} `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
} `json:"lists"`
	DisplayName *string `json:"displayName,omitempty"`
	Notice *string `json:"notice,omitempty"`
	IsRestricted *bool `json:"isRestricted,omitempty"`
	CanRestrict *bool `json:"canRestrict,omitempty"`
	SubscribedBy *bool `json:"subscribedBy,omitempty"`
	SubscribedByExpire *bool `json:"subscribedByExpire,omitempty"`
	SubscribedByExpireDate *string `json:"subscribedByExpireDate,omitempty"`
	SubscribedByAutoprolong *bool `json:"subscribedByAutoprolong,omitempty"`
	SubscribedIsExpiredNow *bool `json:"subscribedIsExpiredNow,omitempty"`
	CurrentSubscribePrice *float64 `json:"currentSubscribePrice,omitempty"`
	SubscribedOn *bool `json:"subscribedOn,omitempty"`
	SubscribedOnExpiredNow *bool `json:"subscribedOnExpiredNow,omitempty"`
	SubscribedOnDuration *string `json:"subscribedOnDuration,omitempty"`
	CanReport *bool `json:"canReport,omitempty"`
	CanReceiveChatMessage *bool `json:"canReceiveChatMessage,omitempty"`
	HideChat *bool `json:"hideChat,omitempty"`
	LastSeen *string `json:"lastSeen,omitempty"`
	IsPerformer *bool `json:"isPerformer,omitempty"`
	IsRealPerformer *bool `json:"isRealPerformer,omitempty"`
	SubscribedByData *struct {
	Price float64 `json:"price"`
	NewPrice float64 `json:"newPrice"`
	RegularPrice float64 `json:"regularPrice"`
	SubscribePrice float64 `json:"subscribePrice"`
	DiscountPercent float64 `json:"discountPercent"`
	DiscountPeriod float64 `json:"discountPeriod"`
	SubscribeAt string `json:"subscribeAt"`
	ExpiredAt string `json:"expiredAt"`
	RenewedAt string `json:"renewedAt"`
	DiscountFinishedAt string `json:"discountFinishedAt"`
	DiscountStartedAt string `json:"discountStartedAt"`
	Status string `json:"status"`
	IsMuted bool `json:"isMuted"`
	UnsubscribeReason string `json:"unsubscribeReason"`
	Duration string `json:"duration"`
	HasActivePaidSubscriptions bool `json:"hasActivePaidSubscriptions"`
	ShowPostsInFeed bool `json:"showPostsInFeed"`
	Subscribes []struct {
	Id float64 `json:"id"`
	UserId float64 `json:"userId"`
	SubscriberId float64 `json:"subscriberId"`
	Date string `json:"date"`
	Duration float64 `json:"duration"`
	StartDate string `json:"startDate"`
	ExpireDate string `json:"expireDate"`
	CancelDate string `json:"cancelDate"`
	Price float64 `json:"price"`
	RegularPrice float64 `json:"regularPrice"`
	Discount float64 `json:"discount"`
	EarningId float64 `json:"earningId"`
	Action string `json:"action"`
	Type string `json:"type"`
	OfferStart string `json:"offerStart"`
	OfferEnd string `json:"offerEnd"`
	IsCurrent bool `json:"isCurrent"`
} `json:"subscribes"`
} `json:"subscribedByData,omitempty"`
	SubscribedOnData *struct {
	Price float64 `json:"price"`
	NewPrice float64 `json:"newPrice"`
	RegularPrice float64 `json:"regularPrice"`
	SubscribePrice float64 `json:"subscribePrice"`
	DiscountPercent float64 `json:"discountPercent"`
	DiscountPeriod float64 `json:"discountPeriod"`
	SubscribeAt string `json:"subscribeAt"`
	ExpiredAt string `json:"expiredAt"`
	RenewedAt string `json:"renewedAt"`
	DiscountFinishedAt string `json:"discountFinishedAt"`
	DiscountStartedAt string `json:"discountStartedAt"`
	Status string `json:"status"`
	IsMuted bool `json:"isMuted"`
	UnsubscribeReason string `json:"unsubscribeReason"`
	Duration string `json:"duration"`
	HasActivePaidSubscriptions bool `json:"hasActivePaidSubscriptions"`
	Subscribes []struct {
	Id float64 `json:"id"`
	UserId float64 `json:"userId"`
	SubscriberId float64 `json:"subscriberId"`
	Date string `json:"date"`
	Duration float64 `json:"duration"`
	StartDate string `json:"startDate"`
	ExpireDate string `json:"expireDate"`
	CancelDate string `json:"cancelDate"`
	Price float64 `json:"price"`
	RegularPrice float64 `json:"regularPrice"`
	Discount float64 `json:"discount"`
	EarningId float64 `json:"earningId"`
	Action string `json:"action"`
	Type string `json:"type"`
	OfferStart string `json:"offerStart"`
	OfferEnd string `json:"offerEnd"`
	IsCurrent bool `json:"isCurrent"`
} `json:"subscribes"`
	TipsSumm float64 `json:"tipsSumm"`
	SubscribesSumm float64 `json:"subscribesSumm"`
	MessagesSumm float64 `json:"messagesSumm"`
	PostsSumm float64 `json:"postsSumm"`
	StreamsSumm float64 `json:"streamsSumm"`
	TotalSumm float64 `json:"totalSumm"`
} `json:"subscribedOnData,omitempty"`
	CanTrialSend *bool `json:"canTrialSend,omitempty"`
	IsBlocked *bool `json:"isBlocked,omitempty"`
	CanUnsubscribe *bool `json:"canUnsubscribe,omitempty"`
	IsPendingAutoprolong *bool `json:"isPendingAutoprolong,omitempty"`
} `json:"list"`
	HasMore bool `json:"hasMore"`
}

// GetAccessUsersListResponse represents the response for GetAccessUsersList
type GetAccessUsersListResponse struct {
	Users []struct {
	Id float64 `json:"id"`
	Username string `json:"username"`
	Name string `json:"name"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs *struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs,omitempty"`
	Lists []struct {
	Id interface{} `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
} `json:"lists"`
	DisplayName *string `json:"displayName,omitempty"`
	Notice *string `json:"notice,omitempty"`
	About *string `json:"about,omitempty"`
	IsMarkdownDisabledForAbout *bool `json:"isMarkdownDisabledForAbout,omitempty"`
	Website *string `json:"website,omitempty"`
	Wishlist *string `json:"wishlist,omitempty"`
	Location *string `json:"location,omitempty"`
	Header *string `json:"header,omitempty"`
	HeaderSize *struct {
	Width float64 `json:"width"`
	Height float64 `json:"height"`
} `json:"headerSize,omitempty"`
	HeaderThumbs *struct {
	W480 string `json:"w480"`
	W760 string `json:"w760"`
} `json:"headerThumbs,omitempty"`
	SubscribersCount *float64 `json:"subscribersCount,omitempty"`
	PostsCount *float64 `json:"postsCount,omitempty"`
	ArchivedPostsCount *float64 `json:"archivedPostsCount,omitempty"`
	PrivateArchivedPostsCount *float64 `json:"privateArchivedPostsCount,omitempty"`
	PhotosCount *float64 `json:"photosCount,omitempty"`
	VideosCount *float64 `json:"videosCount,omitempty"`
	AudiosCount *float64 `json:"audiosCount,omitempty"`
	MediasCount *float64 `json:"mediasCount,omitempty"`
	FavoritesCount *float64 `json:"favoritesCount,omitempty"`
	FavoritedCount *float64 `json:"favoritedCount,omitempty"`
	JoinDate *string `json:"joinDate,omitempty"`
	LastSeen *string `json:"lastSeen,omitempty"`
	SubscribedBy *bool `json:"subscribedBy,omitempty"`
	SubscribedByExpire *bool `json:"subscribedByExpire,omitempty"`
	SubscribedByExpireDate *string `json:"subscribedByExpireDate,omitempty"`
	SubscribedByAutoprolong *bool `json:"subscribedByAutoprolong,omitempty"`
	SubscribedIsExpiredNow *bool `json:"subscribedIsExpiredNow,omitempty"`
	SubscribedByData *struct {
	Price float64 `json:"price"`
	NewPrice float64 `json:"newPrice"`
	RegularPrice float64 `json:"regularPrice"`
	SubscribePrice float64 `json:"subscribePrice"`
	DiscountPercent float64 `json:"discountPercent"`
	DiscountPeriod float64 `json:"discountPeriod"`
	SubscribeAt string `json:"subscribeAt"`
	ExpiredAt string `json:"expiredAt"`
	RenewedAt string `json:"renewedAt"`
	DiscountFinishedAt string `json:"discountFinishedAt"`
	DiscountStartedAt string `json:"discountStartedAt"`
	Status string `json:"status"`
	IsMuted bool `json:"isMuted"`
	UnsubscribeReason string `json:"unsubscribeReason"`
	Duration string `json:"duration"`
	HasActivePaidSubscriptions bool `json:"hasActivePaidSubscriptions"`
	ShowPostsInFeed bool `json:"showPostsInFeed"`
	Subscribes []struct {
	Id float64 `json:"id"`
	UserId float64 `json:"userId"`
	SubscriberId float64 `json:"subscriberId"`
	Date string `json:"date"`
	Duration float64 `json:"duration"`
	StartDate string `json:"startDate"`
	ExpireDate string `json:"expireDate"`
	CancelDate string `json:"cancelDate"`
	Price float64 `json:"price"`
	RegularPrice float64 `json:"regularPrice"`
	Discount float64 `json:"discount"`
	EarningId float64 `json:"earningId"`
	Action string `json:"action"`
	Type string `json:"type"`
	OfferStart string `json:"offerStart"`
	OfferEnd string `json:"offerEnd"`
	IsCurrent bool `json:"isCurrent"`
} `json:"subscribes"`
} `json:"subscribedByData,omitempty"`
	SubscribedOn *bool `json:"subscribedOn,omitempty"`
	SubscribedOnExpiredNow *bool `json:"subscribedOnExpiredNow,omitempty"`
	SubscribedOnDuration *string `json:"subscribedOnDuration,omitempty"`
	SubscribedOnData *struct {
	Price float64 `json:"price"`
	NewPrice float64 `json:"newPrice"`
	RegularPrice float64 `json:"regularPrice"`
	SubscribePrice float64 `json:"subscribePrice"`
	DiscountPercent float64 `json:"discountPercent"`
	DiscountPeriod float64 `json:"discountPeriod"`
	SubscribeAt string `json:"subscribeAt"`
	ExpiredAt string `json:"expiredAt"`
	RenewedAt string `json:"renewedAt"`
	DiscountFinishedAt string `json:"discountFinishedAt"`
	DiscountStartedAt string `json:"discountStartedAt"`
	Status string `json:"status"`
	IsMuted bool `json:"isMuted"`
	UnsubscribeReason string `json:"unsubscribeReason"`
	Duration string `json:"duration"`
	HasActivePaidSubscriptions bool `json:"hasActivePaidSubscriptions"`
	Subscribes []struct {
	Id float64 `json:"id"`
	UserId float64 `json:"userId"`
	SubscriberId float64 `json:"subscriberId"`
	Date string `json:"date"`
	Duration float64 `json:"duration"`
	StartDate string `json:"startDate"`
	ExpireDate string `json:"expireDate"`
	CancelDate string `json:"cancelDate"`
	Price float64 `json:"price"`
	RegularPrice float64 `json:"regularPrice"`
	Discount float64 `json:"discount"`
	EarningId float64 `json:"earningId"`
	Action string `json:"action"`
	Type string `json:"type"`
	OfferStart string `json:"offerStart"`
	OfferEnd string `json:"offerEnd"`
	IsCurrent bool `json:"isCurrent"`
} `json:"subscribes"`
	TipsSumm float64 `json:"tipsSumm"`
	SubscribesSumm float64 `json:"subscribesSumm"`
	MessagesSumm float64 `json:"messagesSumm"`
	PostsSumm float64 `json:"postsSumm"`
	StreamsSumm float64 `json:"streamsSumm"`
	TotalSumm float64 `json:"totalSumm"`
} `json:"subscribedOnData,omitempty"`
	SubscribePrice *float64 `json:"subscribePrice,omitempty"`
	CurrentSubscribePrice *float64 `json:"currentSubscribePrice,omitempty"`
	CanAddSubscriber *bool `json:"canAddSubscriber,omitempty"`
	TipsEnabled *bool `json:"tipsEnabled,omitempty"`
	TipsTextEnabled *bool `json:"tipsTextEnabled,omitempty"`
	TipsMin *float64 `json:"tipsMin,omitempty"`
	TipsMinInternal *float64 `json:"tipsMinInternal,omitempty"`
	TipsMax *float64 `json:"tipsMax,omitempty"`
	CanLookStory *bool `json:"canLookStory,omitempty"`
	CanCommentStory *bool `json:"canCommentStory,omitempty"`
	HasNotViewedStory *bool `json:"hasNotViewedStory,omitempty"`
	HasStories *bool `json:"hasStories,omitempty"`
	IsRestricted *bool `json:"isRestricted,omitempty"`
	CanRestrict *bool `json:"canRestrict,omitempty"`
	IsBlocked *bool `json:"isBlocked,omitempty"`
	CanReport *bool `json:"canReport,omitempty"`
	CanUnsubscribe *bool `json:"canUnsubscribe,omitempty"`
	IsPendingAutoprolong *bool `json:"isPendingAutoprolong,omitempty"`
	IsPerformer *bool `json:"isPerformer,omitempty"`
	IsRealPerformer *bool `json:"isRealPerformer,omitempty"`
	CanReceiveChatMessage *bool `json:"canReceiveChatMessage,omitempty"`
	CanChat *bool `json:"canChat,omitempty"`
	ShowPostsInFeed *bool `json:"showPostsInFeed,omitempty"`
	HasPinnedPosts *bool `json:"hasPinnedPosts,omitempty"`
	HasLabels *bool `json:"hasLabels,omitempty"`
	IsPrivateRestriction *bool `json:"isPrivateRestriction,omitempty"`
	ShowSubscribersCount *bool `json:"showSubscribersCount,omitempty"`
	ShowMediaCount *bool `json:"showMediaCount,omitempty"`
	IsReferrerAllowed *bool `json:"isReferrerAllowed,omitempty"`
	CanCreatePromotion *bool `json:"canCreatePromotion,omitempty"`
	CanCreateTrial *bool `json:"canCreateTrial,omitempty"`
	IsAdultContent *bool `json:"isAdultContent,omitempty"`
	CanTrialSend *bool `json:"canTrialSend,omitempty"`
	IsFriend *bool `json:"isFriend,omitempty"`
	HasScheduledStream *bool `json:"hasScheduledStream,omitempty"`
	HasStream *bool `json:"hasStream,omitempty"`
	CanPayInternal *bool `json:"canPayInternal,omitempty"`
} `json:"users"`
}

// ListAccessUsersListsResponse represents the response for ListAccessUsersLists
type ListAccessUsersListsResponse struct {
	List []struct {
	Id interface{} `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
	UsersCount *float64 `json:"usersCount,omitempty"`
	PostsCount *float64 `json:"postsCount,omitempty"`
	Order *string `json:"order,omitempty"`
	Direction *string `json:"direction,omitempty"`
	CanUpdate *bool `json:"canUpdate,omitempty"`
	CanDelete *bool `json:"canDelete,omitempty"`
	CanManageUsers *bool `json:"canManageUsers,omitempty"`
	CanAddUsers *bool `json:"canAddUsers,omitempty"`
	CanPinnedToFeed *bool `json:"canPinnedToFeed,omitempty"`
	IsPinnedToFeed *bool `json:"isPinnedToFeed,omitempty"`
	CanPinnedToChat *bool `json:"canPinnedToChat,omitempty"`
	IsPinnedToChat *bool `json:"isPinnedToChat,omitempty"`
	SortList []struct {
	Order string `json:"order"`
	Direction string `json:"direction"`
} `json:"sortList"`
	Users []struct {
	Id float64 `json:"id"`
	Username string `json:"username"`
	Name string `json:"name"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs *struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs,omitempty"`
} `json:"users"`
} `json:"list"`
	HasMore bool `json:"hasMore"`
	Order *string `json:"order,omitempty"`
	Sort *string `json:"sort,omitempty"`
}

// CreateAccessUsersListsResponse represents the response for CreateAccessUsersLists
type CreateAccessUsersListsResponse struct {
	Id interface{} `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
	UsersCount *float64 `json:"usersCount,omitempty"`
	PostsCount *float64 `json:"postsCount,omitempty"`
	Order *string `json:"order,omitempty"`
	Direction *string `json:"direction,omitempty"`
	CanUpdate *bool `json:"canUpdate,omitempty"`
	CanDelete *bool `json:"canDelete,omitempty"`
	CanManageUsers *bool `json:"canManageUsers,omitempty"`
	CanAddUsers *bool `json:"canAddUsers,omitempty"`
	CanPinnedToFeed *bool `json:"canPinnedToFeed,omitempty"`
	IsPinnedToFeed *bool `json:"isPinnedToFeed,omitempty"`
	CanPinnedToChat *bool `json:"canPinnedToChat,omitempty"`
	IsPinnedToChat *bool `json:"isPinnedToChat,omitempty"`
	SortList []struct {
	Order string `json:"order"`
	Direction string `json:"direction"`
} `json:"sortList"`
	Users []struct {
	Id float64 `json:"id"`
	Username string `json:"username"`
	Name string `json:"name"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs *struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs,omitempty"`
} `json:"users"`
}

// GetAccessUsersListsResponse represents the response for GetAccessUsersLists
type GetAccessUsersListsResponse struct {
	Id interface{} `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
	UsersCount *float64 `json:"usersCount,omitempty"`
	PostsCount *float64 `json:"postsCount,omitempty"`
	Order *string `json:"order,omitempty"`
	Direction *string `json:"direction,omitempty"`
	CanUpdate *bool `json:"canUpdate,omitempty"`
	CanDelete *bool `json:"canDelete,omitempty"`
	CanManageUsers *bool `json:"canManageUsers,omitempty"`
	CanAddUsers *bool `json:"canAddUsers,omitempty"`
	CanPinnedToFeed *bool `json:"canPinnedToFeed,omitempty"`
	IsPinnedToFeed *bool `json:"isPinnedToFeed,omitempty"`
	CanPinnedToChat *bool `json:"canPinnedToChat,omitempty"`
	IsPinnedToChat *bool `json:"isPinnedToChat,omitempty"`
	SortList []struct {
	Order string `json:"order"`
	Direction string `json:"direction"`
} `json:"sortList"`
	Users []struct {
	Id float64 `json:"id"`
	Username string `json:"username"`
	Name string `json:"name"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs *struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs,omitempty"`
} `json:"users"`
}

// UpdateAccessUsersListsResponse represents the response for UpdateAccessUsersLists
type UpdateAccessUsersListsResponse struct {
	Id interface{} `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
	UsersCount *float64 `json:"usersCount,omitempty"`
	PostsCount *float64 `json:"postsCount,omitempty"`
	Order *string `json:"order,omitempty"`
	Direction *string `json:"direction,omitempty"`
	CanUpdate *bool `json:"canUpdate,omitempty"`
	CanDelete *bool `json:"canDelete,omitempty"`
	CanManageUsers *bool `json:"canManageUsers,omitempty"`
	CanAddUsers *bool `json:"canAddUsers,omitempty"`
	CanPinnedToFeed *bool `json:"canPinnedToFeed,omitempty"`
	IsPinnedToFeed *bool `json:"isPinnedToFeed,omitempty"`
	CanPinnedToChat *bool `json:"canPinnedToChat,omitempty"`
	IsPinnedToChat *bool `json:"isPinnedToChat,omitempty"`
	SortList []struct {
	Order string `json:"order"`
	Direction string `json:"direction"`
} `json:"sortList"`
	Users []struct {
	Id float64 `json:"id"`
	Username string `json:"username"`
	Name string `json:"name"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs *struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs,omitempty"`
} `json:"users"`
}

// DeleteAccessUsersListsResponse represents the response for DeleteAccessUsersLists
type DeleteAccessUsersListsResponse map[string]interface{}

// ListAccessUsersListsUsersResponse represents the response for ListAccessUsersListsUsers
type ListAccessUsersListsUsersResponse struct {
	List []struct {
	Id float64 `json:"id"`
	Username string `json:"username"`
	Name string `json:"name"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs *struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs,omitempty"`
} `json:"list"`
	HasMore bool `json:"hasMore"`
	NextOffset *float64 `json:"nextOffset,omitempty"`
}

// CreateAccessUsersLists1Response represents the response for CreateAccessUsersLists1
type CreateAccessUsersLists1Response struct {
	Success []struct {
	ListId float64 `json:"listId"`
	Success bool `json:"success"`
} `json:"success"`
	Errors []struct {
	ListId float64 `json:"listId"`
	Error string `json:"error"`
} `json:"errors,omitempty"`
}

// CreateAccessUsersRestrictResponse represents the response for CreateAccessUsersRestrict
type CreateAccessUsersRestrictResponse map[string]interface{}

// DeleteAccessUsersRestrictResponse represents the response for DeleteAccessUsersRestrict
type DeleteAccessUsersRestrictResponse map[string]interface{}

// ListAccessVaultMediaResponse represents the response for ListAccessVaultMedia
type ListAccessVaultMediaResponse struct {
	List []struct {
	Id float64 `json:"id"`
	Type string `json:"type"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CanView bool `json:"canView"`
	HasError bool `json:"hasError"`
	CreatedAt string `json:"createdAt"`
	IsReady bool `json:"isReady"`
	Duration *float64 `json:"duration,omitempty"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	View string `json:"view"`
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs"`
	IvStatus string `json:"ivStatus"`
	IsFromGuest bool `json:"isFromGuest"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
	Files *struct {
	Full struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"sources"`
} `json:"full"`
	Thumb *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"thumb,omitempty"`
	Preview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Options []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"options"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"squarePreview,omitempty"`
} `json:"files,omitempty"`
	Counters struct {
	LikesCount float64 `json:"likesCount"`
	TipsSumm float64 `json:"tipsSumm"`
} `json:"counters"`
	HasPosts *bool `json:"hasPosts,omitempty"`
	Lists []struct {
	Id float64 `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
} `json:"lists"`
} `json:"list"`
	HasMore bool `json:"hasMore"`
}

// ListAccessVaultListsResponse represents the response for ListAccessVaultLists
type ListAccessVaultListsResponse struct {
	List []struct {
	Id float64 `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
	HasMedia bool `json:"hasMedia"`
	CanUpdate bool `json:"canUpdate"`
	CanDelete bool `json:"canDelete"`
	Medias []struct {
	Type string `json:"type"`
	Url *string `json:"url,omitempty"`
} `json:"medias"`
} `json:"list"`
	All struct {
	VideosCount float64 `json:"videosCount"`
	PhotosCount float64 `json:"photosCount"`
	GifsCount float64 `json:"gifsCount"`
	AudiosCount float64 `json:"audiosCount"`
	Medias []struct {
	Type interface{} `json:"type"`
} `json:"medias"`
} `json:"all"`
	HasMore bool `json:"hasMore"`
	CanCreateVaultLists bool `json:"canCreateVaultLists"`
	Order string `json:"order"`
	Sort string `json:"sort"`
}

// CreateAccessVaultListsResponse represents the response for CreateAccessVaultLists
type CreateAccessVaultListsResponse struct {
	Id float64 `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
	HasMedia bool `json:"hasMedia"`
	CanUpdate bool `json:"canUpdate"`
	CanDelete bool `json:"canDelete"`
	Medias []struct {
	Type string `json:"type"`
	Url *string `json:"url,omitempty"`
} `json:"medias"`
}

// UpdateAccessVaultListsResponse represents the response for UpdateAccessVaultLists
type UpdateAccessVaultListsResponse struct {
	Id float64 `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
	HasMedia bool `json:"hasMedia"`
	CanUpdate bool `json:"canUpdate"`
	CanDelete bool `json:"canDelete"`
	Medias []struct {
	Type string `json:"type"`
	Url *string `json:"url,omitempty"`
} `json:"medias"`
}

// DeleteAccessVaultListsResponse represents the response for DeleteAccessVaultLists
type DeleteAccessVaultListsResponse map[string]interface{}

// ListAccessVaultListsMediaResponse represents the response for ListAccessVaultListsMedia
type ListAccessVaultListsMediaResponse struct {
	List []struct {
	Id float64 `json:"id"`
	Type string `json:"type"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CanView bool `json:"canView"`
	HasError bool `json:"hasError"`
	CreatedAt string `json:"createdAt"`
	IsReady bool `json:"isReady"`
	Duration *float64 `json:"duration,omitempty"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	View string `json:"view"`
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs"`
	IvStatus string `json:"ivStatus"`
	IsFromGuest bool `json:"isFromGuest"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
	Files *struct {
	Full struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"sources"`
} `json:"full"`
	Thumb *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"thumb,omitempty"`
	Preview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Options []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"options"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"squarePreview,omitempty"`
} `json:"files,omitempty"`
	Counters struct {
	LikesCount float64 `json:"likesCount"`
	TipsSumm float64 `json:"tipsSumm"`
} `json:"counters"`
	HasPosts *bool `json:"hasPosts,omitempty"`
	Lists []struct {
	Id float64 `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
} `json:"lists"`
} `json:"list"`
	HasMore bool `json:"hasMore"`
}

// CreateAccessVaultListsMediaResponse represents the response for CreateAccessVaultListsMedia
type CreateAccessVaultListsMediaResponse struct {
	Id float64 `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
	HasMedia bool `json:"hasMedia"`
	CanUpdate bool `json:"canUpdate"`
	CanDelete bool `json:"canDelete"`
	Medias []struct {
	Type string `json:"type"`
	Url *string `json:"url,omitempty"`
} `json:"medias"`
}

// AccessUploadsCheckResponse represents the response for AccessUploadsCheck
type AccessUploadsCheckResponse struct {
	Exists bool `json:"exists"`
	Media *struct {
	Id float64 `json:"id"`
	Type string `json:"type"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CanView bool `json:"canView"`
	HasError bool `json:"hasError"`
	CreatedAt string `json:"createdAt"`
	IsReady bool `json:"isReady"`
	Duration *float64 `json:"duration,omitempty"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	View string `json:"view"`
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs"`
	IvStatus string `json:"ivStatus"`
	IsFromGuest bool `json:"isFromGuest"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
	Files *struct {
	Full struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"sources"`
} `json:"full"`
	Thumb *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"thumb,omitempty"`
	Preview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Options []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"options"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"squarePreview,omitempty"`
} `json:"files,omitempty"`
} `json:"media,omitempty"`
}

// AccessUploadsInitResponse represents the response for AccessUploadsInit
type AccessUploadsInitResponse struct {
	MediaUploadId string `json:"mediaUploadId"`
}

// ReplaceAccessUploadsPartsResponse represents the response for ReplaceAccessUploadsParts
type ReplaceAccessUploadsPartsResponse struct {
	MediaUploadId string `json:"mediaUploadId"`
	PartNumber int64 `json:"partNumber"`
	Etag string `json:"etag"`
}

// ReplaceAccessUploadsResponse represents the response for ReplaceAccessUploads
type ReplaceAccessUploadsResponse struct {
	MediaUploadId string `json:"mediaUploadId"`
	Media *struct {
	Id float64 `json:"id"`
	Type string `json:"type"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CanView bool `json:"canView"`
	HasError bool `json:"hasError"`
	CreatedAt string `json:"createdAt"`
	IsReady bool `json:"isReady"`
	Duration *float64 `json:"duration,omitempty"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	View string `json:"view"`
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs"`
	IvStatus string `json:"ivStatus"`
	IsFromGuest bool `json:"isFromGuest"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
	Files *struct {
	Full struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"sources"`
} `json:"full"`
	Thumb *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"thumb,omitempty"`
	Preview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Options []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"options"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"squarePreview,omitempty"`
} `json:"files,omitempty"`
} `json:"media,omitempty"`
}

// AccessUploadsCompleteResponse represents the response for AccessUploadsComplete
type AccessUploadsCompleteResponse struct {
	MediaUploadId string `json:"mediaUploadId"`
	Media *struct {
	Id float64 `json:"id"`
	Type string `json:"type"`
	ConvertedToVideo bool `json:"convertedToVideo"`
	CanView bool `json:"canView"`
	HasError bool `json:"hasError"`
	CreatedAt string `json:"createdAt"`
	IsReady bool `json:"isReady"`
	Duration *float64 `json:"duration,omitempty"`
	ReleaseForms []struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	PartnerSource string `json:"partnerSource"`
	Type string `json:"type"`
	User *struct {
	View string `json:"view"`
	Id float64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	IsVerified bool `json:"isVerified"`
	Avatar string `json:"avatar"`
	AvatarThumbs struct {
	C50 string `json:"c50"`
	C144 string `json:"c144"`
} `json:"avatarThumbs"`
	IvStatus string `json:"ivStatus"`
	IsFromGuest bool `json:"isFromGuest"`
} `json:"user,omitempty"`
} `json:"releaseForms"`
	HasCustomPreview *bool `json:"hasCustomPreview,omitempty"`
	VideoSources *struct {
	N240 *string `json:"240,omitempty"`
	N720 *string `json:"720,omitempty"`
} `json:"videoSources,omitempty"`
	Files *struct {
	Full struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Sources []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"sources"`
} `json:"full"`
	Thumb *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"thumb,omitempty"`
	Preview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Options []struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
} `json:"options"`
} `json:"preview,omitempty"`
	SquarePreview *struct {
	Url string `json:"url"`
	Width *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Size *float64 `json:"size,omitempty"`
} `json:"squarePreview,omitempty"`
} `json:"files,omitempty"`
} `json:"media,omitempty"`
}

// LinkInitResponse represents the response for LinkInit
type LinkInitResponse struct {
	Url string `json:"url"`
	ExpiresAt string `json:"expiresAt"`
}

// GetLinkResponse represents the response for GetLink
type GetLinkResponse struct {
	Status string `json:"status"`
	Data *interface{} `json:"data,omitempty"`
}

// ListDynamicRulesResponse represents the response for ListDynamicRules
type ListDynamicRulesResponse struct {
	Rules struct {
	StaticParam string `json:"static_param"`
	Format string `json:"format"`
	Start string `json:"start"`
	End string `json:"end"`
	Prefix string `json:"prefix"`
	Suffix string `json:"suffix"`
	ChecksumConstant float64 `json:"checksum_constant"`
	ChecksumIndexes []float64 `json:"checksum_indexes"`
	AppToken string `json:"app_token"`
	Revision string `json:"revision"`
} `json:"rules"`
	IsCurrent bool `json:"is_current"`
	IsPublic bool `json:"is_public"`
	IsEarlyAccess bool `json:"is_early_access"`
}

// DynamicRulesSignResponse represents the response for DynamicRulesSign
type DynamicRulesSignResponse struct {
	Signed struct {
	Sign string `json:"sign"`
	Time string `json:"time"`
	UserId *string `json:"user-id,omitempty"`
	AppToken string `json:"app-token"`
} `json:"signed"`
	IsPublic bool `json:"is_public"`
	IsEarlyAccess bool `json:"is_early_access"`
}

// GetDynamicRulesStatusResponse represents the response for GetDynamicRulesStatus
type GetDynamicRulesStatusResponse struct {
	Revision string `json:"revision"`
	EarlyAccessRevision string `json:"early_access_revision"`
	PublicRevision string `json:"public_revision"`
	IsCurrent bool `json:"is_current"`
	IsEarlyAccess bool `json:"is_early_access"`
	IsPublic bool `json:"is_public"`
	AccessGranted bool `json:"access_granted"`
}

// CreateVaultPlusStoreListResponse represents the response for CreateVaultPlusStoreList
type CreateVaultPlusStoreListResponse struct {
	Queued bool `json:"queued"`
	ListId string `json:"listId"`
	EstimatedItems *float64 `json:"estimatedItems,omitempty"`
}

// GetVaultPlusStoreStatusResponse represents the response for GetVaultPlusStoreStatus
type GetVaultPlusStoreStatusResponse struct {
	ConnectionId string `json:"connectionId"`
	TotalMedia float64 `json:"totalMedia"`
	StoredCount float64 `json:"storedCount"`
	PendingCount float64 `json:"pendingCount"`
	TotalSizeBytes float64 `json:"totalSizeBytes"`
}

// GetVaultPlusStoreStatsResponse represents the response for GetVaultPlusStoreStats
type GetVaultPlusStoreStatsResponse struct {
	TotalConnections float64 `json:"totalConnections"`
	ConnectionsWithVaultPlus float64 `json:"connectionsWithVaultPlus"`
	TotalStorageBytes float64 `json:"totalStorageBytes"`
	TotalMediaCount float64 `json:"totalMediaCount"`
}

// GetVaultPlusResponse represents the response for GetVaultPlus
type GetVaultPlusResponse struct {
	Id string `json:"id"`
	Type string `json:"type"`
	Duration float64 `json:"duration"`
	Media map[string]struct {
	Id string `json:"id"`
	Status string `json:"status"`
	Quality string `json:"quality"`
	SizeBytes float64 `json:"sizeBytes"`
	ContentType string `json:"contentType"`
	Source string `json:"source"`
	AccessCount float64 `json:"accessCount"`
	CreatedAt float64 `json:"createdAt"`
	ExpiresAt float64 `json:"expiresAt"`
	StoredAt float64 `json:"storedAt"`
	LastAccessedAt float64 `json:"lastAccessedAt"`
	Url string `json:"url"`
} `json:"media"`
}

// DeleteVaultPlusResponse represents the response for DeleteVaultPlus
type DeleteVaultPlusResponse struct {
	Success bool `json:"success"`
	MediaId string `json:"mediaId"`
	FreedBytes float64 `json:"freedBytes"`
}

// CreateVaultPlusBatchResponse represents the response for CreateVaultPlusBatch
type CreateVaultPlusBatchResponse struct {
	Items []struct {
	Id string `json:"id"`
	Type string `json:"type"`
	Duration float64 `json:"duration"`
	Media map[string]struct {
	Id string `json:"id"`
	Status string `json:"status"`
	Quality string `json:"quality"`
	SizeBytes float64 `json:"sizeBytes"`
	ContentType string `json:"contentType"`
	Source string `json:"source"`
	AccessCount float64 `json:"accessCount"`
	CreatedAt float64 `json:"createdAt"`
	ExpiresAt float64 `json:"expiresAt"`
	StoredAt float64 `json:"storedAt"`
	LastAccessedAt float64 `json:"lastAccessedAt"`
	Url string `json:"url"`
} `json:"media"`
} `json:"items"`
}

// GetVaultPlusListResponse represents the response for GetVaultPlusList
type GetVaultPlusListResponse struct {
	Items []struct {
	Id string `json:"id"`
	Type string `json:"type"`
	Duration float64 `json:"duration"`
	Media map[string]struct {
	Id string `json:"id"`
	Status string `json:"status"`
	Quality string `json:"quality"`
	SizeBytes float64 `json:"sizeBytes"`
	ContentType string `json:"contentType"`
	Source string `json:"source"`
	AccessCount float64 `json:"accessCount"`
	CreatedAt float64 `json:"createdAt"`
	ExpiresAt float64 `json:"expiresAt"`
	StoredAt float64 `json:"storedAt"`
	LastAccessedAt float64 `json:"lastAccessedAt"`
	Url string `json:"url"`
} `json:"media"`
} `json:"items"`
	NextCursor *string `json:"nextCursor,omitempty"`
}

// VaultPlusPurgeResponse represents the response for VaultPlusPurge
type VaultPlusPurgeResponse struct {
	Success bool `json:"success"`
	PurgedCount float64 `json:"purgedCount"`
	FreedBytes float64 `json:"freedBytes"`
}


// ============================================================================
// API Methods
// ============================================================================

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
	AutoCacheVault *bool `json:"autoCacheVault,omitempty"`
	AutoCacheMessages *bool `json:"autoCacheMessages,omitempty"`
	AutoCachePosts *bool `json:"autoCachePosts,omitempty"`
	AutoCacheStories *bool `json:"autoCacheStories,omitempty"`
	MinAccessCountVault *float64 `json:"minAccessCountVault,omitempty"`
	MinAccessCountMessages *float64 `json:"minAccessCountMessages,omitempty"`
	MinAccessCountPosts *float64 `json:"minAccessCountPosts,omitempty"`
	MinAccessCountStories *float64 `json:"minAccessCountStories,omitempty"`
	CacheImages *bool `json:"cacheImages,omitempty"`
	CacheVideos *bool `json:"cacheVideos,omitempty"`
	CacheAudio *bool `json:"cacheAudio,omitempty"`
	ImageQualities []string `json:"imageQualities,omitempty"`
	VideoQualities []string `json:"videoQualities,omitempty"`
	RetentionDays *float64 `json:"retentionDays,omitempty"`
	AccessExpiryDays *float64 `json:"accessExpiryDays,omitempty"`
	PresignedUrlTtlSeconds *float64 `json:"presignedUrlTtlSeconds,omitempty"`
	StorageLimitBytes *float64 `json:"storageLimitBytes,omitempty"`
	StorageLimitPurgeStrategy *string `json:"storageLimitPurgeStrategy,omitempty"`
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

// CreateAccountConnectionsImport Import connection
func (c *Client) CreateAccountConnectionsImport(ctx context.Context, body struct {
	Cookie string `json:"cookie"`
	UserAgent string `json:"userAgent"`
	Permissions []string `json:"permissions,omitempty"`
	ClientReferenceId *string `json:"clientReferenceId,omitempty"`
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
	AutoEnableForNewConnections *bool `json:"autoEnableForNewConnections,omitempty"`
	ApplyToExistingConnections *bool `json:"applyToExistingConnections,omitempty"`
	DefaultSettings *struct {
	AutoCacheVault *bool `json:"autoCacheVault,omitempty"`
	AutoCacheMessages *bool `json:"autoCacheMessages,omitempty"`
	AutoCachePosts *bool `json:"autoCachePosts,omitempty"`
	AutoCacheStories *bool `json:"autoCacheStories,omitempty"`
	MinAccessCountVault *float64 `json:"minAccessCountVault,omitempty"`
	MinAccessCountMessages *float64 `json:"minAccessCountMessages,omitempty"`
	MinAccessCountPosts *float64 `json:"minAccessCountPosts,omitempty"`
	MinAccessCountStories *float64 `json:"minAccessCountStories,omitempty"`
	CacheImages *bool `json:"cacheImages,omitempty"`
	CacheVideos *bool `json:"cacheVideos,omitempty"`
	CacheAudio *bool `json:"cacheAudio,omitempty"`
	ImageQualities []string `json:"imageQualities,omitempty"`
	VideoQualities []string `json:"videoQualities,omitempty"`
	RetentionDays *float64 `json:"retentionDays,omitempty"`
	AccessExpiryDays *float64 `json:"accessExpiryDays,omitempty"`
	PresignedUrlTtlSeconds *float64 `json:"presignedUrlTtlSeconds,omitempty"`
	StorageLimitBytes *float64 `json:"storageLimitBytes,omitempty"`
	StorageLimitPurgeStrategy *string `json:"storageLimitPurgeStrategy,omitempty"`
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
	Name *string `json:"name,omitempty"`
	About *string `json:"about,omitempty"`
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

// GetAccessAnalyticsMassMessagesSent Sent mass messages
func (c *Client) GetAccessAnalyticsMassMessagesSent(ctx context.Context, opts *RequestOptions) (*GetAccessAnalyticsMassMessagesSentResponse, error) {
	path := "/v2/access/analytics/mass-messages/sent"
	if opts == nil {
		opts = &RequestOptions{}
	}

	respBody, err := c.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}

	var result GetAccessAnalyticsMassMessagesSentResponse
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
	Text *string `json:"text,omitempty"`
	MediaItems []interface{} `json:"mediaItems,omitempty"`
	IsLockedText *bool `json:"isLockedText,omitempty"`
	Price *float64 `json:"price,omitempty"`
	PreviewMediaCount *int64 `json:"previewMediaCount,omitempty"`
	ReleaseForms *struct {
	Users []int64 `json:"users,omitempty"`
	Partners []int64 `json:"partners,omitempty"`
	Guests []int64 `json:"guests,omitempty"`
} `json:"releaseForms,omitempty"`
	UserTags []int64 `json:"userTags,omitempty"`
	IsMarkdown *bool `json:"isMarkdown,omitempty"`
	ScheduledDate *string `json:"scheduledDate,omitempty"`
	FundRaisingTargetAmount *float64 `json:"fundRaisingTargetAmount,omitempty"`
	FundRaisingTipsPresets []float64 `json:"fundRaisingTipsPresets,omitempty"`
	ExpireAfter *float64 `json:"expireAfter,omitempty"`
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
	Text *string `json:"text,omitempty"`
	MediaItems []interface{} `json:"mediaItems,omitempty"`
	IsLockedText *bool `json:"isLockedText,omitempty"`
	Price *float64 `json:"price,omitempty"`
	PreviewMediaCount *int64 `json:"previewMediaCount,omitempty"`
	ReleaseForms *struct {
	Users []int64 `json:"users,omitempty"`
	Partners []int64 `json:"partners,omitempty"`
	Guests []int64 `json:"guests,omitempty"`
} `json:"releaseForms,omitempty"`
	UserTags []int64 `json:"userTags,omitempty"`
	IsMarkdown *bool `json:"isMarkdown,omitempty"`
	ScheduledDate *string `json:"scheduledDate,omitempty"`
	FundRaisingTargetAmount *float64 `json:"fundRaisingTargetAmount,omitempty"`
	FundRaisingTipsPresets []float64 `json:"fundRaisingTipsPresets,omitempty"`
	ExpireAfter *float64 `json:"expireAfter,omitempty"`
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
	Text *string `json:"text,omitempty"`
	MediaItems []interface{} `json:"mediaItems,omitempty"`
	IsLockedText *bool `json:"isLockedText,omitempty"`
	Price *float64 `json:"price,omitempty"`
	PreviewMediaCount *int64 `json:"previewMediaCount,omitempty"`
	ReleaseForms *struct {
	Users []int64 `json:"users,omitempty"`
	Partners []int64 `json:"partners,omitempty"`
	Guests []int64 `json:"guests,omitempty"`
} `json:"releaseForms,omitempty"`
	UserTags []int64 `json:"userTags,omitempty"`
	IsMarkdown *bool `json:"isMarkdown,omitempty"`
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

// CreateAccessMassMessages Create mass message
func (c *Client) CreateAccessMassMessages(ctx context.Context, body struct {
	IsForwardedMessage *bool `json:"isForwardedMessage,omitempty"`
	Text *string `json:"text,omitempty"`
	MediaItems []interface{} `json:"mediaItems,omitempty"`
	IsLockedText *bool `json:"isLockedText,omitempty"`
	Price *float64 `json:"price,omitempty"`
	PreviewMediaCount *int64 `json:"previewMediaCount,omitempty"`
	ReleaseForms *struct {
	Users []int64 `json:"users,omitempty"`
	Partners []int64 `json:"partners,omitempty"`
	Guests []int64 `json:"guests,omitempty"`
} `json:"releaseForms,omitempty"`
	UserTags []int64 `json:"userTags,omitempty"`
	IsMarkdown *bool `json:"isMarkdown,omitempty"`
	ScheduledDate *interface{} `json:"scheduledDate,omitempty"`
	UserIds []interface{} `json:"userIds,omitempty"`
	UserLists []interface{} `json:"userLists,omitempty"`
	SubscribedAfterDate *interface{} `json:"subscribedAfterDate,omitempty"`
	ExcludeUserLists []interface{} `json:"excludeUserLists,omitempty"`
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
	IsForwardedMessage *bool `json:"isForwardedMessage,omitempty"`
	Text *string `json:"text,omitempty"`
	MediaItems []interface{} `json:"mediaItems,omitempty"`
	IsLockedText *bool `json:"isLockedText,omitempty"`
	Price *float64 `json:"price,omitempty"`
	PreviewMediaCount *int64 `json:"previewMediaCount,omitempty"`
	ReleaseForms *struct {
	Users []int64 `json:"users,omitempty"`
	Partners []int64 `json:"partners,omitempty"`
	Guests []int64 `json:"guests,omitempty"`
} `json:"releaseForms,omitempty"`
	UserTags []int64 `json:"userTags,omitempty"`
	IsMarkdown *bool `json:"isMarkdown,omitempty"`
	ScheduledDate *interface{} `json:"scheduledDate,omitempty"`
	UserIds []interface{} `json:"userIds,omitempty"`
	UserLists []interface{} `json:"userLists,omitempty"`
	SubscribedAfterDate *interface{} `json:"subscribedAfterDate,omitempty"`
	ExcludeUserLists []interface{} `json:"excludeUserLists,omitempty"`
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
	TrialLinkName string `json:"trialLinkName"`
	SubscribeDays int64 `json:"subscribeDays"`
	SubscribeCounts *int64 `json:"subscribeCounts,omitempty"`
	ExpiredAt *string `json:"expiredAt,omitempty"`
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
	TrialLinkName *string `json:"trialLinkName,omitempty"`
	SubscribeDays *int64 `json:"subscribeDays,omitempty"`
	SubscribeCounts *int64 `json:"subscribeCounts,omitempty"`
	ExpiredAt *string `json:"expiredAt,omitempty"`
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
	Message string `json:"message"`
	FinishDays int64 `json:"finishDays"`
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
	Message *string `json:"message,omitempty"`
	FinishDays *int64 `json:"finishDays,omitempty"`
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

// CreateAccessUsersLists Create user list
func (c *Client) CreateAccessUsersLists(ctx context.Context, body struct {
	Name string `json:"name"`
}) (*CreateAccessUsersListsResponse, error) {
	path := "/v2/access/users/lists"
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

// CreateAccessUsersLists1 Add user to multiple lists
func (c *Client) CreateAccessUsersLists1(ctx context.Context, userId string, body struct {
	ListIds []int64 `json:"listIds"`
}) (*CreateAccessUsersLists1Response, error) {
	path := fmt.Sprintf("/v2/access/users/%v/lists", userId)
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
	Name string `json:"name"`
	ClearMedia *bool `json:"clearMedia,omitempty"`
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

// AccessUploadsInit Initialize media upload
func (c *Client) AccessUploadsInit(ctx context.Context, body struct {
	Filename string `json:"filename"`
	Size int64 `json:"size"`
	ContentType string `json:"contentType"`
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

// LinkInit Initialize a Link session
func (c *Client) LinkInit(ctx context.Context, body struct {
	RedirectUrl *string `json:"redirectUrl,omitempty"`
	ClientReferenceId *string `json:"clientReferenceId,omitempty"`
	ConnectionId *string `json:"connectionId,omitempty"`
	Geolocation *struct {
	Country string `json:"country"`
	State *string `json:"state,omitempty"`
	City *string `json:"city,omitempty"`
} `json:"geolocation,omitempty"`
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

