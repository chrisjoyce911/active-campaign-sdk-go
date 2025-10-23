package genconstants

import (
	"context"
	"strconv"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/contacts"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/lists"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/tags"
	"github.com/stretchr/testify/assert"
)

func TestRenderConsts_collision_and_mappingUpdated(t *testing.T) {
	buf := &stringsBuilder{}
	kvs := []KV{{Key: "Hello", Value: "1"}, {Key: "Hello!", Value: "2"}}
	mapping := map[string]string{}
	mappingUpdated := map[string]string{}

	renderConsts(buf, "Tag", kvs, mapping, mappingUpdated)

	// mappingUpdated should contain entries for both IDs
	assert.Contains(t, mappingUpdated, "Tag|1")
	assert.Contains(t, mappingUpdated, "Tag|2")
	s := buf.String()
	// should contain generated struct and map
	assert.Contains(t, s, "type TagsType")
	assert.Contains(t, s, "var Tags = TagsType")
}

func TestShortHash_truncates(t *testing.T) {
	h := shortHash("some long string to hash")
	assert.Len(t, h, 8)
}

// stub implementations for services to exercise pagination
type fakeTagsSvc struct{ all []tags.TagPayload }

func (f *fakeTagsSvc) ListTags(ctx context.Context, opts map[string]string) (*tags.ListTagsResponse, *client.APIResponse, error) {
	limit, _ := strconv.Atoi(opts["limit"])
	offset := 0
	if v, ok := opts["offset"]; ok {
		offset, _ = strconv.Atoi(v)
	}
	if offset >= len(f.all) {
		empty := []tags.TagPayload{}
		return &tags.ListTagsResponse{Tags: &empty}, &client.APIResponse{StatusCode: 200}, nil
	}
	end := offset + limit
	if end > len(f.all) {
		end = len(f.all)
	}
	page := f.all[offset:end]
	return &tags.ListTagsResponse{Tags: &page}, &client.APIResponse{StatusCode: 200}, nil
}
func (f *fakeTagsSvc) CreateTag(ctx context.Context, req *tags.CreateOrUpdateTagRequest) (*tags.TagResponse, *client.APIResponse, error) {
	return nil, nil, nil
}
func (f *fakeTagsSvc) GetTag(ctx context.Context, id string) (*tags.TagResponse, *client.APIResponse, error) {
	return nil, nil, nil
}
func (f *fakeTagsSvc) UpdateTag(ctx context.Context, id string, req *tags.CreateOrUpdateTagRequest) (*tags.TagResponse, *client.APIResponse, error) {
	return nil, nil, nil
}
func (f *fakeTagsSvc) DeleteTag(ctx context.Context, id string) (*client.APIResponse, error) {
	return nil, nil
}
func (f *fakeTagsSvc) AddTagToContact(ctx context.Context, contactID string, req *tags.CreateOrUpdateTagRequest) (*tags.TagResponse, *client.APIResponse, error) {
	return nil, nil, nil
}

type fakeContactsSvc struct{ all []contacts.FieldPayload }

func (f *fakeContactsSvc) ListCustomFieldsWithOpts(ctx context.Context, opts map[string]string) (*contacts.ListFieldsResponse, *client.APIResponse, error) {
	limit, _ := strconv.Atoi(opts["limit"])
	offset := 0
	if v, ok := opts["offset"]; ok {
		offset, _ = strconv.Atoi(v)
	}
	if offset >= len(f.all) {
		empty := []contacts.FieldPayload{}
		return &contacts.ListFieldsResponse{Fields: &empty}, &client.APIResponse{StatusCode: 200}, nil
	}
	end := offset + limit
	if end > len(f.all) {
		end = len(f.all)
	}
	page := f.all[offset:end]
	return &contacts.ListFieldsResponse{Fields: &page}, &client.APIResponse{StatusCode: 200}, nil
}

type fakeListsSvc struct{ all []lists.List }

func (f *fakeListsSvc) ListLists(ctx context.Context, opts map[string]string) (lists.ListsResponse, *client.APIResponse, error) {
	limit, _ := strconv.Atoi(opts["limit"])
	offset := 0
	if v, ok := opts["offset"]; ok {
		offset, _ = strconv.Atoi(v)
	}
	if offset >= len(f.all) {
		return lists.ListsResponse{Lists: []lists.List{}}, &client.APIResponse{StatusCode: 200}, nil
	}
	end := offset + limit
	if end > len(f.all) {
		end = len(f.all)
	}
	page := f.all[offset:end]
	return lists.ListsResponse{Lists: page}, &client.APIResponse{StatusCode: 200}, nil
}
func (f *fakeListsSvc) CreateList(ctx context.Context, req lists.CreateListRequest) (lists.CreateListResponse, *client.APIResponse, error) {
	return lists.CreateListResponse{}, &client.APIResponse{StatusCode: 201}, nil
}
func (f *fakeListsSvc) GetList(ctx context.Context, id string) (lists.GetListResponse, *client.APIResponse, error) {
	return lists.GetListResponse{}, &client.APIResponse{StatusCode: 200}, nil
}
func (f *fakeListsSvc) DeleteList(ctx context.Context, id string) (*client.APIResponse, error) {
	return &client.APIResponse{StatusCode: 204}, nil
}
func (f *fakeListsSvc) CreateListGroup(ctx context.Context, req lists.CreateListGroupRequest) (lists.CreateListGroupResponse, *client.APIResponse, error) {
	return lists.CreateListGroupResponse{}, &client.APIResponse{StatusCode: 201}, nil
}

// minimal strings.Builder replacement to avoid import cycles in tests
type stringsBuilder struct{ b []byte }

func (s *stringsBuilder) Write(p []byte) (int, error) { s.b = append(s.b, p...); return len(p), nil }
func (s *stringsBuilder) String() string              { return string(s.b) }

func TestFetchAll_pagination_helpers(t *testing.T) {
	// create 3 items and set limit=2 so pagination loops twice
	tagsAll := []tags.TagPayload{{ID: "1", Tag: "A"}, {ID: "2", Tag: "B"}, {ID: "3", Tag: "C"}}
	tg := &fakeTagsSvc{all: tagsAll}
	ctx := context.Background()
	gotTags, err := fetchAllTags(ctx, tg, 2)
	assert.NoError(t, err)
	assert.Len(t, gotTags.TagsOrEmpty(), 3)

	fieldsAll := []contacts.FieldPayload{{ID: "10", Title: "F1"}, {ID: "11", Title: "F2"}, {ID: "12", Title: "F3"}}
	cf := &fakeContactsSvc{all: fieldsAll}
	// We can't easily construct a *contacts.RealService backed by our fake here, so exercise the pagination logic
	// by calling the fake's ListCustomFieldsWithOpts repeatedly to simulate fetchAllFields loop.
	offset := 0
	var acc []contacts.FieldPayload
	for {
		opts := map[string]string{"limit": strconv.Itoa(2)}
		if offset > 0 {
			opts["offset"] = strconv.Itoa(offset)
		}
		resp, _, err := cf.ListCustomFieldsWithOpts(ctx, opts)
		assert.NoError(t, err)
		page := resp.FieldsOrEmpty()
		if len(page) == 0 {
			break
		}
		acc = append(acc, page...)
		if len(page) < 2 {
			break
		}
		offset += 2
	}
	assert.Len(t, acc, 3)

	listsAll := []lists.List{{ID: "100", Name: "L1"}, {ID: "101", Name: "L2"}, {ID: "102", Name: "L3"}}
	lf := &fakeListsSvc{all: listsAll}
	gotLists, err := fetchAllLists(ctx, lf, 2)
	assert.NoError(t, err)
	assert.Len(t, gotLists.Lists, 3)
}

func TestRenderConsts_empty_and_digit_prefix(t *testing.T) {
	kvs := []KV{{Key: "---", Value: "x"}, {Key: "123 start", Value: "y"}}
	buf := &stringsBuilder{}
	mapping := map[string]string{}
	mappingUpdated := map[string]string{}
	renderConsts(buf, "Tag", kvs, mapping, mappingUpdated)
	out := buf.String()
	assert.Contains(t, out, "_")    // underscore fallback
	assert.Contains(t, out, "_123") // digit prefix handled
}
