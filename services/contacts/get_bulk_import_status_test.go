package contacts

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRealService_GetBulkImportStatus(t *testing.T) {
	tests := []struct {
		name   string
		id     string
		body   []byte
		status int
	}{
		{name: "found", id: "j1", body: []byte(`{"job":{"id":"j1"}}`), status: 200},
		{name: "notfound", id: "bad", body: []byte(`{"error":"not found"}`), status: 404},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := &mockDoer{Resp: &client.APIResponse{StatusCode: tt.status}, Body: tt.body}
			require := require.New(t)
			require.NotNil(md)

			svc := NewRealServiceFromDoer(md)
			require.NotNil(svc)

			out, apiResp, err := svc.GetBulkImportStatus(context.Background(), tt.id)
			if tt.status >= 400 {
				_ = err
			} else {
				assert.NoError(t, err)
				require.NotNil(apiResp)
				assert.Equal(t, tt.status, apiResp.StatusCode)
				_ = out
			}
		})
	}
}
