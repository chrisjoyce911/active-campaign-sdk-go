package deals

import (
	"context"
	"net/url"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
)

func TestListDealStages_SuccessAndQuery(t *testing.T) {
	body := []byte(`{"dealStages":[{"id":"15","title":"Initial Contact","group":"4"}],"meta":{"total":1}}`)
	hd := &testhelpers.HTTPDoer{BaseURL: "https://example.com/api/3/", RespStatus: 200, RespBody: body}
	svc := NewRealServiceFromDoer(hd)

	opts := map[string]string{"filters[d_groupid]": "2", "filters[title]": "Init", "orders[title]": "ASC"}
	out, apiResp, err := svc.ListDealStages(context.Background(), opts)
	if err != nil {
		t.Fatalf("ListDealStages error: %v", err)
	}
	if apiResp.StatusCode != 200 {
		t.Fatalf("unexpected status: %d", apiResp.StatusCode)
	}
	if out == nil || len(out.DealStages) != 1 || out.DealStages[0].ID != "15" {
		t.Fatalf("unexpected parsed response: %+v", out)
	}

	// Assert request path and query
	if hd.LastRequest == nil {
		t.Fatalf("no request recorded")
	}
	if got := hd.LastRequest.URL.Path; got != "/api/3/dealStages" {
		t.Fatalf("unexpected path: %s", got)
	}
	q, _ := url.QueryUnescape(hd.LastRequest.URL.RawQuery)
	if !(contains(q, "filters[d_groupid]=2") && contains(q, "filters[title]=Init") && contains(q, "orders[title]=ASC")) {
		t.Fatalf("unexpected query: %s", q)
	}
}

func TestListDealStages_Error(t *testing.T) {
	body := []byte(`{"errors":[{"title":"bad"}]}`)
	hd := &testhelpers.HTTPDoer{BaseURL: "https://example.com/api/3/", RespStatus: 400, RespBody: body}
	svc := NewRealServiceFromDoer(hd)
	out, apiResp, err := svc.ListDealStages(context.Background(), nil)
	if err == nil {
		t.Fatalf("expected error")
	}
	if apiResp == nil || apiResp.StatusCode != 400 {
		t.Fatalf("unexpected apiResp: %+v", apiResp)
	}
	if out != nil {
		t.Fatalf("expected nil out on error")
	}
}

// contains is a tiny helper for substring checks without importing strings at top-level
func contains(s, sub string) bool {
	return len(s) >= len(sub) && (s == sub || (len(s) > len(sub) && (indexOf(s, sub) >= 0)))
}
func indexOf(s, sub string) int {
	// very small naive search sufficient for tests
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return i
		}
	}
	return -1
}
