# Legacy adapter example

This directory contains a small example demonstrating the legacy-style client construction and the `legacy` adapter helpers that delegate to the typed services.

Environment

Create a `.env` file in the repository root (or set environment variables directly):

```
ACTIVE_URL=https://your-account.api-us1.com
ACTIVE_TOKEN=your_api_token_here
```

Important note about legacy wiring

Historically this repository provided a lightweight compile-time stub for the root `active_campaign` client so old examples would compile during the migration. As of the current update `active_campaign.NewClient(opts)` will:

- If `opts.BaseUrl` is empty: return a lightweight stub (same as before).
- If `opts.BaseUrl` is provided: create a network-capable `client.CoreClient` and wire the root `Client.Contacts` API to the typed `contacts` service. That means calls like:

```go
client, _ := active_campaign.NewClient(&active_campaign.ClientOpts{BaseUrl: "https://...", Token: "..."})
contact, resp, err := client.Contacts.SearchEmail("me@example.com")
```

will perform a real API request when `BaseUrl`/`Token` are set.

This change preserves backwards compatibility for examples (empty BaseUrl still yields a stub) while allowing legacy call-sites to be migrated in-place to real network behavior if desired.

Run the example

```bash
# run the legacy adapter example (it requires ACTIVE_URL and ACTIVE_TOKEN in the environment)
go run ./examples/legacy_adapter
```

What the example shows

- The example demonstrates both the old legacy-style call (`client.Contacts.SearchEmail`) and the adapter call (`legacy.SearchContacts`). With the wiring above, both will perform real network calls when credentials are provided.

Debugging and safety

- To inspect outgoing JSON from the typed client, set `CLIENT_DEBUG=1` in the environment before running the example.
- If you want to keep examples strictly offline (no network), omit `ACTIVE_URL`/`ACTIVE_TOKEN` — `NewClient` will return a stub and legacy calls will return nils as before.

Extending the example

- The example is intentionally minimal. If you want a full end-to-end flow (search → create → update field → add to list → add tag) with a `--apply` dry-run flag, I can extend the example and add clearer output and error handling.
