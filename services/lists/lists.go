package lists

import (
	"context"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// ListsService handles list related API calls.
type ListsService interface {
	// CreateList creates a new list. Takes a CreateListRequest and returns CreateListResponse.
	CreateList(ctx context.Context, req CreateListRequest) (CreateListResponse, *client.APIResponse, error)

	// ListLists lists lists with optional query params and returns typed ListsResponse.
	ListLists(ctx context.Context, opts map[string]string) (ListsResponse, *client.APIResponse, error)

	// GetList fetches a single list by ID and returns a typed GetListResponse.
	GetList(ctx context.Context, id string) (GetListResponse, *client.APIResponse, error)

	// DeleteList deletes a single list by ID
	DeleteList(ctx context.Context, id string) (*client.APIResponse, error)

	// CreateListGroup associates a list with a group (POST /listGroups)
	CreateListGroup(ctx context.Context, req CreateListGroupRequest) (CreateListGroupResponse, *client.APIResponse, error)
}

type service struct {
	client client.Doer
}
