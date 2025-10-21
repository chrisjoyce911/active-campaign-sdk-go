package contactautomation

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// ListContactAutomations lists automations for a given contact.
// GET /contacts/{contactID}/contactAutomations
func (s *service) ListContactAutomations(ctx context.Context, contactID string) (*ListContactAutomationsResponse, *client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#get-contact-automations")
	}
	var out ListContactAutomationsResponse
	path := "contacts/" + contactID + "/contactAutomations"
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, nil, &out)
	return &out, apiResp, err
}
