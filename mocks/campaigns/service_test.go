package campaignsmock

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/campaigns"
)

func TestService_ListCampaignsFuncCalled(t *testing.T) {
	called := false
	m := &Service{ListCampaignsFunc: func(ctx context.Context, _ interface{}) (*campaigns.ListCampaignsResponse, *client.APIResponse, error) {
		called = true
		return &campaigns.ListCampaignsResponse{Campaigns: []campaigns.Campaign{{ID: "c1", Status: "1"}}}, &client.APIResponse{StatusCode: 200}, nil
	}}
	resp, api, err := m.ListCampaigns(context.Background(), nil)
	if err != nil || api == nil || resp == nil || len(resp.Campaigns) != 1 || !called {
		t.Fatalf("unexpected result: resp=%v api=%v err=%v called=%v", resp, api, err, called)
	}
}

func TestService_Defaults(t *testing.T) {
	m := &Service{}
	resp, api, err := m.ListCampaigns(context.Background(), nil)
	if err != nil || resp == nil || api == nil {
		t.Fatalf("expected non-nil defaults, got resp=%v api=%v err=%v", resp, api, err)
	}
}
