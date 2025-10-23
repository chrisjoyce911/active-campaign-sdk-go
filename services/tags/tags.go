package tags

import (
	"context"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

type service struct {
	client client.Doer
}

type TagsService interface {
	ListTags(ctx context.Context, opts map[string]string) (*ListTagsResponse, *client.APIResponse, error)
	CreateTag(ctx context.Context, req *CreateOrUpdateTagRequest) (*TagResponse, *client.APIResponse, error)
	GetTag(ctx context.Context, id string) (*TagResponse, *client.APIResponse, error)
	UpdateTag(ctx context.Context, id string, req *CreateOrUpdateTagRequest) (*TagResponse, *client.APIResponse, error)
	DeleteTag(ctx context.Context, id string) (*client.APIResponse, error)
	AddTagToContact(ctx context.Context, contactID string, req *CreateOrUpdateTagRequest) (*TagResponse, *client.APIResponse, error)
}

// Accessor to avoid nil checks for tags list
func (l *ListTagsResponse) TagsOrEmpty() []TagPayload {
	if l == nil || l.Tags == nil {
		return []TagPayload{}
	}
	return *l.Tags
}
