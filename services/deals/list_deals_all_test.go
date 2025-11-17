package deals

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
)

// roundTripperFunc lets us customize responses based on the requested URL.
type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func resp(status int, body string) *http.Response {
	return &http.Response{StatusCode: status, Body: io.NopCloser(io.Reader(&stringReader{s: body})), Header: make(http.Header)}
}

// stringReader is a minimal io.Reader over a string (to avoid bytes import).
type stringReader struct{ s string }

func (r *stringReader) Read(p []byte) (int, error) {
	if len(r.s) == 0 {
		return 0, io.EOF
	}
	n := copy(p, r.s)
	r.s = r.s[n:]
	return n, nil
}

func TestListDealsAll_PaginatesByLimitAndOffset(t *testing.T) {
	// Transport that returns two pages when limit=2: offsets 0 and 2
	rt := roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		v, _ := url.ParseQuery(r.URL.RawQuery)
		off := v.Get("offset")
		switch off {
		case "0":
			return resp(200, `{"deals":[{"id":"d1"},{"id":"d2"}],"meta":{"total":4}}`), nil
		case "2":
			return resp(200, `{"deals":[{"id":"d3"},{"id":"d4"}],"meta":{"total":4}}`), nil
		default:
			return resp(200, `{"deals":[],"meta":{"total":4}}`), nil
		}
	})

	hd := &testhelpers.HTTPDoer{BaseURL: "https://example.com/api/3/", RespStatus: 200, Transport: rt}
	svc := NewRealServiceFromDoer(hd)

	// Force small page size to exercise pagination.
	list, apiResp, err := ListDealsAll(context.Background(), svc, map[string]string{"limit": "2"})
	if err != nil {
		t.Fatalf("ListDealsAll error: %v", err)
	}
	if apiResp == nil || apiResp.StatusCode != 200 {
		t.Fatalf("unexpected apiResp: %+v", apiResp)
	}
	if len(list) != 4 {
		t.Fatalf("expected 4 deals, got %d", len(list))
	}
}

func TestListDealsAll_StopsOnShortPageWhenNoTotal(t *testing.T) {
	rt := roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		v, _ := url.ParseQuery(r.URL.RawQuery)
		off := v.Get("offset")
		if off == "0" {
			return resp(200, `{"deals":[{"id":"a"},{"id":"b"}]}`), nil
		}
		return resp(200, `{"deals":[{"id":"c"}]}`), nil
	})

	hd := &testhelpers.HTTPDoer{BaseURL: "https://example.com/api/3/", RespStatus: 200, Transport: rt}
	svc := NewRealServiceFromDoer(hd)
	list, _, err := ListDealsAll(context.Background(), svc, map[string]string{"limit": "2"})
	if err != nil {
		t.Fatalf("ListDealsAll error: %v", err)
	}
	if len(list) != 3 {
		t.Fatalf("expected 3 deals, got %d", len(list))
	}
}

func TestListDealsAll_PropagatesError(t *testing.T) {
	rt := roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		return resp(400, `{"errors":[{"title":"bad"}]}`), nil
	})
	hd := &testhelpers.HTTPDoer{BaseURL: "https://example.com/api/3/", Transport: rt}
	svc := NewRealServiceFromDoer(hd)
	list, apiResp, err := ListDealsAll(context.Background(), svc, map[string]string{"limit": "2"})
	if err == nil {
		t.Fatalf("expected error")
	}
	if list != nil {
		t.Fatalf("expected nil list on error")
	}
	if apiResp == nil || apiResp.StatusCode != 400 {
		t.Fatalf("unexpected apiResp: %+v", apiResp)
	}
}

func TestListDealsAll_DefaultLimitAndInvalidLimit(t *testing.T) {
	// No limit provided -> default 100, invalid limit -> default 100
	sawLimit100 := 0
	rt := roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		v, _ := url.ParseQuery(r.URL.RawQuery)
		if v.Get("limit") == "100" {
			sawLimit100++
		}
		// Return a single result and meta.total=1 so helper stops
		return resp(200, `{"deals":[{"id":"z"}],"meta":{"total":1}}`), nil
	})
	hd := &testhelpers.HTTPDoer{BaseURL: "https://example.com/api/3/", Transport: rt}
	svc := NewRealServiceFromDoer(hd)

	// case 1: no limit in opts
	_, _, err := ListDealsAll(context.Background(), svc, map[string]string{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// case 2: invalid limit
	_, _, err = ListDealsAll(context.Background(), svc, map[string]string{"limit": "bad"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if sawLimit100 < 2 {
		t.Fatalf("expected to see limit=100 at least twice, saw %d", sawLimit100)
	}
}

func TestListDealsAll_RespectsStartingOffset(t *testing.T) {
	firstOffset := ""
	rt := roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		v, _ := url.ParseQuery(r.URL.RawQuery)
		if firstOffset == "" {
			firstOffset = v.Get("offset")
		}
		// Stop in one page using meta.total
		return resp(200, `{"deals":[{"id":"x"}],"meta":{"total":1}}`), nil
	})
	hd := &testhelpers.HTTPDoer{BaseURL: "https://example.com/api/3/", Transport: rt}
	svc := NewRealServiceFromDoer(hd)
	_, _, err := ListDealsAll(context.Background(), svc, map[string]string{"offset": "10"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if firstOffset != "10" {
		t.Fatalf("expected first offset=10, got %s", firstOffset)
	}
}
