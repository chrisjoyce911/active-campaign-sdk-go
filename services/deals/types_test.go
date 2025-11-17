package deals

import (
	"encoding/json"
	"testing"
)

func TestBoolish_UnmarshalAndValue(t *testing.T) {
	cases := []struct {
		name    string
		json    string
		want    bool
		wantErr bool
	}{
		{"true_bool", "true", true, false},
		{"false_bool", "false", false, false},
		{"one_int", "1", true, false},
		{"zero_int", "0", false, false},
		{"null", "null", false, false},
		{"invalid", `"nope"`, false, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var b Boolish
			err := json.Unmarshal([]byte(tc.json), &b)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error for %s", tc.json)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got := b.Value(); got != tc.want {
				t.Fatalf("Value()=%v want %v", got, tc.want)
			}
		})
	}
	// nil receiver safety
	var nb *Boolish
	if nb.Value() != false {
		t.Fatalf("nil Boolish.Value() should be false")
	}
}

func TestIntish_UnmarshalAndValue(t *testing.T) {
	cases := []struct {
		name    string
		json    string
		want    int
		wantErr bool
	}{
		{"number", "5", 5, false},
		{"string_number", `"7"`, 7, false},
		{"empty_string", `""`, 0, false},
		{"null", "null", 0, false},
		{"invalid", `"x"`, 0, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var n Intish
			err := json.Unmarshal([]byte(tc.json), &n)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error for %s", tc.json)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got := n.Value(); got != tc.want {
				t.Fatalf("Value()=%d want %d", got, tc.want)
			}
		})
	}
	// nil receiver safety
	var nn *Intish
	if nn.Value() != 0 {
		t.Fatalf("nil Intish.Value() should be 0")
	}
}
