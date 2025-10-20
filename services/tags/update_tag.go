package tags

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// UpdateTag updates a tag by ID.
//
// It sends a PUT request to /tags/{id} with the provided payload and
// returns the updated TagResponse and API metadata. If the service
// isn't initialized the method returns a not-implemented error to keep
// behaviour consistent with zero-value receivers used elsewhere.
func (s *service) UpdateTag(ctx context.Context, id string, req *CreateOrUpdateTagRequest) (*TagResponse, *client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#update-tag")
	}
	var out TagResponse
	path := "tags/" + id
	apiResp, err := s.client.Do(ctx, http.MethodPut, path, req, &out)
	return &out, apiResp, err
}
