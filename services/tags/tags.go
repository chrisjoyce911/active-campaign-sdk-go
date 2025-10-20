package tags

import (
	"context"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// TagPayload represents the tag resource body used in requests and responses.
type TagPayload struct {
	ID  string `json:"id,omitempty"`
	Tag string `json:"tag,omitempty"`
}

// CreateOrUpdateTagRequest is the request envelope for creating or updating a tag.
type CreateOrUpdateTagRequest struct {
	Tag TagPayload `json:"tag"`
}

// TagResponse is the envelope returned for single-tag endpoints.
type TagResponse struct {
	Tag TagPayload `json:"tag"`
}

// ListTagsResponse is the envelope returned for list endpoints.
type ListTagsResponse struct {
	Tags *[]TagPayload `json:"tags"`
}

// Accessor to avoid nil checks for tags list
func (l *ListTagsResponse) TagsOrEmpty() []TagPayload {
	if l == nil || l.Tags == nil {
		return []TagPayload{}
	}
	return *l.Tags
}

type TagsService interface {
	ListTags(ctx context.Context, opts map[string]string) (*ListTagsResponse, *client.APIResponse, error)
	CreateTag(ctx context.Context, req *CreateOrUpdateTagRequest) (*TagResponse, *client.APIResponse, error)
	GetTag(ctx context.Context, id string) (*TagResponse, *client.APIResponse, error)
	UpdateTag(ctx context.Context, id string, req *CreateOrUpdateTagRequest) (*TagResponse, *client.APIResponse, error)
	DeleteTag(ctx context.Context, id string) (*client.APIResponse, error)
	AddTagToContact(ctx context.Context, contactID string, req *CreateOrUpdateTagRequest) (*TagResponse, *client.APIResponse, error)
}

type service struct {
	client client.Doer
}
