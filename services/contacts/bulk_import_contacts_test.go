package contacts

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// bodyCapturingDoer records the marshalled request so the wire shape can be
// asserted — bulk import uses snake_case field names, unlike contact/sync.
type bodyCapturingDoer struct {
	Resp *client.APIResponse
	Body []byte
	Path string
	Sent []byte
	Opts interface{}
}

func (d *bodyCapturingDoer) Do(ctx context.Context, method, path string, v interface{}, out interface{}) (*client.APIResponse, error) {
	d.Path = path
	d.Opts = v
	if v != nil {
		d.Sent, _ = json.Marshal(v)
	}
	if out != nil && d.Body != nil {
		_ = json.Unmarshal(d.Body, out)
	}
	return d.Resp, nil
}

func TestRealService_BulkImportContacts(t *testing.T) {
	require := require.New(t)

	d := &bodyCapturingDoer{
		Resp: &client.APIResponse{StatusCode: 201},
		Body: []byte(`{"success":1,"queued_contacts":2,"batchId":"b-123"}`),
	}

	req := &BulkImportRequest{Contacts: []BulkImportContact{
		{
			Email:     "a@b.com",
			FirstName: "Ada",
			LastName:  "Lovelace",
			Tags:      []string{"52"},
			Fields:    []BulkImportField{{ID: 39, Value: "USI123"}},
			Subscribe: []BulkImportList{{ListID: 2}},
		},
		{Email: "c@d.com"},
	}}

	out, apiResp, err := NewRealServiceFromDoer(d).BulkImportContacts(context.Background(), req)
	require.NoError(err)
	require.NotNil(apiResp)

	// The path bug that made this method 404 for its whole life.
	assert.Equal(t, "import/bulk_import", d.Path)

	assert.Equal(t, 1, out.Success)
	assert.Equal(t, 2, out.QueuedContacts)
	assert.Equal(t, "b-123", out.BatchID)

	// Wire shape: snake_case, and list membership by listid.
	sent := string(d.Sent)
	assert.Contains(t, sent, `"first_name":"Ada"`)
	assert.Contains(t, sent, `"listid":2`)
	assert.Contains(t, sent, `"id":39`)
	// Empty optional fields must not be sent, or they overwrite real values.
	assert.NotContains(t, sent, `"phone"`)
}

func TestRealService_BulkImportStatus(t *testing.T) {
	require := require.New(t)

	d := &bodyCapturingDoer{
		Resp: &client.APIResponse{StatusCode: 200},
		Body: []byte(`{"success":1,"status":"completed","batchId":"b-123","contacts":{"total":250,"success":249,"failure":1}}`),
	}

	out, _, err := NewRealServiceFromDoer(d).BulkImportStatus(context.Background(), "b-123")
	require.NoError(err)
	assert.Equal(t, "import/info", d.Path)
	assert.Equal(t, map[string]string{"batchId": "b-123"}, d.Opts)
	assert.Equal(t, "completed", out.Status)
	assert.Equal(t, 250, out.Contacts.Total)
	assert.Equal(t, 1, out.Contacts.Failure)
}

func TestBulkImportMaxContacts(t *testing.T) {
	// Callers batch on this; if ActiveCampaign changes it the constant must
	// follow, so pin the documented value.
	assert.Equal(t, 250, BulkImportMaxContacts)
}
