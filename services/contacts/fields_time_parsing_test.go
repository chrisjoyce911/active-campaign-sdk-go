package contacts

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestFieldPayload_TimeParsing(t *testing.T) {
	// Use various time formats observed
	body := []byte(`{"fields":[{"id":"f1","title":"X","cdate":"2022-11-02T02:21:01-05:00","udate":"2023-03-23T02:07:32-05:00","created_timestamp":"2022-11-02 02:21:01","updated_timestamp":"2023-03-23 02:07:32"}]}`)
	md := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: body}
	svc := NewRealServiceFromDoer(md)

	out, apiResp, err := svc.ListCustomFields(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, 200, apiResp.StatusCode)
	fields := out.FieldsOrEmpty()
	if assert.Len(t, fields, 1) {
		f := fields[0]
		// CDate and UDate should be parsed
		assert.False(t, f.CDate == nil)
		assert.False(t, f.UDate == nil)
		// CreatedTimestamp and UpdatedTimestamp should be parsed
		assert.False(t, f.CreatedTimestamp == nil)
		assert.False(t, f.UpdatedTimestamp == nil)
		// Check accessors
		_ = f.CDateOrZero()
		_ = f.UDateOrZero()
	}
}

func TestFieldOptionPayload_TimeParsing_Nulls(t *testing.T) {
	// cdate/udate null and missing cases
	raw := []byte(`{"fieldOptions":[{"id":"1","field":"28","orderid":"1","value":"NSW","label":"NSW","isdefault":"0","cdate":null,"udate":null}],"fields":[],"meta":{"total":"0"}}`)
	var parsed ListFieldsResponse
	if err := json.Unmarshal(raw, &parsed); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}
	// FieldOptions won't be directly on ListFieldsResponse; test FieldOptionPayload unmarshal
	var fo FieldOptionPayload
	// simulate single option
	optBody := []byte(`{"id":"1","field":"28","orderid":"1","value":"NSW","label":"NSW","isdefault":"0","cdate":null,"udate":null}`)
	if err := json.Unmarshal(optBody, &fo); err != nil {
		t.Fatalf("failed to unmarshal FieldOptionPayload: %v", err)
	}
	assert.Nil(t, fo.CDate)
	assert.Nil(t, fo.UDate)
	// Accessors return zero time
	assert.True(t, fo.CDateOrZero().IsZero())
	assert.True(t, fo.UDateOrZero().IsZero())

	// Test a valid RFC3339 timestamp
	optBody2 := []byte(`{"id":"2","field":"12","orderid":"1","value":"Individual","label":"Individual","isdefault":"0","cdate":"2022-11-02T02:23:04-05:00"}`)
	var fo2 FieldOptionPayload
	if err := json.Unmarshal(optBody2, &fo2); err != nil {
		t.Fatalf("failed to unmarshal FieldOptionPayload2: %v", err)
	}
	if assert.NotNil(t, fo2.CDate) {
		// ensure parsed time is not zero
		assert.False(t, fo2.CDate.IsZero())
		// sanity: ensure it's within reasonable year range
		assert.True(t, fo2.CDate.Year() >= 2000)
	}
	// UDate remains nil
	assert.Nil(t, fo2.UDate)
}
