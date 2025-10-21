package lists

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// CreateList creates a new list.
//
// POST /lists
// Request body is CreateListRequest which wraps the List under the "list" key.
func (s *service) CreateList(ctx context.Context, req CreateListRequest) (CreateListResponse, *client.APIResponse, error) {
	var zero CreateListResponse
	if s == nil || s.client == nil {
		return zero, nil, fmt.Errorf("not implemented: see ActiveCampaign API docs for creating lists")
	}
	var out CreateListResponse
	apiResp, err := s.client.Do(ctx, http.MethodPost, "lists", req, &out)
	return out, apiResp, err
}
