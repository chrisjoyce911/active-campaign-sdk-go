package contacts

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRealService_GetContactGeoIP(t *testing.T) {
	tests := []struct {
		name   string
		id     string
		ip     string
		body   []byte
		status int
	}{
		{name: "found", id: "1", ip: "1.2.3.4", body: []byte(`{"geoip":{"ip":"1.2.3.4"}}`), status: 200},
		{name: "none", id: "2", ip: "9.9.9.9", body: []byte(`{"geoip":null}`), status: 200},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := &mockDoer{Resp: &client.APIResponse{StatusCode: tt.status}, Body: tt.body}
			require := require.New(t)
			require.NotNil(md)

			svc := NewRealServiceFromDoer(md)
			require.NotNil(svc)

			out, apiResp, err := svc.GetContactGeoIP(context.Background(), tt.id, tt.ip)
			assert.NoError(t, err)
			require.NotNil(apiResp)
			assert.Equal(t, tt.status, apiResp.StatusCode)
			_ = out
		})
	}
}
