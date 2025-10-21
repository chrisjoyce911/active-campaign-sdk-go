package deals

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// UpdateDealNote updates an existing note on a deal.
//
// What & Why:
//
//	Modify the contents or metadata of a previously created deal note.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#update-a-deal-note
//
// Parameters:
//
//	ctx: context
//	dealID: ID of the deal
//	noteID: ID of the note
//	req: updated note payload
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *RealService) UpdateDealNote(ctx context.Context, dealID, noteID string, req interface{}) (interface{}, *client.APIResponse, error) {
	var out interface{}
	path := "deals/" + dealID + "/notes/" + noteID
	apiResp, err := s.client.Do(ctx, http.MethodPut, path, req, &out)
	return out, apiResp, err
}
