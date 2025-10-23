package custom_objects

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimestampUnmarshal_AlternateLayouts(t *testing.T) {
	var ts Timestamp

	// RFC3339Nano
	err := json.Unmarshal([]byte(`"2020-01-02T15:04:05.123456789Z"`), &ts)
	assert.NoError(t, err)
	if assert.NotNil(t, ts.Time) {
		assert.Equal(t, 2020, ts.Time.Year())
	}

	// space-separated layout
	err = json.Unmarshal([]byte(`"2006-01-02 15:04:05"`), &ts)
	assert.NoError(t, err)
	if assert.NotNil(t, ts.Time) {
		assert.Equal(t, 2006, ts.Time.Year())
	}
}

func TestListSchemasResponse_Unmarshal_SchemasKey(t *testing.T) {
	raw := `{"schemas":[{"id":"ot1","slug":"s"}],"meta":{"total":2}}`
	var lr ListSchemasResponse
	err := json.Unmarshal([]byte(raw), &lr)
	assert.NoError(t, err)
	assert.Len(t, lr.Schemas, 1)
	assert.Equal(t, 2, lr.Meta.Total)
}

func TestRecord_Unmarshal_NoFields_And_NumericID(t *testing.T) {
	// no fields key
	rawNoFields := `{"id":"r-no-fields"}`
	var r Record
	err := json.Unmarshal([]byte(rawNoFields), &r)
	assert.NoError(t, err)
	assert.Nil(t, r.Fields)

	// numeric ID should not populate string ID
	rawNumericID := `{"id":123,"fields":{"a":"b"}}`
	var r2 Record
	err = json.Unmarshal([]byte(rawNumericID), &r2)
	assert.NoError(t, err)
	// ID cannot be decoded into string, so it should remain empty
	assert.Equal(t, "", r2.ID)
	if assert.NotNil(t, r2.Fields) {
		assert.Equal(t, "b", r2.Fields["a"])
	}
}

func TestCreateRecordRequest_MarshalJSON_FieldsArray(t *testing.T) {
	ext := "ext-1"
	req := CreateRecordRequest{
		ID:         "rid",
		ExternalID: &ext,
		Fields: map[string]interface{}{
			"f1": "v1",
			"f2": 2,
		},
		Relationships: map[string][]interface{}{
			"rel": {"x"},
		},
	}
	b, err := json.Marshal(&req)
	assert.NoError(t, err)

	// ensure fields was marshalled as an array of objects with id/value
	var out map[string]interface{}
	err = json.Unmarshal(b, &out)
	assert.NoError(t, err)
	// id and externalId presence
	assert.Equal(t, "rid", out["id"])
	assert.Equal(t, ext, out["externalId"])

	// fields should be an array
	fields, ok := out["fields"].([]interface{})
	if assert.True(t, ok) {
		// convert to map for easier assertions
		got := make(map[string]interface{})
		for _, it := range fields {
			if m, ok := it.(map[string]interface{}); ok {
				k, _ := m["id"].(string)
				got[k] = m["value"]
			}
		}
		assert.Equal(t, "v1", got["f1"])
		// numbers are marshalled as float64
		assert.Equal(t, float64(2), got["f2"])
	}
}

func TestTimestamp_Unmarshal_FallbackParsesRFC3339WhenLayoutsOmitted(t *testing.T) {
	// preserve original
	orig := timestampLayouts
	defer func() { timestampLayouts = orig }()
	// remove RFC3339 from the layouts so the loop won't succeed and fallback is used
	timestampLayouts = []string{time.RFC3339Nano}

	var ts Timestamp
	// this is RFC3339 (no nano) which will fail the loop but succeed in fallback
	err := json.Unmarshal([]byte(`"2020-01-02T15:04:05Z"`), &ts)
	assert.NoError(t, err)
	if assert.NotNil(t, ts.Time) {
		assert.Equal(t, 2020, ts.Time.Year())
	}
}

func TestListSchemasResponse_Unmarshal_MetaOnly(t *testing.T) {
	raw := `{"meta":{"total":5}}`
	var lr ListSchemasResponse
	err := json.Unmarshal([]byte(raw), &lr)
	assert.NoError(t, err)
	assert.Len(t, lr.Schemas, 0)
	assert.Equal(t, 5, lr.Meta.Total)
}

func TestListSchemasResponse_Unmarshal_SchemasBadType(t *testing.T) {
	raw := `{"schemas":"oops"}`
	var lr ListSchemasResponse
	err := json.Unmarshal([]byte(raw), &lr)
	assert.Error(t, err)
}

func TestRecord_Unmarshal_EmptyArrayFields(t *testing.T) {
	raw := `{"id":"r-empty","fields":[]}`
	var r Record
	err := json.Unmarshal([]byte(raw), &r)
	assert.NoError(t, err)
	// should be an empty map
	if assert.NotNil(t, r.Fields) {
		assert.Len(t, r.Fields, 0)
	}
}

func TestTimestamp_Fallback_OnlySpaceLayout(t *testing.T) {
	orig := timestampLayouts
	defer func() { timestampLayouts = orig }()
	// only allow space-separated layout so RFC3339 must be parsed by fallback
	timestampLayouts = []string{"2006-01-02 15:04:05"}

	var ts Timestamp
	err := json.Unmarshal([]byte(`"2020-01-02T15:04:05Z"`), &ts)
	assert.NoError(t, err)
	if assert.NotNil(t, ts.Time) {
		assert.Equal(t, 2020, ts.Time.Year())
	}
}

func TestRecord_Unmarshal_DirectUnmarshalWithNonObjectJSON(t *testing.T) {
	var r Record
	// call UnmarshalJSON directly with a JSON string value (valid JSON) but not an object
	err := r.UnmarshalJSON([]byte(`"just-a-string"`))
	assert.Error(t, err)
}

func TestListSchemasResponse_Unmarshal_InvalidRawDirect(t *testing.T) {
	var lr ListSchemasResponse
	err := lr.UnmarshalJSON([]byte(`not-json`))
	assert.Error(t, err)
}
