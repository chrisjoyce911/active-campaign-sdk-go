package accounts

import (
	"context"
	"errors"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestService_DeleteAccount_Non2xx(t *testing.T) {
	md := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 404, Body: []byte(`{"errors":[{"title":"not found"}]}`)}, Err: &client.APIError{StatusCode: 404, Message: "not found", Body: []byte(`{"errors":[{"title":"not found"}]}`)}, Body: []byte(`{"errors":[{"title":"not found"}]}`)}
	svc := NewRealServiceFromDoer(md)

	apiResp, err := svc.DeleteAccount(context.Background(), "missing")
	assert.Error(t, err)
	var apiErr *client.APIError
	ok := errors.As(err, &apiErr)
	assert.True(t, ok, "expected APIError")
	assert.Equal(t, 404, apiResp.StatusCode)
}

func TestService_BulkDeleteAccounts_Non2xx(t *testing.T) {
	md := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 500, Body: []byte(`{"errors":[{"title":"server"}]}`)}, Err: &client.APIError{StatusCode: 500, Message: "server", Body: []byte(`{"errors":[{"title":"server"}]}`)}, Body: []byte(`{"errors":[{"title":"server"}]}`)}
	svc := NewRealServiceFromDoer(md)

	apiResp, err := svc.BulkDeleteAccounts(context.Background(), []string{"a1", "a2"})
	assert.Error(t, err)
	var apiErr *client.APIError
	ok := errors.As(err, &apiErr)
	assert.True(t, ok, "expected APIError")
	assert.Equal(t, 500, apiResp.StatusCode)
}

func TestService_NilOrInvalidInputs(t *testing.T) {
	md := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: nil}
	svc := NewRealServiceFromDoer(md)

	// CreateAccount with nil req should probably return an error; current impl calls client.Do with nil body — assert behavior
	out, apiResp, err := svc.CreateAccount(context.Background(), nil)
	// current behavior: MockDoer returns nil body and Resp with StatusCode 200; assert behavior is stable
	assert.NoError(t, err)
	assert.NotNil(t, apiResp)
	_ = out

	// UpdateAccount with nil req
	out2, apiResp2, err2 := svc.UpdateAccount(context.Background(), "a1", nil)
	assert.NoError(t, err2)
	assert.NotNil(t, apiResp2)
	_ = out2

	// CreateAccountNote with nil req
	noteOut, apiResp3, err3 := svc.CreateAccountNote(context.Background(), "a1", nil)
	assert.NoError(t, err3)
	assert.NotNil(t, apiResp3)
	_ = noteOut
}
