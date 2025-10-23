package contacts

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
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

// doerWrapper implements a doer that returns a canned fieldsBody for
// GET /fields requests, otherwise delegates to the embedded testDoer.
type doerWrapper struct {
	*testDoer
	fieldsBody []byte
}

func (d *doerWrapper) Do(ctx context.Context, method, path string, v interface{}, out interface{}) (*client.APIResponse, error) {
	if method == "GET" && strings.Contains(path, "fields") {
		if d.fieldsBody != nil && out != nil {
			_ = json.Unmarshal(d.fieldsBody, out)
		}
		return &client.APIResponse{StatusCode: 200, Body: d.fieldsBody}, nil
	}
	return d.testDoer.Do(ctx, method, path, v, out)
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
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if apiResp == nil || apiResp.StatusCode != 200 {
		t.Fatalf("expected 200, got %+v", apiResp)
	}
	if out == nil || out.FieldValue.ID != "fv123" {
		t.Fatalf("unexpected output: %+v", out)
	}
	if len(td.calls) < 2 {
		t.Fatalf("expected at least 2 calls, got %v", td.calls)
	}
	if !strings.HasPrefix(td.calls[1], "PUT ") {
		t.Fatalf("expected second call to be PUT, got %v", td.calls[1])
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
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if apiResp == nil || apiResp.StatusCode != 201 {
		t.Fatalf("expected 201, got %+v", apiResp)
	}
	if out == nil || out.FieldValue.ID != "fv999" {
		t.Fatalf("unexpected output: %+v", out)
	}
	if len(td.calls) < 2 {
		t.Fatalf("expected at least 2 calls, got %v", td.calls)
	}
	if !strings.HasPrefix(td.calls[1], "POST ") {
		t.Fatalf("expected second call to be POST, got %v", td.calls[1])
	}
}

func TestIsAllDigits(t *testing.T) {
	cases := map[string]bool{"": false, "123": true, "a12": false, "0123": true}
	for s, want := range cases {
		if got := isAllDigits(s); got != want {
			t.Fatalf("isAllDigits(%q) = %v, want %v", s, got, want)
		}
	}
}

func TestRealService_UpdateOrCreateFieldValueForContact_ResolveByPerstag_UpdateExisting(t *testing.T) {
	// fields list contains a matching perstag -> resolves to ID 99
	lf := &ListFieldsResponse{Fields: &[]FieldPayload{{ID: "99", Perstag: "mytag", Title: "My Title"}}}
	fieldsBody, _ := json.Marshal(lf)

	// GET contact fieldValues returns existing fv with Field == 99
	fvList := &ListFieldValuesResponse{FieldValues: &[]FieldValuePayload{{ID: "fv99", Field: "99", Value: "old"}}}
	fvBody, _ := json.Marshal(fvList)

	putBody := []byte(`{"fieldValue":{"id":"fv99","value":"new"}}`)

	td := &testDoer{getBody: fvBody, putBody: putBody}
	dw := &doerWrapper{testDoer: td, fieldsBody: fieldsBody}
	svc := NewRealServiceFromDoer(dw)

	out, apiResp, err := svc.UpdateOrCreateFieldValueForContact(context.Background(), "c1", "mytag", "new")
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if apiResp == nil || apiResp.StatusCode != 200 {
		t.Fatalf("expected 200, got %+v", apiResp)
	}
	if out == nil || out.FieldValue.ID != "fv99" {
		t.Fatalf("unexpected output: %+v", out)
	}
}

func TestRealService_UpdateOrCreateFieldValueForContact_MatchByFieldIdentifier_UpdateExisting(t *testing.T) {
	// Simulate ListCustomFields returning empty (no resolution)
	fieldsBody := []byte(`{}`)

	// GET contact fieldValues returns existing fv with Field == the provided identifier
	fvList := &ListFieldValuesResponse{FieldValues: &[]FieldValuePayload{{ID: "fvp", Field: "perstagXYZ", Value: "old"}}}
	fvBody, _ := json.Marshal(fvList)

	putBody := []byte(`{"fieldValue":{"id":"fvp","value":"new"}}`)

	td := &testDoer{getBody: fvBody, putBody: putBody}
	dw := &doerWrapper{testDoer: td, fieldsBody: fieldsBody}
	svc := NewRealServiceFromDoer(dw)

	out, apiResp, err := svc.UpdateOrCreateFieldValueForContact(context.Background(), "c1", "perstagXYZ", "new")
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if apiResp == nil || apiResp.StatusCode != 200 {
		t.Fatalf("expected 200, got %+v", apiResp)
	}
	if out == nil || out.FieldValue.ID != "fvp" {
		t.Fatalf("unexpected output: %+v", out)
	}
}

func TestRealService_UpdateOrCreateFieldValueForContact_EmptyContactID(t *testing.T) {
	svc := NewRealServiceFromDoer(&testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{}`)})
	out, apiResp, err := svc.UpdateOrCreateFieldValueForContact(context.Background(), "", "13", "new")
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if apiResp == nil || apiResp.StatusCode != 400 {
		t.Fatalf("expected 400 for empty contactID, got %+v", apiResp)
	}
	if out != nil {
		t.Fatalf("expected nil out for empty contactID")
	}
}

func TestRealService_UpdateOrCreateFieldValueForContact_EmptyFieldIdentifier_MatchEmptyField(t *testing.T) {
	// GET contact fieldValues returns existing fv with Field == "" and id fvE
	fvList := &ListFieldValuesResponse{FieldValues: &[]FieldValuePayload{{ID: "fvE", Field: "", Value: "old"}}}
	fvBody, _ := json.Marshal(fvList)
	putBody := []byte(`{"fieldValue":{"id":"fvE","value":"new"}}`)

	td := &testDoer{getBody: fvBody, putBody: putBody}
	svc := NewRealServiceFromDoer(td)

	out, apiResp, err := svc.UpdateOrCreateFieldValueForContact(context.Background(), "c1", "", "new")
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if apiResp == nil || apiResp.StatusCode != 200 {
		t.Fatalf("expected 200, got %+v", apiResp)
	}
	if out == nil || out.FieldValue.ID != "fvE" {
		t.Fatalf("unexpected output: %+v", out)
	}
}

// errThenPostDoer returns errors for GET fields and GET fieldValues calls,
// but responds to POST /fieldValues with a created response.
type errThenPostDoer struct{ postBody []byte }

func (d *errThenPostDoer) Do(ctx context.Context, method, path string, v interface{}, out interface{}) (*client.APIResponse, error) {
	if method == "GET" && strings.Contains(path, "fields") {
		return nil, errors.New("fields error")
	}
	if method == "GET" && strings.Contains(path, "fieldValues") {
		return nil, errors.New("fv error")
	}
	if method == "POST" && strings.Contains(path, "fieldValues") {
		if out != nil && d.postBody != nil {
			_ = json.Unmarshal(d.postBody, out)
		}
		return &client.APIResponse{StatusCode: 201, Body: d.postBody}, nil
	}
	md := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{}`)}
	return md.Do(ctx, method, path, v, out)
}

func TestRealService_UpdateOrCreateFieldValueForContact_ListAndGetError_FallsbackToPost(t *testing.T) {
	postBody := []byte(`{"fieldValue":{"id":"fvpost","value":"new"}}`)
	svc := NewRealServiceFromDoer(&errThenPostDoer{postBody: postBody})
	out, apiResp, err := svc.UpdateOrCreateFieldValueForContact(context.Background(), "c1", "nonmatch", "new")
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if apiResp == nil || apiResp.StatusCode != 201 {
		t.Fatalf("expected 201, got %+v", apiResp)
	}
	if out == nil || out.FieldValue.ID != "fvpost" {
		t.Fatalf("unexpected output: %+v", out)
	}
}
