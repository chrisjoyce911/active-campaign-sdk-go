package campaigns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEditCampaignHelpers_ExhaustiveNilAndNonNil(t *testing.T) {
	// non-nil receiver
	r := &EditCampaignRequest{}
	r.WithType("t")
	r.WithSegmentID("s")
	r.WithAddressID(1)
	r.WithListIDs(1,2,3)
	r.WithReplyTrackingEnabled(true)
	r.WithLinkTrackingEnabled(true)
	r.WithGoogleAnalyticsLinkTrackingEnabled(true)
	r.WithGoogleAnalyticsCampaignName("ga")
	r.WithReadTrackingEnabled(true)
	r.WithSendToExistingSubscribers(true)
	r.WithCanSplitContent(true)
	r.WithRecurring(true)
	r.WithResponderDaysOffset(1)
	r.WithResponderHoursOffset(2)
	r.WithScheduledDate("d")
	r.WithReminderField("f")
	r.WithReminderOffset(3)
	r.WithReminderOffsetType("t")
	r.WithReminderType("rt")
	r.WithRSSInterval(5)
	r.WithSplitType("st")
	r.WithSplitWinnerWaitPeriod(6)
	r.WithSplitWinnerWaitPeriodType("wt")
	r.WithPublicCampaignArchive(true)

	assert.NotNil(t, r)

	// nil receiver
	var rn *EditCampaignRequest
	rn = rn.WithType("t")
	rn = rn.WithSegmentID("s")
	rn = rn.WithAddressID(1)
	rn = rn.WithListIDs(1)
	rn = rn.WithReplyTrackingEnabled(false)
	rn = rn.WithLinkTrackingEnabled(false)
	rn = rn.WithGoogleAnalyticsLinkTrackingEnabled(false)
	rn = rn.WithGoogleAnalyticsCampaignName("ga")
	rn = rn.WithReadTrackingEnabled(false)
	rn = rn.WithSendToExistingSubscribers(false)
	rn = rn.WithCanSplitContent(false)
	rn = rn.WithRecurring(false)
	rn = rn.WithResponderDaysOffset(0)
	rn = rn.WithResponderHoursOffset(0)
	rn = rn.WithScheduledDate("")
	rn = rn.WithReminderField("")
	rn = rn.WithReminderOffset(0)
	rn = rn.WithReminderOffsetType("")
	rn = rn.WithReminderType("")
	rn = rn.WithRSSInterval(0)
	rn = rn.WithSplitType("")
	rn = rn.WithSplitWinnerWaitPeriod(0)
	rn = rn.WithSplitWinnerWaitPeriodType("")
	rn = rn.WithPublicCampaignArchive(false)

	assert.NotNil(t, rn)
}
