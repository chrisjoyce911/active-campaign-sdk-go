package contacts

import (
	"context"
	"net/http"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func Test_EnsureSubscribedToList_CreatesWhenNotPresent(t *testing.T) {
	// GET shows no membership -> wrapper should POST subscription with status "1"
	body := []byte(`{"contactLists":[]}`)
	hd := &testhelpers.HTTPDoer{BaseURL: "https://example.com/api/3/", RespStatus: 200, RespBody: body}
	svc := NewRealServiceFromDoer(hd)

	_, apiResp, err := svc.EnsureSubscribedToList(context.Background(), "c1", "l1", false)
	assert.NoError(t, err)
	assert.NotNil(t, apiResp)

	if assert.NotNil(t, hd.LastRequest) {
		assert.Equal(t, http.MethodPost, hd.LastRequest.Method)
		assert.Equal(t, "/api/3/contactLists", hd.LastRequest.URL.Path)
		// ensure defaults to status:"1"
		assert.Contains(t, string(hd.LastRequestBody), `"status":"1"`)
		assert.Contains(t, string(hd.LastRequestBody), `"contact":"c1"`)
		assert.Contains(t, string(hd.LastRequestBody), `"list":"l1"`)
	}
}

func Test_EnsureSubscribedToList_RespectsForceFlag(t *testing.T) {
	// Existing membership is Unsubscribed ("2"). With force=false, wrapper should not POST.
	bodyUnsub := []byte(`{"contactLists":[{"contact":"c2","list":"l2","status":"2"}]}`)
	hd := &testhelpers.HTTPDoer{BaseURL: "https://example.com/api/3/", RespStatus: 200, RespBody: bodyUnsub}
	svc := NewRealServiceFromDoer(hd)

	out, apiResp, err := svc.EnsureSubscribedToList(context.Background(), "c2", "l2", false)
	assert.NoError(t, err)
	assert.Nil(t, out)
	assert.NotNil(t, apiResp)
	if assert.NotNil(t, hd.LastRequest) {
		// Only GET should have been issued
		assert.Equal(t, http.MethodGet, hd.LastRequest.Method)
		assert.Equal(t, "/api/3/contacts/c2/contactLists", hd.LastRequest.URL.Path)
	}

	// With force=true -> should POST
	hd2 := &testhelpers.HTTPDoer{BaseURL: "https://example.com/api/3/", RespStatus: 200, RespBody: bodyUnsub}
	svc2 := NewRealServiceFromDoer(hd2)

	_, apiResp2, err2 := svc2.EnsureSubscribedToList(context.Background(), "c2", "l2", true)
	assert.NoError(t, err2)
	assert.NotNil(t, apiResp2)
	if assert.NotNil(t, hd2.LastRequest) {
		assert.Equal(t, http.MethodPost, hd2.LastRequest.Method)
		assert.Equal(t, "/api/3/contactLists", hd2.LastRequest.URL.Path)
		assert.Contains(t, string(hd2.LastRequestBody), `"status":"1"`)
	}
}

func Test_EnsureSubscribedToList_NoOpWhenAlreadySubscribed(t *testing.T) {
	// Existing membership is already Subscribed ("1"). Even with force=true, no POST should occur.
	bodySub := []byte(`{"contactLists":[{"contact":"c3","list":"l3","status":"1"}]}`)
	hd := &testhelpers.HTTPDoer{BaseURL: "https://example.com/api/3/", RespStatus: 200, RespBody: bodySub}
	svc := NewRealServiceFromDoer(hd)

	out, apiResp, err := svc.EnsureSubscribedToList(context.Background(), "c3", "l3", true)
	assert.NoError(t, err)
	assert.Nil(t, out)
	assert.NotNil(t, apiResp)
	if assert.NotNil(t, hd.LastRequest) {
		assert.Equal(t, http.MethodGet, hd.LastRequest.Method)
		assert.Equal(t, "/api/3/contacts/c3/contactLists", hd.LastRequest.URL.Path)
	}
}
