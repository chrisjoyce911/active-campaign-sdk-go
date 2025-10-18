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
	UpdateListStatus(ctx context.Context, req *UpdateListStatusForContactRequest) (*UpdateContactListStatusResponse, *client.APIResponse, error)
}
