package campaigns

import (
	"context"
	"fmt"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	th "github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_CreateCampaign_table(t *testing.T) {
	tests := []struct {
		name     string
		mockResp *client.APIResponse
		mockBody []byte
		req      interface{}
		wantErr  bool
	}{
		{name: "normal", mockResp: &client.APIResponse{StatusCode: 201}, mockBody: []byte(`{"campaign":{"id":"c1"}}`), req: map[string]interface{}{"title": "t"}, wantErr: false},
		{name: "doer error", mockResp: nil, mockBody: nil, req: nil, wantErr: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			md := &th.MockDoer{Resp: tc.mockResp, Body: tc.mockBody}
			if tc.name == "doer error" {
				md = &th.MockDoer{Err: fmt.Errorf("boom")}
			}
			require := require.New(t)
			require.NotNil(md)

			svc := NewRealServiceFromDoer(md)
			require.NotNil(svc)

			var req *CreateCampaignRequest
			if tc.req != nil {
				if m, ok := tc.req.(map[string]interface{}); ok {
					// map-based tests pass a map; convert minimal fields
					name, _ := m["title"].(string)
					typ, _ := m["type"].(string)
					req = &CreateCampaignRequest{Name: name, Type: typ}
				}
			}
			out, apiResp, err := svc.CreateCampaign(context.Background(), req)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			require.NotNil(apiResp)
			assert.Equal(t, 201, apiResp.StatusCode)
			_ = out
		})
	}
}
