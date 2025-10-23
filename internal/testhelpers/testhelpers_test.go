package testhelpers

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/stretchr/testify/assert"
)

func TestMockDoer_UnmarshalBody(t *testing.T) {
	body := []byte(`{"foo":"bar"}`)
	md := &MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: body}
	var out map[string]string
	resp, err := md.Do(context.Background(), "GET", "p", nil, &out)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "bar", out["foo"])
}

func TestRecordingDoer_RecordsInputs(t *testing.T) {
	r := &RecordingDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"a":1}`)}
	var out map[string]int
	resp, err := r.Do(context.Background(), "POST", "path/x", map[string]string{"x": "y"}, &out)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	// ensure recording fields are populated
	assert.Equal(t, "POST", r.LastMethod)
	assert.Equal(t, "path/x", r.LastPath)
	// ensure LastBody was marshalled
	var mb map[string]string
	_ = json.Unmarshal(r.LastBody, &mb)
	assert.Equal(t, "y", mb["x"])
}
