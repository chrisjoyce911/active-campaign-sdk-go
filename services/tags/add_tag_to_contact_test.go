package tags

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestService_AddTagToContact(t *testing.T) {
	tests := []struct {
		name       string
		mockResp   *client.APIResponse
		mockBody   []byte
		contactID  string
		req        *CreateOrUpdateTagRequest
		wantStatus int
		wantID     string
	}{
		{name: "ok", mockResp: &client.APIResponse{StatusCode: 201}, mockBody: []byte(`{"tag":{"id":"t1","tag":"x"}}`), contactID: "c1", req: &CreateOrUpdateTagRequest{Tag: TagPayload{Tag: "x"}}, wantStatus: 201, wantID: "t1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := &testhelpers.MockDoer{Resp: tt.mockResp, Body: tt.mockBody}
			svc := NewRealServiceFromDoer(md)

			out, apiResp, err := svc.AddTagToContact(context.Background(), tt.contactID, tt.req)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantStatus, apiResp.StatusCode)
			if out != nil {
				assert.Equal(t, tt.wantID, out.Tag.ID)
			}
		})
	}
}
