Examples

This directory contains small, focused example programs that demonstrate how to
use this SDK to call ActiveCampaign APIs. Examples are useful as runnable
reference code and quick-start snippets. Some examples are guarded by the
`examples` build tag — see notes below.

## Available examples

- `campaigns_list` — list campaigns and print their status
- `contact_rto_flow` — example showing contact RTO flow (may require credentials)
- `contacts_get_by_id` — get a contact by id
- `contacts_get_by_search_email` — search contacts by email
- `contacts_get_lists` — legacy adapter example for contact lists
- `contacts_get_tags` — list tags for a contact
- `contacts_mock` — tests/mocked example
- `contacts_real` — real API example (requires environment vars)
- `create_list` — create a list example
- `custom_client` — show how to provide a custom http client
- `custom_objects_demo` — demo for custom objects lifecycle
- `custom_objects_relationships` — custom objects relationships demo
- `genconstants` and `generated_constants` — generation utilities/examples
- `legacy_adapter` — adapter that shows legacy-style calls wired to typed services
- `list_all_custom_fields` — list custom fields example
- `search_by_email`, `search_contact`, `search_contact_by_email`, `search_contact_simple` — contact search examples
- `update_field_value` — update a field value example

## Running examples

Most examples follow the same construction pattern:

```go
core, err := client.NewCoreClient(os.Getenv("ACTIVE_URL"), os.Getenv("ACTIVE_TOKEN"))
if err != nil { /* handle */ }

svc := contacts.NewRealService(core) // or campaigns.NewRealService(core)
```

To build or run a single example from the repository root use:

```bash
# build
go build -tags=examples ./examples/campaigns_list

# run (example may require setting env vars)
ACTIVE_URL="https://youraccount.api-us1.com" ACTIVE_TOKEN="token" \
  go run -tags=examples ./examples/campaigns_list
```

## Notes

- Build tags: many example files include `//go:build examples` at the top. Use
  `-tags=examples` when building or running those packages.
- Environment variables: examples that call the real API require `ACTIVE_URL`
  and `ACTIVE_TOKEN` to be set. If you don't want to call the real API, run
  examples that use built-in mocks or set the env vars to placeholder values.
- Network behavior: some examples perform live HTTP calls during runtime and
  will hang or error if credentials or network access are not available. Run
  those examples manually with the correct env vars when needed.

## Contributing

If you add an example, please:

1. Put it in a directory under `examples/` using snake_case naming.
2. Add `//go:build examples` if the example requires environment variables or
   external network access.
3. Add a brief description to this README (a PR modifying this file is fine).

## Want an accounts example?

If you want, I can add a small runnable `examples/accounts_list` that demonstrates
list/create/get usage for the Accounts service (it will require `ACTIVE_URL` and
`ACTIVE_TOKEN` to call the real API). Tell me if you'd like a minimal no-network
placeholder instead.
TODO: add examples!
