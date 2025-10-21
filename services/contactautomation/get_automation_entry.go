package contactautomation

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetAutomationEntry retrieves a single automation entry by ID.
// GET /contactAutomations/{id}
func (s *service) GetAutomationEntry(ctx context.Context, id string) (*ContactAutomationResponse, *client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#get-an-automation-a-contact-is-in")
	}
	var out ContactAutomationResponse
	path := "contactAutomations/" + id
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, nil, &out)
	return &out, apiResp, err
}
