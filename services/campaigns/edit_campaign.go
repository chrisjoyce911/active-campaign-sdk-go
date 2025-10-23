package campaigns

import (
	"context"
	"net/http"
	"net/url"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// EditCampaign updates a campaign via PUT /campaigns/{id}/edit. The req
// payload should match the API's expected body. The API responds with a
// campaign object (wrapped under "campaign"); this method decodes that
// wrapper and returns the inner *Campaign.
func (s *service) EditCampaign(ctx context.Context, id string, req *EditCampaignRequest) (*Campaign, *client.APIResponse, error) {
	path := "campaigns/" + url.PathEscape(id) + "/edit"

	var wrapper struct {
		Campaign Campaign `json:"campaign"`
	}
	apiResp, err := s.client.Do(ctx, http.MethodPut, path, req, &wrapper)
	if err != nil {
		return nil, apiResp, err
	}
	return &wrapper.Campaign, apiResp, nil
}
