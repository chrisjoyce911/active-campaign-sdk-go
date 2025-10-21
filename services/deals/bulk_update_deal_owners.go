package deals

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// BulkUpdateDealOwners updates owners for multiple deals in a single request.
//
// What & Why:
//
//	Bulk-assign or update deal owners to reduce repeated API calls when
//	updating many deals at once.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#bulk-update-deal-owners
//
// Parameters:
//
//	ctx: context
//	req: bulk update payload
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *RealService) BulkUpdateDealOwners(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error) {
	var out interface{}
	path := "deals/bulk/owners"
	apiResp, err := s.client.Do(ctx, http.MethodPatch, path, req, &out)
	return out, apiResp, err
}
