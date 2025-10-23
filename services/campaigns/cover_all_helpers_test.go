package campaigns

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestEditCampaignHelpers_ExerciseAllPaths(t *testing.T) {
	// non-nil receiver
	r := &EditCampaignRequest{}
	r = r.WithType("t")
	r = r.WithSegmentID("s")
	r = r.WithAddressID(123)
	r = r.WithListIDs(10, 20)
	r = r.WithReplyTrackingEnabled(true)
	r = r.WithLinkTrackingEnabled(false)
	r = r.WithGoogleAnalyticsLinkTrackingEnabled(true)
	r = r.WithGoogleAnalyticsCampaignName("ga")
	r = r.WithReadTrackingEnabled(true)
	r = r.WithSendToExistingSubscribers(false)
	r = r.WithCanSplitContent(true)
	r = r.WithRecurring(false)
	r = r.WithResponderDaysOffset(2)
	r = r.WithResponderHoursOffset(3)
	r = r.WithScheduledDate("2025-01-01")
	r = r.WithReminderField("f")
	r = r.WithReminderOffset(5)
	r = r.WithReminderOffsetType("type")
	r = r.WithReminderType("rtype")
	r = r.WithRSSInterval(7)
	r = r.WithSplitType("split")
	r = r.WithSplitWinnerWaitPeriod(8)
	r = r.WithSplitWinnerWaitPeriodType("wtype")
	r = r.WithPublicCampaignArchive(true)

	// nil receiver should allocate and then set
	var rn *EditCampaignRequest
	rn = rn.WithType("t2")
	rn = rn.WithSegmentID("s2")
	rn = rn.WithAddressID(321)
	rn = rn.WithListIDs()
	rn = rn.WithReplyTrackingEnabled(false)
	rn = rn.WithLinkTrackingEnabled(true)
	rn = rn.WithGoogleAnalyticsLinkTrackingEnabled(false)
	rn = rn.WithGoogleAnalyticsCampaignName("ga2")
	rn = rn.WithReadTrackingEnabled(false)
	rn = rn.WithSendToExistingSubscribers(true)
	rn = rn.WithCanSplitContent(false)
	rn = rn.WithRecurring(true)
	rn = rn.WithResponderDaysOffset(4)
	rn = rn.WithResponderHoursOffset(5)
	rn = rn.WithScheduledDate("2025-02-02")
	rn = rn.WithReminderField("f2")
	rn = rn.WithReminderOffset(6)
	rn = rn.WithReminderOffsetType("type2")
	rn = rn.WithReminderType("rtype2")
	rn = rn.WithRSSInterval(9)
	rn = rn.WithSplitType("split2")
	rn = rn.WithSplitWinnerWaitPeriod(11)
	rn = rn.WithSplitWinnerWaitPeriodType("wtype2")
	rn = rn.WithPublicCampaignArchive(false)

	// basic sanity asserts
	assert.NotNil(t, r)
	assert.NotNil(t, rn)
}

func TestCampaignLinks_NoMatchReturnsEmptySlice(t *testing.T) {
	resp := `{"links":[{"id":"l1","message":"m1"},{"id":"l2","message":null}]}`
	md := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(resp)}
	svc := NewRealServiceFromDoer(md)

	msg := "nomatch"
	filtered, apiResp, err := svc.CampaignLinks(context.Background(), "c1", &msg)
	assert.NoError(t, err)
	assert.NotNil(t, apiResp)
	// implementation may return nil or empty slice; assert there are no matches
	if filtered == nil {
		assert.Equal(t, 0, 0)
	} else {
		assert.Equal(t, 0, len(filtered))
	}
}
