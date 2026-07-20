package contacts

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// pathRecordingDoer captures the request path so tests can assert the
// endpoint — contact/sync (singular) upserts; contacts/sync would hit the
// plain create endpoint and 422 on existing emails.
type pathRecordingDoer struct {
	Resp *client.APIResponse
	Body []byte
	Path string
}

func (m *pathRecordingDoer) Do(ctx context.Context, method, path string, v interface{}, out interface{}) (*client.APIResponse, error) {
	m.Path = path
	if out != nil && m.Body != nil {
		_ = json.Unmarshal(m.Body, out)
	}
	return m.Resp, nil
}

func TestRealService_SyncContact(t *testing.T) {
	require := require.New(t)

	body := []byte(`{"contact":{"id":"c1","email":"a@b.com","cdate":"2026-01-01T00:00:00-05:00","udate":"2026-01-01T00:00:00-05:00"},"fieldValues":[{"field":"39","value":"USI123"}]}`)
	md := &pathRecordingDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: body}

	svc := NewRealServiceFromDoer(md)
	require.NotNil(svc)

	fv := []FieldValue{{Field: "39", Value: "USI123"}}
	req := &CreateContactRequest{Contact: &Contact{Email: "a@b.com", FieldValues: &fv}}

	out, apiResp, err := svc.SyncContact(context.Background(), req)
	require.NoError(err)
	require.NotNil(apiResp)
	assert.Equal(t, 200, apiResp.StatusCode)

	assert.Equal(t, "contact/sync", md.Path, "must use the singular sync endpoint — contacts/sync is plain create and 422s on existing emails")

	require.NotNil(out)
	require.NotNil(out.Contact)
	assert.Equal(t, "c1", out.Contact.ID)
	assert.Equal(t, out.Contact.CDate, out.Contact.UDate, "freshly created contact has cdate == udate")
	require.NotNil(out.FieldValues)
	assert.Equal(t, "USI123", (*out.FieldValues)[0].Value)
}
