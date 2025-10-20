package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// ListCustomFields lists all custom fields for contacts.
// GET /fields
func (s *RealService) ListCustomFields(ctx context.Context) (*ListFieldsResponse, *client.APIResponse, error) {
	out := &ListFieldsResponse{}
	apiResp, err := s.client.Do(ctx, http.MethodGet, "fields", nil, out)
	return out, apiResp, err
}
