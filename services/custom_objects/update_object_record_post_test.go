package custom_objects

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/stretchr/testify/assert"
)

// capturingDoer records the last call parameters for inspection.
type capturingDoer struct {
	Resp *client.APIResponse
	Err  error
	// captured
	Method  string
	Path    string
	Payload json.RawMessage
}

func (c *capturingDoer) Do(ctx context.Context, method, path string, v interface{}, out interface{}) (*client.APIResponse, error) {
	c.Method = method
	c.Path = path
	if v != nil {
		if b, err := json.Marshal(v); err == nil {
			c.Payload = b
		}
	}
	if out != nil && c.Resp != nil && len(c.Resp.Body) > 0 {
		_ = json.Unmarshal(c.Resp.Body, out)
	}
	return c.Resp, c.Err
}

func TestUpdateObjectRecord_UsesPostCreateOrUpdate(t *testing.T) {
	c := &capturingDoer{Resp: &client.APIResponse{StatusCode: 201, Body: []byte(`{"record":{"id":"r1","fields":[{"id":"name","value":"demo-updated"}]}}`)}}
	svc := NewRealServiceFromDoer(c)

	req := &UpdateRecordRequest{Fields: map[string]interface{}{"name": "demo-updated"}}
	out, apiResp, err := svc.UpdateObjectRecord(context.Background(), "schema-123", "r1", req)

	assert.NoError(t, err)
	assert.NotNil(t, apiResp)
	assert.Equal(t, 201, apiResp.StatusCode)

	// Verify the service used POST to the create-or-update endpoint (schema-level)
	assert.Equal(t, "POST", strings.ToUpper(c.Method))
	assert.True(t, strings.HasPrefix(c.Path, "customObjects/records/"))

	// Payload should include a top-level "record" with id and fields array
	var body map[string]json.RawMessage
	err = json.Unmarshal(c.Payload, &body)
	assert.NoError(t, err)
	var rec json.RawMessage
	ok := false
	if v, found := body["record"]; found {
		rec = v
		ok = true
	}
	assert.True(t, ok, "expected payload to contain 'record' key")

	var recObj map[string]json.RawMessage
	err = json.Unmarshal(rec, &recObj)
	assert.NoError(t, err)

	// fields should be an array of objects with id/value
	var fields []struct {
		ID    string      `json:"id"`
		Value interface{} `json:"value"`
	}
	if f, found := recObj["fields"]; found {
		err = json.Unmarshal(f, &fields)
		assert.NoError(t, err)
		assert.Len(t, fields, 1)
		assert.Equal(t, "name", fields[0].ID)
		assert.Equal(t, "demo-updated", fields[0].Value)
	} else {
		t.Fatalf("expected fields key in record payload")
	}

	// ensure returned record parsed correctly
	assert.Equal(t, "r1", out.Record.ID)
}
