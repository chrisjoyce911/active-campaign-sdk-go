package genconstants

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/contacts"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/lists"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/tags"
)

// Generator produces a Go source file of typed constants for ActiveCampaign
// resources (tags, fields, lists). Create with NewGenerator and call
// Configure methods before Generate.
type Generator struct {
	BaseURL     string
	Token       string
	OutPath     string
	MapPath     string
	PackageName string
	Limit       int
}

// KV is a simple key/value pair used during rendering.
type KV struct{ Key, Value string }

// NewGenerator returns a generator configured with sensible defaults.
func NewGenerator(baseURL, token string) *Generator {
	return &Generator{
		BaseURL:     baseURL,
		Token:       token,
		OutPath:     "active/constants.go",
		MapPath:     ".gen-constants.map.json",
		PackageName: "active",
		Limit:       100,
	}
}

// SetOutputPath overrides the output Go file path.
func (g *Generator) SetOutputPath(p string) { g.OutPath = p }

// SetMapPath overrides the mapping file path.
func (g *Generator) SetMapPath(p string) { g.MapPath = p }

// SetPackageName overrides package name used in generated file.
func (g *Generator) SetPackageName(n string) { g.PackageName = n }

// SetLimit sets the page size used when fetching remote lists.
func (g *Generator) SetLimit(l int) { g.Limit = l }

// MappingFile represents the on-disk mapping file structure
type MappingFile struct {
	Package  string            `json:"package,omitempty"`
	Out      string            `json:"out,omitempty"`
	Mappings map[string]string `json:"mappings,omitempty"`
}

// Generate fetches tags, fields and lists from the ActiveCampaign account and
// writes a formatted Go source file to the configured OutPath. It also
// persists a mapping file at MapPath to keep identifier names stable.
func (g *Generator) Generate() error {
	if g.BaseURL == "" || g.Token == "" {
		return fmt.Errorf("base url and token are required")
	}

	cc, err := client.NewCoreClient(g.BaseURL, g.Token)
	if err != nil {
		return fmt.Errorf("new core client: %w", err)
	}
	svc := contacts.NewRealService(cc)
	tagsSvc := tags.NewRealService(cc)
	listsSvc := lists.NewRealService(cc)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	tagsResp, err := fetchAllTags(ctx, tagsSvc, g.Limit)
	if err != nil {
		return fmt.Errorf("fetch tags: %w", err)
	}
	fieldsResp, err := fetchAllFields(ctx, svc, g.Limit)
	if err != nil {
		return fmt.Errorf("fetch fields: %w", err)
	}
	listsResp, err := fetchAllLists(ctx, listsSvc, g.Limit)
	if err != nil {
		return fmt.Errorf("fetch lists: %w", err)
	}

	var tkv, fkv, lkv []KV
	for _, tg := range tagsResp.TagsOrEmpty() {
		tkv = append(tkv, KV{Key: tg.Tag, Value: tg.ID})
	}
	for _, f := range fieldsResp.FieldsOrEmpty() {
		fkv = append(fkv, KV{Key: f.Title, Value: f.ID})
	}
	for _, ls := range listsResp.Lists {
		lkv = append(lkv, KV{Key: ls.Name, Value: ls.ID})
	}

	sort.Slice(tkv, func(i, j int) bool { return tkv[i].Key < tkv[j].Key })
	sort.Slice(fkv, func(i, j int) bool { return fkv[i].Key < fkv[j].Key })

	mf, _ := loadMapping(g.MapPath)
	if mf == nil {
		mf = &MappingFile{Package: g.PackageName, Out: g.OutPath, Mappings: map[string]string{}}
	}
	mappingUpdated := map[string]string{}

	buf := &bytes.Buffer{}
	renderHeader(buf, g.PackageName)
	renderConsts(buf, "Tag", tkv, mf.Mappings, mappingUpdated)
	renderConsts(buf, "Field", fkv, mf.Mappings, mappingUpdated)
	renderConsts(buf, "List", lkv, mf.Mappings, mappingUpdated)

	formatted, err := formatWithGofmt(buf.Bytes())
	if err != nil {
		return fmt.Errorf("format: %w", err)
	}

	// Write if different
	existing, _ := os.ReadFile(g.OutPath)
	if bytes.Equal(existing, formatted) {
		return nil
	}
	if err := os.WriteFile(g.OutPath, formatted, 0644); err != nil {
		return fmt.Errorf("write: %w", err)
	}

	if len(mappingUpdated) > 0 {
		for k, v := range mappingUpdated {
			mf.Mappings[k] = v
		}
		mf.Package = g.PackageName
		mf.Out = g.OutPath
		if err := saveMapping(g.MapPath, mf); err != nil {
			return fmt.Errorf("save mapping: %w", err)
		}
	}

	return nil
}

func loadMapping(path string) (*MappingFile, error) {
	mf := &MappingFile{Mappings: map[string]string{}}
	b, err := os.ReadFile(path)
	if err != nil {
		return mf, nil
	}
	if err := json.Unmarshal(b, mf); err != nil {
		return nil, err
	}
	if mf.Mappings == nil {
		mf.Mappings = map[string]string{}
	}
	return mf, nil
}

// jsonMarshalIndent is a wrapper around json.MarshalIndent so tests can
// override it to simulate marshal failures.
var jsonMarshalIndent = json.MarshalIndent

func saveMapping(path string, mf *MappingFile) error {
	b, err := jsonMarshalIndent(mf, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, 0644)
}

func fetchAllTags(ctx context.Context, svc tags.TagsService, limit int) (*tags.ListTagsResponse, error) {
	offset := 0
	var all []tags.TagPayload
	for {
		opts := map[string]string{"limit": fmt.Sprintf("%d", limit)}
		if offset > 0 {
			opts["offset"] = fmt.Sprintf("%d", offset)
		}
		resp, _, err := svc.ListTags(ctx, opts)
		if err != nil {
			return nil, err
		}
		page := resp.TagsOrEmpty()
		if len(page) == 0 {
			break
		}
		all = append(all, page...)
		if len(page) < limit {
			break
		}
		offset += limit
	}
	return &tags.ListTagsResponse{Tags: &all}, nil
}

func fetchAllFields(ctx context.Context, svc *contacts.RealService, limit int) (*contacts.ListFieldsResponse, error) {
	offset := 0
	var all []contacts.FieldPayload
	for {
		opts := map[string]string{"limit": fmt.Sprintf("%d", limit)}
		if offset > 0 {
			opts["offset"] = fmt.Sprintf("%d", offset)
		}
		resp, _, err := svc.ListCustomFieldsWithOpts(ctx, opts)
		if err != nil {
			return nil, err
		}
		page := resp.FieldsOrEmpty()
		if len(page) == 0 {
			break
		}
		all = append(all, page...)
		if len(page) < limit {
			break
		}
		offset += limit
	}
	return &contacts.ListFieldsResponse{Fields: &all}, nil
}

func fetchAllLists(ctx context.Context, svc lists.ListsService, limit int) (*lists.ListsResponse, error) {
	offset := 0
	var all []lists.List
	for {
		opts := map[string]string{"limit": fmt.Sprintf("%d", limit)}
		if offset > 0 {
			opts["offset"] = fmt.Sprintf("%d", offset)
		}
		resp, _, err := svc.ListLists(ctx, opts)
		if err != nil {
			return nil, err
		}
		if len(resp.Lists) == 0 {
			break
		}
		all = append(all, resp.Lists...)
		if len(resp.Lists) < limit {
			break
		}
		offset += limit
	}
	return &lists.ListsResponse{Lists: all}, nil
}

func renderHeader(w io.Writer, pkg string) {
	fmt.Fprintln(w, "package "+pkg)
	fmt.Fprintln(w, "")
	fmt.Fprintf(w, "// Code generated by genconstants; DO NOT EDIT. Generated at: %s\n", time.Now().UTC().Format(time.RFC3339))
	fmt.Fprintln(w, "")
}

func renderConsts(w io.Writer, prefix string, kvs []KV, mapping map[string]string, mappingUpdated map[string]string) {
	// Build a map of sanitized names to original keys and values, disambiguating collisions
	seen := map[string]string{}
	order := []string{}
	vals := map[string]string{}
	for _, kv := range kvs {
		keyID := prefix + "|" + kv.Value // mapping key uses ID for stability
		var baseName string
		if mapped, ok := mapping[keyID]; ok {
			baseName = mapped
		} else {
			baseName = sanitizeIdentifier(kv.Key)
		}
		name := baseName
		if orig, exists := seen[baseName]; exists {
			suffix := shortHash(orig + "|" + kv.Key)
			name = baseName + "_" + suffix
		}
		seen[baseName] = kv.Key
		order = append(order, name)
		vals[name] = kv.Value
		if _, ok := mapping[keyID]; !ok {
			mappingUpdated[keyID] = baseName
		}
	}

	sort.Strings(order)

	typeName := prefix + "sType"
	varName := prefix + "s"

	fmt.Fprintf(w, "type %s struct {\n", typeName)
	for _, name := range order {
		fmt.Fprintf(w, "    %s string\n", name)
	}
	fmt.Fprintln(w, "}")

	fmt.Fprintf(w, "var %s = %s{\n", varName, typeName)
	for _, name := range order {
		fmt.Fprintf(w, "    %s: \"%s\",\n", name, vals[name])
	}
	fmt.Fprintln(w, "}")
	fmt.Fprintln(w, "")

	revName := varName + "ByID"
	fmt.Fprintf(w, "var %s = map[string]string{\n", revName)
	for _, name := range order {
		fmt.Fprintf(w, "    \"%s\": \"%s\",\n", vals[name], name)
	}
	fmt.Fprintln(w, "}")
	fmt.Fprintln(w, "")
}

func shortHash(s string) string {
	sum := sha1.Sum([]byte(s))
	hex := fmt.Sprintf("%x", sum[:])
	return hex[:8]
}

func sanitizeIdentifier(s string) string {
	parts := regexp.MustCompile(`[^a-zA-Z0-9]+`).Split(strings.ToLower(s), -1)
	for i := range parts {
		if parts[i] == "" {
			continue
		}
		parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
	}
	k := strings.Join(parts, "")
	k = strings.ReplaceAll(k, "Cpr", "CPR")
	k = strings.ReplaceAll(k, "Pfa", "PFA")
	k = strings.ReplaceAll(k, "Efa", "EFA")
	k = strings.ReplaceAll(k, "ContactId", "ContactID")
	k = strings.ReplaceAll(k, "Dob", "DOB")
	k = strings.ReplaceAll(k, "Postcode", "PostCode")
	if strings.Contains(s, "%") {
		// find a purely numeric part and suffix it with Pct
		for i := range parts {
			if matched, _ := regexp.MatchString(`^\d+$`, parts[i]); matched {
				parts[i] = parts[i] + "Pct"
				break
			}
		}
		// recompute k from parts to include the injected Pct
		k = strings.Join(parts, "")
	}
	k = strings.ReplaceAll(k, "Rto", "RTO")
	k = strings.ReplaceAll(k, "Id", "ID")

	reValid := regexp.MustCompile(`[^a-zA-Z0-9_]`)
	k = reValid.ReplaceAllString(k, "")
	if k != "" && unicode.IsDigit(rune(k[0])) {
		k = "_" + k
	}
	if k == "" {
		return "_"
	}
	return k
}

func formatWithGofmt(src []byte) ([]byte, error) {
	cmd := exec.Command("gofmt")
	cmd.Stdin = bytes.NewReader(src)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("gofmt: %v output=%s", err, string(out))
	}
	return out, nil
}
