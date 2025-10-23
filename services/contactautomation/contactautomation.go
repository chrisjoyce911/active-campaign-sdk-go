package contactautomation

import (
	"context"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

type service struct {
	client client.Doer
}

// ContactAutomationService defines the contract for contact automation operations.
type ContactAutomationService interface {
	GetCounts(ctx context.Context) (*AutomationCountsResponse, *client.APIResponse, error)
	GetAutomationEntry(ctx context.Context, id string) (*ContactAutomationResponse, *client.APIResponse, error)
	AddContactToAutomation(ctx context.Context, req *CreateContactAutomationRequest) (*ContactAutomationResponse, *client.APIResponse, error)
	RemoveContactFromAutomation(ctx context.Context, id string) (*client.APIResponse, error)
	ListContactAutomations(ctx context.Context, contactID string) (*ListContactAutomationsResponse, *client.APIResponse, error)
}

// Accessors
func (l *ListContactAutomationsResponse) ContactAutomationsOrEmpty() []ContactAutomationPayload {
	if l == nil || l.ContactAutomations == nil {
		return []ContactAutomationPayload{}
	}
	return *l.ContactAutomations
}

func (a *AutomationCountsResponse) CountsOrEmpty() []AutomationCount {
	if a == nil || a.Counts == nil {
		return []AutomationCount{}
	}
	return *a.Counts
}
