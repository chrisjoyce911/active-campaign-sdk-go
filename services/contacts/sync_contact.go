package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// SyncContact creates or updates a contact by email address ("upsert") in a
// single call, including any inline custom field values.
//
// What & Why:
//
//	POST /api/3/contact/sync matches on the contact's email address: if no
//	contact exists it is created, otherwise the existing contact is updated.
//	FieldValues on the request are written in the same call — an empty Value
//	clears the field. Fields omitted from the request keep their current
//	values. This replaces a search + create/update + per-field write sequence
//	with one request.
//
//	The response contact's CDate/UDate can distinguish outcome: on a freshly
//	created contact they are equal; an update advances UDate only.
//
//	Note the path is contact/sync (singular). POSTing to contacts/sync hits
//	the plain create endpoint, which returns 422 "duplicate" for any existing
//	email — an earlier version of this method had that bug.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#sync-contact-data
//
// Parameters:
//
//	ctx: context for cancellation/timeouts
//	req: contact payload; FieldValues may be set inline
//
// Returns:
//
//	(*CreateContactResponse, *client.APIResponse, error)
func (s *RealService) SyncContact(ctx context.Context, req *CreateContactRequest) (*CreateContactResponse, *client.APIResponse, error) {
	var out CreateContactResponse
	path := "contact/sync"
	apiResp, err := s.client.Do(ctx, http.MethodPost, path, req, &out)
	return &out, apiResp, err
}
