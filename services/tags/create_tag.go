package tags

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// CreateTag creates a new tag resource.
//
// It sends a POST request to the /tags endpoint with the provided
// CreateOrUpdateTagRequest payload and returns the created TagResponse
// and API metadata. When the service is not configured this method
// returns an error indicating the service is not configured.
//
// API reference: https://developers.activecampaign.com/reference#create-tag
func (s *service) CreateTag(ctx context.Context, req *CreateOrUpdateTagRequest) (*TagResponse, *client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, nil, fmt.Errorf("service not configured: CreateTag")
	}
	var out TagResponse
	apiResp, err := s.client.Do(ctx, http.MethodPost, "tags", req, &out)
	return &out, apiResp, err
}
