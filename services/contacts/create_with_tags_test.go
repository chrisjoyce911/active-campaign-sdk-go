package contacts

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
)

// createWithTagsDoer returns a created contact on POST /contacts and then
// returns responses for POST /contactTags. It records calls for assertions.
type createWithTagsDoer struct {
	createBody      []byte
	contactTagBody  []byte
	contactTagError bool
	calls           []string
}

func (d *createWithTagsDoer) Do(ctx context.Context, method, path string, v interface{}, out interface{}) (*client.APIResponse, error) {
	d.calls = append(d.calls, method+" "+path)
	if strings.ToUpper(method) == "POST" && strings.Contains(path, "contacts") {
		if out != nil && d.createBody != nil {
			_ = json.Unmarshal(d.createBody, out)
		}
		return &client.APIResponse{StatusCode: 201, Body: d.createBody}, nil
	}
	if strings.ToUpper(method) == "POST" && strings.Contains(path, "contactTags") {
		if d.contactTagError {
			body := []byte(`{"message":"boom"}`)
			apiErr := &client.APIError{StatusCode: 500, Message: "boom", Body: body}
			return &client.APIResponse{StatusCode: 500, Body: body}, apiErr
		}
		if out != nil && d.contactTagBody != nil {
			_ = json.Unmarshal(d.contactTagBody, out)
		}
		return &client.APIResponse{StatusCode: 201, Body: d.contactTagBody}, nil
	}
	// default delegate
	md := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{}`)}
	return md.Do(ctx, method, path, v, out)
}

func TestRealService_CreateContactWithTags_Success(t *testing.T) {
	createBody := []byte(`{"contact":{"id":"123","email":"jdoe@example.com"}}`)
	contactTagBody := []byte(`{"contactTag":{"contact":"123","tag":"100"}}`)
	td := &createWithTagsDoer{createBody: createBody, contactTagBody: contactTagBody}
	svc := NewRealServiceFromDoer(td)

	req := &CreateContactRequest{Contact: &Contact{Email: "jdoe@example.com"}}
	created, apiResp, err := svc.CreateContactWithTags(context.Background(), req, []string{"100", "101"})
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if apiResp != nil {
		t.Fatalf("expected nil apiResp on success, got %+v", apiResp)
	}
	if created == nil || created.Contact == nil || created.Contact.ID != "123" {
		t.Fatalf("unexpected created: %+v", created)
	}
	if len(td.calls) < 3 {
		t.Fatalf("expected at least 3 calls (create + 2 tags), got %v", td.calls)
	}
	if !strings.Contains(td.calls[0], "contacts") || !strings.Contains(td.calls[1], "contactTags") {
		t.Fatalf("unexpected call sequence: %v", td.calls)
	}
}

func TestRealService_CreateContactWithTags_AttachFailure(t *testing.T) {
	createBody := []byte(`{"contact":{"id":"123","email":"jdoe@example.com"}}`)
	td := &createWithTagsDoer{createBody: createBody, contactTagError: true}
	svc := NewRealServiceFromDoer(td)

	req := &CreateContactRequest{Contact: &Contact{Email: "jdoe@example.com"}}
	created, apiResp, err := svc.CreateContactWithTags(context.Background(), req, []string{"100"})
	if err == nil {
		t.Fatalf("expected attach error, got nil")
	}
	if apiResp == nil || apiResp.StatusCode != 500 {
		t.Fatalf("expected 500 apiResp from attach failure, got %+v", apiResp)
	}
	if created == nil || created.Contact == nil || created.Contact.ID != "123" {
		t.Fatalf("expected created contact despite attach failure, got %+v", created)
	}
}

// createErrorDoer returns an API error for the initial contact create call.
type createErrorDoer struct{}

func (d *createErrorDoer) Do(ctx context.Context, method, path string, v interface{}, out interface{}) (*client.APIResponse, error) {
	if strings.ToUpper(method) == "POST" && strings.Contains(path, "contacts") {
		body := []byte(`{"message":"create failed"}`)
		apiErr := &client.APIError{StatusCode: 500, Message: "create failed", Body: body}
		return &client.APIResponse{StatusCode: 500, Body: body}, apiErr
	}
	md := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{}`)}
	return md.Do(ctx, method, path, v, out)
}

func TestRealService_CreateContactWithTags_CreateError(t *testing.T) {
	svc := NewRealServiceFromDoer(&createErrorDoer{})

	req := &CreateContactRequest{Contact: &Contact{Email: "jdoe@example.com"}}
	created, apiResp, err := svc.CreateContactWithTags(context.Background(), req, []string{"100"})
	if err == nil {
		t.Fatalf("expected create error, got nil")
	}
	if apiResp == nil || apiResp.StatusCode != 500 {
		t.Fatalf("expected 500 apiResp from create failure, got %+v", apiResp)
	}
	// created may be non-nil (Create always returns a response struct), but it must not be treated as success
	if created == nil {
		// acceptable, but just ensure function returned something sensible
		t.Logf("created is nil as returned")
	}
}

// createNoContactDoer returns a 201 with an empty body (no contact) for POST /contacts.
type createNoContactDoer struct{}

func (d *createNoContactDoer) Do(ctx context.Context, method, path string, v interface{}, out interface{}) (*client.APIResponse, error) {
	if strings.ToUpper(method) == "POST" && strings.Contains(path, "contacts") {
		// return empty object -> out.Contact should be nil after unmarshal
		body := []byte(`{}`)
		if out != nil {
			_ = json.Unmarshal(body, out)
		}
		return &client.APIResponse{StatusCode: 201, Body: body}, nil
	}
	md := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{}`)}
	return md.Do(ctx, method, path, v, out)
}

func TestRealService_CreateContactWithTags_NoContactReturned(t *testing.T) {
	svc := NewRealServiceFromDoer(&createNoContactDoer{})

	req := &CreateContactRequest{Contact: &Contact{Email: "jdoe@example.com"}}
	created, apiResp, err := svc.CreateContactWithTags(context.Background(), req, []string{"100"})
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if apiResp == nil || apiResp.StatusCode != 201 {
		t.Fatalf("expected 201 apiResp from create, got %+v", apiResp)
	}
	if created == nil || created.Contact != nil {
		t.Fatalf("expected created.Contact to be nil, got %+v", created)
	}
}

func TestRealService_CreateContactWithTags_NoTags(t *testing.T) {
	createBody := []byte(`{"contact":{"id":"321","email":"jane@example.com"}}`)
	td := &createWithTagsDoer{createBody: createBody}
	svc := NewRealServiceFromDoer(td)

	req := &CreateContactRequest{Contact: &Contact{Email: "jane@example.com"}}
	created, apiResp, err := svc.CreateContactWithTags(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if apiResp != nil {
		t.Fatalf("expected nil apiResp on success, got %+v", apiResp)
	}
	if created == nil || created.Contact == nil || created.Contact.ID != "321" {
		t.Fatalf("unexpected created: %+v", created)
	}
	// Ensure only the create call was made
	if len(td.calls) != 1 || !strings.Contains(td.calls[0], "contacts") {
		t.Fatalf("expected only create call, got %v", td.calls)
	}
}
