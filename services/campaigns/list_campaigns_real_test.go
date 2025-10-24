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

func TestRealService_ListCampaigns(t *testing.T) {
	tests := []struct {
		name       string
		mockResp   *client.APIResponse
		mockBody   []byte
		opts       map[string]string
		wantStatus int
		wantErr    bool
	}{
		{name: "ok", mockResp: &client.APIResponse{StatusCode: 200}, mockBody: []byte(`{"campaigns":[{"id":"c1"}]}`), opts: map[string]string{"limit": "1"}, wantStatus: 200, wantErr: false},
		{name: "server error", mockResp: &client.APIResponse{StatusCode: 500}, mockBody: []byte(`{}`), opts: nil, wantStatus: 500, wantErr: false},
		{name: "doer error", mockResp: nil, mockBody: nil, opts: nil, wantStatus: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := &th.MockDoer{Resp: tt.mockResp, Body: tt.mockBody}
			if tt.name == "doer error" {
				md = &th.MockDoer{Err: fmt.Errorf("boom")}
			}
			require := require.New(t)
			require.NotNil(md)

			svc := NewRealServiceFromDoer(md)
			require.NotNil(svc)

			out, apiResp, err := svc.ListCampaigns(context.Background(), tt.opts)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			require.NotNil(apiResp)
			assert.Equal(t, tt.wantStatus, apiResp.StatusCode)
			_ = out
		})
	}
}
