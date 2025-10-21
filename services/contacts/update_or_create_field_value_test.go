package contacts

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

// testDoer returns different canned bodies depending on method/path so we can
// simulate the sequence of calls made by UpdateOrCreateFieldValueForContact.
type testDoer struct {
	getBody  []byte
	putBody  []byte
	postBody []byte
	calls    []string
}

func (t *testDoer) Do(ctx context.Context, method, path string, v interface{}, out interface{}) (*client.APIResponse, error) {
	t.calls = append(t.calls, method+" "+path)
	var b []byte
	if method == "GET" && strings.Contains(path, "fieldValues") {
		b = t.getBody
		if out != nil && b != nil {
			_ = json.Unmarshal(b, out)
		}
		return &client.APIResponse{StatusCode: 200, Body: b}, nil
	}
	if method == "PUT" && strings.Contains(path, "fieldValues/") {
		b = t.putBody
		if out != nil && b != nil {
			_ = json.Unmarshal(b, out)
		}
		return &client.APIResponse{StatusCode: 200, Body: b}, nil
	}
	if method == "POST" && strings.Contains(path, "fieldValues") {
		b = t.postBody
		if out != nil && b != nil {
			_ = json.Unmarshal(b, out)
		}
		return &client.APIResponse{StatusCode: 201, Body: b}, nil
	}
	// Default: delegate to the simple MockDoer behaviour for other endpoints
	md := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{}`)}
	return md.Do(ctx, method, path, v, out)
}

func TestRealService_UpdateOrCreateFieldValueForContact_UpdateExisting(t *testing.T) {
	// Mock response for GET contact fieldValues
	fvList := &ListFieldValuesResponse{FieldValues: &[]FieldValuePayload{{ID: "fv123", Field: "13", Value: "old"}}}
	fvBody, _ := json.Marshal(fvList)

	// Mock response for PUT /fieldValues/{id}
	putBody := []byte(`{"fieldValue":{"id":"fv123","value":"new"}}`)

	td := &testDoer{getBody: fvBody, putBody: putBody}
	svc := NewRealServiceFromDoer(td)

	out, apiResp, err := svc.UpdateOrCreateFieldValueForContact(context.Background(), "c1", "13", "new")
	assert.NoError(t, err)
	assert.Equal(t, 200, apiResp.StatusCode)
	if assert.NotNil(t, out) {
		assert.Equal(t, "fv123", out.FieldValue.ID)
	}
	// Verify the doer saw a PUT as the second request
	if assert.Len(t, td.calls, 2) {
		assert.True(t, strings.HasPrefix(td.calls[1], "PUT "))
	}
}

func TestRealService_UpdateOrCreateFieldValueForContact_CreateNew(t *testing.T) {
	// Mock response for GET contact fieldValues (empty)
	fvList := &ListFieldValuesResponse{FieldValues: &[]FieldValuePayload{}}
	fvBody, _ := json.Marshal(fvList)

	// Mock response for POST /fieldValues
	postBody := []byte(`{"fieldValue":{"id":"fv999","value":"new"}}`)

	td := &testDoer{getBody: fvBody, postBody: postBody}
	svc := NewRealServiceFromDoer(td)

	out, apiResp, err := svc.UpdateOrCreateFieldValueForContact(context.Background(), "c1", "13", "new")
	assert.NoError(t, err)
	assert.Equal(t, 201, apiResp.StatusCode)
	if assert.NotNil(t, out) {
		assert.Equal(t, "fv999", out.FieldValue.ID)
	}
	// Verify the doer saw a POST as the second request
	if assert.Len(t, td.calls, 2) {
		assert.True(t, strings.HasPrefix(td.calls[1], "POST "))
	}
}
