package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetContactNotes returns notes for a contact.
//
// What & Why:
//
//	Provide notes attached to a contact for CRM/agent context.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#get-contact-notes
//
// Parameters:
//
//	ctx: context
//	id: contact id
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *RealService) GetContactNotes(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	var out interface{}
	path := "contacts/" + id + "/notes"
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, nil, &out)
	return out, apiResp, err
}
