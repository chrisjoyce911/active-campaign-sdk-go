package accounts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// CreateAccountNote creates a note associated with an account.
//
// What & Why:
//
//	Adds a note to an account for audit, comments, or activity tracking.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#create-account-note
//
// Parameters:
//
//	ctx: context
//	accountID: ID of the account
//	req: note payload
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *service) CreateAccountNote(ctx context.Context, accountID string, req *AccountNoteRequest) (*AccountNoteResponse, *client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#create-account-note")
	}
	var out AccountNoteResponse
	path := "accounts/" + accountID + "/notes"
	apiResp, err := s.client.Do(ctx, http.MethodPost, path, req, &out)
	if err != nil {
		return nil, apiResp, err
	}
	return &out, apiResp, nil
}
