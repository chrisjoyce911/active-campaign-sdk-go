package active_campaign

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/services/contacts"
	"github.com/stretchr/testify/assert"
)

func TestNewClient_NoBaseURL_ReturnsStub(t *testing.T) {
	c, err := NewClient(&ClientOpts{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if c == nil || c.Contacts == nil {
		t.Fatalf("expected client with Contacts field")
	}

	// stubbed ContactsAPI should return nils when svc is not wired
	out, resp, err := c.Contacts.SearchEmail("x@example.com")
	assert.Nil(t, out)
	assert.Nil(t, resp)
	assert.Nil(t, err)
}

func TestNewClient_WithBaseURL_WiresContactsSearch(t *testing.T) {
	// Start a fake server that returns a minimal contacts search response
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resp := &contacts.ContactSearchResponse{
			Contacts: []contacts.Contact{{ID: "1", Email: "x@test.com"}},
			Meta:     map[string]interface{}{"total": "1"},
		}
		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer srv.Close()

	c, err := NewClient(&ClientOpts{BaseUrl: srv.URL, Token: "tok"})
	if err != nil {
		t.Fatalf("NewClient error: %v", err)
	}

	out, resp, err := c.Contacts.SearchEmail("x@test.com")
	assert.NoError(t, err)
	if assert.NotNil(t, resp) {
		// underlying APIResponse should be present
		assert.NotZero(t, resp.APIResp.StatusCode)
	}
	// result should be the typed ContactSearchResponse
	csr, ok := out.(*contacts.ContactSearchResponse)
	if assert.True(t, ok, "expected ContactSearchResponse") {
		assert.Equal(t, "x@test.com", csr.Contacts[0].Email)
	}
}
