package contacts

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
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
	svc := NewRealServiceFromDoer(md)

	// Create
	t.Run("Create", func(t *testing.T) {
		req := &CreateContactRequest{Contact: &Contact{Email: "a@b.com", FirstName: "A"}}
		out, apiResp, err := svc.Create(context.Background(), req)
		assert.NoError(t, err)
		assert.Equal(t, 200, apiResp.StatusCode)
		assert.NotNil(t, out)
		assert.Equal(t, "a@b.com", out.Contact.Email)
	})

	// Create: ensure optional nested arrays are nil when absent
	md.Body = createBody
	t.Run("CreateOptionalArraysNil", func(t *testing.T) {
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
	})

	// Create: ensure optional nested arrays unmarshal when present
	withArraysBody := []byte(`{"contact": {"email": "a@b.com", "firstName": "A"},
        "contactAutomations": [{"contact":"1","seriesid":"2","id":"10"}],
        "contactData": [{"contact":"1","tstamp":"x","id":"20"}],
        "contactLists": [{"contact":"1","list":"5","id":"30"}],
        "fieldValues": [{"contact":"1","field":"7","value":"v","id":"40"}],
        "geoAddresses": [{"ip4":"1.2.3.4","id":"50"}],
        "geoIps": [{"contact":"1","id":"60"}],
        "scoreValues": [{"score":"100","id":"70"}]}`)

	md.Body = withArraysBody
	t.Run("CreateOptionalArraysPresent", func(t *testing.T) {
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
	})

	// SearchByEmail
	md.Body = searchBody
	t.Run("SearchByEmail", func(t *testing.T) {
		out, apiResp, err := svc.SearchByEmail(context.Background(), "a@b.com")
		assert.NoError(t, err)
		assert.Equal(t, 200, apiResp.StatusCode)
		assert.Len(t, out.Contacts, 1)
	})

	// GetContact
	md.Body = getBody
	t.Run("GetContact", func(t *testing.T) {
		out, apiResp, err := svc.GetContact(context.Background(), "1")
		assert.NoError(t, err)
		assert.Equal(t, 200, apiResp.StatusCode)
		assert.NotNil(t, out.Contact)
	})

	// ListContacts
	md.Body = listBody
	t.Run("ListContacts", func(t *testing.T) {
		out, apiResp, err := svc.ListContacts(context.Background(), map[string]string{})
		assert.NoError(t, err)
		assert.Equal(t, 200, apiResp.StatusCode)
		assert.Len(t, out.Contacts, 1)
	})
}
