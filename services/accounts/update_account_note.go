package accounts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// UpdateAccountNote updates an existing account note by ID.
//
// What & Why:
//
//	Updates the note text or metadata for a previously created account note.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#update-account-note
//
// Parameters:
//
//	ctx: context
//	noteID: the note ID
//	req: update payload
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *service) UpdateAccountNote(ctx context.Context, noteID string, req interface{}) (interface{}, *client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#update-account-note")
	}
	var out interface{}
	path := "accounts/notes/" + noteID
	apiResp, err := s.client.Do(ctx, http.MethodPut, path, req, &out)
	return out, apiResp, err
}
