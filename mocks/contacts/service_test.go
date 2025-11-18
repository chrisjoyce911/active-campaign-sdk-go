package contactsmock

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/contacts"
)

func TestService_SearchByEmailFuncCalled(t *testing.T) {
	called := false
	m := &Service{SearchByEmailFunc: func(ctx context.Context, email string) (*contacts.ContactSearchResponse, *client.APIResponse, error) {
		called = true
		return &contacts.ContactSearchResponse{Contacts: []contacts.Contact{{ID: "c1"}}}, &client.APIResponse{StatusCode: 200}, nil
	}}
	resp, api, err := m.SearchByEmail(context.Background(), "a@b.com")
	if err != nil || api == nil || resp == nil || len(resp.Contacts) != 1 || !called {
		t.Fatalf("unexpected result: resp=%v api=%v err=%v called=%v", resp, api, err, called)
	}
}

func TestService_Defaults(t *testing.T) {
	m := &Service{}
	// Defaults should return non-nil values where applicable
	resp, api, err := m.SearchByEmail(context.Background(), "a@b.com")
	if err != nil || resp == nil || api == nil {
		t.Fatalf("expected non-nil defaults, got resp=%v api=%v err=%v", resp, api, err)
	}
	fv, fvAPI, fvErr := m.UpdateOrCreateFieldValueForContact(context.Background(), "c1", "cf1", "val")
	if fv == nil || fvAPI == nil || fvErr != nil {
		t.Fatalf("expected non-nil defaults for UpdateOrCreateFieldValueForContact")
	}
	upAPI, upErr := m.UpdateFieldGroup(context.Background(), "g1", nil)
	if upErr != nil || upAPI == nil {
		t.Fatalf("expected non-nil defaults for UpdateFieldGroup")
	}
}
