package custom_objects

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestNilClientMethodsReturnNotImplemented(t *testing.T) {
	svc := &service{}

	// methods that explicitly check s.client == nil
	_, _, err := svc.GetObjectRecord(context.Background(), "ot", "r")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")

	_, err = svc.DeleteObjectRecord(context.Background(), "ot", "r")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")

	_, err = svc.DeleteObjectType(context.Background(), "ot")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")

	_, _, err = svc.ListObjectRecords(context.Background(), "ot", nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")

	_, _, err = svc.ListObjectTypes(context.Background(), nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")

	_, _, err = svc.UpdateObjectRecord(context.Background(), "ot", "r", &UpdateRecordRequest{Fields: map[string]interface{}{"x": "y"}})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")

}

func TestListSchemasResponse_UnmarshalLegacyAndMeta(t *testing.T) {
	raw := `{"objectTypes":[{"id":"ot-legacy"}],"meta":{"total":5,"count":1}}`
	var resp ListSchemasResponse
	err := json.Unmarshal([]byte(raw), &resp)
	assert.NoError(t, err)
	if assert.Len(t, resp.Schemas, 1) {
		assert.Equal(t, "ot-legacy", resp.Schemas[0].ID)
	}
	assert.Equal(t, 5, resp.Meta.Total)
	assert.Equal(t, 1, resp.Meta.Count)
}

func TestRecord_UnmarshalArrayAndMapForms(t *testing.T) {
	// array form
	arr := `{"id":"r1","fields":[{"id":"name","value":"alice"},{"id":"age","value":30}]}`
	var r Record
	err := json.Unmarshal([]byte(arr), &r)
	assert.NoError(t, err)
	assert.Equal(t, "r1", r.ID)
	if assert.NotNil(t, r.Fields) {
		assert.Equal(t, "alice", r.Fields["name"])
		// JSON numbers decode as float64
		assert.Equal(t, float64(30), r.Fields["age"])
	}

	// map form
	mp := `{"id":"r2","fields":{"x":"y","n":10}}`
	var r2 Record
	err = json.Unmarshal([]byte(mp), &r2)
	assert.NoError(t, err)
	assert.Equal(t, "r2", r2.ID)
	if assert.NotNil(t, r2.Fields) {
		assert.Equal(t, "y", r2.Fields["x"])
		assert.Equal(t, float64(10), r2.Fields["n"])
	}

	// no fields
	none := `{"id":"r3"}`
	var r3 Record
	err = json.Unmarshal([]byte(none), &r3)
	assert.NoError(t, err)
	assert.Nil(t, r3.Fields)
}

func TestCreateRecordRequest_MarshalJSON_CoversAllBranches(t *testing.T) {
	ext := "ext-1"
	req := &CreateRecordRequest{
		ID:         "rid",
		ExternalID: &ext,
		Fields:     map[string]interface{}{"f1": "v1"},
		Relationships: map[string][]interface{}{
			"rel": {42},
		},
	}
	b, err := json.Marshal(req)
	assert.NoError(t, err)

	var out map[string]interface{}
	err = json.Unmarshal(b, &out)
	assert.NoError(t, err)
	// id and externalId should appear
	if assert.Contains(t, out, "id") {
		assert.Equal(t, "rid", out["id"])
	}
	if assert.Contains(t, out, "externalId") {
		// externalId was a pointer to string, unmarshalled as string
		assert.Equal(t, "ext-1", out["externalId"])
	}
	// fields should be an array of objects
	if assert.Contains(t, out, "fields") {
		fields := out["fields"].([]interface{})
		assert.Len(t, fields, 1)
	}
	if assert.Contains(t, out, "relationships") {
		rels := out["relationships"].(map[string]interface{})
		v := rels["rel"].([]interface{})
		assert.Len(t, v, 1)
	}
}

func TestTimestamp_UnmarshalAlternateLayouts(t *testing.T) {
	// RFC3339Nano
	var ts Timestamp
	err := json.Unmarshal([]byte(`"2020-01-02T15:04:05.123456789Z"`), &ts)
	assert.NoError(t, err)
	if assert.NotNil(t, ts.Time) {
		assert.Equal(t, 2020, ts.Time.Year())
	}

	// custom layout
	var ts2 Timestamp
	err = json.Unmarshal([]byte(`"2006-01-02 15:04:05"`), &ts2)
	assert.NoError(t, err)
	if assert.NotNil(t, ts2.Time) {
		assert.Equal(t, 2006, ts2.Time.Year())
	}
}

func TestTimestamp_EmptyAndInvalid(t *testing.T) {
	var ts Timestamp
	// empty string should result in nil Time
	err := json.Unmarshal([]byte(`""`), &ts)
	assert.NoError(t, err)
	assert.Nil(t, ts.Time)

	// invalid should return error
	var ts2 Timestamp
	err = json.Unmarshal([]byte(`"not-a-time"`), &ts2)
	assert.Error(t, err)
}

func TestRecord_Unmarshal_ErrorOnBadFields(t *testing.T) {
	// fields is a number (invalid), should cause an error from fallback map unmarshal
	bad := `{"id":"rX","fields":123}`
	var r Record
	err := json.Unmarshal([]byte(bad), &r)
	assert.Error(t, err)
}

func TestListSchemasResponse_SchemasKey(t *testing.T) {
	raw := `{"schemas":[{"id":"ot-new"}],"meta":{"total":2}}`
	var resp ListSchemasResponse
	err := json.Unmarshal([]byte(raw), &resp)
	assert.NoError(t, err)
	if assert.Len(t, resp.Schemas, 1) {
		assert.Equal(t, "ot-new", resp.Schemas[0].ID)
	}
	assert.Equal(t, 2, resp.Meta.Total)
}

func TestUpdateObjectRecord_NilReqReturnsEmpty(t *testing.T) {
	// use a real doer to avoid nil-client check
	svc := NewRealServiceFromDoer(&testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"record":{}}`)})
	out, apiResp, err := svc.UpdateObjectRecord(context.Background(), "schema-1", "r1", nil)
	assert.NoError(t, err)
	assert.Nil(t, apiResp)
	assert.NotNil(t, out)
}

func TestNewRealService_ConstructsService(t *testing.T) {
	c, err := client.NewCoreClient("", "")
	if err != nil {
		t.Fatalf("failed to build core client: %v", err)
	}
	svc := NewRealService(c)
	// basic sanity: should not be nil and should implement CustomObjectsService
	if assert.NotNil(t, svc) {
		// avoid invoking networked methods here; just ensure the service was constructed
	}
}

func TestListSchemasResponse_Unmarshal_BadSchemasValue(t *testing.T) {
	// schemas is a string which should cause unmarshal into []Schema to fail
	raw := `{"schemas":"oops"}`
	var resp ListSchemasResponse
	err := json.Unmarshal([]byte(raw), &resp)
	assert.Error(t, err)
}

func TestGetObjectType_NilClientNonEmptyID(t *testing.T) {
	svc := &service{}
	_, _, err := svc.GetObjectType(context.Background(), "some-id")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")
}

func TestRecord_Unmarshal_IDNotString(t *testing.T) {
	raw := `{"id":123,"fields":[]}`
	var r Record
	err := json.Unmarshal([]byte(raw), &r)
	// ID unmarshal to string will fail but function should continue and not set ID
	assert.NoError(t, err)
	assert.Equal(t, "", r.ID)
}

func TestTimestamp_Unmarshal_NonStringJSON(t *testing.T) {
	var ts Timestamp
	// JSON number instead of string should cause json.Unmarshal into s to fail
	err := json.Unmarshal([]byte(`123`), &ts)
	assert.Error(t, err)
}

func TestListSchemasResponse_Unmarshal_BadMeta(t *testing.T) {
	raw := `{"schemas":[],"meta":"not-an-object"}`
	var resp ListSchemasResponse
	err := json.Unmarshal([]byte(raw), &resp)
	assert.Error(t, err)
}

func TestRecord_Unmarshal_BadFieldsType(t *testing.T) {
	raw := `{"id":"rX","fields":true}`
	var r Record
	err := json.Unmarshal([]byte(raw), &r)
	// fallback map unmarshal should fail and return error
	assert.Error(t, err)
}

func TestListSchemasResponse_EmptyDoesNotError(t *testing.T) {
	raw := `{}`
	var resp ListSchemasResponse
	err := json.Unmarshal([]byte(raw), &resp)
	assert.NoError(t, err)
	assert.Empty(t, resp.Schemas)
}

func TestSchema_FullUnmarshal(t *testing.T) {
	raw := `{
		"id":"otfull",
		"slug":"s1",
		"parentId":"p1",
		"parentVersion":2,
		"visibility":"v",
		"customized":true,
		"labels":{"singular":"S","plural":"P"},
		"description":"desc",
		"appId":"app1",
		"createdTimestamp":"2020-01-02T15:04:05Z",
		"updatedTimestamp":"2006-01-02 15:04:05",
		"parentSchemaCreatedTimestamp":"2020-01-02T15:04:05Z",
		"parentSchemaUpdatedTimestamp":"2020-01-02T15:04:05Z",
		"fields":[{"id":"f1","labels":{"singular":"a"},"description":"d","type":"text","required":true,"inherited":false,"scale":10,"defaultCurrency":"USD","options":[{"id":"o1","value":"v1"}]}],
		"icons": {"small":"i"},
		"relationships": [{"id":"rel1","labels":{"singular":"rs"},"description":"rd","namespace":"ns","hasMany":true,"inherited":false}]
	}`
	var s Schema
	err := json.Unmarshal([]byte(raw), &s)
	assert.NoError(t, err)
	assert.Equal(t, "otfull", s.ID)
	if assert.NotNil(t, s.CreatedTimestamp) {
		assert.Equal(t, 2020, s.CreatedTimestamp.Time.Year())
	}
	if assert.Len(t, s.Fields, 1) {
		assert.Equal(t, "f1", s.Fields[0].ID)
		if assert.Len(t, s.Fields[0].Options, 1) {
			assert.Equal(t, "o1", s.Fields[0].Options[0].ID)
		}
	}
	if assert.Len(t, s.Relationships, 1) {
		assert.Equal(t, "rel1", s.Relationships[0].ID)
	}
}

func TestTimestamp_CoverAllLayouts(t *testing.T) {
	// ensure each configured layout can be parsed by UnmarshalJSON
	base := time.Date(2021, 3, 14, 15, 9, 26, 0, time.UTC)
	for _, layout := range timestampLayouts {
		s := base.Format(layout)
		// wrap in JSON string
		raw, _ := json.Marshal(s)
		var ts Timestamp
		err := json.Unmarshal(raw, &ts)
		assert.NoError(t, err)
		if assert.NotNil(t, ts.Time) {
			assert.Equal(t, 2021, ts.Time.Year())
		}
	}
}

func TestListSchemasResponse_BothKeysPresent(t *testing.T) {
	raw := `{"schemas":[{"id":"a"}],"objectTypes":[{"id":"b"}],"meta":{"total":3}}`
	var resp ListSchemasResponse
	err := json.Unmarshal([]byte(raw), &resp)
	assert.NoError(t, err)
	// when both present, schemas key should be preferred
	if assert.Len(t, resp.Schemas, 1) {
		assert.Equal(t, "a", resp.Schemas[0].ID)
	}
	assert.Equal(t, 3, resp.Meta.Total)
}

func TestRecord_ArrayWithEmptyID(t *testing.T) {
	raw := `{"id":"rZ","fields":[{"id":"","value":"v"}]}`
	var r Record
	err := json.Unmarshal([]byte(raw), &r)
	assert.NoError(t, err)
	// empty id becomes empty key in map
	if assert.NotNil(t, r.Fields) {
		_, ok := r.Fields[""]
		assert.True(t, ok)
	}
}

func TestModels_Exhaustive(t *testing.T) {
	// Timestamp variants
	tsVariants := []string{
		`null`,
		`""`,
		`"2020-01-02T15:04:05.123Z"`,
		`"2020-01-02T15:04:05Z"`,
		`"2006-01-02 15:04:05"`,
		`123`,
		`"not-a-time"`,
	}
	for i, raw := range tsVariants {
		name := fmt.Sprintf("ts#%d", i)
		t.Run(name, func(t *testing.T) {
			var ts Timestamp
			_ = json.Unmarshal([]byte(raw), &ts)
		})
	}

	// ListSchemasResponse variants
	ls := []string{
		`{}`,
		`{"schemas":[]}`,
		`{"objectTypes":[]}`,
		`{"schemas":[],"objectTypes":[]}`,
		`{"schemas":[],"meta":{}}`,
		`{"schemas":0}`,
		`{"objectTypes":0}`,
		`not-json`,
	}
	for i, raw := range ls {
		name := fmt.Sprintf("ls#%d", i)
		t.Run(name, func(t *testing.T) {
			var out ListSchemasResponse
			_ = json.Unmarshal([]byte(raw), &out)
		})
	}

	// Record variants
	recs := []string{
		`{}`,
		`{"id":"x","fields":[]}`,
		`{"id":"x","fields":[{"id":"a","value":1}]}`,
		`{"id":"x","fields":{"a":1}}`,
		`{"id":null,"fields":[]}`,
		`{"id":123,"fields":{"a":1}}`,
		`{"id":"x","fields":true}`,
		`not-json`,
	}
	for i, raw := range recs {
		name := fmt.Sprintf("rec#%d", i)
		t.Run(name, func(t *testing.T) {
			var r Record
			_ = json.Unmarshal([]byte(raw), &r)
		})
	}
}

func TestRecord_FieldsNull(t *testing.T) {
	raw := `{"id":"rN","fields":null}`
	var r Record
	err := json.Unmarshal([]byte(raw), &r)
	assert.NoError(t, err)
	if r.Fields != nil {
		assert.Len(t, r.Fields, 0)
	}
}

func TestRecord_ArrayVariousValues(t *testing.T) {
	raw := `{"id":"rA","fields":[{"id":"a","value":null},{"id":"b","value":{"x":1}}]}`
	var r Record
	err := json.Unmarshal([]byte(raw), &r)
	assert.NoError(t, err)
	if assert.NotNil(t, r.Fields) {
		// null value should be nil
		_, ok := r.Fields["a"]
		assert.True(t, ok)
		// nested object preserved as map
		v := r.Fields["b"].(map[string]interface{})
		assert.Equal(t, float64(1), v["x"])
	}
}

func TestTimestamp_ManyFormats(t *testing.T) {
	samples := []string{
		"2020-01-02T15:04:05Z",
		"2020-01-02T15:04:05.000Z",
		"2006-01-02 15:04:05",
	}
	for _, s := range samples {
		var ts Timestamp
		raw, _ := json.Marshal(s)
		err := json.Unmarshal(raw, &ts)
		assert.NoError(t, err)
		assert.NotNil(t, ts.Time)
	}
}

func TestRecord_Unmarshal_InvalidJSON(t *testing.T) {
	var r Record
	err := json.Unmarshal([]byte(`not-json`), &r)
	assert.Error(t, err)
}

func TestTimestamp_NullExplicit(t *testing.T) {
	var ts Timestamp
	err := json.Unmarshal([]byte(`null`), &ts)
	assert.NoError(t, err)
	assert.Nil(t, ts.Time)
}

func TestTimestamp_FallbackParseExecuted(t *testing.T) {
	// Temporarily remove RFC3339 from layouts so fallback path is used
	orig := timestampLayouts
	timestampLayouts = []string{time.RFC3339Nano, "2006-01-02 15:04:05"}
	defer func() { timestampLayouts = orig }()

	// Provide an RFC3339 formatted time which should be parsed by the fallback
	var ts Timestamp
	err := json.Unmarshal([]byte(`"2020-01-02T15:04:05Z"`), &ts)
	assert.NoError(t, err)
	if assert.NotNil(t, ts.Time) {
		assert.Equal(t, 2020, ts.Time.Year())
	}
}

func TestListSchemasResponse_BadJSON(t *testing.T) {
	var resp ListSchemasResponse
	err := json.Unmarshal([]byte(`not-json`), &resp)
	assert.Error(t, err)
}

func TestListSchemasResponse_ObjectTypesBad(t *testing.T) {
	raw := `{"objectTypes":"oops"}`
	var resp ListSchemasResponse
	err := json.Unmarshal([]byte(raw), &resp)
	assert.Error(t, err)
}
