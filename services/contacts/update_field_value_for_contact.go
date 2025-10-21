package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// UpdateFieldValueForContact updates a custom field value for a contact.
//
// Historically the ActiveCampaign v3 API expects the creation payload to be
// wrapped in a top-level `fieldValue` key, for example:
//
//	{
//	  "fieldValue": {
//	    "contact": 5,
//	    "field": 13,
//	    "value": "Some value"
//	  }
//	}
//
// Some callers previously sent the inner object directly which some accounts
// accept while others reject with "Field id not valid" (HTTP 403). To make the
// client consistently produce the expected shape this method wraps the
// provided `FieldValuePayload` in the required envelope before sending.
//
// The method sends POST /fieldValues and returns the typed FieldValueResponse
// on success along with the low-level *client.APIResponse for inspection.
func (s *RealService) UpdateFieldValueForContact(ctx context.Context, req *FieldValuePayload) (*FieldValueResponse, *client.APIResponse, error) {
	out := &FieldValueResponse{}
	// Envelope expected by the API
	envelope := struct {
		FieldValue *FieldValuePayload `json:"fieldValue"`
	}{FieldValue: req}

	apiResp, err := s.client.Do(ctx, http.MethodPost, "fieldValues", envelope, out)
	return out, apiResp, err
}
