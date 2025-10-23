package tags

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetTag retrieves a tag by ID.
//
// It sends a GET request to /tags/{id} and returns the TagResponse
// containing the tag data and API metadata. If the service or client
// is not configured, a not-implemented error is returned to preserve
// existing zero-value receiver behaviour in tests.
func (s *service) GetTag(ctx context.Context, id string) (*TagResponse, *client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, nil, fmt.Errorf("service not configured: GetTag")
	}
	var out TagResponse
	path := "tags/" + id
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, nil, &out)
	return &out, apiResp, err
}
