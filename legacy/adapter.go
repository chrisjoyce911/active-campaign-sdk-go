package legacy

import (
	"context"
	"fmt"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// Legacy adapters bridge the old root-level API to the new client/services.
// These functions are thin placeholders that preserve the old API surface so
// existing call sites compile. They intentionally return a standardized
// not-implemented error. Replace with real implementations during migration.

// CreateContact mirrors the old root-level Contacts.Create signature.
func CreateContact(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#create-contact")
}

// GetContact mirrors the old root-level Contacts.GetContact signature.
func GetContact(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#get-contact")
}

// DeleteContact mirrors the old root-level Contacts.Delete signature.
func DeleteContact(ctx context.Context, id string) (*client.APIResponse, error) {
	return nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#delete-contact")
}

// SearchContacts mirrors the old root-level search by email.
func SearchContacts(ctx context.Context, email string) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#list-all-contacts")
}

// Add any additional legacy adapter stubs here (Lists, Tags, CustomFields, etc.)
