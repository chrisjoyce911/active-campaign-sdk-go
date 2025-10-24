package contacts

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRealService_GetContactFieldValues(t *testing.T) {
	tests := []struct {
		name string
		id   string
		body []byte
	}{
		{name: "has fields", id: "1", body: []byte(`{"fieldValues":[{"field":"x","value":"y"}]}`)},
		{name: "empty", id: "2", body: []byte(`{"fieldValues":[]}`)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := &mockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: tt.body}
			require := require.New(t)
			require.NotNil(md)

			svc := NewRealServiceFromDoer(md)
			require.NotNil(svc)

			out, apiResp, err := svc.GetContactFieldValues(context.Background(), tt.id)
			assert.NoError(t, err)
			require.NotNil(apiResp)
			assert.Equal(t, 200, apiResp.StatusCode)
			_ = out // future: assert on typed shape
		})
	}
}
