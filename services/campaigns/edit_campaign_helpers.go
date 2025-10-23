package campaigns

// NewEditCampaignRequest creates a new EditCampaignRequest with the provided
// name (optional). Use the fluent With* methods to set additional fields.
func NewEditCampaignRequest(name string) *EditCampaignRequest {
	if name == "" {
		return &EditCampaignRequest{}
	}
	return &EditCampaignRequest{Name: &name}
}

// WithType sets the campaign type.
func (r *EditCampaignRequest) WithType(typ string) *EditCampaignRequest {
	if r == nil {
		r = &EditCampaignRequest{}
	}
	r.Type = &typ
	return r
}

// WithSegmentID sets the segment id (string-based IDs supported).
func (r *EditCampaignRequest) WithSegmentID(seg string) *EditCampaignRequest {
	if r == nil {
		r = &EditCampaignRequest{}
	}
	r.SegmentID = &seg
	return r
}

// WithAddressID sets the selected address id.
func (r *EditCampaignRequest) WithAddressID(id int64) *EditCampaignRequest {
	if r == nil {
		r = &EditCampaignRequest{}
	}
	r.AddressID = &id
	return r
}

// WithListIDs sets the list IDs to target for this campaign.
func (r *EditCampaignRequest) WithListIDs(ids ...int64) *EditCampaignRequest {
	if r == nil {
		r = &EditCampaignRequest{}
	}
	r.ListIDs = append([]int64{}, ids...)
	return r
}

// WithReplyTrackingEnabled toggles reply tracking.
func (r *EditCampaignRequest) WithReplyTrackingEnabled(on bool) *EditCampaignRequest {
	if r == nil {
		r = &EditCampaignRequest{}
	}
	r.ReplyTrackingEnabled = &on
	return r
}

// WithLinkTrackingEnabled toggles link tracking.
func (r *EditCampaignRequest) WithLinkTrackingEnabled(on bool) *EditCampaignRequest {
	if r == nil {
		r = &EditCampaignRequest{}
	}
	r.LinkTrackingEnabled = &on
	return r
}

// WithGoogleAnalyticsLinkTrackingEnabled toggles GA link tracking.
func (r *EditCampaignRequest) WithGoogleAnalyticsLinkTrackingEnabled(on bool) *EditCampaignRequest {
	if r == nil {
		r = &EditCampaignRequest{}
	}
	r.GoogleAnalyticsLinkTrackingEnabled = &on
	return r
}

// WithGoogleAnalyticsCampaignName sets the GA campaign name.
func (r *EditCampaignRequest) WithGoogleAnalyticsCampaignName(name string) *EditCampaignRequest {
	if r == nil {
		r = &EditCampaignRequest{}
	}
	r.GoogleAnalyticsCampaignName = &name
	return r
}

// WithReadTrackingEnabled toggles read tracking.
func (r *EditCampaignRequest) WithReadTrackingEnabled(on bool) *EditCampaignRequest {
	if r == nil {
		r = &EditCampaignRequest{}
	}
	r.ReadTrackingEnabled = &on
	return r
}

// WithSendToExistingSubscribers sets sendToExistingSubscribers.
func (r *EditCampaignRequest) WithSendToExistingSubscribers(on bool) *EditCampaignRequest {
	if r == nil {
		r = &EditCampaignRequest{}
	}
	r.SendToExistingSubscribers = &on
	return r
}

// WithCanSplitContent sets canSplitContent flag.
func (r *EditCampaignRequest) WithCanSplitContent(on bool) *EditCampaignRequest {
	if r == nil {
		r = &EditCampaignRequest{}
	}
	r.CanSplitContent = &on
	return r
}

// WithRecurring sets recurring flag.
func (r *EditCampaignRequest) WithRecurring(on bool) *EditCampaignRequest {
	if r == nil {
		r = &EditCampaignRequest{}
	}
	r.Recurring = &on
	return r
}

// WithResponderDaysOffset sets responderDaysOffset.
func (r *EditCampaignRequest) WithResponderDaysOffset(days int32) *EditCampaignRequest {
	if r == nil {
		r = &EditCampaignRequest{}
	}
	r.ResponderDaysOffset = &days
	return r
}

// WithResponderHoursOffset sets responderHoursOffset.
func (r *EditCampaignRequest) WithResponderHoursOffset(hours int32) *EditCampaignRequest {
	if r == nil {
		r = &EditCampaignRequest{}
	}
	r.ResponderHoursOffset = &hours
	return r
}

// WithScheduledDate sets scheduledDate string.
func (r *EditCampaignRequest) WithScheduledDate(date string) *EditCampaignRequest {
	if r == nil {
		r = &EditCampaignRequest{}
	}
	r.ScheduledDate = &date
	return r
}

// WithReminderField sets reminderField.
func (r *EditCampaignRequest) WithReminderField(field string) *EditCampaignRequest {
	if r == nil {
		r = &EditCampaignRequest{}
	}
	r.ReminderField = &field
	return r
}

// WithReminderOffset sets reminderOffset.
func (r *EditCampaignRequest) WithReminderOffset(offset int32) *EditCampaignRequest {
	if r == nil {
		r = &EditCampaignRequest{}
	}
	r.ReminderOffset = &offset
	return r
}

// WithReminderOffsetType sets reminderOffsetType.
func (r *EditCampaignRequest) WithReminderOffsetType(t string) *EditCampaignRequest {
	if r == nil {
		r = &EditCampaignRequest{}
	}
	r.ReminderOffsetType = &t
	return r
}

// WithReminderType sets reminderType.
func (r *EditCampaignRequest) WithReminderType(t string) *EditCampaignRequest {
	if r == nil {
		r = &EditCampaignRequest{}
	}
	r.ReminderType = &t
	return r
}

// WithRSSInterval sets rssInterval.
func (r *EditCampaignRequest) WithRSSInterval(i int32) *EditCampaignRequest {
	if r == nil {
		r = &EditCampaignRequest{}
	}
	r.RSSInterval = &i
	return r
}

// WithSplitType sets splitType.
func (r *EditCampaignRequest) WithSplitType(t string) *EditCampaignRequest {
	if r == nil {
		r = &EditCampaignRequest{}
	}
	r.SplitType = &t
	return r
}

// WithSplitWinnerWaitPeriod sets splitWinnerWaitPeriod.
func (r *EditCampaignRequest) WithSplitWinnerWaitPeriod(p int32) *EditCampaignRequest {
	if r == nil {
		r = &EditCampaignRequest{}
	}
	r.SplitWinnerWaitPeriod = &p
	return r
}

// WithSplitWinnerWaitPeriodType sets splitWinnerWaitPeriodType.
func (r *EditCampaignRequest) WithSplitWinnerWaitPeriodType(t string) *EditCampaignRequest {
	if r == nil {
		r = &EditCampaignRequest{}
	}
	r.SplitWinnerWaitPeriodType = &t
	return r
}

// WithPublicCampaignArchive sets publicCampaignArchive flag.
func (r *EditCampaignRequest) WithPublicCampaignArchive(on bool) *EditCampaignRequest {
	if r == nil {
		r = &EditCampaignRequest{}
	}
	r.PublicCampaignArchive = &on
	return r
}
