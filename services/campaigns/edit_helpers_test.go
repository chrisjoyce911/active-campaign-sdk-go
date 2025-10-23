package campaigns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEditCampaignHelpers_AllSetters(t *testing.T) {
	r := NewEditCampaignRequest("")
	// call each setter once to exercise branches
	r = r.WithType("t")
	r = r.WithSegmentID("s")
	r = r.WithAddressID(10)
	r = r.WithListIDs(1,2)
	r = r.WithReplyTrackingEnabled(true)
	r = r.WithLinkTrackingEnabled(true)
	r = r.WithGoogleAnalyticsLinkTrackingEnabled(true)
	r = r.WithGoogleAnalyticsCampaignName("ga")
	r = r.WithReadTrackingEnabled(true)
	r = r.WithSendToExistingSubscribers(true)
	r = r.WithCanSplitContent(true)
	r = r.WithRecurring(true)
	r = r.WithResponderDaysOffset(3)
	r = r.WithResponderHoursOffset(4)
	r = r.WithScheduledDate("2025-01-01")
	r = r.WithReminderField("field")
	r = r.WithReminderOffset(5)
	r = r.WithReminderOffsetType("type")
	r = r.WithReminderType("rtype")
	r = r.WithRSSInterval(7)
	r = r.WithSplitType("st")
	r = r.WithSplitWinnerWaitPeriod(8)
	r = r.WithSplitWinnerWaitPeriodType("wtype")
	r = r.WithPublicCampaignArchive(true)

	// verify some fields were set
	assert.NotNil(t, r.Type)
	assert.Equal(t, int64(10), *r.AddressID)
	assert.Equal(t, 2, len(r.ListIDs))
	assert.NotNil(t, r.GoogleAnalyticsCampaignName)
}
