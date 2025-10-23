package campaigns

import (
	"context"
	"fmt"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	th "github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestService_ListCampaigns_table(t *testing.T) {
	tests := []struct {
		name     string
		mockResp *client.APIResponse
		mockBody []byte
		opts     map[string]string
		wantErr  bool
	}{
		{name: "ok", mockResp: &client.APIResponse{StatusCode: 200}, mockBody: []byte(`{"campaigns":[{"id":"c1"}]}`), opts: map[string]string{"limit": "1"}, wantErr: false},
		{name: "doer error", mockResp: nil, mockBody: nil, opts: nil, wantErr: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			md := &th.MockDoer{Resp: tc.mockResp, Body: tc.mockBody}
			if tc.name == "doer error" {
				md = &th.MockDoer{Err: fmt.Errorf("boom")}
			}
			svc := NewRealServiceFromDoer(md)

			out, apiResp, err := svc.ListCampaigns(context.Background(), tc.opts)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			_ = out
			if apiResp != nil {
				assert.Equal(t, tc.mockResp.StatusCode, apiResp.StatusCode)
			}
		})
	}
}
