package contacts

import (
	"encoding/json"
	"testing"
	"time"
)

func TestParseTimePtr_NilAndEmpty(t *testing.T) {
	if tp, err := parseTimePtr(nil); tp != nil || err != nil {
		t.Fatalf("expected nil,nil for nil input; got %v,%v", tp, err)
	}
	empty := ""
	if tp, err := parseTimePtr(&empty); tp != nil || err != nil {
		t.Fatalf("expected nil,nil for empty string; got %v,%v", tp, err)
	}
}

func TestParseTimePtr_SuccessAndFail(t *testing.T) {
	// RFC3339
	s := "2020-01-02T15:04:05Z"
	tp, err := parseTimePtr(&s)
	if err != nil || tp == nil {
		t.Fatalf("expected parsed time, got %v,%v", tp, err)
	}

	// Unsupported format should return an error
	bad := "not a time"
	tp2, err2 := parseTimePtr(&bad)
	if err2 == nil || tp2 != nil {
		t.Fatalf("expected error for bad time, got %v,%v", tp2, err2)
	}
}

func TestFieldPayload_UnmarshalJSON_Times(t *testing.T) {
	// Use multiple layouts to ensure each parsing path is hit
	payload := `{"cdate":"2020-01-02T15:04:05Z","udate":"2006-01-02 15:04:05","created_timestamp":"2006-01-02T15:04:05-07:00","updated_timestamp":""}`
	var f FieldPayload
	if err := json.Unmarshal([]byte(payload), &f); err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}
	if f.CDate == nil {
		t.Fatalf("expected CDate to be parsed")
	}
	if f.UDate == nil {
		t.Fatalf("expected UDate to be parsed")
	}
	// created timestamp should parse
	if f.CreatedTimestamp == nil {
		t.Fatalf("expected CreatedTimestamp to be parsed")
	}
	// updated timestamp was empty string so should remain nil
	if f.UpdatedTimestamp != nil {
		t.Fatalf("expected UpdatedTimestamp to be nil for empty string")
	}
}

func TestFieldOptionPayload_UnmarshalJSON_Times(t *testing.T) {
	payload := `{"cdate":"2020-01-02T15:04:05Z","udate":""}`
	var fo FieldOptionPayload
	if err := json.Unmarshal([]byte(payload), &fo); err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}
	if fo.CDate == nil {
		t.Fatalf("expected CDate to be parsed")
	}
	if fo.UDate != nil {
		t.Fatalf("expected UDate to be nil for empty string")
	}
}

func TestCDateOrZero_UDateOrZero(t *testing.T) {
	var f *FieldPayload
	if z := f.CDateOrZero(); !z.IsZero() {
		t.Fatalf("expected zero time for nil receiver")
	}
	if z := f.UDateOrZero(); !z.IsZero() {
		t.Fatalf("expected zero time for nil receiver")
	}

	// non-nil receiver
	now := time.Now().UTC().Truncate(time.Second)
	f2 := &FieldPayload{CDate: &now, UDate: &now}
	if z := f2.CDateOrZero(); !z.Equal(now) {
		t.Fatalf("expected CDateOrZero to return stored time")
	}
	if z := f2.UDateOrZero(); !z.Equal(now) {
		t.Fatalf("expected UDateOrZero to return stored time")
	}
}

func TestListAccessors_EmptyAndNonEmpty(t *testing.T) {
	var lr *ListFieldsResponse
	if got := lr.FieldsOrEmpty(); len(got) != 0 {
		t.Fatalf("expected empty slice for nil receiver")
	}
	// nil pointer inside
	lr2 := &ListFieldsResponse{Fields: nil}
	if got := lr2.FieldsOrEmpty(); len(got) != 0 {
		t.Fatalf("expected empty slice for nil pointer field")
	}
	// non-empty
	f := FieldPayload{ID: "1"}
	lr3 := &ListFieldsResponse{Fields: &[]FieldPayload{f}}
	if got := lr3.FieldsOrEmpty(); len(got) != 1 || got[0].ID != "1" {
		t.Fatalf("unexpected fields content: %#v", got)
	}

	var lfv *ListFieldValuesResponse
	if got := lfv.FieldValuesOrEmpty(); len(got) != 0 {
		t.Fatalf("expected empty slice for nil receiver")
	}
	fv := FieldValuePayload{ID: "10"}
	lfv2 := &ListFieldValuesResponse{FieldValues: &[]FieldValuePayload{fv}}
	if got := lfv2.FieldValuesOrEmpty(); len(got) != 1 || got[0].ID != "10" {
		t.Fatalf("unexpected field values content: %#v", got)
	}
}
