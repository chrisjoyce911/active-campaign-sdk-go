package testhelpers

import (
	"context"
	"encoding/json"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// RecordingDoer records the incoming method, path and request body (v)
// so tests can assert what was sent. It also behaves like MockDoer and will
// unmarshal the configured Body into the out parameter and return Resp/Err.
type RecordingDoer struct {
	Resp *client.APIResponse
	Err  error
	Body []byte // response body to unmarshal into out

	LastMethod string
	LastPath   string
	LastV      interface{}
	LastBody   []byte // JSON-marshalled form of LastV when non-nil
}

func (r *RecordingDoer) Do(ctx context.Context, method, path string, v interface{}, out interface{}) (*client.APIResponse, error) {
	r.LastMethod = method
	r.LastPath = path
	r.LastV = v
	if v != nil {
		if b, err := json.Marshal(v); err == nil {
			r.LastBody = b
		}
	}
	if out != nil && r.Body != nil {
		_ = json.Unmarshal(r.Body, out)
	}
	return r.Resp, r.Err
}
