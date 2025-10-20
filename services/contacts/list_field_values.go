package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// ListFieldValues lists all custom field values.
// GET /fieldValues
func (s *RealService) ListFieldValues(ctx context.Context) (*ListFieldValuesResponse, *client.APIResponse, error) {
	out := &ListFieldValuesResponse{}
	apiResp, err := s.client.Do(ctx, http.MethodGet, "fieldValues", nil, out)
	return out, apiResp, err
}
