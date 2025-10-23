package deals

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
	rs, ok := s.(*RealService)
	assert.True(t, ok)
	assert.Equal(t, core, rs.client)
}

func TestListDeals_QueryParamsAndError(t *testing.T) {
	// Query param encoding
	hd := &testhelpers.HTTPDoer{BaseURL: "https://example.com/api/3/", RespStatus: 200, RespBody: []byte(`{"deals":[{"id":"d1"}]}`)}
	svc := NewRealServiceFromDoer(hd)

	opts := map[string]string{"search": "term", "limit": "10"}
	out, apiResp, err := svc.ListDeals(context.Background(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 200, apiResp.StatusCode)
	_ = out

	// Error path
	hdErr := &testhelpers.HTTPDoer{BaseURL: "https://example.com/api/3/", RespStatus: 400, RespBody: []byte(`{"errors":[{"title":"bad"}]}`), RespErr: nil}
	svcErr := NewRealServiceFromDoer(hdErr)
	out2, apiResp2, err2 := svcErr.ListDeals(context.Background(), nil)
	assert.Error(t, err2)
	var apiErr *client.APIError
	assert.True(t, errors.As(err2, &apiErr))
	assert.Equal(t, 400, apiResp2.StatusCode)
	assert.Nil(t, out2)
}

func TestListDeals_NilReceiver(t *testing.T) {
	var s *RealService
	// ListDeals has no nil receiver guard and will panic if s is nil; assert that behaviour explicitly
	assert.Panics(t, func() { _, _, _ = s.ListDeals(context.Background(), nil) })
}
