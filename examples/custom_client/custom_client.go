//go:build examples

package main

import (
	"context"
	"log"
	"net/http"
	"os"

	legacy "github.com/chrisjoyce911/active-campaign-sdk-go/legacy"
	"github.com/joho/godotenv"
)

// If you'd like, you can build your httpClient and avoid passing your token through this package entirely.
func main() {

	_ = godotenv.Load()

	client := http.DefaultClient
	rt := WithHeader(client.Transport)
	// load env first so token from .env is available
	_ = godotenv.Load()
	baseURL := os.Getenv("ACTIVE_URL")
	token := os.Getenv("ACTIVE_TOKEN")

	if token == "" {
		rt.Header.Set("Api-Token", "your-token-placeholder")
	} else {
		rt.Header.Set("Api-Token", token)
	}
	client.Transport = rt

	if baseURL == "" {
		log.Printf("ACTIVE_URL not set; running example in placeholder mode")
	} else {
		log.Printf("ACTIVE_URL set to %s", baseURL)
	}

	// Call the legacy adapter SearchContacts as an example placeholder.
	_, _, err := legacy.SearchContacts(context.Background(), "test@example.com")
	if err != nil {
		panic(err)
	}
}

// Credit: https://stackoverflow.com/a/51326483/4544386
type MyClient struct {
	http.Header
	rt http.RoundTripper
}

func WithHeader(rt http.RoundTripper) MyClient {
	if rt == nil {
		rt = http.DefaultTransport
	}

	return MyClient{Header: make(http.Header), rt: rt}
}

func (c MyClient) RoundTrip(req *http.Request) (*http.Response, error) {
	for k, v := range c.Header {
		req.Header[k] = v
	}

	return c.rt.RoundTrip(req)
}
