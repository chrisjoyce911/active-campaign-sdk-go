package campaigns

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestNewRealService_AssignsClient(t *testing.T) {
	core := &client.CoreClient{}
	s := NewRealService(core)
	assert.NotNil(t, s)
	// concrete type assertion
	svc, ok := s.(*service)
	assert.True(t, ok)
	assert.Equal(t, core, svc.client)
}

func TestEditCampaignRequest_FluentSettersAndNilReceiver(t *testing.T) {
	r := NewEditCampaignRequest("name")
	assert.NotNil(t, r)
	assert.NotNil(t, r.Name)
	// chain a few setters
	r = r.WithType("regular").WithReplyTrackingEnabled(true).WithListIDs(1, 2, 3)
	assert.NotNil(t, r.Type)
	assert.NotNil(t, r.ReplyTrackingEnabled)
	assert.Equal(t, 3, len(r.ListIDs))

	// nil receiver behavior — call methods on nil pointer
	var rn *EditCampaignRequest
	rn = rn.WithType("x").WithRecurring(true)
	assert.NotNil(t, rn)
	assert.NotNil(t, rn.Type)
	assert.NotNil(t, rn.Recurring)
}

func TestCampaignStatus_String(t *testing.T) {
	cases := []struct {
		in   CampaignStatus
		want string
	}{
		{CampaignStatusDraft, "Draft"},
		{CampaignStatusScheduled, "Scheduled"},
		{CampaignStatusSending, "Sending"},
		{CampaignStatusPaused, "Paused"},
		{CampaignStatusStopped, "Stopped"},
		{CampaignStatusCompleted, "Completed"},
		{CampaignStatus(999), "Unknown"},
	}
	for _, c := range cases {
		assert.Equal(t, c.want, c.in.String())
	}
}

func TestCampaignHelpers_StatusParsingAndEnum(t *testing.T) {
	c := &Campaign{Status: "2"}
	i, err := c.StatusInt()
	assert.NoError(t, err)
	assert.Equal(t, 2, i)
	assert.Equal(t, CampaignStatusSending, c.StatusEnum())

	// empty status
	c2 := &Campaign{Status: ""}
	i2, err2 := c2.StatusInt()
	assert.Error(t, err2)
	assert.Equal(t, 0, i2)
	assert.Equal(t, CampaignStatusDraft, c2.StatusEnum())

	// nil campaign
	var cnil *Campaign
	i3, err3 := cnil.StatusInt()
	assert.NoError(t, err3)
	assert.Equal(t, 0, i3)
	assert.Equal(t, CampaignStatusDraft, cnil.StatusEnum())
}

func TestGetCampaignLinks_Filtering(t *testing.T) {
	// set up a HTTPDoer that returns links
	resp := `{"links":[{"id":"l1","message":"m1"},{"id":"l2","message":null},{"id":"l3","message":"m2"}]}`
	hd := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(resp)}
	svc := NewRealServiceFromDoer(hd)

	out, apiResp, err := svc.GetCampaignLinks(context.Background(), "c1")
	assert.NoError(t, err)
	assert.NotNil(t, apiResp)
	assert.NotNil(t, out)
	assert.Len(t, out.Links, 3)

	// CampaignLinks convenience
	linksAll, _, err := svc.CampaignLinks(context.Background(), "c1", nil)
	assert.NoError(t, err)
	assert.Len(t, linksAll, 3)

	msg := "m1"
	linksFiltered, _, err := svc.CampaignLinks(context.Background(), "c1", &msg)
	assert.NoError(t, err)
	assert.Len(t, linksFiltered, 1)
	assert.Equal(t, "l1", linksFiltered[0].ID)

	// when GetCampaignLinks returns nil out, CampaignLinks should return nil slice
	hd2 := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: nil}
	svc2 := NewRealServiceFromDoer(hd2)
	linksNil, _, err := svc2.CampaignLinks(context.Background(), "c1", nil)
	assert.NoError(t, err)
	assert.Nil(t, linksNil)
}

func TestUpdateCampaign_NotImplemented(t *testing.T) {
	_, _, err := (&service{}).UpdateCampaign(context.Background(), "id", nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")
}
