package contactautomation

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetCounts retrieves automation entry counts.
// GET /contactAutomations/counts
func (s *service) GetCounts(ctx context.Context) (*AutomationCountsResponse, *client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#get-contact-automation-entry-counts")
	}
	var out AutomationCountsResponse
	apiResp, err := s.client.Do(ctx, http.MethodGet, "contactAutomations/counts", nil, &out)
	return &out, apiResp, err
}
