package contacts

import (
	"context"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// ContactsService defines the behaviour for interacting with contacts in ActiveCampaign.
//
// Methods should accept context.Context for cancellation and timeouts and return
// a typed result, an *client.APIResponse for low-level inspection, and an error.
//
// TODOs:
// - Add go:generate for mock generation (mockgen)
// - Implement the concrete service in impl.go

type ContactsService interface {
	Create(ctx context.Context, req *CreateContactRequest) (*CreateContactResponse, *client.APIResponse, error)
	SearchByEmail(ctx context.Context, email string) (*ContactSearchResponse, *client.APIResponse, error)
	GetContact(ctx context.Context, id string) (*CreateContactResponse, *client.APIResponse, error)
	DeleteContact(ctx context.Context, id string) (*client.APIResponse, error)
	// Update an existing contact by id (PUT /contacts/{id})
	UpdateContact(ctx context.Context, id string, req *CreateContactRequest) (*CreateContactResponse, *client.APIResponse, error)
	UpdateListStatus(ctx context.Context, req *UpdateListStatusForContactRequest) (*UpdateContactListStatusResponse, *client.APIResponse, error)
	// Additional endpoints requested:
	// Get contact data (field values)
	GetContactFieldValues(ctx context.Context, id string) (interface{}, *client.APIResponse, error)
	// Typed helper that returns contact field values as a typed ListFieldValuesResponse
	GetContactFieldValuesTyped(ctx context.Context, id string) (*ListFieldValuesResponse, *client.APIResponse, error)
	// UpdateOrCreateFieldValueForContact will ensure a contact has the given custom
	// field value set. It detects the canonical custom field id (from the
	// provided identifier which may be an id, perstag, or title), checks whether a
	// FieldValue record already exists for the contact and field, and then either
	// updates the existing FieldValue (PUT /fieldValues/{id}) or creates one
	// (POST /fieldValues). Returns the FieldValueResponse or a low-level API
	// response and error.
	UpdateOrCreateFieldValueForContact(ctx context.Context, contactID, fieldIdentifier, value string) (*FieldValueResponse, *client.APIResponse, error)
	// Get contact tags
	TagsGet(ctx context.Context, id string) (*ContactTagsResponse, *client.APIResponse, error)

	// Get contact by email (with tags)
	GetContactByEmailWithTags(ctx context.Context, email string) (interface{}, *client.APIResponse, error)

	// CreateContactWithTags creates a contact and attaches the specified tag
	// IDs. It returns the created contact response and, if one or more tag
	// attachments failed, returns the last non-nil APIResponse and error from
	// the attach calls. If creating the contact fails, the create error is
	// returned directly.
	CreateContactWithTags(ctx context.Context, req *CreateContactRequest, tagIDs []string) (*CreateContactResponse, *client.APIResponse, error)

	// Bounce logs, goals, lists, logs, deals, geo, notes, organization, tracking
	GetContactBounceLogs(ctx context.Context, id string) (interface{}, *client.APIResponse, error)
	GetContactGoals(ctx context.Context, id string) (interface{}, *client.APIResponse, error)
	GetContactLists(ctx context.Context, id string) (*ContactListsResponse, *client.APIResponse, error)
	GetContactLogs(ctx context.Context, id string) (interface{}, *client.APIResponse, error)
	GetContactDealList(ctx context.Context, id string) (interface{}, *client.APIResponse, error)
	GetContactDeals(ctx context.Context, id string) (interface{}, *client.APIResponse, error)
	GetContactGeoIPs(ctx context.Context, id string) (interface{}, *client.APIResponse, error)
	GetContactGeoIP(ctx context.Context, id, ip string) (interface{}, *client.APIResponse, error)
	GetContactNotes(ctx context.Context, id string) (interface{}, *client.APIResponse, error)
	GetContactOrganization(ctx context.Context, id string) (interface{}, *client.APIResponse, error)
	GetContactAccountContacts(ctx context.Context, id string) (interface{}, *client.APIResponse, error)
	GetContactPlusAppend(ctx context.Context, id string) (interface{}, *client.APIResponse, error)
	GetContactTrackingLogs(ctx context.Context, id string) (interface{}, *client.APIResponse, error)

	// Sync, add to list, bulk import and status
	SyncContact(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error)
	AddContactToList(ctx context.Context, req *AddContactToListPayload) (*AddContactToListResponse, *client.APIResponse, error)
	BulkImportContacts(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error)
	GetBulkImportStatus(ctx context.Context, id string) (interface{}, *client.APIResponse, error)
	ListBulkImportStatus(ctx context.Context, opts map[string]string) (interface{}, *client.APIResponse, error)

	// Custom fields & field groups
	CreateCustomField(ctx context.Context, req *FieldPayload) (*FieldResponse, *client.APIResponse, error)
	ListCustomFields(ctx context.Context) (*ListFieldsResponse, *client.APIResponse, error)
	UpdateCustomField(ctx context.Context, id string, req *FieldPayload) (*FieldResponse, *client.APIResponse, error)
	DeleteCustomField(ctx context.Context, id string) (*client.APIResponse, error)

	AddFieldOption(ctx context.Context, req *FieldOptionPayload) (*FieldOptionResponse, *client.APIResponse, error)
	ListFieldValues(ctx context.Context) (*ListFieldValuesResponse, *client.APIResponse, error)
	UpdateFieldValueForContact(ctx context.Context, req *FieldValuePayload) (*FieldValueResponse, *client.APIResponse, error)

	AddFieldToGroup(ctx context.Context, req interface{}) (*client.APIResponse, error)
	GetFieldGroup(ctx context.Context, id string) (*FieldGroupResponse, *client.APIResponse, error)
	UpdateFieldGroup(ctx context.Context, id string, req interface{}) (*client.APIResponse, error)
	DeleteFieldGroup(ctx context.Context, id string) (*client.APIResponse, error)

	// UpdateListStatusManaged helps manage list membership with an optional Force flag.
	// When Force is false, the contact will only be subscribed (status "1") if they are
	// not already explicitly Unsubscribed (status "2"). When Force is true, the contact
	// will be set to Subscribed (status "1") regardless of prior Unsubscribed state.
	UpdateListStatusManaged(ctx context.Context, req *UpdateListStatusHelperRequest) (*UpdateContactListStatusResponse, *client.APIResponse, error)

	// EnsureSubscribedToList is a convenience wrapper that subscribes a contact
	// to the given list using UpdateListStatusManaged. When force is true, the
	// subscription will proceed even if the contact previously unsubscribed.
	EnsureSubscribedToList(ctx context.Context, contactID, listID string, force bool) (*UpdateContactListStatusResponse, *client.APIResponse, error)

	// TagAdd adds a tag to an existing contact.
	TagAdd(ctx context.Context, contactID, tagID string) (*ContactTagResponse, *client.APIResponse, error)

	// TagRemove removes a tag from a contact by finding the association and deleting it.
	TagRemove(ctx context.Context, contactID, tag string) (*client.APIResponse, error)
}
