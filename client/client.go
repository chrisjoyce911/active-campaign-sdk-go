package client

import (
	"context"
	"net/http"
)

// Client is the core interface for the SDK. It wraps HTTP behaviour and exposes
// service factories (Contacts, Accounts, Lists, Tags, ...).
//
// TODOs:
// - Define methods to return service interfaces (Contacts(), Accounts(), ...)
// - Implement concrete httpClient that satisfies this interface
// - Add options for BaseURL, Token, Timeout, Logger
// - Implement Do(ctx, req, v) which handles JSON marshalling/unmarshalling and returns APIResponse

type Client interface {
	Do(ctx context.Context, req *http.Request, v interface{}) (*APIResponse, error)
}

// ...existing code...
