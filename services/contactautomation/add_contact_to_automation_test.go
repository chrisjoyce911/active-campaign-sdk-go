package contactautomation

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestService_AddContactToAutomation(t *testing.T) {
	tests := []struct {
		name     string
		mockResp *client.APIResponse
		mockBody []byte
		req      *CreateContactAutomationRequest
	}{
		{name: "ok", mockResp: &client.APIResponse{StatusCode: 201}, mockBody: []byte(`{"contactAutomation":{"id":"ca1","automation":"a1"}}`), req: &CreateContactAutomationRequest{ContactAutomation: ContactAutomationPayload{Automation: "a1", Contact: "c1"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := &testhelpers.MockDoer{Resp: tt.mockResp, Body: tt.mockBody}
			svc := NewRealServiceFromDoer(md)
			out, apiResp, err := svc.AddContactToAutomation(context.Background(), tt.req)
			assert.NoError(t, err)
			assert.Equal(t, tt.mockResp.StatusCode, apiResp.StatusCode)
			if out != nil {
				assert.Equal(t, "a1", out.ContactAutomation.Automation)
			}
		})
	}
}
