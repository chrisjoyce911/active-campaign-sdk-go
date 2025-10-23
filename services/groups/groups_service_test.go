package groups

import (
	"context"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestGroupsService_Methods(t *testing.T) {
	t.Run("nil service guards", func(t *testing.T) {
		var s *service
		_, _, err := s.CreateGroup(context.Background(), nil)
		assert.Error(t, err)
		_, _, err = s.GetGroup(context.Background(), "1")
		assert.Error(t, err)
		_, _, err = s.ListGroups(context.Background(), nil)
		assert.Error(t, err)
		_, _, err = s.UpdateGroup(context.Background(), "1", nil)
		assert.Error(t, err)
		_, err = s.DeleteGroup(context.Background(), "1")
		assert.Error(t, err)
	})

	t.Run("happy paths", func(t *testing.T) {
		// Create
		md := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 201}, Body: []byte(`{"group":{"id":"g1","title":"T"}}`)}
		svc := NewRealServiceFromDoer(md).(*service)

		outC, apiResp, err := svc.CreateGroup(context.Background(), &CreateGroupRequest{Group: Group{Title: "T"}})
		assert.NoError(t, err)
		assert.Equal(t, 201, apiResp.StatusCode)
		assert.Equal(t, "g1", outC.Group.ID)

		// Get
		md = &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"group":{"id":"g1","title":"T"}}`)}
		svc = NewRealServiceFromDoer(md).(*service)
		outG, apiResp, err := svc.GetGroup(context.Background(), "g1")
		assert.NoError(t, err)
		assert.Equal(t, 200, apiResp.StatusCode)
		assert.Equal(t, "g1", outG.Group.ID)

		// List
		md = &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"groups":[{"id":"g1","title":"T"}]}`)}
		svc = NewRealServiceFromDoer(md).(*service)
		outL, apiResp, err := svc.ListGroups(context.Background(), map[string]string{"q": "x"})
		assert.NoError(t, err)
		assert.Equal(t, 200, apiResp.StatusCode)
		if outL != nil && len(outL.Groups) > 0 {
			assert.Equal(t, "g1", outL.Groups[0].ID)
		}

		// Update
		md = &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"group":{"id":"g1","title":"Updated"}}`)}
		svc = NewRealServiceFromDoer(md).(*service)
		outU, apiResp, err := svc.UpdateGroup(context.Background(), "g1", &UpdateGroupRequest{Group: Group{Title: "Updated"}})
		assert.NoError(t, err)
		assert.Equal(t, 200, apiResp.StatusCode)
		assert.Equal(t, "g1", outU.Group.ID)

		// Delete
		md = &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{}`)}
		svc = NewRealServiceFromDoer(md).(*service)
		apiResp, err = svc.DeleteGroup(context.Background(), "g1")
		assert.NoError(t, err)
		assert.Equal(t, 200, apiResp.StatusCode)
	})
}

func TestGroups_Constructors(t *testing.T) {
	// NewRealService should accept a CoreClient pointer; exercise NewRealServiceFromDoer
	md := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{}`)}
	svc := NewRealServiceFromDoer(md)
	// NewRealServiceFromDoer returns a GroupsService already; just ensure value is not nil
	assert.NotNil(t, svc)
}

func TestNewRealService(t *testing.T) {
	// create a minimal CoreClient and ensure NewRealService returns a GroupsService
	cc, err := client.NewCoreClient("https://example.com/", "")
	assert.NoError(t, err)
	svc := NewRealService(cc)
	// type check: should implement GroupsService
	var _ GroupsService = svc
}
