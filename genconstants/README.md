# genconstants

Package `genconstants` provides a reusable generator for producing a Go
source file containing typed constants for ActiveCampaign resources such as
Tags, Custom Fields and Lists.

The generator is exposed as a programmatic package and there is an example
runner under `examples/gen-constants` that demonstrates how to call it from
the command line or CI.

Features

- Fetches Tags, Custom Fields and Lists from an ActiveCampaign account (paged).
- Deterministic sanitisation and collision handling with a stable mapping file
  (`.gen-constants.map.json`) to preserve identifier names across runs.
- Writes a formatted Go file (default `active/constants.go`) containing typed
  structs and reverse maps (e.g. `Tags`, `TagsByID`).

Usage (programmatically)

```go
g := genconstants.NewGenerator(baseURL, token)
g.SetOutputPath("active/constants.go")
g.SetMapPath(".gen-constants.map.json")
g.SetPackageName("active")
if err := g.Generate(); err != nil {
    // handle error
}
```

See `examples/gen-constants` for a runnable example that demonstrates flags
and usage. Generated output is idempotent — the file is only overwritten when
the formatted contents change.

Collision handling
Mapping file (`.gen-constants.map.json`)

The generator persists a mapping file (by default `.gen-constants.map.json`) that keeps a stable association
between resource IDs and the Go identifiers used in the generated code. This prevents identifiers from
changing when resource names are edited on the ActiveCampaign side.

Format

The mapping file is a simple JSON object with the following shape:

- `package` (string, optional): package name written in the generated file.
- `out` (string, optional): output path written in the generated file.
- `mappings` (object): map of mapping keys to sanitized identifier names.

Mapping keys are formed as `<Prefix>|<ID>` where `<Prefix>` is one of `Tag`, `Field`, or `List` and `<ID>` is the
resource ID returned by the API. Example mapping key: `Tag|12345`.

Example mapping file

```json
{
  "package": "active",
  "out": "active/constants.go",
  "mappings": {
    "Tag|123": "Awfa",
    "List|10": "Students",
    "Field|20": "RtoID"
  }
}
```

Workflow notes

- Dry-run: run `go run ./examples/gen-constants --dry-run` to print the generated file and see new mappings that would be
  added. This does not modify files.
- Apply: run `go run ./examples/gen-constants --apply --map-path=.gen-constants.map.json` to persist new mappings and write the
  generated file if it differs from the existing file.
- Review: check `.gen-constants.map.json` into version control so identifier decisions are auditable and stable across
  CI runs and developer machines.

Tips

- If you rename a resource on the ActiveCampaign side but want to keep the Go identifier, add or keep the mapping entry
  for the resource ID in the mapping file.
- If two resources collide on a sanitized name, the generator will append a deterministic SHA1-based suffix to the name;
  if you'd prefer a human-friendly name, edit the mapping file and re-run with `--apply`.

If two different tags or fields sanitize to the same Go identifier, the
generator will append a short deterministic suffix (first 8 chars of SHA1
hex) to make names unique. This is deterministic so repeated runs yield
stable names.

Notes

- The generator creates a `client.CoreClient` from `ACTIVE_URL` and
  `ACTIVE_TOKEN` and performs real API calls. Run with care.
- Pagination: the current implementation requests one page per resource
  using `--limit`. If you have more than `--limit` items, re-run with a
  higher `--limit` or adjust page size; the generator will page through
  results (offset + limit) and fetch all items automatically.

Example

After running the generator and writing `active/constants.go`, you'll get
grouped variables like:

```go
import "github.com/chrisjoyce911/active-campaign-sdk-go/active"

// use generated constants
fmt.Println("tag id:", active.Tags.Awfa)
fmt.Println("list id:", active.Lists.Students)
fmt.Println("field id:", active.Fields.RtoID)
```

These grouped vars provide nice dot-completion in IDEs and keep the
top-level package tidy.

Generated types

The generator emits a named struct type per resource and a package-level var
of that type. For example:

- `type TagsType struct { ... }` and `var Tags = TagsType{ ... }`
- `type FieldsType struct { ... }` and `var Fields = FieldsType{ ... }`
- `type ListsType struct { ... }` and `var Lists = ListsType{ ... }`

This gives you both a clear type and convenient dot-completion (e.g.
`active.Tags.SomeTag`).

Mapping behavior

- The generator will look for the mapping file at `--map-path` and prefer
  mappings found there (ID -> sanitized name). If an ID is not mapped, the
  generator will create a sanitized name and will add the new mapping to the
  file when you run with `--apply`.
- This ensures that renaming a tag/field/list on the ActiveCampaign side
  does not change the Go identifier once mapped.

Example mapping flow

1. Run a dry-run to see output and new mappings: `go run ./examples/gen-constants --dry-run`
2. Write mappings and generated file: `go run ./examples/gen-constants --apply --map-path=.gen-constants.map.json`
3. Inspect `.gen-constants.map.json` to review the assigned names.

Tests

There are unit and integration-style tests for the generator in the
`genconstants` package. The test suite uses `httptest.Server` and the
`internal/testhelpers` utilities to simulate API responses.

Example generated output (snippet)

Below is a small example showing the shape of the generated `active/constants.go` file.

```go
package active

// Code generated by genconstants; DO NOT EDIT.

type TagsType struct {
  Awfa string
  SomeOtherTag string
}

var Tags = TagsType{
  Awfa: "123",
  SomeOtherTag: "456",
}

var TagsByID = map[string]string{
  "123": "Awfa",
  "456": "SomeOtherTag",
}
```
