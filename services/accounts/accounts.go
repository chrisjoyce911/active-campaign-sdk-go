package accounts

import (
	"context"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// AccountsService handles Account related API calls.
//
// Methods follow the conventional pattern used across the v3 SDK:
// - ctx: context for cancellation/timeouts
// - return typed body (interface{} until models are added), *client.APIResponse for low-level details, and error
//
// See individual method docs in implementation files.
type service struct {
	client client.Doer
}

type AccountsService interface {
	// GetAccount fetches a single account by ID
	GetAccount(ctx context.Context, id string) (*CreateAccountResponse, *client.APIResponse, error)

	// CreateAccount creates a new account
	CreateAccount(ctx context.Context, req *CreateAccountRequest) (*CreateAccountResponse, *client.APIResponse, error)

	// ListAccounts lists accounts with optional query params
	ListAccounts(ctx context.Context, opts map[string]string) (*ListAccountsResponse, *client.APIResponse, error)

	// DeleteAccount deletes a single account by ID
	DeleteAccount(ctx context.Context, id string) (*client.APIResponse, error)

	// BulkDeleteAccounts deletes multiple accounts (by IDs)
	BulkDeleteAccounts(ctx context.Context, ids []string) (*client.APIResponse, error)

	// UpdateAccount updates an account by ID
	UpdateAccount(ctx context.Context, id string, req *UpdateAccountRequest) (*CreateAccountResponse, *client.APIResponse, error)

	// CreateAccountNote creates a note for an account
	CreateAccountNote(ctx context.Context, accountID string, req *AccountNoteRequest) (*AccountNoteResponse, *client.APIResponse, error)

	// UpdateAccountNote updates an existing account note
	UpdateAccountNote(ctx context.Context, noteID string, req *AccountNoteRequest) (*AccountNoteResponse, *client.APIResponse, error)
}
