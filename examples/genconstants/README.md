# gen-constants

this is an example of how to use the gen-constants package tool to generate Go constants for ActiveCampaign resource IDs.

It enables you to run the generator against your ActiveCampaign account and produce a Go source file with typed constants for tags, custom fields, and lists.

This can be added to your project and run periodically to keep your constants up to date or to your CI/CD pipeline to ensure consistency.

This example shows how to use the `genconstants` package from inside the
repository to generate a typed Go source file for ActiveCampaign resources.

Run the example (after setting ACTIVE_URL and ACTIVE_TOKEN in your environment):

```bash
go run ./examples/gen-constants
```

It will write `active/constants.go` and update the mapping file `.gen-constants.map.json`.
\_ = godotenv.Load()
