# Contact Bounce Logs Example

This example demonstrates how to retrieve bounce logs for a specific contact using the ActiveCampaign SDK.

## Features

- Uses typed response objects instead of `interface{}` for better type safety and IDE support
- Shows proper error handling and API response inspection
- Demonstrates the `GetContactBounceLogs` method with typed `BounceLogsResponse`

## Setup

1. Set environment variables:
   ```bash
   export ACTIVE_URL="https://youraccount.activehosted.com"
   export ACTIVE_TOKEN="your-api-token"
   export CONTACT_ID="123"  # Optional, defaults to "1"
   ```

2. Run the example:
   ```bash
   go run main.go
   ```

## Output

The example will output the bounce logs in JSON format, showing details like bounce timestamps, bounce IDs, and associated links.

## Benefits of Typed Responses

This example showcases the improved type safety provided by returning `BounceLogsResponse` instead of `interface{}`:

- Compile-time type checking
- Better IDE autocompletion and documentation
- Clearer code intent and structure
- Easier refactoring and maintenance