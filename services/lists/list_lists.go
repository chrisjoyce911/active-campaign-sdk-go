package lists

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// ListLists retrieves multiple lists.
//
// GET /lists
// Supports optional query params such as filters[name], limit, orders[name], etc.
func (s *service) ListLists(ctx context.Context, opts map[string]string) (ListsResponse, *client.APIResponse, error) {
	var zero ListsResponse
	if s == nil || s.client == nil {
		return zero, nil, fmt.Errorf("not implemented: see ActiveCampaign API docs for listing lists")
	}
	var out ListsResponse
	apiResp, err := s.client.Do(ctx, http.MethodGet, "lists", opts, &out)
	return out, apiResp, err
}
