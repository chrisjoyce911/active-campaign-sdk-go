package contacts

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/stretchr/testify/assert"
)

func TestRealService_BulkImportContacts(t *testing.T) {
	tests := []struct {
		name   string
		body   []byte
		status int
	}{
		{name: "accepted", body: []byte(`{"job":{"id":"j1"}}`), status: 202},
		{name: "bad", body: []byte(`{"error":"bad"}`), status: 400},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := &mockDoer{Resp: &client.APIResponse{StatusCode: tt.status}, Body: tt.body}
			svc := NewRealServiceFromDoer(md)

			out, apiResp, err := svc.BulkImportContacts(context.Background(), map[string]interface{}{"foo": "bar"})
			if tt.status >= 400 {
				_ = err
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.status, apiResp.StatusCode)
				_ = out
			}
		})
	}
}
