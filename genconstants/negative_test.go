package genconstants

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	th "github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/contacts"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/lists"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/tags"
	"github.com/stretchr/testify/assert"
)

func TestGenerator_Generate_api_error(t *testing.T) {
	// Server returns 500 for tags endpoint
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/3/tags" || r.URL.Path == "/tags" {
			http.Error(w, "internal", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.URL.Path == "/api/3/fields" || r.URL.Path == "/fields":
			w.Write([]byte(`{"fields": []}`))
		case r.URL.Path == "/api/3/lists" || r.URL.Path == "/lists":
			w.Write([]byte(`{"lists": []}`))
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	g := NewGenerator(srv.URL, "token")
	g.SetOutputPath("/dev/null/nonexistent.go")
	g.SetMapPath("/dev/null/nonexistent.json")
	err := g.Generate()
	assert.Error(t, err)
}

func TestLoadMapping_invalid_json(t *testing.T) {
	f, err := os.CreateTemp("", "bad-*.json")
	assert.NoError(t, err)
	path := f.Name()
	f.WriteString("not-json")
	f.Close()
	defer os.Remove(path)

	mf, err := loadMapping(path)
	// loadMapping returns error when JSON invalid
	assert.Error(t, err)
	assert.Nil(t, mf)
}

func TestSaveMapping_dir_error(t *testing.T) {
	dir, err := os.MkdirTemp("", "mapdir-*")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	mf := &MappingFile{Package: "p", Out: "o", Mappings: map[string]string{}}
	// Attempting to save mapping to a directory path should error
	err = saveMapping(dir, mf)
	assert.Error(t, err)
}

func TestFetchAllTags_service_error(t *testing.T) {
	md := &th.MockDoer{Err: errors.New("boom")}
	svc := tags.NewRealServiceFromDoer(md)
	_, err := fetchAllTags(context.Background(), svc, 10)
	assert.Error(t, err)
}

func TestFetchAllFields_service_error(t *testing.T) {
	md := &th.MockDoer{Err: errors.New("boom")}
	csvc := contacts.NewRealServiceFromDoer(md)
	_, err := fetchAllFields(context.Background(), csvc, 10)
	assert.Error(t, err)
}

func TestFetchAllLists_service_error(t *testing.T) {
	md := &th.MockDoer{Err: errors.New("boom")}
	lsvc := lists.NewRealServiceFromDoer(md)
	_, err := fetchAllLists(context.Background(), lsvc, 10)
	assert.Error(t, err)
}

func TestRenderHeader_and_renderConsts_edgecases(t *testing.T) {
	// renderHeader basic positive assertion
	buf := &bytes.Buffer{}
	renderHeader(buf, "pkgx")
	out := buf.String()
	assert.Contains(t, out, "package pkgx")

	// renderConsts with empty kvs should still produce type and var
	buf = &bytes.Buffer{}
	mapping := map[string]string{}
	mappingUpdated := map[string]string{}
	renderConsts(buf, "Tag", []KV{}, mapping, mappingUpdated)
	s := buf.String()
	assert.Contains(t, s, "type TagsType")
	assert.Contains(t, s, "var Tags = TagsType{")
}

func TestFormatWithGofmt_missing_gofmt(t *testing.T) {
	// Temporarily clear PATH so gofmt is not found
	old := os.Getenv("PATH")
	os.Setenv("PATH", "")
	defer os.Setenv("PATH", old)

	_, err := formatWithGofmt([]byte("package x; func() {"))
	assert.Error(t, err)
}
