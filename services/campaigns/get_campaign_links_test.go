package campaigns

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	th "github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_GetCampaignLinks_table(t *testing.T) {
	tests := []struct {
		name     string
		mockResp *client.APIResponse
		mockBody []byte
		id       string
		wantErr  bool
	}{
		{name: "ok", mockResp: &client.APIResponse{StatusCode: 200}, mockBody: []byte(`{"links":[{"campaignid":"1","messageid":"3","link":"open","name":"Read Tracker","ref":"","tracked":"1","links":{"campaign":"https://:account.api-us1.com/api/3/links/1/campaign","message":"https://:account.api-us1.com/api/3/links/1/message"},"id":"1","campaign":"1","message":"3"}]}`), id: "1", wantErr: false},
		{name: "not found", mockResp: &client.APIResponse{StatusCode: 404}, mockBody: []byte(`{"message":"No Result found for Campaign with id 2"}`), id: "2", wantErr: false},
		{name: "doer error", mockResp: nil, mockBody: nil, id: "x", wantErr: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			md := &th.MockDoer{Resp: tc.mockResp, Body: tc.mockBody}
			if tc.name == "doer error" {
				md = &th.MockDoer{Err: assert.AnError}
			}
			require := require.New(t)
			require.NotNil(md)

			svc := NewRealServiceFromDoer(md)
			require.NotNil(svc)

			out, apiResp, err := svc.GetCampaignLinks(context.Background(), tc.id)
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
