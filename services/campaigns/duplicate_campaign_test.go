package campaigns

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	th "github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
)

type stubDoerDup struct {
	resp *client.APIResponse
	body []byte
	err  error
}

func (s stubDoerDup) Do(ctx context.Context, method, path string, v interface{}, out interface{}) (*client.APIResponse, error) {
	if s.err != nil {
		return nil, s.err
	}
	if out != nil && s.body != nil {
		_ = json.Unmarshal(s.body, out)
	}
	return s.resp, nil
}

func TestDuplicateCampaign_Table(t *testing.T) {
	cases := []struct {
		name    string
		doer    client.Doer
		wantErr bool
	}{
		{"success", stubDoerDup{resp: &client.APIResponse{StatusCode: http.StatusOK}, body: []byte(`{"succeeded":1,"message":"ok","newCampaignId":217}`)}, false},
		{"doer_error", stubDoerDup{err: errors.New("boom")}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			svc := &service{client: tc.doer}
			got, apiResp, err := svc.DuplicateCampaign(context.Background(), "2")
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if apiResp == nil || apiResp.StatusCode != http.StatusOK {
				t.Fatalf("expected 200 got %+v", apiResp)
			}
			if got == nil || got.NewCampaignID == 0 {
				t.Fatalf("expected newCampaignId populated")
			}
		})
	}
}

func TestDuplicateCampaign_Recording_PathAndMethod(t *testing.T) {
	rd := &th.RecordingDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"succeeded":1,"message":"Campaign draft copied.","newCampaignId":217}`)}
	svc := &service{client: rd}

	_, _, err := svc.DuplicateCampaign(context.Background(), "2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if rd.LastMethod != http.MethodPost {
		t.Fatalf("expected POST got %s", rd.LastMethod)
	}
	if rd.LastPath != "campaigns/2/copy" {
		t.Fatalf("expected campaigns/2/copy got %s", rd.LastPath)
	}
}
