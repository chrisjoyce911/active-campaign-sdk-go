package lists

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListLists(t *testing.T) {
	tests := []struct {
		name       string
		mockResp   *client.APIResponse
		mockBody   []byte
		opts       map[string]string
		wantStatus int
		wantErr    bool
	}{
		{name: "ok", mockResp: &client.APIResponse{StatusCode: 200}, mockBody: []byte(`{"lists":[{"id":"1","name":"L1"}],"meta":{"total":"1"}}`), opts: nil, wantStatus: 200, wantErr: false},
		{name: "err-422", mockResp: &client.APIResponse{StatusCode: 422}, mockBody: []byte(`{"error":"invalid"}`), opts: nil, wantStatus: 422, wantErr: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			md := &testhelpers.MockDoer{Resp: tc.mockResp, Body: tc.mockBody}
			if tc.wantErr {
				md.Err = assert.AnError
			}
			require := require.New(t)
			require.NotNil(md)

			svc := NewRealServiceFromDoer(md)
			require.NotNil(svc)

			out, apiResp, err := svc.ListLists(context.Background(), tc.opts)
			if tc.wantErr {
				assert.Error(t, err)
				if apiResp != nil {
					assert.Equal(t, tc.wantStatus, apiResp.StatusCode)
				}
				return
			}
			assert.NoError(t, err)
			require.NotNil(apiResp)
			assert.Equal(t, tc.wantStatus, apiResp.StatusCode)
			if assert.NotNil(t, out) {
				if tc.name == "ok" {
					if assert.NotEmpty(t, out.Lists) {
						assert.Equal(t, "1", out.Lists[0].ID)
					}
				}
			}
		})
	}
}
