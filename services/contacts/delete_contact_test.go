package contacts

import (
	"context"
	"net/http"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/stretchr/testify/assert"
)

func TestRealService_DeleteContact(t *testing.T) {
	tests := []struct {
		name string
		id   string
		resp *client.APIResponse
		err  error
	}{
		{name: "success", id: "123", resp: &client.APIResponse{StatusCode: 200, HTTP: &http.Response{StatusCode: 200}}, err: nil},
		{name: "not found", id: "404", resp: &client.APIResponse{StatusCode: 404, HTTP: &http.Response{StatusCode: 404}}, err: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := &mockDoer{Resp: tt.resp, Err: tt.err}
			svc := NewRealServiceFromDoer(md)

			apiResp, err := svc.DeleteContact(context.Background(), tt.id)
			if tt.err != nil {
				assert.Equal(t, tt.err, err)
				return
			}
			assert.NoError(t, err)
			if apiResp != nil {
				assert.Equal(t, tt.resp.StatusCode, apiResp.StatusCode)
			}
		})
	}
}
