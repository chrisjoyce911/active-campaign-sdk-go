package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// BulkImportMaxContacts is the number of contacts ActiveCampaign accepts in a
// single bulk import request.
const BulkImportMaxContacts = 250

// BulkImportContact is one contact in a bulk import request.
//
// Note the field names differ from the single-contact endpoints: bulk import
// uses snake_case (first_name) where contact/sync uses camelCase (firstName).
type BulkImportContact struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Phone     string `json:"phone,omitempty"`

	// Tags are added; bulk import cannot remove a tag.
	Tags []string `json:"tags,omitempty"`

	// Fields are custom field values, keyed by field ID.
	Fields []BulkImportField `json:"fields,omitempty"`

	// Subscribe and Unsubscribe manage list membership by list ID.
	Subscribe   []BulkImportList `json:"subscribe,omitempty"`
	Unsubscribe []BulkImportList `json:"unsubscribe,omitempty"`
}

// BulkImportField is a custom field value on a bulk-imported contact.
type BulkImportField struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

// BulkImportList is a list membership change on a bulk-imported contact.
type BulkImportList struct {
	ListID int `json:"listid"`
}

// BulkImportCallback asks ActiveCampaign to POST the outcome somewhere when
// the batch finishes, instead of the caller polling BulkImportStatus.
type BulkImportCallback struct {
	URL     string              `json:"url"`
	Method  string              `json:"requestType,omitempty"`
	Params  []map[string]string `json:"params,omitempty"`
	Headers []map[string]string `json:"headers,omitempty"`
}

// BulkImportRequest is the payload for a bulk import.
type BulkImportRequest struct {
	Contacts []BulkImportContact `json:"contacts"`
	Callback *BulkImportCallback `json:"callback,omitempty"`
}

// BulkImportResponse acknowledges a bulk import. The work is asynchronous:
// Success only means the batch was accepted, so BatchID must be polled with
// BulkImportStatus to learn whether the contacts actually imported.
type BulkImportResponse struct {
	Success        int    `json:"success"`
	QueuedContacts int    `json:"queued_contacts"`
	BatchID        string `json:"batchId"`
	Message        string `json:"message,omitempty"`
	FailureReasons []any  `json:"failureReasons,omitempty"`
}

// BulkImportStatusResponse reports how a submitted batch fared.
type BulkImportStatusResponse struct {
	Success  int    `json:"success"`
	Status   string `json:"status,omitempty"`
	BatchID  string `json:"batchId,omitempty"`
	Message  string `json:"message,omitempty"`
	Contacts struct {
		Total   int `json:"total"`
		Success int `json:"success"`
		Failure int `json:"failure"`
	} `json:"contacts"`
	FailureReasons []any `json:"failureReasons,omitempty"`
}

// BulkImportContacts submits up to BulkImportMaxContacts contacts in one
// request.
//
// What & Why:
//
//	Importing contacts one at a time costs a request each; this endpoint takes
//	250 at once, which is the difference between hours and minutes for a large
//	reconciliation.
//
//	It is asynchronous — a 2xx means the batch was accepted, not that the
//	contacts imported. Poll BulkImportStatus with the returned BatchID, or
//	supply a Callback.
//
//	Capability differs from the per-contact endpoints: tags can be added but
//	not removed, and list membership is subscribe/unsubscribe only. Work that
//	must remove a tag still needs TagRemove per contact.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference/bulk-import
//
//	The path is import/bulk_import. An earlier version of this method posted
//	to bulkImport, which 404s.
//
// Parameters:
//
//	ctx: context for cancellation/timeouts
//	req: the contacts to import
//
// Returns:
//
//	(*BulkImportResponse, *client.APIResponse, error)
func (s *RealService) BulkImportContacts(ctx context.Context, req *BulkImportRequest) (*BulkImportResponse, *client.APIResponse, error) {
	out := &BulkImportResponse{}
	apiResp, err := s.client.Do(ctx, http.MethodPost, "import/bulk_import", req, out)
	return out, apiResp, err
}

// BulkImportStatus reports on a batch previously accepted by
// BulkImportContacts.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference/bulk-import-status
func (s *RealService) BulkImportStatus(ctx context.Context, batchID string) (*BulkImportStatusResponse, *client.APIResponse, error) {
	out := &BulkImportStatusResponse{}
	apiResp, err := s.client.Do(ctx, http.MethodGet, "import/info", map[string]string{"batchId": batchID}, out)
	return out, apiResp, err
}
