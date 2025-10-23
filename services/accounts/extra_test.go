package accounts

import (
	"context"
	"errors"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestNewRealService_AssignsClient(t *testing.T) {
	core := &client.CoreClient{}
	s := NewRealService(core)
	assert.NotNil(t, s)

	srv, ok := s.(*service)
	assert.True(t, ok, "expected returned service to be concrete *service")
	// underlying client should be the same pointer we passed in
	assert.Equal(t, core, srv.client)
}

func TestService_ListAccounts_QueryParamsAndErrorPaths(t *testing.T) {
	// Query params encoding
	hd := &testhelpers.HTTPDoer{BaseURL: "https://example.com/api/3/", RespStatus: 200, RespBody: []byte(`{"accounts":[{"id":"a1"}]}`)}
	svc := NewRealServiceFromDoer(hd)

	opts := map[string]string{"search": "term", "limit": "10"}
	out, apiResp, err := svc.ListAccounts(context.Background(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 200, apiResp.StatusCode)
	// verify the outgoing request had the expected query params
	q := hd.LastRequest.URL.Query()
	assert.Equal(t, "term", q.Get("search"))
	assert.Equal(t, "10", q.Get("limit"))
	_ = out

	// Error response path
	hdErr := &testhelpers.HTTPDoer{BaseURL: "https://example.com/api/3/", RespStatus: 400, RespBody: []byte(`{"errors":[{"title":"bad"}]}`)}
	svcErr := NewRealServiceFromDoer(hdErr)
	out2, apiResp2, err2 := svcErr.ListAccounts(context.Background(), nil)
	assert.Error(t, err2)
	var apiErr *client.APIError
	assert.True(t, errors.As(err2, &apiErr))
	assert.Equal(t, 400, apiResp2.StatusCode)
	assert.Nil(t, out2)
}

func TestService_ListAccounts_NilReceiver(t *testing.T) {
	var s *service
	out, apiResp, err := s.ListAccounts(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, apiResp)
	assert.Nil(t, out)
}
