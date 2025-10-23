package tags

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// DeleteTag deletes a tag by ID.
//
// It sends a DELETE request to /tags/{id}. The APIResponse contains
// the HTTP status and any error information. When the service is not
// configured this method returns an error indicating the service is
// not configured.
//
// API reference: https://developers.activecampaign.com/reference#delete-tag
func (s *service) DeleteTag(ctx context.Context, id string) (*client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, fmt.Errorf("service not configured: DeleteTag")
	}
	path := "tags/" + id
	apiResp, err := s.client.Do(ctx, http.MethodDelete, path, nil, nil)
	return apiResp, err
}
