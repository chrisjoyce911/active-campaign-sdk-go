package contactsmock

import (
	"context"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/contacts"
)

// Service is a function-field mock implementing contacts.ContactsService.
type Service struct {
	CreateFunc                             func(ctx context.Context, req *contacts.CreateContactRequest) (*contacts.CreateContactResponse, *client.APIResponse, error)
	SearchByEmailFunc                      func(ctx context.Context, email string) (*contacts.ContactSearchResponse, *client.APIResponse, error)
	GetContactFunc                         func(ctx context.Context, id string) (*contacts.CreateContactResponse, *client.APIResponse, error)
	DeleteContactFunc                      func(ctx context.Context, id string) (*client.APIResponse, error)
	UpdateContactFunc                      func(ctx context.Context, id string, req *contacts.CreateContactRequest) (*contacts.CreateContactResponse, *client.APIResponse, error)
	UpdateListStatusFunc                   func(ctx context.Context, req *contacts.UpdateListStatusForContactRequest) (*contacts.UpdateContactListStatusResponse, *client.APIResponse, error)
	GetContactFieldValuesFunc              func(ctx context.Context, id string) (interface{}, *client.APIResponse, error)
	GetContactFieldValuesTypedFunc         func(ctx context.Context, id string) (*contacts.ListFieldValuesResponse, *client.APIResponse, error)
	UpdateOrCreateFieldValueForContactFunc func(ctx context.Context, contactID, fieldIdentifier, value string) (*contacts.FieldValueResponse, *client.APIResponse, error)
	TagsGetFunc                            func(ctx context.Context, id string) (*contacts.ContactTagsResponse, *client.APIResponse, error)
	GetContactByEmailWithTagsFunc          func(ctx context.Context, email string) (interface{}, *client.APIResponse, error)
	CreateContactWithTagsFunc              func(ctx context.Context, req *contacts.CreateContactRequest, tagIDs []string) (*contacts.CreateContactResponse, *client.APIResponse, error)
	GetContactBounceLogsFunc               func(ctx context.Context, id string) (interface{}, *client.APIResponse, error)
	GetContactGoalsFunc                    func(ctx context.Context, id string) (interface{}, *client.APIResponse, error)
	GetContactListsFunc                    func(ctx context.Context, id string) (*contacts.ContactListsResponse, *client.APIResponse, error)
	GetContactLogsFunc                     func(ctx context.Context, id string) (interface{}, *client.APIResponse, error)
	GetContactDealListFunc                 func(ctx context.Context, id string) (interface{}, *client.APIResponse, error)
	GetContactDealsFunc                    func(ctx context.Context, id string) (interface{}, *client.APIResponse, error)
	GetContactGeoIPsFunc                   func(ctx context.Context, id string) (interface{}, *client.APIResponse, error)
	GetContactGeoIPFunc                    func(ctx context.Context, id, ip string) (interface{}, *client.APIResponse, error)
	GetContactNotesFunc                    func(ctx context.Context, id string) (interface{}, *client.APIResponse, error)
	GetContactOrganizationFunc             func(ctx context.Context, id string) (interface{}, *client.APIResponse, error)
	GetContactAccountContactsFunc          func(ctx context.Context, id string) (interface{}, *client.APIResponse, error)
	GetContactPlusAppendFunc               func(ctx context.Context, id string) (interface{}, *client.APIResponse, error)
	GetContactTrackingLogsFunc             func(ctx context.Context, id string) (interface{}, *client.APIResponse, error)
	SyncContactFunc                        func(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error)
	AddContactToListFunc                   func(ctx context.Context, req *contacts.AddContactToListPayload) (*contacts.AddContactToListResponse, *client.APIResponse, error)
	BulkImportContactsFunc                 func(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error)
	GetBulkImportStatusFunc                func(ctx context.Context, id string) (interface{}, *client.APIResponse, error)
	ListBulkImportStatusFunc               func(ctx context.Context, opts map[string]string) (interface{}, *client.APIResponse, error)
	CreateCustomFieldFunc                  func(ctx context.Context, req *contacts.FieldPayload) (*contacts.FieldResponse, *client.APIResponse, error)
	ListCustomFieldsFunc                   func(ctx context.Context) (*contacts.ListFieldsResponse, *client.APIResponse, error)
	UpdateCustomFieldFunc                  func(ctx context.Context, id string, req *contacts.FieldPayload) (*contacts.FieldResponse, *client.APIResponse, error)
	DeleteCustomFieldFunc                  func(ctx context.Context, id string) (*client.APIResponse, error)
	AddFieldOptionFunc                     func(ctx context.Context, req *contacts.FieldOptionPayload) (*contacts.FieldOptionResponse, *client.APIResponse, error)
	ListFieldValuesFunc                    func(ctx context.Context) (*contacts.ListFieldValuesResponse, *client.APIResponse, error)
	UpdateFieldValueForContactFunc         func(ctx context.Context, req *contacts.FieldValuePayload) (*contacts.FieldValueResponse, *client.APIResponse, error)
	AddFieldToGroupFunc                    func(ctx context.Context, req interface{}) (*client.APIResponse, error)
	GetFieldGroupFunc                      func(ctx context.Context, id string) (*contacts.FieldGroupResponse, *client.APIResponse, error)
	UpdateFieldGroupFunc                   func(ctx context.Context, id string, req interface{}) (*client.APIResponse, error)
	DeleteFieldGroupFunc                   func(ctx context.Context, id string) (*client.APIResponse, error)
	UpdateListStatusManagedFunc            func(ctx context.Context, req *contacts.UpdateListStatusHelperRequest) (*contacts.UpdateContactListStatusResponse, *client.APIResponse, error)
	EnsureSubscribedToListFunc             func(ctx context.Context, contactID, listID string, force bool) (*contacts.UpdateContactListStatusResponse, *client.APIResponse, error)
}

var _ contacts.ContactsService = (*Service)(nil)

func (m *Service) Create(ctx context.Context, req *contacts.CreateContactRequest) (*contacts.CreateContactResponse, *client.APIResponse, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, req)
	}
	return &contacts.CreateContactResponse{}, &client.APIResponse{}, nil
}
func (m *Service) SearchByEmail(ctx context.Context, email string) (*contacts.ContactSearchResponse, *client.APIResponse, error) {
	if m.SearchByEmailFunc != nil {
		return m.SearchByEmailFunc(ctx, email)
	}
	return &contacts.ContactSearchResponse{}, &client.APIResponse{}, nil
}
func (m *Service) GetContact(ctx context.Context, id string) (*contacts.CreateContactResponse, *client.APIResponse, error) {
	if m.GetContactFunc != nil {
		return m.GetContactFunc(ctx, id)
	}
	return &contacts.CreateContactResponse{}, &client.APIResponse{}, nil
}
func (m *Service) DeleteContact(ctx context.Context, id string) (*client.APIResponse, error) {
	if m.DeleteContactFunc != nil {
		return m.DeleteContactFunc(ctx, id)
	}
	return &client.APIResponse{}, nil
}
func (m *Service) UpdateContact(ctx context.Context, id string, req *contacts.CreateContactRequest) (*contacts.CreateContactResponse, *client.APIResponse, error) {
	if m.UpdateContactFunc != nil {
		return m.UpdateContactFunc(ctx, id, req)
	}
	return &contacts.CreateContactResponse{}, &client.APIResponse{}, nil
}
func (m *Service) UpdateListStatus(ctx context.Context, req *contacts.UpdateListStatusForContactRequest) (*contacts.UpdateContactListStatusResponse, *client.APIResponse, error) {
	if m.UpdateListStatusFunc != nil {
		return m.UpdateListStatusFunc(ctx, req)
	}
	return &contacts.UpdateContactListStatusResponse{}, &client.APIResponse{}, nil
}
func (m *Service) GetContactFieldValues(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	if m.GetContactFieldValuesFunc != nil {
		return m.GetContactFieldValuesFunc(ctx, id)
	}
	return nil, &client.APIResponse{}, nil
}
func (m *Service) GetContactFieldValuesTyped(ctx context.Context, id string) (*contacts.ListFieldValuesResponse, *client.APIResponse, error) {
	if m.GetContactFieldValuesTypedFunc != nil {
		return m.GetContactFieldValuesTypedFunc(ctx, id)
	}
	return &contacts.ListFieldValuesResponse{}, &client.APIResponse{}, nil
}
func (m *Service) UpdateOrCreateFieldValueForContact(ctx context.Context, contactID, fieldIdentifier, value string) (*contacts.FieldValueResponse, *client.APIResponse, error) {
	if m.UpdateOrCreateFieldValueForContactFunc != nil {
		return m.UpdateOrCreateFieldValueForContactFunc(ctx, contactID, fieldIdentifier, value)
	}
	return &contacts.FieldValueResponse{}, &client.APIResponse{}, nil
}
func (m *Service) TagsGet(ctx context.Context, id string) (*contacts.ContactTagsResponse, *client.APIResponse, error) {
	if m.TagsGetFunc != nil {
		return m.TagsGetFunc(ctx, id)
	}
	return &contacts.ContactTagsResponse{}, &client.APIResponse{}, nil
}
func (m *Service) GetContactByEmailWithTags(ctx context.Context, email string) (interface{}, *client.APIResponse, error) {
	if m.GetContactByEmailWithTagsFunc != nil {
		return m.GetContactByEmailWithTagsFunc(ctx, email)
	}
	return nil, &client.APIResponse{}, nil
}
func (m *Service) CreateContactWithTags(ctx context.Context, req *contacts.CreateContactRequest, tagIDs []string) (*contacts.CreateContactResponse, *client.APIResponse, error) {
	if m.CreateContactWithTagsFunc != nil {
		return m.CreateContactWithTagsFunc(ctx, req, tagIDs)
	}
	return &contacts.CreateContactResponse{}, &client.APIResponse{}, nil
}
func (m *Service) GetContactBounceLogs(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	if m.GetContactBounceLogsFunc != nil {
		return m.GetContactBounceLogsFunc(ctx, id)
	}
	return nil, &client.APIResponse{}, nil
}
func (m *Service) GetContactGoals(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	if m.GetContactGoalsFunc != nil {
		return m.GetContactGoalsFunc(ctx, id)
	}
	return nil, &client.APIResponse{}, nil
}
func (m *Service) GetContactLists(ctx context.Context, id string) (*contacts.ContactListsResponse, *client.APIResponse, error) {
	if m.GetContactListsFunc != nil {
		return m.GetContactListsFunc(ctx, id)
	}
	return &contacts.ContactListsResponse{}, &client.APIResponse{}, nil
}
func (m *Service) GetContactLogs(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	if m.GetContactLogsFunc != nil {
		return m.GetContactLogsFunc(ctx, id)
	}
	return nil, &client.APIResponse{}, nil
}
func (m *Service) GetContactDealList(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	if m.GetContactDealListFunc != nil {
		return m.GetContactDealListFunc(ctx, id)
	}
	return nil, &client.APIResponse{}, nil
}
func (m *Service) GetContactDeals(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	if m.GetContactDealsFunc != nil {
		return m.GetContactDealsFunc(ctx, id)
	}
	return nil, &client.APIResponse{}, nil
}
func (m *Service) GetContactGeoIPs(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	if m.GetContactGeoIPsFunc != nil {
		return m.GetContactGeoIPsFunc(ctx, id)
	}
	return nil, &client.APIResponse{}, nil
}
func (m *Service) GetContactGeoIP(ctx context.Context, id, ip string) (interface{}, *client.APIResponse, error) {
	if m.GetContactGeoIPFunc != nil {
		return m.GetContactGeoIPFunc(ctx, id, ip)
	}
	return nil, &client.APIResponse{}, nil
}
func (m *Service) GetContactNotes(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	if m.GetContactNotesFunc != nil {
		return m.GetContactNotesFunc(ctx, id)
	}
	return nil, &client.APIResponse{}, nil
}
func (m *Service) GetContactOrganization(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	if m.GetContactOrganizationFunc != nil {
		return m.GetContactOrganizationFunc(ctx, id)
	}
	return nil, &client.APIResponse{}, nil
}
func (m *Service) GetContactAccountContacts(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	if m.GetContactAccountContactsFunc != nil {
		return m.GetContactAccountContactsFunc(ctx, id)
	}
	return nil, &client.APIResponse{}, nil
}
func (m *Service) GetContactPlusAppend(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	if m.GetContactPlusAppendFunc != nil {
		return m.GetContactPlusAppendFunc(ctx, id)
	}
	return nil, &client.APIResponse{}, nil
}
func (m *Service) GetContactTrackingLogs(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	if m.GetContactTrackingLogsFunc != nil {
		return m.GetContactTrackingLogsFunc(ctx, id)
	}
	return nil, &client.APIResponse{}, nil
}
func (m *Service) SyncContact(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error) {
	if m.SyncContactFunc != nil {
		return m.SyncContactFunc(ctx, req)
	}
	return nil, &client.APIResponse{}, nil
}
func (m *Service) AddContactToList(ctx context.Context, req *contacts.AddContactToListPayload) (*contacts.AddContactToListResponse, *client.APIResponse, error) {
	if m.AddContactToListFunc != nil {
		return m.AddContactToListFunc(ctx, req)
	}
	return &contacts.AddContactToListResponse{}, &client.APIResponse{}, nil
}
func (m *Service) BulkImportContacts(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error) {
	if m.BulkImportContactsFunc != nil {
		return m.BulkImportContactsFunc(ctx, req)
	}
	return nil, &client.APIResponse{}, nil
}
func (m *Service) GetBulkImportStatus(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	if m.GetBulkImportStatusFunc != nil {
		return m.GetBulkImportStatusFunc(ctx, id)
	}
	return nil, &client.APIResponse{}, nil
}
func (m *Service) ListBulkImportStatus(ctx context.Context, opts map[string]string) (interface{}, *client.APIResponse, error) {
	if m.ListBulkImportStatusFunc != nil {
		return m.ListBulkImportStatusFunc(ctx, opts)
	}
	return nil, &client.APIResponse{}, nil
}
func (m *Service) CreateCustomField(ctx context.Context, req *contacts.FieldPayload) (*contacts.FieldResponse, *client.APIResponse, error) {
	if m.CreateCustomFieldFunc != nil {
		return m.CreateCustomFieldFunc(ctx, req)
	}
	return &contacts.FieldResponse{}, &client.APIResponse{}, nil
}
func (m *Service) ListCustomFields(ctx context.Context) (*contacts.ListFieldsResponse, *client.APIResponse, error) {
	if m.ListCustomFieldsFunc != nil {
		return m.ListCustomFieldsFunc(ctx)
	}
	return &contacts.ListFieldsResponse{}, &client.APIResponse{}, nil
}
func (m *Service) UpdateCustomField(ctx context.Context, id string, req *contacts.FieldPayload) (*contacts.FieldResponse, *client.APIResponse, error) {
	if m.UpdateCustomFieldFunc != nil {
		return m.UpdateCustomFieldFunc(ctx, id, req)
	}
	return &contacts.FieldResponse{}, &client.APIResponse{}, nil
}
func (m *Service) DeleteCustomField(ctx context.Context, id string) (*client.APIResponse, error) {
	if m.DeleteCustomFieldFunc != nil {
		return m.DeleteCustomFieldFunc(ctx, id)
	}
	return &client.APIResponse{}, nil
}
func (m *Service) AddFieldOption(ctx context.Context, req *contacts.FieldOptionPayload) (*contacts.FieldOptionResponse, *client.APIResponse, error) {
	if m.AddFieldOptionFunc != nil {
		return m.AddFieldOptionFunc(ctx, req)
	}
	return &contacts.FieldOptionResponse{}, &client.APIResponse{}, nil
}
func (m *Service) ListFieldValues(ctx context.Context) (*contacts.ListFieldValuesResponse, *client.APIResponse, error) {
	if m.ListFieldValuesFunc != nil {
		return m.ListFieldValuesFunc(ctx)
	}
	return &contacts.ListFieldValuesResponse{}, &client.APIResponse{}, nil
}
func (m *Service) UpdateFieldValueForContact(ctx context.Context, req *contacts.FieldValuePayload) (*contacts.FieldValueResponse, *client.APIResponse, error) {
	if m.UpdateFieldValueForContactFunc != nil {
		return m.UpdateFieldValueForContactFunc(ctx, req)
	}
	return &contacts.FieldValueResponse{}, &client.APIResponse{}, nil
}
func (m *Service) AddFieldToGroup(ctx context.Context, req interface{}) (*client.APIResponse, error) {
	if m.AddFieldToGroupFunc != nil {
		return m.AddFieldToGroupFunc(ctx, req)
	}
	return &client.APIResponse{}, nil
}
func (m *Service) GetFieldGroup(ctx context.Context, id string) (*contacts.FieldGroupResponse, *client.APIResponse, error) {
	if m.GetFieldGroupFunc != nil {
		return m.GetFieldGroupFunc(ctx, id)
	}
	return &contacts.FieldGroupResponse{}, &client.APIResponse{}, nil
}
func (m *Service) UpdateFieldGroup(ctx context.Context, id string, req interface{}) (*client.APIResponse, error) {
	if m.UpdateFieldGroupFunc != nil {
		return m.UpdateFieldGroupFunc(ctx, id, req)
	}
	return &client.APIResponse{}, nil
}
func (m *Service) DeleteFieldGroup(ctx context.Context, id string) (*client.APIResponse, error) {
	if m.DeleteFieldGroupFunc != nil {
		return m.DeleteFieldGroupFunc(ctx, id)
	}
	return &client.APIResponse{}, nil
}
func (m *Service) UpdateListStatusManaged(ctx context.Context, req *contacts.UpdateListStatusHelperRequest) (*contacts.UpdateContactListStatusResponse, *client.APIResponse, error) {
	if m.UpdateListStatusManagedFunc != nil {
		return m.UpdateListStatusManagedFunc(ctx, req)
	}
	return &contacts.UpdateContactListStatusResponse{}, &client.APIResponse{}, nil
}
func (m *Service) EnsureSubscribedToList(ctx context.Context, contactID, listID string, force bool) (*contacts.UpdateContactListStatusResponse, *client.APIResponse, error) {
	if m.EnsureSubscribedToListFunc != nil {
		return m.EnsureSubscribedToListFunc(ctx, contactID, listID, force)
	}
	return &contacts.UpdateContactListStatusResponse{}, &client.APIResponse{}, nil
}
