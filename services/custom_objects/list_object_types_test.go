package custom_objects

import (
	"context"
	"errors"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestListObjectTypes(t *testing.T) {
	tests := []struct {
		name       string
		doer       *testhelpers.MockDoer
		wantStatus int
		wantErr    bool
	}{
		{"ok-populated", &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"objectTypes": [{"id":"ot1","name":"OT1"}]}`)}, 200, false},
		{"ok-empty", &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"objectTypes": []}`)}, 200, false},
		{"ok-empty-legacy", &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"schemas": []}`)}, 200, false},
		{"err-500", &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 500}, Body: []byte(`{"error":"boom"}`), Err: errors.New("boom")}, 500, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc := NewRealServiceFromDoer(tc.doer)
			resp, apiResp, err := svc.ListObjectTypes(context.Background(), nil)
			if tc.wantErr {
				assert.Error(t, err)
				if apiResp != nil {
					assert.Equal(t, tc.wantStatus, apiResp.StatusCode)
				}
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.wantStatus, apiResp.StatusCode)
			if tc.name == "ok-populated" {
				if assert.NotNil(t, resp) && assert.NotEmpty(t, resp.Schemas) {
					assert.Equal(t, "ot1", resp.Schemas[0].ID)
				}
			} else {
				assert.NotNil(t, resp)
			}
		})
	}
}
