# active-campaign-sdk-go

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/chrisjoyce911/active-campaign-sdk-go?tab=overview)
[![codecov](https://codecov.io/gh/chrisjoyce911/active-campaign-sdk-go/graph/badge.svg?token=RM7LL6MFUO)](https://codecov.io/gh/chrisjoyce911/active-campaign-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/chrisjoyce911/active-campaign-sdk-go)](https://goreportcard.com/report/github.com/chrisjoyce911/active-campaign-sdk-go)

## Summary

active-campaign-sdk-go is a typed Go client for the ActiveCampaign V3 API.
It provides `CoreClient` primitives and per-resource typed service packages
under `services/` (many of which have unit tests and examples).

## Table of contents

- Quick start
- Services (links)
- Examples
- Usage patterns / Client basics
- Testing & coverage
- Integration tests
- Development & contributing
- Generator (genconstants)
- Troubleshooting / Debugging
- Roadmap / Status
- License

## Usage

```go
package main

import "github.com/benkrig/active-campaign-sdk-go"
```

Construct a new client, then use the services available within the client to access the Active Campaign API.

```go
package main

import (
    ac "github.com/benkrig/active-campaign-sdk-go"
    "os"
)

func main() {
    baseURL := os.Getenv("YOUR_BASE_URL_KEY")
    token := os.Getenv("YOUR_TOKEN_KEY")

    a, err := ac.NewClient(
        &ac.ClientOpts{
            BaseUrl: baseURL,
            Token: token,
        },
    )
    if err != nil {
        panic(err)
    }

    c := ac.CreateContactRequest{
        &ac.Contact{
            Email: "test@email.com",
            FirstName: "testf",
            LastName: "testl",
            Phone: "1234567890",
        },
    }

    contact, _, err := a.Contacts.Create(&c)
    if err != nil {
        panic(err)
    }
}
```

## Code structure

The code structure of this package was inspired by [google/go-github](https://github.com/google/go-github) and [andygrunwald/go-jira](https://github.com/andygrunwald/go-jira).

Everything is based around the Client. The Client contains various services for resources found in the Active Campaign API, like Contacts, or Automations. Each service implements actions for its respective resource(s).

## Contribution

PR's are always welcome! The SDK is still being heavily developed and is missing many entities.

It doesn't matter if you are not able to write code.
Creating issues or holding talks and helping other people use the SDK is contribution as well!
A few examples:

- Correct typos in the README / documentation
- Reporting bugs
- Implement a new feature or endpoint

If you are new to pull requests, checkout [Collaborating on projects using issues and pull requests / Creating a pull request](https://help.github.com/articles/creating-a-pull-request/).

## Code coverage

This repository includes CI that runs the test suite and produces a coverage
report. The GitHub Actions workflow (`.github/workflows/ci.yml`) will run
`go test ./... -coverprofile=coverage.out` and upload `coverage.out` as a
workflow artifact. If you configure the secret `CODECOV_TOKEN` the job will
also upload coverage to Codecov.

Run coverage locally:

```bash
# run tests with coverage and open an HTML report
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
open coverage.html
```

If you'd like Codecov badges added to the README, set up the repository on
Codecov and provide `CODECOV_TOKEN` as a repository secret (private repos).

## License

This project is released under the terms of the [MIT license](http://en.wikipedia.org/wiki/MIT_License).

## Services (links)

Each service is implemented under `services/<name>` and includes a README
with examples and endpoint details. Quick links:

- Accounts — [services/accounts/README.md](services/accounts/README.md)
- Campaigns — [services/campaigns/README.md](services/campaigns/README.md)
- Contact Automations — [services/contactautomation/README.md](services/contactautomation/README.md)
- Contacts — [services/contacts/README.md](services/contacts/README.md)
- Custom Objects — [services/custom_objects/README.md](services/custom_objects/README.md)
- Deals — [services/deals/README.md](services/deals/README.md)
- E-Commerce — [services/ecommerce/README.md](services/ecommerce/README.md)
- Groups — [services/groups/README.md](services/groups/README.md)
- Lists — [services/lists/README.md](services/lists/README.md)
- Messages — [services/messages/README.md](services/messages/README.md)
- Tags — [services/tags/README.md](services/tags/README.md)
- Tracking — [services/tracking/README.md](services/tracking/README.md)
- Users — [services/users/README.md](services/users/README.md)
- Webhooks — [services/webhooks/README.md](services/webhooks/README.md)

## Generator (genconstants)

This repository includes a small generator that fetches Tags, Custom Fields
and Lists from an ActiveCampaign account and emits a typed Go source file.
See `genconstants/README.md` for usage examples and details.

Example generated output (short)

```go
package active

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

## Examples

Runnable examples live under `examples/`. Many are small CLI programs that
demonstrate wiring a `CoreClient` and calling a typed service. The
`examples/campaigns_list` example includes a unit test and a build/run
integration test (gated by the `integration` build tag).

## Debugging outgoing requests

The SDK `CoreClient` includes a simple, opt-in debug facility to dump outgoing
request bodies for easier troubleshooting. This is helpful when developing
against the ActiveCampaign API and wanting to inspect the exact JSON the
client sends.

Usage examples:

- Print debug output to stdout:

```go
cc.SetDebug(true, os.Stdout)
```

- Capture debug output in tests (inspect later):

```go
var buf bytes.Buffer
cc.SetDebug(true, &buf)
// perform calls
// assert strings.Contains(buf.String(), "DEBUG OUTGOING")
```

If you pass a nil writer the SDK falls back to the standard logger
(`log.Printf`). Debug output is gated by a boolean flag so it must be
explicitly enabled. In the current SDK behavior, when Debug is enabled the
client will emit debug output for all outgoing requests by default. If you
prefer to only log a subset of requests, provide a filter using
`SetDebugFilter` (see example below).

The debug header looks like:

```
DEBUG OUTGOING <METHOD> <URL> body:
{...json body...}
```

If you prefer different formatting or structured logs (JSON), you can pass a
custom writer that formats the bytes any way you like.

Example: restrict debug to only POSTs to custom objects endpoints.

```go
cc.SetDebug(true, os.Stdout)
cc.SetDebugFilter(func(method, path string) bool {
    if method != "POST" {
        return false
    }
    return strings.Contains(path, "/customObjects/")
})
```

[![Codecov sunburst](https://codecov.io/gh/chrisjoyce911/active-campaign-sdk-go/graphs/sunburst.svg?token=RM7LL6MFUO)](https://codecov.io/gh/chrisjoyce911/active-campaign-sdk-go)
