package custom_objects

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecord_Unmarshal_ArrayAndMapForms(t *testing.T) {
	// array form
	arr := `{"id":"r1","fields":[{"id":"f1","value":"v1"},{"id":"f2","value":2}]}`
	var r Record
	err := json.Unmarshal([]byte(arr), &r)
	assert.NoError(t, err)
	if assert.NotNil(t, r.Fields) {
		assert.Equal(t, "v1", r.Fields["f1"])
		assert.Equal(t, float64(2), r.Fields["f2"])
	}

	// map form
	mp := `{"id":"r2","fields":{"f3":"x","f4":4}}`
	var r2 Record
	err = json.Unmarshal([]byte(mp), &r2)
	assert.NoError(t, err)
	if assert.NotNil(t, r2.Fields) {
		assert.Equal(t, "x", r2.Fields["f3"])
		assert.Equal(t, float64(4), r2.Fields["f4"])
	}
}

func TestListSchemasResponse_Unmarshal_LegacyKey(t *testing.T) {
	raw := `{"objectTypes":[{"id":"ot1","slug":"s"}],"meta":{"total":1}}`
	var lr ListSchemasResponse
	err := json.Unmarshal([]byte(raw), &lr)
	assert.NoError(t, err)
	assert.Len(t, lr.Schemas, 1)
	assert.Equal(t, 1, lr.Meta.Total)
}
