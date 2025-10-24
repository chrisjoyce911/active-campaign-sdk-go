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

func TestCreateCampaign_Mocked(t *testing.T) {
	tests := []struct {
		name     string
		mockResp *client.APIResponse
		mockBody []byte
		wantErr  bool
	}{
		{name: "created", mockResp: &client.APIResponse{StatusCode: 201}, mockBody: []byte(`{"campaign":{"id":"c1"}}`), wantErr: false},
		{name: "doer error", mockResp: nil, mockBody: nil, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := &th.MockDoer{Resp: tt.mockResp, Body: tt.mockBody}
			if tt.name == "doer error" {
				md = &th.MockDoer{Err: fmt.Errorf("boom")}
			}
			svc := NewRealServiceFromDoer(md)
			require.NotNil(t, svc)

			req := &CreateCampaignRequest{Name: "x", Type: "single"}
			out, apiResp, err := svc.CreateCampaign(context.Background(), req)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			if apiResp != nil {
				assert.Equal(t, 201, apiResp.StatusCode)
			}
			_ = out
		})
	}
}
