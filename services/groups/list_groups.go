package groups

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// ListGroups lists groups.
func (s *service) ListGroups(ctx context.Context, opts map[string]string) (*ListGroupsResponse, *client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, nil, fmt.Errorf("service not configured: ListGroups")
	}
	var out ListGroupsResponse
	apiResp, err := s.client.Do(ctx, http.MethodGet, "groups", opts, &out)
	return &out, apiResp, err
}
