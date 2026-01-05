package active_campaign

// root_stub.go provides minimal, buildable stubs for the module root package
// so example programs that import the root module can compile during migration.
//
// This file now wires the legacy-style root Client to the typed services so
// legacy call sites (for example `client.Contacts.SearchEmail`) can perform
// real API calls when a BaseUrl/Token are provided.

import (
	"context"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/contacts"
)

// ClientOpts are a simplified placeholder for example compatibility.
type ClientOpts struct {
	BaseUrl    string
	Token      string
	HttpClient interface{}
}

// Response wraps the low-level *client.APIResponse so legacy examples can
// inspect status/body while we keep the typed APIResponse available.
type Response struct {
	APIResp *client.APIResponse
}

// Client is a minimal shim that exposes the legacy-ish Contacts and Tags
// APIs. Contacts will delegate to the typed contacts service when possible.
type Client struct {
	Contacts *ContactsAPI
	Tags     *TagsAPI
}

// ContactsAPI wraps the typed contacts service and adapts its SearchByEmail
// into the legacy SearchEmail signature used by older examples.
type ContactsAPI struct {
	svc *contacts.RealService
}

// SearchEmail performs a search by email using the typed contacts service.
// It returns the raw typed response as interface{} and a wrapped Response.
func (c *ContactsAPI) SearchEmail(email string) (interface{}, *Response, error) {
	if c == nil || c.svc == nil {
		return nil, nil, nil
	}
	out, apiResp, err := c.svc.SearchByEmail(context.Background(), email)
	var resp *Response
	if apiResp != nil {
		resp = &Response{APIResp: apiResp}
	}
	return out, resp, err
}

// GetContactLists placeholder - adapt later if needed
func (c *ContactsAPI) GetContactLists(contactID string) (interface{}, *Response, error) {
	return nil, nil, nil
}

// TagsGetContact placeholder
func (c *ContactsAPI) TagsGet(contactID string) (interface{}, *Response, error) {
	return nil, nil, nil
}

// RemoveContactTag placeholder
func (c *ContactsAPI) RemoveContactTag(contactID, tagID string) (*Response, error) {
	return nil, nil
}

// NewClient returns a Client wired to a CoreClient using the provided opts.
// If BaseUrl/Token are empty the function falls back to returning a client
// with stubbed (nil-backed) Contacts to preserve backwards-compatibility.
func NewClient(opts *ClientOpts) (*Client, error) {
	if opts == nil || opts.BaseUrl == "" {
		// return a lightweight stub to preserve existing example behavior
		return &Client{Contacts: &ContactsAPI{}, Tags: &TagsAPI{}}, nil
	}
	core, err := client.NewCoreClient(opts.BaseUrl, opts.Token)
	if err != nil {
		return nil, err
	}
	contactsSvc := contacts.NewRealService(core)
	return &Client{Contacts: &ContactsAPI{svc: contactsSvc}, Tags: &TagsAPI{}}, nil
}

// TagsAPI placeholder
type TagsAPI struct{}

// RemoveTag placeholder
func (t *TagsAPI) RemoveTag(tagID string) (*Response, error) {
	return nil, nil
}
