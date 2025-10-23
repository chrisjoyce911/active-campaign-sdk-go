package campaigns

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// CreateCampaign creates a campaign by POSTing the provided payload to the API.
// The API returns the created campaign; when successful this method returns a
// value of type *Campaign (as interface{}). Callers can type-assert the
// response: out.(*Campaign).
func (s *service) CreateCampaign(ctx context.Context, req *CreateCampaignRequest) (*Campaign, *client.APIResponse, error) {
	// API returns wrapper {"campaign": {...}}
	var wrapper struct {
		Campaign Campaign `json:"campaign"`
	}
	apiResp, err := s.client.Do(ctx, http.MethodPost, "campaigns", req, &wrapper)
	if err != nil {
		return nil, apiResp, err
	}
	return &wrapper.Campaign, apiResp, nil
}
