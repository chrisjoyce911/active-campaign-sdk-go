package contacts

import (
	"encoding/json"
	"time"
)

// Models for contact custom fields

type FieldPayload struct {
	ID               string            `json:"id,omitempty"`
	Title            string            `json:"title,omitempty"`
	Descript         *string           `json:"descript,omitempty"`
	Type             string            `json:"type,omitempty"`
	IsRequired       string            `json:"isrequired,omitempty"`
	Perstag          string            `json:"perstag,omitempty"`
	Defval           *string           `json:"defval,omitempty"`
	ShowInList       string            `json:"show_in_list,omitempty"`
	Rows             string            `json:"rows,omitempty"`
	Cols             string            `json:"cols,omitempty"`
	Visible          string            `json:"visible,omitempty"`
	Service          string            `json:"service,omitempty"`
	OrderNum         string            `json:"ordernum,omitempty"`
	CDate            *time.Time        `json:"cdate,omitempty"`
	UDate            *time.Time        `json:"udate,omitempty"`
	CreatedTimestamp *time.Time        `json:"created_timestamp,omitempty"`
	UpdatedTimestamp *time.Time        `json:"updated_timestamp,omitempty"`
	CreatedBy        *string           `json:"created_by,omitempty"`
	UpdatedBy        *string           `json:"updated_by,omitempty"`
	Options          []string          `json:"options,omitempty"`
	Relations        []string          `json:"relations,omitempty"`
	Links            map[string]string `json:"links,omitempty"`
}

type FieldResponse struct {
	Field FieldPayload `json:"field"`
}

type ListFieldsResponse struct {
	Fields *[]FieldPayload `json:"fields"`
}

// Accessor for ListFieldsResponse
func (l *ListFieldsResponse) FieldsOrEmpty() []FieldPayload {
	if l == nil || l.Fields == nil {
		return []FieldPayload{}
	}
	return *l.Fields
}

type FieldOptionPayload struct {
	ID        string            `json:"id,omitempty"`
	Field     string            `json:"field,omitempty"`
	OrderID   string            `json:"orderid,omitempty"`
	Value     string            `json:"value,omitempty"`
	Label     string            `json:"label,omitempty"`
	IsDefault string            `json:"isdefault,omitempty"`
	CDate     *time.Time        `json:"cdate,omitempty"`
	UDate     *time.Time        `json:"udate,omitempty"`
	Links     map[string]string `json:"links,omitempty"`
}

type FieldOptionResponse struct {
	FieldOption FieldOptionPayload `json:"fieldOption"`
}

type FieldValuePayload struct {
	ID      string `json:"id,omitempty"`
	Contact string `json:"contact,omitempty"`
	Field   string `json:"field,omitempty"`
	Value   string `json:"value,omitempty"`
}

type FieldValueResponse struct {
	FieldValue FieldValuePayload `json:"fieldValue"`
}

type ListFieldValuesResponse struct {
	FieldValues *[]FieldValuePayload `json:"fieldValues"`
}

// Accessor for ListFieldValuesResponse
func (l *ListFieldValuesResponse) FieldValuesOrEmpty() []FieldValuePayload {
	if l == nil || l.FieldValues == nil {
		return []FieldValuePayload{}
	}
	return *l.FieldValues
}

// FieldGroup models
type FieldGroupPayload struct {
	ID    string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

type FieldGroupResponse struct {
	FieldGroup FieldGroupPayload `json:"fieldGroup"`
}

// --- Time parsing helpers -------------------------------------------------

var timeLayouts = []string{
	time.RFC3339,
	"2006-01-02 15:04:05", // created_timestamp format observed in API
	"2006-01-02T15:04:05-07:00",
}

func parseTimePtr(s *string) (*time.Time, error) {
	if s == nil || *s == "" {
		return nil, nil
	}
	v := *s
	var parsed time.Time
	var err error
	for _, layout := range timeLayouts {
		parsed, err = time.Parse(layout, v)
		if err == nil {
			t := parsed
			return &t, nil
		}
	}
	return nil, err
}

// UnmarshalJSON for FieldPayload attempts to parse timestamp strings into time.Time pointers
func (f *FieldPayload) UnmarshalJSON(data []byte) error {
	// Create an alias to avoid recursion
	type Alias FieldPayload
	aux := &struct {
		CDate            *string `json:"cdate,omitempty"`
		UDate            *string `json:"udate,omitempty"`
		CreatedTimestamp *string `json:"created_timestamp,omitempty"`
		UpdatedTimestamp *string `json:"updated_timestamp,omitempty"`
		*Alias
	}{Alias: (*Alias)(f)}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// parse times
	if t, err := parseTimePtr(aux.CDate); err == nil {
		f.CDate = t
	}
	if t, err := parseTimePtr(aux.UDate); err == nil {
		f.UDate = t
	}
	if t, err := parseTimePtr(aux.CreatedTimestamp); err == nil {
		f.CreatedTimestamp = t
	}
	if t, err := parseTimePtr(aux.UpdatedTimestamp); err == nil {
		f.UpdatedTimestamp = t
	}

	return nil
}

// UnmarshalJSON for FieldOptionPayload attempts to parse timestamp strings into time.Time pointers
func (fo *FieldOptionPayload) UnmarshalJSON(data []byte) error {
	type Alias FieldOptionPayload
	aux := &struct {
		CDate *string `json:"cdate,omitempty"`
		UDate *string `json:"udate,omitempty"`
		*Alias
	}{Alias: (*Alias)(fo)}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if t, err := parseTimePtr(aux.CDate); err == nil {
		fo.CDate = t
	}
	if t, err := parseTimePtr(aux.UDate); err == nil {
		fo.UDate = t
	}
	return nil
}

// Safe accessors
func (f *FieldPayload) CDateOrZero() time.Time {
	if f == nil || f.CDate == nil {
		return time.Time{}
	}
	return *f.CDate
}

func (f *FieldPayload) UDateOrZero() time.Time {
	if f == nil || f.UDate == nil {
		return time.Time{}
	}
	return *f.UDate
}

func (fo *FieldOptionPayload) CDateOrZero() time.Time {
	if fo == nil || fo.CDate == nil {
		return time.Time{}
	}
	return *fo.CDate
}

func (fo *FieldOptionPayload) UDateOrZero() time.Time {
	if fo == nil || fo.UDate == nil {
		return time.Time{}
	}
	return *fo.UDate
}
