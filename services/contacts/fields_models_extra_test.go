package contacts

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
)

func TestFieldPayload_UnmarshalJSON_BadTimesAndNilAccessors(t *testing.T) {
	// Bad timestamp formats should not cause Unmarshal to fail; times remain nil
	b := []byte(`{"field":{"id":"1","cdate":"badtime","udate":""}}`)
	var fr FieldResponse
	if err := json.Unmarshal(b, &fr); err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}

	// The FieldPayload CDate/UDate should be nil; accessors return zero
	if fr.Field.CDateOrZero().IsZero() == false {
		t.Fatalf("expected zero CDate")
	}
	if fr.Field.UDateOrZero().IsZero() == false {
		t.Fatalf("expected zero UDate")
	}

	// Test nil-receiver accessors
	var fp *FieldPayload
	if !fp.CDateOrZero().IsZero() {
		t.Fatalf("expected zero time from nil receiver CDateOrZero")
	}
	if !fp.UDateOrZero().IsZero() {
		t.Fatalf("expected zero time from nil receiver UDateOrZero")
	}
}

func TestFieldOptionPayload_UnmarshalJSON_NilAndBad(t *testing.T) {
	b := []byte(`{"fieldOption":{"id":"o1","cdate":"2020-01-01T00:00:00Z","udate":"2020-01-02T00:00:00Z"}}`)
	var fop FieldOptionResponse
	if err := json.Unmarshal(b, &fop); err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}
	// CDate should parse correctly
	if fop.FieldOption.CDateOrZero().IsZero() {
		t.Fatalf("expected non-zero CDate")
	}

	if fop.FieldOption.UDateOrZero().IsZero() {
		t.Fatalf("expected non-zero UDate")
	}

	// nil receiver
	var fo *FieldOptionPayload
	if !fo.CDateOrZero().IsZero() {
		t.Fatalf("expected zero from nil receiver CDateOrZero")
	}
	if !fo.UDateOrZero().IsZero() {
		t.Fatalf("expected zero from nil receiver UDateOrZero")
	}
}

func TestRealService_UpdateOrCreateFieldValueForContact_ResolveByTitle_UpdateExisting(t *testing.T) {
	lf := &ListFieldsResponse{Fields: &[]FieldPayload{{ID: "77", Perstag: "ptag", Title: "TargetTitle"}}}
	fieldsBody, _ := json.Marshal(lf)

	fvList := &ListFieldValuesResponse{FieldValues: &[]FieldValuePayload{{ID: "fv77", Field: "77", Value: "old"}}}
	fvBody, _ := json.Marshal(fvList)

	putBody := []byte(`{"fieldValue":{"id":"fv77","value":"new"}}`)
	td := &testDoer{getBody: fvBody, putBody: putBody}
	dw := &doerWrapper{testDoer: td, fieldsBody: fieldsBody}
	svc := NewRealServiceFromDoer(dw)

	out, apiResp, err := svc.UpdateOrCreateFieldValueForContact(context.Background(), "c1", "TargetTitle", "new")
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if apiResp == nil || apiResp.StatusCode != 200 {
		t.Fatalf("expected 200, got %+v", apiResp)
	}
	if out == nil || out.FieldValue.ID != "fv77" {
		t.Fatalf("unexpected output: %+v", out)
	}
}

func TestUpdateListStatus_Wrapper_Contacts(t *testing.T) {
	hd := &testhelpers.HTTPDoer{BaseURL: "https://example.api-us1.com/api/3/", Token: "tok", RespStatus: 200, RespBody: []byte(`{"contactList":{"listid":1,"status":1}}`)}
	svc := NewRealServiceFromDoer(hd)

	req := &UpdateListStatusForContactRequest{ContactList: &ContactList{Contact: "1", List: ListID("1"), Status: 1}}
	out, apiResp, err := svc.UpdateListStatus(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if apiResp == nil || apiResp.StatusCode != 200 {
		t.Fatalf("expected 200, got %+v", apiResp)
	}
	if out == nil {
		t.Fatalf("expected non-nil out")
	}
	// ensure request went to contactLists endpoint
	if hd.LastRequest == nil {
		t.Fatalf("expected last request recorded")
	}
}

func TestFieldPayload_UnmarshalJSON_AllTimestamps(t *testing.T) {
	// Use multiple timestamp formats to exercise all parse branches
	b := []byte(`{"field":{"id":"1","cdate":"2020-01-01T00:00:00Z","udate":"2006-01-02 15:04:05","created_timestamp":"2006-01-02 15:04:05","updated_timestamp":"2006-01-02T15:04:05-07:00"}}`)
	var fr FieldResponse
	if err := json.Unmarshal(b, &fr); err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}

	if fr.Field.CDateOrZero().IsZero() {
		t.Fatalf("expected non-zero CDate")
	}
	if fr.Field.UDateOrZero().IsZero() {
		t.Fatalf("expected non-zero UDate")
	}
	if fr.Field.CreatedTimestamp == nil {
		t.Fatalf("expected CreatedTimestamp to be set")
	}
	if fr.Field.UpdatedTimestamp == nil {
		t.Fatalf("expected UpdatedTimestamp to be set")
	}
}

func TestFieldPayload_UnmarshalJSON_BadJSON_ReturnsError(t *testing.T) {
	var fp FieldPayload
	if err := fp.UnmarshalJSON([]byte("not-a-json")); err == nil {
		t.Fatalf("expected error for bad json")
	}
}

func TestFieldOptionPayload_UnmarshalJSON_BadJSON_ReturnsError(t *testing.T) {
	var fo FieldOptionPayload
	if err := fo.UnmarshalJSON([]byte("{bad json")); err == nil {
		t.Fatalf("expected error for bad json in FieldOptionPayload")
	}
}
