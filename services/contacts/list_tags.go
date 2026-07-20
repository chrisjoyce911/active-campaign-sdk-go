package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// Tag is a tag defined on the account, as returned by GET /tags. This is the
// tag definition itself — ContactTag is the association between a tag and a
// contact.
type Tag struct {
	ID          string            `json:"id,omitempty"`
	Tag         string            `json:"tag,omitempty"`
	TagType     string            `json:"tagType,omitempty"`
	Description string            `json:"description,omitempty"`
	CDate       string            `json:"cdate,omitempty"`
	Subscriber  string            `json:"subscriber_count,omitempty"`
	Links       map[string]string `json:"links,omitempty"`
}

// ListTagsResponse is the payload returned by GET /tags.
type ListTagsResponse struct {
	Tags *[]Tag `json:"tags"`
}

// TagsOrEmpty returns the tags, or an empty slice when none were returned.
func (l *ListTagsResponse) TagsOrEmpty() []Tag {
	if l == nil || l.Tags == nil {
		return []Tag{}
	}
	return *l.Tags
}

// ListTags lists the tags defined on the account.
//
// Note ActiveCampaign pages this endpoint (default 20), so accounts with more
// tags than that need ListTagsWithOpts to raise the limit.
//
// GET /tags
func (s *RealService) ListTags(ctx context.Context) (*ListTagsResponse, *client.APIResponse, error) {
	return s.ListTagsWithOpts(ctx, nil)
}

// ListTagsWithOpts lists tags with optional query parameters, e.g.
// opts["limit"] = "100" or opts["offset"] for paging.
func (s *RealService) ListTagsWithOpts(ctx context.Context, opts map[string]string) (*ListTagsResponse, *client.APIResponse, error) {
	out := &ListTagsResponse{}
	apiResp, err := s.client.Do(ctx, http.MethodGet, "tags", opts, out)
	return out, apiResp, err
}
