package contacts

import (
	"context"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// EnsureSubscribedToList subscribes a contact to a list, with optional force.
// When force is true, the contact will be subscribed (status "1") even if they
// were previously Unsubscribed (status "2"). When false, Unsubscribed contacts
// are left unchanged.
func (s *RealService) EnsureSubscribedToList(ctx context.Context, contactID, listID string, force bool) (*UpdateContactListStatusResponse, *client.APIResponse, error) {
	req := &UpdateListStatusHelperRequest{ContactList: &ContactList{Contact: contactID, List: listID, Status: 1}, Force: force}
	return s.UpdateListStatusManaged(ctx, req)
}
