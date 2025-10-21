package lists

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// DeleteList deletes a single list by ID.
//
// DELETE /lists/{id}
func (s *service) DeleteList(ctx context.Context, id string) (*client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, fmt.Errorf("not implemented: see ActiveCampaign API docs for deleting lists")
	}
	apiResp, err := s.client.Do(ctx, http.MethodDelete, "lists/"+id, nil, nil)
	return apiResp, err
}
