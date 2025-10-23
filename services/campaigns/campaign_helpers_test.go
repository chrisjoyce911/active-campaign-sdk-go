package campaigns

import (
	"testing"
)

func TestStatusHelpers_EmptyAndInvalid(t *testing.T) {
	var c Campaign

	// empty status should now return an error from StatusInt
	_, err := c.StatusInt()
	if err == nil {
		t.Fatalf("expected error parsing empty status, got nil")
	}
	if c.StatusEnum() != CampaignStatusDraft {
		t.Fatalf("expected Draft for empty status")
	}

	// invalid string should return error from StatusInt and Draft from StatusEnum
	c.Status = "notanint"
	_, err = c.StatusInt()
	if err == nil {
		t.Fatalf("expected error parsing invalid status")
	}
	if c.StatusEnum() != CampaignStatusDraft {
		t.Fatalf("expected Draft for invalid status")
	}
}

func TestStatusHelpers_NumericValues(t *testing.T) {
	cases := []struct {
		in      string
		want    CampaignStatus
		wantErr bool
	}{
		{"0", CampaignStatusDraft, false},
		{"1", CampaignStatusScheduled, false},
		{"2", CampaignStatusSending, false},
		{"3", CampaignStatusPaused, false},
		{"4", CampaignStatusStopped, false},
		{"5", CampaignStatusCompleted, false},
		// edge cases
		{"10", CampaignStatusDraft, false},         // multi-digit: unknown -> Draft
		{"-1", CampaignStatusDraft, false},         // negative -> unknown -> Draft
		{"9999999999", CampaignStatusDraft, false}, // large number parses on 64-bit, enum should fallback to Draft
	}

	for _, tc := range cases {
		var c Campaign
		c.Status = tc.in
		i, err := c.StatusInt()
		if tc.wantErr {
			if err == nil {
				t.Fatalf("expected error parsing %q, got nil", tc.in)
			}
			// ensure StatusEnum still falls back to Draft
			if c.StatusEnum() != tc.want {
				t.Fatalf("for %q expected enum %v got %v", tc.in, tc.want, c.StatusEnum())
			}
			continue
		}
		if err != nil {
			t.Fatalf("unexpected error parsing %q: %v", tc.in, err)
		}
		// StatusEnum should map known single-digit values; unknowns -> Draft
		if c.StatusEnum() != tc.want {
			t.Fatalf("for %q expected enum %v got %v", tc.in, tc.want, c.StatusEnum())
		}
		// For known single-digit statuses, ensure numeric conversion matches
		if len(tc.in) == 1 && tc.in[0] >= '0' && tc.in[0] <= '9' {
			if i != int(tc.in[0]-'0') {
				t.Fatalf("for %q expected int %d got %d", tc.in, int(tc.in[0]-'0'), i)
			}
		}
	}
}
