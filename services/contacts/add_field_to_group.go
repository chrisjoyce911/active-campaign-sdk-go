package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// AddFieldToGroup associates a custom field with a field group.
// POST /fieldGroups/fields
func (s *RealService) AddFieldToGroup(ctx context.Context, req interface{}) (*client.APIResponse, error) {
	apiResp, err := s.client.Do(ctx, http.MethodPost, "fieldGroups/fields", req, nil)
	return apiResp, err
}
