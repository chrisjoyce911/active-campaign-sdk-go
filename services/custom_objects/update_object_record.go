package custom_objects

import (
	"context"
	"errors"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// UpdateObjectRecord updates an existing record for a custom object type.
// PUT /customObjects/schemas/{id}/records/{recordId}
func (s *service) UpdateObjectRecord(ctx context.Context, objectTypeID, recordID string, req *UpdateRecordRequest) (*UpdateRecordResponse, *client.APIResponse, error) {
	if s.client == nil {
		return nil, nil, errors.New("not implemented")
	}
	var out UpdateRecordResponse
	// Helper to build array form of fields (used by many API shapes).
	type fieldItem struct {
		ID    string      `json:"id"`
		Value interface{} `json:"value"`
	}

	buildArrayFields := func(m map[string]interface{}) []fieldItem {
		items := make([]fieldItem, 0, len(m))
		for k, v := range m {
			items = append(items, fieldItem{ID: k, Value: v})
		}
		return items
	}

	// The ActiveCampaign API exposes a create-or-update endpoint that accepts
	// POST /customObjects/records/{schemaId} with a payload like:
	// { "record": { "id": "<id>", "fields": [{"id":"name","value":"x"}], ... } }
	// Use that endpoint (no recordId in the path) to perform updates.
	if req == nil {
		return &out, nil, nil
	}

	rec := make(map[string]interface{})
	// include id so the POST will update the existing record
	if recordID != "" {
		rec["id"] = recordID
	}
	// schemaId is optional for the POST but harmless to include
	rec["schemaId"] = objectTypeID
	if req.Fields != nil {
		rec["fields"] = buildArrayFields(req.Fields)
	}
	if req.Relationships != nil {
		rec["relationships"] = req.Relationships
	}

	body := map[string]interface{}{"record": rec}
	// POST to create-or-update endpoint for the schema
	postPath := "customObjects/records/" + objectTypeID
	apiResp, err := s.client.Do(ctx, http.MethodPost, postPath, body, &out)
	return &out, apiResp, err
}
