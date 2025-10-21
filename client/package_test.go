package client

import (
	"net/url"
	"testing"
)

func TestAPIError_Error(t *testing.T) {
	e := &APIError{StatusCode: 400, Message: "bad request", Body: []byte("err")}
	if e.Error() != "bad request" {
		t.Fatalf("expected Error() to return message, got %q", e.Error())
	}
}

func TestBuildContactsSearchPath(t *testing.T) {
	cases := []struct {
		email string
		want  string
	}{
		{"user@example.com", "/api/3/contacts?email=" + url.QueryEscape("user@example.com")},
		{"a+b@x.y", "/api/3/contacts?email=" + url.QueryEscape("a+b@x.y")},
		{"", "/api/3/contacts?email="},
	}

	for _, c := range cases {
		got := BuildContactsSearchPath(c.email)
		if got != c.want {
			t.Fatalf("BuildContactsSearchPath(%q) = %q; want %q", c.email, got, c.want)
		}
	}
}
