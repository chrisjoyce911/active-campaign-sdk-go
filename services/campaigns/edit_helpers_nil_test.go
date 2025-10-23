package campaigns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEditCampaignHelpers_NilReceiverAllSetters(t *testing.T) {
	// call every With* on a nil receiver to exercise nil receiver branch
	var r *EditCampaignRequest
	r = r.WithType("t")
	assert.NotNil(t, r)

	r = nil
	r = r.WithSegmentID("s")
	assert.NotNil(t, r.SegmentID)

	r = nil
	r = r.WithAddressID(1)
	assert.NotNil(t, r.AddressID)

	r = nil
	r = r.WithListIDs()
	// should handle empty list
	assert.NotNil(t, r)

	r = nil
	r = r.WithReplyTrackingEnabled(true)
	assert.NotNil(t, r.ReplyTrackingEnabled)

	r = nil
	r = r.WithLinkTrackingEnabled(true)
	assert.NotNil(t, r.LinkTrackingEnabled)

	r = nil
	r = r.WithGoogleAnalyticsLinkTrackingEnabled(true)
	assert.NotNil(t, r.GoogleAnalyticsLinkTrackingEnabled)

	r = nil
	r = r.WithGoogleAnalyticsCampaignName("ga")
	assert.NotNil(t, r.GoogleAnalyticsCampaignName)

	r = nil
	r = r.WithReadTrackingEnabled(true)
	assert.NotNil(t, r.ReadTrackingEnabled)

	r = nil
	r = r.WithSendToExistingSubscribers(true)
	assert.NotNil(t, r.SendToExistingSubscribers)

	r = nil
	r = r.WithCanSplitContent(true)
	assert.NotNil(t, r.CanSplitContent)

	r = nil
	r = r.WithRecurring(true)
	assert.NotNil(t, r.Recurring)

	r = nil
	r = r.WithResponderDaysOffset(2)
	assert.NotNil(t, r.ResponderDaysOffset)

	r = nil
	r = r.WithResponderHoursOffset(3)
	assert.NotNil(t, r.ResponderHoursOffset)

	r = nil
	r = r.WithScheduledDate("d")
	assert.NotNil(t, r.ScheduledDate)

	r = nil
	r = r.WithReminderField("f")
	assert.NotNil(t, r.ReminderField)

	r = nil
	r = r.WithReminderOffset(4)
	assert.NotNil(t, r.ReminderOffset)

	r = nil
	r = r.WithReminderOffsetType("t")
	assert.NotNil(t, r.ReminderOffsetType)

	r = nil
	r = r.WithReminderType("rt")
	assert.NotNil(t, r.ReminderType)

	r = nil
	r = r.WithRSSInterval(5)
	assert.NotNil(t, r.RSSInterval)

	r = nil
	r = r.WithSplitType("st")
	assert.NotNil(t, r.SplitType)

	r = nil
	r = r.WithSplitWinnerWaitPeriod(6)
	assert.NotNil(t, r.SplitWinnerWaitPeriod)

	r = nil
	r = r.WithSplitWinnerWaitPeriodType("wt")
	assert.NotNil(t, r.SplitWinnerWaitPeriodType)

	r = nil
	r = r.WithPublicCampaignArchive(true)
	assert.NotNil(t, r.PublicCampaignArchive)
}
