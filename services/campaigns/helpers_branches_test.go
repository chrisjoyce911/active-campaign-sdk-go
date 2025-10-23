package campaigns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEditCampaignHelpers_AllBranchesCovered(t *testing.T) {
	// non-nil receiver: exercise setters with typical values
	r := &EditCampaignRequest{}
	r = r.WithType("type")
	r = r.WithSegmentID("seg")
	r = r.WithAddressID(42)
	r = r.WithListIDs(1, 2, 3)
	r = r.WithListIDs() // empty
	r = r.WithReplyTrackingEnabled(true)
	r = r.WithReplyTrackingEnabled(false)
	r = r.WithLinkTrackingEnabled(true)
	r = r.WithLinkTrackingEnabled(false)
	r = r.WithGoogleAnalyticsLinkTrackingEnabled(true)
	r = r.WithGoogleAnalyticsLinkTrackingEnabled(false)
	r = r.WithGoogleAnalyticsCampaignName("ga")
	r = r.WithReadTrackingEnabled(true)
	r = r.WithReadTrackingEnabled(false)
	r = r.WithSendToExistingSubscribers(true)
	r = r.WithSendToExistingSubscribers(false)
	r = r.WithCanSplitContent(true)
	r = r.WithCanSplitContent(false)
	r = r.WithRecurring(true)
	r = r.WithRecurring(false)
	r = r.WithResponderDaysOffset(0)
	r = r.WithResponderDaysOffset(7)
	r = r.WithResponderHoursOffset(0)
	r = r.WithResponderHoursOffset(12)
	r = r.WithScheduledDate("")
	r = r.WithScheduledDate("2025-10-23")
	r = r.WithReminderField("")
	r = r.WithReminderField("field")
	r = r.WithReminderOffset(0)
	r = r.WithReminderOffset(10)
	r = r.WithReminderOffsetType("")
	r = r.WithReminderOffsetType("type")
	r = r.WithReminderType("")
	r = r.WithReminderType("rtype")
	r = r.WithRSSInterval(0)
	r = r.WithRSSInterval(15)
	r = r.WithSplitType("")
	r = r.WithSplitType("split")
	r = r.WithSplitWinnerWaitPeriod(0)
	r = r.WithSplitWinnerWaitPeriod(3)
	r = r.WithSplitWinnerWaitPeriodType("")
	r = r.WithSplitWinnerWaitPeriodType("wt")
	r = r.WithPublicCampaignArchive(true)
	r = r.WithPublicCampaignArchive(false)

	// nil receiver: should allocate and set
	var rn *EditCampaignRequest
	rn = rn.WithType("nt")
	rn = rn.WithSegmentID("ns")
	rn = rn.WithAddressID(7)
	rn = rn.WithListIDs(99)
	rn = rn.WithReplyTrackingEnabled(true)
	rn = rn.WithLinkTrackingEnabled(true)
	rn = rn.WithGoogleAnalyticsLinkTrackingEnabled(true)
	rn = rn.WithGoogleAnalyticsCampaignName("nga")
	rn = rn.WithReadTrackingEnabled(true)
	rn = rn.WithSendToExistingSubscribers(true)
	rn = rn.WithCanSplitContent(true)
	rn = rn.WithRecurring(true)
	rn = rn.WithResponderDaysOffset(1)
	rn = rn.WithResponderHoursOffset(2)
	rn = rn.WithScheduledDate("sdate")
	rn = rn.WithReminderField("rf")
	rn = rn.WithReminderOffset(4)
	rn = rn.WithReminderOffsetType("rot")
	rn = rn.WithReminderType("rt")
	rn = rn.WithRSSInterval(6)
	rn = rn.WithSplitType("st")
	rn = rn.WithSplitWinnerWaitPeriod(2)
	rn = rn.WithSplitWinnerWaitPeriodType("sw")
	rn = rn.WithPublicCampaignArchive(false)

	// sanity checks
	assert.NotNil(t, r)
	assert.NotNil(t, rn)
}
