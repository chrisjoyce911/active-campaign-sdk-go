package tags

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
