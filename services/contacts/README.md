## services/contacts

This package provides a thin, typed client for the ActiveCampaign Contacts API. It contains:

- The `ContactsService` interface describing the client-facing operations (create, update, search, field operations, tags, etc.).
- A minimal concrete implementation `RealService` that is backed by a `client.CoreClient` (or any `client.Doer` for tests).
- Models for contact payloads and related resources in `models.go`.
- Small helpers for mapping application structs into API payloads (see `mapper.go`).

This README documents the package purpose, some important helpers, testing guidance, and examples.

### Purpose and design

The package is organized to make the HTTP endpoints easy to use by callers while remaining testable:

- Each endpoint is implemented as a separate small source file (for clearer diffs and easier testing).
- The public `ContactsService` interface declares method signatures returning typed results along with a low-level `*client.APIResponse` and `error` (so callers can inspect HTTP-level details when needed).
- `RealService` is a lightweight adapter that delegates to a `client.Doer` (usually `*client.CoreClient`). Tests use `NewRealServiceFromDoer` to inject test doubles.

### Key helpers

- `MapToContact(src, fieldIDByName, tagNameToID)` (in `mapper.go`)

  - Reflectively maps a user struct into a `contacts.Contact` and a list of tag IDs.
  - Uses struct tags on the source type:
    - `contact:"<ContactField>[,omitempty]"` — map to core contact properties such as `Email`, `FirstName`.
    - `fieldValues:"<CustomFieldName>[,omitempty]"` — map to custom field values; requires `fieldIDByName` (map of field title -> id).
    - `tags:"<FieldName>[,omitempty]"` — expects a comma-separated list of tag names or IDs; resolves by name using `tagNameToID` if available.

- `BuildFieldIDByName(f)` (in `mapper.go`)

  - Convenience to convert generated `Fields` structs (from `genconstants`) into a map[name]->id.

- `CreateContactWithTags(ctx, req, tagIDs)` (in `create_with_tags.go`)
  - Convenience wrapper that creates a contact and then attaches the provided tag IDs via `CreateContactTag`.
  - Returns the created contact response and, if any attach fails, the last attach `*client.APIResponse` and error.

### List membership management

ActiveCampaign encodes list membership status as:

- Subscribed: `"1"`
- Unsubscribed: `"2"`

This package provides helpers to manage a contact’s list memberships safely:

- `UpdateListStatusManaged(ctx, req *UpdateListStatusHelperRequest)`
  - Reads current memberships via `GetContactLists` and applies the following logic:
    - If the contact is already Unsubscribed on the target list and `Force == false`, no change is made.
    - If `Force == true`, the contact will be set to Subscribed (`status: "1"`) regardless of prior Unsubscribed state.
    - If there is no existing membership, a new Subscribed membership is created.
  - Example:

```go
out, resp, err := contactsSvc.UpdateListStatusManaged(ctx, &contacts.UpdateListStatusHelperRequest{
    ContactList: &contacts.ContactList{Contact: contactID, List: listID, Status: "1"},
    Force:       false, // set to true to force subscription even if previously unsubscribed
})
```

- `EnsureSubscribedToList(ctx context.Context, contactID, listID string, force bool)`
  - Convenience wrapper that constructs the helper request (defaults to `status: "1"`).
  - Example:

```go
out, resp, err := contactsSvc.EnsureSubscribedToList(ctx, contactID, listID, false)
```

Usage scenarios and behavior

```go
// 1) No existing membership -> subscribes (POST /contactLists, status: "1")
_, resp, err := contactsSvc.EnsureSubscribedToList(ctx, "c1", "l1", false)

// 2) Already unsubscribed and force=false -> no-op (only GET); out == nil
_, resp, err = contactsSvc.EnsureSubscribedToList(ctx, "c2", "l2", false)

// 3) Already unsubscribed and force=true -> re-subscribe (POST)
_, resp, err = contactsSvc.EnsureSubscribedToList(ctx, "c2", "l2", true)

// 4) Already subscribed -> no-op (only GET); out == nil
_, resp, err = contactsSvc.EnsureSubscribedToList(ctx, "c3", "l3", true)
```

Notes

- When the helper decides no change is needed (no-op), it returns `out == nil` and the `resp` from the prior GET call so you can still inspect HTTP details.
- `GetContactLists` returns a typed `*ContactListsResponse`. Use `ContactListsOrEmpty()` to avoid nil checks when iterating:

```go
lists, _, _ := contactsSvc.GetContactLists(ctx, contactID)
for _, cl := range lists.ContactListsOrEmpty() {
    // cl.List, cl.Status, cl.UnsubReason, etc.
}
```

### Examples

- See `examples/contact_create_with_tags` for a runnable example that demonstrates `MapToContact`, creating a contact, and attaching tags. The example is build-tagged with `//go:build examples` so it won't be included in normal test runs.

### Testing

- The package includes many unit tests that use small test doubles:
  - `internal/testhelpers.MockDoer` and `internal/testhelpers.HTTPDoer` provide easy ways to stub `Do` responses and inspect outgoing requests.
  - `NewRealServiceFromDoer` allows injecting those doers into a `RealService` for tests.
- To run the package tests and see coverage for the mapper helpers:

```sh
go test ./services/contacts -v
go test ./services/contacts -coverprofile=services/contacts/coverage.out
go tool cover -func=services/contacts/coverage.out | grep mapper.go -A2
```

### Troubleshooting & notes

- `MapToContact` currently only inspects string-backed struct fields. Non-string fields are ignored.
- When mapping custom field values you must provide `fieldIDByName` (use `BuildFieldIDByName(Fields)` from generated constants or fetch fields via the API).
- Tag resolution prefers name lookup via `tagNameToID`; when a token is not found it is treated as a raw tag ID.
- `CreateContactWithTags` returns the created contact even when tag attachment fails; the last attach error and its `APIResponse` are returned to the caller so they can inspect what went wrong and decide how to recover.

### Contribution and future work

- Consider adding richer error aggregation for `CreateContactWithTags` (to return all attach failures, not only the last one).
- If you expose `MapToContact` to external users, add more comprehensive validation and support for more field types.
- Add `go:generate` to produce a mock for `ContactsService` to improve consumer testing.

If you'd like, I can also (pick one):

- add a compile-only test to ensure the examples build (guarded by `//go:build examples`),
- add more unit tests to push the entire package coverage to 100%, or
- change `CreateContactWithTags` to aggregate all attach errors and return them as a single combined error value.
