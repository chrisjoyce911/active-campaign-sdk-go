# contact_add_to_list example

This example demonstrates how to subscribe an existing contact to a list using the
`contacts.ContactsService.AddContactToList` method and the typed `AddContactToListPayload`.

## Prerequisites

Set the following environment variables (or pass the corresponding flags):

- `ACTIVE_URL` – Your ActiveCampaign API base URL.
- `ACTIVE_TOKEN` – The API token.
- `CONTACT_ID` – The contact to subscribe.
- `LIST_ID` – The list to subscribe the contact to.
- `LIST_STATUS` (optional) – Defaults to `1` (subscribe). Use `2` to unsubscribe.

## Running the example

By default the script runs in dry-run mode and only prints the request it would
send. Pass `-apply` to perform the mutation:

```bash
go run ./examples/contact_add_to_list -apply
```

You can override the contact, list, or status via flags:

```bash
go run ./examples/contact_add_to_list \
  -contact-id 123 \
  -list-id 5 \
  -status 1 \
  -apply
```

## Testing

The tests exercise the `run` helper directly without hitting the network:

```bash
go test ./examples/contact_add_to_list
```
