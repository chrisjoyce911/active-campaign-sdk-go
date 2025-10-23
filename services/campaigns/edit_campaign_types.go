package campaigns

// EditCampaignRequest represents the fields accepted by PUT /campaigns/{id}/edit.
// Many fields are optional; use pointers when you need to distinguish zero vs unset.
type EditCampaignRequest struct {
	Name                               *string `json:"name,omitempty"`
	Type                               *string `json:"type,omitempty"`
	SegmentID                          *string `json:"segmentId,omitempty"`
	AddressID                          *int64  `json:"addressId,omitempty"`
	ListIDs                            []int64 `json:"listIds,omitempty"`
	ReplyTrackingEnabled               *bool   `json:"replyTrackingEnabled,omitempty"`
	LinkTrackingEnabled                *bool   `json:"linkTrackingEnabled,omitempty"`
	GoogleAnalyticsLinkTrackingEnabled *bool   `json:"googleAnalyticsLinkTrackingEnabled,omitempty"`
	GoogleAnalyticsCampaignName        *string `json:"googleAnalyticsCampaignName,omitempty"`
	ReadTrackingEnabled                *bool   `json:"readTrackingEnabled,omitempty"`
	SendToExistingSubscribers          *bool   `json:"sendToExistingSubscribers,omitempty"`
	CanSplitContent                    *bool   `json:"canSplitContent,omitempty"`
	Recurring                          *bool   `json:"recurring,omitempty"`
	ResponderDaysOffset                *int32  `json:"responderDaysOffset,omitempty"`
	ResponderHoursOffset               *int32  `json:"responderHoursOffset,omitempty"`
	ScheduledDate                      *string `json:"scheduledDate,omitempty"`
	ReminderField                      *string `json:"reminderField,omitempty"`
	ReminderOffset                     *int32  `json:"reminderOffset,omitempty"`
	ReminderOffsetType                 *string `json:"reminderOffsetType,omitempty"`
	ReminderType                       *string `json:"reminderType,omitempty"`
	RSSInterval                        *int32  `json:"rssInterval,omitempty"`
	SplitType                          *string `json:"splitType,omitempty"`
	SplitWinnerWaitPeriod              *int32  `json:"splitWinnerWaitPeriod,omitempty"`
	SplitWinnerWaitPeriodType          *string `json:"splitWinnerWaitPeriodType,omitempty"`
	PublicCampaignArchive              *bool   `json:"publicCampaignArchive,omitempty"`
}
