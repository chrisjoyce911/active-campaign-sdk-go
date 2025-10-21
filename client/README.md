# client

Contains the core Client interface and implementation that handles HTTP transport, base URL, authentication, and request/response handling.

TODO:

- Define `Client` interface
- Implement `NewClient(opts...)` constructor
- Implement `APIResponse` and `APIError` types
- Support injecting custom `http.RoundTripper` for testing
- Add tests for `Do` and request building
