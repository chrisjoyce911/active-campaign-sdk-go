package campaignsmock

import (
	"context"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/campaigns"
)

// Service is a function-field mock implementing campaigns.CampaignsService.
type Service struct {
	CreateCampaignFunc    func(ctx context.Context, req *campaigns.CreateCampaignRequest) (*campaigns.Campaign, *client.APIResponse, error)
	GetCampaignFunc       func(ctx context.Context, id string) (*campaigns.Campaign, *client.APIResponse, error)
	ListCampaignsFunc     func(ctx context.Context, opts interface{}) (*campaigns.ListCampaignsResponse, *client.APIResponse, error)
	GetCampaignLinksFunc  func(ctx context.Context, id string) (*campaigns.CampaignLinksResponse, *client.APIResponse, error)
	CampaignLinksFunc     func(ctx context.Context, id string, messageID *string) ([]campaigns.CampaignLink, *client.APIResponse, error)
	EditCampaignFunc      func(ctx context.Context, id string, req *campaigns.EditCampaignRequest) (*campaigns.Campaign, *client.APIResponse, error)
	DuplicateCampaignFunc func(ctx context.Context, id string) (*campaigns.DuplicateCampaignResponse, *client.APIResponse, error)
}

var _ campaigns.CampaignsService = (*Service)(nil)

func (m *Service) CreateCampaign(ctx context.Context, req *campaigns.CreateCampaignRequest) (*campaigns.Campaign, *client.APIResponse, error) {
	if m.CreateCampaignFunc != nil {
		return m.CreateCampaignFunc(ctx, req)
	}
	return nil, nil, nil
}
func (m *Service) GetCampaign(ctx context.Context, id string) (*campaigns.Campaign, *client.APIResponse, error) {
	if m.GetCampaignFunc != nil {
		return m.GetCampaignFunc(ctx, id)
	}
	return nil, nil, nil
}
func (m *Service) ListCampaigns(ctx context.Context, opts interface{}) (*campaigns.ListCampaignsResponse, *client.APIResponse, error) {
	if m.ListCampaignsFunc != nil {
		return m.ListCampaignsFunc(ctx, opts)
	}
	return &campaigns.ListCampaignsResponse{}, &client.APIResponse{}, nil
}
func (m *Service) GetCampaignLinks(ctx context.Context, id string) (*campaigns.CampaignLinksResponse, *client.APIResponse, error) {
	if m.GetCampaignLinksFunc != nil {
		return m.GetCampaignLinksFunc(ctx, id)
	}
	return nil, nil, nil
}
func (m *Service) CampaignLinks(ctx context.Context, id string, messageID *string) ([]campaigns.CampaignLink, *client.APIResponse, error) {
	if m.CampaignLinksFunc != nil {
		return m.CampaignLinksFunc(ctx, id, messageID)
	}
	return nil, nil, nil
}
func (m *Service) EditCampaign(ctx context.Context, id string, req *campaigns.EditCampaignRequest) (*campaigns.Campaign, *client.APIResponse, error) {
	if m.EditCampaignFunc != nil {
		return m.EditCampaignFunc(ctx, id, req)
	}
	return nil, nil, nil
}
func (m *Service) DuplicateCampaign(ctx context.Context, id string) (*campaigns.DuplicateCampaignResponse, *client.APIResponse, error) {
	if m.DuplicateCampaignFunc != nil {
		return m.DuplicateCampaignFunc(ctx, id)
	}
	return nil, nil, nil
}
