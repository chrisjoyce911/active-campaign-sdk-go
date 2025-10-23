package testhelpers

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/stretchr/testify/assert"
)

func TestHTTPDoer_GetWithParamsAndToken(t *testing.T) {
	h := &HTTPDoer{BaseURL: "https://api.example.com/", Token: "tok", RespStatus: 200, RespBody: []byte(`{"ok":true}`)}
	var out map[string]bool
	resp, err := h.Do(context.Background(), "GET", "path/resource", map[string]string{"filters[name]": "x"}, &out)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	// ensure LastRequest contains resolved URL with query
	assert.Contains(t, h.LastRequest.URL.String(), "filters%5Bname%5D=x")
	assert.Equal(t, "tok", h.LastRequest.Header.Get("Api-Token"))
}

func TestHTTPDoer_PostJsonAndErrorResponse(t *testing.T) {
	h := &HTTPDoer{BaseURL: "https://api.example.com/", RespStatus: 400, RespBody: []byte(`{"error":"bad"}`)}
	req := map[string]string{"a": "b"}
	_, err := h.Do(context.Background(), "POST", "p", req, nil)
	assert.Error(t, err)
	if apiErr, ok := err.(*client.APIError); ok {
		assert.Equal(t, 400, apiErr.StatusCode)
	}
}
