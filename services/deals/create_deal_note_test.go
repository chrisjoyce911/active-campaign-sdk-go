package deals

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestRealService_CreateDealNote(t *testing.T) {
	tests := []struct {
		name       string
		mockResp   *client.APIResponse
		mockBody   []byte
		dealID     string
		req        interface{}
		wantStatus int
	}{
		{name: "success", mockResp: &client.APIResponse{StatusCode: 201}, mockBody: []byte(`{"note":{"id":"n1"}}`), dealID: "d1", req: map[string]interface{}{"note": "hello"}, wantStatus: 201},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := &testhelpers.MockDoer{Resp: tt.mockResp, Body: tt.mockBody}
			svc := NewRealServiceFromDoer(md)

			out, apiResp, err := svc.CreateDealNote(context.Background(), tt.dealID, tt.req)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantStatus, apiResp.StatusCode)
			_ = out
		})
	}
}
