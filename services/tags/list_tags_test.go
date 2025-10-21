package tags

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestService_ListTags(t *testing.T) {
	tests := []struct {
		name       string
		mockResp   *client.APIResponse
		mockBody   []byte
		opts       map[string]string
		wantStatus int
	}{
		{name: "ok", mockResp: &client.APIResponse{StatusCode: 200}, mockBody: []byte(`{"tags":[{"id":"t1"}]}`), opts: map[string]string{}, wantStatus: 200},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := &testhelpers.MockDoer{Resp: tt.mockResp, Body: tt.mockBody}
			svc := NewRealServiceFromDoer(md)

			out, apiResp, err := svc.ListTags(context.Background(), tt.opts)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantStatus, apiResp.StatusCode)
			if out != nil && len(out.TagsOrEmpty()) > 0 {
				assert.Equal(t, "t1", out.TagsOrEmpty()[0].ID)
			}
		})
	}
}
