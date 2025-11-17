//go:build examples

package main

import (
	"os"
	"testing"
)

func TestMain_GetLists(t *testing.T) {
	oldURL := os.Getenv("ACTIVE_URL")
	oldTok := os.Getenv("ACTIVE_TOKEN")
	t.Cleanup(func() {
		_ = os.Setenv("ACTIVE_URL", oldURL)
		_ = os.Setenv("ACTIVE_TOKEN", oldTok)
	})
	_ = os.Setenv("ACTIVE_URL", "http://example.com")
	_ = os.Setenv("ACTIVE_TOKEN", "test-token")

	main()
}
