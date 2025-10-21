package lists

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetList retrieves a single list by ID.
//
// GET /lists/{id}
func (s *service) GetList(ctx context.Context, id string) (GetListResponse, *client.APIResponse, error) {
	var zero GetListResponse
	if s == nil || s.client == nil {
		return zero, nil, fmt.Errorf("not implemented: see ActiveCampaign API docs for getting lists")
	}
	var out GetListResponse
	apiResp, err := s.client.Do(ctx, http.MethodGet, "lists/"+id, nil, &out)
	return out, apiResp, err
}
