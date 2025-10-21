package contacts

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestRealService_UpdateFieldValueForContact_RequestShape(t *testing.T) {
	tests := []struct {
		name             string
		req              *FieldValuePayload
		respBody         []byte
		respStatus       int
		wantBodyContains string
	}{
		{name: "basic", req: &FieldValuePayload{Contact: "1", Field: "2", Value: "Blue"}, respBody: []byte(`{"fieldValue":{"id":"fv1","value":"Blue"}}`), respStatus: 201, wantBodyContains: `"value":"Blue"`},
		{name: "validation 422", req: &FieldValuePayload{Contact: "x", Field: "y", Value: ""}, respBody: []byte(`{"errors":[{"title":"invalid"}]}`), respStatus: 422, wantBodyContains: `"contact":"x"`},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			hd := &testhelpers.HTTPDoer{BaseURL: "https://example.api-us1.com/api/3/", Token: "tok", RespStatus: tc.respStatus, RespBody: tc.respBody}
			svc := NewRealServiceFromDoer(hd)

			out, apiResp, err := svc.UpdateFieldValueForContact(context.Background(), tc.req)
			if tc.respStatus >= 200 && tc.respStatus < 300 {
				assert.NoError(t, err)
				assert.NotNil(t, out)
			}
			if apiResp != nil {
				assert.Equal(t, tc.respStatus, apiResp.StatusCode)
			}

			assert.Equal(t, "POST", hd.LastRequest.Method)
			assert.Contains(t, hd.LastRequest.URL.String(), "/api/3/fieldValues")
			assert.Equal(t, "application/json", hd.LastRequest.Header.Get("Content-Type"))
			assert.Equal(t, "tok", hd.LastRequest.Header.Get("Api-Token"))
			if hd.LastRequestBody == nil {
				t.Fatalf("no recorded request body")
			}
			expB, _ := json.Marshal(tc.req)
			assert.JSONEq(t, string(expB), string(hd.LastRequestBody))
		})
	}
}
