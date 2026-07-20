package contacts

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRealService_TagRemove(t *testing.T) {
	tests := []struct {
		name       string
		contactID  string
		tag        string
		tagsBody   []byte
		statusCode int
	}{
		{
			name:       "remove tag from contact",
			contactID:  "1",
			tag:        "foo",
			tagsBody:   []byte(`{"contactTags":[{"id":"22","tag":"foo","contact":"1","cdate":"2025-01-01T00:00:00-06:00"}]}`),
			statusCode: 200,
		},
		{
			name:       "tag not found",
			contactID:  "1",
			tag:        "bar",
			tagsBody:   []byte(`{"contactTags":[{"id":"22","tag":"foo","contact":"1","cdate":"2025-01-01T00:00:00-06:00"}]}`),
			statusCode: 404,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := &mockDoer{Resp: &client.APIResponse{StatusCode: tt.statusCode}, Body: tt.tagsBody}
			require := require.New(t)
			require.NotNil(md)

			svc := NewRealServiceFromDoer(md)
			require.NotNil(svc)

			apiResp, err := svc.TagRemove(context.Background(), tt.contactID, tt.tag)
			assert.NoError(t, err)
			require.NotNil(apiResp)
			assert.Equal(t, tt.statusCode, apiResp.StatusCode)
		})
	}
}

// statusDoer returns a canned status/error for the DELETE, so the 404
// idempotency path can be exercised.
type statusDoer struct {
	Status int
	Err    error
	Path   string
	Calls  int
}

func (d *statusDoer) Do(ctx context.Context, method, path string, v interface{}, out interface{}) (*client.APIResponse, error) {
	d.Calls++
	d.Path = path
	return &client.APIResponse{StatusCode: d.Status}, d.Err
}

func TestTagRemoveByAssociation_SingleRequest(t *testing.T) {
	d := &statusDoer{Status: 200}
	svc := NewRealServiceFromDoer(d)

	resp, err := svc.TagRemoveByAssociation(context.Background(), "5942809")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("status = %d", resp.StatusCode)
	}
	if d.Path != "contactTags/5942809" {
		t.Fatalf("path = %q", d.Path)
	}
	if d.Calls != 1 {
		t.Fatalf("want a single request, got %d", d.Calls)
	}
}

// A tag removed between lookup and delete must not surface as an error —
// production saw "No Result found for SubscriberTag" requeue a message.
func TestTagRemoveByAssociation_AlreadyGoneIsSuccess(t *testing.T) {
	d := &statusDoer{Status: 404, Err: &client.APIError{StatusCode: 404, Message: "No Result found for SubscriberTag with id 5942809"}}
	svc := NewRealServiceFromDoer(d)

	if _, err := svc.TagRemoveByAssociation(context.Background(), "5942809"); err != nil {
		t.Fatalf("404 on delete should be treated as already-removed, got %v", err)
	}
}

func TestTagRemoveByAssociation_RealErrorPropagates(t *testing.T) {
	d := &statusDoer{Status: 500, Err: &client.APIError{StatusCode: 500, Message: "boom"}}
	svc := NewRealServiceFromDoer(d)

	if _, err := svc.TagRemoveByAssociation(context.Background(), "1"); err == nil {
		t.Fatal("want error for a 500")
	}
}
