package contacts

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestRealService_UpdateContact(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		body := []byte(`{"contact": {"id":"c1","email":"j@doe.com","firstName":"J"}}`)
		md := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: body}
		svc := NewRealServiceFromDoer(md)

		req := &CreateContactRequest{Contact: &Contact{Email: "j@doe.com", FirstName: "J"}}
		out, apiResp, err := svc.UpdateContact(context.Background(), "c1", req)
		assert.NoError(t, err)
		if assert.NotNil(t, apiResp) {
			assert.Equal(t, 200, apiResp.StatusCode)
		}
		if assert.NotNil(t, out) {
			assert.Equal(t, "c1", out.Contact.ID)
			assert.Equal(t, "j@doe.com", out.Contact.Email)
		}
	})

	t.Run("not found", func(t *testing.T) {
		// simulate 404 body
		md := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 404}, Body: nil}
		svc := NewRealServiceFromDoer(md)

		req := &CreateContactRequest{Contact: &Contact{Email: "missing@example.com"}}
		out, apiResp, err := svc.UpdateContact(context.Background(), "9999", req)
		// The client returns apiResp even for non-2xx; our MockDoer doesn't return an error
		assert.NoError(t, err)
		if assert.NotNil(t, apiResp) {
			assert.Equal(t, 404, apiResp.StatusCode)
		}
		// out may be empty or nil depending on response; ensure no panic
		_ = out
	})
}
