# contact_create_with_tags example

This example shows how to map a simple Go struct into an ActiveCampaign contact payload (including custom field values) and attach tags to the contact using the SDK helper `MapToContact` found in `services/contacts`.

The example is intentionally small and pragmatic. It demonstrates the recommended pattern for:

- mapping application structs into `contacts.Contact` + a list of tag IDs
- creating a contact via the `contacts` service
- attaching tags to the created contact

This example is build-tagged with `examples` so it won't be included in normal test runs or builds unless you explicitly build/run it with that tag.

## Files

- `main.go` - runnable example (guarded by `//go:build examples`). It uses `contacts.MapToContact` to convert a user struct to the SDK `Contact` payload and then calls `svc.Create` followed by `svc.CreateContactTag` to attach tags.

## Prerequisites

- Go >= 1.19 (matches the project `go.mod`).
- A working ActiveCampaign account with API URL and API key if you intend to call the real API.
- The example imports `github.com/joho/godotenv` so if you rely on a local `.env` file, make sure your module has that dependency (run `go mod tidy` in the repo root if needed).

## Environment variables

The example reads two environment variables:

- `ACTIVE_URL` – your ActiveCampaign base URL (for example `https://youraccount.api-us1.com`).
- `ACTIVE_TOKEN` – your ActiveCampaign API token.

If those env vars are not set the example logs a message and will run in a placeholder/dry-run mode; however the example still constructs the client and will attempt API calls. For safety in automated runs, either set up a local httptest server or ensure you have a throwaway account.

You can optionally place these variables in a `.env` file at the repository root and the example will attempt to load it (the example uses `godotenv.Load()`). Example `.env`:

```
ACTIVE_URL=https://youraccount.api-us1.com
ACTIVE_TOKEN=0123456789abcdef
```

## Mapping rules (what MapToContact expects)

The helper `contacts.MapToContact` maps fields from a user struct using reflection and a few simple struct tags. The function signature is:

```
MapToContact(src interface{}, fieldIDByName map[string]string, tagNameToID map[string]string) (contacts.Contact, []string, error)
```

Key points:

- `src` can be any struct type. MapToContact inspects struct tags to decide which fields are contact properties, which are custom field values, and which hold tag information.
- Supported struct tags used by the helper:
  - `contact:"<FieldName>[,omitempty]"` — maps the struct field to core contact properties (Email, FirstName, LastName, etc.). The optional `omitempty` flag will omit zero values.
  - `fieldValues:"<FieldName>[,omitempty]"` — maps the struct field to a custom field value. The helper requires a `fieldIDByName` map that maps your human-friendly field name to the numeric field id used by the API.
  - `tags:"<FieldName>[,omitempty]"` — treats the struct field as a comma-separated list of tag names or numeric IDs. Each token is resolved to a tag ID by first checking if it is numeric (treated as id) or by looking it up in `tagNameToID`.

Examples:

```go
type MyContact struct {
  Email string `contact:"Email"`
  First string `contact:"FirstName,omitempty"`
  Star  string `fieldValues:"Star"`
  Tags  string `tags:"Tags,omitempty"` // e.g. "VIP,101"
}

// fieldIDByName is a map built from generated constants or via listing fields from the API.
fieldIDByName := map[string]string{"Star": "21"}
tagNameToID := map[string]string{"VIP": "100"}

contact, tagIDs, err := contacts.MapToContact(src, fieldIDByName, tagNameToID)
```

The helper returns:

- a `contacts.Contact` value populated with core contact properties and `FieldValues` populated
- a slice of tag IDs (strings) that you can feed to `CreateContactTag`

## Running the example

Build the example binary:

```sh
go build -tags=examples -o bin/contact_create_with_tags ./examples/contact_create_with_tags
```

Or run it directly:

```sh
go run -tags=examples ./examples/contact_create_with_tags
```

If you want to run the example against a local test server instead of the real API, create an `httptest.Server` that implements the routes used by the example and point `ACTIVE_URL` at it.

## Expected output (example)

On success you'll see something like:

```
Created contact id=123 email=jdoe@example.com
Attached tag id=100 to contact 123
Attached tag id=101 to contact 123
```

If creating the contact fails, the example will log the HTTP status and response body to help debugging.

## Dependency and build notes

- The example depends on `github.com/joho/godotenv` only for optional `.env` support. If you prefer not to pull that dependency, remove or change the call to `godotenv.Load()` in `main.go`.
- If `go build` complains about missing modules, run `go mod tidy` from the repository root to synchronize dependencies.

## Troubleshooting

- If you see an error about missing or invalid `ACTIVE_TOKEN`/`ACTIVE_URL`, confirm the environment variables are set and valid.
- If tag names in the `Tags` field are not found in `tagNameToID`, `MapToContact` will skip those entries unless they are numeric IDs. Ensure your `tagNameToID` map is populated (you can list tags via the API).
- To avoid touching real accounts while experimenting, run a local `httptest.Server` or use a sandbox account.

## Next steps / suggestions

- Add a small unit test for `MapToContact` that verifies mapping rules for contact fields, field values, and tags.
- Add a convenience method on the contacts service (e.g. `CreateContactWithTags`) which wraps create + attach-tag calls for a single call site.
- Consider removing the runtime dependency on `godotenv` in `main.go` if you prefer zero external deps for examples.

If you'd like I can add a compile-only test that ensures this example builds (guarded by the `examples` tag), and/or implement `CreateContactWithTags` in the service layer and a corresponding unit test. Tell me which you'd prefer next.
