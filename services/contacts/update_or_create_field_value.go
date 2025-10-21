package contacts

import (
	"context"
	"net/http"
	"strings"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// UpdateOrCreateFieldValueForContact will set a custom field value for a contact.
// fieldIdentifier may be a numeric id string, a perstag, or a title. The method
// will attempt to resolve the canonical field id, then find any existing
// FieldValue for that contact+field and update it (PUT /fieldValues/{id}),
// otherwise create a new FieldValue (POST /fieldValues).
func (s *RealService) UpdateOrCreateFieldValueForContact(ctx context.Context, contactID, fieldIdentifier, value string) (*FieldValueResponse, *client.APIResponse, error) {
	// quick path: if fieldIdentifier looks like a numeric id, prefer it
	fieldID := strings.TrimSpace(fieldIdentifier)

	// If not a plausible id, try to detect by perstag/title
	if fieldID == "" || !isAllDigits(fieldID) {
		// call ListCustomFields to search by perstag/title (limit larger)
		lf, _, err := s.ListCustomFieldsWithOpts(ctx, map[string]string{"limit": "100"})
		if err == nil && lf != nil {
			wantLower := strings.ToLower(strings.TrimSpace(fieldIdentifier))
			for _, f := range lf.FieldsOrEmpty() {
				if strings.ToLower(strings.TrimSpace(f.Perstag)) == wantLower || strings.ToLower(strings.TrimSpace(f.Title)) == wantLower {
					fieldID = f.ID
					break
				}
			}
		}
	}

	// If we still don't have an id, return an error via POST attempt (let server validate)
	// Next, inspect existing fieldValues for the contact to see if a FieldValue record exists
	if contactID == "" {
		return nil, &client.APIResponse{StatusCode: http.StatusBadRequest}, nil
	}

	fvList, _, err := s.GetContactFieldValuesTyped(ctx, contactID)
	if err == nil && fvList != nil {
		for _, fv := range fvList.FieldValuesOrEmpty() {
			// match either by exact field id or by perstag/title fallback
			if fieldID != "" && fv.Field == fieldID {
				// update this fieldValue by id
				req := &FieldValuePayload{Contact: contactID, Field: fieldID, Value: value}
				out, resp, err := s.UpdateFieldValueByID(ctx, fv.ID, req)
				return out, resp, err
			}
			// if fieldID is empty, try matching by field reference
			if fieldID == "" && fv.Field == fieldIdentifier {
				req := &FieldValuePayload{Contact: contactID, Field: fv.Field, Value: value}
				out, resp, err := s.UpdateFieldValueByID(ctx, fv.ID, req)
				return out, resp, err
			}
		}
	}

	// No existing fieldValue found; attempt to create one via POST /fieldValues
	postReq := &FieldValuePayload{Contact: contactID, Field: fieldID, Value: value}
	out, resp, err := s.UpdateFieldValueForContact(ctx, postReq)
	return out, resp, err
}

// isAllDigits reports whether s consists only of digits.
func isAllDigits(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
