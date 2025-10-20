package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// UpdateFieldValueForContact updates a custom field value for a contact.
// POST /fieldValues
func (s *RealService) UpdateFieldValueForContact(ctx context.Context, req *FieldValuePayload) (*FieldValueResponse, *client.APIResponse, error) {
	out := &FieldValueResponse{}
	apiResp, err := s.client.Do(ctx, http.MethodPost, "fieldValues", req, out)
	return out, apiResp, err
}
