package custom_objects

import (
	"context"
	"errors"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestGetObjectRecord(t *testing.T) {
	tests := []struct {
		name    string
		doer    *testhelpers.MockDoer
		wantErr bool
	}{
		{"ok", &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"record":{"id":"r1","fields":{"name":"x"}}}`)}, false},
		{"ok-empty", &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"record":{}}`)}, false},
		{"err-404", &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 404}, Body: []byte(`{"error":"not found"}`), Err: errors.New("not found")}, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc := NewRealServiceFromDoer(tc.doer)
			out, apiResp, err := svc.GetObjectRecord(context.Background(), "ot1", "r1")
			if tc.wantErr {
				assert.Error(t, err)
				if apiResp != nil {
					assert.Equal(t, 404, apiResp.StatusCode)
				}
				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, apiResp)
			if tc.name == "ok" {
				assert.Equal(t, "r1", out.ID)
				assert.Equal(t, "x", out.Fields["name"])
			} else if tc.name == "ok-empty" {
				assert.Equal(t, "", out.ID)
				assert.Nil(t, out.Fields)
			}
		})
	}
}
