package genconstants

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	th "github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/contacts"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/lists"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/tags"
	"github.com/stretchr/testify/assert"
)

func TestGenerator_Generate_requires_credentials(t *testing.T) {
	g := NewGenerator("", "")
	err := g.Generate()
	assert.Error(t, err)
}

func TestLoadSaveMapping_roundtrip(t *testing.T) {
	f, err := os.CreateTemp("", "mapping-*.json")
	assert.NoError(t, err)
	path := f.Name()
	f.Close()
	defer os.Remove(path)

	mf := &MappingFile{Package: "x", Out: "o", Mappings: map[string]string{"Tag|1": "T"}}
	assert.NoError(t, saveMapping(path, mf))
	got, err := loadMapping(path)
	assert.NoError(t, err)
	assert.Equal(t, mf.Out, got.Out)
	assert.Equal(t, mf.Mappings["Tag|1"], got.Mappings["Tag|1"])
}

func TestFormatAndHeader(t *testing.T) {
	buf := &bytes.Buffer{}
	renderHeader(buf, "mypkg")
	out := buf.Bytes()
	formatted, err := formatWithGofmt(out)
	assert.NoError(t, err)
	assert.Contains(t, string(formatted), "package mypkg")
}

func TestFetchAllTagsFieldsLists_with_mock(t *testing.T) {
	// Mock tags response
	tagsBody := `{"tags": [{"id":"1","tag":"Foo"}]}`
	md := &th.MockDoer{Body: []byte(tagsBody), Resp: &client.APIResponse{}}
	// tags service via factory returns tags.TagsService, but fetchAllTags expects tags.TagsService
	tagsSvc := tags.NewRealServiceFromDoer(md)
	ctx := context.Background()
	_, err := fetchAllTags(ctx, tagsSvc, 10)
	assert.NoError(t, err)

	// Mock fields response via contacts service
	fieldsBody := `{"fields": [{"id":"10","title":"Bar"}]}`
	md2 := &th.MockDoer{Body: []byte(fieldsBody), Resp: &client.APIResponse{}}
	csvc := contacts.NewRealServiceFromDoer(md2)
	_, err = fetchAllFields(ctx, csvc, 10)
	assert.NoError(t, err)

	// Mock lists response
	listsBody := `{"lists": [{"id":"100","name":"Baz"}]}`
	md3 := &th.MockDoer{Body: []byte(listsBody), Resp: &client.APIResponse{}}
	lsvc := lists.NewRealServiceFromDoer(md3)
	_, err = fetchAllLists(ctx, lsvc, 10)
	assert.NoError(t, err)
}

func TestGenerator_Generate_happy_path(t *testing.T) {
	// Create an httptest server to simulate the ActiveCampaign API
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// r.URL.Path will include /api/3/... depending on CoreClient normalization
		switch {
		case strings.HasSuffix(r.URL.Path, "/tags"):
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"tags": [{"id":"1","tag":"Foo"}]}`))
		case strings.HasSuffix(r.URL.Path, "/fields"):
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"fields": [{"id":"10","title":"Bar"}]}`))
		case strings.HasSuffix(r.URL.Path, "/lists"):
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"lists": [{"id":"100","name":"Baz"}]}`))
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	// Temp output and mapping files
	outF, err := os.CreateTemp("", "gen-constants-*.go")
	assert.NoError(t, err)
	outPath := outF.Name()
	outF.Close()
	defer os.Remove(outPath)

	mapF, err := os.CreateTemp("", "gen-constants-*.json")
	assert.NoError(t, err)
	mapPath := mapF.Name()
	mapF.Close()
	defer os.Remove(mapPath)

	g := NewGenerator(srv.URL, "token")
	g.SetOutputPath(outPath)
	g.SetMapPath(mapPath)

	err = g.Generate()
	assert.NoError(t, err)

	// Check output file contains package and values
	b, err := os.ReadFile(outPath)
	assert.NoError(t, err)
	s := string(b)
	assert.Contains(t, s, "package active")
	assert.Contains(t, s, "Foo")
	assert.Contains(t, s, "Bar")
	assert.Contains(t, s, "Baz")

	// Check mapping file contains entries for Tag|1, Field|10, List|100
	mb, err := os.ReadFile(mapPath)
	assert.NoError(t, err)
	ms := string(mb)
	assert.Contains(t, ms, "Tag|1")
	assert.Contains(t, ms, "Field|10")
	assert.Contains(t, ms, "List|100")
}
