# services/lists

Lists service interface and implementation.

Lists

POST Create a list
GET Retrieve a list
DELETE Delete a list
GET Retrieve all lists
POST Create a list group permission

Example: create a list

There is a small example under `examples/create_list` which demonstrates creating a list using
the SDK. It expects the environment variables `ACTIVE_URL` and `ACTIVE_TOKEN` to be set.

By default the example is safe — set `LISTS_SAFE=false` to bypass any additional warnings and
allow the example to create and delete data without an interactive prompt. The example also
supports a `--delete` flag (default: off) that will delete the created list after the run. Note:
if `LISTS_SAFE=false` the example defaults `--delete=true`.

Usage (locally):

```bash
# set these for your account before running the example
export ACTIVE_URL="https://youraccount.api-us1.com/"
export ACTIVE_TOKEN="<your-api-token>"

# optional: use an existing contact as the list owner and subscriber
export ACTIVE_CONTACTID="5"      # prefer contact ID
# or
export ACTIVE_EMAIL="user@example.com"  # will search and use first match

# run the example (it will delete the created list if LISTS_SAFE=false or --delete passed)
go run ./examples/create_list
```

Notes

- The API expects list payloads wrapped under a top-level `list` key. Use `CreateListRequest`.
- The example will prefer `ACTIVE_CONTACTID` when setting the list owner (`List.User`). If only
  `ACTIVE_EMAIL` is provided the example searches by email and uses the first matching contact.
- The SDK provides typed models for the important list fields; additional fields are unmarshalled
  into the `List` struct where available.
