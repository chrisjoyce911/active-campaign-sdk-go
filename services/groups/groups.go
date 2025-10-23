package groups

import (
	"context"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

type service struct {
	client client.Doer
}

// GroupsService defines the operations available for the Groups API.
type GroupsService interface {
	ListGroups(ctx context.Context, opts map[string]string) (*ListGroupsResponse, *client.APIResponse, error)
	CreateGroup(ctx context.Context, req *CreateGroupRequest) (*CreateGroupResponse, *client.APIResponse, error)
	GetGroup(ctx context.Context, id string) (*GetGroupResponse, *client.APIResponse, error)
	UpdateGroup(ctx context.Context, id string, req *UpdateGroupRequest) (*UpdateGroupResponse, *client.APIResponse, error)
	DeleteGroup(ctx context.Context, id string) (*client.APIResponse, error)
}
