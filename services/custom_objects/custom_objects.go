package custom_objects

import (
	"context"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

type service struct {
	client client.Doer
}

// CustomObjectsService defines behaviour for working with ActiveCampaign custom
// object schemas (types) and their records. Methods return a generic
// interface{} containing the parsed JSON response; callers may type-assert or
// unmarshal into concrete types if desired.
type CustomObjectsService interface {
	// Schemas
	ListObjectTypes(ctx context.Context, opts map[string]string) (*ListSchemasResponse, *client.APIResponse, error)
	CreateObjectType(ctx context.Context, req *CreateObjectTypeRequest) (*Schema, *client.APIResponse, error)
	GetObjectType(ctx context.Context, id string) (*Schema, *client.APIResponse, error)
	DeleteObjectType(ctx context.Context, id string) (*client.APIResponse, error)

	// Records
	ListObjectRecords(ctx context.Context, objectTypeID string, opts map[string]string) (*ListRecordsResponse, *client.APIResponse, error)
	GetObjectRecord(ctx context.Context, objectTypeID, recordID string) (*Record, *client.APIResponse, error)
	CreateObjectRecord(ctx context.Context, objectTypeID string, req *CreateRecordRequest) (*CreateRecordResponse, *client.APIResponse, error)
	UpdateObjectRecord(ctx context.Context, objectTypeID, recordID string, req *UpdateRecordRequest) (*UpdateRecordResponse, *client.APIResponse, error)
	DeleteObjectRecord(ctx context.Context, objectTypeID, recordID string) (*client.APIResponse, error)
}
