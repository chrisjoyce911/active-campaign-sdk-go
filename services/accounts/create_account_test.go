package accounts

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestService_CreateAccount(t *testing.T) {
	tests := []struct {
		name       string
		mockResp   *client.APIResponse
		mockBody   []byte
		req        interface{}
		wantStatus int
	}{
		{name: "ok", mockResp: &client.APIResponse{StatusCode: 201}, mockBody: []byte(`{"account":{"id":"a1"}}`), req: map[string]interface{}{"name": "X"}, wantStatus: 201},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := &testhelpers.MockDoer{Resp: tt.mockResp, Body: tt.mockBody}
			svc := NewRealServiceFromDoer(md)

			out, apiResp, err := svc.CreateAccount(context.Background(), tt.req)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantStatus, apiResp.StatusCode)
			_ = out
		})
	}
}
