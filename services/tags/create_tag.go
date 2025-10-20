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
// and API metadata. If the service is nil or not wired with a client
// a not-implemented error is returned so callers using zero-value
// receivers continue to behave as before.
func (s *service) CreateTag(ctx context.Context, req *CreateOrUpdateTagRequest) (*TagResponse, *client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#create-tag")
	}
	var out TagResponse
	apiResp, err := s.client.Do(ctx, http.MethodPost, "tags", req, &out)
	return &out, apiResp, err
}
