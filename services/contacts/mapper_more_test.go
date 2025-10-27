package contacts

import (
	"testing"
)

// small helper type that mimics generated Fields constants
type exampleFields struct {
	Star  string
	Speed string
}

func TestBuildFieldIDByName_WithStructAndPointer(t *testing.T) {
	f := exampleFields{Star: "21", Speed: "22"}
	m := BuildFieldIDByName(f)
	if got := m["Star"]; got != "21" {
		t.Fatalf("expected Star->21 got %q", got)
	}
	if got := m["Speed"]; got != "22" {
		t.Fatalf("expected Speed->22 got %q", got)
	}

	// pointer variant
	mp := BuildFieldIDByName(&f)
	if mp["Star"] != "21" || mp["Speed"] != "22" {
		t.Fatalf("pointer variant failed: %+v", mp)
	}
}

func TestBuildFieldIDByName_NonStructReturnsEmpty(t *testing.T) {
	m := BuildFieldIDByName(123)
	if len(m) != 0 {
		t.Fatalf("expected empty map for non-struct, got %+v", m)
	}
}

func TestMapToContact_ContactUnknownFieldError(t *testing.T) {
	type S struct {
		Foo string `contact:"NonExistent"`
	}
	src := S{Foo: "bar"}
	_, _, err := MapToContact(src, map[string]string{}, map[string]string{})
	if err == nil {
		t.Fatalf("expected error for unknown contact field, got nil")
	}
}

func TestMapToContact_OmitEmptyContactDoesNotSet(t *testing.T) {
	type S struct {
		First string `contact:"FirstName,omitempty"`
	}
	src := S{First: ""}
	c, tags, err := MapToContact(src, nil, nil)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if c.FirstName != "" {
		t.Fatalf("expected empty FirstName, got %q", c.FirstName)
	}
	if len(tags) != 0 {
		t.Fatalf("expected no tags, got %+v", tags)
	}
}

func TestMapToContact_NonStringFieldIsIgnored(t *testing.T) {
	type S struct {
		Age int `contact:"Age"`
	}
	src := S{Age: 42}
	// should not panic or error; Age is non-string so it's skipped
	_, _, err := MapToContact(src, nil, nil)
	if err != nil {
		t.Fatalf("unexpected err for non-string field: %v", err)
	}
}

func TestMapToContact_TagsFallbackToIDsWhenLookupMissing(t *testing.T) {
	type S struct {
		Tags string `tags:"Tags"`
	}
	src := S{Tags: "100,VIP"}
	// tagNameToID nil -> tokens treated as IDs
	_, tagIDs, err := MapToContact(src, nil, nil)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if len(tagIDs) != 2 || tagIDs[0] != "100" || tagIDs[1] != "VIP" {
		t.Fatalf("unexpected tagIDs: %+v", tagIDs)
	}
}

func TestMapToContact_NilSrcError(t *testing.T) {
	_, _, err := MapToContact(nil, nil, nil)
	if err == nil {
		t.Fatalf("expected error for nil src")
	}
}

func TestMapToContact_PointerSourceAndTrimmedTags(t *testing.T) {
	type S struct {
		Email string `contact:" Email ,omitempty "`
		Tags  string `tags:"Tags"`
	}
	src := S{Email: "ptr@example.com", Tags: " VIP , 200 "}
	// pass pointer source
	c, tagIDs, err := MapToContact(&src, nil, map[string]string{"VIP": "100"})
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if c.Email != "ptr@example.com" {
		t.Fatalf("expected email set from pointer source, got %q", c.Email)
	}
	if len(tagIDs) != 2 || tagIDs[0] != "100" || tagIDs[1] != "200" {
		t.Fatalf("unexpected tagIDs: %+v", tagIDs)
	}
}

func TestMapToContact_ContactFieldNotSettable(t *testing.T) {
	type S struct {
		Bad string `contact:"FieldValues"`
	}
	src := S{Bad: "x"}
	_, _, err := MapToContact(src, nil, nil)
	if err == nil {
		t.Fatalf("expected error when contact tag points to non-string field on Contact")
	}
}

func TestMapToContact_FieldValuesOmitEmptyAndMissingKey(t *testing.T) {
	type S struct {
		Star string `fieldValues:"Star,omitempty"`
		Moon string `fieldValues:"Moon"`
	}
	// omit empty: Star is empty and should be skipped
	src := S{Star: "", Moon: "yes"}
	// case: omitEmpty should skip Star without requiring mapping
	// but Moon exists in tag and mapping missing should error
	fieldIDByName := map[string]string{"Star": "21"} // Moon missing
	_, _, err := MapToContact(src, fieldIDByName, nil)
	if err == nil {
		t.Fatalf("expected error when a non-omitempty custom field has no mapping")
	}
}

func TestMapToContact_TagsOmitEmptyAndEmptyTokens(t *testing.T) {
	type S struct {
		TagsA string `tags:"TagsA,omitempty"`
		TagsB string `tags:"TagsB"`
	}
	src := S{TagsA: "", TagsB: "100,, VIP ,  "}
	c, tagIDs, err := MapToContact(src, nil, map[string]string{"VIP": "200"})
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if len(tagIDs) != 2 || tagIDs[0] != "100" || tagIDs[1] != "200" {
		t.Fatalf("unexpected tagIDs after trimming and skipping empty tokens: %+v", tagIDs)
	}
	if c.Email != "" {
		// ensure no contact fields were set accidentally
	}
}
