package accounts

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccountNote_ErrorAndNilReceiver(t *testing.T) {
	// error path from Doer
	md := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 500}, Err: &client.APIError{StatusCode: 500, Message: "err"}}
	svc := NewRealServiceFromDoer(md)

	out, apiResp, err := svc.CreateAccountNote(context.Background(), "a1", &AccountNoteRequest{Note: map[string]interface{}{"note": "x"}})
	assert.Error(t, err)
	assert.Nil(t, out)
	if apiResp != nil {
		assert.Equal(t, 500, apiResp.StatusCode)
	}

	// nil receiver should return a not-implemented style error (per implementation)
	var s *service
	out2, apiResp2, err2 := s.CreateAccountNote(context.Background(), "a1", &AccountNoteRequest{Note: map[string]interface{}{"note": "x"}})
	assert.Error(t, err2)
	assert.Nil(t, out2)
	assert.Nil(t, apiResp2)
}

func TestGetAccount_ErrorAndNilReceiver(t *testing.T) {
	// error from Doer
	md := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 404}, Err: &client.APIError{StatusCode: 404, Message: "not found"}}
	svc := NewRealServiceFromDoer(md)

	out, apiResp, err := svc.GetAccount(context.Background(), "missing")
	assert.Error(t, err)
	assert.Nil(t, out)
	if apiResp != nil {
		assert.Equal(t, 404, apiResp.StatusCode)
	}

	// nil receiver behaviour
	var s *service
	out2, apiResp2, err2 := s.GetAccount(context.Background(), "id")
	assert.Error(t, err2)
	assert.Nil(t, out2)
	assert.Nil(t, apiResp2)
}
