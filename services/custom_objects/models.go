package custom_objects

import (
	"encoding/json"
	"time"
)

// Timestamp wraps a *time.Time and provides tolerant unmarshalling.
type Timestamp struct {
	Time *time.Time
}

var timestampLayouts = []string{
	time.RFC3339Nano,
	time.RFC3339,
	"2006-01-02 15:04:05",
}

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	// allow null
	if string(b) == "null" {
		t.Time = nil
		return nil
	}
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if s == "" {
		t.Time = nil
		return nil
	}
	for _, layout := range timestampLayouts {
		if tt, err := time.Parse(layout, s); err == nil {
			t.Time = &tt
			return nil
		}
	}
	// fallback
	tt, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return err
	}
	t.Time = &tt
	return nil
}

// LabelPair holds singular/plural labels.
type LabelPair struct {
	Singular string `json:"singular,omitempty"`
	Plural   string `json:"plural,omitempty"`
}

type FieldOption struct {
	ID    string `json:"id,omitempty"`
	Value string `json:"value,omitempty"`
}

type SchemaField struct {
	ID              string        `json:"id,omitempty"`
	Labels          LabelPair     `json:"labels,omitempty"`
	Description     *string       `json:"description,omitempty"`
	Type            string        `json:"type,omitempty"`
	Required        bool          `json:"required,omitempty"`
	Inherited       bool          `json:"inherited,omitempty"`
	Scale           *int          `json:"scale,omitempty"`
	DefaultCurrency *string       `json:"defaultCurrency,omitempty"`
	Options         []FieldOption `json:"options,omitempty"`
}

type SchemaRelationship struct {
	ID          string    `json:"id,omitempty"`
	Labels      LabelPair `json:"labels,omitempty"`
	Description *string   `json:"description,omitempty"`
	Namespace   string    `json:"namespace,omitempty"`
	HasMany     bool      `json:"hasMany,omitempty"`
	Inherited   bool      `json:"inherited,omitempty"`
}

type Schema struct {
	ID                           string               `json:"id,omitempty"`
	Slug                         string               `json:"slug,omitempty"`
	ParentID                     *string              `json:"parentId,omitempty"`
	ParentVersion                *int                 `json:"parentVersion,omitempty"`
	Visibility                   string               `json:"visibility,omitempty"`
	Customized                   bool                 `json:"customized,omitempty"`
	Labels                       LabelPair            `json:"labels,omitempty"`
	Description                  *string              `json:"description,omitempty"`
	AppID                        *string              `json:"appId,omitempty"`
	CreatedTimestamp             *Timestamp           `json:"createdTimestamp,omitempty"`
	UpdatedTimestamp             *Timestamp           `json:"updatedTimestamp,omitempty"`
	ParentSchemaCreatedTimestamp *Timestamp           `json:"parentSchemaCreatedTimestamp,omitempty"`
	ParentSchemaUpdatedTimestamp *Timestamp           `json:"parentSchemaUpdatedTimestamp,omitempty"`
	Fields                       []SchemaField        `json:"fields,omitempty"`
	Icons                        map[string]string    `json:"icons,omitempty"`
	Relationships                []SchemaRelationship `json:"relationships,omitempty"`
}

type Meta struct {
	Total  int `json:"total,omitempty"`
	Count  int `json:"count,omitempty"`
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
}

type ListSchemasResponse struct {
	Schemas []Schema `json:"schemas,omitempty"`
	Meta    Meta     `json:"meta,omitempty"`
}

// UnmarshalJSON accepts responses that may use either "schemas" (new) or
// "objectTypes" (legacy) as the top-level key for the list of schemas.
func (r *ListSchemasResponse) UnmarshalJSON(b []byte) error {
	var aux map[string]json.RawMessage
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}
	// try new key
	if v, ok := aux["schemas"]; ok {
		if err := json.Unmarshal(v, &r.Schemas); err != nil {
			return err
		}
	} else if v, ok := aux["objectTypes"]; ok {
		if err := json.Unmarshal(v, &r.Schemas); err != nil {
			return err
		}
	}
	if v, ok := aux["meta"]; ok {
		if err := json.Unmarshal(v, &r.Meta); err != nil {
			return err
		}
	}
	return nil
}

// Record types
type Record struct {
	ID     string                 `json:"id,omitempty"`
	Fields map[string]interface{} `json:"fields,omitempty"`
}

// UnmarshalJSON accepts either an array of {id,value} objects (API response)
// or a map[string]interface{} form. It normalizes both into Fields map.
func (r *Record) UnmarshalJSON(b []byte) error {
	type rawRecord struct {
		ID json.RawMessage `json:"id,omitempty"`
		// keep raw fields
		Fields json.RawMessage `json:"fields,omitempty"`
	}
	var rr rawRecord
	if err := json.Unmarshal(b, &rr); err != nil {
		return err
	}
	// extract ID
	var id string
	if len(rr.ID) > 0 {
		if err := json.Unmarshal(rr.ID, &id); err == nil {
			r.ID = id
		}
	}
	// handle fields which may be an array of {id,value} or a map
	if len(rr.Fields) == 0 {
		r.Fields = nil
		return nil
	}
	// try array form
	var arr []struct {
		ID    string      `json:"id"`
		Value interface{} `json:"value"`
	}
	if err := json.Unmarshal(rr.Fields, &arr); err == nil {
		m := make(map[string]interface{}, len(arr))
		for _, it := range arr {
			m[it.ID] = it.Value
		}
		r.Fields = m
		return nil
	}
	// fallback to map form
	var mp map[string]interface{}
	if err := json.Unmarshal(rr.Fields, &mp); err != nil {
		return err
	}
	r.Fields = mp
	return nil
}

type ListRecordsResponse struct {
	Records []Record `json:"records,omitempty"`
	Meta    Meta     `json:"meta,omitempty"`
}

// Create/Update request/response types for object types and records
type CreateObjectTypeRequest struct {
	Name   string                 `json:"name"`
	Slug   string                 `json:"slug,omitempty"`
	Fields []SchemaField          `json:"fields,omitempty"`
	Labels *LabelPair             `json:"labels,omitempty"`
	Extras map[string]interface{} `json:"-"`
}

type CreateObjectTypeResponse struct {
	Schema Schema `json:"schema,omitempty"`
}

type CreateRecordRequest struct {
	ID         string                 `json:"id,omitempty"`
	ExternalID *string                `json:"externalId,omitempty"`
	Fields     map[string]interface{} `json:"-"`
	// Relationships maps relationship id -> []values (API expects arrays)
	Relationships map[string][]interface{} `json:"relationships,omitempty"`
}

// MarshalJSON converts the Fields map into the API-expected array of objects
// with `id` and `value` keys. Relationships are included as-is (arrays).
func (r *CreateRecordRequest) MarshalJSON() ([]byte, error) {
	type fieldItem struct {
		ID    string      `json:"id"`
		Value interface{} `json:"value"`
	}
	out := make(map[string]interface{})
	if r.ID != "" {
		out["id"] = r.ID
	}
	if r.ExternalID != nil {
		out["externalId"] = r.ExternalID
	}
	if r.Fields != nil {
		items := make([]fieldItem, 0, len(r.Fields))
		for k, v := range r.Fields {
			items = append(items, fieldItem{ID: k, Value: v})
		}
		out["fields"] = items
	}
	if len(r.Relationships) > 0 {
		out["relationships"] = r.Relationships
	}
	return json.Marshal(out)
}

type CreateRecordResponse struct {
	Record Record `json:"record,omitempty"`
}

type UpdateRecordRequest = CreateRecordRequest
type UpdateRecordResponse = CreateRecordResponse
