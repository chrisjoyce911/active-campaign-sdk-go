package contacts

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// alias mockDoer to the shared test helper type so existing tests can keep
// using the md variable name without changing many files.
type mockDoer = testhelpers.MockDoer

func init() {
	// ensure the json package symbol stays referenced similarly to prior files
	_ = json.Unmarshal
}

func TestContacts_RealServiceSuccessPaths(t *testing.T) {
	// Provide example JSON bodies for a few endpoints
	createBody := []byte(`{"contact": {"email": "a@b.com", "firstName": "A"}}`)
	searchBody := []byte(`{"contacts": [{"email": "a@b.com"}]}`)
	getBody := []byte(`{"contact": {"email": "a@b.com", "firstName": "A"}}`)
	listBody := []byte(`{"contacts": [{"email": "a@b.com"}]}`)

	md := &mockDoer{Resp: &client.APIResponse{StatusCode: 200}, Err: nil, Body: createBody}
	require := require.New(t)
	require.NotNil(md)

	svc := NewRealServiceFromDoer(md)
	require.NotNil(svc)

	tests := []struct {
		name string
		body []byte
		call string
	}{
		{name: "Create", body: createBody, call: "create"},
		{name: "CreateOptionalArraysNil", body: createBody, call: "create-nil-arrays"},
		{name: "CreateOptionalArraysPresent", body: []byte(`{"contact": {"email": "a@b.com", "firstName": "A"},
		"contactAutomations": [{"contact":"1","seriesid":"2","id":"10"}],
		"contactData": [{"contact":"1","tstamp":"x","id":"20"}],
		"contactLists": [{"contact":"1","list":"5","id":"30"}],
		"fieldValues": [{"contact":"1","field":"7","value":"v","id":"40"}],
		"geoAddresses": [{"ip4":"1.2.3.4","id":"50"}],
		"geoIps": [{"contact":"1","id":"60"}],
		"scoreValues": [{"score":"100","id":"70"}]}`), call: "create-with-arrays"},
		{name: "SearchByEmail", body: searchBody, call: "search"},
		{name: "GetContact", body: getBody, call: "get"},
		{name: "ListContacts", body: listBody, call: "list"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			md.Body = tc.body
			switch tc.call {
			case "create":
				req := &CreateContactRequest{Contact: &Contact{Email: "a@b.com", FirstName: "A"}}
				out, apiResp, err := svc.Create(context.Background(), req)
				assert.NoError(t, err)
				assert.Equal(t, 200, apiResp.StatusCode)
				assert.NotNil(t, out)
				assert.Equal(t, "a@b.com", out.Contact.Email)

			case "create-nil-arrays":
				req := &CreateContactRequest{Contact: &Contact{Email: "a@b.com", FirstName: "A"}}
				out, apiResp, err := svc.Create(context.Background(), req)
				assert.NoError(t, err)
				assert.Equal(t, 200, apiResp.StatusCode)
				assert.NotNil(t, out)
				// After changing CreateContactResponse fields to pointer slices, absent arrays should be nil
				assert.Nil(t, out.ContactAutomations)
				assert.Nil(t, out.ContactData)
				assert.Nil(t, out.ContactLists)
				assert.Nil(t, out.FieldValues)
				assert.Nil(t, out.GeoAddresses)
				assert.Nil(t, out.GeoIps)
				assert.Nil(t, out.ScoreValues)

			case "create-with-arrays":
				req := &CreateContactRequest{Contact: &Contact{Email: "a@b.com", FirstName: "A"}}
				out, apiResp, err := svc.Create(context.Background(), req)
				assert.NoError(t, err)
				assert.Equal(t, 200, apiResp.StatusCode)
				assert.NotNil(t, out)
				// pointers should be non-nil and contain the expected single item
				if assert.NotNil(t, out.ContactAutomations) {
					assert.Len(t, *out.ContactAutomations, 1)
				}
				if assert.NotNil(t, out.ContactData) {
					assert.Len(t, *out.ContactData, 1)
				}
				if assert.NotNil(t, out.ContactLists) {
					assert.Len(t, *out.ContactLists, 1)
				}
				if assert.NotNil(t, out.FieldValues) {
					assert.Len(t, *out.FieldValues, 1)
				}
				if assert.NotNil(t, out.GeoAddresses) {
					assert.Len(t, *out.GeoAddresses, 1)
				}
				if assert.NotNil(t, out.GeoIps) {
					assert.Len(t, *out.GeoIps, 1)
				}
				if assert.NotNil(t, out.ScoreValues) {
					assert.Len(t, *out.ScoreValues, 1)
				}

			case "search":
				out, apiResp, err := svc.SearchByEmail(context.Background(), "a@b.com")
				assert.NoError(t, err)
				assert.Equal(t, 200, apiResp.StatusCode)
				assert.Len(t, out.Contacts, 1)

			case "get":
				out, apiResp, err := svc.GetContact(context.Background(), "1")
				assert.NoError(t, err)
				assert.Equal(t, 200, apiResp.StatusCode)
				assert.NotNil(t, out.Contact)

			case "list":
				out, apiResp, err := svc.ListContacts(context.Background(), map[string]string{})
				assert.NoError(t, err)
				assert.Equal(t, 200, apiResp.StatusCode)
				assert.Len(t, out.Contacts, 1)
			}
		})
	}
}
