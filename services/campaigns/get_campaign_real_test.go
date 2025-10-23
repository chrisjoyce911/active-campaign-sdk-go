package campaigns

import (
	"context"
	"fmt"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	th "github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestRealService_GetCampaign(t *testing.T) {
	tests := []struct {
		name       string
		mockResp   *client.APIResponse
		mockBody   []byte
		id         string
		wantStatus int
		wantErr    bool
	}{
		{name: "found", mockResp: &client.APIResponse{StatusCode: 200}, mockBody: []byte(`{"campaign":{"id":"c1"}}`), id: "c1", wantStatus: 200, wantErr: false},
		{name: "not found", mockResp: &client.APIResponse{StatusCode: 404}, mockBody: []byte(`{}`), id: "nope", wantStatus: 404, wantErr: false},
		{name: "doer error", mockResp: nil, mockBody: nil, id: "err", wantStatus: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := &th.MockDoer{Resp: tt.mockResp, Body: tt.mockBody}
			if tt.name == "doer error" {
				md = &th.MockDoer{Err: errExample}
			}
			svc := NewRealServiceFromDoer(md)

			out, apiResp, err := svc.GetCampaign(context.Background(), tt.id)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			if apiResp != nil {
				assert.Equal(t, tt.wantStatus, apiResp.StatusCode)
			}
			_ = out
		})
	}
}

// errExample is a sentinel error used by MockDoer tests.
var errExample = fmt.Errorf("boom")
