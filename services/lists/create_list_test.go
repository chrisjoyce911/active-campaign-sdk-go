package lists

import (
	"context"
	"errors"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateList(t *testing.T) {
	tests := []struct {
		name       string
		mockResp   *client.APIResponse
		mockBody   []byte
		req        CreateListRequest
		wantStatus int
		wantErr    bool
	}{
		{name: "ok", mockResp: &client.APIResponse{StatusCode: 201}, mockBody: []byte(`{"list":{"id":"10","name":"x"}}`), req: CreateListRequest{List: List{Name: "x", StringID: "x"}}, wantStatus: 201, wantErr: false},
		{name: "ok-empty", mockResp: &client.APIResponse{StatusCode: 201}, mockBody: []byte(`{"list":{}}`), req: CreateListRequest{List: List{Name: "y"}}, wantStatus: 201, wantErr: false},
		{name: "err-422", mockResp: &client.APIResponse{StatusCode: 422}, mockBody: []byte(`{"error":"invalid"}`), req: CreateListRequest{List: List{}}, wantStatus: 422, wantErr: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			md := &testhelpers.MockDoer{Resp: tc.mockResp, Body: tc.mockBody}
			if tc.wantErr {
				md.Err = errors.New("invalid")
			}
			require := require.New(t)
			require.NotNil(md)

			svc := NewRealServiceFromDoer(md)
			require.NotNil(svc)

			out, apiResp, err := svc.CreateList(context.Background(), tc.req)
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
					assert.Equal(t, "10", out.List.ID)
				}
			}
		})
	}
}
