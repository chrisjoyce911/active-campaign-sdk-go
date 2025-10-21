package deals

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// DeleteDeal deletes a deal.
//
// What & Why:
//
//	Permanently remove a deal from the account.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#delete-a-deal
//
// Parameters:
//
//	ctx: context
//	id: deal ID
//
// Returns:
//
//	(*client.APIResponse, error)
func (s *RealService) DeleteDeal(ctx context.Context, id string) (*client.APIResponse, error) {
	path := "deals/" + id
	apiResp, err := s.client.Do(ctx, http.MethodDelete, path, nil, nil)
	return apiResp, err
}
