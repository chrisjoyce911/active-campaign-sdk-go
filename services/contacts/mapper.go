package contacts

import (
	"fmt"
	"reflect"
	"strings"
)

// BuildFieldIDByName builds a map[name]id for a generated Fields-like struct.
// It accepts any struct value where each exported field is a string containing
// the ID. Example: pass the generated `Fields` value from genconstants.
func BuildFieldIDByName(f interface{}) map[string]string {
	out := make(map[string]string)
	v := reflect.ValueOf(f)
	// accept pointer to struct
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return out
	}
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Name
		id := v.Field(i).String()
		if id != "" {
			out[name] = id
		}
	}
	return out
}

// parseTag parses a tag of the form "Name[,omitempty]" returning the name
// and whether the omitempty flag was present.
func parseTag(tag string) (name string, omitEmpty bool) {
	parts := strings.Split(tag, ",")
	name = strings.TrimSpace(parts[0])
	for _, p := range parts[1:] {
		if strings.TrimSpace(p) == "omitempty" {
			omitEmpty = true
		}
	}
	return
}

// MapToContact maps a user-defined struct into a contacts.Contact and a
// slice of tag IDs to attach. Mapping rules are controlled by struct tags on
// the source type:
//   - contact:"<ContactField>[,omitempty]"  -> sets a field on Contact (e.g. Email)
//   - fieldValues:"<CustomFieldName>[,omitempty]" -> appends a FieldValue
//   - tags:"<TagNameOrID>[,other]" -> comma-separated tags; each item is
//     resolved using tagNameToID (preferred) or treated as an ID if not found.
//
// fieldIDByName is a map of custom field name -> field id (from genconstants
// or other source). tagNameToID is a map of tag name -> tag id. Either map
// may be nil but then field or tag lookups will fail when referenced.
func MapToContact(src interface{}, fieldIDByName map[string]string, tagNameToID map[string]string) (Contact, []string, error) {
	var out Contact

	sv := reflect.ValueOf(src)
	if !sv.IsValid() {
		return out, nil, fmt.Errorf("src is not valid")
	}
	if sv.Kind() == reflect.Ptr {
		sv = sv.Elem()
	}
	st := sv.Type()

	dv := reflect.ValueOf(&out).Elem()
	dt := dv.Type()

	var fvals []FieldValue
	var tagIDs []string

	for i := 0; i < st.NumField(); i++ {
		sf := st.Field(i)
		fv := sv.Field(i)

		// only handle string-backed fields in the source struct for now
		if sf.Type.Kind() != reflect.String {
			continue
		}
		s := fv.String()

		// contact:"<ContactField>[,omitempty]"
		if tag := sf.Tag.Get("contact"); tag != "" {
			name, omitEmpty := parseTag(tag)
			if omitEmpty && s == "" {
				// skip
			} else {
				df, ok := dt.FieldByName(name)
				if !ok {
					return out, nil, fmt.Errorf("contact tag refers to unknown field on Contact: %q", name)
				}
				dvf := dv.FieldByName(df.Name)
				if !dvf.CanSet() || dvf.Kind() != reflect.String {
					return out, nil, fmt.Errorf("contact field %q not a settable string", df.Name)
				}
				dvf.SetString(s)
			}
		}

		// fieldValues:"<CustomFieldName>[,omitempty]"
		if tag := sf.Tag.Get("fieldValues"); tag != "" {
			name, omitEmpty := parseTag(tag)
			if omitEmpty && s == "" {
				continue
			}
			if fieldIDByName == nil {
				return out, nil, fmt.Errorf("no field ID mapping provided for custom field %q", name)
			}
			id, ok := fieldIDByName[name]
			if !ok || id == "" {
				return out, nil, fmt.Errorf("no custom field ID configured for %q", name)
			}
			fvals = append(fvals, FieldValue{Field: id, Value: s})
		}

		// tags:"tagA,tagB[,omitempty]"  -- comma-separated tag names or ids
		if tag := sf.Tag.Get("tags"); tag != "" {
			name, omitEmpty := parseTag(tag)
			if omitEmpty && s == "" {
				continue
			}
			// source field value expected to be comma-separated tag tokens
			tokens := strings.Split(s, ",")
			for _, tok := range tokens {
				tok = strings.TrimSpace(tok)
				if tok == "" {
					continue
				}
				// prefer lookup by name
				if tagNameToID != nil {
					if id, ok := tagNameToID[tok]; ok && id != "" {
						tagIDs = append(tagIDs, id)
						continue
					}
				}
				// otherwise treat token as ID
				tagIDs = append(tagIDs, tok)
			}
			_ = name // reserved for future use (e.g. specifying destination field)
		}
	}

	if len(fvals) > 0 {
		out.FieldValues = &fvals
	}
	return out, tagIDs, nil
}
