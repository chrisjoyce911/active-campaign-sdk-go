package contacts

import (
    "context"
    "encoding/json"
    "testing"

    "github.com/chrisjoyce911/active-campaign-sdk-go/client"
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
            // The client now wraps the payload in a top-level `fieldValue` envelope
            envelope := struct {
                FieldValue *FieldValuePayload `json:"fieldValue"`
            }{FieldValue: tc.req}
            expB, _ := json.Marshal(envelope)
            assert.JSONEq(t, string(expB), string(hd.LastRequestBody))
        })
    }
}

func TestRealService_UpdateFieldValueForContact_EnvelopeShape(t *testing.T) {
    req := &FieldValuePayload{Contact: "42", Field: "13", Value: "Acme"}
    // HTTPDoer records the final request and body so we can assert the JSON shape.
    hd := &testhelpers.HTTPDoer{BaseURL: "https://example.api-us1.com/api/3/", Token: "tok", RespStatus: 201, RespBody: []byte(`{"fieldValue":{"id":"fv1","value":"Acme"}}`)}
    svc := NewRealServiceFromDoer(hd)

    out, apiResp, err := svc.UpdateFieldValueForContact(context.Background(), req)
    assert.NoError(t, err)
    assert.NotNil(t, apiResp)
    assert.NotNil(t, out)

    if hd.LastRequestBody == nil {
        t.Fatalf("expected request body to be recorded")
    }

    // The request body should be a JSON object with top-level fieldValue key.
    var got map[string]json.RawMessage
    if err := json.Unmarshal(hd.LastRequestBody, &got); err != nil {
        t.Fatalf("invalid request JSON: %v", err)
    }

    raw, ok := got["fieldValue"]
    if !ok {
        t.Fatalf("expected top-level 'fieldValue' key in request JSON; got: %s", string(hd.LastRequestBody))
    }

    var fv FieldValuePayload
    if err := json.Unmarshal(raw, &fv); err != nil {
        t.Fatalf("invalid fieldValue payload: %v", err)
    }
    assert.Equal(t, req.Contact, fv.Contact)
    assert.Equal(t, req.Field, fv.Field)
    assert.Equal(t, req.Value, fv.Value)
}

func TestRealService_UpdateFieldValueForContact_Non2xx(t *testing.T) {
    // Simulate a validation error from the API (422) with a JSON body
    body := []byte(`{"errors":[{"title":"invalid contact"}]}`)
    hd := &testhelpers.HTTPDoer{BaseURL: "https://example.api-us1.com/api/3/", Token: "tok", RespStatus: 422, RespBody: body}
    svc := NewRealServiceFromDoer(hd)

    req := &FieldValuePayload{Contact: "x", Field: "y", Value: ""}
    out, apiResp, err := svc.UpdateFieldValueForContact(context.Background(), req)

    // Different Doer implementations may or may not populate `out` on non-2xx.
    // We primarily assert that callers receive the low-level apiResp and an error.
    _ = out

    // apiResp should still be returned so callers can inspect status/body
    if assert.NotNil(t, apiResp) {
        assert.Equal(t, 422, apiResp.StatusCode)
        assert.Equal(t, body, apiResp.Body)
    }

    // Expect an error of type *client.APIError with status/body preserved
    if assert.Error(t, err) {
        if apiErr, ok := err.(*client.APIError); ok {
            assert.Equal(t, 422, apiErr.StatusCode)
            assert.Equal(t, body, apiErr.Body)
        } else {
            t.Fatalf("expected *client.APIError, got %T", err)
        }
    }
}
