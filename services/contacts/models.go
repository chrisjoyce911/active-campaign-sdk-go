package contacts

// Models for contacts API - these are placeholders and should match ActiveCampaign's API

// CreateContactRequest is the payload when creating a contact.
// TODO: expand fields per https://developers.activecampaign.com/reference#create-contact
type CreateContactRequest struct {
	Contact *Contact `json:"contact"`
}

// Contact represents the contact object.
// TODO: expand fields and tags, fieldValues, links, etc.
type Contact struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}

// CreateContactResponse is returned when a contact is created.
// TODO: align with real response fields.
type CreateContactResponse struct {
	Contact *Contact `json:"contact"`
}

// ContactSearchResponse for searching contacts by email.
// TODO: align with real response.
type ContactSearchResponse struct {
	Contacts []Contact `json:"contacts"`
}

// UpdateListStatusForContactRequest - TODO
type UpdateListStatusForContactRequest struct {
	ContactList *ContactList `json:"contactList"`
}

type ContactList struct {
	List    string `json:"list"`
	Contact string `json:"contact"`
	Status  string `json:"status"`
}

// UpdateContactListStatusResponse - TODO: align with real response
type UpdateContactListStatusResponse struct {
	// ...existing code...
}
