package ecommerce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt64String_EdgeCases(t *testing.T) {
	var i Int64String

	// leading whitespace before quoted string -> should parse via the fallback string branch
	err := i.UnmarshalJSON([]byte("   \"789\""))
	assert.NoError(t, err)
	assert.Equal(t, Int64String(789), i)

	// empty quoted string -> should become zero
	err = i.UnmarshalJSON([]byte("\"\""))
	assert.NoError(t, err)
	assert.Equal(t, Int64String(0), i)

	// non-numeric quoted string -> should return an error
	err = i.UnmarshalJSON([]byte("\"abc\""))
	assert.Error(t, err)

	// non-quoted boolean -> should return an error (cannot unmarshal to int or string)
	err = i.UnmarshalJSON([]byte("true"))
	assert.Error(t, err)

	// non-quoted number with surrounding spaces -> should parse as number
	err = i.UnmarshalJSON([]byte("  321  "))
	assert.NoError(t, err)
	assert.Equal(t, Int64String(321), i)

	// empty input -> treated as zero
	err = i.UnmarshalJSON([]byte(""))
	assert.NoError(t, err)
	assert.Equal(t, Int64String(0), i)

	// fallback: leading spaces then empty quoted string -> should become zero via string branch
	err = i.UnmarshalJSON([]byte("   \"\""))
	assert.NoError(t, err)
	assert.Equal(t, Int64String(0), i)

	// malformed quoted JSON should return an error from json.Unmarshal
	err = i.UnmarshalJSON([]byte("\"\\x\""))
	assert.Error(t, err)

	// leading spaces then a quoted non-numeric string should return a parse error
	err = i.UnmarshalJSON([]byte("   \"abc\""))
	assert.Error(t, err)
}
