package genconstants

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenderConsts_collision_suffix(t *testing.T) {
	tc := struct {
		name string
		kvs  []KV
	}{
		name: "collisions",
		kvs:  []KV{{Key: "Example", Value: "1"}, {Key: "Example", Value: "2"}},
	}

	t.Run(tc.name, func(t *testing.T) {
		buf := &bytes.Buffer{}
		mapping := map[string]string{}
		updated := map[string]string{}
		renderConsts(buf, "Tag", tc.kvs, mapping, updated)

		out := buf.String()
		assert.Contains(t, out, "type TagsType")
		assert.Contains(t, out, "var Tags = TagsType{")
		assert.Contains(t, out, "\"1\":")
		assert.Contains(t, out, "\"2\":")
		assert.Equal(t, 2, len(updated))
	})
}
