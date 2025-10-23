package genconstants

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/contacts"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/lists"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/tags"
	"github.com/stretchr/testify/assert"
)

func TestFormatWithGofmt_success(t *testing.T) {
	src := []byte("package x\nfunc F() {}\n")
	out, err := formatWithGofmt(src)
	assert.NoError(t, err)
	assert.Contains(t, string(out), "package x")
}

func TestFetchAll_with_real_services_pagination(t *testing.T) {
	// create slices of 5 items, use limit=2 to force multiple pages
	tagsAll := []tags.TagPayload{{ID: "1", Tag: "A"}, {ID: "2", Tag: "B"}, {ID: "3", Tag: "C"}, {ID: "4", Tag: "D"}, {ID: "5", Tag: "E"}}
	fieldsAll := []contacts.FieldPayload{{ID: "10", Title: "F1"}, {ID: "11", Title: "F2"}, {ID: "12", Title: "F3"}, {ID: "13", Title: "F4"}, {ID: "14", Title: "F5"}}
	listsAll := []lists.List{{ID: "100", Name: "L1"}, {ID: "101", Name: "L2"}, {ID: "102", Name: "L3"}, {ID: "103", Name: "L4"}, {ID: "104", Name: "L5"}}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		limit, _ := strconv.Atoi(q.Get("limit"))
		offset, _ := strconv.Atoi(q.Get("offset"))
		switch {
		case r.URL.Path == "/api/3/tags" || r.URL.Path == "/tags":
			end := offset + limit
			if end > len(tagsAll) {
				end = len(tagsAll)
			}
			page := tagsAll[offset:end]
			b, _ := json.Marshal(map[string][]tags.TagPayload{"tags": page})
			w.Write(b)
		case r.URL.Path == "/api/3/fields" || r.URL.Path == "/fields":
			end := offset + limit
			if end > len(fieldsAll) {
				end = len(fieldsAll)
			}
			page := fieldsAll[offset:end]
			b, _ := json.Marshal(map[string][]contacts.FieldPayload{"fields": page})
			w.Write(b)
		case r.URL.Path == "/api/3/lists" || r.URL.Path == "/lists":
			end := offset + limit
			if end > len(listsAll) {
				end = len(listsAll)
			}
			page := listsAll[offset:end]
			b, _ := json.Marshal(map[string][]lists.List{"lists": page})
			w.Write(b)
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	cc, err := client.NewCoreClient(srv.URL, "token")
	assert.NoError(t, err)

	ctx := context.Background()
	tgSvc := tags.NewRealService(cc)
	gotTags, err := fetchAllTags(ctx, tgSvc, 2)
	assert.NoError(t, err)
	assert.Len(t, gotTags.TagsOrEmpty(), 5)

	cfSvc := contacts.NewRealService(cc)
	gotFields, err := fetchAllFields(ctx, cfSvc, 2)
	assert.NoError(t, err)
	assert.Len(t, gotFields.FieldsOrEmpty(), 5)

	lSvc := lists.NewRealService(cc)
	gotLists, err := fetchAllLists(ctx, lSvc, 2)
	assert.NoError(t, err)
	assert.Len(t, gotLists.Lists, 5)
}

func TestLoadMapping_nil_mappings_are_normalized(t *testing.T) {
	f, err := os.CreateTemp("", "mapnil-*.json")
	assert.NoError(t, err)
	path := f.Name()
	// write mappings:null
	_, err = f.Write([]byte(`{"package":"p","out":"o","mappings":null}`))
	f.Close()
	defer os.Remove(path)

	mf, err := loadMapping(path)
	assert.NoError(t, err)
	assert.NotNil(t, mf)
	assert.NotNil(t, mf.Mappings)
}

func TestRenderConsts_with_existing_mapping(t *testing.T) {
	buf := &stringsBuilder{}
	kvs := []KV{{Key: "X", Value: "1"}}
	mapping := map[string]string{"Tag|1": "PreMapName"}
	mappingUpdated := map[string]string{}
	renderConsts(buf, "Tag", kvs, mapping, mappingUpdated)
	out := buf.String()
	assert.Contains(t, out, "PreMapName")
	// mappingUpdated should not include key since mapping provided
	_, ok := mappingUpdated["Tag|1"]
	assert.False(t, ok)
}

func TestSanitizeIdentifier_additional_cases(t *testing.T) {
	cases := map[string]string{
		"cpr number": "CPRNumber",
		"pfa test":   "PFATest",
		"efa val":    "EFAVal",
		"rto id":     "RTOID",
		"id value":   "IDValue",
		"123abc":     "_123abc",
	}
	for in, want := range cases {
		got := sanitizeIdentifier(in)
		assert.Equal(t, want, got)
	}
}

func TestNewGenerator_defaults_and_loadMapping_missing(t *testing.T) {
	g := NewGenerator("http://example.com", "tok")
	// defaults set
	assert.Equal(t, "active/constants.go", g.OutPath)
	assert.Equal(t, ".gen-constants.map.json", g.MapPath)
	assert.Equal(t, "active", g.PackageName)
	assert.Equal(t, 100, g.Limit)

	// loadMapping on missing file should return non-nil MappingFile and no error
	mf, err := loadMapping("/does/not/exist/hopefully.json")
	assert.NoError(t, err)
	assert.NotNil(t, mf)
	assert.NotNil(t, mf.Mappings)
}

func TestExercise_sanitize_and_render_paths(t *testing.T) {
	// various inputs to trigger many sanitizeIdentifier branches
	inputs := []string{
		"", "---", "First name", "50% off", "123numeric", "Cpr code", "pfa thing", "Efa val", "contact id", "dob date", "post-code", "rto id", "with@Id",
	}
	for _, in := range inputs {
		_ = sanitizeIdentifier(in)
	}

	// create a mix of KV entries that will cause collisions and suffixing
	kvs := []KV{
		{Key: "Example", Value: "1"},
		{Key: "example", Value: "2"},
		{Key: "Example!!", Value: "3"},
		{Key: "50% off", Value: "4"},
		{Key: "123", Value: "5"},
		{Key: "", Value: "6"},
	}
	buf := &stringsBuilder{}
	mapping := map[string]string{"Tag|6": "MappedEmpty"}
	mappingUpdated := map[string]string{}
	renderConsts(buf, "Tag", kvs, mapping, mappingUpdated)
	out := buf.String()
	// assertions to ensure output contains expected tokens
	assert.Contains(t, out, "TagsType")
	assert.Contains(t, out, "MappedEmpty")

	// saveMapping to a real temp file path
	mf := &MappingFile{Package: "p", Out: "o", Mappings: map[string]string{"a": "b"}}
	tf, err := os.CreateTemp("", "mapsave-*.json")
	assert.NoError(t, err)
	tf.Close()
	defer os.Remove(tf.Name())
	err = saveMapping(tf.Name(), mf)
	assert.NoError(t, err)
}

func TestLoadMapping_large_invalid_json(t *testing.T) {
	f, err := os.CreateTemp("", "badmap-*.json")
	assert.NoError(t, err)
	_, err = f.Write([]byte("not a json payload"))
	f.Close()
	defer os.Remove(f.Name())

	mf, err := loadMapping(f.Name())
	assert.Error(t, err)
	assert.Nil(t, mf)
}

func TestRenderConsts_many_collisions(t *testing.T) {
	// generate many KV entries that normalize to the same identifier
	var kvs []KV
	for i := 0; i < 20; i++ {
		kvs = append(kvs, KV{Key: "Example!", Value: strconv.Itoa(i)})
	}
	buf := &stringsBuilder{}
	mapping := map[string]string{}
	mappingUpdated := map[string]string{}
	renderConsts(buf, "Tag", kvs, mapping, mappingUpdated)
	// mappingUpdated should have entries for each
	assert.Equal(t, 20, len(mappingUpdated))
}

func TestGenerator_Generate_with_existing_mapping_and_small_limit(t *testing.T) {
	tagsAll := []tags.TagPayload{{ID: "1", Tag: "A"}, {ID: "2", Tag: "B"}}
	fieldsAll := []contacts.FieldPayload{{ID: "10", Title: "F1"}}
	listsAll := []lists.List{{ID: "100", Name: "L1"}}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		limit, _ := strconv.Atoi(q.Get("limit"))
		offset, _ := strconv.Atoi(q.Get("offset"))
		switch {
		case r.URL.Path == "/api/3/tags" || r.URL.Path == "/tags":
			end := offset + limit
			if end > len(tagsAll) {
				end = len(tagsAll)
			}
			page := tagsAll[offset:end]
			b, _ := json.Marshal(map[string][]tags.TagPayload{"tags": page})
			w.Write(b)
		case r.URL.Path == "/api/3/fields" || r.URL.Path == "/fields":
			end := offset + limit
			if end > len(fieldsAll) {
				end = len(fieldsAll)
			}
			page := fieldsAll[offset:end]
			b, _ := json.Marshal(map[string][]contacts.FieldPayload{"fields": page})
			w.Write(b)
		case r.URL.Path == "/api/3/lists" || r.URL.Path == "/lists":
			end := offset + limit
			if end > len(listsAll) {
				end = len(listsAll)
			}
			page := listsAll[offset:end]
			b, _ := json.Marshal(map[string][]lists.List{"lists": page})
			w.Write(b)
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	// prepare a mapping file with pre-existing mapping so mappingUpdated remains empty
	mf := &MappingFile{Package: "p", Out: "o", Mappings: map[string]string{"Tag|1": "AName", "Field|10": "FName"}}
	mfF, err := os.CreateTemp("", "mf-*.json")
	assert.NoError(t, err)
	b, _ := json.Marshal(mf)
	mfF.Write(b)
	mfF.Close()
	defer os.Remove(mfF.Name())

	outF, err := os.CreateTemp("", "out-*.go")
	assert.NoError(t, err)
	outF.Close()
	defer os.Remove(outF.Name())

	g := NewGenerator(srv.URL, "token")
	g.SetMapPath(mfF.Name())
	g.SetOutputPath(outF.Name())
	g.SetLimit(1)

	err = g.Generate()
	assert.NoError(t, err)
}
