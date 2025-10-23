package campaigns

import (
	"context"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

type service struct {
	client client.Doer
}

// CampaignsService describes the operations available for campaigns. The
// concrete implementation is provided by NewRealService / NewRealServiceFromDoer.
type CampaignsService interface {
	CreateCampaign(ctx context.Context, req *CreateCampaignRequest) (*Campaign, *client.APIResponse, error)
	GetCampaign(ctx context.Context, id string) (*Campaign, *client.APIResponse, error)
	ListCampaigns(ctx context.Context, opts interface{}) (*ListCampaignsResponse, *client.APIResponse, error)
	GetCampaignLinks(ctx context.Context, id string) (*CampaignLinksResponse, *client.APIResponse, error)
	// CampaignLinks is a convenience helper that returns the list of links
	// for a campaign; if messageID is non-nil the result is filtered to links
	// where the link's message matches the provided messageID.
	CampaignLinks(ctx context.Context, id string, messageID *string) ([]CampaignLink, *client.APIResponse, error)
	// EditCampaign updates a campaign via PUT /campaigns/{id}/edit. The req
	// payload should match the API's expected body (see docs). On success it
	// returns the updated *Campaign and raw *client.APIResponse.
	EditCampaign(ctx context.Context, id string, req *EditCampaignRequest) (*Campaign, *client.APIResponse, error)
	// DuplicateCampaign creates a copy of a campaign via POST /campaigns/{id}/copy
	// and returns the API's acknowledgement including the new campaign ID.
	DuplicateCampaign(ctx context.Context, id string) (*DuplicateCampaignResponse, *client.APIResponse, error)
}
