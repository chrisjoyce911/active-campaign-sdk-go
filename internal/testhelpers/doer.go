package testhelpers

import (
	"context"
	"encoding/json"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// MockDoer is a simple test double implementing client.Doer. It returns a
// canned APIResponse and optional JSON body which will be unmarshalled into
// the provided out parameter when non-nil.
type MockDoer struct {
	Resp *client.APIResponse
	Err  error
	Body []byte
}

func (m *MockDoer) Do(ctx context.Context, method, path string, v interface{}, out interface{}) (*client.APIResponse, error) {
	if out != nil && m.Body != nil {
		_ = json.Unmarshal(m.Body, out)
	}
	return m.Resp, m.Err
}
