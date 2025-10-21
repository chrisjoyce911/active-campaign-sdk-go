package deals

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// CreateDealNote creates a note attached to a deal.
//
// What & Why:
//
//	Add a textual note to a deal for audit, comments, or activity tracking.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#create-a-deal-note
//
// Parameters:
//
//	ctx: context
//	dealID: ID of the deal to attach the note to
//	req: note payload
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *RealService) CreateDealNote(ctx context.Context, dealID string, req interface{}) (interface{}, *client.APIResponse, error) {
	var out interface{}
	path := "deals/" + dealID + "/notes"
	apiResp, err := s.client.Do(ctx, http.MethodPost, path, req, &out)
	return out, apiResp, err
}
