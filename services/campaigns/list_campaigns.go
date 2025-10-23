package campaigns

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// ListCampaigns fetches a list of campaigns. It delegates to the Doer and
// returns a typed *ListCampaignsResponse when successful. Callers may
// type-assert the returned interface{} to *ListCampaignsResponse.
func (s *service) ListCampaigns(ctx context.Context, opts interface{}) (*ListCampaignsResponse, *client.APIResponse, error) {
	var out ListCampaignsResponse
	apiResp, err := s.client.Do(ctx, http.MethodGet, "campaigns", opts, &out)
	if err != nil {
		return nil, apiResp, err
	}
	return &out, apiResp, nil
}
