package deals

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestRealService_UpdateDealNote(t *testing.T) {
	tests := []struct {
		name       string
		mockResp   *client.APIResponse
		mockBody   []byte
		dealID     string
		noteID     string
		wantStatus int
	}{
		{name: "ok", mockResp: &client.APIResponse{StatusCode: 200}, mockBody: []byte(`{"note":{"id":"n1"}}`), dealID: "d1", noteID: "n1", wantStatus: 200},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := &testhelpers.MockDoer{Resp: tt.mockResp, Body: tt.mockBody}
			svc := NewRealServiceFromDoer(md)

			out, apiResp, err := svc.UpdateDealNote(context.Background(), tt.dealID, tt.noteID, map[string]interface{}{"body": "x"})
			assert.NoError(t, err)
			assert.Equal(t, tt.wantStatus, apiResp.StatusCode)
			_ = out
		})
	}
}
