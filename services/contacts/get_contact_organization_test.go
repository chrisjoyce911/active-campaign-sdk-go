package contacts

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRealService_GetContactOrganization(t *testing.T) {
	tests := []struct {
		name   string
		id     string
		body   []byte
		status int
	}{
		{name: "org", id: "1", body: []byte(`{"organization":{"id":"o1"}}`), status: 200},
		{name: "none", id: "2", body: []byte(`{"organization":null}`), status: 200},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := &mockDoer{Resp: &client.APIResponse{StatusCode: tt.status}, Body: tt.body}
			require := require.New(t)
			require.NotNil(md)

			svc := NewRealServiceFromDoer(md)
			require.NotNil(svc)

			out, apiResp, err := svc.GetContactOrganization(context.Background(), tt.id)
			assert.NoError(t, err)
			require.NotNil(apiResp)
			assert.Equal(t, tt.status, apiResp.StatusCode)
			_ = out
		})
	}
}
