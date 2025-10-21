package custom_objects

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/stretchr/testify/assert"
)

// reuse capturingDoer from the previous test file (but redeclare here to keep tests self-contained)
type capturingDoer2 struct {
	Resp    *client.APIResponse
	Err     error
	Method  string
	Path    string
	Payload json.RawMessage
}

func (c *capturingDoer2) Do(ctx context.Context, method, path string, v interface{}, out interface{}) (*client.APIResponse, error) {
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

func TestUpdateObjectRecord_RelationshipsAsArray(t *testing.T) {
	c := &capturingDoer2{Resp: &client.APIResponse{StatusCode: 201, Body: []byte(`{"record":{"id":"r1"}}`)}}
	svc := NewRealServiceFromDoer(c)

	// relationships should be map[string][]interface{}
	req := &UpdateRecordRequest{Fields: map[string]interface{}{"name": "x"}, Relationships: map[string][]interface{}{"primary-contact": {42}}}
	out, apiResp, err := svc.UpdateObjectRecord(context.Background(), "schema-xyz", "r1", req)
	assert.NoError(t, err)
	assert.NotNil(t, apiResp)
	assert.Equal(t, 201, apiResp.StatusCode)

	// inspect payload
	var body map[string]json.RawMessage
	assert.NoError(t, json.Unmarshal(c.Payload, &body))
	var rec map[string]json.RawMessage
	assert.NoError(t, json.Unmarshal(body["record"], &rec))

	// relationships should be present and be an object mapping to arrays
	var rels map[string][]interface{}
	assert.NoError(t, json.Unmarshal(rec["relationships"], &rels))
	v, ok := rels["primary-contact"]
	assert.True(t, ok)
	assert.Len(t, v, 1)
	assert.Equal(t, float64(42), v[0]) // JSON numbers decode as float64

	assert.Equal(t, "r1", out.Record.ID)
}
