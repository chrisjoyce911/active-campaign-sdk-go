package lists

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestGetList(t *testing.T) {
	tests := []struct {
		name       string
		mockResp   *client.APIResponse
		mockBody   []byte
		id         string
		wantStatus int
		wantErr    bool
	}{
		{name: "ok", mockResp: &client.APIResponse{StatusCode: 200}, mockBody: []byte(`{"list":{"id":"1","name":"L1"}}`), id: "1", wantStatus: 200, wantErr: false},
		{name: "err-404", mockResp: &client.APIResponse{StatusCode: 404}, mockBody: []byte(`{"error":"not found"}`), id: "999", wantStatus: 404, wantErr: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			md := &testhelpers.MockDoer{Resp: tc.mockResp, Body: tc.mockBody}
			if tc.wantErr {
				md.Err = assert.AnError
			}
			svc := NewRealServiceFromDoer(md)

			out, apiResp, err := svc.GetList(context.Background(), tc.id)
			if tc.wantErr {
				assert.Error(t, err)
				if apiResp != nil {
					assert.Equal(t, tc.wantStatus, apiResp.StatusCode)
				}
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.wantStatus, apiResp.StatusCode)
			if assert.NotNil(t, out) {
				assert.Equal(t, "1", out.List.ID)
			}
		})
	}
}
