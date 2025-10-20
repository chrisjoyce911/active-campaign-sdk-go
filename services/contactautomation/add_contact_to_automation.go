package contactautomation

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// AddContactToAutomation adds a contact to an automation.
// POST /contactAutomations
func (s *service) AddContactToAutomation(ctx context.Context, req *CreateContactAutomationRequest) (*ContactAutomationResponse, *client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#create-new-contactautomation")
	}
	var out ContactAutomationResponse
	apiResp, err := s.client.Do(ctx, http.MethodPost, "contactAutomations", req, &out)
	return &out, apiResp, err
}
