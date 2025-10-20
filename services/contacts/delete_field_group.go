package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// DeleteFieldGroup deletes a field group by id.
// POST /fieldGroups/delete
func (s *RealService) DeleteFieldGroup(ctx context.Context, id string) (*client.APIResponse, error) {
	// ActiveCampaign sometimes uses POST for delete group endpoints; adapt accordingly.
	path := "fieldGroups/delete"
	payload := map[string]string{"id": id}
	apiResp, err := s.client.Do(ctx, http.MethodPost, path, payload, nil)
	return apiResp, err
}
