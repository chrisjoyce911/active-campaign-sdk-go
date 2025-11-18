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

func TestService_AllFuncs(t *testing.T) {
	called := 0
	m := &Service{
		CreateCampaignFunc: func(ctx context.Context, req *campaigns.CreateCampaignRequest) (*campaigns.Campaign, *client.APIResponse, error) {
			called++
			return &campaigns.Campaign{ID: "c"}, &client.APIResponse{StatusCode: 200}, nil
		},
		GetCampaignFunc: func(ctx context.Context, id string) (*campaigns.Campaign, *client.APIResponse, error) {
			called++
			return &campaigns.Campaign{ID: id}, &client.APIResponse{StatusCode: 200}, nil
		},
		ListCampaignsFunc: func(ctx context.Context, _ interface{}) (*campaigns.ListCampaignsResponse, *client.APIResponse, error) {
			called++
			return &campaigns.ListCampaignsResponse{}, &client.APIResponse{StatusCode: 200}, nil
		},
		GetCampaignLinksFunc: func(ctx context.Context, id string) (*campaigns.CampaignLinksResponse, *client.APIResponse, error) {
			called++
			return &campaigns.CampaignLinksResponse{}, &client.APIResponse{StatusCode: 200}, nil
		},
		CampaignLinksFunc: func(ctx context.Context, id string, messageID *string) ([]campaigns.CampaignLink, *client.APIResponse, error) {
			called++
			return []campaigns.CampaignLink{}, &client.APIResponse{StatusCode: 200}, nil
		},
		EditCampaignFunc: func(ctx context.Context, id string, req *campaigns.EditCampaignRequest) (*campaigns.Campaign, *client.APIResponse, error) {
			called++
			return &campaigns.Campaign{ID: id}, &client.APIResponse{StatusCode: 200}, nil
		},
		DuplicateCampaignFunc: func(ctx context.Context, id string) (*campaigns.DuplicateCampaignResponse, *client.APIResponse, error) {
			called++
			return &campaigns.DuplicateCampaignResponse{}, &client.APIResponse{StatusCode: 200}, nil
		},
	}

	_, _, _ = m.CreateCampaign(context.Background(), &campaigns.CreateCampaignRequest{})
	_, _, _ = m.GetCampaign(context.Background(), "c1")
	_, _, _ = m.ListCampaigns(context.Background(), nil)
	_, _, _ = m.GetCampaignLinks(context.Background(), "c1")
	_, _, _ = m.CampaignLinks(context.Background(), "c1", nil)
	_, _, _ = m.EditCampaign(context.Background(), "c1", &campaigns.EditCampaignRequest{})
	_, _, _ = m.DuplicateCampaign(context.Background(), "c1")

	if called != 7 {
		t.Fatalf("expected 7 calls, got %d", called)
	}
}

func TestService_AllDefaults(t *testing.T) {
	m := &Service{}
	ctx := context.Background()
	_, _, _ = m.CreateCampaign(ctx, &campaigns.CreateCampaignRequest{})
	_, _, _ = m.GetCampaign(ctx, "c1")
	_, _, _ = m.ListCampaigns(ctx, nil)
	_, _, _ = m.GetCampaignLinks(ctx, "c1")
	_, _, _ = m.CampaignLinks(ctx, "c1", nil)
	_, _, _ = m.EditCampaign(ctx, "c1", &campaigns.EditCampaignRequest{})
	_, _, _ = m.DuplicateCampaign(ctx, "c1")
}
