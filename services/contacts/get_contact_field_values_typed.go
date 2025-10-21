package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetContactFieldValuesTyped returns custom field values for a contact as a typed response.
// GET /contacts/{id}/fieldValues
func (s *RealService) GetContactFieldValuesTyped(ctx context.Context, id string) (*ListFieldValuesResponse, *client.APIResponse, error) {
	out := &ListFieldValuesResponse{}
	path := "contacts/" + id + "/fieldValues"
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, nil, out)
	return out, apiResp, err
}
