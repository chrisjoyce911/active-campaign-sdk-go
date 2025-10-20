package contacts

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/stretchr/testify/assert"
)

func TestRealService_SearchByEmail(t *testing.T) {
	tests := []struct {
		name   string
		email  string
		body   []byte
		status int
	}{
		{name: "found", email: "a@b.com", body: []byte(`{"contacts":[{"email":"a@b.com"}]}`), status: 200},
		{name: "none", email: "notfound@example.com", body: []byte(`{"contacts":[]}`), status: 200},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := &mockDoer{Resp: &client.APIResponse{StatusCode: tt.status}, Body: tt.body}
			svc := NewRealServiceFromDoer(md)

			out, apiResp, err := svc.SearchByEmail(context.Background(), tt.email)
			assert.NoError(t, err)
			assert.Equal(t, tt.status, apiResp.StatusCode)
			if tt.name == "found" {
				assert.Len(t, out.Contacts, 1)
				assert.Equal(t, "a@b.com", out.Contacts[0].Email)
			} else {
				assert.Len(t, out.Contacts, 0)
			}
		})
	}
}
