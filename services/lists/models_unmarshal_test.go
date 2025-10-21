package lists

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalMixedTypes(t *testing.T) {
	cases := []struct {
		name    string
		jsonStr string
	}{
		{"numbers", `{"list":{"id":"1","send_last_broadcast":0,"userid":5,"non_deleted_subscribers":10,"active_subscribers":7,"created_by":2}}`},
		{"strings", `{"list":{"id":"1","send_last_broadcast":"0","userid":"5","non_deleted_subscribers":"10","active_subscribers":"7","created_by":"2"}}`},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var out GetListResponse
			err := json.Unmarshal([]byte(c.jsonStr), &out)
			if !assert.NoError(t, err) {
				return
			}
			// Ensure fields are populated and canonicalized
			if assert.NotNil(t, out.List) {
				// send_last_broadcast as string
				s := string(out.List.SendLastBroadcast)
				assert.Equal(t, "0", s)
				// userid
				assert.Equal(t, StringOrInt("5"), out.List.UserID)
				// non_deleted_subscribers and active_subscribers integer values
				assert.Equal(t, IntOrString(10), out.List.NonDeletedSubscribers)
				assert.Equal(t, IntOrString(7), out.List.ActiveSubscribers)
				// created_by
				assert.Equal(t, StringOrInt("2"), out.List.CreatedBy)
			}
		})
	}
}

func TestStringOrIntUnmarshalVariants(t *testing.T) {
	var s StringOrInt

	// null -> empty
	err := json.Unmarshal([]byte("null"), &s)
	assert.NoError(t, err)
	assert.Equal(t, StringOrInt(""), s)

	// quoted string
	err = json.Unmarshal([]byte("\"abc\""), &s)
	assert.NoError(t, err)
	assert.Equal(t, StringOrInt("abc"), s)

	// number
	err = json.Unmarshal([]byte("123"), &s)
	assert.NoError(t, err)
	assert.Equal(t, StringOrInt("123"), s)
}

func TestIntOrStringUnmarshalVariants(t *testing.T) {
	var i IntOrString

	// null -> zero
	err := json.Unmarshal([]byte("null"), &i)
	assert.NoError(t, err)
	assert.Equal(t, IntOrString(0), i)

	// number
	err = json.Unmarshal([]byte("42"), &i)
	assert.NoError(t, err)
	assert.Equal(t, IntOrString(42), i)

	// quoted numeric string
	err = json.Unmarshal([]byte("\"7\""), &i)
	assert.NoError(t, err)
	assert.Equal(t, IntOrString(7), i)

	// quoted non-numeric should error
	err = json.Unmarshal([]byte("\"notanint\""), &i)
	assert.Error(t, err)
}

func TestUnmarshalErrorCases(t *testing.T) {
	var s StringOrInt
	// invalid JSON for number/string: object
	err := s.UnmarshalJSON([]byte("{}"))
	assert.Error(t, err)

	var i IntOrString
	// floating point number should fail Atoi path
	err = i.UnmarshalJSON([]byte("3.14"))
	assert.Error(t, err)
}
