package contacts

import (
	"testing"
)

func TestMapToContact_SimpleMapping(t *testing.T) {
	type S struct {
		Email string `contact:"Email"`
		Star  string `fieldValues:"Star"`
		Tags  string `tags:"Tags"`
	}

	src := S{Email: "jdoe@example.com", Star: "Gold", Tags: "VIP,101"}
	fieldIDByName := map[string]string{"Star": "21"}
	tagNameToID := map[string]string{"VIP": "100"}

	c, tagIDs, err := MapToContact(src, fieldIDByName, tagNameToID)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if c.Email != "jdoe@example.com" {
		t.Fatalf("email not mapped, got %q", c.Email)
	}
	if c.FieldValues == nil || len(*c.FieldValues) != 1 {
		t.Fatalf("expected 1 fieldValue, got %+v", c.FieldValues)
	}
	fv := (*c.FieldValues)[0]
	if fv.Field != "21" || fv.Value != "Gold" {
		t.Fatalf("unexpected fieldValue: %+v", fv)
	}
	if len(tagIDs) != 2 || tagIDs[0] != "100" || tagIDs[1] != "101" {
		t.Fatalf("unexpected tagIDs: %+v", tagIDs)
	}
}

func TestMapToContact_MissingFieldID(t *testing.T) {
	type S struct {
		Star string `fieldValues:"Star"`
	}
	src := S{Star: "Gold"}
	_, _, err := MapToContact(src, nil, nil)
	if err == nil {
		t.Fatalf("expected error when no fieldIDByName provided")
	}
}
