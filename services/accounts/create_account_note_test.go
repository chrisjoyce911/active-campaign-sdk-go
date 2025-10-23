package accounts

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestService_CreateAccountNote(t *testing.T) {
	tests := []struct {
		name       string
		mockResp   *client.APIResponse
		mockBody   []byte
		accountID  string
		req        *AccountNoteRequest
		wantStatus int
	}{
		{name: "ok", mockResp: &client.APIResponse{StatusCode: 201}, mockBody: []byte(`{"note":{"id":"n1"}}`), accountID: "a1", req: &AccountNoteRequest{Note: map[string]interface{}{"note": "x"}}, wantStatus: 201},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := &testhelpers.MockDoer{Resp: tt.mockResp, Body: tt.mockBody}
			svc := NewRealServiceFromDoer(md)

			out, apiResp, err := svc.CreateAccountNote(context.Background(), tt.accountID, tt.req)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantStatus, apiResp.StatusCode)
			_ = out
		})
	}
}
