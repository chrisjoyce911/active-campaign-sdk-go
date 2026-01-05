package contacts

import (
	"context"
	"net/http"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func Test_UpdateListStatusManaged_SkipWhenUnsubscribedAndNotForce(t *testing.T) {
	// Existing membership is Unsubscribed (status "2"). Force=false -> skip POST
	body := []byte(`{"contactLists":[{"contact":"c1","list":"l1","status":2}]}`)
	hd := &testhelpers.HTTPDoer{BaseURL: "https://example.com/api/3/", RespStatus: 200, RespBody: body}
	svc := NewRealServiceFromDoer(hd)

	req := &UpdateListStatusHelperRequest{ContactList: &ContactList{Contact: "c1", List: ListID("l1"), Status: 1}, Force: false}
	out, apiResp, err := svc.UpdateListStatusManaged(context.Background(), req)
	assert.NoError(t, err)
	assert.Nil(t, out)
	assert.NotNil(t, apiResp)
	// Should only issue GET
	if assert.NotNil(t, hd.LastRequest) {
		assert.Equal(t, http.MethodGet, hd.LastRequest.Method)
		assert.Equal(t, "/api/3/contacts/c1/contactLists", hd.LastRequest.URL.Path)
	}
}

func Test_UpdateListStatusManaged_ForceWhenUnsubscribed(t *testing.T) {
	// Existing membership is Unsubscribed (status "2"). Force=true -> POST subscribe
	body := []byte(`{"contactLists":[{"contact":"c1","list":"l1","status":2}]}`)
	hd := &testhelpers.HTTPDoer{BaseURL: "https://example.com/api/3/", RespStatus: 200, RespBody: body}
	svc := NewRealServiceFromDoer(hd)

	req := &UpdateListStatusHelperRequest{ContactList: &ContactList{Contact: "c1", List: ListID("l1"), Status: 1}, Force: true}
	out, apiResp, err := svc.UpdateListStatusManaged(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, apiResp)
	_ = out
	if assert.NotNil(t, hd.LastRequest) {
		assert.Equal(t, http.MethodPost, hd.LastRequest.Method)
		assert.Equal(t, "/api/3/contactLists", hd.LastRequest.URL.Path)
		// ensure request body contains desired payload
		assert.Contains(t, string(hd.LastRequestBody), `"contactList"`)
		assert.Contains(t, string(hd.LastRequestBody), `"contact":"c1"`)
		assert.Contains(t, string(hd.LastRequestBody), `"list":"l1"`)
		assert.Contains(t, string(hd.LastRequestBody), `"status":1`)
	}
}

func Test_UpdateListStatusManaged_CreateWhenNotPresent(t *testing.T) {
	// No membership present -> POST subscribe
	body := []byte(`{"contactLists":[]}`)
	hd := &testhelpers.HTTPDoer{BaseURL: "https://example.com/api/3/", RespStatus: 200, RespBody: body}
	svc := NewRealServiceFromDoer(hd)

	req := &UpdateListStatusHelperRequest{ContactList: &ContactList{Contact: "c1", List: ListID("l2"), Status: 1}, Force: false}
	_, apiResp, err := svc.UpdateListStatusManaged(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, apiResp)
	if assert.NotNil(t, hd.LastRequest) {
		assert.Equal(t, http.MethodPost, hd.LastRequest.Method)
		assert.Equal(t, "/api/3/contactLists", hd.LastRequest.URL.Path)
	}
}

func Test_UpdateListStatusManaged_NoOpWhenAlreadySubscribed(t *testing.T) {
	// Already subscribed (status "1") -> no POST
	body := []byte(`{"contactLists":[{"contact":"c1","list":"l1","status":1}]}`)
	hd := &testhelpers.HTTPDoer{BaseURL: "https://example.com/api/3/", RespStatus: 200, RespBody: body}
	svc := NewRealServiceFromDoer(hd)

	req := &UpdateListStatusHelperRequest{ContactList: &ContactList{Contact: "c1", List: ListID("l1"), Status: 1}}
	out, apiResp, err := svc.UpdateListStatusManaged(context.Background(), req)
	assert.NoError(t, err)
	assert.Nil(t, out)
	assert.NotNil(t, apiResp)
	if assert.NotNil(t, hd.LastRequest) {
		assert.Equal(t, http.MethodGet, hd.LastRequest.Method)
		assert.Equal(t, "/api/3/contacts/c1/contactLists", hd.LastRequest.URL.Path)
	}
}

func Test_UpdateListStatusManaged_DefaultDesiredStatusWhenEmpty(t *testing.T) {
	// No existing membership, desired status left empty -> defaults to "1" in POST body
	body := []byte(`{"contactLists":[]}`)
	hd := &testhelpers.HTTPDoer{BaseURL: "https://example.com/api/3/", RespStatus: 200, RespBody: body}
	svc := NewRealServiceFromDoer(hd)

	req := &UpdateListStatusHelperRequest{ContactList: &ContactList{Contact: "c9", List: ListID("l9"), Status: 0}, Force: false}
	_, apiResp, err := svc.UpdateListStatusManaged(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, apiResp)
	if assert.NotNil(t, hd.LastRequest) {
		assert.Equal(t, http.MethodPost, hd.LastRequest.Method)
		assert.Contains(t, string(hd.LastRequestBody), `"status":1`)
	}
}

func Test_UpdateListStatusManaged_ErrorsAndValidation(t *testing.T) {
	svc := NewRealServiceFromDoer(&testhelpers.HTTPDoer{})

	// nil request
	out, apiResp, err := svc.UpdateListStatusManaged(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, out)
	assert.Nil(t, apiResp)

	// nil ContactList
	out, apiResp, err = svc.UpdateListStatusManaged(context.Background(), &UpdateListStatusHelperRequest{})
	assert.Error(t, err)
	assert.Nil(t, out)
	assert.Nil(t, apiResp)

	// missing contact and list
	out, apiResp, err = svc.UpdateListStatusManaged(context.Background(), &UpdateListStatusHelperRequest{ContactList: &ContactList{}})
	assert.Error(t, err)
	assert.Nil(t, out)
	assert.Nil(t, apiResp)

	// GetContactLists error propagation (e.g., 404)
	hd := &testhelpers.HTTPDoer{BaseURL: "https://example.com/api/3/", RespStatus: 404, RespBody: []byte(`{"message":"No Result found for Subscriber with id 121"}`)}
	svc = NewRealServiceFromDoer(hd)
	out, apiResp, err = svc.UpdateListStatusManaged(context.Background(), &UpdateListStatusHelperRequest{ContactList: &ContactList{Contact: "121", List: ListID("25"), Status: 1}})
	assert.Error(t, err)
	assert.Nil(t, out)
	if assert.NotNil(t, apiResp) {
		// ensure we performed a GET and bubbled the response status
		assert.Equal(t, 404, apiResp.StatusCode)
	}
}
