package lists

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// CreateListGroup associates a list with a group (list group permission).
//
// POST /listGroups
// Request body should be CreateListGroupRequest with a top-level "listGroup" object.
func (s *service) CreateListGroup(ctx context.Context, req CreateListGroupRequest) (CreateListGroupResponse, *client.APIResponse, error) {
	var zero CreateListGroupResponse
	if s == nil || s.client == nil {
		return zero, nil, fmt.Errorf("not implemented: see ActiveCampaign API docs for creating list group permissions")
	}
	var out CreateListGroupResponse
	apiResp, err := s.client.Do(ctx, http.MethodPost, "listGroups", req, &out)
	return out, apiResp, err
}
