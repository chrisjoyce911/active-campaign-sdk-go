package campaigns

import (
	"context"
	"net/http"
	"net/url"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetCampaign retrieves a campaign by ID from /campaigns/{id} and decodes
// the JSON into a Campaign value. On success the returned *Campaign is
// populated from the API response and the raw *client.APIResponse is
// also returned for callers that need access to headers/status/body.
func (s *service) GetCampaign(ctx context.Context, id string) (*Campaign, *client.APIResponse, error) {
	// Ensure the id is safely encoded for inclusion in a path segment.
	path := "campaigns/" + url.PathEscape(id)

	// The API returns a wrapper object: { "campaign": { ... } }
	var wrapper struct {
		Campaign Campaign `json:"campaign"`
	}
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, nil, &wrapper)
	if err != nil {
		return nil, apiResp, err
	}
	return &wrapper.Campaign, apiResp, nil
}
