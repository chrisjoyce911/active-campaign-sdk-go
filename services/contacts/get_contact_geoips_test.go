package contacts

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/stretchr/testify/assert"
)

func TestRealService_GetContactGeoIPs(t *testing.T) {
	tests := []struct {
		name   string
		id     string
		body   []byte
		status int
	}{
		{name: "listed", id: "1", body: []byte(`{"geoips":[{"ip":"1.2.3.4"}]}`), status: 200},
		{name: "none", id: "2", body: []byte(`{"geoips":[]}`), status: 200},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := &mockDoer{Resp: &client.APIResponse{StatusCode: tt.status}, Body: tt.body}
			svc := NewRealServiceFromDoer(md)

			out, apiResp, err := svc.GetContactGeoIPs(context.Background(), tt.id)
			assert.NoError(t, err)
			assert.Equal(t, tt.status, apiResp.StatusCode)
			_ = out
		})
	}
}
