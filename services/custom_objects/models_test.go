package custom_objects

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimestampUnmarshal(t *testing.T) {
	var ts Timestamp
	err := json.Unmarshal([]byte(`"2020-01-02T15:04:05Z"`), &ts)
	assert.NoError(t, err)
	if assert.NotNil(t, ts.Time) {
		assert.Equal(t, 2020, ts.Time.Year())
	}

	// null
	err = json.Unmarshal([]byte(`null`), &ts)
	assert.NoError(t, err)
	assert.Nil(t, ts.Time)
}

func TestSchemaOptionsAndRelationships(t *testing.T) {
	raw := `{
        "id": "ot1",
        "fields": [{"id":"f1","options":[{"id":"o1","value":"v1"}] }],
        "relationships": [{"id":"rel1","namespace":"ns","hasMany":true}]
    }`
	var s Schema
	err := json.Unmarshal([]byte(raw), &s)
	assert.NoError(t, err)
	if assert.Len(t, s.Fields, 1) {
		assert.Len(t, s.Fields[0].Options, 1)
		assert.Equal(t, "o1", s.Fields[0].Options[0].ID)
	}
	if assert.Len(t, s.Relationships, 1) {
		assert.Equal(t, "rel1", s.Relationships[0].ID)
		assert.Equal(t, "ns", s.Relationships[0].Namespace)
		assert.True(t, s.Relationships[0].HasMany)
	}
}
