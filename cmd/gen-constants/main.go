package main

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
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
	"github.com/joho/godotenv"
)

var (
	outPath = flag.String("out", "active/constants.go", "output file path")
	limit   = flag.Int("limit", 100, "page size to request from API")
	dryRun  = flag.Bool("dry-run", false, "render to stdout instead of writing file")
	apply   = flag.Bool("apply", false, "write the generated file if different")
	mapPath = flag.String("map-path", ".gen-constants.map.json", "path to mapping file (JSON)")
)

type KV struct {
	Key   string
	Value string
}

// MappingFile represents the on-disk mapping file structure
type MappingFile struct {
	Package  string            `json:"package,omitempty"`
	Out      string            `json:"out,omitempty"`
	Mappings map[string]string `json:"mappings,omitempty"`
}

func main() {
	flag.Parse()
	_ = godotenv.Load()

	base := os.Getenv("ACTIVE_URL")
	token := os.Getenv("ACTIVE_TOKEN")
	if base == "" || token == "" {
		log.Fatalf("ACTIVE_URL and ACTIVE_TOKEN must be set")
	}

	cc, err := client.NewCoreClient(base, token)
	if err != nil {
		log.Fatalf("NewCoreClient: %v", err)
	}
	svc := contacts.NewRealService(cc)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Fetch tags and fields (simple pagination)
	tags, err := fetchAllTags(ctx, *limit)
	if err != nil {
		log.Fatalf("fetch tags: %v", err)
	}
	fields, err := fetchAllFields(ctx, svc, *limit)
	if err != nil {
		log.Fatalf("fetch fields: %v", err)
	}
	listsResp, err := fetchAllLists(ctx, *limit)
	if err != nil {
		log.Fatalf("fetch lists: %v", err)
	}

	// Convert to KV and sanitize
	var tkv, fkv []KV
	for _, tg := range tags.TagsOrEmpty() {
		tkv = append(tkv, KV{Key: tg.Tag, Value: tg.ID})
	}
	// ListCustomFieldsWithOpts returns ListFieldsResponse which has Fields
	for _, f := range fields.FieldsOrEmpty() {
		fkv = append(fkv, KV{Key: f.Title, Value: f.ID})
	}
	var lkv []KV
	for _, l := range listsResp.Lists {
		lkv = append(lkv, KV{Key: l.Name, Value: l.ID})
	}

	// Sort for deterministic output
	sort.Slice(tkv, func(i, j int) bool { return tkv[i].Key < tkv[j].Key })
	sort.Slice(fkv, func(i, j int) bool { return fkv[i].Key < fkv[j].Key })

	// Load mapping file (if present)
	mappingFile, _ := loadMapping(*mapPath)
	if mappingFile == nil {
		mappingFile = &MappingFile{Package: "active", Out: *outPath, Mappings: map[string]string{}}
	}
	mappingUpdated := map[string]string{}

	// Render template
	buf := &bytes.Buffer{}
	renderHeader(buf)
	renderConsts(buf, "Tag", tkv, mappingFile.Mappings, mappingUpdated)
	renderConsts(buf, "Field", fkv, mappingFile.Mappings, mappingUpdated)
	renderConsts(buf, "List", lkv, mappingFile.Mappings, mappingUpdated)

	// Format output using gofmt
	formatted, err := formatWithGofmt(buf.Bytes())
	if err != nil {
		log.Fatalf("format: %v", err)
	}

	if *dryRun {
		io.Copy(os.Stdout, bytes.NewReader(formatted))
		return
	}

	// Write only if different
	existing, _ := os.ReadFile(*outPath)
	if bytes.Equal(existing, formatted) {
		fmt.Println("no changes")
		return
	}
	if *apply {
		if err := os.WriteFile(*outPath, formatted, 0644); err != nil {
			log.Fatalf("write: %v", err)
		}
		// persist mapping updates
		if len(mappingUpdated) > 0 {
			for k, v := range mappingUpdated {
				mappingFile.Mappings[k] = v
			}
			// update metadata
			mappingFile.Package = "active"
			mappingFile.Out = *outPath
			if err := saveMapping(*mapPath, mappingFile); err != nil {
				log.Fatalf("save mapping: %v", err)
			}
		}
		fmt.Println("wrote", *outPath)
	} else {
		fmt.Println("changes detected; run with --apply to write")
	}
}

// mapping file is a JSON object mapping "<prefix>|<originalKey>" -> sanitizedName
func loadMapping(path string) (*MappingFile, error) {
	mf := &MappingFile{Mappings: map[string]string{}}
	b, err := os.ReadFile(path)
	if err != nil {
		return mf, nil // not fatal
	}
	if err := json.Unmarshal(b, mf); err != nil {
		return nil, err
	}
	if mf.Mappings == nil {
		mf.Mappings = map[string]string{}
	}
	return mf, nil
}

func saveMapping(path string, mf *MappingFile) error {
	b, err := json.MarshalIndent(mf, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, 0644)
}

func fetchAllTags(ctx context.Context, limit int) (*tags.ListTagsResponse, error) {
	base := os.Getenv("ACTIVE_URL")
	token := os.Getenv("ACTIVE_TOKEN")
	cc, err := client.NewCoreClient(base, token)
	if err != nil {
		return nil, err
	}
	tagsSvc := tags.NewRealService(cc)

	offset := 0
	var all []tags.TagPayload
	for {
		opts := map[string]string{"limit": fmt.Sprintf("%d", limit)}
		if offset > 0 {
			opts["offset"] = fmt.Sprintf("%d", offset)
		}
		resp, _, err := tagsSvc.ListTags(ctx, opts)
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

func fetchAllLists(ctx context.Context, limit int) (*lists.ListsResponse, error) {
	base := os.Getenv("ACTIVE_URL")
	token := os.Getenv("ACTIVE_TOKEN")
	cc, err := client.NewCoreClient(base, token)
	if err != nil {
		return nil, err
	}
	listsSvc := lists.NewRealService(cc)

	offset := 0
	var all []lists.List
	for {
		opts := map[string]string{"limit": fmt.Sprintf("%d", limit)}
		if offset > 0 {
			opts["offset"] = fmt.Sprintf("%d", offset)
		}
		resp, _, err := listsSvc.ListLists(ctx, opts)
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

func renderHeader(w io.Writer) {
	fmt.Fprintln(w, "package active")
	fmt.Fprintln(w, "")
	fmt.Fprintf(w, "// Code generated by cmd/gen-constants; DO NOT EDIT. Generated at: %s\n", time.Now().UTC().Format(time.RFC3339))
	fmt.Fprintln(w, "")
}

func renderConsts(w io.Writer, prefix string, kvs []KV, mapping map[string]string, mappingUpdated map[string]string) {
	// Build a map of sanitized names to original keys and values, disambiguating collisions
	seen := map[string]string{}
	order := []string{}
	vals := map[string]string{}
	for _, kv := range kvs {
		keyID := prefix + "|" + kv.Value // mapping key now uses ID for stability
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
		// record mapping if not present
		if _, ok := mapping[keyID]; !ok {
			mappingUpdated[keyID] = baseName
		}
	}

	// Sort fields for deterministic struct ordering
	sort.Strings(order)

	// Use a named type per resource so packages have a clear type.
	typeName := prefix + "sType" // e.g., TagsType, FieldsType, ListsType
	varName := prefix + "s"      // e.g., Tags, Fields, Lists

	// type declaration
	fmt.Fprintf(w, "type %s struct {\n", typeName)
	for _, name := range order {
		fmt.Fprintf(w, "    %s string\n", name)
	}
	fmt.Fprintln(w, "}")

	// var with composite literal
	fmt.Fprintf(w, "var %s = %s{\n", varName, typeName)
	for _, name := range order {
		fmt.Fprintf(w, "    %s: \"%s\",\n", name, vals[name])
	}
	fmt.Fprintln(w, "}")
	fmt.Fprintln(w, "")

	// Emit reverse map ID->name
	revName := varName + "ByID"
	fmt.Fprintf(w, "var %s = map[string]string{\n", revName)
	for _, name := range order {
		fmt.Fprintf(w, "    \"%s\": \"%s\",\n", vals[name], name)
	}
	fmt.Fprintln(w, "}")
	fmt.Fprintln(w, "")
}

// shortHash returns the first 8 hex chars of SHA1 of input
func shortHash(s string) string {
	sum := sha1.Sum([]byte(s))
	hex := fmt.Sprintf("%x", sum[:])
	if len(hex) < 8 {
		return hex
	}
	return hex[:8]
}

// simplistic sanitiser; unit tests will be added
// sanitizeIdentifier converts an arbitrary string into a safe Go identifier
// fragment (no spaces or punctuation). It applies title-casing and a set of
// common acronym fixes, strips invalid characters, and prefixes an underscore
// if the identifier would start with a digit.
func sanitizeIdentifier(s string) string {
	// simple Title-case: split on non-letter/digit, uppercase first letter of words
	parts := regexp.MustCompile(`[^a-zA-Z0-9]+`).Split(strings.ToLower(s), -1)
	for i := range parts {
		if parts[i] == "" {
			continue
		}
		parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
	}
	k := strings.Join(parts, "")
	// common acronym fixes
	k = strings.ReplaceAll(k, "Cpr", "CPR")
	k = strings.ReplaceAll(k, "Pfa", "PFA")
	k = strings.ReplaceAll(k, "Efa", "EFA")
	k = strings.ReplaceAll(k, "ContactId", "ContactID")
	k = strings.ReplaceAll(k, "Dob", "DOB")
	k = strings.ReplaceAll(k, "Postcode", "PostCode")
	// If original string contained percent sign, append Pct
	if strings.Contains(s, "%") {
		k = k + "Pct"
	}
	// common small words that should be uppercased
	k = strings.ReplaceAll(k, "Rto", "RTO")
	k = strings.ReplaceAll(k, "Id", "ID")

	// Remove all characters not allowed in Go identifiers (letters, digits, _)
	reValid := regexp.MustCompile(`[^a-zA-Z0-9_]`)
	k = reValid.ReplaceAllString(k, "")

	// If key starts with a digit, prefix with '_'
	if k != "" && unicode.IsDigit(rune(k[0])) {
		k = "_" + k
	}
	if k == "" {
		return "_" // fallback
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
