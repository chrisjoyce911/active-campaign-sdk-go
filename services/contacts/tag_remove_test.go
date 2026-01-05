package contacts

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRealService_TagRemove(t *testing.T) {
	tests := []struct {
		name         string
		contactTagID string
		statusCode   int
	}{
		{
			name:         "remove tag from contact",
			contactTagID: "1",
			statusCode:   200,
		},
		{
			name:         "tag not found",
			contactTagID: "999",
			statusCode:   404,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := &mockDoer{Resp: &client.APIResponse{StatusCode: tt.statusCode}}
			require := require.New(t)
			require.NotNil(md)

			svc := NewRealServiceFromDoer(md)
			require.NotNil(svc)

			apiResp, err := svc.TagRemove(context.Background(), tt.contactTagID)
			assert.NoError(t, err)
			require.NotNil(apiResp)
			assert.Equal(t, tt.statusCode, apiResp.StatusCode)
		})
	}
}
