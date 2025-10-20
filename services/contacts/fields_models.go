package contacts

import "time"

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
	CreatedTimestamp *string           `json:"created_timestamp,omitempty"`
	UpdatedTimestamp *string           `json:"updated_timestamp,omitempty"`
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
