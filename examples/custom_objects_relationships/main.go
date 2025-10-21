//go:build examples

package main

import (
	"encoding/json"
	"fmt"
)

// Small standalone example that demonstrates how to construct relationship
// payloads for custom objects. This file does not call the SDK; it simply
// prints the JSON shape the SDK will marshal when given the equivalent
// Go data structure.
func main() {
	// Relationships are encoded as map[string][]interface{} where each key is
	// the related namespace (for example "contacts") and the value is a list
	// of IDs (strings or numbers) to relate.
	rel := map[string][]interface{}{
		"contacts": []interface{}{"22"},
		"orders":   []interface{}{123, "456"},
	}

	payload := map[string]interface{}{"record": map[string]interface{}{"relationships": rel}}
	b, _ := json.MarshalIndent(payload, "", "  ")
	fmt.Println(string(b))
}
