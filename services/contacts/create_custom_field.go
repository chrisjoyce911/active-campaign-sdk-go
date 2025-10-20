package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// CreateCustomField creates a custom field for contacts.
// https://developers.activecampaign.com/reference/create-a-contact-custom-field
// POST /fields
func (s *RealService) CreateCustomField(ctx context.Context, req *FieldPayload) (*FieldResponse, *client.APIResponse, error) {
	out := &FieldResponse{}
	apiResp, err := s.client.Do(ctx, http.MethodPost, "fields", req, out)
	return out, apiResp, err
}
