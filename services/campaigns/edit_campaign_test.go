package campaigns

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

func ptrString(s string) *string { return &s }

type stubDoerEdit struct {
	resp *client.APIResponse
	err  error
}

func (s stubDoerEdit) Do(ctx context.Context, method, path string, v interface{}, out interface{}) (*client.APIResponse, error) {
	if s.err != nil {
		return nil, s.err
	}
	if out != nil {
		// populate out if it's the wrapper type
		if w, ok := out.(*struct {
			Campaign Campaign `json:"campaign"`
		}); ok {
			*w = struct {
				Campaign Campaign `json:"campaign"`
			}{Campaign: Campaign{ID: "29", Name: "MY SUPER CAMPAIGN"}}
		}
	}
	return s.resp, nil
}

func TestEditCampaign_Table(t *testing.T) {
	cases := []struct {
		name    string
		doer    client.Doer
		wantErr bool
	}{
		{"success", stubDoerEdit{resp: &client.APIResponse{StatusCode: http.StatusOK}}, false},
		{"doer_error", stubDoerEdit{err: errors.New("boom")}, true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			svc := &service{client: tc.doer}
			req := &EditCampaignRequest{Name: ptrString("x")}
			got, apiResp, err := svc.EditCampaign(context.Background(), "29", req)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if apiResp == nil {
				t.Fatalf("expected apiResp")
			}
			if got == nil || got.ID != "29" {
				t.Fatalf("expected campaign id 29 got %+v", got)
			}
		})
	}
}
