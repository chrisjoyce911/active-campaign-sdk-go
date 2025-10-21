package client

import (
	"context"
)

// Doer is a minimal interface around the CoreClient.Do method so callers can
// inject test doubles in unit tests.
type Doer interface {
	Do(ctx context.Context, method, path string, v interface{}, out interface{}) (*APIResponse, error)
}
