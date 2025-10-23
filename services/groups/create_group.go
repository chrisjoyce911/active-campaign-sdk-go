package groups

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// CreateGroup creates a group resource.
func (s *service) CreateGroup(ctx context.Context, req *CreateGroupRequest) (*CreateGroupResponse, *client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, nil, fmt.Errorf("service not configured: CreateGroup")
	}
	var out CreateGroupResponse
	apiResp, err := s.client.Do(ctx, http.MethodPost, "groups", req, &out)
	return &out, apiResp, err
}
