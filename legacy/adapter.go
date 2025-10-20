package legacy

import (
	"context"
	"fmt"
	"os"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/contactautomation"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/contacts"
)

// getCoreClientFromEnv creates a CoreClient using environment variables
// ACTIVE_URL and ACTIVE_TOKEN. Examples should set these before calling.
func getCoreClientFromEnv() (*client.CoreClient, error) {
	base := os.Getenv("ACTIVE_URL")
	token := os.Getenv("ACTIVE_TOKEN")
	if base == "" {
		return nil, fmt.Errorf("ACTIVE_URL not set")
	}
	return client.NewCoreClient(base, token)
}

// CreateContact creates a contact using the CoreClient. The request body is
// marshalled and sent to POST /api/3/contacts. Returns the raw unmarshalled
// response (map[string]interface{}) for flexibility.
func CreateContact(ctx context.Context, req *contacts.CreateContactRequest) (*contacts.CreateContactResponse, *client.APIResponse, error) {
	c, err := getCoreClientFromEnv()
	if err != nil {
		return nil, nil, err
	}
	svc := contacts.NewRealService(c)
	return svc.Create(ctx, req)
}

// GetContact retrieves a contact by ID using GET /api/3/contacts/{id}
func GetContact(ctx context.Context, id string) (*contacts.CreateContactResponse, *client.APIResponse, error) {
	c, err := getCoreClientFromEnv()
	if err != nil {
		return nil, nil, err
	}
	svc := contacts.NewRealService(c)
	return svc.GetContact(ctx, id)
}

// DeleteContact deletes a contact by ID using DELETE /api/3/contacts/{id}
func DeleteContact(ctx context.Context, id string) (*client.APIResponse, error) {
	c, err := getCoreClientFromEnv()
	if err != nil {
		return nil, err
	}
	svc := contacts.NewRealService(c)
	return svc.DeleteContact(ctx, id)
}

// SearchContacts performs a basic email search using GET /api/3/contacts?email={email}
// It returns the raw unmarshalled response as map[string]interface{}.
func SearchContacts(ctx context.Context, email string) (*contacts.ContactSearchResponse, *client.APIResponse, error) {
	c, err := getCoreClientFromEnv()
	if err != nil {
		return nil, nil, err
	}
	// Use the typed contacts service to perform the search so the legacy adapter
	// delegates to the v3 service implementation.
	svc := contacts.NewRealService(c)
	return svc.SearchByEmail(ctx, email)
}

// Add additional legacy adapter implementations here as needed.

// GetAutomationCounts retrieves automation entry counts using the contactautomation service.
func GetAutomationCounts(ctx context.Context) (*contactautomation.AutomationCountsResponse, *client.APIResponse, error) {
	c, err := getCoreClientFromEnv()
	if err != nil {
		return nil, nil, err
	}
	svc := contactautomation.NewRealService(c)
	return svc.GetCounts(ctx)
}

// GetAutomationEntry retrieves a contact automation entry by id.
func GetAutomationEntry(ctx context.Context, id string) (*contactautomation.ContactAutomationResponse, *client.APIResponse, error) {
	c, err := getCoreClientFromEnv()
	if err != nil {
		return nil, nil, err
	}
	svc := contactautomation.NewRealService(c)
	return svc.GetAutomationEntry(ctx, id)
}

// AddContactToAutomation adds a contact to an automation.
func AddContactToAutomation(ctx context.Context, req *contactautomation.CreateContactAutomationRequest) (*contactautomation.ContactAutomationResponse, *client.APIResponse, error) {
	c, err := getCoreClientFromEnv()
	if err != nil {
		return nil, nil, err
	}
	svc := contactautomation.NewRealService(c)
	return svc.AddContactToAutomation(ctx, req)
}

// RemoveContactFromAutomation removes a contact from an automation.
func RemoveContactFromAutomation(ctx context.Context, id string) (*client.APIResponse, error) {
	c, err := getCoreClientFromEnv()
	if err != nil {
		return nil, err
	}
	svc := contactautomation.NewRealService(c)
	return svc.RemoveContactFromAutomation(ctx, id)
}

// ListContactAutomations lists automations for a contact.
func ListContactAutomations(ctx context.Context, contactID string) (*contactautomation.ListContactAutomationsResponse, *client.APIResponse, error) {
	c, err := getCoreClientFromEnv()
	if err != nil {
		return nil, nil, err
	}
	svc := contactautomation.NewRealService(c)
	return svc.ListContactAutomations(ctx, contactID)
}
