package contacts

import (
	"context"
	"net/http"
	"net/url"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

func (s *RealService) ListContacts(ctx context.Context, opts map[string]string) (*ContactSearchResponse, *client.APIResponse, error) {
	out := &ContactSearchResponse{}
	base := "contacts"
	if len(opts) > 0 {
		vals := url.Values{}
		for k, v := range opts {
			vals.Add(k, v)
		}
		base = base + "?" + vals.Encode()
	}
	apiResp, err := s.client.Do(ctx, http.MethodGet, base, nil, out)
	return out, apiResp, err
}
