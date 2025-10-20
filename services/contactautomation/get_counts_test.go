package contactautomation

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestService_GetCounts(t *testing.T) {
	tests := []struct {
		name     string
		mockResp *client.APIResponse
		mockBody []byte
	}{
		{name: "ok", mockResp: &client.APIResponse{StatusCode: 200}, mockBody: []byte(`{"counts":[{"automation":"a1","count":3}]}`)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := &testhelpers.MockDoer{Resp: tt.mockResp, Body: tt.mockBody}
			svc := NewRealServiceFromDoer(md)
			out, apiResp, err := svc.GetCounts(context.Background())
			assert.NoError(t, err)
			assert.Equal(t, tt.mockResp.StatusCode, apiResp.StatusCode)
			if out != nil && len(out.CountsOrEmpty()) > 0 {
				assert.Equal(t, "a1", out.CountsOrEmpty()[0].Automation)
			}
		})
	}
}
