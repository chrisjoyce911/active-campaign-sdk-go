package accounts

// Typed request/response shapes for the Accounts service.

// Account is a minimal representation used in examples and tests.
type Account struct {
	ID   string  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// CreateAccountRequest wraps the account payload for create calls.
type CreateAccountRequest struct {
	Account Account `json:"account"`
}

// CreateAccountResponse represents the response containing the account.
type CreateAccountResponse struct {
	Account Account `json:"account"`
}

// UpdateAccountRequest is used for account update payloads.
type UpdateAccountRequest struct {
	Account Account `json:"account"`
}

// ListAccountsResponse models a basic accounts list response.
type ListAccountsResponse struct {
	Accounts []Account              `json:"accounts"`
	Meta     map[string]interface{} `json:"meta,omitempty"`
}

// AccountNoteRequest for creating/updating notes.
type AccountNoteRequest struct {
	Note map[string]interface{} `json:"note"`
}

// AccountNoteResponse models the response containing note/account info.
type AccountNoteResponse struct {
	// Leave flexible: tests only validate http status; include raw fields
	// for potential future assertions.
	Account interface{} `json:"account,omitempty"`
	Note    interface{} `json:"note,omitempty"`
}
