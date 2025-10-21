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
// the HTTP status and any error information. When called against a
// zero-value receiver a not-implemented error is returned to preserve
// previous behaviour during migration.
func (s *service) DeleteTag(ctx context.Context, id string) (*client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#delete-tag")
	}
	path := "tags/" + id
	apiResp, err := s.client.Do(ctx, http.MethodDelete, path, nil, nil)
	return apiResp, err
}
