# services/custom_objects

Custom Objects endpoints: object types (schemas) and their records. Below is
what our production examples and package tests exercise and are known to work
against ActiveCampaign's API (verified end-to-end during example runs):

Schemas
. https://developers.activecampaign.com/reference/custom-object-schemas

- GET List all schemas (list)
- POST Create Schema (create)
- GET Get Schema by ID (get)
- PUT Update Schema (update)
- DELETE Delete Schema (delete)

Parent and Child Schemas
. https://developers.activecampaign.com/reference/parent-and-child-schemas

- POST Create a public schema
- POST Create a child schema

Records (per-schema)
. https://developers.activecampaign.com/reference/custom-object-records

- GET List records for a Schema (list)
- POST Create a record (create)
- GET Get record by ID (get)
- POST Create-or-Update a record (upsert) — used by the SDK for updates
- DELETE Delete a record by ID (delete)
- GET Get record by external ID
- DELETE Delete record by external ID

Notes and SDK behaviors (important for integrators)

- Field encoding: ActiveCampaign's API returns and expects custom object
  record fields as an array of objects (e.g. [{"id":"name","value":"x"}])
  while Go callers often prefer map[string]interface{}. The SDK performs
  tolerant unmarshalling and marshals requests into the API-expected array
  format automatically. Consumers can work with `Record.Fields` as a map and
  the SDK will convert to/from the API shape.

- Update (mutate) behavior: the live API accepts updates via the POST
  "create-or-update" record endpoint for the payload shape the SDK uses.
  Attempts to send PUT/PATCH in earlier iterations failed with 400s. The
  SDK's `UpdateObjectRecord` implementation therefore uses the POST
  create-or-update flow and constructs the request body as
  {"record": {"id": "<id>", "schemaId": "<schema>", "fields": [...]}}
  which has been verified by production tests and the example demo.

- Relationships: when setting relationships (parent/child links) the SDK
  encodes them as arrays (map[string][]interface{}) inside the
  `record.relationships` object, matching the API's expected shape. Tests
  assert that relationships are serialized as arrays.

- Debugging: the SDK `CoreClient` exposes a `Debug` toggle and
  `SetDebugFilter` to dump outgoing request bodies for troubleshooting; see
  the root README for usage examples. By default, when `Debug` is enabled the
  client logs all outgoing requests unless a `DebugFilter` is provided.

Tests & Examples

- Examples: `examples/custom_objects_demo` exercises a full lifecycle
  (create schema, create record, get, update via POST upsert, delete record,
  delete schema). Run it with the environment variables set:

  ```bash
  CUSTOM_OBJECTS_SAFE=false ACTIVE_URL=<your_url> ACTIVE_TOKEN=<token> \
      go run -tags=examples ./examples/custom_objects_demo
  ```

- Package tests: see `services/custom_objects` unit tests which include:
  - tests verifying create/list/get/delete flows
  - tests asserting fields array marshaling/unmarshaling
  - tests asserting update uses POST create-or-update and relationships
    serialize as array
