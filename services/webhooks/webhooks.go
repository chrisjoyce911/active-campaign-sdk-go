package webhooks

import "github.com/chrisjoyce911/active-campaign-sdk-go/client"

type service struct {
	client *client.CoreClient
}

type WebhooksService interface{}
