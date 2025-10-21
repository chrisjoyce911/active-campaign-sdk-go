package custom_objects

import (
	"context"
	"errors"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestGetObjectType(t *testing.T) {
	tests := []struct {
		name    string
		doer    *testhelpers.MockDoer
		id      string
		wantErr bool
	}{
		{"ok", &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"schema":{"id":"ot1","slug":"s1"}}`)}, "ot1", false},
		{"ok-empty", &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"schema":{}}`)}, "ot2", false},
		{"err-missing", nil, "", true},
		{"err-404", &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 404}, Body: []byte(`{"error":"not found"}`), Err: errors.New("not found")}, "ot-missing", true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var svc CustomObjectsService
			if tc.doer != nil {
				svc = NewRealServiceFromDoer(tc.doer)
			} else {
				svc = &service{}
			}
			out, apiResp, err := svc.GetObjectType(context.Background(), tc.id)
			if tc.wantErr {
				assert.Error(t, err)
				_ = apiResp
				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, apiResp)
			switch tc.name {
			case "ok":
				assert.Equal(t, "ot1", out.ID)
			case "ok-empty":
				assert.Equal(t, "", out.ID)
			}
		})
	}
}
