package contacts

import "testing"

func TestCreateContactResponseAccessors_NilReceiver(t *testing.T) {
	var c *CreateContactResponse
	if got := c.ContactAutomationsOrEmpty(); len(got) != 0 {
		t.Fatalf("ContactAutomationsOrEmpty nil receiver = %v, want empty", got)
	}
	if got := c.ContactDataOrEmpty(); len(got) != 0 {
		t.Fatalf("ContactDataOrEmpty nil receiver = %v, want empty", got)
	}
	if got := c.ContactListsOrEmpty(); len(got) != 0 {
		t.Fatalf("ContactListsOrEmpty nil receiver = %v, want empty", got)
	}
	if got := c.FieldValuesOrEmpty(); len(got) != 0 {
		t.Fatalf("FieldValuesOrEmpty nil receiver = %v, want empty", got)
	}
	if got := c.GeoAddressesOrEmpty(); len(got) != 0 {
		t.Fatalf("GeoAddressesOrEmpty nil receiver = %v, want empty", got)
	}
	if got := c.GeoIpsOrEmpty(); len(got) != 0 {
		t.Fatalf("GeoIpsOrEmpty nil receiver = %v, want empty", got)
	}
	if got := c.ScoreValuesOrEmpty(); len(got) != 0 {
		t.Fatalf("ScoreValuesOrEmpty nil receiver = %v, want empty", got)
	}
}

func TestCreateContactResponseAccessors_NonNil(t *testing.T) {
	ca := ContactAutomation{ID: "a1"}
	cd := ContactData{ID: "d1"}
	cl := ContactList{ID: "l1"}
	fv := FieldValue{ID: "f1"}
	ga := GeoAddress{ID: "g1"}
	gi := GeoIp{ID: "gi1"}
	sv := ScoreValue{ID: "s1"}

	c := &CreateContactResponse{
		ContactAutomations: &[]ContactAutomation{ca},
		ContactData:        &[]ContactData{cd},
		ContactLists:       &[]ContactList{cl},
		FieldValues:        &[]FieldValue{fv},
		GeoAddresses:       &[]GeoAddress{ga},
		GeoIps:             &[]GeoIp{gi},
		ScoreValues:        &[]ScoreValue{sv},
	}

	if got := c.ContactAutomationsOrEmpty(); len(got) != 1 || got[0].ID != "a1" {
		t.Fatalf("ContactAutomationsOrEmpty = %v, want length 1 with ID a1", got)
	}
	if got := c.ContactDataOrEmpty(); len(got) != 1 || got[0].ID != "d1" {
		t.Fatalf("ContactDataOrEmpty = %v, want length 1 with ID d1", got)
	}
	if got := c.ContactListsOrEmpty(); len(got) != 1 || got[0].ID != "l1" {
		t.Fatalf("ContactListsOrEmpty = %v, want length 1 with ID l1", got)
	}
	if got := c.FieldValuesOrEmpty(); len(got) != 1 || got[0].ID != "f1" {
		t.Fatalf("FieldValuesOrEmpty = %v, want length 1 with ID f1", got)
	}
	if got := c.GeoAddressesOrEmpty(); len(got) != 1 || got[0].ID != "g1" {
		t.Fatalf("GeoAddressesOrEmpty = %v, want length 1 with ID g1", got)
	}
	if got := c.GeoIpsOrEmpty(); len(got) != 1 || got[0].ID != "gi1" {
		t.Fatalf("GeoIpsOrEmpty = %v, want length 1 with ID gi1", got)
	}
	if got := c.ScoreValuesOrEmpty(); len(got) != 1 || got[0].ID != "s1" {
		t.Fatalf("ScoreValuesOrEmpty = %v, want length 1 with ID s1", got)
	}
}

func TestContactTagsAndSearchResponseAccessors(t *testing.T) {
	var ct *ContactTagsResponse
	if got := ct.ContactTagsOrEmpty(); len(got) != 0 {
		t.Fatalf("ContactTagsOrEmpty nil receiver = %v, want empty", got)
	}

	ctr := &ContactTagsResponse{ContactTags: &[]ContactTag{{ID: "tag1"}}}
	if got := ctr.ContactTagsOrEmpty(); len(got) != 1 || got[0].ID != "tag1" {
		t.Fatalf("ContactTagsOrEmpty = %v, want ID tag1", got)
	}

	var sr *ContactSearchResponse
	if got := sr.ScoreValuesOrEmpty(); len(got) != 0 {
		t.Fatalf("ScoreValuesOrEmpty nil receiver = %v, want empty", got)
	}
	srr := &ContactSearchResponse{ScoreValues: &[]ScoreValue{{ID: "sv1"}}}
	if got := srr.ScoreValuesOrEmpty(); len(got) != 1 || got[0].ID != "sv1" {
		t.Fatalf("ScoreValuesOrEmpty = %v, want ID sv1", got)
	}
}
