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

func TestService_GetCampaign_table(t *testing.T) {
	tests := []struct {
		name     string
		mockResp *client.APIResponse
		mockBody []byte
		ctx      context.Context
		id       string
		wantErr  bool
	}{
		{name: "found", mockResp: &client.APIResponse{StatusCode: 200}, mockBody: []byte(`{"campaign":{"id":"c1"}}`), ctx: context.Background(), id: "c1", wantErr: false},
		{name: "bad request", mockResp: &client.APIResponse{StatusCode: 400}, mockBody: []byte(`{}`), ctx: context.Background(), id: "", wantErr: false},
		{name: "doer error", mockResp: nil, mockBody: nil, ctx: func() context.Context { c, cancel := context.WithCancel(context.Background()); cancel(); return c }(), id: "1", wantErr: true},
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

			out, apiResp, err := svc.GetCampaign(tc.ctx, tc.id)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			require.NotNil(apiResp)
			assert.Equal(t, tc.mockResp.StatusCode, apiResp.StatusCode)
			_ = out
		})
	}
}
