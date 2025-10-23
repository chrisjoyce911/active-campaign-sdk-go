package genconstants

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerator_Setters(t *testing.T) {
	g := NewGenerator("u", "t")
	g.SetOutputPath("out.go")
	g.SetMapPath("map.json")
	g.SetPackageName("pkg")
	g.SetLimit(5)

	assert.Equal(t, "out.go", g.OutPath)
	assert.Equal(t, "map.json", g.MapPath)
	assert.Equal(t, "pkg", g.PackageName)
	assert.Equal(t, 5, g.Limit)
}

func TestGenerator_Generate_idempotent(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/api/3/tags" || r.URL.Path == "/tags":
			w.Write([]byte(`{"tags":[{"id":"1","tag":"Foo"}]}`))
		case r.URL.Path == "/api/3/fields" || r.URL.Path == "/fields":
			w.Write([]byte(`{"fields":[{"id":"10","title":"Bar"}]}`))
		case r.URL.Path == "/api/3/lists" || r.URL.Path == "/lists":
			w.Write([]byte(`{"lists":[{"id":"100","name":"Baz"}]}`))
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	outF, err := os.CreateTemp("", "gen-*.go")
	assert.NoError(t, err)
	outPath := outF.Name()
	outF.Close()
	defer os.Remove(outPath)

	mapF, err := os.CreateTemp("", "map-*.json")
	assert.NoError(t, err)
	mapPath := mapF.Name()
	mapF.Close()
	defer os.Remove(mapPath)

	g := NewGenerator(srv.URL, "token")
	g.SetOutputPath(outPath)
	g.SetMapPath(mapPath)

	// First generate writes file
	err = g.Generate()
	assert.NoError(t, err)

	// Second generate should detect identical file and return nil early
	err = g.Generate()
	assert.NoError(t, err)
}

func TestGenerator_Generate_format_error(t *testing.T) {
	// clear PATH to simulate missing gofmt when formatWithGofmt is called
	old := os.Getenv("PATH")
	os.Setenv("PATH", "")
	defer os.Setenv("PATH", old)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/api/3/tags" || r.URL.Path == "/tags":
			w.Write([]byte(`{"tags":[]} `))
		case r.URL.Path == "/api/3/fields" || r.URL.Path == "/fields":
			w.Write([]byte(`{"fields":[]} `))
		case r.URL.Path == "/api/3/lists" || r.URL.Path == "/lists":
			w.Write([]byte(`{"lists":[]} `))
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	outF, err := os.CreateTemp("", "gen-*.go")
	assert.NoError(t, err)
	outPath := outF.Name()
	outF.Close()
	defer os.Remove(outPath)

	mapF, err := os.CreateTemp("", "map-*.json")
	assert.NoError(t, err)
	mapPath := mapF.Name()
	mapF.Close()
	defer os.Remove(mapPath)

	g := NewGenerator(srv.URL, "token")
	g.SetOutputPath(outPath)
	g.SetMapPath(mapPath)

	err = g.Generate()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "format:")
}

func TestGenerator_Generate_saveMapping_error(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/api/3/tags" || r.URL.Path == "/tags":
			w.Write([]byte(`{"tags":[{"id":"1","tag":"Foo"}]}`))
		case r.URL.Path == "/api/3/fields" || r.URL.Path == "/fields":
			w.Write([]byte(`{"fields":[{"id":"10","title":"Bar"}]}`))
		case r.URL.Path == "/api/3/lists" || r.URL.Path == "/lists":
			w.Write([]byte(`{"lists":[{"id":"100","name":"Baz"}]}`))
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	outF, err := os.CreateTemp("", "gen-*.go")
	assert.NoError(t, err)
	outPath := outF.Name()
	outF.Close()
	defer os.Remove(outPath)

	// make map path a directory so saveMapping will fail
	mapDir, err := os.MkdirTemp("", "mapdir-*")
	assert.NoError(t, err)
	defer os.RemoveAll(mapDir)

	g := NewGenerator(srv.URL, "token")
	g.SetOutputPath(outPath)
	g.SetMapPath(mapDir)

	err = g.Generate()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "save mapping")
}

func TestGenerator_Generate_saveMapping_success_and_format(t *testing.T) {
	// ensure PATH includes system gofmt (do not modify PATH)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/api/3/tags" || r.URL.Path == "/tags":
			w.Write([]byte(`{"tags":[{"id":"1","tag":"Foo"}]}`))
		case r.URL.Path == "/api/3/fields" || r.URL.Path == "/fields":
			w.Write([]byte(`{"fields":[{"id":"10","title":"Bar"}]}`))
		case r.URL.Path == "/api/3/lists" || r.URL.Path == "/lists":
			w.Write([]byte(`{"lists":[{"id":"100","name":"Baz"}]}`))
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	outF, err := os.CreateTemp("", "gen-*.go")
	assert.NoError(t, err)
	outPath := outF.Name()
	outF.Close()
	defer os.Remove(outPath)

	mapF, err := os.CreateTemp("", "map-*.json")
	assert.NoError(t, err)
	mapPath := mapF.Name()
	mapF.Close()
	defer os.Remove(mapPath)

	g := NewGenerator(srv.URL, "token")
	g.SetOutputPath(outPath)
	g.SetMapPath(mapPath)

	// Generate should create mapping and write files
	err = g.Generate()
	assert.NoError(t, err)

	// mapping file should now exist and contain mappings
	b, err := os.ReadFile(mapPath)
	assert.NoError(t, err)
	assert.Contains(t, string(b), "mappings")
}

func TestGenerator_Generate_newCoreClient_error(t *testing.T) {
	g := NewGenerator(":", "tok") // invalid URL will cause url.Parse error
	g.SetOutputPath("/dev/null/x.go")
	g.SetMapPath("/dev/null/x.json")
	err := g.Generate()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "new core client")
}

func TestGenerator_Generate_write_error_dir(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"tags":[],"fields":[],"lists":[]}`))
	}))
	defer srv.Close()

	dir, err := os.MkdirTemp("", "outdir-*")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	mapF, err := os.CreateTemp("", "map-*.json")
	assert.NoError(t, err)
	mapPath := mapF.Name()
	mapF.Close()
	defer os.Remove(mapPath)

	g := NewGenerator(srv.URL, "token")
	g.SetOutputPath(dir) // point to directory to cause write error
	g.SetMapPath(mapPath)

	err = g.Generate()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "write:")
}

func TestGenerator_Generate_saveMapping_marshal_error(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/api/3/tags" || r.URL.Path == "/tags":
			w.Write([]byte(`{"tags":[{"id":"1","tag":"Foo"}]}`))
		case r.URL.Path == "/api/3/fields" || r.URL.Path == "/fields":
			w.Write([]byte(`{"fields":[{"id":"10","title":"Bar"}]}`))
		case r.URL.Path == "/api/3/lists" || r.URL.Path == "/lists":
			w.Write([]byte(`{"lists":[{"id":"100","name":"Baz"}]}`))
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	outF, err := os.CreateTemp("", "gen-*.go")
	assert.NoError(t, err)
	outPath := outF.Name()
	outF.Close()
	defer os.Remove(outPath)

	mapF, err := os.CreateTemp("", "map-*.json")
	assert.NoError(t, err)
	mapPath := mapF.Name()
	mapF.Close()
	defer os.Remove(mapPath)

	// override jsonMarshalIndent to return an error
	old := jsonMarshalIndent
	jsonMarshalIndent = func(v interface{}, prefix, indent string) ([]byte, error) {
		return nil, fmt.Errorf("marshal fail")
	}
	defer func() { jsonMarshalIndent = old }()

	g := NewGenerator(srv.URL, "token")
	g.SetOutputPath(outPath)
	g.SetMapPath(mapPath)

	err = g.Generate()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "save mapping")
}
