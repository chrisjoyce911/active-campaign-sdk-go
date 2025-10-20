package contactautomation

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestService_RemoveContactFromAutomation(t *testing.T) {
	tests := []struct {
		name     string
		mockResp *client.APIResponse
		mockBody []byte
		id       string
	}{
		{name: "ok", mockResp: &client.APIResponse{StatusCode: 200}, mockBody: []byte(`{"success":true}`), id: "ca1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := &testhelpers.MockDoer{Resp: tt.mockResp, Body: tt.mockBody}
			svc := NewRealServiceFromDoer(md)
			apiResp, err := svc.RemoveContactFromAutomation(context.Background(), tt.id)
			assert.NoError(t, err)
			assert.Equal(t, tt.mockResp.StatusCode, apiResp.StatusCode)
		})
	}
}
