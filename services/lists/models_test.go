package lists

import (
	"encoding/json"
	"testing"
)

func TestStringOrInt_UnmarshalJSON(t *testing.T) {
	var s StringOrInt

	if err := json.Unmarshal([]byte(`"abc"`), &s); err != nil {
		t.Fatalf("unexpected error unmarshalling quoted string: %v", err)
	}
	if string(s) != "abc" {
		t.Fatalf("expected 'abc', got %q", string(s))
	}

	if err := json.Unmarshal([]byte(`12345`), &s); err != nil {
		t.Fatalf("unexpected error unmarshalling number: %v", err)
	}
	if string(s) != "12345" {
		t.Fatalf("expected '12345', got %q", string(s))
	}

	if err := json.Unmarshal([]byte(`null`), &s); err != nil {
		t.Fatalf("unexpected error unmarshalling null: %v", err)
	}
	if string(s) != "" {
		t.Fatalf("expected empty string for null, got %q", string(s))
	}
}

func TestIntOrString_UnmarshalJSON(t *testing.T) {
	var i IntOrString

	if err := json.Unmarshal([]byte(`"42"`), &i); err != nil {
		t.Fatalf("unexpected error unmarshalling quoted numeric string: %v", err)
	}
	if int(i) != 42 {
		t.Fatalf("expected 42, got %d", int(i))
	}

	if err := json.Unmarshal([]byte(`99`), &i); err != nil {
		t.Fatalf("unexpected error unmarshalling number: %v", err)
	}
	if int(i) != 99 {
		t.Fatalf("expected 99, got %d", int(i))
	}

	// Quoted non-numeric should return an error
	var ierr IntOrString
	if err := json.Unmarshal([]byte(`"notanint"`), &ierr); err == nil {
		t.Fatalf("expected error unmarshalling quoted non-numeric, got nil")
	}

	// Very large number should eventually error through strconv.Atoi path
	var big IntOrString
	if err := json.Unmarshal([]byte(`999999999999999999999999999999`), &big); err == nil {
		t.Fatalf("expected error unmarshalling overly-large number, got nil")
	}
}

func TestStringToUser(t *testing.T) {
	got := StringToUser("user123")
	if string(got) != "user123" {
		t.Fatalf("StringToUser returned %q, want %q", string(got), "user123")
	}
}

func TestUnmarshal_EmptyAndNull(t *testing.T) {
	var s StringOrInt
	if err := s.UnmarshalJSON([]byte{}); err != nil {
		t.Fatalf("unexpected error for empty bytes StringOrInt: %v", err)
	}
	if string(s) != "" {
		t.Fatalf("expected empty string from empty bytes, got %q", string(s))
	}

	var i IntOrString
	if err := i.UnmarshalJSON([]byte{}); err != nil {
		t.Fatalf("unexpected error for empty bytes IntOrString: %v", err)
	}
	if int(i) != 0 {
		t.Fatalf("expected zero IntOrString from empty bytes, got %d", int(i))
	}

	if err := s.UnmarshalJSON([]byte("null")); err != nil {
		t.Fatalf("unexpected error for null StringOrInt: %v", err)
	}
	if string(s) != "" {
		t.Fatalf("expected empty string from null, got %q", string(s))
	}

	if err := i.UnmarshalJSON([]byte("null")); err != nil {
		t.Fatalf("unexpected error for null IntOrString: %v", err)
	}
	if int(i) != 0 {
		t.Fatalf("expected zero from null IntOrString, got %d", int(i))
	}
}

func TestIntOrString_InvalidNumberPaths(t *testing.T) {
	var i IntOrString
	// invalid JSON token (not quoted, not number) should return an error
	if err := i.UnmarshalJSON([]byte("notjson")); err == nil {
		t.Fatalf("expected error unmarshalling invalid token, got nil")
	}
}

func TestQuotedUnmarshal_ErrorPaths(t *testing.T) {
	var s StringOrInt
	// malformed JSON (unterminated quote) should produce an error
	if err := s.UnmarshalJSON([]byte("\"unterminated")); err == nil {
		t.Fatalf("expected error unmarshalling malformed quoted StringOrInt, got nil")
	}

	var i IntOrString
	// malformed JSON for number should produce an error
	if err := i.UnmarshalJSON([]byte("{notvalidjson}")); err == nil {
		t.Fatalf("expected error unmarshalling malformed IntOrString, got nil")
	}
}

func TestInvalidEscapeSequences(t *testing.T) {
	var s StringOrInt
	// invalid unicode escape inside quoted string
	if err := s.UnmarshalJSON([]byte("\"\\uZZZZ\"")); err == nil {
		t.Fatalf("expected error for invalid escape in StringOrInt, got nil")
	}

	var i IntOrString
	if err := i.UnmarshalJSON([]byte("\"\\uZZZZ\"")); err == nil {
		t.Fatalf("expected error for invalid escape in IntOrString, got nil")
	}
}

func TestIntOrString_JsonNumberFallback(t *testing.T) {
	// Create a numeric value that is too large to decode directly into int on some
	// platforms so that json.Unmarshal into int fails but into json.Number succeeds.
	largeNum := `92233720368547758079223372036854775807`
	var i IntOrString
	if err := i.UnmarshalJSON([]byte(largeNum)); err == nil {
		t.Fatalf("expected error or fallback behaviour for very large number, got nil")
	}
}
