package contacts

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRealService_CreateContact_RequestShape(t *testing.T) {
	tests := []struct {
		name             string
		req              *CreateContactRequest
		respBody         []byte
		respStatus       int
		wantBodyContains string
	}{
		{name: "basic create", req: &CreateContactRequest{Contact: &Contact{Email: "x@y.com", FirstName: "X"}}, respBody: []byte(`{"contact":{"id":"10","email":"x@y.com"}}`), respStatus: 201, wantBodyContains: `"email":"x@y.com"`},
		{name: "with fieldValues", req: &CreateContactRequest{Contact: &Contact{Email: "f@v.com", FieldValues: &[]FieldValue{{Field: "1", Value: "ABC"}}}}, respBody: []byte(`{"contact":{"id":"11","email":"f@v.com","fieldValues":[{"field":"1","value":"ABC"}]}}`), respStatus: 201, wantBodyContains: `"fieldValues"`},
		{name: "validation error 422", req: &CreateContactRequest{Contact: &Contact{Email: "bad"}}, respBody: []byte(`{"errors":[{"title":"invalid","detail":"email invalid"}]}`), respStatus: 422, wantBodyContains: `"email":"bad"`},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			hd := &testhelpers.HTTPDoer{BaseURL: "https://example.api-us1.com/api/3/", Token: "tok", RespStatus: tc.respStatus, RespBody: tc.respBody}
			require := require.New(t)
			require.NotNil(hd)

			svc := NewRealServiceFromDoer(hd)
			require.NotNil(svc)

			out, apiResp, err := svc.Create(context.Background(), tc.req)
			if tc.respStatus >= 200 && tc.respStatus < 300 {
				assert.NoError(t, err)
				assert.NotNil(t, out)
			}
			if apiResp != nil {
				require.NotNil(apiResp)
				assert.Equal(t, tc.respStatus, apiResp.StatusCode)
			}

			// assertions on final HTTP request
			assert.Equal(t, "POST", hd.LastRequest.Method)
			assert.Contains(t, hd.LastRequest.URL.String(), "/api/3/contacts")
			// headers
			assert.Equal(t, "application/json", hd.LastRequest.Header.Get("Content-Type"))
			assert.Equal(t, "tok", hd.LastRequest.Header.Get("Api-Token"))
			if hd.LastRequestBody == nil {
				t.Fatalf("no recorded request body")
			}
			// expected request body is the marshalled CreateContactRequest
			expB, err := json.Marshal(tc.req)
			if err != nil {
				t.Fatalf("marshal expected: %v", err)
			}
			assert.JSONEq(t, string(expB), string(hd.LastRequestBody))
		})
	}
}
