package contactautomation

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// RemoveContactFromAutomation removes a contact from an automation.
// DELETE /contactAutomations/{id}
func (s *service) RemoveContactFromAutomation(ctx context.Context, id string) (*client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#remove-a-contact-from-an-automation")
	}
	path := "contactAutomations/" + id
	apiResp, err := s.client.Do(ctx, http.MethodDelete, path, nil, nil)
	return apiResp, err
}
