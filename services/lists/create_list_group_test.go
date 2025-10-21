package lists

import (
	"context"
	"errors"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestCreateListGroup(t *testing.T) {
	tests := []struct {
		name       string
		mockResp   *client.APIResponse
		mockBody   []byte
		req        CreateListGroupRequest
		wantStatus int
		wantErr    bool
	}{
		{name: "ok", mockResp: &client.APIResponse{StatusCode: 201}, mockBody: []byte(`{"listGroup":{"listid":19,"groupid":1}}`), req: CreateListGroupRequest{ListGroup: ListGroup{ListID: 19, GroupID: 1}}, wantStatus: 201, wantErr: false},
		{name: "err-422", mockResp: &client.APIResponse{StatusCode: 422}, mockBody: []byte(`{"error":"invalid"}`), req: CreateListGroupRequest{ListGroup: ListGroup{}}, wantStatus: 422, wantErr: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			md := &testhelpers.MockDoer{Resp: tc.mockResp, Body: tc.mockBody}
			if tc.wantErr {
				md.Err = errors.New("invalid")
			}
			svc := NewRealServiceFromDoer(md)

			out, apiResp, err := svc.CreateListGroup(context.Background(), tc.req)
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
				if tc.name == "ok" {
					assert.Equal(t, 19, out.ListGroup.ListID)
				}
			}
		})
	}
}
