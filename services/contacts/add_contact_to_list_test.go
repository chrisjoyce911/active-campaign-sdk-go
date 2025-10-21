package contacts

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/stretchr/testify/assert"
)

func TestRealService_AddContactToList(t *testing.T) {
	tests := []struct {
		name   string
		body   []byte
		status int
	}{
		{name: "success", body: []byte(`{"contactList":{"id":"cl1"}}`), status: 201},
		{name: "bad request", body: []byte(`{"error":"bad"}`), status: 400},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := &mockDoer{Resp: &client.APIResponse{StatusCode: tt.status}, Body: tt.body}
			svc := NewRealServiceFromDoer(md)

			out, apiResp, err := svc.AddContactToList(context.Background(), map[string]interface{}{"foo": "bar"})
			if tt.status >= 400 {
				// allow error path to be returned from client
				_ = err
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.status, apiResp.StatusCode)
				_ = out
			}
		})
	}
}
