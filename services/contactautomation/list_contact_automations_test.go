package contactautomation

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_ListContactAutomations(t *testing.T) {
	tests := []struct {
		name      string
		mockResp  *client.APIResponse
		mockBody  []byte
		contactID string
	}{
		{name: "ok", mockResp: &client.APIResponse{StatusCode: 200}, mockBody: []byte(`{"contactAutomations":[{"id":"ca1","automation":"a1"}]}`), contactID: "c1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := &testhelpers.MockDoer{Resp: tt.mockResp, Body: tt.mockBody}
			require := require.New(t)
			require.NotNil(md)

			svc := NewRealServiceFromDoer(md)
			require.NotNil(svc)

			out, apiResp, err := svc.ListContactAutomations(context.Background(), tt.contactID)
			assert.NoError(t, err)
			require.NotNil(apiResp)
			assert.Equal(t, tt.mockResp.StatusCode, apiResp.StatusCode)
			if out != nil && len(out.ContactAutomationsOrEmpty()) > 0 {
				assert.Equal(t, "ca1", out.ContactAutomationsOrEmpty()[0].ID)
			}
		})
	}
}
