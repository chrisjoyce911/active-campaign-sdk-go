package webhooks

import (
	"context"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

type service struct {
	client client.Doer
}

// WebhooksService defines available webhook-related methods.
type WebhooksService interface {
	CreateWebhook(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error)
	UpdateWebhook(ctx context.Context, id string, req interface{}) (interface{}, *client.APIResponse, error)
	GetWebhook(ctx context.Context, id string) (interface{}, *client.APIResponse, error)
	ListWebhooks(ctx context.Context, opts map[string]string) (interface{}, *client.APIResponse, error)
	ListWebhookEvents(ctx context.Context, opts map[string]string) (interface{}, *client.APIResponse, error)
	DeleteWebhook(ctx context.Context, id string) (*client.APIResponse, error)
}
