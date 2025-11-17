package contacts

import "testing"

func TestContactListsResponse_Accessors(t *testing.T) {
	// nil receiver returns empty slice
	var r *ContactListsResponse
	if got := r.ContactListsOrEmpty(); len(got) != 0 {
		t.Fatalf("ContactListsOrEmpty nil receiver = %v, want empty", got)
	}

	// non-nil with populated slice returns underlying data
	items := []ContactList{{ID: "l1"}}
	r = &ContactListsResponse{ContactLists: &items}
	got := r.ContactListsOrEmpty()
	if len(got) != 1 || got[0].ID != "l1" {
		t.Fatalf("ContactListsOrEmpty = %v, want length 1 with ID l1", got)
	}
}
