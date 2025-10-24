package contacts

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListCustomFields_Unmarshal_HappyPath(t *testing.T) {
	body := []byte(`{"fields":[{"id":"f1","title":"X","type":"text"}],"meta":{"total":"1"}}`)
	md := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: body}
	require := require.New(t)
	require.NotNil(md)

	svc := NewRealServiceFromDoer(md)
	require.NotNil(svc)

	out, apiResp, err := svc.ListCustomFields(context.Background())
	assert.NoError(t, err)
	require.NotNil(apiResp)
	assert.Equal(t, 200, apiResp.StatusCode)
	if assert.NotNil(t, out) {
		// Fields should be present and FieldsOrEmpty should return the item
		fields := out.FieldsOrEmpty()
		if assert.Len(t, fields, 1) {
			assert.Equal(t, "f1", fields[0].ID)
			assert.Equal(t, "X", fields[0].Title)
			assert.Equal(t, "text", fields[0].Type)
		}
	}
}

func TestListCustomFields_Unmarshal_MissingFieldsKey(t *testing.T) {
	// body missing the 'fields' key entirely
	body := []byte(`{"meta":{"total":"0"}}`)
	md := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: body}
	require := require.New(t)
	require.NotNil(md)

	svc := NewRealServiceFromDoer(md)
	require.NotNil(svc)

	out, apiResp, err := svc.ListCustomFields(context.Background())
	assert.NoError(t, err)
	require.NotNil(apiResp)
	assert.Equal(t, 200, apiResp.StatusCode)
	if assert.NotNil(t, out) {
		// Fields may be nil, but FieldsOrEmpty should return empty slice
		assert.Nil(t, out.Fields)
		fields := out.FieldsOrEmpty()
		assert.NotNil(t, fields)
		assert.Len(t, fields, 0)
	}

	// Also ensure we can unmarshal raw JSON into the struct type directly
	var parsed ListFieldsResponse
	if err := json.Unmarshal(body, &parsed); err != nil {
		t.Fatalf("failed to unmarshal raw body: %v", err)
	}
	// parsed.Fields should be nil and FieldsOrEmpty returns empty
	assert.Nil(t, parsed.Fields)
	assert.Len(t, parsed.FieldsOrEmpty(), 0)
}
