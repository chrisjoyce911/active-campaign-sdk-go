package campaigns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEditCampaignHelpers_ExhaustiveCalls(t *testing.T) {
	// non-nil receiver
	r := &EditCampaignRequest{}
	_ = r.WithType("a")
	_ = r.WithSegmentID("b")
	_ = r.WithAddressID(1)
	_ = r.WithListIDs(1,2)
	_ = r.WithReplyTrackingEnabled(true)
	_ = r.WithLinkTrackingEnabled(true)
	_ = r.WithGoogleAnalyticsLinkTrackingEnabled(true)
	_ = r.WithGoogleAnalyticsCampaignName("ga")
	_ = r.WithReadTrackingEnabled(true)
	_ = r.WithSendToExistingSubscribers(true)
	_ = r.WithCanSplitContent(true)
	_ = r.WithRecurring(true)
	_ = r.WithResponderDaysOffset(1)
	_ = r.WithResponderHoursOffset(2)
	_ = r.WithScheduledDate("d")
	_ = r.WithReminderField("rf")
	_ = r.WithReminderOffset(3)
	_ = r.WithReminderOffsetType("rot")
	_ = r.WithReminderType("rt")
	_ = r.WithRSSInterval(4)
	_ = r.WithSplitType("st")
	_ = r.WithSplitWinnerWaitPeriod(5)
	_ = r.WithSplitWinnerWaitPeriodType("sw")
	_ = r.WithPublicCampaignArchive(true)

	// nil receiver
	var rn *EditCampaignRequest
	rn = rn.WithType("a")
	rn = rn.WithSegmentID("b")
	rn = rn.WithAddressID(1)
	rn = rn.WithListIDs(1,2)
	rn = rn.WithReplyTrackingEnabled(true)
	rn = rn.WithLinkTrackingEnabled(true)
	rn = rn.WithGoogleAnalyticsLinkTrackingEnabled(true)
	rn = rn.WithGoogleAnalyticsCampaignName("ga")
	rn = rn.WithReadTrackingEnabled(true)
	rn = rn.WithSendToExistingSubscribers(true)
	rn = rn.WithCanSplitContent(true)
	rn = rn.WithRecurring(true)
	rn = rn.WithResponderDaysOffset(1)
	rn = rn.WithResponderHoursOffset(2)
	rn = rn.WithScheduledDate("d")
	rn = rn.WithReminderField("rf")
	rn = rn.WithReminderOffset(3)
	rn = rn.WithReminderOffsetType("rot")
	rn = rn.WithReminderType("rt")
	rn = rn.WithRSSInterval(4)
	rn = rn.WithSplitType("st")
	rn = rn.WithSplitWinnerWaitPeriod(5)
	rn = rn.WithSplitWinnerWaitPeriodType("sw")
	rn = rn.WithPublicCampaignArchive(false)

	assert.NotNil(t, r)
	assert.NotNil(t, rn)
}
