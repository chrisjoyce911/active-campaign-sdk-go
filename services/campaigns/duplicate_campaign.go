package campaigns

import (
	"context"
	"net/http"
	"net/url"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// DuplicateCampaign creates a copy of the campaign identified by id. The API
// responds with a simple acknowledgement which includes the new campaign id.
func (s *service) DuplicateCampaign(ctx context.Context, id string) (*DuplicateCampaignResponse, *client.APIResponse, error) {
	path := "campaigns/" + url.PathEscape(id) + "/copy"
	var out DuplicateCampaignResponse
	apiResp, err := s.client.Do(ctx, http.MethodPost, path, nil, &out)
	if err != nil {
		return nil, apiResp, err
	}
	return &out, apiResp, nil
}
