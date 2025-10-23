package custom_objects

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimestampUnmarshal_EmptyStringAndInvalid(t *testing.T) {
	var ts Timestamp

	// empty string should produce nil Time and no error
	err := json.Unmarshal([]byte(`""`), &ts)
	assert.NoError(t, err)
	assert.Nil(t, ts.Time)

	// invalid string should return an error
	err = json.Unmarshal([]byte(`"not-a-time"`), &ts)
	assert.Error(t, err)
}

func TestRecord_Unmarshal_InvalidFieldsType(t *testing.T) {
	// fields is a boolean which should cause both array and map unmarshal to fail
	raw := `{"id":"r-bad","fields":true}`
	var r Record
	err := json.Unmarshal([]byte(raw), &r)
	assert.Error(t, err)
}
