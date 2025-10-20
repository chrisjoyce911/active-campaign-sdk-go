package tags

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestService_UpdateTag(t *testing.T) {
	tests := []struct {
		name       string
		mockResp   *client.APIResponse
		mockBody   []byte
		id         string
		req        *CreateOrUpdateTagRequest
		wantStatus int
		wantTag    string
	}{
		{name: "ok", mockResp: &client.APIResponse{StatusCode: 200}, mockBody: []byte(`{"tag":{"id":"t1","tag":"updated"}}`), id: "t1", req: &CreateOrUpdateTagRequest{Tag: TagPayload{Tag: "updated"}}, wantStatus: 200, wantTag: "updated"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := &testhelpers.MockDoer{Resp: tt.mockResp, Body: tt.mockBody}
			svc := NewRealServiceFromDoer(md)

			out, apiResp, err := svc.UpdateTag(context.Background(), tt.id, tt.req)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantStatus, apiResp.StatusCode)
			if out != nil {
				assert.Equal(t, tt.wantTag, out.Tag.Tag)
			}
		})
	}
}
