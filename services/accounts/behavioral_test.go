package accounts

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestService_CreateAccount_RequestShapeAndErrorPaths(t *testing.T) {
	// Assert request shape and method/path
	hd := &testhelpers.HTTPDoer{BaseURL: "https://example.com/api/3/", RespStatus: 201, RespBody: []byte(`{"account":{"id":"a1","name":"X"}}`)}
	svc := NewRealServiceFromDoer(hd)

	req := &CreateAccountRequest{Account: Account{ID: "", Name: strptr("X")}}
	out, apiResp, err := svc.CreateAccount(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, 201, apiResp.StatusCode)
	// inspect recorded request
	reqMethod := hd.LastRequest.Method
	assert.Equal(t, "POST", reqMethod)
	assert.Equal(t, "/api/3/accounts", hd.LastRequest.URL.Path)
	// assert body JSON contains account.name
	var body map[string]map[string]interface{}
	err = json.Unmarshal(hd.LastRequestBody, &body)
	assert.NoError(t, err)
	assert.Equal(t, "X", body["account"]["name"])
	_ = out

	// Error path: non-2xx from API
	hdErr := &testhelpers.HTTPDoer{BaseURL: "https://example.com/api/3/", RespStatus: 400, RespBody: []byte(`{"errors":[{"title":"bad"}]}`)}
	svcErr := NewRealServiceFromDoer(hdErr)
	out2, apiResp2, err2 := svcErr.CreateAccount(context.Background(), req)
	assert.Error(t, err2)
	var apiErr *client.APIError
	ok := errors.As(err2, &apiErr)
	assert.True(t, ok, "expected error to be an APIError")
	assert.Equal(t, 400, apiResp2.StatusCode)
	assert.Nil(t, out2)
}

func TestService_UpdateAccount_RequestShapeAndBadJSON(t *testing.T) {
	// Assert request method/path & behavior when response has invalid JSON
	hd := &testhelpers.HTTPDoer{BaseURL: "https://example.com/api/3/", RespStatus: 200, RespBody: []byte(`{"account":{"id":"a1","name":"X"}}`)}
	svc := NewRealServiceFromDoer(hd)

	req := &UpdateAccountRequest{Account: Account{ID: "a1", Name: strptr("New")}}
	out, apiResp, err := svc.UpdateAccount(context.Background(), "a1", req)
	assert.NoError(t, err)
	assert.Equal(t, 200, apiResp.StatusCode)
	assert.Equal(t, "PUT", hd.LastRequest.Method)
	assert.Equal(t, "/api/3/accounts/a1", hd.LastRequest.URL.Path)
	_ = out

	// Now simulate invalid JSON in a 200 response
	hdBad := &testhelpers.HTTPDoer{BaseURL: "https://example.com/api/3/", RespStatus: 200, RespBody: []byte(`{invalid-json}`)}
	svcBad := NewRealServiceFromDoer(hdBad)
	out2, apiResp2, err2 := svcBad.UpdateAccount(context.Background(), "a1", req)
	assert.Error(t, err2)
	// in this case, apiResp2 should be non-nil and contain the raw body, out2 should be nil
	assert.NotNil(t, apiResp2)
	assert.Nil(t, out2)
}

func TestService_CreateAccountNote_RequestShape(t *testing.T) {
	hd := &testhelpers.HTTPDoer{BaseURL: "https://example.com/api/3/", RespStatus: 201, RespBody: []byte(`{"note":{"id":"n1"}}`)}
	svc := NewRealServiceFromDoer(hd)

	req := &AccountNoteRequest{Note: map[string]interface{}{"note": "hello"}}
	out, apiResp, err := svc.CreateAccountNote(context.Background(), "a1", req)
	assert.NoError(t, err)
	assert.Equal(t, 201, apiResp.StatusCode)
	assert.Equal(t, "POST", hd.LastRequest.Method)
	assert.Equal(t, "/api/3/accounts/a1/notes", hd.LastRequest.URL.Path)
	var body map[string]map[string]interface{}
	err = json.Unmarshal(hd.LastRequestBody, &body)
	assert.NoError(t, err)
	assert.Equal(t, "hello", body["note"]["note"])
	_ = out
}

// helper
func strptr(s string) *string { return &s }
