package groups

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// UpdateGroup updates a group.
func (s *service) UpdateGroup(ctx context.Context, id string, req *UpdateGroupRequest) (*UpdateGroupResponse, *client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, nil, fmt.Errorf("service not configured: UpdateGroup")
	}
	var out UpdateGroupResponse
	path := "groups/" + id
	apiResp, err := s.client.Do(ctx, http.MethodPut, path, req, &out)
	return &out, apiResp, err
}
