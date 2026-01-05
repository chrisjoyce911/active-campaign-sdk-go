package contacts

import (
	"context"
	"errors"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// UpdateListStatusHelperRequest describes the intent to subscribe a contact to a list
// with optional Force semantics.
//
// If Force is false, the helper will not change a contact that is already
// explicitly Unsubscribed (status "2"). If Force is true, the helper will
// set the contact to Subscribed (status "1") regardless of previous state.
type UpdateListStatusHelperRequest struct {
	ContactList *ContactList
	Force       bool
}

// UpdateListStatusManaged ensures the desired list status using the rules above.
func (s *RealService) UpdateListStatusManaged(ctx context.Context, req *UpdateListStatusHelperRequest) (*UpdateContactListStatusResponse, *client.APIResponse, error) {
	if req == nil || req.ContactList == nil {
		return nil, nil, errors.New("UpdateListStatusManaged: nil request or contact list")
	}
	cl := req.ContactList
	if cl.Contact == "" || cl.List == "" {
		return nil, nil, errors.New("UpdateListStatusManaged: contact and list are required")
	}

	// Fetch current membership state
	listsResp, apiResp, err := s.GetContactLists(ctx, cl.Contact)
	if err != nil {
		// Bubble up GET error (including 404 for unknown contact)
		return nil, apiResp, err
	}

	// Find existing membership for the target list
	var existing *ContactList
	for i := range listsResp.ContactListsOrEmpty() {
		item := listsResp.ContactListsOrEmpty()[i]
		if item.List == cl.List {
			existing = &item
			break
		}
	}

	// Desired subscribed state (default to subscribe if unset)
	desiredStatus := cl.Status
	if desiredStatus == 0 {
		desiredStatus = 1
	}

	// Decision matrix
	if existing != nil {
		if existing.Status == 2 && !req.Force {
			return nil, apiResp, nil // skip when unsubscribed and not forced
		}
		if existing.Status == desiredStatus {
			return nil, apiResp, nil // already desired
		}
	}

	// Proceed to create/update membership via underlying endpoint
	payload := &UpdateListStatusForContactRequest{ContactList: &ContactList{
		Contact: cl.Contact,
		List:    cl.List,
		Status:  desiredStatus,
	}}
	return s.UpdateListStatus(ctx, payload)
}
