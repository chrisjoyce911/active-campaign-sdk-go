package contactautomation

import (
	"context"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// ContactAutomationPayload represents a contact-automation relationship.
type ContactAutomationPayload struct {
	ID         string `json:"id,omitempty"`
	Contact    string `json:"contact,omitempty"`
	Automation string `json:"automation,omitempty"`
	Status     string `json:"status,omitempty"`
}

// CreateContactAutomationRequest is the request envelope for adding a contact to an automation.
type CreateContactAutomationRequest struct {
	ContactAutomation ContactAutomationPayload `json:"contactAutomation"`
}

// ContactAutomationResponse is the envelope returned for single contact-automation endpoints.
type ContactAutomationResponse struct {
	ContactAutomation ContactAutomationPayload `json:"contactAutomation"`
}

// ListContactAutomationsResponse represents a list response for contact automations.
type ListContactAutomationsResponse struct {
	ContactAutomations *[]ContactAutomationPayload `json:"contactAutomations"`
}

// AutomationCount represents a simple automation count.
type AutomationCount struct {
	Automation string `json:"automation"`
	Count      int    `json:"count"`
}

// AutomationCountsResponse is the envelope for counts endpoint.
type AutomationCountsResponse struct {
	Counts *[]AutomationCount `json:"counts"`
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

// ContactAutomationService defines the contract for contact automation operations.
type ContactAutomationService interface {
	GetCounts(ctx context.Context) (*AutomationCountsResponse, *client.APIResponse, error)
	GetAutomationEntry(ctx context.Context, id string) (*ContactAutomationResponse, *client.APIResponse, error)
	AddContactToAutomation(ctx context.Context, req *CreateContactAutomationRequest) (*ContactAutomationResponse, *client.APIResponse, error)
	RemoveContactFromAutomation(ctx context.Context, id string) (*client.APIResponse, error)
	ListContactAutomations(ctx context.Context, contactID string) (*ListContactAutomationsResponse, *client.APIResponse, error)
}

type service struct {
	client client.Doer
}
