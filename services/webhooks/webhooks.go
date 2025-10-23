package webhooks

import (
	"context"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

type service struct {
	client client.Doer
}

// Webhook represents a webhook resource.
type Webhook struct {
	ID      string   `json:"id,omitempty"`
	Name    string   `json:"name,omitempty"`
	URL     string   `json:"url,omitempty"`
	Events  []string `json:"events,omitempty"`
	Sources []string `json:"sources,omitempty"`
	ListID  string   `json:"listid,omitempty"`
	Links   []string `json:"links,omitempty"`
	CDate   string   `json:"cdate,omitempty"`
}

// CreateWebhookRequest is the body for creating a webhook.
type CreateWebhookRequest struct {
	Webhook *Webhook `json:"webhook"`
}

// CreateWebhookResponse is the API response for CreateWebhook.
type CreateWebhookResponse struct {
	Webhook *Webhook `json:"webhook"`
}

// GetWebhookResponse is the API response for GetWebhook.
type GetWebhookResponse struct {
	Webhook *Webhook `json:"webhook"`
}

// ListWebhooksResponse is the API response for listing webhooks.
type ListWebhooksResponse struct {
	Webhooks []*Webhook        `json:"webhooks"`
	Meta     map[string]string `json:"meta,omitempty"`
}

// ListWebhookEventsResponse lists available events
type ListWebhookEventsResponse struct {
	Events []string `json:"events"`
}

// UpdateWebhookRequest wraps webhook update payload.
type UpdateWebhookRequest struct {
	Webhook *Webhook `json:"webhook"`
}

// UpdateWebhookResponse is the API response for UpdateWebhook.
type UpdateWebhookResponse struct {
	Webhook *Webhook `json:"webhook"`
}

// WebhooksService defines available webhook-related methods.
type WebhooksService interface {
	CreateWebhook(ctx context.Context, req *CreateWebhookRequest) (*CreateWebhookResponse, *client.APIResponse, error)
	UpdateWebhook(ctx context.Context, id string, req *UpdateWebhookRequest) (*UpdateWebhookResponse, *client.APIResponse, error)
	GetWebhook(ctx context.Context, id string) (*GetWebhookResponse, *client.APIResponse, error)
	ListWebhooks(ctx context.Context, opts map[string]string) (*ListWebhooksResponse, *client.APIResponse, error)
	ListWebhookEvents(ctx context.Context, opts map[string]string) (*ListWebhookEventsResponse, *client.APIResponse, error)
	DeleteWebhook(ctx context.Context, id string) (*client.APIResponse, error)
}
