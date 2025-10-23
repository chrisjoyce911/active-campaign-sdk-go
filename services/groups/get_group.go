package groups

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetGroup retrieves a group by ID.
func (s *service) GetGroup(ctx context.Context, id string) (*GetGroupResponse, *client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, nil, fmt.Errorf("service not configured: GetGroup")
	}
	var out GetGroupResponse
	path := "groups/" + id
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, nil, &out)
	return &out, apiResp, err
}
