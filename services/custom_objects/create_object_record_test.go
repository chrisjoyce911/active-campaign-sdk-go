package custom_objects

import (
	"context"
	"errors"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestCreateObjectRecord(t *testing.T) {
	tests := []struct {
		name       string
		doer       *testhelpers.MockDoer
		payload    *CreateRecordRequest
		wantStatus int
		wantErr    bool
	}{
		{"ok", &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 201}, Body: []byte(`{"record":{"id":"r1"}}`)}, &CreateRecordRequest{Fields: map[string]interface{}{"name": "x"}}, 201, false},
		{"ok-empty", &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 201}, Body: []byte(`{"record":{}}`)}, &CreateRecordRequest{Fields: map[string]interface{}{"name": "y"}}, 201, false},
		{"err-422", &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 422}, Body: []byte(`{"error":"invalid"}`), Err: errors.New("invalid")}, &CreateRecordRequest{Fields: map[string]interface{}{}}, 422, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc := NewRealServiceFromDoer(tc.doer)
			out, apiResp, err := svc.CreateObjectRecord(context.Background(), "ot1", tc.payload)
			if tc.wantErr {
				assert.Error(t, err)
				if apiResp != nil {
					assert.Equal(t, tc.wantStatus, apiResp.StatusCode)
				}
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.wantStatus, apiResp.StatusCode)
			if assert.NotNil(t, out) {
				switch tc.name {
				case "ok":
					assert.Equal(t, "r1", out.Record.ID)
				case "ok-empty":
					assert.Equal(t, "", out.Record.ID)
				}
			}
		})
	}
}
