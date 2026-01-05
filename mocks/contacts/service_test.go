package contactsmock

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/contacts"
)

func TestService_SearchByEmailFuncCalled(t *testing.T) {
	called := false
	m := &Service{SearchByEmailFunc: func(ctx context.Context, email string) (*contacts.ContactSearchResponse, *client.APIResponse, error) {
		called = true
		return &contacts.ContactSearchResponse{Contacts: []contacts.Contact{{ID: "c1"}}}, &client.APIResponse{StatusCode: 200}, nil
	}}
	resp, api, err := m.SearchByEmail(context.Background(), "a@b.com")
	if err != nil || api == nil || resp == nil || len(resp.Contacts) != 1 || !called {
		t.Fatalf("unexpected result: resp=%v api=%v err=%v called=%v", resp, api, err, called)
	}
}

func TestService_Defaults(t *testing.T) {
	m := &Service{}
	// Defaults should return non-nil values where applicable
	resp, api, err := m.SearchByEmail(context.Background(), "a@b.com")
	if err != nil || resp == nil || api == nil {
		t.Fatalf("expected non-nil defaults, got resp=%v api=%v err=%v", resp, api, err)
	}
	fv, fvAPI, fvErr := m.UpdateOrCreateFieldValueForContact(context.Background(), "c1", "cf1", "val")
	if fv == nil || fvAPI == nil || fvErr != nil {
		t.Fatalf("expected non-nil defaults for UpdateOrCreateFieldValueForContact")
	}
	upAPI, upErr := m.UpdateFieldGroup(context.Background(), "g1", nil)
	if upErr != nil || upAPI == nil {
		t.Fatalf("expected non-nil defaults for UpdateFieldGroup")
	}
}

func TestService_Defaults_Many(t *testing.T) {
	m := &Service{}
	ctx := context.Background()
	// Call a broad set of default-returning methods to exercise code paths
	_, _, _ = m.GetContact(ctx, "c1")
	_, _ = m.DeleteContact(ctx, "c1")
	_, _, _ = m.UpdateContact(ctx, "c1", &contacts.CreateContactRequest{})
	_, _, _ = m.UpdateListStatus(ctx, &contacts.UpdateListStatusForContactRequest{})
	_, _, _ = m.GetContactFieldValues(ctx, "c1")
	_, _, _ = m.GetContactFieldValuesTyped(ctx, "c1")
	_, _, _ = m.TagsGet(ctx, "c1")
	_, _, _ = m.GetContactByEmailWithTags(ctx, "a@b.com")
	_, _, _ = m.CreateContactWithTags(ctx, &contacts.CreateContactRequest{}, []string{"t1"})
	_, _, _ = m.GetContactBounceLogs(ctx, "c1")
	_, _, _ = m.GetContactGoals(ctx, "c1")
	_, _, _ = m.GetContactLists(ctx, "c1")
	_, _, _ = m.GetContactLogs(ctx, "c1")
	_, _, _ = m.GetContactDealList(ctx, "c1")
	_, _, _ = m.GetContactDeals(ctx, "c1")
	_, _, _ = m.GetContactGeoIPs(ctx, "c1")
	_, _, _ = m.GetContactGeoIP(ctx, "c1", "127.0.0.1")
	_, _, _ = m.GetContactNotes(ctx, "c1")
	_, _, _ = m.GetContactOrganization(ctx, "c1")
	_, _, _ = m.GetContactAccountContacts(ctx, "c1")
	_, _, _ = m.GetContactPlusAppend(ctx, "c1")
	_, _, _ = m.GetContactTrackingLogs(ctx, "c1")
	_, _, _ = m.SyncContact(ctx, nil)
	_, _, _ = m.AddContactToList(ctx, nil)
	_, _, _ = m.BulkImportContacts(ctx, nil)
	_, _, _ = m.GetBulkImportStatus(ctx, "id1")
	_, _, _ = m.ListBulkImportStatus(ctx, nil)
	_, _, _ = m.CreateCustomField(ctx, &contacts.FieldPayload{})
	_, _, _ = m.ListCustomFields(ctx)
	_, _, _ = m.UpdateCustomField(ctx, "f1", &contacts.FieldPayload{})
	_, _ = m.DeleteCustomField(ctx, "f1")
	_, _, _ = m.AddFieldOption(ctx, &contacts.FieldOptionPayload{})
	_, _, _ = m.ListFieldValues(ctx)
	_, _, _ = m.UpdateFieldValueForContact(ctx, &contacts.FieldValuePayload{})
	_, _ = m.AddFieldToGroup(ctx, nil)
	_, _, _ = m.GetFieldGroup(ctx, "g1")
	_, _ = m.UpdateFieldGroup(ctx, "g1", nil)
	_, _ = m.DeleteFieldGroup(ctx, "g1")
	_, _, _ = m.UpdateListStatusManaged(ctx, &contacts.UpdateListStatusHelperRequest{})
	_, _, _ = m.EnsureSubscribedToList(ctx, "c1", "l1", true)
}

func TestService_FunctionsSubset(t *testing.T) {
	called := 0
	m := &Service{
		CreateFunc: func(ctx context.Context, req *contacts.CreateContactRequest) (*contacts.CreateContactResponse, *client.APIResponse, error) {
			called++
			return &contacts.CreateContactResponse{}, &client.APIResponse{StatusCode: 200}, nil
		},
		UpdateListStatusManagedFunc: func(ctx context.Context, req *contacts.UpdateListStatusHelperRequest) (*contacts.UpdateContactListStatusResponse, *client.APIResponse, error) {
			called++
			return &contacts.UpdateContactListStatusResponse{}, &client.APIResponse{StatusCode: 200}, nil
		},
		EnsureSubscribedToListFunc: func(ctx context.Context, contactID, listID string, force bool) (*contacts.UpdateContactListStatusResponse, *client.APIResponse, error) {
			called++
			return &contacts.UpdateContactListStatusResponse{}, &client.APIResponse{StatusCode: 200}, nil
		},
	}
	_, _, _ = m.Create(context.Background(), &contacts.CreateContactRequest{})
	_, _, _ = m.UpdateListStatusManaged(context.Background(), &contacts.UpdateListStatusHelperRequest{})
	_, _, _ = m.EnsureSubscribedToList(context.Background(), "c1", "l1", true)
	if called != 3 {
		t.Fatalf("expected 3 calls, got %d", called)
	}
}

func TestService_FunctionsMany(t *testing.T) {
	called := 0
	m := &Service{
		SearchByEmailFunc: func(ctx context.Context, email string) (*contacts.ContactSearchResponse, *client.APIResponse, error) {
			called++
			return &contacts.ContactSearchResponse{}, &client.APIResponse{StatusCode: 200}, nil
		},
		GetContactFunc: func(ctx context.Context, id string) (*contacts.CreateContactResponse, *client.APIResponse, error) {
			called++
			return &contacts.CreateContactResponse{}, &client.APIResponse{StatusCode: 200}, nil
		},
		DeleteContactFunc: func(ctx context.Context, id string) (*client.APIResponse, error) {
			called++
			return &client.APIResponse{StatusCode: 204}, nil
		},
		UpdateContactFunc: func(ctx context.Context, id string, req *contacts.CreateContactRequest) (*contacts.CreateContactResponse, *client.APIResponse, error) {
			called++
			return &contacts.CreateContactResponse{}, &client.APIResponse{StatusCode: 200}, nil
		},
		TagsGetFunc: func(ctx context.Context, id string) (*contacts.ContactTagsResponse, *client.APIResponse, error) {
			called++
			return &contacts.ContactTagsResponse{}, &client.APIResponse{StatusCode: 200}, nil
		},
		GetContactByEmailWithTagsFunc: func(ctx context.Context, email string) (interface{}, *client.APIResponse, error) {
			called++
			return nil, &client.APIResponse{StatusCode: 200}, nil
		},
		CreateContactWithTagsFunc: func(ctx context.Context, req *contacts.CreateContactRequest, tagIDs []string) (*contacts.CreateContactResponse, *client.APIResponse, error) {
			called++
			return &contacts.CreateContactResponse{}, &client.APIResponse{StatusCode: 200}, nil
		},
		GetContactBounceLogsFunc: func(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
			called++
			return nil, &client.APIResponse{StatusCode: 200}, nil
		},
		GetContactGoalsFunc: func(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
			called++
			return nil, &client.APIResponse{StatusCode: 200}, nil
		},
		GetContactListsFunc: func(ctx context.Context, id string) (*contacts.ContactListsResponse, *client.APIResponse, error) {
			called++
			return &contacts.ContactListsResponse{}, &client.APIResponse{StatusCode: 200}, nil
		},
		GetContactLogsFunc: func(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
			called++
			return nil, &client.APIResponse{StatusCode: 200}, nil
		},
		GetContactDealListFunc: func(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
			called++
			return nil, &client.APIResponse{StatusCode: 200}, nil
		},
		GetContactDealsFunc: func(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
			called++
			return nil, &client.APIResponse{StatusCode: 200}, nil
		},
		GetContactGeoIPsFunc: func(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
			called++
			return nil, &client.APIResponse{StatusCode: 200}, nil
		},
		GetContactGeoIPFunc: func(ctx context.Context, id, ip string) (interface{}, *client.APIResponse, error) {
			called++
			return nil, &client.APIResponse{StatusCode: 200}, nil
		},
		GetContactNotesFunc: func(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
			called++
			return nil, &client.APIResponse{StatusCode: 200}, nil
		},
		GetContactOrganizationFunc: func(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
			called++
			return nil, &client.APIResponse{StatusCode: 200}, nil
		},
		GetContactAccountContactsFunc: func(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
			called++
			return nil, &client.APIResponse{StatusCode: 200}, nil
		},
		GetContactPlusAppendFunc: func(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
			called++
			return nil, &client.APIResponse{StatusCode: 200}, nil
		},
		GetContactTrackingLogsFunc: func(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
			called++
			return nil, &client.APIResponse{StatusCode: 200}, nil
		},
		SyncContactFunc: func(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error) {
			called++
			return nil, &client.APIResponse{StatusCode: 200}, nil
		},
		AddContactToListFunc: func(ctx context.Context, req *contacts.AddContactToListPayload) (*contacts.AddContactToListResponse, *client.APIResponse, error) {
			called++
			return &contacts.AddContactToListResponse{}, &client.APIResponse{StatusCode: 200}, nil
		},
		BulkImportContactsFunc: func(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error) {
			called++
			return nil, &client.APIResponse{StatusCode: 200}, nil
		},
		GetBulkImportStatusFunc: func(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
			called++
			return nil, &client.APIResponse{StatusCode: 200}, nil
		},
		ListBulkImportStatusFunc: func(ctx context.Context, opts map[string]string) (interface{}, *client.APIResponse, error) {
			called++
			return nil, &client.APIResponse{StatusCode: 200}, nil
		},
		CreateCustomFieldFunc: func(ctx context.Context, req *contacts.FieldPayload) (*contacts.FieldResponse, *client.APIResponse, error) {
			called++
			return &contacts.FieldResponse{}, &client.APIResponse{StatusCode: 200}, nil
		},
		ListCustomFieldsFunc: func(ctx context.Context) (*contacts.ListFieldsResponse, *client.APIResponse, error) {
			called++
			return &contacts.ListFieldsResponse{}, &client.APIResponse{StatusCode: 200}, nil
		},
		AddFieldOptionFunc: func(ctx context.Context, req *contacts.FieldOptionPayload) (*contacts.FieldOptionResponse, *client.APIResponse, error) {
			called++
			return &contacts.FieldOptionResponse{}, &client.APIResponse{StatusCode: 200}, nil
		},
		ListFieldValuesFunc: func(ctx context.Context) (*contacts.ListFieldValuesResponse, *client.APIResponse, error) {
			called++
			return &contacts.ListFieldValuesResponse{}, &client.APIResponse{StatusCode: 200}, nil
		},
		AddFieldToGroupFunc: func(ctx context.Context, req interface{}) (*client.APIResponse, error) {
			called++
			return &client.APIResponse{StatusCode: 200}, nil
		},
		GetFieldGroupFunc: func(ctx context.Context, id string) (*contacts.FieldGroupResponse, *client.APIResponse, error) {
			called++
			return &contacts.FieldGroupResponse{}, &client.APIResponse{StatusCode: 200}, nil
		},
		DeleteFieldGroupFunc: func(ctx context.Context, id string) (*client.APIResponse, error) {
			called++
			return &client.APIResponse{StatusCode: 204}, nil
		},
	}

	ctx := context.Background()
	// Call each function-backed method
	_, _, _ = m.SearchByEmail(ctx, "a@b.com")
	_, _, _ = m.GetContact(ctx, "c1")
	_, _ = m.DeleteContact(ctx, "c1")
	_, _, _ = m.UpdateContact(ctx, "c1", &contacts.CreateContactRequest{})
	_, _, _ = m.TagsGet(ctx, "c1")
	_, _, _ = m.GetContactByEmailWithTags(ctx, "a@b.com")
	_, _, _ = m.CreateContactWithTags(ctx, &contacts.CreateContactRequest{}, []string{"t1"})
	_, _, _ = m.GetContactBounceLogs(ctx, "c1")
	_, _, _ = m.GetContactGoals(ctx, "c1")
	_, _, _ = m.GetContactLists(ctx, "c1")
	_, _, _ = m.GetContactLogs(ctx, "c1")
	_, _, _ = m.GetContactDealList(ctx, "c1")
	_, _, _ = m.GetContactDeals(ctx, "c1")
	_, _, _ = m.GetContactGeoIPs(ctx, "c1")
	_, _, _ = m.GetContactGeoIP(ctx, "c1", "127.0.0.1")
	_, _, _ = m.GetContactNotes(ctx, "c1")
	_, _, _ = m.GetContactOrganization(ctx, "c1")
	_, _, _ = m.GetContactAccountContacts(ctx, "c1")
	_, _, _ = m.GetContactPlusAppend(ctx, "c1")
	_, _, _ = m.GetContactTrackingLogs(ctx, "c1")
	_, _, _ = m.SyncContact(ctx, nil)
	_, _, _ = m.AddContactToList(ctx, nil)
	_, _, _ = m.BulkImportContacts(ctx, nil)
	_, _, _ = m.GetBulkImportStatus(ctx, "id")
	_, _, _ = m.ListBulkImportStatus(ctx, nil)
	_, _, _ = m.CreateCustomField(ctx, &contacts.FieldPayload{})
	_, _, _ = m.ListCustomFields(ctx)
	_, _, _ = m.AddFieldOption(ctx, &contacts.FieldOptionPayload{})
	_, _, _ = m.ListFieldValues(ctx)
	_, _ = m.AddFieldToGroup(ctx, nil)
	_, _, _ = m.GetFieldGroup(ctx, "g1")
	_, _ = m.DeleteFieldGroup(ctx, "g1")

	if called < 30 { // ensure we exercised most paths
		t.Fatalf("expected many calls, got %d", called)
	}
}
